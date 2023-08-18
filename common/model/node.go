package model

import "time"

type Node struct {
	ID        uint        `gorm:"primarykey" json:"id"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	Instances []*Instance `gorm:"foreignKey:NodeID" json:"instances"`
	*DeviceInfo
}

type NodeRegisterInfo struct {
	NodeID     uint
	DeviceInfo *DeviceInfo
	Instances  []*Instance
}
