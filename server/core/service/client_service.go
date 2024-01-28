package service

import (
	"bytes"
	"errors"
	"github.com/bluele/gcache"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"math"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/model"
	"thingue-launcher/common/request"
	"thingue-launcher/server/core/provider"
	"thingue-launcher/server/global"
	"time"
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
	global.ServerDB.Create(&client)
}

func (s *clientService) RegisterClient(client *model.Client, agentInfo *pb.GetAgentInfoResponse) {
	_ = mapstructure.Decode(agentInfo.DeviceInfo, client)
	var serverInstances = make([]*model.Instance, 0)
	for _, instance := range agentInfo.Instances {
		var serverInstance = &model.Instance{}
		_ = mapstructure.Decode(instance, serverInstance)
		_ = mapstructure.Decode(instance.Config, serverInstance)
		serverInstance.ID = instance.Id
		if serverInstance.StreamerId == "" {
			streamerId, _ := uuid.NewUUID()
			serverInstance.StreamerId = streamerId.String()
		} else {
			serverInstance.StreamerId = instance.StreamerId
		}
		serverInstances = append(serverInstances, serverInstance)
	}
	client.Instances = serverInstances
	global.ServerDB.Save(&client)
	provider.AdminConnProvider.BroadcastUpdate()
}

func (s *clientService) DeleteClient(client *model.Client) {
	global.ServerDB.Delete(&client)
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

func (s *clientService) CollectLogs(req request.LogsCollect) error {
	//s.WsIdMap[req.TraceId] = req.WsId
	//return provider.ClientConnProvider.SendToClient(req.ClientId, &message.Message{
	//	Type: types.ServerCollectClientLogs,
	//	Data: req.TraceId,
	//})
	// todo
	return nil
}

func (s *clientService) UploadLogs(traceId string, buf *bytes.Buffer) error {
	err := s.BufferCache.SetWithExpire(traceId, buf, time.Second*60)
	if err != nil {
		logger.Zap.Error(err)
		return err
	}
	provider.AdminConnProvider.NotifyDownloadComplete(s.WsIdMap[traceId])
	delete(s.WsIdMap, traceId)
	return nil
}

func (s *clientService) DownloadLogs(traceId string) (error, *bytes.Buffer) {
	value, err := s.BufferCache.Get(traceId)
	if err != nil {
		return errors.New("缓冲区数据超时被清除"), nil
	} else {
		return err, value.(*bytes.Buffer)
	}
}
