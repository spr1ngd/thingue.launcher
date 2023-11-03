package model

import "time"

type CloudResource struct {
	Name         string    `json:"name" gorm:"primarykey"`
	LastUpdateAt time.Time `json:"lastUpdateAt"`
}
