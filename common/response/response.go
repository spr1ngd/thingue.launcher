package response

import (
	"github.com/mitchellh/mapstructure"
	"k8s.io/apimachinery/pkg/labels"
	"thingue-launcher/common/domain"
	"thingue-launcher/common/model"
)

type InstanceTicket struct {
	ID                uint                `json:"cid"`
	StreamerId        string              `json:"sid"`
	Name              string              `json:"name"`
	ExecPath          string              `json:"execPath"`
	Pid               int                 `json:"pid"`
	StateCode         int8                `json:"stateCode"`
	StreamerConnected bool                `json:"streamerConnected"`
	Labels            labels.Labels       `json:"labels"`
	Ticket            string              `json:"ticket"`
	PlayerConfig      domain.PlayerConfig `json:"playerConfig"`
}

func (t *InstanceTicket) SetInstanceInfo(instance *model.Instance) {
	mapstructure.Decode(instance, t)
}

type ServerInstance struct {
	CID                 uint32              `json:"cid"`
	ClientID            uint32              `json:"ClientID"`
	SID                 string              `json:"sid"`
	Name                string              `json:"name"`
	ExecPath            string              `json:"execPath"`
	Metadata            string              `json:"metadata"`
	FaultRecover        bool                `json:"faultRecover"`
	EnableRelay         bool                `json:"enableRelay"`
	EnableRenderControl bool                `json:"enableRenderControl"`
	Pid                 int32               `json:"pid"`
	StateCode           int32               `json:"stateCode"`
	StreamerConnected   bool                `json:"streamerConnected"`
	PlayerIds           []uint32            `json:"playerIds"`
	PlayerCount         uint32              `json:"playerCount"`
	CurrentPak          string              `json:"currentPak"`
	Rendering           bool                `json:"rendering"`
	AutoControl         bool                `json:"autoControl"`
	StopDelay           uint32              `json:"stopDelay"`
	Labels              labels.Labels       `json:"labels"`
	PakName             string              `json:"pakName"`
	CloudRes            string              `json:"cloudRes"`
	PlayerConfig        domain.PlayerConfig `json:"playerConfig"`
}

func (s *ServerInstance) FromModel(instance *model.Instance) {
	s.CID = instance.ID
	s.ClientID = instance.ClientID
	s.SID = instance.StreamerId

	s.PlayerConfig.HideUI = instance.HideUI
	s.PlayerConfig.MatchViewportRes = instance.MatchViewportRes
	s.PlayerConfig.IdleDisconnect = instance.IdleDisconnect
	s.PlayerConfig.IdleTimeout = instance.IdleTimeout

	_ = mapstructure.Decode(instance, s)
}
