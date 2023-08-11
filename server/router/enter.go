package router

import (
	"thingue-launcher/server/router/api"
	"thingue-launcher/server/router/static"
	"thingue-launcher/server/router/ws"
)

type RouterGroup struct {
	Api    api.RouterGroup
	Ws     ws.RouterGroup
	Static static.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
