package initialize

import (
	"context"
	"embed"
	"fmt"
	"thingue-launcher/client/api"
	"thingue-launcher/common/provider"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func InitApp(assets embed.FS) {
	myLog := NewFileLogger(fmt.Sprintf("thingue-launcher/info-%v.log", time.Now().Format("2006-01-02")))
	// 初始化wails app
	err := wails.Run(&options.App{
		Title:  "ThingUE启动器 v" + provider.VersionInfo.Version,
		Width:  820,
		Height: 501,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup: func(ctx context.Context) {
			api.InstanceApi.Init(ctx)
			api.ServerApi.Init(ctx)
			api.SystemApi.Init(ctx)
		},
		Logger: myLog,
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
