package instance

import (
	"errors"
	"thingue-launcher/common/message"
	"thingue-launcher/common/model"
	"thingue-launcher/common/request"
	"thingue-launcher/server/global"
	"thingue-launcher/server/service/sdp/provider"
	"thingue-launcher/server/service/ws"
)

type instanceService struct{}

var InstanceService = new(instanceService)

func (s *instanceService) GetInstanceBySid(sid string) model.ServerInstance {
	var instance model.ServerInstance
	global.SERVER_DB.Where("s_id = ?", sid).First(&instance)
	return instance
}

func (s *instanceService) AddPlayer(sid string, playerId string) {
	instance := s.GetInstanceBySid(sid)
	instance.PlayerIds = append(instance.PlayerIds, playerId)
	instance.PlayerCount = instance.PlayerCount + 1
	global.SERVER_DB.Save(&instance)
	ws.AdminWsManager.Broadcast()
}

func (s *instanceService) RemovePlayer(sid string, playerId string) {
	instance := s.GetInstanceBySid(sid)
	for i, id := range instance.PlayerIds {
		if id == playerId {
			instance.PlayerIds = append(instance.PlayerIds[:i], instance.PlayerIds[i+1:]...)
		}
	}
	instance.PlayerCount = instance.PlayerCount - 1
	global.SERVER_DB.Save(&instance)
	ws.AdminWsManager.Broadcast()
}

func (s *instanceService) UpdateStreamerConnected(sid string, connected bool) {
	global.SERVER_DB.Model(&model.ServerInstance{}).Where("s_id = ?", sid).Update("streamer_connected", connected)
	ws.AdminWsManager.Broadcast()
	instance := model.ServerInstance{}
	global.SERVER_DB.Where("s_id = ?", sid).First(&instance)
	updateMsg := message.ServerStreamerConnectedUpdate{
		CID:       instance.CID,
		Connected: connected,
	}
	ws.NodeWsManager.SendToNode(instance.NodeID, updateMsg.Pack())
}

func (s *instanceService) UpdateProcessState(msg *message.NodeProcessStateUpdate) {
	global.SERVER_DB.Model(&model.ServerInstance{}).Where("s_id = ?", msg.SID).Update("state_code", msg.StateCode)
	ws.AdminWsManager.Broadcast()
}

func (s *instanceService) UpdateRendering(sid string, rendering bool) {
	global.SERVER_DB.Model(&model.ServerInstance{}).Where("s_id = ?", sid).Update("rendering", rendering)
}

func (s *instanceService) UpdatePak(sid string, pak string) {
	global.SERVER_DB.Model(&model.ServerInstance{}).Where("s_id = ?", sid).Update("current_pak", pak)
	ws.AdminWsManager.Broadcast()
}

func (s *instanceService) ProcessControl(processControl request.ProcessControl) {
	var instance model.ServerInstance
	global.SERVER_DB.Where("s_id = ?", processControl.SID).First(&instance)
	control := message.ServerProcessControl{
		CID:     instance.CID,
		Command: processControl.Command,
	}
	ws.NodeWsManager.SendToNode(instance.NodeID, control.Pack())
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
	return provider.StreamerConnProvider.SendCommand(control.SID, &command)
}
