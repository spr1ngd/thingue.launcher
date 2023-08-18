package handler

import (
	"thingue-launcher/server/handler/instance"
	"thingue-launcher/server/handler/ws"
)

var (
	InstanceGroup = new(instance.HandlerGroup)
	WsGroup       = new(ws.HandlerGroup)
)
