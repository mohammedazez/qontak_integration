package usecase

import (
	"fmt"
	"github.com/ewinjuman/go-lib/session"
	"qontak_integration/app/domain/entities"
	"qontak_integration/app/domain/queries"
	"qontak_integration/app/models"
	"qontak_integration/platform/http/qontak"
	"sync"
)

type (
	TemplateUsecaseService interface {
		GetTemplate(request *models.GetTemplateRequest) (interface{}, error)
	}

	templateUsecase struct {
		session    *session.Session
		query      queries.QueriesService
		qontakHttp qontak.QontakHttpService
	}
)

func NewTemplateUsecase(session *session.Session) (item TemplateUsecaseService) {
	return &templateUsecase{
		session:    session,
		query:      queries.NewQueries(session),
		qontakHttp: qontak.NewQontakHttp(session),
	}
}

func (o *templateUsecase) GetTemplate(request *models.GetTemplateRequest) (interface{}, error) {
	var wg sync.WaitGroup
	var channel entities.ClientChannel
	var session entities.Sessions
	var err1, err2 error

	wg.Add(2)
	go func() {
		defer wg.Done()
		channel, err1 = o.query.GetClientChannel(request.ClientID, "ListWhatsappTemplate")
	}()

	go func() {
		defer wg.Done()
		session, err2 = o.query.GetSession(request.ClientID)
	}()
	wg.Wait()

	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}

	requestHttp := qontak.GetTemplateRequest{
		Limit:  request.Limit,
		Offset: request.Offset,
	}
	resp, err := o.qontakHttp.WaTemplate(requestHttp, fmt.Sprintf("Bearer %s", session.Token))
	if err != nil {
		return nil, err
	}
	response := models.GetTemplateResponse{
		Data: resp.Data,
		Meta: resp.Meta,
	}
	response.ClientVendor.ChannelID = channel.ChannelID

	return response, nil
}
