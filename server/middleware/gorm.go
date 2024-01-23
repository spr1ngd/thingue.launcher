package middleware

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"thingue-launcher/common/model"
	"thingue-launcher/server/global"
)

var GormInitialed bool

func InitGorm() {
	if !GormInitialed {
		openServerDB()
		openStorageDB()
		global.ServerDB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Client{})
		global.ServerDB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.ServerInstance{})
	}
}

func openServerDB() {
	dsn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	global.ServerDB = db
	if err = db.AutoMigrate(
		&model.Client{},
		&model.ServerInstance{},
	); err != nil {
		fmt.Println(err)
	}
}

func openStorageDB() {
	dsn := "./thingue-launcher/storage.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	global.StorageDB = db
	if err = db.AutoMigrate(
		&model.CloudFile{},
		&model.CloudRes{},
	); err != nil {
		fmt.Println(err)
	}
}
