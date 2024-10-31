package forward

import "time"

type InboundRequest struct {
	Channel        string `json:"channel"`
	ContactID      string `json:"contact_id"`
	ContactName    string `json:"contact_name"`
	ConversationID string `json:"conversation_id"`
	MessageID      string `json:"message_id"`
	To             string `json:"to"`
	From           string `json:"from"`
	Type           string `json:"type"`
	Content        struct {
		Text string `json:"text"`
	} `json:"content"`
	Status          string    `json:"status"`
	CreatedDatetime time.Time `json:"createdDatetime"`
	UpdatedDatetime time.Time `json:"updatedDatetime"`
}

//type QontakInbound struct {
//	Contact      ContactInbound      `json:"contact"`
//	Conversation ConversationInbound `json:"conversation"`
//	Message      MessageInbound      `json:"message"`
//	Type         string              `json:"type"`
//	MessageID    string              `json:"message_id"`
//	ClientID     int                 `json:"client_id"`
//	VendorAlias  string
////}
//
//type ContactInbound struct {
//	ID            string `json:"id"`
//	Href          string `json:"href"`
//	Msisdn        string `json:"msisdn"`
//	Displayname   string `json:"displayName"`
//	Firstname     string `json:"firstName"`
//	Lastname      string `json:"lastName"`
//	Customdetails struct {
//	} `json:"customDetails"`
//	Attributes struct {
//	} `json:"attributes"`
//	Createddatetime time.Time `json:"createdDatetime"`
//	Updateddatetime time.Time `json:"updatedDatetime"`
//}
//
//type ConversationInbound struct {
//	ID                   string    `json:"id"`
//	Contactid            string    `json:"contactId"`
//	Status               string    `json:"status"`
//	Createddatetime      time.Time `json:"createdDatetime"`
//	Updateddatetime      time.Time `json:"updatedDatetime"`
//	Lastreceiveddatetime time.Time `json:"lastReceivedDatetime"`
//	Lastusedchannelid    string    `json:"lastUsedChannelId"`
//	Messages             struct {
//		Totalcount int    `json:"totalCount"`
//		Href       string `json:"href"`
//	} `json:"messages"`
//}
//
//type MessageInbound struct {
//	ID              string      `json:"id"`
//	Conversationid  string      `json:"conversationId"`
//	Platform        string      `json:"platform"`
//	To              string      `json:"to"`
//	From            string      `json:"from"`
//	Cc              string      `json:"cc"`
//	Bcc             string      `json:"bcc"`
//	Subject         string      `json:"subject"`
//	Channelid       string      `json:"channelId"`
//	Type            string      `json:"type"`
//	Content         interface{} `json:"content"`
//	Direction       string      `json:"direction"`
//	Status          string      `json:"status"`
//	Createddatetime time.Time   `json:"createdDatetime"`
//	Updateddatetime time.Time   `json:"updatedDatetime"`
//}
