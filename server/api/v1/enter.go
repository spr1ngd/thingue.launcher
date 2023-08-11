package v1

import (
	"thingue-launcher/server/api/v1/admin"
	"thingue-launcher/server/api/v1/agent"
)

type ApiGroup struct {
	Agent agent.RouterGroup
	Admin admin.RouterGroup
}

var ApiGroupApp = new(ApiGroup)
