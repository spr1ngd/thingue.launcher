package service

import (
	"context"
	"github.com/bluele/gcache"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/protobuf/types/known/emptypb"
	"math"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/common/model"
	"thingue-launcher/server/core/provider"
	"thingue-launcher/server/global"
)

type clientService struct {
	WsIdMap     map[string]int
	BufferCache gcache.Cache
}

var ClientService = clientService{
	WsIdMap:     make(map[string]int),
	BufferCache: gcache.New(math.MaxInt64).LRU().Build(),
}

func (s *clientService) ClientList() []*model.Client {
	var clients []*model.Client
	global.ServerDB.Preload("Instances").Find(&clients)
	return clients
}

func (s *clientService) CreateClient(client *model.Client) {
	global.ServerDB.Create(client)
}

func (s *clientService) RegisterClient(client *model.Client, agentInfo *pb.GetAgentInfoResponse) {
	_ = mapstructure.Decode(agentInfo.DeviceInfo, client)
	var serverInstances []*model.Instance
	for _, instance := range agentInfo.Instances {
		var serverInstance = &model.Instance{}
		_ = mapstructure.Decode(instance, serverInstance)
		_ = mapstructure.Decode(instance.Config, serverInstance)
		_ = mapstructure.Decode(instance.PlayerConfig, serverInstance)
		serverInstance.ID = instance.Id
		serverInstance.StreamerId = instance.Config.Name
		serverInstances = append(serverInstances, serverInstance)
	}
	client.Instances = serverInstances
	global.ServerDB.Save(client)
	provider.AdminConnProvider.BroadcastUpdate()
}

func (s *clientService) DeleteClient(client *model.Client) {
	global.ServerDB.Delete(client)
	global.ServerDB.Where("client_id = ?", client.ID).Delete(&model.Instance{})
	provider.AdminConnProvider.BroadcastUpdate()
}

func (s *clientService) GetInstanceStreamerId(clientId, instanceId uint32) (string, error) {
	var instance model.Instance
	err := global.ServerDB.Where("client_id = ? AND id = ?", clientId, instanceId).First(&instance).Error
	if err == nil {
		return instance.StreamerId, err
	} else {
		return "", err
	}
}

func (s *clientService) CollectLogs(clientId uint32) ([]byte, error) {
	client, err := provider.GrpcClientProvider.GetClient(clientId)
	if err != nil {
		return nil, err
	}
	resp, err := client.GetInstanceLogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
