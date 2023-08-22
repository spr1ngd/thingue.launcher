package response

import "thingue-launcher/common/model"

type InstanceTicket struct {
	*model.ServerInstance
	Ticket string `json:"ticket"`
}
