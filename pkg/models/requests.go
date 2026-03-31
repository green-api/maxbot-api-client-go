package models

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

/*
ToValues converts a struct to url.Values by inspecting its "url" struct tags
It supports omitempty and handles slices by adding multiple values for the same key
*/
func ToValues(v any) url.Values {
	values := url.Values{}
	val := reflect.Indirect(reflect.ValueOf(v))
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("url")
		if tag == "" || tag == "-" {
			continue
		}

		parts := strings.Split(tag, ",")
		key := parts[0]
		omitempty := false
		for _, p := range parts[1:] {
			if p == "omitempty" {
				omitempty = true
				break
			}
		}

		fieldVal := val.Field(i)

		if omitempty && fieldVal.IsZero() {
			continue
		}

		switch fieldVal.Kind() {
		case reflect.Slice:
			for j := 0; j < fieldVal.Len(); j++ {
				values.Add(key, fmt.Sprintf("%v", fieldVal.Index(j).Interface()))
			}
		default:
			values.Set(key, fmt.Sprintf("%v", fieldVal.Interface()))
		}
	}
	return values
}

/* GetChatsReq represents a request to retrieve a paginated list of chats */
type GetChatsReq struct {
	Count  int   `url:"count,omitempty" json:"-"`
	Marker int64 `url:"marker,omitempty" json:"-"`
}

/* GetChatsResp contains a list of chats and a marker for pagination */
type GetChatsResp struct {
	Chats  []Chat `json:"chats"`
	Marker int64  `json:"marker,omitempty"`
}

/* GetChatReq represents a request to retrieve information about a specific chat */
type GetChatReq struct {
	ChatID int64 `url:"chatId" json:"-"`
}

/* EditChatReq represents a request to modify chat metadata (e.g., title, icon) */
type EditChatReq struct {
	ChatID int64  `url:"chatId" json:"-"`
	Icon   *Image `json:"icon,omitempty"`
	Title  string `json:"title,omitempty"`
	Pin    string `json:"pin,omitempty"`
	Notify bool   `json:"notify,omitempty"`
}

/* DeleteChatReq represents a request to delete a chat */
type DeleteChatReq struct {
	ChatID int64 `url:"chatId" json:"-"`
}

/* SendActionReq represents a request to broadcast a typing/uploading status in a chat */
type SendActionReq struct {
	ChatID int64        `url:"chatId" json:"-"`
	Action SenderAction `json:"action"`
}

/* PinMessageReq represents a request to pin a specific message in a chat */
type PinMessageReq struct {
	ChatID    int64  `url:"chatId" json:"-"`
	MessageID string `json:"message_id"`
	Notify    bool   `json:"notify,omitempty"`
}

/* UnpinMessageReq represents a request to remove the pinned message in a chat */
type UnpinMessageReq struct {
	ChatID int64 `url:"chatId" json:"-"`
}

/* GetPinnedMessageReq represents a request to retrieve the currently pinned message */
type GetPinnedMessageReq struct {
	ChatID int64 `url:"chatId" json:"-"`
}

/* GetChatMembershipReq represents a request to check membership status in a chat */
type GetChatMembershipReq struct {
	ChatID int64 `url:"chatId" json:"-"`
}

/* LeaveChatReq represents a request for the current user to leave a chat */
type LeaveChatReq struct {
	ChatID int64 `url:"chatId" json:"-"`
}

/* GetChatAdminsReq represents a request to get a list of administrators for a chat */
type GetChatAdminsReq struct {
	ChatID int64 `url:"chatId" json:"-"`
}

/* GetChatAdminsResp contains the list of chat administrators */
type GetChatAdminsResp struct {
	Members []ChatMember `json:"members"`
	Marker  int64        `json:"marker,omitempty"`
}

/* SetChatAdminsReq represents a request to update or set administrators for a chat */
type SetChatAdminsReq struct {
	ChatID int64       `url:"chatId" json:"-"`
	Admins []ChatAdmin `json:"admins"`
	Marker int64       `json:"marker,omitempty"`
}

/* DeleteAdminReq represents a request to revoke administrator privileges from a user */
type DeleteAdminReq struct {
	ChatID int64 `url:"chatId" json:"-"`
	UserID int64 `url:"userId" json:"-"`
}

/* GetChatMembersReq represents a request to fetch members of a chat, optionally filtered by IDs */
type GetChatMembersReq struct {
	ChatID  int64 `url:"chatId" json:"-"`
	UserIDs []int `url:"user_ids,omitempty" json:"-"`
	Marker  int64 `url:"marker,omitempty" json:"-"`
	Count   int   `url:"count,omitempty" json:"-"` /* default: 20 */
}

/* AddMembersReq represents a request to add users to a chat */
type AddMembersReq struct {
	ChatID  int64 `url:"chatId" json:"-"`
	UserIDs []int `json:"user_ids,omitempty"`
}

/* AddMembersResp contains the result of an add members operation, including any failures */
type AddMembersResp struct {
	SimpleQueryResult
	FailedUserIDs     []int               `url:"failed_user_ids,omitempty" json:"-"`
	FailedUserDetails []FailedUserDetails `url:"failed_user_details,omitempty" json:"-"`
}

/* DeleteMemberReq represents a request to remove or block a user from a chat */
type DeleteMemberReq struct {
	ChatID int64 `url:"chatId" json:"-"`
	UserID int64 `url:"userId" json:"-"`
	Block  bool  `url:"block,omitempty" json:"-"`
}

/* GetSubscriptionsResp contains a list of active webhook subscriptions */
type GetSubscriptionsResp struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

