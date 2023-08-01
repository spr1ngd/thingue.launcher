package unreal

import (
	"context"
	"thingue-launcher/agent/global"
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

func (u *Unreal) ListInstance() []Instance {
	var instances []Instance
	//global.APP_DB.Find(&instances)
	return instances
}

func (u *Unreal) AddInstance(instance Instance) uint {
	global.APP_DB.Create(&instance)
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
		global.APP_DB.Delete(&Instance{}, id)
	}
	return nil
}

func (u *Unreal) GetInstanceById(id uint) *Instance {
	instance := Instance{}
	global.APP_DB.First(&instance, id)
	return &instance
}

func (u *Unreal) UpdateInstance(instance Instance) {
	global.APP_DB.Save(instance)
}

func (u *Unreal) StartInstance(id uint) error {
	instance := u.GetInstanceById(id)
	process := NewProcess(instance)
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
