package main

import (
	"embed"
	"os"
	"thingue-launcher/client"
	"thingue-launcher/common/constants"
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
	provider.SetVersionBuildInfo(AppVersion, GitCommit, BuildDate)
	_, err := os.Stat(constants.SAVE_DIR)
	if os.IsNotExist(err) {
		os.MkdirAll(constants.SAVE_DIR, 0755)
	}
	provider.InitConfigFromFile()
	provider.SetWebStatic("server/frontend/dist", staticFiles)
	client.Startup()
	client.Shutdown()
}
