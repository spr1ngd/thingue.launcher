package app

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os/exec"
	"path/filepath"
	"thingue-launcher/agent/global"
	"thingue-launcher/common/app"
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
	InitRestartTask(ctx)
}

func (a *App) OpenFileDialog(title string, displayName string, pattern string) (string, error) {
	return runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: title,
		Filters: []runtime.FileFilter{
			{
				DisplayName: displayName,
				Pattern:     pattern,
			},
		},
	})
}

func (a *App) OpenExplorer(path string) error {
	cmd := exec.Command("explorer", filepath.Dir(path))
	err := cmd.Run()
	return err
}

func (a *App) GetAppConfig() *app.Config {
	return app.GetAppConfig()
}

func (a *App) ControlRestartTask(enable bool) error {
	var err error
	appConfig := app.GetAppConfig()
	if enable {
		err = EnableRestartTask()
	} else {
		DisableRestartTask()
	}
	if err == nil {
		appConfig.EnableRestartTask = enable
		app.WriteConfig()
	}
	return err
}

func (a *App) UpdateSystemSettings(systemSettings app.SystemSettings) {
	fmt.Println(systemSettings)
	appConfig := app.GetAppConfig()
	appConfig.SystemSettings = systemSettings
	app.WriteConfig()
}

func (a *App) GetVersionInfo() *VersionInfo {
	return &VersionInfo{
		Version:   global.APP_VERSION,
		GitCommit: global.APP_GITCOMMIT,
		BuildDate: global.APP_BUILDDATE,
	}
}
