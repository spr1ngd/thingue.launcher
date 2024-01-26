package model

import (
	"time"
)

type ClientInstance struct {
	ID                     uint32       `json:"id" gorm:"primarykey;column:c_id"`
	Name                   string       `json:"name"`
	ExecPath               string       `json:"execPath"`
	LaunchArguments        StringSlice  `json:"launchArguments"`
	Metadata               string       `json:"metadata"`
	PaksConfig             string       `json:"paksConfig"`
	FaultRecover           bool         `json:"faultRecover"`
	EnableRelay            bool         `json:"enableRelay"`
	EnableRenderControl    bool         `json:"enableRenderControl"`
	EnableMultiuserControl bool         `json:"enableMultiuserControl"`
	LastStartAt            time.Time    `json:"lastStartAt"`
	LastStopAt             time.Time    `json:"lastStopAt"`
	AutoControl            bool         `json:"autoControl"`
	StopDelay              int32        `json:"stopDelay"`
	AutoResizeRes          bool         `json:"autoResizeRes"`
	IsInternal             bool         `json:"isInternal"`
	CloudRes               string       `json:"cloudRes"`
	PlayerConfig           PlayerConfig `json:"playerConfig" gorm:"serializer:json"`
}
