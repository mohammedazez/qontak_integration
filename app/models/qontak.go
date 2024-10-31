package models

import "time"

type QontakGeneralMessage struct {
	ID           string `json:"id"`
	DataEvent    string `json:"data_event"`
	WebhookEvent string `json:"webhook_event"`
}

type QontakInboundMessage struct {
	ID                 string                   `json:"id"`
	Type               string                   `json:"type"`
	RoomID             string                   `json:"room_id"`
	File               QontakFile               `json:"file"`
	FileUniqID         string                   `json:"file_uniq_id"`
	IsCampaign         bool                     `json:"is_campaign"`
	SenderID           string                   `json:"sender_id"`
	SenderType         string                   `json:"sender_type"`
	ParticipantID      string                   `json:"participant_id"`
	OrganizationID     string                   `json:"organization_id"`
	Text               string                   `json:"text"`
	Status             string                   `json:"status"`
	ExternalID         string                   `json:"external_id"`
	LocalID            interface{}              `json:"local_id"`
	CreatedAt          time.Time                `json:"created_at"`
	Buttons            []interface{}            `json:"buttons"`
	Reply              interface{}              `json:"reply"`
	Room               QontakRoom               `json:"room"`
	Sender             QontakSender             `json:"sender"`
	ParticipantType    string                   `json:"participant_type"`
	ChannelIntegration QontakChannelIntegration `json:"channel_integration"`
	LastCampaign       map[string]interface{}   `json:"last_campaign"`
	DataEvent          string                   `json:"data_event"`
	WebhookEvent       string                   `json:"webhook_event"`
}

type QontakFile struct {
	Filename string `json:"filename"`
	Large    struct {
		URL *string `json:"url"`
	} `json:"large"`
	Medium struct {
		URL *string `json:"url"`
	} `json:"medium"`
	Size  int `json:"size"`
	Small struct {
		URL *string `json:"url"`
	} `json:"small"`
	URL string `json:"url"`
}

type QontakRoom struct {
	ID                   string        `json:"id"`
	Name                 string        `json:"name"`
	Description          string        `json:"description"`
	Status               string        `json:"status"`
	Type                 string        `json:"type"`
	Tags                 []interface{} `json:"tags"`
	Channel              string        `json:"channel"`
	ChannelAccount       string        `json:"channel_account"`
	OrganizationID       string        `json:"organization_id"`
	AccountUniqID        string        `json:"account_uniq_id"`
	ChannelIntegrationID string        `json:"channel_integration_id"`
	SessionAt            time.Time     `json:"session_at"`
	UnreadCount          int           `json:"unread_count"`
	CreatedAt            time.Time     `json:"created_at"`
	UpdatedAt            time.Time     `json:"updated_at"`
	Avatar               QontakAvatar  `json:"avatar"`
	ResolvedAt           interface{}   `json:"resolved_at"`
	ExternalID           string        `json:"external_id"`
	ResolvedByID         interface{}   `json:"resolved_by_id"`
	ResolvedByType       interface{}   `json:"resolved_by_type"`
	Extra                QontakExtra   `json:"extra"`
	IsDontAutoResolve    bool          `json:"is_dont_auto_resolve"`
	IsBlocked            bool          `json:"is_blocked"`
	LastMessageAt        time.Time     `json:"last_message_at"`
	LastActivityAt       time.Time     `json:"last_activity_at"`
	IsUnresponded        bool          `json:"is_unresponded"`
	DivisionID           interface{}   `json:"division_id"`
}

type QontakSender struct {
	Name   string       `json:"name"`
	Avatar QontakAvatar `json:"avatar"`
}

type QontakChannelIntegration struct {
	ID            string `json:"id"`
	TargetChannel string `json:"target_channel"`
}

type QontakAvatar struct {
	URL      string            `json:"url"`
	Large    QontakSizeVariant `json:"large"`
	Filename interface{}       `json:"filename"`
	Size     int               `json:"size"`
	Small    QontakSizeVariant `json:"small"`
	Medium   QontakSizeVariant `json:"medium"`
}

// SizeVariant struct definition
type QontakSizeVariant struct {
	URL string `json:"url"`
}

// Extra struct definition
type QontakExtra struct {
	IsParticipantOnline bool `json:"is_participant_online"`
}

type QontakBroadcastLog struct {
	ContactFullName      string           `json:"contact_full_name"`
	ContactPhoneNumber   string           `json:"contact_phone_number"`
	CreatedAt            time.Time        `json:"created_at"`
	DataEvent            string           `json:"data_event"`
	ID                   string           `json:"id"`
	IsPacing             bool             `json:"is_pacing"`
	ChannelIntegrationID string           `json:"channel_integration_id"`
	TrID                 string           `json:"tr_id,omitempty"`
	Messages             Messages         `json:"messages"`
	MessagesBroadcastID  string           `json:"messages_broadcast_id"`
	MessagesResponse     MessagesResponse `json:"messages_response"`
	OrganizationID       string           `json:"organization_id"`
	Status               string           `json:"status"`
	WebhookEvent         string           `json:"webhook_event"`
	WhatsappErrorMessage string           `json:"whatsapp_error_message"`
	WhatsappMessageID    string           `json:"whatsapp_message_id"`
}

type Messages struct {
	Body    MessageBody    `json:"body"`
	Buttons MessageButtons `json:"buttons"`
	Header  MessageHeader  `json:"header"`
}

type MessageBody struct {
	Parameters interface{} `json:"parameters"`
	Template   string      `json:"template"`
	Type       string      `json:"type"`
}

type MessageButtons struct {
	Parameters []interface{}    `json:"parameters"`
	Template   []ButtonTemplate `json:"template"`
	Type       string           `json:"type"`
}

type ButtonTemplate struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type MessageHeader struct {
	Parameters []HeaderParameter `json:"parameters"`
	Template   HeaderTemplate    `json:"template"`
	Type       string            `json:"type"`
}

type HeaderParameter struct {
	Image Image  `json:"image"`
	Type  string `json:"type"`
}

type Image struct {
	Link string `json:"link"`
}

type HeaderTemplate struct {
	Example string `json:"example"`
	Format  string `json:"format"`
}

type MessagesResponse struct {
	Contacts         []Contact       `json:"contacts"`
	Delivered        DeliveryStatus  `json:"delivered"`
	Messages         []MessageStatus `json:"messages"`
	MessagingProduct string          `json:"messaging_product"`
	Sent             DeliveryStatus  `json:"sent"`
}

type Contact struct {
	Input string `json:"input"`
	WaID  string `json:"wa_id"`
}

type DeliveryStatus struct {
	Statuses []Status `json:"statuses"`
	Webhook  string   `json:"webhook"`
}

type Status struct {
	Conversation Conversation `json:"conversation"`
	ID           string       `json:"id"`
	Pricing      Pricing      `json:"pricing"`
	RecipientID  string       `json:"recipient_id"`
	Status       string       `json:"status"`
	Timestamp    string       `json:"timestamp"`
}

type Conversation struct {
	ID                  string `json:"id"`
	Origin              Origin `json:"origin"`
	ExpirationTimestamp string `json:"expiration_timestamp,omitempty"`
}

type Origin struct {
	Type string `json:"type"`
}

type Pricing struct {
	Billable     bool   `json:"billable"`
	Category     string `json:"category"`
	PricingModel string `json:"pricing_model"`
}

type MessageStatus struct {
	ID            string `json:"id"`
	MessageStatus string `json:"message_status"`
}
