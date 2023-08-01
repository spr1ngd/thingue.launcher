package app

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"thingue-launcher/common/config"
)

// App struct
type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) SelectExePath(name string) string {
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择文件",
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

func (a *App) GetAppConfig() config.AppConfig {
	return *config.GetAppConfig()
}
