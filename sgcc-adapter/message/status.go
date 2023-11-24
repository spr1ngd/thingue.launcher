package message

import "encoding/json"

type Status struct {
	Type      string    `json:"type"`
	Node      string    `json:"node"`
	AssetId   string    `json:"assetId"`
	Status    int       `json:"status"`
	Statistic Statistic `json:"statistic"`
}

type Statistic struct {
	Gpu    int `json:"gpu"`
	Cpu    int `json:"cpu"`
	Memory int `json:"memory"`
}

func (m *Status) GetBytes() []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}
