package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	cl "github.com/green-api/maxbot-api-client-go/pkg/client"
	m "github.com/green-api/maxbot-api-client-go/pkg/models"
)

type helpers struct {
	client *cl.Client
}

/*
SendFile is a helper that simplifies sending files to a chat
It automatically determines whether the provided FileSource is a direct URL or a local file path

Example:

Sending a file via URL:

	msg, err := api.Helpers.SendFile(ctx, m.SendFileReq{
		ChatID:     123456789,
		Text:       "Check out this image!",
		FileSource: "https://example.com/image.png",
	})

Sending a local file:

	msg, err := api.Helpers.SendFile(ctx, m.SendFileReq{
		ChatID:     123456789,
		Text:       "Here is the report.",
		FileSource: "/local/path/to/report.pdf",
	})
*/
func (h *helpers) SendFile(ctx context.Context, req m.SendFileReq) (*m.Message, error) {
	if isURL(req.FileSource) {
		return h.SendFileByUrl(ctx, req)
	}
	return h.SendFileByUpload(ctx, req)
}

/*
SendFileByUrl sends a file, image, audio, or video directly via an external URL

Example:

	msg, err := api.Helpers.SendFileByUrl(ctx, m.SendFileReq{
		ChatID:     123456789,
		Text:       "Look at this web image!",
		FileSource: "https://example.com/photo.jpg?width=500&height=500",
	})
*/
func (h *helpers) SendFileByUrl(ctx context.Context, req m.SendFileReq) (*m.Message, error) {
	extInUrl := getExtension(req.FileSource)
	if determineUploadType(extInUrl) == m.UploadImage {
		attachment := m.Attachment{
			Type:    m.AttachmentImage,
			Payload: m.PhotoAttachmentPayload{URL: req.FileSource},
		}
		return h.sendFile(ctx, req, attachment)
	}

	tempPath, realExt, err := downloadTempFile(req.FileSource)

	uploadType := determineUploadType(realExt)

	uploader := &uploads{client: h.client}
	uploadResp, err := uploader.UploadFile(ctx, m.UploadFileReq{
		Type:     uploadType,
		FilePath: tempPath,
	})
	if err != nil {
		return nil, err
	}

	attachment := buildAttachmentFromToken(uploadType, uploadResp.Token, filepath.Base(req.FileSource))
	return h.sendFile(ctx, req, attachment)
}

/*
SendFileByUpload handles the two-step process of uploading a local file to MAX Bot servers
and then sending it as an attachment in a message

Example:

	msg, err := api.Helpers.SendFileByUpload(ctx, m.SendFileReq{
		ChatID:     987654321,
		Text:       "Here is the video you requested.",
		FileSource: "/Users/admin/Downloads/presentation.mp4",
	})
*/
func (h *helpers) SendFileByUpload(ctx context.Context, req m.SendFileReq) (*m.Message, error) {
	ext := getExtension(req.FileSource)
	uploadType := determineUploadType(ext)

	uploader := &uploads{client: h.client}
	uploadResp, err := uploader.UploadFile(ctx, m.UploadFileReq{
		Type:     uploadType,
		FilePath: req.FileSource,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to auto-upload file: %w", err)
	}

	attachment := buildAttachmentFromToken(uploadType, uploadResp.Token, filepath.Base(req.FileSource))
	return h.sendFile(ctx, req, attachment)
}

/*
sendFile is an internal helper that constructs the final message payload and executes the API request
*/
func (h *helpers) sendFile(ctx context.Context, req m.SendFileReq, attachment m.Attachment) (*m.Message, error) {
	allAttachments := append(req.Attachments, attachment)

	request := m.SendMessageReq{
		UserID:             req.UserID,
		ChatID:             req.ChatID,
		Text:               req.Text,
		Format:             req.Format,
		Attachments:        allAttachments,
		Notify:             req.Notify,
		Link:               req.Link,
		DisableLinkPreview: req.DisableLinkPreview,
	}

	var lastErr error
	for i := 0; i < 3; i++ {
		msg, err := decode[m.Message](
			ctx, h.client, http.MethodPost, m.PathMessages,
			cl.WithQuery(request),
			cl.WithBody(request),
		)

		if err == nil {
			return msg, nil
		}

		lastErr = err
		if strings.Contains(err.Error(), "not.ready") {
			time.Sleep(3 * time.Second)
			continue
		}

		break
	}

	return nil, lastErr
}

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func getExtension(source string) string {
	if isURL(source) {
		if u, err := url.Parse(source); err == nil {
			return strings.ToLower(filepath.Ext(u.Path))
		}
	}
	return strings.ToLower(filepath.Ext(source))
}

func determineUploadType(ext string) m.UploadType {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".webp":
		return m.UploadImage
	case ".mp4", ".avi", ".mov":
		return m.UploadVideo
	case ".mp3", ".ogg", ".wav":
		return m.UploadAudio
	default:
		return m.UploadFile
	}
}

func buildAttachmentFromToken(uType m.UploadType, token, filename string) m.Attachment {
	switch uType {
	case m.UploadImage:
		return m.AttachImage(token, "")
	case m.UploadVideo:
		return m.AttachVideo(token, "")
	case m.UploadAudio:
		return m.AttachAudio(token, "")
	default:
		return m.AttachFile(token, "", filename)
	}
}

func downloadTempFile(urlStr string) (string, string, error) {
	resp, err := http.DefaultClient.Get(urlStr)
	if err != nil {
		return "", "", fmt.Errorf("failed to create GET request: %w", err)
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	contentType = strings.Split(contentType, ";")[0]

	ext, ok := mimeToExt[contentType]

	if !ok || ext == "" {
		ext = getExtension(urlStr)
	}

	if ext == "" {
		ext = ".bin"
	}

	tempFile, err := os.CreateTemp("", "maxbot_*"+ext)
	if err != nil {
		return "", "", fmt.Errorf("failed to create temp file: %w", err)
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, resp.Body)
	return tempFile.Name(), ext, err
}

var mimeToExt = map[string]string{
	"image/jpeg":      ".jpg",
	"image/png":       ".png",
	"image/webp":      ".webp",
	"video/mp4":       ".mp4",
	"audio/mpeg":      ".mp3",
	"audio/ogg":       ".ogg",
	"audio/wav":       ".wav",
	"application/pdf": ".pdf",
}