/* SubscribeReq represents a request to register a new webhook URL for updates */
type SubscribeReq struct {
	Url         string       `json:"url"`
	UpdateTypes []UpdateType `json:"update_types,omitempty"`
	Secret      string       `json:"secret,omitempty"`
}

/* UnsubscribeReq represents a request to remove a previously registered webhook */
type UnsubscribeReq struct {
	Url string `url:"url" json:"-"`
}

/* GetUpdatesReq represents a request to long-poll for recent events or messages */
type GetUpdatesReq struct {
	Limit   int          `url:"limit,omitempty" json:"-"`
	Timeout int          `url:"timeout,omitempty" json:"-"`
	Marker  int          `url:"marker,omitempty" json:"-"`
	Types   []UpdateType `url:"types,omitempty" json:"-"`
}

/* GetUpdatesResp contains the list of fetched events/updates */
type GetUpdatesResp struct {
	Updates []Update `json:"updates"`
	Marker  int64    `json:"marker"`
}

/* UploadFileReq represents a request to upload a file to a specific URL */
type UploadFileReq struct {
	Type      UploadType `url:"type" json:"-"`
	UploadURL string
	FilePath  string
}

/* UploadTypeReq represents a request to get an upload URL for a specific file type */
type UploadTypeReq struct {
	Type UploadType `url:"type" json:"-"`
}

/* UploadFileMultipartReq represents a multipart file upload payload */
type UploadFileMultipartReq struct {
	UploadURL string
	FilePath  string
}

/*
UploadedInfo contains metadata and details about an uploaded audio or video file
*/
type UploadedInfo struct {
	FileID int64  `json:"file_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

/* PhotoAttachmentRequestPayload represents the payload to attach an uploaded photo */
type PhotoAttachmentRequestPayload struct {
	Url    string               `json:"url"`
	Token  string               `json:"token,omitempty"`
	Photos map[string]PhotoData `json:"photos,omitempty"`
}

/* GetMessagesReq represents a request to fetch messages from a chat */
type GetMessagesReq struct {
	ChatID     int64    `url:"chat_id,omitempty" json:"-"`
	MessageIDs []string `url:"message_ids,omitempty" json:"-"`
	From       int64    `url:"from,omitempty" json:"-"`
	To         int64    `url:"to,omitempty" json:"-"`
	Count      int      `url:"count,omitempty" json:"-"`
}

/* SendMessageReq represents a request to send a new message to a chat or user */
type SendMessageReq struct {
	UserID             int64           `url:"user_id,omitempty" json:"-"`
	ChatID             int64           `url:"chat_id,omitempty" json:"-"`
	Text               string          `json:"text,omitempty"`
	Format             Format          `json:"format,omitempty"`
	Notify             bool            `json:"notify,omitempty"`
	Attachments        []Attachment    `json:"attachments,omitempty"`
	Link               *NewMessageLink `json:"link,omitempty"`
	DisableLinkPreview bool            `url:"disable_link_preview,omitempty" json:"-"`
}

/* SendMessageResp contains the resulting message object after successfully sending it */
type SendMessageResp struct {
	Message Message `json:"message"`
}

/* EditMessageReq represents a request to edit the content of an existing message */
type EditMessageReq struct {
	MessageID   string          `url:"message_id" json:"-"`
	Text        string          `json:"text,omitempty"`
	Attachments []Attachment    `json:"attachments,omitempty"`
	Link        *NewMessageLink `json:"link,omitempty"`
	Notify      bool            `json:"notify,omitempty"`
	Format      Format          `json:"format,omitempty"`
}

/* DeleteMessageReq represents a request to delete a specific message */
type DeleteMessageReq struct {
	MessageID string `url:"message_id" json:"-"`
}

/* GetMessageReq represents a request to fetch a specific message by its ID */
type GetMessageReq struct {
	MessageID string `url:"message_id" json:"-"`
}

/* GetVideoInfoReq represents a request to retrieve details about an uploaded video */
type GetVideoInfoReq struct {
	VideoToken string `url:"video_token" json:"-"`
}

/* GetVideoInfoResp contains the metadata and available playback URLs for a video */
type GetVideoInfoResp struct {
	Token     string                 `json:"token"`
	Urls      VideoUrls              `json:"urls,omitempty"`
	Thumbnail PhotoAttachmentPayload `json:"thumbnail,omitempty"`
	Width     int                    `json:"width,omitempty"`
	Height    int                    `json:"height,omitempty"`
	Duration  int                    `json:"duration,omitempty"`
}

/* AnswerCallbackReq represents a response payload sent back after a user clicks an inline button */
type AnswerCallbackReq struct {
	CallbackID   string          `url:"callback_id" json:"-"`
	Message      *NewMessageBody `json:"message,omitempty"`
	Notification string          `json:"notification,omitempty"`
}

/* SimpleQueryResult represents a basic boolean result indicating success or failure of an operation */
type SimpleQueryResult struct {
	Success bool   `url:"success" json:"-"`
	Message string `url:"message,omitempty" json:"-"`
}

type SendFileReq struct {
	UserID             int64           `url:"user_id,omitempty" json:"-"`
	ChatID             int64           `url:"chat_id,omitempty" json:"-"`
	Text               string          `json:"text,omitempty"`
	Format             Format          `json:"format,omitempty"`
	Notify             bool            `json:"notify,omitempty"`
	FileSource         string          `json:"file_source"`
	Link               *NewMessageLink `json:"link,omitempty"`
	DisableLinkPreview bool            `url:"disable_link_preview,omitempty" json:"-"`
	Attachments        []Attachment    `json:"attachments,omitempty"`
}
