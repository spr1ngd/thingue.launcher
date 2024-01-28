package utils

import (
	"thingue-launcher/common/logger"
	"thingue-launcher/common/model"
)

func GetInstanceNodeStatus(instance *model.Instance) int {
	if instance.StateCode == 1 {
		if instance.PakName == "" {
			return 0
		} else {
			if instance.PlayerCount > 0 {
				return 1
			} else {
				return 2
			}
		}
	} else if instance.StateCode == 0 {
		return 3
	} else if instance.StateCode == -1 {
		return 4
	}
	logger.Zap.Error("获取实例状态失败", instance.StreamerId)
	return -1
}
