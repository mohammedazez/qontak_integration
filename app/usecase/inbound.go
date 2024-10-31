package usecase

import (
	"encoding/json"
	"fmt"
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/ewinjuman/go-lib/helper/convert"
	"github.com/ewinjuman/go-lib/session"
	"log"
	"qontak_integration/app/domain/queries"
	"qontak_integration/app/models"
	"qontak_integration/pkg/repository"
	"qontak_integration/pkg/utils"
	"qontak_integration/platform/http/forward"
	"qontak_integration/platform/http/qontak"
	"strings"
)

type (
	InboundUsecaseService interface {
		QontakInbound(generalRequest *models.QontakGeneralMessage, request []byte) (interface{}, error)
		QontakArchived(request *models.ArchivedRequest) (interface{}, error)
	}

	inboundUsecase struct {
		session           *session.Session
		query             queries.QueriesService
		qontakHttp        qontak.QontakHttpService
		forwardStatusHttp forward.ForwardStatusHttpService
	}
)

func NewInboundUsecase(session *session.Session) (item InboundUsecaseService) {
	//userGrpc, err := user.NewServerContext()
	return &inboundUsecase{
		session:           session,
		query:             queries.NewQueries(session),
		qontakHttp:        qontak.NewQontakHttp(session),
		forwardStatusHttp: forward.NewForwardHttp(session),
	}
}

func (i *inboundUsecase) QontakArchived(request *models.ArchivedRequest) (response interface{}, err error) {
	resolvedRequest := qontak.ResolvedRequest{Status: "archived"}
	session, err := i.query.GetSession(request.ClientID)
	if err != nil {
		return
	}
	result, err := i.qontakHttp.Resolved(resolvedRequest, request.ConversationID, fmt.Sprintf("Bearer %s", session.Token))
	if err != nil {
		return
	}
	response = result
	return
}

func (i *inboundUsecase) QontakInbound(generalRequest *models.QontakGeneralMessage, request []byte) (response interface{}, err error) {

	if generalRequest.WebhookEvent != "message_interaction" {
		err = Error.NewError(repository.NotFoundCode, repository.FailedStatus, "event not found")
		return
	}
	//var payload interface{}
	payload, channelIntegrationID, err := i.createPayload(generalRequest.DataEvent, request)
	if err != nil {
		return nil, err
	}
	webhooks, err := i.getWebhooksBasedOnEvent(generalRequest.DataEvent, channelIntegrationID)
	if len(webhooks) <= 0 {
		err = Error.NewError(repository.NotFoundCode, repository.FailedStatus, "webhook not found")
		return
	}

	return i.processWebhooks(webhooks, payload)
}

func (uc *inboundUsecase) getWebhooksBasedOnEvent(event, channelIntegrationID string) ([]models.WebhookResult, error) {
	var (
		webhookType  string
		eventchannel string
	)

	switch event {
	case "receive_message_from_customer", "receive_message_from_agent":
		webhookType = "Inbound"
		eventchannel = "Whatsapp"
	case "status_message":
		webhookType = "ForwardStatus"
		eventchannel = "WhatsappStatus"
	case "broadcast_log_status":
		webhookType = "ForwardStatus"
		eventchannel = "WhatsappStatus"
	default:
		return nil, Error.NewError(repository.NotFoundCode, repository.FailedStatus, "event not found")
	}

	webhooks, err := uc.getWebhook(webhookType, eventchannel, channelIntegrationID)
	if err != nil || len(webhooks) == 0 {
		return nil, Error.NewError(repository.NotFoundCode, repository.FailedStatus, "webhook not found")
	}

	return webhooks, nil
}

func (i *inboundUsecase) getWebhook(trxType, eventChannel, to string) (filteredWebhooks []models.WebhookResult, err error) {

	result, err := i.query.GetWebhookBychannel(trxType)
	if err != nil {
		return
	}
	//webhooks := []models.WebhookResult{}
	for _, w := range result {
		webhook := models.WebhookResult{}
		var events []models.Event
		convert.ObjectToObject(w, &webhook)

		err := json.Unmarshal(w.Events, &events)
		if err != nil {
			log.Printf("Error unmarshalling events for webhook ID %d: %v", webhook.ID, err)
			continue
		}
		webhook.Events = events

		for _, event := range webhook.Events {
			//println(event)
			if event.ChannelName == eventChannel && event.To == to {
				filteredWebhooks = append(filteredWebhooks, webhook)
				break
			}
		}

	}
	return
}

func (uc *inboundUsecase) createPayload(dataEvent string, request []byte) (interface{}, string, error) {
	switch dataEvent {
	case "receive_message_from_customer", "receive_message_from_agent":
		inboundMessage := models.QontakInboundMessage{}
		convert.StringToObject(string(request), &inboundMessage)
		return forward.InboundRequest{
			Channel:        "Whatsapp",
			ContactID:      inboundMessage.SenderID,
			ContactName:    inboundMessage.Sender.Name,
			ConversationID: inboundMessage.RoomID,
			MessageID:      inboundMessage.ID,
			To:             "",
			From:           "+" + inboundMessage.Room.AccountUniqID,
			Type:           inboundMessage.Type,
			Content: struct {
				Text string `json:"text"`
			}{inboundMessage.Text},
			Status:          inboundMessage.Room.Status,
			CreatedDatetime: inboundMessage.Room.LastMessageAt,
			UpdatedDatetime: inboundMessage.Room.LastMessageAt,
		}, inboundMessage.Room.ChannelIntegrationID, nil
	case "status_message":
		inboundMessage := models.QontakInboundMessage{}
		convert.StringToObject(string(request), &inboundMessage)
		return inboundMessage, inboundMessage.Room.ChannelIntegrationID, nil
	case "broadcast_log_status":
		inboundMessage := models.QontakBroadcastLog{}
		convert.StringToObject(string(request), &inboundMessage)
		if strings.Contains(inboundMessage.Messages.Body.Template, "#{{") {
			keyTrID := utils.GetStringBetween(inboundMessage.Messages.Body.Template, "#{{", "}}")
			parameterMaps := inboundMessage.Messages.Body.Parameters.(map[string]interface{})
			inboundMessage.Messages.Body.Template = utils.FulfillTemplate(inboundMessage.Messages.Body.Template, parameterMaps)
			//qim.Text = qim.Messages.Body.Template
			if keyTrID != "" && parameterMaps[keyTrID] != nil {
				inboundMessage.TrID = parameterMaps[keyTrID].(string)
			}
		}
		return inboundMessage, inboundMessage.ChannelIntegrationID, nil
	default:
		return nil, "", Error.NewError(repository.NotFoundCode, repository.FailedStatus, "event not found")
	}
	return nil, "", nil
}

func (uc *inboundUsecase) processWebhooks(webhooks []models.WebhookResult, payload interface{}) (interface{}, error) {
	for _, webhook := range webhooks {
		result, err := uc.forwardStatusHttp.Forward(payload, webhook)
		if err != nil {
			uc.session.Info(err.Error())
		}
		// TODO: save to database
		println(result)
	}
	return nil, nil
}
