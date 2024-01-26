package response

import (
	"github.com/mitchellh/mapstructure"
	"k8s.io/apimachinery/pkg/labels"
	"thingue-launcher/common/model"
)

type InstanceTicket struct {
	ID                uint               `json:"id"`
	StreamerId        string             `json:"streamerId"`
	Name              string             `json:"name"`
	ExecPath          string             `json:"execPath"`
	Pid               int                `json:"pid"`
	StateCode         int8               `json:"stateCode"`
	StreamerConnected bool               `json:"streamerConnected"`
	Labels            labels.Labels      `json:"labels"`
	Ticket            string             `json:"ticket"`
	PlayerConfig      model.PlayerConfig `json:"playerConfig"`
}

func (t *InstanceTicket) SetInstanceInfo(instance *model.ServerInstance) {
	mapstructure.Decode(instance, t)
}
