package api

import (
	"context"
	"os/exec"
	"path/filepath"
	"thingue-launcher/client/service"
	"thingue-launcher/common/domain"
	"thingue-launcher/common/provider"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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

func (a *systemApi) GetAppConfig() *provider.Config {
	return provider.AppConfig
}

func (a *systemApi) ControlRestartTask(enable bool) error {
	var err error
	if enable {
		err = service.RunnerRestartTaskManager.Start()
	} else {
		service.RunnerRestartTaskManager.Stop()
	}
	if err == nil {
		provider.AppConfig.SystemSettings.EnableRestartTask = enable
		provider.WriteConfigToFile()
	}
	return err
}

func (a *systemApi) UpdateSystemSettings(systemSettings provider.SystemSettings) {
	appConfig := provider.AppConfig
	appConfig.SystemSettings = systemSettings
	provider.WriteConfigToFile()
}

func (a *systemApi) GetVersionInfo() *domain.VersionInfo {
	return provider.VersionInfo
}
