package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"k8s.io/apimachinery/pkg/labels"
	"sync"
	"thingue-launcher/common/domain"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	types "thingue-launcher/common/gen/proto/go/types/v1"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/message"
	"thingue-launcher/common/model"
	"thingue-launcher/common/request"
	"thingue-launcher/server/core/provider"
	"thingue-launcher/server/global"
)

type instanceService struct {
	updateLock sync.Mutex
}

var InstanceService = new(instanceService)

func (s *instanceService) GetInstanceByStreamerId(streamerId string) *model.Instance {
	instance := &model.Instance{}
	global.ServerDB.Where("streamer_id = ?", streamerId).First(instance)
	return instance
}

func (s *instanceService) UpdatePlayers(streamer *provider.StreamerConnector) *model.Instance {
	s.updateLock.Lock()
	defer s.updateLock.Unlock()
	instance := s.GetInstanceByStreamerId(streamer.StreamerId)
	var playerIds []uint
	for _, connector := range streamer.PlayerConnectors {
		playerIds = append(playerIds, connector.PlayerId)
	}
	instance.PlayerIds = playerIds
	instance.PlayerCount = uint32(len(streamer.PlayerConnectors))
	global.ServerDB.Save(&instance)
	provider.AdminConnProvider.BroadcastUpdate()
	return instance
}

func (s *instanceService) UpdateStreamerConnected(streamerId string, connected bool) {
	s.updateLock.Lock()
	defer s.updateLock.Unlock()
	global.ServerDB.Model(&model.Instance{}).Where("streamer_id = ?", streamerId).Update("streamer_connected", connected)
	provider.AdminConnProvider.BroadcastUpdate()
	instance := model.Instance{}
	global.ServerDB.Where("streamer_id = ?", streamerId).First(&instance)
	getClient, err := provider.GrpcClientProvider.GetClient(instance.ClientID)
	if err == nil {
		_, err := getClient.UpdateStreamerState(context.Background(), &pb.UpdateStreamerStateRequest{
			InstanceId:    instance.ID,
			StreamerState: connected,
		})
		if err != nil {
			logger.Zap.Error(err)
		}
	}
}

func (s *instanceService) UpdateRenderingState(streamerId string, rendering bool) {
	global.ServerDB.Model(&model.Instance{}).Where("streamer_id = ?", streamerId).Update("rendering", rendering)
}

func (s *instanceService) UpdatePak(streamerId, pakValue string) {
	if pakValue != "" {
		instance := model.Instance{StreamerId: streamerId}
		global.ServerDB.Find(&instance)
		found := false
		for _, pak := range instance.Paks {
			if pak.Value == pakValue {
				found = true
				break
			}
		}
		if !found {
			logger.Zap.Debug("未配置的pak值", pakValue)
			return
		}
	}
	global.ServerDB.Model(&model.Instance{}).Where("streamer_id = ?", streamerId).Update("pak_value", pakValue)
	provider.AdminConnProvider.BroadcastUpdate()
}

func (s *instanceService) ProcessControl(processControl request.ProcessControl) {
	var instance model.Instance
	global.ServerDB.Where("streamer_id = ?", processControl.StreamerId).First(&instance)
	req := &pb.ControlProcessRequest{
		InstanceId: instance.ID,
	}
	if processControl.Command == "STOP" {
		req.Command = types.Command_COMMAND_STOP
		s.UpdatePak(instance.StreamerId, "")
	} else {
		req.Command = types.Command_COMMAND_START
	}
	getClient, err := provider.GrpcClientProvider.GetClient(instance.ClientID)
	if err == nil {
		_, err := getClient.ControlProcess(context.Background(), req)
		if err != nil {
			logger.Zap.Error(err)
		}
	}
}

func (s *instanceService) PakControl(control request.PakControl) error {
	command := message.Command{}
	if control.Type == "load" {
		instance := s.GetInstanceByStreamerId(control.StreamerId)
		if instance.PakValue != control.Pak {
			command.BuildBundleLoadCommand(message.BundleLoadParams{Bundle: control.Pak})
		} else {
			return errors.New("已经加载当前Pak")
		}
	} else if control.Type == "unload" {
		command.BuildBundleUnloadCommand()
	} else {
		return errors.New("不支持的消息类型")
	}
	streamer, err := provider.SdpConnProvider.GetStreamer(control.StreamerId)
	if err == nil {
		streamer.SendCommand(&command)
	}
	return err
}

func (s *instanceService) InstanceList() []*model.Instance {
	var instances []*model.Instance
	global.ServerDB.Find(&instances)
	return instances
}

