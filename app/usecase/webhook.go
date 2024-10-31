package usecase

import (
	"encoding/json"
	"github.com/ewinjuman/go-lib/helper/convert"
	"github.com/ewinjuman/go-lib/session"
	"gorm.io/datatypes"
	"qontak_integration/app/domain/entities"
	"qontak_integration/app/domain/queries"
	"qontak_integration/app/models"
	"qontak_integration/pkg/utils"
	"qontak_integration/platform/http/qontak"
)

type (
	WebhookUsecaseService interface {
		ListWebhook() (response interface{}, err error)
		GetWebhook(ID int) (interface{}, error)
		SaveWebhook(request *models.WebhookResult) (response interface{}, err error)
	}

	webhookUsecase struct {
		session    *session.Session
		query      queries.QueriesService
		qontakHttp qontak.QontakHttpService
	}
)

func NeWebhookUsecase(session *session.Session) (item WebhookUsecaseService) {
	//userGrpc, err := user.NewServerContext()
	return &webhookUsecase{
		session:    session,
		query:      queries.NewQueries(session),
		qontakHttp: qontak.NewQontakHttp(session),
	}
}

func (w *webhookUsecase) GetWebhook(id int) (response interface{}, err error) {
	result, err := w.query.GetWebhook(id)
	webhook := models.WebhookResult{}
	event := []models.Event{}

	json.Unmarshal(result.Events, &event)

	utils.ObjectToObject(result, &webhook)
	webhook.Events = event
	response = webhook
	return
}

func (w *webhookUsecase) ListWebhook() (response interface{}, err error) {
	result, err := w.query.ListWebhook()
	webhooks := []models.WebhookResult{}
	for _, e := range result {
		webhook := models.WebhookResult{}
		event := []models.Event{}

		json.Unmarshal(e.Events, &event)

		convert.ObjectToObject(e, &webhook)
		webhook.Events = event

		webhooks = append(webhooks, webhook)
	}

	response = webhooks
	return
}

func (w *webhookUsecase) SaveWebhook(request *models.WebhookResult) (response interface{}, err error) {
	//result, err := w.query.ListWebhook()

	webhook := entities.Webhook{}
	//event := []models.Event{}

	convert.ObjectToObject(request, &webhook)

	jsonData, err := json.Marshal(request.Events)
	if err != nil {
		return
	}

	jsonDatatypes := datatypes.JSON(jsonData)
	webhook.Events = jsonDatatypes

	result, err := w.query.SaveWebhook(webhook)

	event := []models.Event{}
	webhookresult := models.WebhookResult{}
	json.Unmarshal(result.Events, &event)
	convert.ObjectToObject(result, &webhookresult)
	webhookresult.Events = event
	response = webhookresult

	return
}
