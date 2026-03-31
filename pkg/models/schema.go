package models

import "fmt"

/* APIError represents an error response returned by the API */
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

/* Attachment represents a generic schema for message attachment */
type Attachment struct {
	Type      AttachmentType `json:"type"`
	Payload   any            `json:"payload,omitempty"`
	Latitude  *float64       `json:"latitude,omitempty"`
	Longitude *float64       `json:"longitude,omitempty"`
}

/* ImagePayload represents a collection of photo data indexed by key */
type ImagePayload struct {
	Photos map[string]PhotoData `json:"photos"`
}

/* PhotoData contains a unique token for a specific photo version */
type PhotoData struct {
	Token string `json:"token"`
}

/* AttachmentRequest represents a request to attach data to a message */
type AttachmentRequest struct {
	Type AttachmentType `json:"type"`
}

/* PhotoAttachmentPayload represents the payload for an "image" attachment */
type PhotoAttachmentPayload struct {
	PhotoID int64  `json:"photo_id,omitempty"`
	Token   string `json:"token,omitempty"`
	URL     string `json:"url,omitempty"`
}

/* AttachImage creates an image attachment using a pre-uploaded file token

Example:
	attachment := m.AttachImage("v2_01GHS...")

	msg := m.SendMessageReq{
		Attachments: []m.Attachment{attachment},
	}
*/
func AttachImage(token, url string) Attachment {
	return Attachment{
		Type: AttachmentImage,
		Payload: PhotoAttachmentPayload{
			Token: token,
			URL:   url,
		},
	}
}

/* MediaPayload contains the upload token and URL for a media file */
type MediaPayload struct {
	URL   string `json:"url,omitempty"`
	Token string `json:"token,omitempty"`
}

/* VideoAttachmentPayload represents the payload for a "video" attachment */
type VideoAttachmentPayload struct {
	MediaPayload
	Thumbnail *string `json:"thumbnail,omitempty"`
	Width     *int    `json:"width,omitempty"`
	Height    *int    `json:"height,omitempty"`
	Duration  *int    `json:"duration,omitempty"`
}

/*\
AttachVideo creates a video attachment. Requires a file token or a direct URL

Example:
	attachment := m.AttachVideo("vid_token_123", "https:/*cdn.max.ru/v/123.mp4")

	msg := m.SendMessageReq{
		Attachments: []m.Attachment{attachment},
	}
*/
func AttachVideo(token, url string) Attachment {
	return Attachment{
		Type: AttachmentVideo,
		Payload: VideoAttachmentPayload{
			MediaPayload: MediaPayload{Token: token, URL: url},
		},
	}
}

/* AudioAttachmentPayload represents the payload for an "audio" attachment */
type AudioAttachmentPayload struct {
	MediaPayload
	Transcription *string `json:"transcription,omitempty"`
}

/*
AttachAudio creates an audio attachment, typically used for voice messages or music files

Example:
	attachment := m.AttachAudio("audio_token", "https:/*cdn.max.ru/a/audio.mp3")

	msg := m.SendMessageReq{
		Attachments: []m.Attachment{attachment},
	}
*/
func AttachAudio(token, url string) Attachment {
	return Attachment{
		Type: AttachmentAudio,
		Payload: AudioAttachmentPayload{
			MediaPayload: MediaPayload{Token: token, URL: url},
		},
	}
}

/* FileAttachmentPayload represents the payload for a "file" attachment */
type FileAttachmentPayload struct {
	MediaPayload
	Filename string `json:"filename,omitempty"`
	Size     int64  `json:"size,omitempty"`
}

/*
AttachFile creates a general file attachment (document, archive, etc.)

Example:
	attachment := m.AttachFile("doc_token", "https:/*cdn.max.ru/f/report.pdf", "Monthly_Report.pdf")

	msg := m.SendMessageReq{
		Attachments: []m.Attachment{attachment},
	}
*/
func AttachFile(token, url, filename string) Attachment {
	return Attachment{
		Type: AttachmentFile,
		Payload: FileAttachmentPayload{
			MediaPayload: MediaPayload{Token: token, URL: url},
			Filename:     filename,
		},
	}
}

/* StickerData contains the basic information needed to render or identify a sticker */
type StickerData struct {
	URL  string `json:"url,omitempty"`
	Code string `json:"code,omitempty"`
}

