package instance

import (
	"errors"
	"thingue-launcher/agent/global"
	"thingue-launcher/common/model"
)

type instanceManager struct{}

var InstanceManager = new(instanceManager)

func (m *instanceManager) List() []model.ClientInstance {
	var instances []model.ClientInstance
	global.APP_DB.Find(&instances)
	return instances
}

func (m *instanceManager) Create(instance *model.ClientInstance) uint {
	global.APP_DB.Create(&instance)
	return instance.ID
}

func (m *instanceManager) GetById(id uint) *model.ClientInstance {
	instance := model.ClientInstance{}
	global.APP_DB.First(&instance, id)
	return &instance
}

func (m *instanceManager) Save(instance *model.ClientInstance) error {
	runner := RunnerManager.GetRunnerById(instance.ID)
	if runner.StateCode == 1 {
		return errors.New("实例运行中无法修改配置")
	}
	global.APP_DB.Save(instance)
	runner.ClientInstance = instance
	return nil
}

func (m *instanceManager) Delete(id uint) {
	global.APP_DB.Delete(&model.ClientInstance{}, id)
}
