package service

import (
	"github.com/labstack/gommon/log"
	"thingue-launcher/common/model"
	"thingue-launcher/common/util"
	"thingue-launcher/sgcc-adapter/message"
	"thingue-launcher/sgcc-adapter/provider"
	"time"
)

type sgcc struct{}

var SGCC = sgcc{}

func (s *sgcc) GetNodeStatus(node string) int {
	instance, err := ThingUE.GetInstanceBySID(node)
	if err == nil {
		return s.GetInstanceStatus(instance)
	}
	log.Error("获取实例状态失败", instance.SID)
	return -1
}

func (s *sgcc) GetInstanceStatus(instance *model.ServerInstance) int {
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
	log.Error("获取实例状态失败", instance.SID)
	return -1
}

func (s *sgcc) Register() {
	maxNode := ThingUE.GetMaxNode()
	registerMessage := message.NewRegister(maxNode)
	provider.SendCloudMessage(registerMessage.GetBytes())
}

func (s *sgcc) Init() {
	// instance start + // 获取nodes状态
	var callback *message.InitCallback
	var nodes []message.Node
	instances, err := ThingUE.GetInstances()
	if err == nil {
		for _, instance := range instances {
			ThingUE.SendPrecessCommand(instance.SID, "START")
			node := message.Node{
				Id:       instance.SID,
				Status:   s.GetInstanceStatus(instance),
				Datetime: util.DateFormat(time.Now()),
				Station:  instance.PakName,
				LoadType: 3,
			}
			nodes = append(nodes, node)
		}
		callback = message.NewInitCallback(true, nodes)
	} else {
		callback = message.NewInitCallback(false, nil)
	}
	provider.SendCloudMessage(callback.GetBytes())
}

func (s *sgcc) Deploy(deploy *message.Deploy) {
	// instance pakload
	var callback *message.DeployCallback
	ThingUE.SendPakControl(deploy.Node, "load", deploy.Station)
	callback = message.NewDeployCallback(true)
	callback.Datetime = util.DateFormat(time.Now())
	callback.Node = deploy.Node
	callback.User = deploy.User
	callback.Status = s.GetNodeStatus(deploy.Node)
	provider.SendCloudMessage(callback.GetBytes())
}

func (s *sgcc) Release(nodes []string) {
	// instance pakunload
	var callback *message.ReleaseCallback
	var callbackNodes []*message.CallBackNode
	for _, node := range nodes {
		ThingUE.SendPakControl(node, "unload", "")
		callbackNodes = append(callbackNodes, &message.CallBackNode{
			Id:       node,
			Status:   s.GetNodeStatus(node),
			Datetime: util.DateFormat(time.Now()),
		})
	}
	callback = message.NewReleaseCallback(true, callbackNodes)
	callback.Nodes = callbackNodes
	provider.SendCloudMessage(callback.GetBytes())
}

func (s *sgcc) Status() {
	statistic := message.Statistic{}
	status := message.Status{
		Statistic: statistic,
	}
	provider.SendCloudMessage(status.GetBytes())
}

func (s *sgcc) Restart(nodes []string) {
	// instance restart
	var callback *message.RestartCallback
	var callbackNodes []*message.CallBackNode
	for _, node := range nodes {
		ThingUE.SendPrecessCommand(node, "RESTART")
		callbackNodes = append(callbackNodes, &message.CallBackNode{
			Id:       node,
			Status:   s.GetNodeStatus(node),
			Datetime: util.DateFormat(time.Now()),
		})
	}
	callback = message.NewRestartCallback(true, callbackNodes)
	provider.SendCloudMessage(callback.GetBytes())
}

func (s *sgcc) Kill(nodes []string) {
	// instance stop
	var callback *message.KillCallback
	var callbackNodes []*message.CallBackNode
	for _, node := range nodes {
		ThingUE.SendPrecessCommand(node, "STOP")
		callbackNodes = append(callbackNodes, &message.CallBackNode{
			Id:       node,
			Status:   s.GetNodeStatus(node),
			Datetime: util.DateFormat(time.Now()),
		})
	}
	callback = message.NewKillCallback(true, callbackNodes)
	provider.SendCloudMessage(callback.GetBytes())
}
