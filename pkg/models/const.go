package models

const DefaultBaseURL = "https://platform-api.max.ru"

const (
	PathMe            = "me"
	PathChats         = "chats"
	PathAnswers       = "answers"
	PathUpdates       = "updates"
	PathUploads       = "uploads"
	PathMessages      = "messages"
	PathSubscriptions = "subscriptions"
	PathVideos        = "videos"

	FormatPathChatsID           = "chats/%d"
	FormatPathChatsActions      = "chats/%d/actions"
	FormatPathChatsMembers      = "chats/%d/members"
	FormatPathChatsMembersMe    = "chats/%d/members/me"
	FormatPathChatsMembersAdmin = "chats/%d/members/admins"
)

/* Type of attachments in message */
type AttachmentType string

/* List of AttachmentType */
const (
	AttachmentImage    AttachmentType = "image"
	AttachmentVideo    AttachmentType = "video"
	AttachmentAudio    AttachmentType = "audio"
	AttachmentFile     AttachmentType = "file"
	AttachmentSticker  AttachmentType = "sticker"
	AttachmentContact  AttachmentType = "contact"
	AttachmentKeyboard AttachmentType = "inline_keyboard"
	AttachmentShare    AttachmentType = "share"
	AttachmentLocation AttachmentType = "location"
)

/* ChatType : Type of chat. Dialog (one-on-one), chat or channel */
type ChatType string

/* List of ChatType */
const (
	CHANNEL ChatType = "channel"
	CHAT    ChatType = "chat"
	DIALOG  ChatType = "dialog"
)

/* ChatStatus : Chat status for current bots */
type ChatStatus string

/* List of ChatStatus */
const (
	ACTIVE  ChatStatus = "active"
	CLOSED  ChatStatus = "closed"
	LEFT    ChatStatus = "left"
	REMOVED ChatStatus = "removed"
)

/* MarkupType : Type of markups */
type MarkupType string

/* List of MarkupType */
const (
	MarkupStrong        MarkupType = "strong"
	MarkupEmphasized    MarkupType = "emphasized"
	MarkupMonospaced    MarkupType = "monospaced"
	MarkupLink          MarkupType = "link"
	MarkupStrikethrough MarkupType = "strikethrough"
	MarkupUnderline     MarkupType = "underline"
	MarkupUser          MarkupType = "user_mention"
)

/* Format types */
type Format string

/* List of Format */
const (
	HTML     Format = "html"
	Markdown Format = "markdown"
)

/* Types of update events */
type UpdateType string

/* List of UpdateType */
const (
	TypeBotAdded         UpdateType = "bot_added"
	TypeBotRemoved       UpdateType = "bot_removed"
	TypeBotStarted       UpdateType = "bot_started"
	TypeBotStopped       UpdateType = "bot_stopped"
	TypeChatTitleChanged UpdateType = "chat_title_changed"
	TypeDialogMuted      UpdateType = "dialog_muted"
	TypeDialogUnmuted    UpdateType = "dialog_unmuted"
	TypeDialogCleared    UpdateType = "dialog_cleared"
	TypeDialogRemoved    UpdateType = "dialog_removed"
	TypeMessageCreated   UpdateType = "message_created"
	TypeMessageCallback  UpdateType = "message_callback"
	TypeMessageEdited    UpdateType = "message_edited"
	TypeMessageRemoved   UpdateType = "message_removed"
	TypeUserAdded        UpdateType = "user_added"
	TypeUserRemoved      UpdateType = "user_removed"
)

/* Type of linked message */
type LinkedMessageType string

/* List of LinkedMessageType */
const (
	FORWARD LinkedMessageType = "forward"
	REPLY   LinkedMessageType = "reply"
)

/* SenderAction : Different actions to send to chat members */
type SenderAction string

/* List of SenderAction */
const (
	TYPING_ON     SenderAction = "typing_on"
	TYPING_OFF    SenderAction = "typing_off"
	SENDING_PHOTO SenderAction = "sending_photo"
	SENDING_VIDEO SenderAction = "sending_video"
	SENDING_AUDIO SenderAction = "sending_audio"
	SENDING_FILE  SenderAction = "sending_file"
	MARK_SEEN     SenderAction = "mark_seen"
)

/* Admin rights granted for user */
type ChatAdminPermission string

/* List of ChatAdminPermission */
const (
	PermissionReadAllMessages ChatAdminPermission = "read_all_messages"
	PermissionAddRemoveUsers  ChatAdminPermission = "add_remove_members"
	PermissionAddAdmins       ChatAdminPermission = "add_admins"
	PermissionChangeChatPhoto ChatAdminPermission = "change_chat_info"
	PermissionPinMessage      ChatAdminPermission = "pin_message"
	PermissionWrite           ChatAdminPermission = "write"
	PermissionEditLink        ChatAdminPermission = "edit_link"
)

/* Types of buttons */
type ButtonType string

/* List of ButtonType */
const (
	ButtonCallback        ButtonType = "callback"
	ButtonLink            ButtonType = "link"
	ButtonRequestLocation ButtonType = "request_geo_location"
	ButtonRequestContact  ButtonType = "request_contact"
	ButtonOpenApp         ButtonType = "open_app"
	ButtonMessage         ButtonType = "message"
)

/* Types of files for uploading */
type UploadType string

/* List of UploadType */
const (
	UploadImage UploadType = "image"
	UploadVideo UploadType = "video"
	UploadAudio UploadType = "audio"
	UploadFile  UploadType = "file"
)
