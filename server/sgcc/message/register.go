package message

import (
	"encoding/json"
	"thingue-launcher/sgcc-adapter/provider"
)

type RegisterMessage struct {
	Type       string   `json:"type"`
	ResType    int      `json:"resType"`
	Maintainer string   `json:"maintainer"`
	MaxNode    int      `json:"maxNode"`
	Url        string   `json:"url"`
	Stations   []string `json:"stations"`
}

type RegisterCallback struct {
	Type string `json:"type"`
	Code string `json:"code"`
	Id   string `json:"id"`
}

func NewRegister(maxNode int) *RegisterMessage {
	return &RegisterMessage{
		Type:       "register",
		ResType:    0,
		MaxNode:    maxNode,
		Maintainer: provider.Config.Register.Maintainer,
		Url:        provider.Config.Register.Url,
		Stations:   provider.Config.Register.Stations,
	}
}

func (m *RegisterMessage) GetBytes() []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}
