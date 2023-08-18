package instance

import (
	"fmt"
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
	node.Instances = registerInfo.Instances
	global.SERVER_DB.Save(&node)
	return err
}

func (s *nodeService) NodeList() []model.Node {
	var nodeList []model.Node
	global.SERVER_DB.Preload("Instances").Find(&nodeList)
	return nodeList
}
