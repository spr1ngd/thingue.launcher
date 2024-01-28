package model

import (
	"time"
)

type InstanceConfig struct {
	ID                     uint32      `json:"id" gorm:"primarykey"`
	Name                   string      `json:"name"`
	ExecPath               string      `json:"execPath"`
	LaunchArguments        StringSlice `json:"launchArguments"`
	Metadata               string      `json:"metadata"`
	PaksConfig             string      `json:"paksConfig"`
	FaultRecover           bool        `json:"faultRecover"`
	EnableRelay            bool        `json:"enableRelay"`
	EnableRenderControl    bool        `json:"enableRenderControl"`
	EnableMultiuserControl bool        `json:"enableMultiuserControl"`
	LastStartAt            time.Time   `json:"lastStartAt"`
	LastStopAt             time.Time   `json:"lastStopAt"`
	AutoControl            bool        `json:"autoControl"`
	StopDelay              int32       `json:"stopDelay"`
	AutoResizeRes          bool        `json:"autoResizeRes"`
	IsInternal             bool        `json:"isInternal"`
	CloudRes               string      `json:"cloudRes"`
	MatchViewportRes       bool        `json:"matchViewportRes"`
	HideUI                 bool        `json:"hideUI"`
	IdleDisconnect         bool        `json:"idleDisconnect"`
	IdleTimeout            uint32      `json:"idleTimeout"`
}
