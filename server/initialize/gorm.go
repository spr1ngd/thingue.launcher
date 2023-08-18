package initialize

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"thingue-launcher/common/model"
	"thingue-launcher/server/global"
)

func InitGorm() {
	db, err := gorm.Open(sqlite.Open("server.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	global.SERVER_DB = db
	if err = db.AutoMigrate(
		&model.Node{},
		&model.Instance{},
	); err != nil {
		fmt.Println(err)
	}
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Node{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Instance{})
}
