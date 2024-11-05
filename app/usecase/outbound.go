package usecase

import (
	"qontak_integration/app/domain/queries"
	"qontak_integration/app/models"
	"qontak_integration/app/vendors"
	"qontak_integration/platform/http/qontak"

	"github.com/ewinjuman/go-lib/session"
)

type (
	OutboundUsecaseService interface {
		SendWaMessage(request *models.SendWhatsappRequest) (response interface{}, err error)
		SendWaBroadCastDirect(request *models.BroadcastDirectRequest) (response interface{}, err error)
		SendWaBroadCastDirectBulk(request *models.BroadcastDirectBulkRequest) (response interface{}, err error)
		SendInstagramMessage(request *models.SendInstagramRequest) (response interface{}, err error)
		SendTelegramMessage(request *models.SendTelegramRequest) (response interface{}, err error)
		//SendWaMessageToQueue(request *models.OutboundRequest) (response interface{}, err error)
		//SendWaMessageToQueueBulk(request *models.OutboundBulkRequest) (response interface{}, err error)
	}

	outboundUsecase struct {
		session    *session.Session
		query      queries.QueriesService
		qontakHttp qontak.QontakHttpService
	}
)

func NewOutboundUsecase(session *session.Session) (item OutboundUsecaseService) {
	//userGrpc, err := user.NewServerContext()
	return &outboundUsecase{
		session: session,
		query:   queries.NewQueries(session),
	}
}

func (o *outboundUsecase) SendWaMessage(request *models.SendWhatsappRequest) (response interface{}, err error) {
	channel, err := o.query.GetClientChannel(request.ClientID, "Whatsapp")
	if err != nil {
		return nil, err
	}
	vendor, err := vendors.GetService(o.session, channel.Vendor.Code)
	if err != nil {
		return nil, err
	}

	sessions, err := o.query.GetSession(request.ClientID)
	if err != nil {
		return nil, err
	}

	return vendor.WaSendMessage(vendors.CredentialObject{
		ClientID: request.ClientID,
		Channel:  "Whatsapp",
		Token:    sessions.Token,
	}, request)
}

func (o *outboundUsecase) SendInstagramMessage(request *models.SendInstagramRequest) (response interface{}, err error) {
	channel, err := o.query.GetClientChannel(request.ClientID, "Whatsapp")
	if err != nil {
		return nil, err
	}
	vendor, err := vendors.GetService(o.session, channel.Vendor.Code)
	if err != nil {
		return nil, err
	}

	sessions, err := o.query.GetSession(request.ClientID)
	if err != nil {
		return nil, err
	}

	return vendor.InstagramSendMessage(vendors.CredentialObject{
		ClientID: request.ClientID,
		Channel:  "Instagram",
		Token:    sessions.Token,
	}, request)
}

func (o *outboundUsecase) SendWaBroadCastDirect(request *models.BroadcastDirectRequest) (response interface{}, err error) {

	channel, err := o.query.GetClientChannel(request.ClientID, request.Channel)
	if err != nil {
		return nil, err
	}
	vendor, err := vendors.GetService(o.session, channel.Vendor.Code)
	if err != nil {
		return nil, err
	}
	sessions, err := o.query.GetSession(request.ClientID)
	if err != nil {
		return nil, err
	}
	return vendor.WaBroadcastDirect(vendors.CredentialObject{
		ClientID: request.ClientID,
		Channel:  request.Channel,
		Token:    sessions.Token,
	}, request)
}

func (o *outboundUsecase) SendWaBroadCastDirectBulk(request *models.BroadcastDirectBulkRequest) (response interface{}, err error) {
	channel, err := o.query.GetClientChannel(request.ClientID, request.Channel)
	if err != nil {
		return nil, err
	}

	vendor, err := vendors.GetService(o.session, channel.Vendor.Code)
	if err != nil {
		return nil, err
	}
	sessions, err := o.query.GetSession(request.ClientID)
	if err != nil {
		return nil, err
	}
	return vendor.WaBroadcastDirectBulk(vendors.CredentialObject{
		ClientID: request.ClientID,
		Channel:  request.Channel,
		Token:    sessions.Token,
	}, request)
}

func (o *outboundUsecase) SendTelegramMessage(request *models.SendTelegramRequest) (response interface{}, err error) {
	channel, err := o.query.GetClientChannel(request.ClientID, "Telegram")
	if err != nil {
		return nil, err
	}
	vendor, err := vendors.GetService(o.session, channel.Vendor.Code)
	if err != nil {
		return nil, err
	}

	sessions, err := o.query.GetSession(request.ClientID)
	if err != nil {
		return nil, err
	}

	return vendor.TelegramSendMessage(vendors.CredentialObject{
		ClientID: request.ClientID,
		Channel:  "Telegram",
		Token:    sessions.Token,
	}, request)
}
