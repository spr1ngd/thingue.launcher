package model

import (
	"thingue-launcher/agent/model"
)

type AgentRegisterInfo struct {
	DeviceInfo *DeviceInfo
	Instances  []*model.Instance
}
