package server

import (
	"embed"
	"thingue-launcher/server/core"
)

//go:embed all:frontend/dist
var staticFiles embed.FS

func Init() {
	core.ServerApp.Init(staticFiles)
}

func Startup() {
	core.ServerApp.Start()
}

func Shutdown() {
	core.ServerApp.Stop()
}
