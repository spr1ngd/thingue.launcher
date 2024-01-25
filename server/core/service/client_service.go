package service

import (
	"bytes"
	"errors"
	"github.com/bluele/gcache"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"math"
	"thingue-launcher/common/domain"
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

func (s *clientService) ClientRegister(registerInfo *request.ClientRegisterInfo) error {
	var client model.Client
	global.ServerDB.Find(&client, registerInfo.ClientID)
	client.SetDeviceInfo(registerInfo.DeviceInfo)
	var serverInstances = make([]*model.ServerInstance, 0)
	for _, instance := range registerInfo.Instances {
		var serverInstance = &model.ServerInstance{}
		_ = mapstructure.Decode(instance, serverInstance)
		if serverInstance.SID == "" {
			sid, _ := uuid.NewUUID()
			serverInstance.SID = sid.String()
		}
		serverInstances = append(serverInstances, serverInstance)
	}
	client.Instances = serverInstances
	global.ServerDB.Save(&client)
	provider.AdminConnProvider.BroadcastUpdate()
	return nil
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
	deviceInfo := &domain.DeviceInfo{}
	_ = mapstructure.Decode(agentInfo.DeviceInfo, deviceInfo)
	client.SetDeviceInfo(deviceInfo)
	var serverInstances = make([]*model.ServerInstance, 0)
	for _, instance := range agentInfo.Instances {
		var serverInstance = &model.ServerInstance{}
		_ = mapstructure.Decode(instance, serverInstance)
		_ = mapstructure.Decode(instance.InstanceConfig, serverInstance)
		serverInstance.CID = uint(instance.Id)
		serverInstance.SID = instance.StreamerId
		if serverInstance.SID == "" {
			sid, _ := uuid.NewUUID()
			serverInstance.SID = sid.String()
		}
		serverInstances = append(serverInstances, serverInstance)
	}
	client.Instances = serverInstances
	global.ServerDB.Save(&client)
	provider.AdminConnProvider.BroadcastUpdate()
}

func (s *clientService) DeleteClient(client *model.Client) {
	global.ServerDB.Delete(&client)
	global.ServerDB.Where("client_id = ?", client.ID).Delete(&model.ServerInstance{})
	provider.AdminConnProvider.BroadcastUpdate()
}

func (s *clientService) GetInstanceSid(clientId string, instanceId string) (string, error) {
	var instance model.ServerInstance
	err := global.ServerDB.Where("client_id = ? AND c_id = ?", clientId, instanceId).First(&instance).Error
	if err == nil {
		return instance.SID, err
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
