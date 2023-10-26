package service

import (
	"thingue-launcher/agent/service/instance"
	"thingue-launcher/agent/service/server"
)

var (
	ServerConnManager        = server.ConnManager
	RunnerManager            = instance.RunnerManager
	InstanceManager          = instance.InstanceManager
	RunnerRestartTaskManager = new(instance.RunnerRestartTaskManager)
)
