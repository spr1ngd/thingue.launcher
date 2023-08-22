package instance

import (
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"thingue-launcher/common/model"
	"thingue-launcher/common/request"
	"thingue-launcher/server/global"
)

type nodeService struct{}

var NodeService = new(nodeService)

func (s *nodeService) NodeRegister(registerInfo *request.NodeRegisterInfo) error {
	var node model.Node
	global.SERVER_DB.Find(&node, registerInfo.NodeID)
	node.SetDeviceInfo(*registerInfo.DeviceInfo)
	var serverInstances = make([]*model.ServerInstance, 0)
	for _, instance := range registerInfo.Instances {
		var s *model.ServerInstance
		mapstructure.Decode(instance, &s)
		sid, _ := uuid.NewUUID()
		s.SID = sid.String()
		serverInstances = append(serverInstances, s)
	}
	node.Instances = serverInstances
	global.SERVER_DB.Save(&node)
	return nil
}

func (s *nodeService) NodeList() []model.Node {
	var nodeList []model.Node
	global.SERVER_DB.Preload("Instances").Find(&nodeList)
	return nodeList
}

func (s *nodeService) GetInstanceSid(nodeId string, instanceId string) (string, error) {
	var instance model.ServerInstance
	err := global.SERVER_DB.Where("node_id = ? AND id = ?", nodeId, instanceId).First(&instance).Error
	if err == nil {
		return instance.SID, err
	} else {
		return "", err
	}
}
