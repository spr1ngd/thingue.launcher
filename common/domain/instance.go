package domain

import (
	"github.com/mitchellh/mapstructure"
	"thingue-launcher/common/model"
	"time"
)

type Instance struct {
	ID                uint32             `json:"id"`
	Pid               int32              `json:"pid"`
	StateCode         int32              `json:"stateCode"`
	StreamerConnected bool               `json:"streamerConnected"`
	StreamerId        string             `json:"streamerId"`
	LastStartAt       time.Time          `json:"lastStartAt"`
	LastStopAt        time.Time          `json:"lastStopAt"`
	PlayerIds         []string           `json:"playerIds"`
	PlayerCount       uint32             `json:"playerCount"`
	IsInternal        bool               `json:"isInternal"`
	InstanceConfig    InstanceConfig     `json:"instanceConfig"`
	PlayerConfig      model.PlayerConfig `json:"playerConfig"`
}

type InstanceConfig struct {
	Name                   string   `json:"name"`
	CloudRes               string   `json:"cloudRes"`
	ExecPath               string   `json:"execPath"`
	LaunchArguments        []string `json:"launchArguments"`
	Metadata               string   `json:"metadata"`
	PaksConfig             string   `json:"paksConfig"`
	FaultRecover           bool     `json:"faultRecover"`
	EnableRelay            bool     `json:"enableRelay"`
	EnableRenderControl    bool     `json:"enableRenderControl"`
	EnableMultiuserControl bool     `json:"enableMultiuserControl"`
	AutoControl            bool     `json:"autoControl"`
	StopDelay              int32    `json:"stopDelay"`
}

func (i *Instance) ToServerModel() *model.ServerInstance { return nil }

func (i *Instance) FromServerModel(client *model.ServerInstance) {}

func (i *Instance) ToClientModel() *model.ClientInstance { return nil }

func (i *Instance) FromClientModel(client *model.ClientInstance) {
	var config InstanceConfig
	_ = mapstructure.Decode(client, &config)
	i.InstanceConfig = config
	i.ID = client.ID
	i.LastStartAt = client.LastStartAt
	i.LastStopAt = client.LastStopAt
	i.IsInternal = client.IsInternal
	i.PlayerConfig = client.PlayerConfig
}
