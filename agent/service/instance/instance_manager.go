package instance

import (
	"errors"
	"thingue-launcher/agent/global"
	"thingue-launcher/common/model"
)

type instanceManager struct{}

var InstanceManager = new(instanceManager)

func (m *instanceManager) List() []model.Instance {
	var instances []model.Instance
	global.APP_DB.Find(&instances)
	return instances
}

func (m *instanceManager) Create(instance *model.Instance) uint {
	global.APP_DB.Create(&instance)
	return instance.ID
}

func (m *instanceManager) GetById(id uint) *model.Instance {
	instance := model.Instance{}
	global.APP_DB.First(&instance, id)
	return &instance
}

func (m *instanceManager) Save(instance *model.Instance) error {
	runner := RunnerManager.GetRunnerById(instance.ID)
	if runner.StateCode == 1 {
		return errors.New("实例运行中无法修改配置")
	}
	global.APP_DB.Save(instance)
	runner.Instance = instance
	return nil
}

func (m *instanceManager) Delete(id uint) {
	global.APP_DB.Delete(&model.Instance{}, id)
}
