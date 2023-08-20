package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"sync"
)

var (
	APP_DB        *gorm.DB
	APP_VP        *viper.Viper
	lock          sync.RWMutex
	APP_VERSION   string
	APP_GITCOMMIT string
	APP_BUILDDATE string
	NODE_ID       uint
)

func SetAppVersion(version, gitCommit, buildDate string) {
	APP_VERSION = version
	APP_GITCOMMIT = gitCommit
	APP_BUILDDATE = buildDate
}
