package initialize

import (
	"thingue-launcher/client/global"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitGorm() {
	db, err := gorm.Open(sqlite.Open(constants.SAVE_DIR+"config.db"), &gorm.Config{})
	if err != nil {
		logger.Zap.Panic("failed to connect database")
	}
	global.APP_DB = db
	if err = db.AutoMigrate(
		&model.RemoteServer{},
		&model.ClientInstance{},
	); err != nil {
		logger.Zap.Error(err)
	}
}
