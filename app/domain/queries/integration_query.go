package queries

import (
	"github.com/ewinjuman/go-lib/session"
	"qontak_integration/app/domain/entities"
	"qontak_integration/pkg/repository"
	"qontak_integration/platform/database"
)

type (
	QueriesService interface {
		GetClient(ID int) (client entities.Client, err error)
		GetClientChannel(clientID int, channelName string) (client entities.ClientChannel, err error)
		GetSession(clientID int) (session entities.Sessions, err error)
		GetWebhook(ID int) (webhook entities.Webhook, err error)
		ListWebhook() (webhooks []entities.Webhook, err error)
		SaveWebhook(webhook entities.Webhook) (result entities.Webhook, err error)
		GetWebhookBychannel(trxType string) (webhooks []entities.Webhook, err error)
		InsertApiLog(data entities.ApiLogs) error
	}

	queries struct {
		session *session.Session
	}
)

func NewQueries(session *session.Session) (rep QueriesService) {
	return &queries{session: session}
}

func (q *queries) GetClient(ID int) (client entities.Client, err error) {
	db, err := database.MysqlConnection(q.session)
	if err != nil {
		return
	}
	client.ID = ID
	err = db.First(&client).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return
}

func (q *queries) GetClientChannel(clientID int, channelName string) (client entities.ClientChannel, err error) {
	db, err := database.MysqlConnection(q.session)
	if err != nil {
		return
	}
	err = db.Joins("JOIN vendors ON client_channel.vendor_id = vendors.id").
		Where("client_channel.client_id = ? and client_channel.channel = ?", clientID, channelName).
		Preload("Vendor").
		Find(&client).Error
	//err = db.Where("client_id = ? and channel = ?", clientID, channelName).First(&client).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return
}

func (q *queries) GetSession(clientID int) (session entities.Sessions, err error) {
	db, err := database.MysqlConnection(q.session)
	if err != nil {
		return
	}
	err = db.Where("client_id = ? and status = ?", clientID, 1).First(&session).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return
}

func (q *queries) GetWebhook(ID int) (webhook entities.Webhook, err error) {
	db, err := database.MysqlConnection(q.session)
	if err != nil {
		return
	}
	err = db.Where("id = ? and status = ?", ID, 1).First(&webhook).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return
}
func (q *queries) ListWebhook() (webhooks []entities.Webhook, err error) {
	db, err := database.MysqlConnection(q.session)
	if err != nil {
		return
	}
	err = db.Where("status = ?", 1).Find(&webhooks).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return
}

func (q *queries) SaveWebhook(webhook entities.Webhook) (result entities.Webhook, err error) {
	db, err := database.MysqlConnection(q.session)
	if err != nil {
		return
	}
	err = db.Save(&webhook).Scan(&result).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return
}

func (q *queries) GetWebhookBychannel(trxType string) (webhooks []entities.Webhook, err error) {
	db, err := database.MysqlConnection(q.session)
	if err != nil {
		return
	}
	err = db.Where("trx_type = ? and status = ?", trxType, 1).Find(&webhooks).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return
}

func (q *queries) InsertApiLog(data entities.ApiLogs) error {

	db, err := database.MysqlConnection(q.session)
	if err != nil {
		return err
	}
	err = db.Create(&data).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return err
	}
	return nil

}
