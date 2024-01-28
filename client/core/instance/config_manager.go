package instance

import (
	"errors"
	"thingue-launcher/client/global"
	"thingue-launcher/common/domain"
	"thingue-launcher/common/model"

	"gorm.io/gorm"
)

type configManager struct{}

var ConfigManager = new(configManager)

func (m *configManager) List() []*model.InstanceConfig {
	var instances []*model.InstanceConfig
	global.AppDB.Find(&instances)
	return instances
}

func (m *configManager) Create(instance *domain.Instance) uint32 {
	instanceConfig := instance.ToInstanceConfig()
	global.AppDB.Create(&instanceConfig)
	return instanceConfig.ID
}

func (m *configManager) GetById(id uint) *domain.Instance {
	var instanceModel model.InstanceConfig
	result := global.AppDB.First(&instanceModel, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	var instance domain.Instance
	instance.FromInstanceConfig(&instanceModel)
	return &instance
}

func (m *configManager) GetInternal() (*domain.Instance, error) {
	var instanceModel model.InstanceConfig
	result := global.AppDB.Where(&model.InstanceConfig{IsInternal: true}).First(&instanceModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	var instance domain.Instance
	instance.FromInstanceConfig(&instanceModel)
	return &instance, nil
}

func (m *configManager) Update(instance *domain.Instance) error {
	runner, err := RunnerManager.GetRunnerById(instance.ID)
	if err != nil {
		return err
	}
	if runner.StateCode == 1 {
		return errors.New("实例运行中无法修改配置")
	}
	global.AppDB.Save(instance.ToInstanceConfig())
	runner.Config = instance.Config
	return nil
}

func (m *configManager) Delete(id uint32) {
	global.AppDB.Delete(&model.InstanceConfig{}, id)
}

func (m *configManager) GetByCloudRes(res string) []model.InstanceConfig {
	var instances []model.InstanceConfig
	global.AppDB.Where(&model.InstanceConfig{CloudRes: res}).Find(&instances)
	return instances
}

func (m *configManager) GetDefault() *domain.Instance {
	return &domain.Instance{
		Config: domain.InstanceConfig{
			LaunchArguments:        []string{"-AudioMixer", "-RenderOffScreen", "-ForceRes", "-ResX=1920", "-ResY=1080"},
			FaultRecover:           false,
			EnableRelay:            true,
			EnableRenderControl:    false,
			EnableMultiuserControl: false,
			AutoControl:            false,
			StopDelay:              3,
		},
		PlayerConfig: domain.PlayerConfig{
			MatchViewportRes: true,
			HideUI:           false,
			IdleDisconnect:   false,
			IdleTimeout:      5,
		},
	}
}
