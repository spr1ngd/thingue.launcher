package initialize

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"thingue-launcher/agent/app"
	"thingue-launcher/agent/server"
	"thingue-launcher/agent/unreal"
)

func InitApp(assets embed.FS) {
	// 初始化wails app
	newApp := app.NewApp()
	newServer := server.NewServer()
	newUnreal := unreal.NewUnreal()
	err := wails.Run(&options.App{
		Title:  "ThingUE启动器",
		Width:  800,
		Height: 480,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup: func(ctx context.Context) {
			newApp.SetContext(ctx)
			newUnreal.SetContext(ctx)
			newServer.SetContext(ctx)
		},
		Bind: []interface{}{
			newApp,
			newUnreal,
			newServer,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
