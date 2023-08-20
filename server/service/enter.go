package service

import (
	"thingue-launcher/server/service/instance"
	"thingue-launcher/server/service/sdp"
)

var (
	NodeService         = instance.NodeService
	InstanceService     = instance.InstanceService
	TicketService       = instance.TicketService
	StreamerConnManager = sdp.StreamerConnManager
	PlayerConnManager   = sdp.PlayerConnManager
)
