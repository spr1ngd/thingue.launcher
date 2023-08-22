package server

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/service/instance"
	"thingue-launcher/common/message"
)

func MsgReceive(msg map[string]interface{}) error {
	var err error
	switch msg["type"].(string) {
	case "ConnectCallback":
		id := msg["data"].(float64)
		global.NODE_ID = uint(id)
		instance.NodeService.RegisterNode(global.NODE_ID)
	case "control":
		var msgData message.ControlMsg
		err = mapstructure.Decode(msg["data"], &msgData)
		if err == nil {
			instance.NodeService.Control()
		}
	case "update":
		var msgData message.UpdateMsg
		err = mapstructure.Decode(msg["data"], &msgData)
		if err == nil {
			instance.NodeService.Update()
		}
	default:
		return errors.New("不支持的消息类型")
	}
	return nil
}
