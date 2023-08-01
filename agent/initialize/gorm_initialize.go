package initialize

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/model"
)

func InitGorm() {
	db, err := gorm.Open(sqlite.Open("config.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	global.APP_DB = db
	if err = db.AutoMigrate(
		&model.RemoteServer{},
		&model.UnrealRunner{},
	); err != nil {
		fmt.Println(err)
	}
}
