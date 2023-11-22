package global

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"gorm.io/gorm"
)

var (
	SERVER_DB   *gorm.DB
	STORAGE_DB  *gorm.DB
	MQTT_SERVER *mqtt.Server
)
