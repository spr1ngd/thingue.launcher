package manager

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"thingue-launcher/common/app"
)

type RunnerRestartTaskManager struct {
	restartCron        *cron.Cron
	restartTaskEntryID cron.EntryID
}

func (t *RunnerRestartTaskManager) Init() {
	t.restartCron = cron.New()
	if app.GetAppConfig().EnableRestartTask {
		err := t.Start()
		if err != nil {
			// 如果开启失败将设置改为false
			app.GetAppConfig().EnableRestartTask = false
			app.WriteConfig()
		}
	}
}

func (t *RunnerRestartTaskManager) Start() error {
	var err error
	appConfig := app.GetAppConfig()
	t.restartTaskEntryID, err = t.restartCron.AddFunc(appConfig.SystemSettings.RestartTaskCron, func() {
		fmt.Println("重启定时任务执行开始")
		RunnerManager.RestartAllRunner()
		fmt.Println("重启定时任务执行结束")
	})
	t.restartCron.Start()
	return err
}

func (t *RunnerRestartTaskManager) Stop() {
	t.restartCron.Remove(t.restartTaskEntryID)
	t.restartCron.Stop()
}
