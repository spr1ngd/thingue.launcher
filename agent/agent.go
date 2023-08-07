package agent

import (
	"embed"
	"fmt"
	"thingue-launcher/agent/initialize"
)

//go:embed all:frontend/dist
var assets embed.FS

func Startup() {
	//初始化Gorm
	initialize.InitGorm()
	//初始化App
	initialize.InitApp(assets)
}

func Shutdown() {
	fmt.Println("agent关闭")
}
