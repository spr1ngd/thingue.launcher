package domain

import (
	"time"
)

type Instance struct {
	CID               uint      `json:"cid"`
	SID               string    `json:"sid"`
	Name              string    `json:"name"`
	ExecPath          string    `json:"execPath"`
	LaunchArguments   []string  `json:"launchArguments"`
	Metadata          string    `json:"metadata"`
	PaksConfig        string    `json:"paksConfig"`
	FaultRecover      bool      `json:"faultRecover"`
	EnableRelay       bool      `json:"enableRelay"`
	LastStartAt       time.Time `json:"lastStartAt"`
	LastStopAt        time.Time `json:"lastStopAt"`
	AutoControl       bool      `json:"autoControl"`
	StopDelay         int       `json:"stopDelay"`
	EnableH265        bool      `json:"enableH265"`
	AutoResizeRes     bool      `json:"autoResizeRes"`
	Pid               int       `json:"pid"`
	StateCode         int8      `json:"stateCode"`
	StreamerConnected bool      `json:"streamerConnected"`
	PlayerIds         []string  `json:"playerIds"`
	PlayerCount       uint      `json:"playerCount"`
	IsInternal        bool      `json:"isInternal"`
	CloudRes          string    `json:"cloudRes"`
}
