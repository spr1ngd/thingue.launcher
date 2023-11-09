package request

import (
	"thingue-launcher/common/domain"
)

type SelectorCond struct {
	SID               string `json:"sid"`
	Name              string `json:"name"`
	PlayerCount       *int   `json:"playerCount"`
	LabelSelector     string `json:"labelSelector"`
	StreamerConnected bool   `json:"streamerConnected"`
}

type ClientRegisterInfo struct {
	ClientID   uint
	DeviceInfo *domain.DeviceInfo
	Instances  []*domain.Instance
}

type ProcessControl struct {
	SID     string `json:"sid"`
	Command string `json:"command"`
}

type PakControl struct {
	SID  string `json:"sid"`
	Type string `json:"type"`
	Pak  string `json:"pak"`
}

type LogsCollect struct {
	WsId     int    `json:"wsId"`
	TraceId  string `json:"traceId"`
	ClientId uint   `json:"clientId"`
}
