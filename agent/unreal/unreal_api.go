package unreal

import (
	"context"
	"errors"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/model"
)

type Unreal struct {
	ctx context.Context
}

func NewUnreal() *Unreal {
	return &Unreal{}
}

func (u *Unreal) SetContext(ctx context.Context) {
	u.ctx = ctx
	instances := u.ListInstance()

	for index := range instances {
		NewRunner(&instances[index])
	}
}

func (u *Unreal) GetInstanceById(id uint) *model.Instance {
	instance := model.Instance{}
	global.APP_DB.First(&instance, id)
	return &instance
}

func (u *Unreal) ListInstance() []model.Instance {
	var instances []model.Instance
	global.APP_DB.Find(&instances)
	return instances
}

func (u *Unreal) ListRunner() []*Runner {
	var runners = make([]*Runner, 0)
	instances := u.ListInstance()
	for _, instance := range instances {
		runners = append(runners, idRunnerMap[instance.ID])
	}
	return runners
}

func (u *Unreal) CreateInstance(instance *model.Instance) uint {
	global.APP_DB.Create(&instance)
	NewRunner(instance)
	return instance.ID
}

func (u *Unreal) SaveInstance(instance *model.Instance) error {
	runner := GetRunnerById(instance.ID)
	if runner.IsRunning {
		return errors.New("实例运行中无法修改配置")
	}
	global.APP_DB.Save(instance)
	runner.Instance = instance
	return nil
}

func (u *Unreal) DeleteInstance(id uint) error {
	runner := GetRunnerById(id)
	if runner != nil {
		if runner.IsRunning {
			return errors.New("实例正在运行")
		}
		err := runner.delete()
		global.APP_DB.Delete(&model.Instance{}, id)
		return err
	} else {
		return errors.New("实例不存在")
	}
}

func (u *Unreal) StartInstance(id uint) error {
	runner := GetRunnerById(id)
	if runner != nil {
		return runner.start(u.ctx)
	} else {
		return errors.New("实例不存在")
	}
}

func (u *Unreal) StopInstance(id uint) error {
	runner := GetRunnerById(id)
	if runner != nil {
		return runner.stop()
	} else {
		return errors.New("实例不存在")
	}
}
