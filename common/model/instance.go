package model

import (
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"k8s.io/apimachinery/pkg/labels"
	"time"
)

type Instance struct {
	ID                uint32    `json:"id" gorm:"primarykey;autoIncrement:false"`
	ClientID          uint32    `json:"clientID" gorm:"primarykey;autoIncrement:false"`
	StreamerId        string    `json:"streamerId" gorm:"unique"`
	Pid               int32     `json:"pid"`
	StateCode         int32     `json:"stateCode"`
	StreamerConnected bool      `json:"streamerConnected"`
	PlayerIds         UintSlice `json:"playerIds"`
	PlayerCount       uint32    `json:"playerCount"`
	PakName           string    `json:"pakName" gorm:"-"`
	PakValue          string    `json:"pakValue"`
	Rendering         bool      `json:"rendering"`
	LastStartAt       time.Time `json:"lastStartAt"`
	LastStopAt        time.Time `json:"lastStopAt"`

	Name                string      `json:"name"`
	CloudRes            string      `json:"cloudRes"`
	ExecPath            string      `json:"execPath"`
	LaunchArguments     StringSlice `json:"launchArguments"`
	Metadata            string      `json:"metadata"`
	PaksConfig          string      `json:"paksConfig"`
	FaultRecover        bool        `json:"faultRecover"`
	EnableRelay         bool        `json:"enableRelay"`
	EnableRenderControl bool        `json:"enableRenderControl"`
	AutoControl         bool        `json:"autoControl"`
	StopDelay           uint32      `json:"stopDelay"`

	MatchViewportRes bool   `json:"matchViewportRes"`
	HideUI           bool   `json:"hideUI"`
	IdleDisconnect   bool   `json:"idleDisconnect"`
	IdleTimeout      uint32 `json:"idleTimeout"`

	Labels labels.Labels `json:"labels" gorm:"-"`
	Paks   []Pak         `json:"paks" gorm:"-"`
}

func (instance *Instance) AfterFind(tx *gorm.DB) (err error) {
	if instance.Metadata != "" {
		var metaData MetaData
		err := yaml.Unmarshal([]byte(instance.Metadata), &metaData)
		if err == nil {
			instance.Labels = labels.Set(metaData.Labels)
		}
	}
	if instance.PaksConfig != "" {
		var paksConfig PakConfig
		err = yaml.Unmarshal([]byte(instance.PaksConfig), &paksConfig)
		if err == nil {
			instance.Paks = paksConfig.Paks
			if instance.PakValue != "" {
				for _, pak := range paksConfig.Paks {
					if pak.Value == instance.PakValue {
						instance.PakName = pak.Name
					}
				}
			}
		}
	}
	return
}
