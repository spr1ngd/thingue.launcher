package model

import "time"

type CloudRes struct {
	Name         string    `json:"name" gorm:"primarykey"`
	LastUpdateAt time.Time `json:"lastUpdateAt"`
}
