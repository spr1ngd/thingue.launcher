package initialize

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"thingue-launcher/common/model"
	"thingue-launcher/server/global"
)

func InitGorm() {
	//dsn := "./thingue-launcher/server.db"
	dsn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	global.SERVER_DB = db
	if err = db.AutoMigrate(
		&model.Node{},
		&model.ServerInstance{},
	); err != nil {
		fmt.Println(err)
	}
}
