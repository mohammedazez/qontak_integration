package qontak

import (
	"encoding/json"
	"fmt"
	"net/http"
	"qontak_integration/pkg/utils"
	"strconv"
	"strings"

	"qontak_integration/app/domain/entities"
	"qontak_integration/app/domain/queries"
	"qontak_integration/pkg/configs"
	"qontak_integration/pkg/repository"

	Error "github.com/ewinjuman/go-lib/error"
	"github.com/ewinjuman/go-lib/helper/convert"
	Rest "github.com/ewinjuman/go-lib/http"
	Session "github.com/ewinjuman/go-lib/session"
)

type (
	QontakHttpService interface {
		CreateInstagramMessage(request interface{}, token string) (response CreateMessageResponse, err error)
		CreateWaMessage(request interface{}, token string) (response CreateMessageResponse, err error)
		CreateTelegramMessage(request interface{}, token string) (response CreateMessageResponse, err error)
		WaBroadcastDirect(request interface{}, token string) (response BroadcastDirectResponse, err error)
		WaTemplate(request GetTemplateRequest, token string) (response GeneralResponse, err error)
		Resolved(request interface{}, roomID, token string) (response GeneralResponse, err error)
	}

	qontakHttp struct {
		session      *Session.Session
		qontakRest   Rest.RestClient
		qontakConfig configs.Qontak
	}
)

// New creates a new http request to Qontak
func NewQontakHttp(session *Session.Session) QontakHttpService {
	option := Rest.Options{}
	convert.ObjectToObject(configs.Config.Qontak.Option, &option)
	return &qontakHttp{
		session:      session,
		qontakRest:   Rest.New(option),
		qontakConfig: configs.Config.Qontak,
	}
}

func qontakHeaders(token string, contentType string) http.Header {
	headers := http.Header{}
	headers.Set("Content-Type", contentType)
	headers.Set("Authorization", token)
	return headers
}

func (o *qontakHttp) handleApiLog(url string, method string, headers http.Header, requestBody interface{}, responseBody []byte, httpStatus int) {
	apiLog := entities.ApiLogs{
		Url:            url,
		Method:         method,
		RequestHeader:  fmt.Sprintf("%v", headers),
		RequestBody:    utils.ObjectToString(requestBody),
		ResponseStatus: strconv.Itoa(httpStatus),
		ResponseHeader: fmt.Sprintf("%v", headers),
		ResponseBody:   string(responseBody),
	}

	q := queries.NewQueries(o.session)
	go q.InsertApiLog(apiLog)
}

func (o *qontakHttp) executeRequest(method, path, token string, requestBody interface{}, params map[string]string, contentType string) (result []byte, httpStatus int, err error) {
	headers := qontakHeaders(token, contentType)
	return o.qontakRest.Execute(o.session, o.qontakConfig.Host, path, method, headers, requestBody, params, nil)
}

func (o *qontakHttp) handleError(result []byte, httpStatus int) error {
	if httpStatus != http.StatusOK && httpStatus != http.StatusCreated {
		var errResponse ErrorResponse
		if err := json.Unmarshal(result, &errResponse); err == nil {
			msg := "Error hit to Qontak"
			if errResponse.Error.Code != 0 {
				msg = errResponse.Error.Messages[0]
			}
			return Error.New(httpStatus, repository.FailedStatus, msg)
		}
	}
	return nil
}

func (o *qontakHttp) doHttpRequest(method, path string, token string, requestBody interface{}, params map[string]string, responseBody interface{}, contentType string) error {
	result, httpStatus, err := o.executeRequest(method, path, token, requestBody, params, contentType)
	if httpStatus == 0 {
		httpStatus = http.StatusInternalServerError
	}
	if err != nil {
		o.handleApiLog(o.qontakConfig.Host+path, method, qontakHeaders(token, contentType), requestBody, []byte(err.Error()), httpStatus)
		return err
	}

	o.handleApiLog(o.qontakConfig.Host+path, method, qontakHeaders(token, contentType), requestBody, result, httpStatus)

	if err = o.handleError(result, httpStatus); err != nil {
		return err
	}

	return json.Unmarshal(result, responseBody)
}

func (o *qontakHttp) WaTemplate(request GetTemplateRequest, token string) (response GeneralResponse, err error) {
	params := map[string]string{
		"cursor":     request.Cursor,
		"hsm_chat":   strconv.FormatBool(request.HSMChat),
		"is_counted": strconv.FormatBool(request.IsCounted),
		"limit":      strconv.Itoa(request.Limit),
		"offset":     strconv.Itoa(request.Offset),
		"query":      request.Query,
		"status":     request.Status,
	}

	err = o.doHttpRequest(http.MethodGet, o.qontakConfig.Path.WaTemplateList, token, nil, params, &response, "application/json")
	return
}

func (o *qontakHttp) CreateWaMessage(request interface{}, token string) (response CreateMessageResponse, err error) {
	err = o.doHttpRequest(http.MethodPost, o.qontakConfig.Path.WaSendMessage, token, request, nil, &response, "multipart/form-data")
	return
}

func (o *qontakHttp) WaBroadcastDirect(request interface{}, token string) (response BroadcastDirectResponse, err error) {
	err = o.doHttpRequest(http.MethodPost, o.qontakConfig.Path.WaBroadcastDirect, token, request, nil, &response, "application/json")
	return
}

func (o *qontakHttp) Resolved(request interface{}, roomID, token string) (response GeneralResponse, err error) {
	path := strings.Replace(o.qontakConfig.Path.Resolved, "{room_id}", roomID, -1)
	err = o.doHttpRequest(http.MethodPut, path, token, request, nil, &response, "application/json")
	return
}

func (o *qontakHttp) CreateInstagramMessage(request interface{}, token string) (response CreateMessageResponse, err error) {
	err = o.doHttpRequest(http.MethodPost, o.qontakConfig.Path.WaSendMessage, token, request, nil, &response, "multipart/form-data")
	return
}

func (o *qontakHttp) CreateTelegramMessage(request interface{}, token string) (response CreateMessageResponse, err error) {
	err = o.doHttpRequest(http.MethodPost, o.qontakConfig.Path.TelegramSendMessage, token, request, nil, &response, "multipart/form-data")
	return
}
