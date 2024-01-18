package message

import (
	"encoding/json"
)

type Init struct {
	Type     string `json:"type"`
	MaxNode  int    `json:"maxNode"`
	Datetime string `json:"datetime"`
}

type Node struct {
	Id       string `json:"id"`
	Status   int    `json:"status"`
	Datetime string `json:"datetime"`
	Station  string `json:"station"`
	LoadType int    `json:"loadType"`
}

type InitCallback struct {
	Type  string `json:"type"`
	Code  string `json:"code"`
	Nodes []Node `json:"nodes"`
}

func NewInitCallback(success bool, nodes []Node) *InitCallback {
	callback := &InitCallback{
		Type:  "init_callback",
		Code:  "0",
		Nodes: nodes,
	}
	if success {
		callback.Code = "0"
	} else {
		callback.Code = "-1"
	}
	return callback
}

func (m *InitCallback) GetBytes() []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}
