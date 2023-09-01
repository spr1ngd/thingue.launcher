package main

import (
	"embed"
	"os"
	"thingue-launcher/agent"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/provider"
)

var (
	GitCommit string
	BuildDate string
	//go:embed all:server/frontend/dist
	staticFiles embed.FS
)

func main() {
	provider.SetVersionBuildInfo(GitCommit, BuildDate)
	_, err := os.Stat(constants.SAVE_DIR)
	if os.IsNotExist(err) {
		os.MkdirAll(constants.SAVE_DIR, 0755)
	}
	provider.InitConfigFromFile()
	provider.SetWebStatic("server/frontend/dist", staticFiles)
	agent.Startup()
	agent.Shutdown()
}
