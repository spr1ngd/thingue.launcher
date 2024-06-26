package service

import (
	"errors"
	"k8s.io/apimachinery/pkg/labels"
	"sync"
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

func (s *instanceService) GetInstanceBySid(sid string) *model.ServerInstance {
	instance := &model.ServerInstance{}
	global.SERVER_DB.Where("s_id = ?", sid).First(instance)
	return instance
}

func (s *instanceService) UpdatePlayers(streamer *provider.StreamerConnector) *model.ServerInstance {
	s.updateLock.Lock()
	defer s.updateLock.Unlock()
	instance := s.GetInstanceBySid(streamer.SID)
	var playerIds []uint
	for _, connector := range streamer.PlayerConnectors {
		playerIds = append(playerIds, connector.PlayerId)
	}
	instance.PlayerIds = playerIds
	instance.PlayerCount = uint(len(streamer.PlayerConnectors))
	global.SERVER_DB.Save(&instance)
	provider.AdminConnProvider.BroadcastUpdate()
	return instance
}

func (s *instanceService) UpdateStreamerConnected(sid string, connected bool) {
	s.updateLock.Lock()
	defer s.updateLock.Unlock()
	global.SERVER_DB.Model(&model.ServerInstance{}).Where("s_id = ?", sid).Update("streamer_connected", connected)
	provider.AdminConnProvider.BroadcastUpdate()
	instance := model.ServerInstance{}
	global.SERVER_DB.Where("s_id = ?", sid).First(&instance)
	updateMsg := message.ServerStreamerConnectedUpdate{
		CID:       instance.CID,
		Connected: connected,
	}
	provider.ClientConnProvider.SendToClient(instance.ClientID, updateMsg.Pack())
}

func (s *instanceService) UpdateProcessState(msg *message.ClientProcessStateUpdate) {
	global.SERVER_DB.Model(&model.ServerInstance{}).Where("s_id = ?", msg.SID).Updates(map[string]any{"Pid": msg.Pid, "StateCode": msg.StateCode})
	provider.AdminConnProvider.BroadcastUpdate()
}

func (s *instanceService) UpdateRenderingState(sid string, rendering bool) {
	global.SERVER_DB.Model(&model.ServerInstance{}).Where("s_id = ?", sid).Update("rendering", rendering)
}

func (s *instanceService) UpdatePak(sid, currentPakValue string) {
	if currentPakValue != "" {
		instance := model.ServerInstance{}
		global.SERVER_DB.Where("s_id = ?", sid).First(&instance)
		found := false
		for _, pak := range instance.Paks {
			if pak.Value == currentPakValue {
				found = true
				break
			}
		}
		if !found {
			logger.Zap.Debug("未配置的pak值", currentPakValue)
			return
		}
	}
	global.SERVER_DB.Model(&model.ServerInstance{}).Where("s_id = ?", sid).Update("current_pak", currentPakValue)
	provider.AdminConnProvider.BroadcastUpdate()
}

func (s *instanceService) ProcessControl(processControl request.ProcessControl) {
	var instance model.ServerInstance
	global.SERVER_DB.Where("s_id = ?", processControl.SID).First(&instance)
	control := message.ServerProcessControl{
		CID:     instance.CID,
		Command: processControl.Command,
	}
	provider.ClientConnProvider.SendToClient(instance.ClientID, control.Pack())
	if processControl.Command == "STOP" {
		s.UpdatePak(instance.SID, "")
	}

}

func (s *instanceService) PakControl(control request.PakControl) error {
	command := message.Command{}
	if control.Type == "load" {
		instance := s.GetInstanceBySid(control.SID)
		if instance.CurrentPak != control.Pak {
			command.BuildBundleLoadCommand(message.BundleLoadParams{Bundle: control.Pak})
		} else {
			return errors.New("已经加载当前Pak")
		}
	} else if control.Type == "unload" {
		command.BuildBundleUnloadCommand()
	} else {
		return errors.New("不支持的消息类型")
	}
	streamer, err := provider.SdpConnProvider.GetStreamer(control.SID)
	if err == nil {
		streamer.SendCommand(&command)
	}
	return err
}

func (s *instanceService) InstanceList() ([]*model.ServerInstance, error) {
	var instances []*model.ServerInstance
	global.SERVER_DB.Find(&instances)
	return instances, nil
}

func (s *instanceService) InstanceSelect(selectCond request.SelectorCond) ([]*model.ServerInstance, error) {
	// 数据库查询
	//query := global.SERVER_DB.Where("state_code = ? or auto_control = ?", 1, true)
	query := global.SERVER_DB
	if selectCond.StreamerConnected == true {
		query = query.Where("streamer_connected = ?", selectCond.StreamerConnected)
	}
	if selectCond.SID != "" {
		query = query.Where("s_id = ?", selectCond.SID)
	}
	if selectCond.Name != "" {
		query = query.Where("name = ?", selectCond.Name)
	}
	if selectCond.PlayerCount != nil && *selectCond.PlayerCount >= 0 {
		query = query.Where("player_count = ?", selectCond.PlayerCount)
	}
	var findInstances []*model.ServerInstance
	query.Find(&findInstances)
	// 筛选掉未启动且未开启自动启停的实例
	var readyInstances []*model.ServerInstance
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
		var matchInstances []*model.ServerInstance
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

func (s *instanceService) GetInstanceByHostnameAndPid(hostname string, pid int) (*model.ServerInstance, error) {
	db := global.SERVER_DB
	instance := &model.ServerInstance{}
	tx := db.Debug().Select("server_instances.*").Joins("JOIN clients ON server_instances.client_id=clients.id AND clients.hostname = ? AND server_instances.pid = ?",
		hostname, pid).First(instance)
	if tx.Error == nil {
		return instance, nil
	} else {
		return nil, tx.Error
	}
}
