package instance

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"thingue-launcher/common/model"
	"thingue-launcher/server/global"
)

type nodeService struct{}

var NodeService = new(nodeService)

func (s *nodeService) NodeRegister(registerInfo *model.NodeRegisterInfo) error {
	var node model.Node
	global.SERVER_DB.Find(&node, registerInfo.NodeID)
	deviceInfoMap := make(map[string]interface{})
	err := mapstructure.Decode(registerInfo.DeviceInfo, &deviceInfoMap)
	err = mapstructure.Decode(deviceInfoMap, &node.DeviceInfo)
	fmt.Printf("%+v\n", node)
	if err != nil {
		return err
	}
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
	return err
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
