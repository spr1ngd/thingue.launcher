package api

import (
	"context"
	"thingue-launcher/client/core"
	"thingue-launcher/client/global"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/domain"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type instanceApi struct {
	ctx context.Context
}

var InstanceApi = new(instanceApi)

func (u *instanceApi) Init(ctx context.Context) {
	u.ctx = ctx
	core.RunnerManager.Init()
	// 启动实例异常退出监听
	go func() {
		for {
			id := <-core.RunnerManager.RunnerUnexpectedExitChanel
			runtime.EventsEmit(ctx, constants.RUNNER_UNEXPECTED_EXIT, id)
		}
	}()
	// 启动实例状态变化监听
	go func() {
		for {
			id := <-core.RunnerManager.RunnerStatusUpdateChanel
			runtime.EventsEmit(ctx, constants.RUNNER_STATUS_UPDATE, id)
		}
	}()
}

func (u *instanceApi) GetInstanceById(id uint) *model.ClientInstance {
	return core.InstanceManager.GetById(id)
}

func (u *instanceApi) ListInstance() []*domain.Instance {
	return core.RunnerManager.List()
}

func (u *instanceApi) CreateInstance(instance *model.ClientInstance) error {
	core.InstanceManager.Create(instance)
	err := core.RunnerManager.NewRunner(instance)
	if err == nil {
		// todo
		//service.ServerConnManager.Reconnect()
	} else {
		core.InstanceManager.Delete(instance.CID)
	}
	return err
}

func (u *instanceApi) SaveInstance(instance *model.ClientInstance) error {
	err := core.InstanceManager.SaveConfig(instance)
	if err == nil {
		// todo
		//service.ServerConnManager.Reconnect()
	}
	return err
}

func (u *instanceApi) DeleteInstance(cid uint) error {
	err := core.RunnerManager.DeleteRunner(cid)
	if err == nil {
		// todo
		//service.ServerConnManager.Reconnect()
		core.InstanceManager.Delete(cid)
	}
	return err
}

func (u *instanceApi) StartInstance(id uint) error {
	runner, err := core.RunnerManager.GetRunnerById(id)
	if err != nil {
		return err
	}
	return runner.Start()
}

func (u *instanceApi) StopInstance(id uint) error {
	runner, err := core.RunnerManager.GetRunnerById(id)
	if err != nil {
		return err
	}
	_, err = global.GrpcClient.ClearPakState(context.Background(), &pb.ClearPakStateRequest{
		ClientId:   uint32(global.ClientId),
		InstanceId: uint32(runner.CID),
	})
	if err != nil {
		logger.Zap.Error(err)
	}
	return runner.Stop()
}

func (u *instanceApi) OpenInstanceLog(id uint) error {
	runner, err := core.RunnerManager.GetRunnerById(id)
	if err != nil {
		return err
	}
	return runner.OpenLog()
}

func (u *instanceApi) StartUpload(id uint) (string, error) {
	return core.SyncManager.StartUpload(id)
}

func (u *instanceApi) StartDownload(id uint) (string, error) {
	return core.SyncManager.StartUpdate(id)
}
