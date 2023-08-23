package model

import (
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"k8s.io/apimachinery/pkg/labels"
	"thingue-launcher/common/domain"
	"time"
)

type ServerInstance struct {
	CID               uint          `json:"cid" gorm:"primarykey"`
	NodeID            uint          `json:"nodeID" gorm:"primarykey"`
	SID               string        `json:"sid" gorm:"unique"`
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
	PlayerCount       uint          `json:"playerCount"`
	Labels            labels.Labels `json:"labels" gorm:"-"`
	Paks              []domain.Pak  `json:"paks" gorm:"-"`
}

func (i *ServerInstance) AfterFind(tx *gorm.DB) (err error) {
	if i.Metadata != "" {
		var metaData domain.MetaData
		err := yaml.Unmarshal([]byte(i.Metadata), &metaData)
		if err == nil {
			i.Labels = labels.Set(metaData.Labels)
		}
	}
	if i.PaksConfig != "" {
		var paksConfig domain.PakConfig
		err = yaml.Unmarshal([]byte(i.PaksConfig), &paksConfig)
		if err == nil {
			i.Paks = paksConfig.Paks
		}
	}
	return
}
