package model

import (
	"time"
)

type Instance struct {
	ID              uint `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string
	ExecPath        string
	LaunchArguments StringSlice
	Metadata        string
	PaksConfig      string
	FaultRecover    bool
	TimeRestart     bool
	RestartCron     string
}