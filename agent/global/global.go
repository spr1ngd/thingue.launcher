package global

import (
	"github.com/spf13/viper"
	"golang.org/x/net/websocket"
	"gorm.io/gorm"
	"sync"
)

var (
	APP_DB *gorm.DB
	APP_VP *viper.Viper
	WS     *websocket.Conn
	lock   sync.RWMutex
)
