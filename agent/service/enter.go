package service

import "thingue-launcher/agent/service/manager"

var (
	ServerConnManager        = new(manager.ServerConnManager)
	InstanceManager          = new(manager.InstanceManager)
	UeRunnerManager          = manager.RunnerManager
	RunnerRestartTaskManager = new(manager.RunnerRestartTaskManager)
)
