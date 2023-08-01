package model

import "gorm.io/gorm"

type RemoteServer struct {
	gorm.Model
	Url    string
	Enable bool
}
