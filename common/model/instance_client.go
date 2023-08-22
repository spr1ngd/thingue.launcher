package model

import "time"

type ClientInstance struct {
	ID              uint        `json:"id" gorm:"primarykey"`
	Name            string      `json:"name"`
	ExecPath        string      `json:"execPath"`
	LaunchArguments StringSlice `json:"launchArguments"`
	Metadata        string      `json:"metadata"`
	PaksConfig      string      `json:"paksConfig"`
	FaultRecover    bool        `json:"faultRecover"`
	Pid             int         `json:"pid" gorm:"-"`
	StateCode       int8        `json:"stateCode" gorm:"-"`
	SID             string      `json:"sid" gorm:"-"`
	LastStartAt     time.Time   `json:"lastStartAt"`
	LastStopAt      time.Time   `json:"lastStopAt"`
}
