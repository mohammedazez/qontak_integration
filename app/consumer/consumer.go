package consumer

import (
	"context"
	"fmt"
	"github.com/ewinjuman/go-lib/helper/convert"
	Logger "github.com/ewinjuman/go-lib/logger"
	"github.com/ewinjuman/go-lib/session"
	Session "github.com/ewinjuman/go-lib/session"
	"golang.org/x/time/rate"
	"log"
	"os"
	"qontak_integration/app/models"
	"qontak_integration/app/vendors"
	"qontak_integration/pkg/configs"
	"qontak_integration/pkg/repository"
	"qontak_integration/platform/cache"
	"time"
)

var ctx = context.Background()

const requestsPerMinute = 60
const burstLimit = 1 // Max burst size

// Rate limiter setup
var limiter = rate.NewLimiter(rate.Every(time.Minute/requestsPerMinute), burstLimit)

func OutboundSendMessageConsumer(session *session.Session) {
	println("start consumer")
	redis, errC := cache.RedisConnection()
	if errC != nil {
		session.Error(errC.Error())
	}
	for {
		result, err := redis.Conn().BRPop(context.Background(), 0*time.Second, repository.SendMessageQueue).Result()
		if err != nil {
			log.Printf("Error saat menerima pesan: %v", err)
			continue
		}

		session = Session.New(Logger.New(configs.Config.Logger)).
			SetInstitutionID(configs.Config.Apps.Name).
			SetAppName(configs.Config.Apps.Name).
			SetURL("/send-message-consumer").
			SetMethod("QUEUE")
		message := result[1] // Pesan yang diterima
		session.Info(fmt.Sprintf("Pesan diterima dari Redis Queue '%s': %s", repository.SendMessageQueue, message))

		var request models.MessageWrapping
		convert.StringToObject(message, &request)

		_, _ = sendMessage(session, request)
	}
}

func OutboundBroadcastConsumer(session *session.Session) {
	redis, errC := cache.RedisConnection()
	if errC != nil {
		session.Error(errC.Error())
		os.Exit(1)
	}
	for {
		result, err := redis.Conn().BRPop(context.Background(), 0*time.Second, repository.BroadcastQueue).Result()
		if err != nil {
			log.Printf("Error saat menerima pesan: %v", err)
			continue
		}

		session = Session.New(Logger.New(configs.Config.Logger)).
			SetInstitutionID(configs.Config.Apps.Name).
			SetAppName(configs.Config.Apps.Name).
			SetURL("/send-broadcast-message-consumer").
			SetMethod("QUEUE")
		message := result[1] // Pesan yang diterima
		session.Info(fmt.Sprintf("Pesan diterima dari Redis Queue '%s': %s", repository.BroadcastQueue, message))

		var request models.MessageWrapping
		convert.StringToObject(message, &request)
		limiter.Wait(ctx)
		_, _ = sendMessage(session, request)
	}
}
func sendMessage(session *session.Session, wrapping models.MessageWrapping) (result interface{}, err error) {
	vendor, _ := vendors.GetService(session, wrapping.VendorCode)
	result, err = vendor.SendMessage(wrapping.Message, vendors.Credential{
		Channel: wrapping.Channel,
		Token:   wrapping.ClientToken,
	})
	return
}
