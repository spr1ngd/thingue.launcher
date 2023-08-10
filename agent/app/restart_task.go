package app

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"thingue-launcher/agent/unreal"
	"thingue-launcher/common/app"
)

var restartCron *cron.Cron
var restartTaskEntryID cron.EntryID
var ctx context.Context

func InitRestartTask(context context.Context) {
	ctx = context
	restartCron = cron.New()
	appConfig := app.GetAppConfig()
	if appConfig.EnableRestartTask {
		err := EnableRestartTask()
		if err != nil {
			// 如果开启失败将设置改为false
			appConfig.EnableRestartTask = false
			app.WriteConfig()
		}
	}
}

func EnableRestartTask() error {
	var err error
	appConfig := app.GetAppConfig()
	restartTaskEntryID, err = restartCron.AddFunc(appConfig.SystemSettings.RestartTaskCron, func() {
		fmt.Println("重启定时任务执行开始")
		unreal.RestartAllRunner(ctx)
		fmt.Println("重启定时任务执行结束")
	})
	restartCron.Start()
	return err
}

func DisableRestartTask() {
	restartCron.Remove(restartTaskEntryID)
	restartCron.Stop()
}
