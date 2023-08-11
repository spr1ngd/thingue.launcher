package server

import (
	"embed"
	"thingue-launcher/server/core"
)

//go:embed all:frontend/dist
var staticFiles embed.FS

func Startup() {
	core.ServerApp.Start(staticFiles)
}

func Shutdown() {
	core.ServerApp.Stop()
}
