package model

import "time"

type CloudRes struct {
	Name         string      `json:"name" gorm:"primarykey"`
	Configs      StringSlice `json:"configs"`
	LastUpdateAt time.Time   `json:"lastUpdateAt"`
}
