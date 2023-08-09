package app

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os/exec"
	"path/filepath"
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

func (a *App) GetAppConfig() config.AppConfig {
	return *config.GetAppConfig()
}

func (a *App) ControlRestartTask(enable bool) error {
	var err error
	appConfig := config.GetAppConfig()
	if enable {
		err = EnableRestartTask()
	} else {
		DisableRestartTask()
	}
	if err == nil {
		appConfig.EnableRestartTask = enable
		config.WriteConfig()
	}
	return err
}

func (a *App) UpdateSystemSettings(systemSettings config.SystemSettings) {
	fmt.Println(systemSettings)
	appConfig := config.GetAppConfig()
	appConfig.SystemSettings = systemSettings
	config.WriteConfig()
}
