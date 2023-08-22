package initialize

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"thingue-launcher/agent/global"
	"thingue-launcher/common/model"
)

func InitGorm() {
	db, err := gorm.Open(sqlite.Open("./thingue-launcher/config.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	global.APP_DB = db
	if err = db.AutoMigrate(
		&model.RemoteServer{},
		&model.ClientInstance{},
	); err != nil {
		fmt.Println(err)
	}
}
