package server

import (
	"github.com/spf13/viper"
	"golang.org/x/net/websocket"
	"gorm.io/gorm"
	"sync"
)

var (
	APP_DB        *gorm.DB
	APP_VP        *viper.Viper
	WS            *websocket.Conn
	lock          sync.RWMutex
	APP_VERSION   string
	APP_GITCOMMIT string
	APP_BUILDDATE string
)

func SetAppVersion(version, gitCommit, buildDate string) {
	APP_VERSION = version
	APP_GITCOMMIT = gitCommit
	APP_BUILDDATE = buildDate
}
