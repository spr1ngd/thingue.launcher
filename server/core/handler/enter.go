package handler

import (
	"thingue-launcher/server/core/handler/mqtt"
	"thingue-launcher/server/core/handler/rest"
	"thingue-launcher/server/core/handler/ws"
)

var (
	InstanceGroup = new(rest.InstanceGroup)
	SyncGroup     = new(rest.SyncGroup)
	WsGroup       = new(ws.HandlerGroup)
	MqttHandler   = mqtt.MqttHandler
)
