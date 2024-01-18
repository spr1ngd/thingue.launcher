package global

import (
	"gorm.io/gorm"
)

var (
	AppDB    *gorm.DB
	ClientId uint
)
