package message

import (
	"github.com/mitchellh/mapstructure"
	"thingue-launcher/common/message/types"
)

// ServerProcessControl 消息体1
type ServerProcessControl struct {
	CID     uint   `json:"cid"`
	Command string `json:"command"`
}

func (data *ServerProcessControl) Pack() *Message {
	msg := Message{}
	msg.Type = types.ServerProcessControl
	msg.Data = data
	return &msg
}

func (msg *Message) RecvServerProcessControl() *ServerProcessControl {
	data := ServerProcessControl{}
	mapstructure.Decode(msg.Data, &data)
	return &data
}

// ServerConnectCallback 消息体2
type ServerConnectCallback uint

func (data ServerConnectCallback) Pack() *Message {
	msg := Message{}
	msg.Type = types.ServerConnectCallback
	msg.Data = data
	return &msg
}

// NodeProcessStateUpdate 消息体3
type NodeProcessStateUpdate struct {
	SID       string `json:"sid"`
	StateCode int8   `json:"stateCode"`
}

// ServerStreamerConnectedUpdate 消息体4
type ServerStreamerConnectedUpdate struct {
	CID       uint `json:"cid"`
	Connected bool `json:"connected"`
}

func (data *ServerStreamerConnectedUpdate) Pack() *Message {
	msg := Message{}
	msg.Type = types.ServerStreamerConnectedUpdate
	msg.Data = data
	return &msg
}

func (msg *Message) RecvServerStreamerConnectedUpdate() *ServerStreamerConnectedUpdate {
	data := ServerStreamerConnectedUpdate{}
	mapstructure.Decode(msg.Data, &data)
	return &data
}

type SyncUpdate string

func (data SyncUpdate) Pack() *Message {
	msg := Message{}
	msg.Type = types.SyncUpdate
	msg.Data = data
	return &msg
}
