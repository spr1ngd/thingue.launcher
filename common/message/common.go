package message

import "encoding/json"

type Message struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

func (msg *Message) GetBytes() []byte {
	bytes, _ := json.Marshal(msg)
	return bytes
}
