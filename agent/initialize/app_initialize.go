package initialize

import (
	"context"
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"thingue-launcher/agent/api"
	"time"
)

func InitApp(assets embed.FS) {
	myLog := NewFileLogger(fmt.Sprintf("thingue-launcher/info-%v.log", time.Now().Format("2006-01-02")))
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
			fmt.Println("88888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888")
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
