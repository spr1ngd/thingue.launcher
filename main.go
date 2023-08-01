package main

import (
	"os"
	"thingue-launcher/agent"
	"thingue-launcher/common/config"
	"thingue-launcher/server"
)

func main() {
	config.InitConfig()
	if len(os.Args) > 1 && os.Args[1] == "server" {
		server.Start()
	} else {
		agent.Startup()
	}
}
