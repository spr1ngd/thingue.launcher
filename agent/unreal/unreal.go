package unreal

import (
	"context"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/model"
	"thingue-launcher/common/config"
)

type Unreal struct {
	ctx context.Context
}

func NewUnreal() *Unreal {
	return &Unreal{}
}

func (u *Unreal) SetContext(ctx context.Context) {
	u.ctx = ctx
}

func (u *Unreal) ListInstance() []model.Instance {
	var instances []model.Instance
	global.APP_DB.Find(&instances)
	return instances
}

func (u *Unreal) CreateInstance(instance model.Instance) uint {
	global.APP_DB.Create(&instance)
	return instance.ID
}

func (u *Unreal) SaveInstance(instance model.Instance) uint {
	global.APP_DB.Save(&instance)
	return instance.ID
}

func (u *Unreal) DeleteInstance(id uint) error {
	process := GetProcessById(id)
	if process != nil {
		err := process.stop()
		if err != nil {
			return err
		}
	} else {
		global.APP_DB.Delete(&model.Instance{}, id)
	}
	return nil
}

func (u *Unreal) GetInstanceById(id uint) *model.Instance {
	instance := model.Instance{}
	global.APP_DB.First(&instance, id)
	return &instance
}

func (u *Unreal) StartInstance(id uint) error {
	instance := u.GetInstanceById(id)
	process := NewProcess(instance)
	appConfig := config.GetAppConfig()
	if appConfig.ServerUrl != "" {
		process.LaunchArguments = append(process.LaunchArguments, "-PixelStreamingURL="+appConfig.ServerUrl+"/ws/streamer/"+instance.Name)
	}
	err := process.start()
	if err != nil {
		instance.Status = 1
	}
	return err
}

func (u *Unreal) StopInstance(id uint) error {
	process := GetProcessById(id)
	err := process.stop()
	if err != nil {
		instance := u.GetInstanceById(id)
		instance.Status = 0
	}
	return err
}
