package model

import (
	"k8s.io/apimachinery/pkg/labels"
	"time"
)

type ServerInstance struct {
	ID                uint          `json:"id" gorm:"primarykey"`
	NodeID            uint          `json:"nodeID" gorm:"primarykey"`
	SID               string        `json:"sid" gorm:"unique"`
	CreatedAt         time.Time     `json:"createdAt"`
	UpdatedAt         time.Time     `json:"updatedAt"`
	Name              string        `json:"name"`
	ExecPath          string        `json:"execPath"`
	LaunchArguments   StringSlice   `json:"launchArguments"`
	Metadata          string        `json:"metadata"`
	PaksConfig        string        `json:"paksConfig"`
	FaultRecover      bool          `json:"faultRecover"`
	LastStartAt       time.Time     `json:"lastStartAt"`
	LastStopAt        time.Time     `json:"lastStopAt"`
	Pid               int           `json:"pid"`
	StateCode         int8          `json:"stateCode"`
	StreamerConnected bool          `json:"streamerConnected"`
	PlayerIds         StringSlice   `json:"playerIds"`
	Labels            labels.Labels `json:"labels" gorm:"-"`
}
