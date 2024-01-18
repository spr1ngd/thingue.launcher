package main

import (
	"embed"
	"os"
	"thingue-launcher/client"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/provider"
)

var (
	GitCommit  string
	BuildDate  string
	AppVersion string
	//go:embed all:server/frontend/dist
	staticFiles embed.FS
)

func main() {
	constants.SetEmbed("server/frontend/dist", staticFiles)
	constants.SetVersionInfo(AppVersion, GitCommit, BuildDate)
	_, err := os.Stat(constants.SaveDir)
	if os.IsNotExist(err) {
		os.MkdirAll(constants.SaveDir, 0755)
	}
	provider.InitConfigFromFile()
	logger.InitZapLogger(provider.AppConfig.SystemSettings.LogLevel, "app.log")
	client.Startup()
	client.Shutdown()
}
