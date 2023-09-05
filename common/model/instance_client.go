package model

import (
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"time"
)

type ClientInstance struct {
	CID               uint        `json:"cid" gorm:"primarykey"`
	SID               string      `json:"sid" gorm:"-"`
	Name              string      `json:"name"`
	ExecPath          string      `json:"execPath"`
	LaunchArguments   StringSlice `json:"launchArguments"`
	Metadata          string      `json:"metadata"`
	PaksConfig        string      `json:"paksConfig"`
	FaultRecover      bool        `json:"faultRecover"`
	LastStartAt       time.Time   `json:"lastStartAt"`
	LastStopAt        time.Time   `json:"lastStopAt"`
	AutoControl       bool        `json:"autoControl"`
	StopDelay         int         `json:"stopDelay"`
	Pid               int         `json:"pid" gorm:"-"`
	StateCode         int8        `json:"stateCode" gorm:"-"`
	StreamerConnected bool        `json:"streamerConnected" gorm:"-"`
	PlayerIds         []string    `json:"playerIds" gorm:"-"`
	PlayerCount       uint        `json:"playerCount" gorm:"-"`
}

func (clientInstance *ClientInstance) ToServerInstance() *ServerInstance {
	var serverInstance *ServerInstance
	mapstructure.Decode(clientInstance, &serverInstance)
	if serverInstance.SID == "" {
		sid, _ := uuid.NewUUID()
		serverInstance.SID = sid.String()
	}
	return serverInstance
}
