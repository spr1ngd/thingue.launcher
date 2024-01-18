package global

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"gorm.io/gorm"
	"thingue-launcher/server/sgcc/service"
)

var (
	ServerDB    *gorm.DB
	StorageDB   *gorm.DB
	MqttServer  *mqtt.Server
	SgccService *service.SgccService
)
