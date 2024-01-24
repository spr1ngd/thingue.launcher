package initialize

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"thingue-launcher/client/api"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/logger"
)

func InitApp(assets embed.FS) {
	// 初始化wails app
	err := wails.Run(&options.App{
		Title:  "ThingUE启动器 v" + constants.VersionInfo.Version,
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
		Windows: &windows.Options{
			Theme: windows.Light,
		},
		Bind: []any{
			api.InstanceApi,
			api.ServerApi,
			api.SystemApi,
		},
		Logger: &ZapLogger{},
	})
	if err != nil {
		logger.Zap.Fatalf("App初始化失败%s", err.Error())
	}
}
