package handler

import (
	"thingue-launcher/server/web/handler/instance"
	"thingue-launcher/server/web/handler/ws"
)

var (
	InstanceGroup = new(instance.HandlerGroup)
	WsGroup       = new(ws.HandlerGroup)
)