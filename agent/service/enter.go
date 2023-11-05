package service

import (
	"thingue-launcher/agent/service/instance"
	"thingue-launcher/agent/service/server"
)

var (
	ServerConnManager        = server.ConnManager
	RunnerManager            = instance.RunnerManager
	InstanceManager          = instance.InstanceManager
	SyncManager              = instance.SyncManager
	RunnerRestartTaskManager = new(instance.RunnerRestartTaskManager)
)
