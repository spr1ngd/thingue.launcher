package service

import (
	"github.com/labstack/gommon/log"
	"thingue-launcher/common/model"
	"thingue-launcher/common/request"
	"thingue-launcher/common/util"
	"thingue-launcher/server/core/service"
	"thingue-launcher/server/sgcc/message"
	"thingue-launcher/server/sgcc/provider"
	"time"
)

type sgccService struct{}

var SgccService = &sgccService{}

func (s *sgccService) GetNodeStatus(node string) int {
	instance := service.InstanceService.GetInstanceBySid(node)
	if instance != nil {
		return s.GetInstanceStatus(instance)
	}
	log.Error("获取实例状态失败", instance.SID)
	return -1
}

func (s *sgccService) GetInstanceStatus(instance *model.ServerInstance) int {
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

func (s *sgccService) Register() {
	maxNode := len(service.InstanceService.InstanceList())
	registerMessage := message.NewRegister(maxNode)
	provider.SendCloudMessage(registerMessage.GetBytes())
}

func (s *sgccService) Init() {
	// instance start + // 获取nodes状态
	var callback *message.InitCallback
	var nodes []message.Node
	instances := service.InstanceService.InstanceList()
	for _, instance := range instances {
		service.InstanceService.ProcessControl(request.ProcessControl{
			SID:     instance.SID,
			Command: "START",
		})
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
	//callback = message.NewInitCallback(false, nil)
	provider.SendCloudMessage(callback.GetBytes())
}

func (s *sgccService) Deploy(deploy *message.Deploy) {
	// instance pakload
	var callback *message.DeployCallback
	_ = service.InstanceService.PakControl(request.PakControl{
		SID:  deploy.Node,
		Type: "load",
		Pak:  deploy.Station,
	})
	callback = message.NewDeployCallback(true)
	callback.Datetime = util.DateFormat(time.Now())
	callback.Node = deploy.Node
	callback.User = deploy.User
	callback.Status = s.GetNodeStatus(deploy.Node)
	provider.SendCloudMessage(callback.GetBytes())
}

func (s *sgccService) Release(nodes []string) {
	// instance pakunload
	var callback *message.ReleaseCallback
	var callbackNodes []*message.CallBackNode
	for _, node := range nodes {
		_ = service.InstanceService.PakControl(request.PakControl{
			SID:  node,
			Type: "unload",
			Pak:  "",
		})
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

func (s *sgccService) Status() {
	statistic := message.Statistic{}
	status := message.Status{
		Statistic: statistic,
	}
	provider.SendCloudMessage(status.GetBytes())
}

func (s *sgccService) Restart(nodes []string) {
	// instance restart
	var callback *message.RestartCallback
	var callbackNodes []*message.CallBackNode
	for _, node := range nodes {
		service.InstanceService.ProcessControl(request.ProcessControl{
			SID:     node,
			Command: "RESTART",
		})
		callbackNodes = append(callbackNodes, &message.CallBackNode{
			Id:       node,
			Status:   s.GetNodeStatus(node),
			Datetime: util.DateFormat(time.Now()),
		})
	}
	callback = message.NewRestartCallback(true, callbackNodes)
	provider.SendCloudMessage(callback.GetBytes())
}

func (s *sgccService) Kill(nodes []string) {
	// instance stop
	var callback *message.KillCallback
	var callbackNodes []*message.CallBackNode
	for _, node := range nodes {
		service.InstanceService.ProcessControl(request.ProcessControl{
			SID:     node,
			Command: "STOP",
		})
		callbackNodes = append(callbackNodes, &message.CallBackNode{
			Id:       node,
			Status:   s.GetNodeStatus(node),
			Datetime: util.DateFormat(time.Now()),
		})
	}
	callback = message.NewKillCallback(true, callbackNodes)
	provider.SendCloudMessage(callback.GetBytes())
}
