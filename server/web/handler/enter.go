package handler

import (
	"thingue-launcher/server/web/handler/rest"
	"thingue-launcher/server/web/handler/ws"
)

var (
	InstanceGroup = new(rest.InstanceGroup)
	SyncGroup     = new(rest.SyncGroup)
	WsGroup       = new(ws.HandlerGroup)
)
