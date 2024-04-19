package client

import (
	"embed"
	"os"
	"thingue-launcher/client/initialize"
	"thingue-launcher/client/service"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/provider"
)

//go:embed all:frontend/dist
var assets embed.FS

func Init() {
	_, err := os.Stat(constants.SAVE_DIR)
	if os.IsNotExist(err) {
		os.MkdirAll(constants.SAVE_DIR, 0755)
	}
	provider.InitConfigFromFile()
	initialize.InitGorm()
}

func RunApp() {
	//初始化并运行App
	initialize.InitApp(assets)
	//App关闭时清理进程
	service.RunnerManager.CloseAllRunner()
}
