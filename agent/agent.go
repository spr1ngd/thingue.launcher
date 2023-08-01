package agent

import (
	"embed"
	"thingue-launcher/agent/initialize"
)

//go:embed all:frontend/dist
var assets embed.FS

func Startup() {
	//初始化App
	initialize.InitApp(assets)
	//初始化Gorm
	initialize.InitGorm()
}

func Shutdown() {

}
