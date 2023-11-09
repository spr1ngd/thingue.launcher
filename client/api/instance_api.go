package api

import (
	"context"
	"errors"
	"thingue-launcher/client/service"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/domain"
	"thingue-launcher/common/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type instanceApi struct {
	ctx context.Context
}

var InstanceApi = new(instanceApi)

func (u *instanceApi) Init(ctx context.Context) {
	u.ctx = ctx
	service.RunnerManager.Init()
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

func (u *instanceApi) GetInstanceById(id uint) *model.ClientInstance {
	return service.InstanceManager.GetById(id)
}

func (u *instanceApi) ListInstance() []*domain.Instance {
	return service.RunnerManager.List()
}

func (u *instanceApi) CreateInstance(instance *model.ClientInstance) error {
	service.InstanceManager.Create(instance)
	err := service.RunnerManager.NewRunner(instance)
	if err == nil {
		service.ServerConnManager.Reconnect()
	} else {
		service.InstanceManager.Delete(instance.CID)
	}
	return err
}

func (u *instanceApi) SaveInstance(instance *model.ClientInstance) error {
	err := service.InstanceManager.SaveConfig(instance)
	if err == nil {
		service.ServerConnManager.Reconnect()
	}
	return err
}

func (u *instanceApi) DeleteInstance(cid uint) error {
	err := service.RunnerManager.DeleteRunner(cid)
	if err == nil {
		service.ServerConnManager.Reconnect()
		service.InstanceManager.Delete(cid)
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

func (u *instanceApi) OpenInstanceLog(id uint) error {
	runner := service.RunnerManager.GetRunnerById(id)
	return runner.OpenLog()
}

func (u *instanceApi) StartUpload(id uint) (string, error) {
	return service.SyncManager.StartUpload(id)
}

func (u *instanceApi) StartDownload(id uint) (string, error) {
	return service.SyncManager.StartUpdate(id)
}
