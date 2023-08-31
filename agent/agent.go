package agent

import (
	"embed"
	"thingue-launcher/agent/initialize"
	"thingue-launcher/agent/service"
)

//go:embed all:frontend/dist
var assets embed.FS

func Startup() {
	//初始化Gorm
	initialize.InitGorm()
	//初始化并运行App
	initialize.InitApp(assets)
	//App关闭时清理进程
	service.RunnerManager.CloseAllRunner()
}

func Shutdown() {
}
