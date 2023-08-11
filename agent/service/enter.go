package service

import "thingue-launcher/agent/service/manager"

var (
	ServerManager            = new(manager.ServerManager)
	InstanceManager          = new(manager.InstanceManager)
	UeRunnerManager          = new(manager.RunnerManager)
	RunnerRestartTaskManager = new(manager.RunnerRestartTaskManager)
)
