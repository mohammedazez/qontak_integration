package qontak

import "time"

type ErrorResponse struct {
	Status string `json:"status"`
	Error  struct {
		Code     int      `json:"code"`
		Messages []string `json:"messages"`
	} `json:"error"`
}

type GeneralResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Meta   interface{} `json:"meta,omitempty"`
}

type GetTemplateRequest struct {
	Cursor          string `json:"cursor,omitempty"`
	CursorDirection string `json:"cursor_direction,omitempty"` // default after; before | after
	HSMChat         bool   `json:"hsm_chat,omitempty"`
	IsCounted       bool   `json:"is_counted,omitempty"`
	Limit           int    `json:"limit,omitempty"`
	Offset          int    `json:"offset,omitempty"`
	OrderBy         string `json:"order_by,omitempty"`
	OrderDirection  string `json:"order_direction,omitempty"`
	Query           string `json:"query,omitempty"`
	Status          string `json:"status,omitempty"`
}

type CreateMassageRequest struct {
	RoomID string `json:"room_id"`
	Type   string `json:"type"`
	Text   string `json:"text"`
}

type CreateMessageResponse struct {
	Status string `json:"status"`
	Data   struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		RoomID     string `json:"room_id"`
		IsCampaign bool   `json:"is_campaign"`
		SenderID   string `json:"sender_id"`
		SenderType string `json:"sender_type"`
		Sender     struct {
			Name   string `json:"name"`
			Avatar struct {
				URL   string `json:"url"`
				Large struct {
					URL string `json:"url"`
				} `json:"large"`
				Filename interface{} `json:"filename"`
				Size     int         `json:"size"`
				Small    struct {
					URL string `json:"url"`
				} `json:"small"`
				Medium struct {
					URL string `json:"url"`
				} `json:"medium"`
			} `json:"avatar"`
			Username string `json:"username"`
		} `json:"sender"`
		ParticipantID   string        `json:"participant_id"`
		OrganizationID  string        `json:"organization_id"`
		Text            string        `json:"text"`
		Status          string        `json:"status"`
		ParticipantType string        `json:"participant_type"`
		ExternalID      interface{}   `json:"external_id"`
		LocalID         interface{}   `json:"local_id"`
		CreatedAt       time.Time     `json:"created_at"`
		Buttons         []interface{} `json:"buttons"`
		Reply           interface{}   `json:"reply"`
		Room            struct {
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
			Avatar               struct {
				URL   string `json:"url"`
				Large struct {
					URL string `json:"url"`
				} `json:"large"`
				Filename interface{} `json:"filename"`
				Size     int         `json:"size"`
				Small    struct {
					URL string `json:"url"`
				} `json:"small"`
				Medium struct {
					URL string `json:"url"`
				} `json:"medium"`
			} `json:"avatar"`
			ResolvedAt     interface{} `json:"resolved_at"`
			ExternalID     string      `json:"external_id"`
			ResolvedByID   interface{} `json:"resolved_by_id"`
			ResolvedByType interface{} `json:"resolved_by_type"`
			Extra          struct {
				IsParticipantOnline bool `json:"is_participant_online"`
			} `json:"extra"`
			IsDontAutoResolve bool        `json:"is_dont_auto_resolve"`
			IsBlocked         bool        `json:"is_blocked"`
			LastMessageAt     time.Time   `json:"last_message_at"`
			LastActivityAt    time.Time   `json:"last_activity_at"`
			IsUnresponded     bool        `json:"is_unresponded"`
			DivisionID        interface{} `json:"division_id"`
		} `json:"room"`
		URL        string `json:"url"`
		IsEdited   bool   `json:"is_edited"`
		ReviewStar int    `json:"review_star"`
	} `json:"data"`
}

type HitAPIBulkRequest struct {
	URL          string        `json:"url,omitempty"`
	Method       string        `json:"method,omitempty"`
	BodyStr      string        `json:"body,omitempty"`
	Token        string        `json:"str_token,omitempty"`
	AuthType     string        `json:"auth_type,omitempty"`
	Alias        string        `json:"alias,omitempty"`
	Timeout      time.Duration `json:"timeout,omitempty"`
	Simulation   bool          `json:"Simulation,omitempty"`
	InsertApiLog bool          `json:"insertApiLog,omitempty"`
}

