package initialize

import (
	"fmt"
	"thingue-launcher/client/global"
	"thingue-launcher/common/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
