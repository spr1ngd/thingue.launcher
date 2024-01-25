package client

import (
	"embed"
	"thingue-launcher/client/core"
	"thingue-launcher/client/initialize"
)

//go:embed all:frontend/dist
var assets embed.FS

func Startup() {
	//初始化Gorm
	initialize.InitGorm()
	//初始化并运行App
	initialize.InitApp(assets)
	//App关闭时清理进程
	core.RunnerManager.CloseAllRunner()
}

func Shutdown() {
}
