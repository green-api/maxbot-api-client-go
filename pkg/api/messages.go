package api

import (
	"context"
	"net/http"
	"net/url"

	cl "github.com/green-api/maxbot-api-client-go/pkg/client"
	m "github.com/green-api/maxbot-api-client-go/pkg/models"
)

type messages struct {
	client *cl.Client
}

/*
GetMessages retrieves a list of messages
It can fetch messages belonging to a specific ChatID
or by an exact list of MessageIDs.

Example:

	resp, err := api.Messages.GetMessages(ctx, GetMessagesReq{
		ChatID: 123456,
			Alternatively, fetch specific messages:
			MessageIDs: []string{"mid:1", "mid:2"},
	})
*/
func (api *messages) GetMessages(ctx context.Context, req m.GetMessagesReq) (*m.MessagesList, error) {
	return decode[m.MessagesList](ctx, api.client, http.MethodGet, m.PathMessages, cl.WithQuery(req))
}

/*
SendMessage sends a text message or attachment to a specific user or chat
If Notify is false, no push notification will be sent to the user

Example:

	resp, err := api.Messages.SendMessage(ctx, SendMessageReq{
		ChatID: 123456,
		Text:   "Hello, world!",
		Notify: true,
	})
*/
func (api *messages) SendMessage(ctx context.Context, req m.SendMessageReq) (*m.SendMessageResp, error) {
	return decode[m.SendMessageResp](ctx, api.client, http.MethodPost, m.PathMessages, cl.WithQuery(req), cl.WithBody(req))
}

/*
EditMessage modifies the content of a previously sent message

Example:

	msg, err := api.Messages.EditMessage(ctx, EditMessageReq{
		MessageID: "mid:987654321...",
		Text:      "Updated message text!",
		Format:    HTML,
	})
*/
func (api *messages) EditMessage(ctx context.Context, req m.EditMessageReq) (*m.SimpleQueryResult, error) {
	return decode[m.SimpleQueryResult](ctx, api.client, http.MethodPut, m.PathMessages, cl.WithQuery(req), cl.WithBody(req))
}

/*
DeleteMessage removes a previously sent message by its ID

Example:

	success, err := api.Messages.DeleteMessage(ctx, DeleteMessageReq{
		MessageID: "mid:987654321...",
	})
*/
func (api *messages) DeleteMessage(ctx context.Context, req m.DeleteMessageReq) (*m.SimpleQueryResult, error) {
	return decode[m.SimpleQueryResult](ctx, api.client, http.MethodDelete, m.PathMessages, cl.WithQuery(req))
}

/*
GetMessage retrieves full information about a message by its ID

Example:
	msg, err := api.Messages.GetMessage(ctx, GetMessageReq{
		MessageID: "mid:987654321...",
	})
*/

func (api *messages) GetMessage(ctx context.Context, req m.GetMessageReq) (*m.Message, error) {
	path, _ := url.JoinPath(m.PathMessages, req.MessageID)
	return decode[m.Message](ctx, api.client, http.MethodGet, path)
}

/*
GetVideoInfo retrieves metadata and the processing status for an uploaded video

Example:

	info, err := api.Messages.GetVideoInfo(ctx, GetVideoInfoReq{
		VideoToken: "vtok_abc123xyz...",
	})
*/
func (api *messages) GetVideoInfo(ctx context.Context, req m.GetVideoInfoReq) (*m.GetVideoInfoResp, error) {
	path, _ := url.JoinPath(m.PathVideos, req.VideoToken)
	return decode[m.GetVideoInfoResp](ctx, api.client, http.MethodGet, path)
}

/*
AnswerCallback acknowledges and responds to a user clicking an inline button

Example:

	success, err := api.Messages.AnswerCallback(ctx, AnswerCallbackReq{
		CallbackID: "cbk_12345...",
		Message: &NewMessageBody{
			Text: "Action confirmed!",
		},
	})
*/
func (api *messages) AnswerCallback(ctx context.Context, req m.AnswerCallbackReq) (*m.SimpleQueryResult, error) {
	return decode[m.SimpleQueryResult](ctx, api.client, http.MethodPost, m.PathAnswers, cl.WithQuery(req), cl.WithBody(&req))
}
