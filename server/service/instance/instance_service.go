package instance

import (
	"thingue-launcher/common/message"
	"thingue-launcher/common/model"
	"thingue-launcher/common/request"
	"thingue-launcher/server/global"
	"thingue-launcher/server/service/ws"
)

type instanceService struct{}

var InstanceService = new(instanceService)

func (s *instanceService) AddPlayer(sid string, playerId string) {
	var instance model.ServerInstance
	global.SERVER_DB.Where("s_id = ?", sid).First(&instance)
	instance.PlayerIds = append(instance.PlayerIds, playerId)
	instance.PlayerCount = instance.PlayerCount + 1
	global.SERVER_DB.Save(&instance)
	ws.AdminWsManager.Broadcast()
}

func (s *instanceService) RemovePlayer(sid string, playerId string) {
	var instance model.ServerInstance
	global.SERVER_DB.Where("s_id = ?", sid).First(&instance)
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

func (s *instanceService) ProcessControl(processControl request.ProcessControl) {
	var instance model.ServerInstance
	global.SERVER_DB.Where("s_id = ?", processControl.SID).First(&instance)
	control := message.ServerProcessControl{
		CID:     instance.CID,
		Command: processControl.Command,
	}
	ws.NodeWsManager.SendToNode(instance.NodeID, control.Pack())
}

func (s *instanceService) PakControl() {

}
