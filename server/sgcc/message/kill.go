package message

import (
	"encoding/json"
)

type Kill struct {
	Type     string   `json:"type"`
	Nodes    []string `json:"nodes"`
	Datetime string   `json:"datetime"`
}

type KillCallback struct {
	Type  string          `json:"type"`
	Code  string          `json:"code"`
	Nodes []*CallBackNode `json:"nodes"`
}

func NewKillCallback(success bool, nodes []*CallBackNode) *KillCallback {
	callback := &KillCallback{
		Type:  "kill_callback",
		Nodes: nodes,
	}
	if success {
		callback.Code = "0"
	} else {
		callback.Code = "-1"
	}
	return callback
}

func (m *KillCallback) GetBytes() []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}
