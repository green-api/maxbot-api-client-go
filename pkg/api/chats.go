package api

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	cl "github.com/green-api/maxbot-api-client-go/pkg/client"
	"github.com/green-api/maxbot-api-client-go/pkg/models"
)

type chats struct {
	client *cl.Client
}

/*
GetChats returns information about chats that bot participated in: a result list and marker points to the next page

Example:

	chats, err := api.Chats.GetChats(ctx, &models.GetChatsReq{
		Count: 20,
	})
*/
func (api *chats) GetChats(ctx context.Context, req *models.GetChatsReq) (*models.GetChatsResp, error) {
	return decode[models.GetChatsResp](ctx, api.client, http.MethodGet, models.PathChats, cl.WithQuery(req))
}

/*
GetChat returns info about chat

Example:

	info, err := api.Chats.GetChat(ctx, &models.GetChatReq{
		ChatID: 123456789,
	})
*/
func (api *chats) GetChat(ctx context.Context, req *models.GetChatReq) (*models.ChatInfo, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10))
	return decode[models.ChatInfo](ctx, api.client, http.MethodGet, path)
}

/*
EditChat modifies the properties of a chat, such as its title, icon, or notification settings

Example:

	info, err := api.Chats.EditChat(ctx, &models.EditChatReq{
		ChatID: 123456789,
		Title:  "Updated Chat Title",
		Notify: true,
	})
*/
func (api *chats) EditChat(ctx context.Context, req *models.EditChatReq) (*models.ChatInfo, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10))
	return decode[models.ChatInfo](ctx, api.client, http.MethodPatch, path, cl.WithBody(req))
}

/*
DeleteChat permanently deletes a chat for the bot

Example :

	response, err:= api.Chats.DeleteChat(ctx, &models.DeleteChatReq{
		ChatID: 123456789,
	})
*/
func (api *chats) DeleteChat(ctx context.Context, req *models.DeleteChatReq) (*models.SimpleQueryResult, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10))
	return decode[models.SimpleQueryResult](ctx, api.client, http.MethodDelete, path)
}

/*
SendAction broadcasts a temporary status action (e.g., "typing...", "recording video...") to the chat participants

Example:

	response, err:= api.Chats.SendAction(ctx, &models.SendActionReq{
		ChatID: 123456789,
		Action: "mark_seen",
	})
*/
func (api *chats) SendAction(ctx context.Context, req *models.SendActionReq) (*models.SimpleQueryResult, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10), "actions")
	return decode[models.SimpleQueryResult](ctx, api.client, http.MethodPost, path, cl.WithBody(req))
}

/*
GetPinnedMessage retrieves the currently pinned message in the specified chat

Example:

	msg, err := api.Chats.GetPinnedMessage(ctx, &models.GetPinnedMessageReq{
		ChatID: 123456789,
	})
*/
func (api *chats) GetPinnedMessage(ctx context.Context, req *models.GetPinnedMessageReq) (*models.Message, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10), "pin")
	return decode[models.Message](ctx, api.client, http.MethodGet, path)
}

/*
PinMessage pins a specific message in the chat
You can optionally specify whether to notify chat members about the new pinned message

Example:

	response, err:= api.Chats.PinMessage(ctx, &models.PinMessageReq{
		ChatID:    123456789,
		MessageID: "mid:987654321...",
		Notify:    true,
	})
*/
func (api *chats) PinMessage(ctx context.Context, req *models.PinMessageReq) (*models.SimpleQueryResult, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10), "pin")
	return decode[models.SimpleQueryResult](ctx, api.client, http.MethodPut, path, cl.WithBody(req))
}

/*
UnpinMessage removes the pinned message from the specified chat

Example:

	response, err:= api.Chats.UnpinMessage(ctx, &models.UnpinMessageReq{
		ChatID: 123456789,
	})
*/
func (api *chats) UnpinMessage(ctx context.Context, req *models.UnpinMessageReq) (*models.SimpleQueryResult, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10), "pin")
	return decode[models.SimpleQueryResult](ctx, api.client, http.MethodDelete, path)
}

