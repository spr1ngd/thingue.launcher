package api

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os/exec"
	"path/filepath"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/service"
	"thingue-launcher/agent/system"
	"thingue-launcher/common/app"
)

type System struct {
	ctx context.Context
}

var AppApi = new(System)

func (a *System) Init(ctx context.Context) {
	a.ctx = ctx
	service.RunnerRestartTaskManager.Init()
}

func (a *System) OpenFileDialog(title string, displayName string, pattern string) (string, error) {
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

func (a *System) OpenExplorer(path string) error {
	cmd := exec.Command("explorer", filepath.Dir(path))
	err := cmd.Run()
	return err
}

func (a *System) GetAppConfig() *app.Config {
	return app.GetAppConfig()
}

func (a *System) ControlRestartTask(enable bool) error {
	var err error
	appConfig := app.GetAppConfig()
	if enable {
		err = service.RunnerRestartTaskManager.Start()
	} else {
		service.RunnerRestartTaskManager.Stop()
	}
	if err == nil {
		appConfig.EnableRestartTask = enable
		app.WriteConfig()
	}
	return err
}

func (a *System) UpdateSystemSettings(systemSettings app.SystemSettings) {
	appConfig := app.GetAppConfig()
	appConfig.SystemSettings = systemSettings
	app.WriteConfig()
}

func (a *System) GetVersionInfo() *system.VersionInfo {
	return &system.VersionInfo{
		Version:   global.APP_VERSION,
		GitCommit: global.APP_GITCOMMIT,
		BuildDate: global.APP_BUILDDATE,
	}
}
