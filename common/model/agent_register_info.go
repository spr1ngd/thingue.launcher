package model

import (
	"thingue-launcher/common/model"
)

type AgentRegisterInfo struct {
	DeviceInfo *DeviceInfo
	Instances  []*model.Instance
}
