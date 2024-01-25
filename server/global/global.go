package global

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"gorm.io/gorm"
)

var (
	ServerDB   *gorm.DB
	StorageDB  *gorm.DB
	MqttServer *mqtt.Server
	//SgccService *service.SgccService
)
