package usecase

import (
	"github.com/ewinjuman/go-lib/session"
	"qontak_integration/app/domain/queries"
	"qontak_integration/app/models"
	"qontak_integration/app/vendors"
	"qontak_integration/platform/http/qontak"
)

type (
	OutboundUsecaseService interface {
		SendWaMessage(request *models.SendWhatsappRequest) (response interface{}, err error)
		SendWaBroadCastDirect(request *models.BroadcastDirectRequest) (response interface{}, err error)
		SendWaBroadCastDirectBulk(request *models.BroadcastDirectBulkRequest) (response interface{}, err error)
		SendInstagramMessage(request *models.SendInstagramRequest) (response interface{}, err error)
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

//
//func (o *outboundUsecase) SendWaMessageToQueue(request *models.OutboundRequest) (response interface{}, err error) {
//	redisClient, err := cache.RedisConnection()
//	if err != nil {
//		return nil, err
//	}
//
//	sessions, err := o.query.GetSession(request.ClientID)
//	if err != nil {
//		return nil, err
//	}
//
//	messageWrapper := createMessageWrapper(request.ClientID, request.Channel, sessions.Token, nil)
//
//	switch request.Channel {
//	case "Whatsapp":
//		return o.enqueueWhatsAppMessage(redisClient, request.Messages, messageWrapper)
//	case "BlastWhatsapp":
//		return o.enqueueBroadcastMessage(redisClient, request.Messages, messageWrapper, convertPayload)
//	case "BlastWhatsappHeader":
//		return o.enqueueBroadcastMessage(redisClient, request.Messages, messageWrapper, convertPayloadHeader)
//	default:
//		return nil, repository.NotFoundErr
//	}
//}
//
//func (o *outboundUsecase) SendWaMessageToQueueBulk(request *models.OutboundBulkRequest) (response interface{}, err error) {
//	redis, err := cache.RedisConnection()
//	if err != nil {
//		return nil, err
//	}
//	sessions, err := o.query.GetSession(request.ClientID)
//	if err != nil {
//		return nil, err
//	}
//
//	messageWrapper := createMessageWrapper(request.ClientID, request.Channel, sessions.Token, nil)
//	switch request.Channel {
//	case "Whatsapp":
//		return o.enqueueBulkWhatsappMessages(redis, messageWrapper, request.Messages)
//	case "BlastWhatsapp":
//		return o.enqueueBulkBroadcastWhatsappMessages(redis, messageWrapper, request.Messages, convertPayload)
//	case "BlastWhatsappHeader":
//		return o.enqueueBulkBroadcastWhatsappMessages(redis, messageWrapper, request.Messages, convertPayloadHeader)
//	default:
//		return nil, repository.NotFoundErr
//	}
//
//	return
//}

//func createMessageWrapper(clientID int, channel, clientToken string, message interface{}) models.MessageWrapping {
//	return models.MessageWrapping{
//		ClientID:    clientID,
//		Channel:     channel,
//		ClientToken: clientToken,
//		Message:     message,
//	}
//}

//func (o *outboundUsecase) enqueueWhatsAppMessage(redisClient *redis.Storage, request models.OutMessage, messageWrapper models.MessageWrapping) (interface{}, error) {
//	message := qontak.CreateWaMassageRequest{
//		RoomID: request.From,
//		Type:   request.Type,
//		Text:   request.Content.Text,
//	}
//	messageWrapper.Message = message
//	return o.pushToQueue(redisClient, repository.SendMessageQueue, messageWrapper)
//}
//
//func (o *outboundUsecase) enqueueBroadcastMessage(redisClient *redis.Storage, request models.OutMessage, messageWrapper models.MessageWrapping, payloadConverter func(message *models.OutMessage) interface{}) (interface{}, error) {
//	messageWrapper.Message = payloadConverter(&request)
//	return o.pushToQueue(redisClient, repository.BroadcastQueue, messageWrapper)
//}
//
//func (o *outboundUsecase) enqueueBulkWhatsappMessages(redisClient *redis.Storage, messageWrapper models.MessageWrapping, messages []models.OutMessage) (interface{}, error) {
//	for _, msg := range messages {
//		message := qontak.CreateWaMassageRequest{
//			RoomID: msg.From,
//			Type:   msg.Type,
//			Text:   msg.Content.Text,
//		}
//		messageWrapper.Message = message
//		if _, err := o.pushToQueue(redisClient, repository.SendMessageQueue, messageWrapper); err != nil {
//			return nil, err
//		}
//	}
//	return nil, nil
//}
//
//func (o *outboundUsecase) enqueueBulkBroadcastWhatsappMessages(redisClient *redis.Storage, messageWrapper models.MessageWrapping, messages []models.OutMessage, payloadConverter func(message *models.OutMessage) interface{}) (interface{}, error) {
//	for _, msg := range messages {
//		messageWrapper.Message = payloadConverter(&msg)
//		if _, err := o.pushToQueue(redisClient, repository.BroadcastQueue, messageWrapper); err != nil {
//			return nil, err
//		}
//	}
//	return nil, nil
//}

//func (o *outboundUsecase) pushToQueue(redisClient *redis.Storage, queueName string, messageWrapper models.MessageWrapping) (interface{}, error) {
//	err := redisClient.Conn().RPush(context.Background(), queueName, convert.ObjectToString(messageWrapper)).Err()
//	if err != nil {
//		return nil, err
//	}
//	fmt.Printf("Message '%s' added to queue '%s'\n", convert.ObjectToString(messageWrapper), queueName)
//	return nil, nil
//}
//
//func convertPayload(in *models.OutMessage) interface{} {
//	if in.Components != nil {
//		return convertPayloadHeader(in)
//	}
//	var payloadQontak qontak.QontakMessageWABlast
//
//	payloadQontak.ToName = in.To
//	payloadQontak.ToNumber = in.To
//	payloadQontak.MessageTemplateID = in.WATemplateID
//	payloadQontak.ChannelIntegrationID = in.From
//	if in.Type == "hsm" {
//		payloadQontak.Language.Code = in.Hsm.Lang.Code
//
//		var parameterBody []qontak.QontakBodyItem
//		var prms models.Params
//		if in.Hsm.Localizable != nil {
//			for i, val := range in.Hsm.Localizable {
//				prms.Default = val.Default
//				parameterBody = append(parameterBody, qontak.QontakBodyItem{
//					Key:       fmt.Sprintf("%d", i+1),
//					ValueText: prms.Default,
//					Value:     fmt.Sprintf("key%d", i+1),
//				})
//			}
//		}
//
//		if len(parameterBody) > 0 {
//			payloadQontak.Parameters.Body = parameterBody
//		}
//		//jsonMsg, _ = json.Marshal(payloadQontak)
//	}
//	return payloadQontak
//}

//func convertPayloadHeader(in *models.OutMessage) interface{} {
//	var params models.Params
//	var payloadQontak qontak.QontakMessageWABlastWithHeader
//	//var jsonMsg []byte
//
//	payloadQontak.ToName = in.To
//	payloadQontak.ToNumber = in.To
//	payloadQontak.MessageTemplateID = in.WATemplateID
//	payloadQontak.ChannelIntegrationID = in.From
//	if in.Type == "hsm" {
//		payloadQontak.Language.Code = in.Hsm.Lang.Code
//
//		var parameterHeader *qontak.QontakHeaderItem
//		parameterBody := make([]qontak.QontakBodyItem, 0)
//		parameterButtons := make([]qontak.QontakButton, 0)
//
//		if in.Hsm.Localizable != nil {
//			for i, val := range in.Hsm.Localizable {
//				params.Default = val.Default
//				parameterBody = append(parameterBody, qontak.QontakBodyItem{
//					Key:       fmt.Sprintf("%d", i+1),
//					ValueText: params.Default,
//					Value:     fmt.Sprintf("key%d", i+1),
//				})
//			}
//		}
//
//		if in.Components != nil {
//			for _, val := range in.Components {
//				if val.Type == "header" {
//					if len(val.Parameters) != 0 {
//						if val.Parameters[0].Type == "document" {
//							parameterHeader = &qontak.QontakHeaderItem{
//								Format: "DOCUMENT",
//								Params: []qontak.QontakHeaderParam{
//									{Key: "url", Value: val.Parameters[0].Document.Link},
//									{Key: "filename", Value: val.Parameters[0].Document.Filename},
//								},
//							}
//						} else if val.Parameters[0].Type == "image" {
//							parsedURL, _ := url.Parse(val.Parameters[0].Image.Url)
//							fileName := path.Base(parsedURL.Path)
//							parameterHeader = &qontak.QontakHeaderItem{
//								Format: "IMAGE",
//								Params: []qontak.QontakHeaderParam{
//									{Key: "url", Value: val.Parameters[0].Image.Url},
//									{Key: "filename", Value: fileName},
//								},
//							}
//						} else {
//							parsedURL, _ := url.Parse(val.Parameters[0].Video.Url)
//							fileName := path.Base(parsedURL.Path)
//							parameterHeader = &qontak.QontakHeaderItem{
//								Format: "VIDEO",
//								Params: []qontak.QontakHeaderParam{
//									{Key: "url", Value: val.Parameters[0].Video.Url},
//									{Key: "filename", Value: fileName},
//								},
//							}
//						}
//					}
//				}
//
//				if val.Type == "button" {
//					if len(val.Parameters) > 0 {
//						idxStr := strconv.FormatInt(int64(val.Index), 10) // get index from payload body
//						parameterButtons = []qontak.QontakButton{{
//							Index: idxStr,
//							Type:  val.Subtype,
//							Value: val.Parameters[0].Text,
//						}}
//					} else if in.Interactive != nil {
//						parameterButtons = []qontak.QontakButton{
//							{Index: "0", Type: "url", Value: *in.Interactive.Body.Text},
//						}
//					}
//				}
//			}
//		}
//
//		if len(parameterBody) > 0 {
//			payloadQontak.Parameters.Body = parameterBody
//		}
//
//		if parameterHeader != nil {
//			payloadQontak.Parameters.Header = parameterHeader
//		}
//
//		if len(parameterButtons) > 0 {
//			payloadQontak.Parameters.Buttons = parameterButtons
//		}
//
//	}
//
//	return payloadQontak
//}
