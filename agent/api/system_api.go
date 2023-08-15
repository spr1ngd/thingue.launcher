package api

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os/exec"
	"path/filepath"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/service"
	"thingue-launcher/common/config"
	"thingue-launcher/common/model"
)

type systemApi struct {
	ctx context.Context
}

var SystemApi = new(systemApi)

func (a *systemApi) Init(ctx context.Context) {
	a.ctx = ctx
	service.RunnerRestartTaskManager.Init()
}

func (a *systemApi) OpenFileDialog(title string, displayName string, pattern string) (string, error) {
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

func (a *systemApi) OpenExplorer(path string) error {
	cmd := exec.Command("explorer", filepath.Dir(path))
	err := cmd.Run()
	return err
}

func (a *systemApi) GetAppConfig() *config.Config {
	return config.AppConfig
}

func (a *systemApi) ControlRestartTask(enable bool) error {
	var err error
	appConfig := config.AppConfig
	if enable {
		err = service.RunnerRestartTaskManager.Start()
	} else {
		service.RunnerRestartTaskManager.Stop()
	}
	if err == nil {
		appConfig.EnableRestartTask = enable
		config.WriteConfig()
	}
	return err
}

func (a *systemApi) UpdateSystemSettings(systemSettings config.SystemSettings) {
	appConfig := config.AppConfig
	appConfig.SystemSettings = systemSettings
	config.WriteConfig()
}

func (a *systemApi) GetVersionInfo() *model.VersionInfo {
	return &model.VersionInfo{
		Version:   global.APP_VERSION,
		GitCommit: global.APP_GITCOMMIT,
		BuildDate: global.APP_BUILDDATE,
	}
}
