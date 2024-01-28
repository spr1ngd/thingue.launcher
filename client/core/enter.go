package core

import (
	"thingue-launcher/client/core/conn"
	"thingue-launcher/client/core/instance"
	"thingue-launcher/client/core/sync"
)

var (
	ConnManager              = conn.ConnManager
	RunnerManager            = instance.RunnerManager
	ConfigManager            = instance.ConfigManager
	SyncManager              = sync.SyncManager
	RunnerRestartTaskManager = new(instance.RunnerRestartTask)
)
