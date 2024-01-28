package initialize

import (
	"path"
	"thingue-launcher/client/global"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitGorm() {
	db, err := gorm.Open(sqlite.Open(path.Join(constants.SaveDir, "config.db")), &gorm.Config{})
	if err != nil {
		logger.Zap.Panic("failed to connect database")
	}
	global.AppDB = db
	if err = db.AutoMigrate(
		&model.RemoteServer{},
		&model.InstanceConfig{},
	); err != nil {
		logger.Zap.Panic(err)
	}
}
