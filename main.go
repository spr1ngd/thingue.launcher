package main

import (
	"os"
	"thingue-launcher/agent"
	"thingue-launcher/agent/global"
	"thingue-launcher/common/app"
	"thingue-launcher/server"
)

var (
	Version   = "0.0.1"
	GitCommit string
	BuildDate string
)

func main() {
	global.SetAppVersion(Version, GitCommit, BuildDate)
	app.InitConfig()
	if len(os.Args) > 1 && os.Args[1] == "server" {
		server.Startup()
	} else {
		agent.Startup()
		agent.Shutdown()
	}
}
