package initialize

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"thingue-launcher/agent/api"
)

func InitApp(assets embed.FS) {
	// 初始化wails app
	err := wails.Run(&options.App{
		Title:  "ThingUE启动器",
		Width:  820,
		Height: 500,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup: func(ctx context.Context) {
			api.InstanceApi.Init(ctx)
			api.ServerApi.Init(ctx)
			api.SystemApi.Init(ctx)
		},
		Bind: []interface{}{
			api.InstanceApi,
			api.ServerApi,
			api.SystemApi,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
