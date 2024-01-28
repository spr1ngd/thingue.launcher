package domain

import (
	"github.com/mitchellh/mapstructure"
	types "thingue-launcher/common/gen/proto/go/types/v1"
	"thingue-launcher/common/model"
	"time"
)

type Instance struct {
	ID                uint32         `json:"id"`
	Pid               int32          `json:"pid"`
	StateCode         int32          `json:"stateCode"`
	StreamerConnected bool           `json:"streamerConnected"`
	StreamerId        string         `json:"streamerId"`
	LastStartAt       time.Time      `json:"lastStartAt"`
	LastStopAt        time.Time      `json:"lastStopAt"`
	PlayerIds         []string       `json:"playerIds"`
	PlayerCount       uint32         `json:"playerCount"`
	IsInternal        bool           `json:"isInternal"`
	Config            InstanceConfig `json:"instanceConfig"`
	PlayerConfig      PlayerConfig   `json:"playerConfig"`
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

type PlayerConfig struct {
	MatchViewportRes bool   `json:"matchViewportRes"`
	HideUI           bool   `json:"hideUI"`
	IdleDisconnect   bool   `json:"idleDisconnect"`
	IdleTimeout      uint32 `json:"idleTimeout"`
}

func (i *Instance) ToServerModel() *model.ServerInstance { return nil }

func (i *Instance) FromServerModel(client *model.ServerInstance) {}

func (i *Instance) ToInstanceInfoTypes() *types.InstanceInfo {
	var instanceInfo types.InstanceInfo
	_ = mapstructure.Decode(i, &instanceInfo)
	return &instanceInfo
}

func (i *Instance) ToInstanceConfig() *model.InstanceConfig {
	var instanceConfig model.InstanceConfig
	_ = mapstructure.Decode(i, &instanceConfig)
	_ = mapstructure.Decode(i.Config, &instanceConfig)
	_ = mapstructure.Decode(i.PlayerConfig, &instanceConfig)
	return &instanceConfig
}

func (i *Instance) FromInstanceConfig(instanceConfig *model.InstanceConfig) {
	var config InstanceConfig
	_ = mapstructure.Decode(instanceConfig, &config)
	i.Config = config
	var playerConfig PlayerConfig
	_ = mapstructure.Decode(instanceConfig, &playerConfig)
	i.PlayerConfig = playerConfig
	i.ID = instanceConfig.ID
	i.LastStartAt = instanceConfig.LastStartAt
	i.LastStopAt = instanceConfig.LastStopAt
	i.IsInternal = instanceConfig.IsInternal
}
