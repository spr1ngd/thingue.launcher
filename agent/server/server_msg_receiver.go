package server

import (
	"encoding/json"
	"fmt"
)

type Message[T ControlMsg | UpdateMsg | any] struct {
	Type string
	Data T
}

type ControlMsg struct {
	InstanceName string
	ControlType  string
}

type UpdateMsg struct {
}

func MsgReceive(msgData []byte) {
	msg := Message[ControlMsg]{}
	err := json.Unmarshal(msgData, &msg)
	if err != nil {
		fmt.Println("消息解析失败")
		return
	}
	switch msg.Type {
	case "Control":
		Control(msg.Data)
	default:
		return
	}

}

func Control(msg ControlMsg) {

}
