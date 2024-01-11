package initialize

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"thingue-launcher/client/api"
	"thingue-launcher/common/provider"
)

func InitApp(assets embed.FS) {
	zapLogger := ZapLogger{}
	// 初始化wails app
	err := wails.Run(&options.App{
		Title:  "ThingUE启动器 v" + provider.VersionInfo.Version,
		Width:  820,
		Height: 510,
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
		Logger: &zapLogger,
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
