package api

import (
	"context"
	"errors"
	"thingue-launcher/agent/core"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/model"
	"thingue-launcher/agent/service"
)

type Instance struct {
	ctx context.Context
}

func (u *Instance) SetContext(ctx context.Context) {
	u.ctx = ctx
	instances := u.ListInstance()

	for index := range instances {
		service.UeRunnerManager.NewRunner(&instances[index])
	}
}

func (u *Instance) GetInstanceById(id uint) *model.Instance {
	instance := model.Instance{}
	global.APP_DB.First(&instance, id)
	return &instance
}

func (u *Instance) ListInstance() []model.Instance {
	var instances []model.Instance
	global.APP_DB.Find(&instances)
	return instances
}

func (u *Instance) ListRunner() []*core.Runner {
	var runners = make([]*core.Runner, 0)
	instances := u.ListInstance()
	for _, instance := range instances {
		runner := service.UeRunnerManager.GetRunnerById(instance.ID)
		runners = append(runners, runner)
	}
	return runners
}

func (u *Instance) CreateInstance(instance *model.Instance) uint {
	global.APP_DB.Create(&instance)
	service.UeRunnerManager.NewRunner(instance)
	return instance.ID
}

func (u *Instance) SaveInstance(instance *model.Instance) error {
	runner := service.UeRunnerManager.GetRunnerById(instance.ID)
	if runner.IsRunning {
		return errors.New("实例运行中无法修改配置")
	}
	global.APP_DB.Save(instance)
	runner.Instance = instance
	return nil
}

func (u *Instance) DeleteInstance(id uint) error {
	err := service.UeRunnerManager.DeleteRunner(id)
	if err == nil {
		global.APP_DB.Delete(&model.Instance{}, id)
	}
	return err
}

func (u *Instance) StartInstance(id uint) error {
	runner := service.UeRunnerManager.GetRunnerById(id)
	if runner != nil {
		return runner.Start()
	} else {
		return errors.New("实例不存在")
	}
}

func (u *Instance) StopInstance(id uint) error {
	runner := service.UeRunnerManager.GetRunnerById(id)
	if runner != nil {
		return runner.Stop()
	} else {
		return errors.New("实例不存在")
	}
}
