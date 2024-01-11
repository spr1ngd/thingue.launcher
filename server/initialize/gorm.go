package initialize

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/model"
	"thingue-launcher/server/global"
)

func initServerDB() {
	dsn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	global.SERVER_DB = db
	if err = db.AutoMigrate(
		&model.Client{},
		&model.ServerInstance{},
	); err != nil {
		logger.Zap.Error(err)
	}
}

func initStorageDB() {
	dsn := constants.SAVE_DIR + "storage.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	global.STORAGE_DB = db
	if err = db.AutoMigrate(
		&model.CloudFile{},
		&model.CloudRes{},
	); err != nil {
		logger.Zap.Error(err)
	}
}
