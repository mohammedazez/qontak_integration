package models

import "time"

type GetTemplateRequest struct {
	ClientID int `json:"client_id"`
	Offset   int `json:"offset"`
	Limit    int `json:"limit"`
}

type GetTemplateResponse struct {
	ClientVendor struct {
		ChannelID string `json:"channel_id"`
	} `json:"client_vendor"`
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

type OutboundRequest struct {
	Messages OutMessage `json:"messages" validate:"required"`
	ClientID int        `json:"client_id" validate:"required"`
	Channel  string     `json:"channel" validate:"required"`
}

type OutboundBulkRequest struct {
	Messages []OutMessage `json:"messages" validate:"required"`
	ClientID int          `json:"client_id" validate:"required"`
	Channel  string       `json:"channel" validate:"required"`
}

type OutMessage struct {
	Channel         string      `json:"channel" validate:"required,oneof=Whatsapp BlastWhatsapp BlastWhatsappHeader"`
	ChannelID       string      `json:"channel_id,omitempty"` //for MessageBird
	From            string      `json:"from,omitempty"`
	SesameID        string      `json:"sesame_id,omitempty"`
	MessageID       string      `json:"id,omitempty"`
	SesameMessageID string      `json:"sesame_message_id,omitempty"`
	Timestamp       int64       `json:"timestamp" validate:"required"`
	Content         Content     `json:"content,omitempty"`
	To              string      `json:"to,omitempty"`
	WATemplateID    string      `json:"wa_template_id,omitempty"`
	ConversationID  string      `json:"conversation_id,omitempty"` // QONTAK PROJECT: di qontak ini room id
	Type            string      `json:"type" validate:"required,oneof=template hsm text audio document image video sticker flow interactive text_template button_template"`
	Hsm             Hsm         `json:"hsm,omitempty"`
	Name            string      `json:"name,omitempty"`
	Components      []Component `json:"components,omitempty"`
	//URLAttachmengts []*Attachments `json:"url_attachments,omitempty"`
	CampaignId string `json:"campaign_id,omitempty"`
	//CallBackUrl     string         `json:"call_back_url,omitempty"`
	//CallBackAuth  CallBackAuth `json:"call_back_auth,omitempty"`
	//TrxId         string                        `json:"trx_id,omitempty"`
	//RecipientType string                        `json:"recipient_type,omitempty"`
	Interactive *Outbound_Interactive `json:"interactive,omitempty"`
}
type Outbound_Interactive struct {
	Type   *string             `json:"type,omitempty"`
	Body   *Interactive_Body   `json:"body,omitempty"`
	Action *Interactive_Action `json:"action,omitempty"`
}

type Interactive_Body struct {
	Text *string `json:"text,omitempty"`
}

type Interactive_Action struct {
	Button   *string          `json:"button,omitempty"`
	Sections *[]Section       `json:"sections,omitempty"`
	Buttons  *[]Action_Button `json:"buttons,omitempty"`
}
type Action_Button struct {
	Type  *string              `json:"type,omitempty"`
	Reply *Action_Button_reply `json:"reply,omitempty"`
}

type Action_Button_reply struct {
	Id    *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
}

type Section struct {
	Title *string        `json:"title,omitempty"`
	Rows  *[]Section_Row `json:"rows,omitempty"`
}

type Section_Row struct {
	Id          *string `json:"id,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}
type Attachments struct {
	Location *OutLocation `json:"location,omitempty"`
	Image    *OutImage    `json:"image,omitempty"`
	Document *OutDocument `json:"document,omitempty"`
	Audio    *OutAudio    `json:"audio,omitempty"`
	Video    *OutVideo    `json:"video,omitempty"`
}
type Component struct {
	Type       string    `json:"type"` //Type for header or body
	Subtype    string    `json:"sub_type,omiyempty"`
	Index      int       `json:"index,omitempty"`
	Parameters []PrmType `json:"parameters"`
}

type PrmType struct {
	Type     string        `json:"type"` //Type for document, image, text, currency, date_time
	Text     string        `json:"text,omitempty"`
	Document MediaDoc      `json:"document,omitempty"`
	Image    MediaImage    `json:"image,omitempty"`
	Video    MediaVideo    `json:"video,omitempty"`
	Currency MediaCurrency `json:"currency,omitempty"`
	DateTime MediaDate     `json:"date_time,omitempty"`
}

type MediaDate struct {
	Fallback   string `json:"fallback_value"`
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth string `json:"day_of_month"`
	Year       int    `json:"year"`
	Month      int    `json:"month"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Timestamp  int64  `json:"timestamp"`
}

type MediaCurrency struct {
	Fallback string `json:"fallback_value"`
	Code     string `json:"code"`
	Amount   string `json:"amount_1000"`
}

type MediaImage struct {
	Link string `json:"link"`
	Name string `json:"name"`
	Url  string `json:"url,omitempty"`
}

type MediaVideo struct {
	Url string `json:"url,omitempty"`
}

type MediaDoc struct {
	Link     string   `json:"url"`
	Provider Provider `json:"provider"`
	Filename string   `json:"filename"`
}

type Provider struct {
	Name string `json:"name"`
}

type Content struct {
	Hsm       *Hsm         `json:"hsm,omitempty"`
	Text      string       `json:"text,omitempty"`
	Audio     *OutAudio    `json:"audio,omitempty"`
	Image     *OutImage    `json:"image,omitempty"`
	Document  *OutDocument `json:"document,omitempty"`
	Video     *OutVideo    `json:"video,omitempty"`
	Sticker   *OutSticker  `json:"sticker,omitempty"`
	TextEmail string       `json:"text_email,omitempty"`
	Name      string       `json:"name,omitempty"`
}

type Hsm struct {
	Namespace    string   `json:"namespace"`
	TemplateName string   `json:"templateName,omitempty"`
	ElementName  string   `json:"element_name,omitempty"`
	Lang         Language `json:"language"`
	Localizable  []Params `json:"localizable_params,omitempty"`
	Params       []Params `json:"params,omitempty"`
}

type Params struct {
	Default string `json:"default,omitempty"`
}

type Language struct {
	Policy string `json:"policy"`
	Code   string `json:"code"`
}

type OutText struct {
	Body string `json:"body"`
}

type OutSticker struct {
	Caption string `json:"caption,omitempty"`
	Link    string `json:"url,omitempty"`
}

type OutLocation struct {
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

type OutAudio struct {
	Caption string `json:"caption,omitempty"`
	Link    string `json:"url,omitempty"`
}

type OutVideo struct {
	Caption string `json:"caption,omitempty"`
	Link    string `json:"url,omitempty"`
}

type OutImage struct {
	Caption string `json:"caption,omitempty"`
	Link    string `json:"url,omitempty"`
}

type OutDocument struct {
	Caption string `json:"caption,omitempty"`
	Link    string `json:"url,omitempty"`
}

//======================================

type MessageWrapping struct {
	ClientID    int         `json:"client_id" validate:"required"`
	Channel     string      `json:"channel" validate:"required"`
	ClientToken string      `json:"client_token" validate:"required"`
	VendorCode  string      `json:"vendor_code"`
	Message     interface{} `json:"message" validate:"required"`
}

type WebhookResult struct {
	ID               int       `json:"id"`
	ClientID         string    `json:"client_id"`
	ClientCode       string    `json:"client_code"`
	TrxType          string    `json:"trx_type"`
	Url              string    `json:"url"`
	Method           string    `json:"method"`
	HeaderPrefix     string    `json:"header_prefix"`
	Token            string    `json:"token"`
	Events           []Event   `json:"events"`
	ExpectedHttpCode int       `json:"expected_http_code"`
	Retry            int       `json:"retry"`
	Timeout          int       `json:"timeout"`
	Status           int       `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	CreatedBy        string    `json:"created_by"`
}

type Event struct {
	//EventName string `json:"event_name"`
	ChannelName      string `json:"channel_name,omitempty"`
	From             string `json:"from,omitempty"`
	To               string `json:"to,omitempty"`
	WaBusinessNumber string `json:"wa_business_number,omitempty"`
}

type ArchivedRequest struct {
	ConversationID string `json:"conversation_id" validate:"required"`
	ClientID       int    `json:"client_id" validate:"required"`
}
