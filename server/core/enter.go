package core

import (
	"thingue-launcher/server/core/sdp"
	"thingue-launcher/server/core/service"
)

var (
	NodeService         = service.NodeService
	InstanceService     = service.InstanceService
	TicketService       = service.TicketService
	SyncService         = service.SyncService
	StreamerConnManager = sdp.StreamerConnManager
	PlayerConnManager   = sdp.PlayerConnManager
)
