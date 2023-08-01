package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"sync"
)

var (
	APP_DB *gorm.DB
	APP_VP *viper.Viper
	lock   sync.RWMutex
)
