package message

import (
	"encoding/json"
)

type Release struct {
	Type     string   `json:"type"`
	Nodes    []string `json:"nodes"`
	Datetime string   `json:"datetime"`
}

type CallBackNode struct {
	Id       string `json:"id"`
	Status   int    `json:"status"`
	Datetime string `json:"datetime"`
}

type ReleaseCallback struct {
	Type  string          `json:"type"`
	Code  string          `json:"code"`
	Nodes []*CallBackNode `json:"nodes"`
}

func NewReleaseCallback(success bool, nodes []*CallBackNode) *ReleaseCallback {
	callback := &ReleaseCallback{
		Type:  "release_callback",
		Nodes: nodes,
	}
	if success {
		callback.Code = "0"
	} else {
		callback.Code = "-1"
	}
	return callback
}

func (m *ReleaseCallback) GetBytes() []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}
