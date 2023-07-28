package agent

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"thingue-launcher/agent/config"
	"thingue-launcher/agent/unreal"
	"thingue-launcher/server"
)

// App struct
type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	config.InitAppConfig()
	a.ctx = ctx
}

func (a *App) SelectExePath(name string) string {
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "ThingUE (*.exe)",
				Pattern:     "*.exe",
			},
		},
	})
	if err != nil {
		return ""
	}
	return selection
}

func (a *App) ServerStart() {
	appConfig := config.GetAppConfig()
	runtime.EventsEmit(a.ctx, "local_server_status_update", true)
	server.Start(appConfig.LocalServer.BindAddr, appConfig.LocalServer.BasePath)
	runtime.EventsEmit(a.ctx, "local_server_status_update", false)
}

func (a *App) ServerShutdown() {
	server.Shutdown()
}

func (a *App) GetServerStatus() bool {
	return server.GetServerStatus()
}

func (a *App) UnrealStart(exePath string, params []string) int {
	return unreal.UnrealStart(exePath, params...)
}

func (a *App) UnrealStop(pid int) {
	unreal.UnrealStop(pid)
}

func (a *App) GetAppConfig() config.AppConfig {
	return *config.GetAppConfig()
}

func (a *App) UpdateLocalServerConfig(localServerConfig config.LocalServer) {
	appConfig := config.GetAppConfig()
	appConfig.LocalServer = localServerConfig
	config.WriteConfig()
}
