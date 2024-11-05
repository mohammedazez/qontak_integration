package vendors

import (
	"fmt"
	"log"
	"net/url"
	"path"
	"qontak_integration/app/domain/queries"
	"qontak_integration/app/models"
	"qontak_integration/pkg/repository"
	"qontak_integration/platform/cache"
	"qontak_integration/platform/http/qontak"
	"strconv"
	"sync"

	"github.com/ewinjuman/go-lib/session"
	"github.com/gofiber/storage/redis/v3"
)

type (
	QontakVendor struct {
		session    *session.Session
		query      queries.QueriesService
		qontakHttp qontak.QontakHttpService
	}
)

func (w *QontakVendor) SendMessage(payload interface{}, credential Credential) (result interface{}, err error) {
	switch credential.Channel {
	case "Whatsapp":
		return w.qontakHttp.CreateWaMessage(payload, credential.Token)
	case "Instagram":
		return w.qontakHttp.CreateInstagramMessage(payload, credential.Token)
	case "Telegram":
		return w.qontakHttp.CreateTelegramMessage(payload, credential.Token)
	case "BlastWhatsapp":
		return w.qontakHttp.WaBroadcastDirect(payload, credential.Token)
	case "BlastWhatsappHeader":
		return w.qontakHttp.WaBroadcastDirect(payload, credential.Token)
	default:
		return nil, fmt.Errorf("Unsupported channel")
	}
}

func (w *QontakVendor) WaSendMessage(credentialObject CredentialObject, request *models.SendWhatsappRequest) (response interface{}, err error) {
	qontakRequest := qontak.CreateMassageRequest{
		RoomID: request.To,
		Type:   request.Message.Type,
		Text:   request.Message.Content.Text,
	}
	redisClient, err := cache.RedisConnection()
	if err != nil {
		return nil, err
	}
	messageWrapper := createMessageWrapper(request.ClientID, "QONTAK", "Whatsapp", credentialObject.Token, qontakRequest)
	return pushToQueue(redisClient, repository.SendMessageQueue, messageWrapper)
}

func (w *QontakVendor) WaBroadcastDirect(credentialObject CredentialObject, request *models.BroadcastDirectRequest) (response interface{}, err error) {
	messageWrapper := createMessageWrapper(credentialObject.ClientID, "QONTAK", credentialObject.Channel, credentialObject.Token, nil)
	redisClient, err := cache.RedisConnection()
	if err != nil {
		return nil, err
	}

	switch request.Channel {
	case "BlastWhatsapp":
		return w.enqueueBroadcastMessage(redisClient, request.Messages, messageWrapper, w.convertPayload)
	case "BlastWhatsappHeader":
		return w.enqueueBroadcastMessage(redisClient, request.Messages, messageWrapper, w.convertPayloadHeader)
	default:
		return nil, repository.NotFoundErr
	}
}

func (w *QontakVendor) WaBroadcastDirectBulk(CredentialObject CredentialObject, request *models.BroadcastDirectBulkRequest) (response interface{}, err error) {
	redis, err := cache.RedisConnection()
	if err != nil {
		return nil, err
	}
	messageWrapper := createMessageWrapper(CredentialObject.ClientID, "QONTAK", CredentialObject.Channel, CredentialObject.Token, nil)
	switch request.Channel {
	case "BlastWhatsapp":
		go w.enqueueBulkBroadcastWhatsappMessagesWithChannel(redis, messageWrapper, request.Messages, w.convertPayload)
		return
	case "BlastWhatsappHeader":
		go w.enqueueBulkBroadcastWhatsappMessagesWithChannel(redis, messageWrapper, request.Messages, w.convertPayloadHeader)
		return
	default:
		return nil, repository.NotFoundErr
	}
}

