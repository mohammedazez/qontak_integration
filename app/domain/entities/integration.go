package entities

import (
	"gorm.io/datatypes"
	"time"
)

type Client struct {
	ID         int    `json:"id"`
	ClientCode string `json:"client_code"`
	ClientName string `json:"client_name"`
	Status     int    `json:"status"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  string `json:"created_by"`
}

func (Client) TableName() string {
	return "client"
}

type Sessions struct {
	ID        int    `json:"id"`
	ClientID  string `json:"client_id"`
	Token     string `json:"token"`
	Status    int    `json:"status"`
	ExpiredAt string `json:"expired_at"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
}

func (Sessions) TableName() string {
	return "sessions"
}

type Channel struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Uri          string `json:"uri"`
	Method       string `json:"method"`
	HeaderPrefix string `json:"header_prefix"`
	Status       int    `json:"status"`
	CreatedAt    string `json:"created_at"`
	CreatedBy    string `json:"created_by"`
}

func (Channel) TableName() string {
	return "channel"
}

type ClientChannel struct {
	ID        int    `json:"id"`
	ChannelID string `json:"channel_id"`
	Channel   string `json:"channel"`
	ClientID  int    `json:"client_id"`
	Status    int    `json:"status"`
	VendorID  uint   `gorm:"column:vendor_id"`
	Vendor    Vendor `gorm:"foreignKey:VendorID"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
}

func (ClientChannel) TableName() string {
	return "client_channel"
}

type Webhook struct {
	ID               int            `json:"id"`
	ClientID         string         `json:"client_id"`
	ClientCode       string         `json:"client_code"`
	TrxType          string         `json:"trx_type"`
	Url              string         `json:"url"`
	Method           string         `json:"method"`
	HeaderPrefix     string         `json:"header_prefix"`
	Token            string         `json:"token"`
	Events           datatypes.JSON `json:"events"`
	ExpectedHttpCode int            `json:"expected_http_code"`
	Retry            int            `json:"retry"`
	Timeout          int            `json:"timeout"`
	Status           int            `json:"status"`
	CreatedAt        time.Time      `json:"created_at"`
	CreatedBy        string         `json:"created_by"`
}

func (Webhook) TableName() string {
	return "webhooks"
}

type Event struct {
	//EventName string `json:"event_name"`
	ChannelName      string `json:"channel_name,omitempty"`
	From             string `json:"from,omitempty"`
	To               string `json:"to,omitempty"`
	WaBusinessNumber string `json:"wa_business_number,omitempty"`
}

type Vendor struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"column:code"`
}

func (Vendor) TableName() string {
	return "vendors"
}

type ApiLogs struct {
	ID             int64     `json:"id" gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Url            string    `json:"url" gorm:"column:url;NOT NULL"`
	Method         string    `json:"method" gorm:"column:method;NOT NULL"`
	RequestHeader  string    `json:"requestHeader" gorm:"column:request_header;NOT NULL"`
	RequestBody    string    `json:"requestBody" gorm:"column:request_body;NOT NULL"`
	ResponseStatus string    `json:"responseStatus" gorm:"column:response_status;NOT NULL"`
	ResponseHeader string    `json:"responseHeader" gorm:"column:response_header;NOT NULL"`
	ResponseBody   string    `json:"responseBody" gorm:"column:response_body;NOT NULL"`
	CreatedAt      time.Time `json:"createdAt" gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	CreatedBy      string    `json:"createdBy" gorm:"column:created_by;NOT NULL"`
}

func (ApiLogs) TableName() string {
	return "api_logs"
}
