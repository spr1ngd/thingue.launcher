package server

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"thingue-launcher/common/model/message"
)

func MsgReceive(msg map[string]interface{}) error {
	var err error
	switch msg["type"].(string) {
	case "ConnectCallback":
		id := msg["data"].(float64)
		RegisterNode(uint(id))
	case "control":
		var controlMsg message.ControlMsg
		err = mapstructure.Decode(msg["data"], &controlMsg)
		if err == nil {
			Control(controlMsg)
		}
	case "update":
		var updateMsg message.UpdateMsg
		err = mapstructure.Decode(msg["data"], &updateMsg)
		if err == nil {
			Update(updateMsg)
		}
	default:
		return errors.New("不支持的消息类型")
	}
	return nil
}

func Control(msg message.ControlMsg) {
	//todo
}

func Update(msg message.UpdateMsg) {
	//todo
}
