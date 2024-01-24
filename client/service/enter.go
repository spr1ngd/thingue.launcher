package service

import (
	"thingue-launcher/client/service/instance"
	"thingue-launcher/client/service/server"
)

var (
	ConnManager              = server.ConnManager
	GrpcClient               = server.GrpcClient
	RunnerManager            = instance.RunnerManager
	InstanceManager          = instance.InstanceManager
	SyncManager              = instance.SyncManager
	RunnerRestartTaskManager = new(instance.RunnerRestartTaskManager)
)
