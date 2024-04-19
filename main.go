package main

import (
	"thingue-launcher/client"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/provider"
	"thingue-launcher/server"
)

var (
	GitCommit  string
	BuildDate  string
	AppVersion string
)

func main() {
	provider.SetBuildInfo(AppVersion, GitCommit, BuildDate)

	client.Init()
	logger.InitZapLogger(provider.AppConfig.SystemSettings.LogLevel, "app.log")

	server.Init()

	client.RunApp()
}
