package server

import (
	"embed"
	"thingue-launcher/common/provider"
)

var (
	//go:embed all:frontend/dist
	staticFiles embed.FS
)

func Init() {
	provider.SetWebStatic("frontend/dist", staticFiles)
	if provider.AppConfig.LocalServer.StaticDir != "" {
		provider.AppConfig.LocalServer.UseExternalStatic = true
	}
}
