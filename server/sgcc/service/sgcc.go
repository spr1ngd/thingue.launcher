package service

import (
	"context"
	"math"
	"math/rand"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	types "thingue-launcher/common/gen/proto/go/types/v1"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/model"
	"thingue-launcher/common/request"
	"thingue-launcher/common/util"
	coreprovider "thingue-launcher/server/core/provider"
	"thingue-launcher/server/core/service"
	"thingue-launcher/server/sgcc/message"
	"thingue-launcher/server/sgcc/provider"
	"time"
)

type sgccService struct{}

var SgccService = &sgccService{}

func (s *sgccService) GetNodeStatus(node string) int {
	instance := service.InstanceService.GetInstanceByStreamerId(node)
	if instance != nil {
		return s.GetInstanceStatus(instance)
	}
	logger.Zap.Error("获取实例状态失败", instance.StreamerId)
	return -1
}

func (s *sgccService) GetInstanceStatus(instance *model.Instance) int {
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
		client, err := coreprovider.GrpcClientProvider.GetClient(instance.ClientID)
		if err == nil {
			_, err := client.ControlProcess(context.Background(), &pb.ControlProcessRequest{
				InstanceId: instance.ID,
				Command:    types.Command_COMMAND_START,
			})
			if err != nil {
				logger.Zap.Error("负载初始化时启动实例失败")
			}
		}
		node := message.Node{
			Id:       instance.StreamerId,
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
	_ = service.InstanceService.PakControl(deploy.Node, "load", deploy.Station)
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
		_ = service.InstanceService.PakControl(node, "unload", "")
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

func (s *sgccService) SendStatus(node string) {
	statistic := message.Statistic{
		Gpu:    int(math.Round(rand.Float64()*4 + 20)),
		Cpu:    int(math.Round(rand.Float64()*10 + 5)),
		Memory: rand.Float32()*1.5 + 1.5,
	}
	msg := message.Status{
		Type:      "status",
		Node:      node,
		Status:    s.GetNodeStatus(node),
		Statistic: statistic,
	}
	provider.SendCloudMessage(msg.GetBytes())
}

func (s *sgccService) Restart(nodes []string) {
	// instance restart
	var callback *message.RestartCallback
	var callbackNodes []*message.CallBackNode
	for _, node := range nodes {
		service.InstanceService.ProcessControl(request.ProcessControl{
			StreamerId: node,
			Command:    "STOP",
		})
		service.InstanceService.ProcessControl(request.ProcessControl{
			StreamerId: node,
			Command:    "START",
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
			StreamerId: node,
			Command:    "STOP",
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
