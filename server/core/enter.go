package core

import (
	"thingue-launcher/server/core/sdp"
	"thingue-launcher/server/core/service"
)

var (
	NodeService         = service.NodeService
	InstanceService     = service.InstanceService
	TicketService       = service.TicketService
	StreamerConnManager = sdp.StreamerConnManager
	PlayerConnManager   = sdp.PlayerConnManager
)
