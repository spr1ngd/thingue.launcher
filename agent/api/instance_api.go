package api

import (
	"context"
	"errors"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"thingue-launcher/agent/constants"
	"thingue-launcher/agent/core"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/model"
	"thingue-launcher/agent/service"
)

type instanceApi struct {
	ctx context.Context
}

var InstanceApi = new(instanceApi)

func (u *instanceApi) Init(ctx context.Context) {
	u.ctx = ctx
	// 从持久化数据中实例化Runners
	instances := u.ListInstance()
	for index := range instances {
		service.UeRunnerManager.NewRunner(&instances[index])
	}
	// 启动实例异常退出监听
	global.RunnerUnexpectedExitChanel = make(chan uint)
	go func() {
		for {
			id := <-global.RunnerUnexpectedExitChanel
			runtime.EventsEmit(ctx, constants.RUNNER_UNEXPECTED_EXIT, id)
		}
	}()
}

func (u *instanceApi) GetInstanceById(id uint) *model.Instance {
	instance := model.Instance{}
	global.APP_DB.First(&instance, id)
	return &instance
}

func (u *instanceApi) ListInstance() []model.Instance {
	var instances []model.Instance
	global.APP_DB.Find(&instances)
	return instances
}

func (u *instanceApi) ListRunner() []*core.Runner {
	var runners = make([]*core.Runner, 0)
	instances := u.ListInstance()
	for _, instance := range instances {
		runner := service.UeRunnerManager.GetRunnerById(instance.ID)
		runners = append(runners, runner)
	}
	return runners
}

func (u *instanceApi) CreateInstance(instance *model.Instance) uint {
	global.APP_DB.Create(&instance)
	service.UeRunnerManager.NewRunner(instance)
	return instance.ID
}

func (u *instanceApi) SaveInstance(instance *model.Instance) error {
	runner := service.UeRunnerManager.GetRunnerById(instance.ID)
	if runner.StateCode == 1 {
		return errors.New("实例运行中无法修改配置")
	}
	global.APP_DB.Save(instance)
	runner.Instance = instance
	return nil
}

func (u *instanceApi) DeleteInstance(id uint) error {
	err := service.UeRunnerManager.DeleteRunner(id)
	if err == nil {
		global.APP_DB.Delete(&model.Instance{}, id)
	}
	return err
}

func (u *instanceApi) StartInstance(id uint) error {
	runner := service.UeRunnerManager.GetRunnerById(id)
	if runner != nil {
		return runner.Start()
	} else {
		return errors.New("实例不存在")
	}
}

func (u *instanceApi) StopInstance(id uint) error {
	runner := service.UeRunnerManager.GetRunnerById(id)
	if runner != nil {
		return runner.Stop()
	} else {
		return errors.New("实例不存在")
	}
}
