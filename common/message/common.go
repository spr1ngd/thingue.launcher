package message

import "encoding/json"

// Message Server与Node通信消息格式
type Message struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

func (msg *Message) GetBytes() []byte {
	bytes, _ := json.Marshal(msg)
	return bytes
}

// Command Sever与UE通信消息格式
type Command struct {
	Type    string `json:"type"`
	Command string `json:"command"`
	Params  any    `json:"params"`
}

func (c *Command) GetBytes() []byte {
	bytes, _ := json.Marshal(c)
	return bytes
}
