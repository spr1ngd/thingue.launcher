package service

import (
	"thingue-launcher/server/service/instance"
	"thingue-launcher/server/service/sdp"
)

var (
	NodeService     = instance.NodeService
	InstanceService = instance.Service
	StreamerManager = sdp.StreamerManager
	PlayerManager   = sdp.PlayerManager
)
