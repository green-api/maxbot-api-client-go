package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	cl "github.com/green-api/maxbot-api-client-go/pkg/client"
	m "github.com/green-api/maxbot-api-client-go/pkg/models"
)

type uploads struct {
	client *cl.Client
}

/*
UploadFile uploads a file to the server and returns the upload metadata.
It seamlessly handles both STEP 1 (obtaining the URL) and STEP 2 (streaming the file).

Example:

	resp, err := api.Uploads.UploadFile(ctx, m.UploadFileReq{
		Type:     m.UploadImage,
		FilePath: "/path/to/image.png",
	})
*/
func (u *uploads) UploadFile(ctx context.Context, req m.UploadFileReq) (*m.UploadedInfo, error) {
	initResp, err := u.getUploadURL(ctx, req.Type)
	if err != nil {
		return nil, err
	}

	multipartResp, err := u.uploadMultipart(ctx, initResp.Url, req.FilePath)

	if err == nil && multipartResp != nil && multipartResp.Token != "" {
		return multipartResp, nil
	}

	if initResp.Token != "" {
		return &m.UploadedInfo{Token: initResp.Token}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("upload failed and no token found: %w", err)
	}

	return nil, fmt.Errorf("server did not return token after upload")
}

// STEP 1: Obtaining the target URL
func (u *uploads) getUploadURL(ctx context.Context, uploadType m.UploadType) (*m.PhotoAttachmentRequestPayload, error) {
	return decode[m.PhotoAttachmentRequestPayload](
		ctx,
		u.client,
		http.MethodPost,
		m.PathUploads,
		cl.WithQuery(m.UploadTypeReq{Type: uploadType}),
	)
}

// STEP 2: Streaming file transfer
func (u *uploads) uploadMultipart(ctx context.Context, uploadURL string, filePath string) (*m.UploadedInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	return decode[m.UploadedInfo](
		ctx,
		u.client,
		http.MethodPost,
		uploadURL,
		cl.WithMultipart(filepath.Base(filePath), file),
	)
}
