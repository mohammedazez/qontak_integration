package vendors

import (
	"github.com/ewinjuman/go-lib/session"
	"qontak_integration/app/domain/queries"
	"qontak_integration/app/models"
)

type (
	KommoVendor struct {
		session *session.Session
		query   queries.QueriesService
	}

	//komoVendor struct {
	//	session *session.Session
	//	query   queries.QueriesService
	//}
)

func (w *KommoVendor) SendMessage(payload interface{}, credential Credential) (result interface{}, err error) {

	return
}

func (w *KommoVendor) WaSendMessage(CredentialObject CredentialObject, request *models.SendWhatsappRequest) (response interface{}, err error) {
	// Implementasi untuk kirim pesan melalui WhatsApp
	panic("implement me")
}

func (w *KommoVendor) WaBroadcastDirect(CredentialObject CredentialObject, request *models.BroadcastDirectRequest) (response interface{}, err error) {
	// Implementasi untuk kirim pesan melalui WhatsApp
	panic("implement me")
}

func (w *KommoVendor) WaBroadcastDirectBulk(CredentialObject CredentialObject, request *models.BroadcastDirectBulkRequest) (response interface{}, err error) {
	// Implementasi untuk kirim pesan melalui WhatsApp
	panic("implement me")
}

func (w *KommoVendor) InstagramSendMessage(credentialObject CredentialObject, request *models.SendInstagramRequest) (response interface{}, err error) {
	panic("implement me")
}
