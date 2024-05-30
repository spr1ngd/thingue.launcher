package message

import (
	"encoding/json"
)

type Deploy struct {
	Type     string `json:"type"`
	Stage    any    `json:"stage"`
	Station  string `json:"station"`
	AssetId  string `json:"assetId"`
	User     string `json:"user"`
	Node     string `json:"node"`
	Datetime string `json:"datetime"`
}

type DeployCallback struct {
	Type     string `json:"type"`
	Code     string `json:"code"`
	Node     string `json:"node"`
	User     string `json:"user"`
	Status   int    `json:"status"`
	Datetime string `json:"datetime"`
}

func NewDeployCallback(success bool) *DeployCallback {
	callback := &DeployCallback{
		Type: "deploy_callback",
	}
	if success {
		callback.Code = "0"
	} else {
		callback.Code = "-1"
	}
	return callback
}

func (m *DeployCallback) GetBytes() []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}