/*
GetChatMembership returns chat membership info for the current bot

Example:

	memberInfo, err := api.Chats.GetChatMembership(ctx, &models.GetChatMembershipReq{
		ChatID: 123456789,
	})
*/
func (api *chats) GetChatMembership(ctx context.Context, req *models.GetChatMembershipReq) (*models.ChatMember, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10), "members", "me")
	return decode[models.ChatMember](ctx, api.client, http.MethodGet, path)
}

/*
LeaveChat removes bot from chat members

Example:

	response, err:= api.Chats.LeaveChat(ctx, &models.LeaveChatReq{
		ChatID: 123456789,
	})
*/
func (api *chats) LeaveChat(ctx context.Context, req *models.LeaveChatReq) (*models.SimpleQueryResult, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10), "members", "me")
	return decode[models.SimpleQueryResult](ctx, api.client, http.MethodDelete, path)
}

/*
GetChatAdmins retrieves a list of administrators for the specified group chat

Example:

	admins, err := api.Chats.GetChatAdmins(ctx, &models.GetChatAdminsReq{
		ChatID: 123456789,
	})
*/
func (api *chats) GetChatAdmins(ctx context.Context, req *models.GetChatAdminsReq) (*models.GetChatAdminsResp, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10), "members", "admins")
	return decode[models.GetChatAdminsResp](ctx, api.client, http.MethodGet, path)
}

/*
SetChatAdmins assigns administrator rights to specific users in a group chat

Example:

	response, err:= api.Chats.SetChatAdmins(ctx, &models.SetChatAdminsReq{
		ChatID: 123456789,
		Admins: []models.ChatAdmin{
			{
				UserID: 98765,
				Role: "admin",
			},
			{
				UserID: 43210,
				Role: "admin",
			},
		}
	})
*/
func (api *chats) SetChatAdmins(ctx context.Context, req *models.SetChatAdminsReq) (*models.SimpleQueryResult, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10), "/members/admins")
	return decode[models.SimpleQueryResult](ctx, api.client, http.MethodPost, path, cl.WithBody(req))
}

/*
DeleteAdmin revokes administrator rights from a specific user in a group chat

Example:

	response, err:= api.Chats.DeleteAdmin(ctx, &models.DeleteAdminReq{
		ChatID: 123456789,
		UserID: 98765,
	})
*/
func (api *chats) DeleteAdmin(ctx context.Context, req *models.DeleteAdminReq) (*models.SimpleQueryResult, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10), "/members/admins", strconv.FormatInt(req.UserID, 10))
	return decode[models.SimpleQueryResult](ctx, api.client, http.MethodDelete, path)
}

/*
GetChatMembers returns users participated in chat.

Example:

	participants, err := api.Chats.GetChatMembers(ctx, &models.GetChatMembersReq{
		ChatID: 123456789,
		Count:  20,
	})
*/
func (api *chats) GetChatMembers(ctx context.Context, req *models.GetChatMembersReq) (*models.GetChatAdminsResp, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10), "/members")
	return decode[models.GetChatAdminsResp](ctx, api.client, http.MethodGet, path, cl.WithQuery(req))
}

/*
AddMembers adds one or more users to a group chat

Example:

	response, err:= api.Chats.AddMembers(ctx, &models.AddMembersReq{
		ChatID:  123456789,
		UserIDs: []int{11111, 22222},
	})
*/
func (api *chats) AddMembers(ctx context.Context, req *models.AddMembersReq) (*models.AddMembersResp, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10), "/members")
	return decode[models.AddMembersResp](ctx, api.client, http.MethodPost, path, cl.WithBody(req))
}

/*
DeleteMember removes a specific user from a group chat
You can optionally block the user from rejoining

Example:

	response, err:= api.Chats.DeleteMember(ctx, &models.DeleteMemberReq{
		ChatID: 123456789,
		UserID: 98765,
		Block: true,
	})
*/
func (api *chats) DeleteMember(ctx context.Context, req *models.DeleteMemberReq) (*models.SimpleQueryResult, error) {
	path, _ := url.JoinPath(models.PathChats, strconv.FormatInt(req.ChatID, 10), "/members")
	return decode[models.SimpleQueryResult](ctx, api.client, http.MethodDelete, path, cl.WithQuery(req))
}
