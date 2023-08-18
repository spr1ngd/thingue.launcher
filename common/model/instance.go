package model

import (
	"time"
)

type Instance struct {
	ID              uint        `gorm:"primarykey" json:"id"`
	NodeID          uint        `gorm:"primarykey" json:"nodeID"`
	SID             string      `gorm:"unique;default:gen_random_uuid()"`
	CreatedAt       time.Time   `json:"createdAt"`
	UpdatedAt       time.Time   `json:"updatedAt"`
	Name            string      `json:"name"`
	ExecPath        string      `json:"execPath"`
	LaunchArguments StringSlice `json:"launchArguments"`
	Metadata        string      `json:"metadata"`
	PaksConfig      string      `json:"paksConfig"`
	FaultRecover    bool        `json:"faultRecover"`
	TimeRestart     bool        `json:"timeRestart"`
	RestartCron     string      `json:"restartCron"`
	LastStartAt     time.Time   `json:"lastStartAt"`
	LastStopAt      time.Time   `json:"lastStopAt"`
	Pid             int         `gorm:"-" json:"pid"`
	StateCode       int8        `gorm:"-" json:"stateCode"`
}
