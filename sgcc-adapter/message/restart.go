package message

import (
	"encoding/json"
)

type Restart struct {
	Type     string   `json:"type"`
	Nodes    []string `json:"nodes"`
	Datetime string   `json:"datetime"`
}

type RestartCallback struct {
	Type  string          `json:"type"`
	Code  string          `json:"code"`
	Nodes []*CallBackNode `json:"nodes"`
}

func NewRestartCallback(success bool, nodes []*CallBackNode) *RestartCallback {
	callback := &RestartCallback{
		Type:  "restart_callback",
		Nodes: nodes,
	}
	if success {
		callback.Code = "0"
	} else {
		callback.Code = "-1"
	}
	return callback
}

func (m *RestartCallback) GetBytes() []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}
