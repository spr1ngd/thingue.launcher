package model

import (
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"k8s.io/apimachinery/pkg/labels"
	"thingue-launcher/common/domain"
	"time"
)

type ServerInstance struct {
	CID                 uint                `json:"cid" gorm:"primarykey"`
	ClientID            uint                `json:"ClientID" gorm:"primarykey"`
	SID                 string              `json:"sid" gorm:"unique"`
	Name                string              `json:"name"`
	ExecPath            string              `json:"execPath"`
	LaunchArguments     StringSlice         `json:"launchArguments"`
	Metadata            string              `json:"metadata"`
	PaksConfig          string              `json:"paksConfig"`
	FaultRecover        bool                `json:"faultRecover"`
	EnableRelay         bool                `json:"enableRelay"`
	EnableRenderControl bool                `json:"enableRenderControl"`
	LastStartAt         time.Time           `json:"lastStartAt"`
	LastStopAt          time.Time           `json:"lastStopAt"`
	Pid                 int                 `json:"pid"`
	StateCode           int8                `json:"stateCode"`
	StreamerConnected   bool                `json:"streamerConnected"`
	PlayerIds           UintSlice           `json:"playerIds"`
	PlayerCount         uint                `json:"playerCount"`
	CurrentPak          string              `json:"currentPak"`
	Rendering           bool                `json:"rendering"`
	AutoControl         bool                `json:"autoControl"`
	StopDelay           uint                `json:"stopDelay"`
	AutoResizeRes       bool                `json:"autoResizeRes"`
	Labels              labels.Labels       `json:"labels" gorm:"-"`
	Paks                []domain.Pak        `json:"paks" gorm:"-"`
	PakName             string              `json:"pakName" gorm:"-"`
	CloudRes            string              `json:"cloudRes"`
	PlayerConfig        domain.PlayerConfig `json:"playerConfig" gorm:"serializer:json"`
}

func (instance *ServerInstance) AfterFind(tx *gorm.DB) (err error) {
	if instance.Metadata != "" {
		var metaData domain.MetaData
		err := yaml.Unmarshal([]byte(instance.Metadata), &metaData)
		if err == nil {
			instance.Labels = labels.Set(metaData.Labels)
		}
	}
	if instance.PaksConfig != "" {
		var paksConfig domain.PakConfig
		err = yaml.Unmarshal([]byte(instance.PaksConfig), &paksConfig)
		if err == nil {
			instance.Paks = paksConfig.Paks
			if instance.CurrentPak != "" {
				for _, pak := range paksConfig.Paks {
					if pak.Value == instance.CurrentPak {
						instance.PakName = pak.Name
					}
				}
			}
		}
	}
	return
}
