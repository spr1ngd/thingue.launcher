package api

import (
	"context"
	"errors"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"thingue-launcher/agent/constants"
	"thingue-launcher/agent/service"
	"thingue-launcher/common/model"
)

type instanceApi struct {
	ctx context.Context
}

var InstanceApi = new(instanceApi)

func (u *instanceApi) Init(ctx context.Context) {
	u.ctx = ctx
	// 从持久化数据中实例化Runners
	instances := service.InstanceManager.List()
	for index := range instances {
		service.RunnerManager.NewRunner(&instances[index])
	}
	// 启动实例异常退出监听
	go func() {
		for {
			id := <-service.RunnerManager.RunnerUnexpectedExitChanel
			runtime.EventsEmit(ctx, constants.RUNNER_UNEXPECTED_EXIT, id)
		}
	}()
	// 启动实例状态变化监听
	go func() {
		for {
			id := <-service.RunnerManager.RunnerStatusUpdateChanel
			runtime.EventsEmit(ctx, constants.RUNNER_STATUS_UPDATE, id)
		}
	}()
}

func (u *instanceApi) GetInstanceById(id uint) *model.Instance {
	return service.InstanceManager.GetById(id)
}

func (u *instanceApi) ListInstance() []*model.Instance {
	return service.RunnerManager.List()
}

func (u *instanceApi) CreateInstance(instance *model.Instance) error {
	service.InstanceManager.Create(instance)
	err := service.RunnerManager.NewRunner(instance)
	if err != nil {
		service.InstanceManager.Delete(instance.ID)
	}
	return err
}

func (u *instanceApi) SaveInstance(instance *model.Instance) error {
	return service.InstanceManager.Save(instance)
}

func (u *instanceApi) DeleteInstance(id uint) error {
	err := service.RunnerManager.DeleteRunner(id)
	if err == nil {
		service.InstanceManager.Delete(id)
	}
	return err
}

func (u *instanceApi) StartInstance(id uint) error {
	runner := service.RunnerManager.GetRunnerById(id)
	if runner != nil {
		return runner.Start()
	} else {
		return errors.New("实例不存在")
	}
}

func (u *instanceApi) StopInstance(id uint) error {
	runner := service.RunnerManager.GetRunnerById(id)
	if runner != nil {
		return runner.Stop()
	} else {
		return errors.New("实例不存在")
	}
}
