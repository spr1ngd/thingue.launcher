package instance

import (
	"errors"
	"thingue-launcher/client/global"
	"thingue-launcher/common/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type instanceManager struct{}

var InstanceManager = new(instanceManager)

func (m *instanceManager) List() []model.ClientInstance {
	var instances []model.ClientInstance
	global.AppDB.Find(&instances)
	return instances
}

func (m *instanceManager) Create(instance *model.ClientInstance) uint {
	global.AppDB.Create(&instance)
	return instance.CID
}

func (m *instanceManager) GetById(id uint) *model.ClientInstance {
	var instance model.ClientInstance
	result := global.AppDB.First(&instance, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &instance
}

func (m *instanceManager) GetInternal() (*model.ClientInstance, error) {
	var instance model.ClientInstance
	result := global.AppDB.Where(&model.ClientInstance{IsInternal: true}).First(&instance)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &instance, result.Error
	}
	return &instance, nil
}

func (m *instanceManager) SaveConfig(instance *model.ClientInstance) error {
	runner, err := RunnerManager.GetRunnerById(instance.CID)
	if err != nil {
		return err
	}
	if runner.StateCode == 1 {
		return errors.New("实例运行中无法修改配置")
	}
	global.AppDB.Save(instance)
	return copier.Copy(runner.Instance, instance)
}

func (m *instanceManager) Delete(id uint) {
	global.AppDB.Delete(&model.ClientInstance{}, id)
}

func (m *instanceManager) GetByCloudRes(res string) []model.ClientInstance {
	var instances []model.ClientInstance
	global.AppDB.Where(&model.ClientInstance{CloudRes: res}).Find(&instances)
	return instances
}
