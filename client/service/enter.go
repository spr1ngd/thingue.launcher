package service

import (
	"thingue-launcher/client/service/instance"
	"thingue-launcher/client/service/server"
)

var (
	ServerConnManager        = server.ConnManager
	RunnerManager            = instance.RunnerManager
	InstanceManager          = instance.InstanceManager
	SyncManager              = instance.SyncManager
	RunnerRestartTaskManager = new(instance.RunnerRestartTaskManager)
)