type QontakMessageWABlastWithHeader struct {
	ToName               string                        `json:"to_name"`
	ToNumber             string                        `json:"to_number"`
	MessageTemplateID    string                        `json:"message_template_id"`
	ChannelIntegrationID string                        `json:"channel_integration_id"`
	Language             QontakLanguage                `json:"language"`
	Parameters           QontakParametersWABlastHeader `json:"parameters,omitempty"`
}

type QontakMessageWABlast struct {
	ToName               string                  `json:"to_name"`
	ToNumber             string                  `json:"to_number"`
	MessageTemplateID    string                  `json:"message_template_id"`
	ChannelIntegrationID string                  `json:"channel_integration_id"`
	Language             QontakLanguage          `json:"language"`
	Parameters           QontakParametersWABlast `json:"parameters,omitempty"`
}

type QontakLanguage struct {
	Code string `json:"code"`
}

type QontakParametersWABlast struct {
	Body []QontakBodyItem `json:"body,omitempty"`
}

type QontakParametersWABlastHeader struct {
	Header  *QontakHeaderItem `json:"header,omitempty"`
	Body    []QontakBodyItem  `json:"body,omitempty"`
	Buttons []QontakButton    `json:"buttons,omitempty"`
}

type QontakParameters struct {
	Header  QontakHeaderItem `json:"header"`
	Body    []QontakBodyItem `json:"body"`
	Buttons []QontakButton   `json:"buttons"`
}

type QontakHeaderItem struct {
	Format string              `json:"format"`
	Params []QontakHeaderParam `json:"params"`
}

type QontakHeaderParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type QontakBodyItem struct {
	Key       string `json:"key"`
	ValueText string `json:"value_text"`
	Value     string `json:"value"`
}

type QontakButton struct {
	Index string `json:"index"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type BroadcastDirectResponse struct {
	Data struct {
		ChannelAccountName    string      `json:"channel_account_name"`
		ChannelIntegrationID  string      `json:"channel_integration_id"`
		ChannelPhoneNumber    string      `json:"channel_phone_number"`
		ContactExtra          interface{} `json:"contact_extra"`
		ContactID             string      `json:"contact_id"`
		ContactListID         interface{} `json:"contact_list_id"`
		CreatedAt             time.Time   `json:"created_at"`
		DivisionID            interface{} `json:"division_id"`
		ExecuteStatus         string      `json:"execute_status"`
		ExecuteType           string      `json:"execute_type"`
		ID                    string      `json:"id"`
		MessageBroadcastError string      `json:"message_broadcast_error"`
		MessageStatusCount    struct {
			Delivered int `json:"delivered"`
			Failed    int `json:"failed"`
			Pending   int `json:"pending"`
			Read      int `json:"read"`
			Sent      int `json:"sent"`
		} `json:"message_status_count"`
		MessageTemplate struct {
			Body    string `json:"body"`
			Buttons []struct {
				Text string `json:"text"`
				Type string `json:"type"`
			} `json:"buttons"`
			Category string      `json:"category"`
			Footer   interface{} `json:"footer"`
			Header   struct {
				Example string `json:"example"`
				Format  string `json:"format"`
			} `json:"header"`
			ID                string      `json:"id"`
			Language          string      `json:"language"`
			Name              string      `json:"name"`
			OrganizationID    string      `json:"organization_id"`
			QualityRating     string      `json:"quality_rating"`
			QualityRatingText interface{} `json:"quality_rating_text"`
			Status            string      `json:"status"`
		} `json:"message_template"`
		Name           string `json:"name"`
		OrganizationID string `json:"organization_id"`
		Parameters     struct {
			Body    interface{} `json:"body"`
			Buttons struct {
			} `json:"buttons"`
			Header struct {
				Format string `json:"format"`
				Params struct {
					Filename string `json:"filename"`
					URL      string `json:"url"`
				} `json:"params"`
			} `json:"header"`
		} `json:"parameters"`
		SendAt        time.Time `json:"send_at"`
		SenderEmail   string    `json:"sender_email"`
		SenderName    string    `json:"sender_name"`
		TargetChannel string    `json:"target_channel"`
	} `json:"data"`
	Status string `json:"status"`
}

type ResolvedRequest struct {
	Status string `json:"status"`
}
