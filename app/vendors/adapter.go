package vendors

import (
	"context"
	"fmt"
	"qontak_integration/app/domain/queries"
	"qontak_integration/app/models"
	"qontak_integration/pkg/repository"
	"qontak_integration/platform/http/qontak"
	"strings"

	Error "github.com/ewinjuman/go-lib/error"
	"github.com/ewinjuman/go-lib/helper/convert"
	"github.com/ewinjuman/go-lib/session"
	"github.com/gofiber/storage/redis/v3"
)

type (
	VendorClient interface {
		SendMessage(payload interface{}, credential Credential) (result interface{}, err error)
		InstagramSendMessage(credentialObject CredentialObject, request *models.SendInstagramRequest) (response interface{}, err error)
		TelegramSendMessage(credentialObject CredentialObject, request *models.SendTelegramRequest) (response interface{}, err error)
		WaSendMessage(CredentialObject CredentialObject, request *models.SendWhatsappRequest) (response interface{}, err error)
		WaBroadcastDirect(CredentialObject CredentialObject, request *models.BroadcastDirectRequest) (response interface{}, err error)
		WaBroadcastDirectBulk(CredentialObject CredentialObject, request *models.BroadcastDirectBulkRequest) (response interface{}, err error)
	}
)

func GetService(session *session.Session, vendor string) (VendorClient, error) {
	vendor = strings.ToUpper(vendor)
	switch vendor {
	case "QONTAK":
		return &QontakVendor{
			session:    session,
			query:      queries.NewQueries(session),
			qontakHttp: qontak.NewQontakHttp(session),
		}, nil
	case "KOMMO":
		return &KommoVendor{
			session: session,
			query:   queries.NewQueries(session),
		}, nil
	default:
		return nil, Error.NewError(repository.BadRequestCode, repository.FailedStatus, "invalid Vendor")
	}
}

func createMessageWrapper(clientID int, vendorCode, channel, clientToken string, message interface{}) models.MessageWrapping {
	return models.MessageWrapping{
		ClientID:    clientID,
		Channel:     channel,
		ClientToken: clientToken,
		Message:     message,
		VendorCode:  vendorCode,
	}
}

func pushToQueue(redisClient *redis.Storage, queueName string, messageWrapper models.MessageWrapping) (interface{}, error) {
	err := redisClient.Conn().RPush(context.Background(), queueName, convert.ObjectToString(messageWrapper)).Err()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Message '%s' added to queue '%s'\n", convert.ObjectToString(messageWrapper), queueName)
	return nil, nil
}
