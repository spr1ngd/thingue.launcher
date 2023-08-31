package main

import (
	"embed"
	"thingue-launcher/agent"
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
	provider.InitConfig()
	provider.SetWebStatic("server/frontend/dist", staticFiles)
	agent.Startup()
	agent.Shutdown()
}
