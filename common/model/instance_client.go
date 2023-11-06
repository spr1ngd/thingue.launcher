package model

import (
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"time"
)

type ClientInstance struct {
	CID             uint        `json:"cid" gorm:"primarykey"`
	SID             string      `json:"sid" gorm:"-"`
	Name            string      `json:"name"`
	ExecPath        string      `json:"execPath"`
	LaunchArguments StringSlice `json:"launchArguments"`
	Metadata        string      `json:"metadata"`
	PaksConfig      string      `json:"paksConfig"`
	FaultRecover    bool        `json:"faultRecover"`
	LastStartAt     time.Time   `json:"lastStartAt"`
	LastStopAt      time.Time   `json:"lastStopAt"`
	AutoControl     bool        `json:"autoControl"`
	StopDelay       int         `json:"stopDelay"`
	EnableH265      bool        `json:"enableH265"`
	AutoResizeRes   bool        `json:"autoResizeRes"`
	IsInternal      bool        `json:"isInternal"`
	CloudRes        string      `json:"cloudRes"`
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
