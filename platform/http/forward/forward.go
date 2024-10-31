package forward

import (
	"fmt"
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/ewinjuman/go-lib/helper/convert"
	Rest "github.com/ewinjuman/go-lib/http"
	Session "github.com/ewinjuman/go-lib/session"
	"net/http"
	"net/url"
	"qontak_integration/app/models"
	"qontak_integration/pkg/configs"
	"qontak_integration/pkg/repository"
)

type (
	ForwardStatusHttpService interface {
		Forward(request interface{}, webhook models.WebhookResult) (response interface{}, err error)
	}

	forwardHttp struct {
		session       *Session.Session
		forwardRest   Rest.RestClient
		forwardConfig configs.ForwardOption
	}
)

func NewForwardHttp(session *Session.Session) ForwardStatusHttpService {
	option := Rest.Options{}
	convert.ObjectToObject(configs.Config.ForwardOption.Option, &option)
	return &forwardHttp{
		session:       session,
		forwardRest:   Rest.New(option),
		forwardConfig: configs.Config.ForwardOption,
	}
}

func (f *forwardHttp) Forward(request interface{}, webhook models.WebhookResult) (response interface{}, err error) {
	parsedURL, err := url.Parse(webhook.Url)
	if err != nil {
		err = Error.New(repository.BadRequestCode, repository.FailedStatus, err.Error())
		return
	}

	// Get host and path
	host := parsedURL.Scheme + "://" + parsedURL.Host
	path := parsedURL.Path

	header := http.Header{}
	if webhook.Token != "" {
		header.Add("Authorization", webhook.Token)
	}

	result, httpStatus, err := f.forwardRest.Execute(f.session, host, path, webhook.Method, header, request, nil, nil)
	//do something if err is not nil
	if err != nil {
		if Error.IsTimeout(err) {
			//.... do something if needed
		}
		return
	}

	if httpStatus != webhook.ExpectedHttpCode {
		return result, Error.New(httpStatus, repository.FailedStatus, fmt.Sprintf("http status code %d is not expected", httpStatus))
	}
	return
}