func (w *QontakVendor) InstagramSendMessage(credentialObject CredentialObject, request *models.SendInstagramRequest) (response interface{}, err error) {
	qontakRequest := qontak.CreateMassageRequest{
		RoomID: request.To,
		Type:   request.Message.Type,
		Text:   request.Message.Content.Text,
	}
	redisClient, err := cache.RedisConnection()
	if err != nil {
		return nil, err
	}
	messageWrapper := createMessageWrapper(request.ClientID, "QONTAK", "Whatsapp", credentialObject.Token, qontakRequest)
	return pushToQueue(redisClient, repository.SendMessageQueue, messageWrapper)
}

func (w *QontakVendor) TelegramSendMessage(credentialObject CredentialObject, request *models.SendTelegramRequest) (response interface{}, err error) {
	qontakRequest := qontak.CreateMassageRequest{
		RoomID: request.To,
		Type:   request.Message.Type,
		Text:   request.Message.Content.Text,
	}
	redisClient, err := cache.RedisConnection()
	if err != nil {
		return nil, err
	}
	messageWrapper := createMessageWrapper(request.ClientID, "QONTAK", "Telegram", credentialObject.Token, qontakRequest)
	return pushToQueue(redisClient, repository.SendMessageQueue, messageWrapper)
}

// Util====================================
func (o *QontakVendor) enqueueBroadcastMessage(redisClient *redis.Storage, request models.OutMessage, messageWrapper models.MessageWrapping, payloadConverter func(message *models.OutMessage) interface{}) (interface{}, error) {
	messageWrapper.Message = payloadConverter(&request)
	return pushToQueue(redisClient, repository.BroadcastQueue, messageWrapper)
}
func producer(data []models.OutMessage, ch chan models.OutMessage) {
	defer close(ch)
	for _, item := range data {
		ch <- item
	}
}

func consumer(ch chan models.OutMessage, wg *sync.WaitGroup, redisClient *redis.Storage, messageWrapper models.MessageWrapping, payloadConverter func(message *models.OutMessage) interface{}) {
	defer wg.Done()
	for item := range ch {
		messageWrapper.Message = payloadConverter(&item)
		if _, err := pushToQueue(redisClient, repository.BroadcastQueue, messageWrapper); err != nil {
			log.Println(err.Error())
		}
	}
}

func (o *QontakVendor) enqueueBulkBroadcastWhatsappMessagesWithChannel(redisClient *redis.Storage, messageWrapper models.MessageWrapping, messages []models.OutMessage, payloadConverter func(message *models.OutMessage) interface{}) (interface{}, error) {
	ch := make(chan models.OutMessage, 100) // buffer channel
	var wg sync.WaitGroup

	wg.Add(1)
	go consumer(ch, &wg, redisClient, messageWrapper, payloadConverter)
	go producer(messages, ch)
	wg.Wait()
	return nil, nil
}
func (o *QontakVendor) enqueueBulkBroadcastWhatsappMessages(redisClient *redis.Storage, messageWrapper models.MessageWrapping, messages []models.OutMessage, payloadConverter func(message *models.OutMessage) interface{}) (interface{}, error) {
	go func() {
		for _, msg := range messages {
			messageWrapper.Message = payloadConverter(&msg)
			if _, err := pushToQueue(redisClient, repository.BroadcastQueue, messageWrapper); err != nil {
				log.Println(err.Error())
			}
		}
	}()
	return nil, nil
}
func (w *QontakVendor) convertPayload(in *models.OutMessage) interface{} {
	if in.Components != nil {
		return w.convertPayloadHeader(in)
	}
	var payloadQontak qontak.QontakMessageWABlast

	payloadQontak.ToName = in.To
	payloadQontak.ToNumber = in.To
	payloadQontak.MessageTemplateID = in.WATemplateID
	payloadQontak.ChannelIntegrationID = in.From
	if in.Type == "hsm" {
		payloadQontak.Language.Code = in.Hsm.Lang.Code

		var parameterBody []qontak.QontakBodyItem
		var prms models.Params
		if in.Hsm.Localizable != nil {
			for i, val := range in.Hsm.Localizable {
				prms.Default = val.Default
				parameterBody = append(parameterBody, qontak.QontakBodyItem{
					Key:       fmt.Sprintf("%d", i+1),
					ValueText: prms.Default,
					Value:     fmt.Sprintf("key%d", i+1),
				})
			}
		}

		if len(parameterBody) > 0 {
			payloadQontak.Parameters.Body = parameterBody
		}
		//jsonMsg, _ = json.Marshal(payloadQontak)
	}
	return payloadQontak
}

