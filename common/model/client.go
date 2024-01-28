package model

import (
	"time"
)

type Client struct {
	ID         uint32            `json:"id" gorm:"primarykey"`
	CreatedAt  time.Time         `json:"createdAt"`
	UpdatedAt  time.Time         `json:"updatedAt"`
	Instances  []*ServerInstance `json:"instances" gorm:"foreignKey:ClientID"`
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
