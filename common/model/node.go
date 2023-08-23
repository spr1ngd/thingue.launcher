package model

import (
	"github.com/mitchellh/mapstructure"
	"thingue-launcher/common/domain"
	"time"
)

type Node struct {
	ID         uint              `json:"id" gorm:"primarykey"`
	CreatedAt  time.Time         `json:"createdAt"`
	UpdatedAt  time.Time         `json:"updatedAt"`
	Instances  []*ServerInstance `json:"instances" gorm:"foreignKey:NodeID"`
	Version    string            `json:"version"`
	Workdir    string            `json:"workdir"`
	Hostname   string            `json:"hostname"`
	Memory     string            `json:"memory"`
	Cpus       StringSlice       `json:"cpus"`
	Gpus       StringSlice       `json:"gpus"`
	Ips        StringSlice       `json:"ips"`
	OsArch     string            `json:"osArch"`
	OsType     string            `json:"osType"`
	SystemUser string            `json:"systemUser"`
}

func (node *Node) SetDeviceInfo(info domain.DeviceInfo) {
	mapstructure.Decode(&info, node)
}