func (w *QontakVendor) convertPayloadHeader(in *models.OutMessage) interface{} {
	var params models.Params
	var payloadQontak qontak.QontakMessageWABlastWithHeader
	//var jsonMsg []byte

	payloadQontak.ToName = in.To
	payloadQontak.ToNumber = in.To
	payloadQontak.MessageTemplateID = in.WATemplateID
	payloadQontak.ChannelIntegrationID = in.From
	if in.Type == "hsm" {
		payloadQontak.Language.Code = in.Hsm.Lang.Code

		var parameterHeader *qontak.QontakHeaderItem
		parameterBody := make([]qontak.QontakBodyItem, 0)
		parameterButtons := make([]qontak.QontakButton, 0)

		if in.Hsm.Localizable != nil {
			for i, val := range in.Hsm.Localizable {
				params.Default = val.Default
				parameterBody = append(parameterBody, qontak.QontakBodyItem{
					Key:       fmt.Sprintf("%d", i+1),
					ValueText: params.Default,
					Value:     fmt.Sprintf("key%d", i+1),
				})
			}
		}

		if in.Components != nil {
			for _, val := range in.Components {
				if val.Type == "header" {
					if len(val.Parameters) != 0 {
						if val.Parameters[0].Type == "document" {
							parameterHeader = &qontak.QontakHeaderItem{
								Format: "DOCUMENT",
								Params: []qontak.QontakHeaderParam{
									{Key: "url", Value: val.Parameters[0].Document.Link},
									{Key: "filename", Value: val.Parameters[0].Document.Filename},
								},
							}
						} else if val.Parameters[0].Type == "image" {
							parsedURL, _ := url.Parse(val.Parameters[0].Image.Url)
							fileName := path.Base(parsedURL.Path)
							parameterHeader = &qontak.QontakHeaderItem{
								Format: "IMAGE",
								Params: []qontak.QontakHeaderParam{
									{Key: "url", Value: val.Parameters[0].Image.Url},
									{Key: "filename", Value: fileName},
								},
							}
						} else {
							parsedURL, _ := url.Parse(val.Parameters[0].Video.Url)
							fileName := path.Base(parsedURL.Path)
							parameterHeader = &qontak.QontakHeaderItem{
								Format: "VIDEO",
								Params: []qontak.QontakHeaderParam{
									{Key: "url", Value: val.Parameters[0].Video.Url},
									{Key: "filename", Value: fileName},
								},
							}
						}
					}
				}

				if val.Type == "button" {
					if len(val.Parameters) > 0 {
						idxStr := strconv.FormatInt(int64(val.Index), 10) // get index from payload body
						parameterButtons = []qontak.QontakButton{{
							Index: idxStr,
							Type:  val.Subtype,
							Value: val.Parameters[0].Text,
						}}
					} else if in.Interactive != nil {
						parameterButtons = []qontak.QontakButton{
							{Index: "0", Type: "url", Value: *in.Interactive.Body.Text},
						}
					}
				}
			}
		}

		if len(parameterBody) > 0 {
			payloadQontak.Parameters.Body = parameterBody
		}

		if parameterHeader != nil {
			payloadQontak.Parameters.Header = parameterHeader
		}

		if len(parameterButtons) > 0 {
			payloadQontak.Parameters.Buttons = parameterButtons
		}

	}

	return payloadQontak
}
