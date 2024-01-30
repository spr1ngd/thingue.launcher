package api

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"thingue-launcher/client/core"
	"thingue-launcher/client/global"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/domain"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/common/logger"
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

func (u *instanceApi) GetInstanceById(id uint) *domain.Instance {
	return core.ConfigManager.GetById(id)
}

func (u *instanceApi) ListInstance() []*domain.Instance {
	return core.RunnerManager.List()
}

func (u *instanceApi) GetDefaultConfig() *domain.Instance {
	return core.ConfigManager.GetDefault()
}

func (u *instanceApi) CreateInstance(instance *domain.Instance) error {
	id := core.ConfigManager.Create(instance)
	instance.ID = id
	err := core.RunnerManager.NewRunner(instance)
	if err != nil {
		core.ConfigManager.Delete(id)
		return err
	}
	if global.GrpcClient != nil {
		_, err := global.GrpcClient.AddInstance(context.Background(),
			&pb.AddInstanceRequest{
				ClientId:     global.ClientId,
				InstanceInfo: instance.ToInstanceInfoTypes(),
			})
		if err != nil {
			logger.Zap.Error(err)
		}
	}
	return nil
}

func (u *instanceApi) UpdateConfig(instance *domain.Instance) error {
	//fmt.Printf("更新%+v", instance)
	runner, err := core.RunnerManager.GetRunnerById(instance.ID)
	if err != nil {
		return err
	}
	runner.Config = instance.Config
	runner.PlayerConfig = instance.PlayerConfig
	err = core.ConfigManager.Update(instance)
	if err != nil {
		return err
	}
	if global.GrpcClient != nil {
		instanceInfo := instance.ToInstanceInfoTypes()
		_, err := global.GrpcClient.UpdateConfig(context.Background(), &pb.UpdateConfigRequest{
			ClientId:       global.ClientId,
			InstanceId:     instance.ID,
			InstanceConfig: instanceInfo.Config,
			PlayerConfig:   instanceInfo.PlayerConfig,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *instanceApi) DeleteInstance(id uint32) error {
	if global.GrpcClient != nil {
		_, err := global.GrpcClient.DeleteInstance(context.Background(), &pb.DeleteInstanceRequest{
			ClientId:   global.ClientId,
			InstanceId: id,
		})
		if err != nil {
			return err
		}
	}
	err := core.RunnerManager.DeleteRunner(id)
	if err != nil {
		return err
	}
	core.ConfigManager.Delete(id)
	return nil
}

func (u *instanceApi) StartInstance(id uint32) error {
	runner, err := core.RunnerManager.GetRunnerById(id)
	if err != nil {
		return err
	}
	return runner.Start()
}

func (u *instanceApi) StopInstance(id uint32) error {
	runner, err := core.RunnerManager.GetRunnerById(id)
	if err != nil {
		return err
	}
	_, err = global.GrpcClient.ClearPakState(context.Background(), &pb.ClearPakStateRequest{
		ClientId:   global.ClientId,
		InstanceId: runner.ID,
	})
	if err != nil {
		logger.Zap.Error(err)
	}
	return runner.Stop()
}

func (u *instanceApi) OpenInstanceLog(id uint32) error {
	runner, err := core.RunnerManager.GetRunnerById(id)
	if err != nil {
		return err
	}
	return runner.OpenLog()
}

func (u *instanceApi) StartUpload(id uint32) (string, error) {
	return core.SyncManager.StartUpload(id)
}

func (u *instanceApi) StartDownload(id uint32) (string, error) {
	return core.SyncManager.StartUpdate(id)
}