/* StickerAttachmentPayload represents the payload for a "sticker" attachment */
type StickerAttachmentPayload struct {
	StickerData
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

/* AttachSticker creates a sticker attachment using its unique code or URL

Example:
	attachment := m.AttachSticker("https:/*cdn.max.ru/s/smile.webp", "smile_code")

	msg := m.SendMessageReq{
		Attachments: []m.Attachment{attachment},
	}
*/
func AttachSticker(url, code string) Attachment {
	return Attachment{
		Type: AttachmentSticker,
		Payload: StickerAttachmentPayload{
			StickerData: StickerData{URL: url, Code: code},
		},
	}
}

/* ContactAttachmentPayload represents the payload for a "contact" attachment */
type ContactAttachmentPayload struct {
	Name      *string `json:"name,omitempty"`
	ContactID *int64  `json:"contact_id,omitempty"`
	VCFInfo   *string `json:"vcf_info,omitempty"`
	VCFPhone  *string `json:"vcf_phone,omitempty"`
}

/* AttachContact creates a contact card attachment using VCF (vCard) format

Example:
	attachment := m.AttachContact("John Doe", "79991234567")

	msg := m.SendMessageReq{
		Attachments: []m.Attachment{attachment},
	}
*/
func AttachContact(name, phone string, contactID *int64) Attachment {
	vcfInfo := fmt.Sprintf("BEGIN:VCARD\nVERSION:3.0\nFN:%s\nTEL:%s\nEND:VCARD", name, phone)

	namePtr := name
	phonePtr := phone

	return Attachment{
		Type: AttachmentContact,
		Payload: ContactAttachmentPayload{
			Name:      &namePtr,
			ContactID: contactID,
			VCFInfo:   &vcfInfo,
			VCFPhone:  &phonePtr,
		},
	}
}

/* Keyboard represents the payload for an "inline_keyboard" attachment */
type Keyboard struct {
	Buttons [][]KeyboardButton `json:"buttons"`
}

/* KeyboardButton represents a single button in an inline keyboard */
type KeyboardButton struct {
	Type      ButtonType `json:"type"`
	Text      string     `json:"text"`              /* 1 to 128 characters */
	Payload   string     `json:"payload,omitempty"` /* up to 1024 characters */
	URL       string     `json:"url,omitempty"`
	Quick     bool       `json:"quick,omitempty"`
	WebApp    string     `json:"web_app,omitempty"`
	ContactID int        `json:"contact_id,omitempty"`
}

/* AttachKeyboard creates an inline keyboard attachment with custom buttons

Example:
	attachment := m.AttachKeyboard([][]m.KeyboardButton{
		{Type: m.ButtonLink, Text: "Visit Site", Payload: "https:/*max.ru"},
		{Type: m.ButtonCallback, Text: "Accept", Payload: "btn_accept"},
	})

	msg := m.SendMessageReq{
		Attachments: []m.Attachment{attachment},
	}
*/
func AttachKeyboard(buttons [][]KeyboardButton) Attachment {
	return Attachment{
		Type:    AttachmentKeyboard,
		Payload: Keyboard{Buttons: buttons},
	}
}

/* ShareAttachmentPayload represents the payload for a "share" attachment */
type ShareAttachmentPayload struct {
	URL         string `json:"url,omitempty"`
	Token       string `json:"token,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
}

/* AttachShare creates a link preview (share) attachment with an optional title and description.

Example:
	attachment := m.AttachShare("https:/*google.com", "Google Search", "The world's engine")

	msg := m.SendMessageReq{
		Attachments: []m.Attachment{attachment},
	}
*/
func AttachShare(url, title, desc string) Attachment {
	return Attachment{
		Type: AttachmentShare,
		Payload: ShareAttachmentPayload{
			URL:         url,
			Title:       title,
			Description: desc,
		},
	}
}

/* AttachLocation creates a location attachment with specified latitude and longitude

Example:
	attachment := m.AttachLocation(55.751244, 37.618423)

	msg := m.SendMessageReq{
		Attachments: []m.Attachment{attachment},
	}
*/
func AttachLocation(lat, lon float64) Attachment {
	return Attachment{
		Type:      AttachmentLocation,
		Latitude:  &lat,
		Longitude: &lon,
	}
}

/* BotInfo provides detailed information about a bot, extending the base User model */
type BotInfo struct {
	User
	Description   *string      `json:"description,omitempty"`
	AvatarURL     string       `json:"avatar_url,omitempty"`
	FullAvatarURL string       `json:"full_avatar_url,omitempty"`
	Commands      []BotCommand `json:"commands,omitempty"`
}

/* BotPatch provides bot parameters that can be changed by BotPatch method */
type BotPatch struct {
	Name        string                         `json:"name,omitempty"`
	Username    string                         `json:"username,omitempty"`
	Description string                         `json:"description,omitempty"`
	Commands    []BotCommand                   `json:"commands,omitempty"`
	Photo       *PhotoAttachmentRequestPayload `json:"photo,omitempty"`
}

/* BotCommand represents a specific command that the bot supports and can execute */
type BotCommand struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

/* User represents a participant (human or bot) in the messaging platform */
type User struct {
	UserID           int64  `json:"user_id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name,omitempty"`
	Username         string `json:"username,omitempty"`
	IsBot            bool   `json:"is_bot"`
	LastActivityTime int64  `json:"last_activity_time"`
}

/* Chat represents a conversation, which can be a dialog, a group chat, or a channel */
type Chat struct {
	ChatID            int64           `json:"chat_id"`
	Type              ChatType        `json:"type"`
	Status            ChatStatus      `json:"status"`
	Title             string          `json:"title,omitempty"`
	Icon              *Image          `json:"icon,omitempty"`
	LastEventTime     int64           `json:"last_event_time"`
	ParticipantsCount int32           `json:"participants_count"`
	OwnerID           int64           `json:"owner_id,omitempty"`
	Participants      *map[string]int `json:"participants,omitempty"`
	IsPublic          bool            `json:"is_public"`
	Link              string          `json:"link,omitempty"`
	Description       string          `json:"description,omitempty"`
	DialogWithUser    *DialogWithUser `json:"dialog_with_user,omitempty"`
	ChatMessageID     string          `json:"chat_message_id"`
	PinnedMessage     *Message        `json:"pinned_message,omitempty"`
}

/* ChatMember represents a participant (human or bot) in a group chat or a channel */
type ChatMember struct {
	DialogWithUser
	LastAccessTime int64                 `json:"last_access_time"`
	IsOwner        bool                  `json:"is_owner"`
	IsAdmin        bool                  `json:"is_admin"`
	JoinTime       int64                 `json:"join_time"`
	Permissions    []ChatAdminPermission `json:"permissions"`
	Alias          string                `json:"alias,omitempty"`
}

/* ChatAdmin represents an administrative user in a chat along with their specific permissions */
type ChatAdmin struct {
	UserID      int64                 `json:"user_id"`
	Permissions []ChatAdminPermission `json:"permissions"`
	Alias       string                `json:"alias,omitempty"`
}

/* Image is a generic schema for an image object */
type Image struct {
	URL string `json:"url"`
}

/* DialogWithUser contains details about the other participant in a dialog */
type DialogWithUser struct {
	User
	Description   string `json:"description,omitempty"`
	AvatarURL     string `json:"avatar_url"`
	FullAvatarURL string `json:"full_avatar_url"`
}

/* MessageList represents a paginated list of messages */
type MessagesList struct {
	Messages []Message `json:"messages"`
}

/* MessageStat contains statistics about messages */
type MessageStat struct {
	Views int `json:"views"`
}

/* Message in chat */
type Message struct {
	Sender        User           `json:"sender"`
	Recipient     Recipient      `json:"recipient"`
	Timestamp     int64          `json:"timestamp"`
	LinkedMessage *LinkedMessage `json:"link,omitempty"`
	Body          MessageBody    `json:"body"`
	Stat          *MessageStat   `json:"stat,omitempty"`
	URL           string         `json:"url,omitempty"`
}

/* Recipient New message recipient. Could be user or chat */
type Recipient struct {
	ChatID   int64    `json:"chat_id,omitempty"`
	ChatType ChatType `json:"chat_type"`
	UserID   int64    `json:"user_id,omitempty"`
}

/* Link represents a reference to another message, such as a reply or a forwarded message */
type LinkedMessage struct {
	Type    LinkedMessageType `json:"type"`
	Sender  *User             `json:"sender,omitempty"`
	ChatID  int64             `json:"chat_id,omitempty"`
	Message *MessageBody      `json:"message,omitempty"`
}

/* MessageBody represents the body of a message */
type MessageBody struct {
	Mid         string          `json:"mid"`
	Seq         int64           `json:"seq"`
	Text        string          `json:"text,omitempty"`
	Attachments []Attachment    `json:"attachments,omitempty"`
	Markup      []MarkupElement `json:"markup,omitempty"`
}

/* Photos represents a tokenized reference to an uploaded photo */
type Photos struct {
	Token string `json:"token"`
}

/* MarkupElement represents a generic message formatting schema */
type MarkupElement struct {
	Type   MarkupType `json:"type"`
	From   int        `json:"from"`
	Length int        `json:"length"`
}

/* VideoInfo provides details and playback status for an uploaded video attachment */
type VideoInfo struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Duration int    `json:"duration"`
	URL      string `json:"url"`
}

/* VideoUrls provides a collection of URLs for a video at various resolutions and formats */
type VideoUrls struct {
	Mp4_1080 string `json:"mp4_1080,omitempty"`
	Mp4_720  string `json:"mp4_720,omitempty"`
	Mp4_480  string `json:"mp4_480,omitempty"`
	Mp4_360  string `json:"mp4_360,omitempty"`
	Mp4_240  string `json:"mp4_240,omitempty"`
	Mp4_144  string `json:"mp4_144,omitempty"`
	Hls      string `json:"hls,omitempty"`
}

/* NewMessageBody contains the payload required to send or edit a message */
type NewMessageBody struct {
	Text        string          `json:"text,omitempty"`
	Attachments []Attachment    `json:"attachments,omitempty"`
	Link        *NewMessageLink `json:"link,omitempty"`
	Notify      bool            `json:"notify,omitempty"`
	Format      Format          `json:"format,omitempty"`
}

/* NewMessageLink represents a request to link a new message to an existing one */
type NewMessageLink struct {
	Type LinkedMessageType `json:"type"`
	Mid  string            `json:"mid"`
}

/* Update represents different types of events that occurred in the chat */
type Update struct {
	UpdateType UpdateType `json:"update_type"`
	Timestamp  int64      `json:"timestamp"`
	Callback   *Callback  `json:"callback,omitempty"`
	Message    Message    `json:"message"`
	MessageID  string     `json:"message_id,omitempty"`
	ChatID     int        `json:"chat_id,omitempty"`
	UserID     int        `json:"user_id,omitempty"`
	MutedUntil int        `json:"muted_until,omitempty"`
	UserLocale string     `json:"user_locale,omitempty"`
	IsChannel  bool       `json:"is_channel,omitempty"`
}

/* Subscription represents a webhook configuration for receiving updates */
type Subscription struct {
	Url         string       `json:"url"`
	Time        int64        `json:"time"`
	UpdateTypes []UpdateType `json:"update_types,omitempty"`
}

/* Callback is an object sent to bots when a user presses a button */
type Callback struct {
	Timestamp  int64  `json:"timestamp"`
	CallbackID string `json:"callback_id"`
	Payload    string `json:"payload,omitempty"`
	User       User   `json:"user"`
}

/* MessageCallbackUpdate is triggered when a user presses a button */
type MessageCallbackUpdate struct {
	Update
	Callback Callback `json:"callback"`
	Message  *Message `json:"message,omitempty"`
}

/* FailedUserDetails contains information about users for whom a specific operation failed */
type FailedUserDetails struct {
	ErrorCode string `json:"error_code"`
	UserIDs   []int  `json:"user_ids"`
}

/* ChatInfo provides detailed metadata and state information for a specific chat instance */
type ChatInfo struct {
	ChatID            int              `json:"chat_id"`
	Type              ChatType         `json:"type"`
	Status            ChatStatus       `json:"status"`
	Title             string           `json:"title,omitempty"`
	Icon              Image            `json:"icon,omitempty"`
	LastEventTime     int64            `json:"last_event_time"`
	ParticipantsCount int32            `json:"participants_count,omitempty"`
	OwnerID           int64            `json:"owner_id,omitempty"`
	Participants      map[string]int   `json:"participants,omitempty"`
	IsPublic          bool             `json:"is_public"`
	Link              string           `json:"link,omitempty"`
	Description       string           `json:"description,omitempty"`
	DialogWithUser    []DialogWithUser `json:"dialog_with_user,omitempty"`
	ChatMessageID     string           `json:"chat_message_id,omitempty"`
	PinnedMessage     []Message        `json:"pinned_message,omitempty"`
}
