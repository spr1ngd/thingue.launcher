package domain

import (
	"time"
)

type Instance struct {
	CID                    uint         `json:"cid"`
	SID                    string       `json:"sid"`
	Name                   string       `json:"name"`
	ExecPath               string       `json:"execPath"`
	LaunchArguments        []string     `json:"launchArguments"`
	Metadata               string       `json:"metadata"`
	PaksConfig             string       `json:"paksConfig"`
	FaultRecover           bool         `json:"faultRecover"`
	EnableRelay            bool         `json:"enableRelay"`
	EnableRenderControl    bool         `json:"enableRenderControl"`
	EnableMultiuserControl bool         `json:"enableMultiuserControl"`
	LastStartAt            time.Time    `json:"lastStartAt"`
	LastStopAt             time.Time    `json:"lastStopAt"`
	AutoControl            bool         `json:"autoControl"`
	StopDelay              int          `json:"stopDelay"`
	Pid                    int          `json:"pid"`
	StateCode              int8         `json:"stateCode"`
	StreamerConnected      bool         `json:"streamerConnected"`
	PlayerIds              []string     `json:"playerIds"`
	PlayerCount            uint         `json:"playerCount"`
	IsInternal             bool         `json:"isInternal"`
	CloudRes               string       `json:"cloudRes"`
	PlayerConfig           PlayerConfig `json:"playerConfig"`
}

type PlayerConfig struct {
	MatchViewportRes bool `json:"matchViewportRes"`
	HideUI           bool `json:"hideUI"`
	IdleDisconnect   bool `json:"idleDisconnect"`
	IdleTimeout      uint `json:"idleTimeout"`
}
