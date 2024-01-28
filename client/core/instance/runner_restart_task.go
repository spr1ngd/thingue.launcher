package instance

import (
	"github.com/robfig/cron/v3"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/provider"
)

type RunnerRestartTask struct {
	restartCron        *cron.Cron
	restartTaskEntryID cron.EntryID
}

func (t *RunnerRestartTask) Init() {
	t.restartCron = cron.New()
	if provider.AppConfig.SystemSettings.EnableRestartTask {
		err := t.Start()
		if err != nil {
			// 如果开启失败将设置改为false
			provider.AppConfig.SystemSettings.EnableRestartTask = false
			provider.WriteConfigToFile()
		}
	}
}

func (t *RunnerRestartTask) Start() error {
	var err error
	appConfig := provider.AppConfig
	t.restartTaskEntryID, err = t.restartCron.AddFunc(appConfig.SystemSettings.RestartTaskCron, func() {
		logger.Zap.Debug("重启定时任务执行开始")
		RunnerManager.RestartAllRunner()
		logger.Zap.Debug("重启定时任务执行结束")
	})
	t.restartCron.Start()
	return err
}

func (t *RunnerRestartTask) Stop() {
	t.restartCron.Remove(t.restartTaskEntryID)
	t.restartCron.Stop()
}
