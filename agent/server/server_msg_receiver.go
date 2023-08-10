package server

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"thingue-launcher/common/model"
	"thingue-launcher/common/model/message"
)

func MsgReceive(msgData []byte) error {
	msg := model.MsgStruct{}
	err := json.Unmarshal(msgData, &msg)
	if err != nil {
		return err
	}
	switch msg.Type {
	case "control":
		var controlMsg message.ControlMsg
		err = mapstructure.Decode(msg.Data, &controlMsg)
		if err == nil {
			Control(controlMsg)
		}
	case "update":
		var updateMsg message.UpdateMsg
		err = mapstructure.Decode(msg.Data, &updateMsg)
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