func (s *instanceService) InstanceSelect(selectCond request.SelectorCond) ([]*model.Instance, error) {
	// 数据库查询
	//query := global.SERVER_DB.Where("state_code = ? or auto_control = ?", 1, true)
	query := global.ServerDB
	if selectCond.StreamerConnected == true {
		query = query.Where("streamer_connected = ?", selectCond.StreamerConnected)
	}
	if selectCond.StreamerId != "" {
		query = query.Where("streamer_id = ?", selectCond.StreamerId)
	}
	if selectCond.Name != "" {
		query = query.Where("name = ?", selectCond.Name)
	}
	if selectCond.PlayerCount != nil && *selectCond.PlayerCount >= 0 {
		query = query.Where("player_count = ?", selectCond.PlayerCount)
	}
	var findInstances []*model.Instance
	query.Find(&findInstances)
	// 筛选掉未启动且未开启自动启停的实例
	var readyInstances []*model.Instance
	for _, instance := range findInstances {
		if instance.StateCode == 1 || instance.AutoControl == true {
			readyInstances = append(readyInstances, instance)
		}
	}
	if len(readyInstances) > 0 && selectCond.LabelSelector != "" {
		// label匹配
		selector, err := labels.Parse(selectCond.LabelSelector)
		if err != nil {
			return nil, err // label解析失败
		}
		var matchInstances []*model.Instance
		for _, instance := range readyInstances {
			if selector.Matches(instance.Labels) {
				matchInstances = append(matchInstances, instance)
			}
		}
		return matchInstances, nil
	} else {
		return readyInstances, nil
	}
}

func (s *instanceService) GetInstanceByHostnameAndPid(hostname string, pid int) (*model.Instance, error) {
	db := global.ServerDB
	instance := &model.Instance{}
	tx := db.Debug().Select("server_instances.*").Joins("JOIN clients ON server_instances.client_id=clients.id AND clients.hostname = ? AND server_instances.pid = ?",
		hostname, pid).First(instance)
	if tx.Error == nil {
		return instance, nil
	} else {
		return nil, tx.Error
	}
}

func (s *instanceService) DeleteInstance(request *pb.DeleteInstanceRequest) error {
	tx := global.ServerDB.Where("client_id = ? and id = ?", request.ClientId, request.InstanceId).Delete(&model.Instance{})
	provider.AdminConnProvider.BroadcastUpdate()
	return tx.Error
}

func (s *instanceService) AddInstance(req *pb.AddInstanceRequest) error {
	var serverInstance model.Instance
	_ = mapstructure.Decode(req.InstanceInfo, &serverInstance)
	_ = mapstructure.Decode(req.InstanceInfo.Config, &serverInstance)
	serverInstance.ID = req.InstanceInfo.Id
	serverInstance.ClientID = req.ClientId
	serverInstance.StreamerId = req.InstanceInfo.StreamerId
	if serverInstance.StreamerId == "" {
		streamerId, _ := uuid.NewUUID()
		serverInstance.StreamerId = streamerId.String()
	} else {
		serverInstance.StreamerId = req.InstanceInfo.StreamerId
	}
	tx := global.ServerDB.Create(&serverInstance)
	provider.AdminConnProvider.BroadcastUpdate()
	return tx.Error
}

func (s *instanceService) UpdateInstanceConfig(req *pb.UpdateConfigRequest) error {
	instance := model.Instance{
		ID:       req.InstanceId,
		ClientID: req.ClientId,
	}
	_ = mapstructure.Decode(req.InstanceConfig, &instance)
	_ = mapstructure.Decode(req.PlayerConfig, &instance)
	tx := global.ServerDB.Omit("ID", "ClientID", "Pid", "StateCode", "StreamerConnected", "PlayerIds", "PlayerCount",
		"PakName", "PakValue", "Rendering", "LastStartAt", "LastStopAt").Updates(instance)
	provider.AdminConnProvider.BroadcastUpdate()
	return tx.Error
}

func (s *instanceService) UpdatePlayerConfig(req *pb.UpdateConfigRequest) error {
	playerConfig := domain.PlayerConfig{}
	_ = mapstructure.Decode(req.PlayerConfig, &playerConfig)
	tx := global.ServerDB.Model(&model.Instance{}).Where(model.Instance{
		ID:       req.InstanceId,
		ClientID: req.ClientId,
	}).Updates(playerConfig)
	provider.AdminConnProvider.BroadcastUpdate()
	return tx.Error
}

func (s *instanceService) UpdateProcessState(req *pb.UpdateProcessStateRequest) error {
	instance := model.Instance{
		ClientID:  req.ClientId,
		ID:        req.InstanceId,
		StateCode: req.StateCode,
		Pid:       req.Pid,
	}
	tx := global.ServerDB.Select("StateCode", "Pid").Updates(&instance)
	provider.AdminConnProvider.BroadcastUpdate()
	return tx.Error
}

func (s *instanceService) UpdateRestarting(req *pb.UpdateRestartingRequest) {
	provider.SdpConnProvider.SetStreamerRestartingState(req.StreamerId, req.Restarting)
}

func (s *instanceService) ClearPakState(req *pb.ClearPakStateRequest) error {
	instance := model.Instance{
		ClientID: req.ClientId,
		ID:       req.InstanceId,
		PakValue: "",
	}
	tx := global.ServerDB.Select("PakValue").Updates(&instance)
	provider.AdminConnProvider.BroadcastUpdate()
	return tx.Error
}
