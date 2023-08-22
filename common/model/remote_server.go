package model

import "time"

type RemoteServer struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Url       string
	Enable    bool
}
