package request

import (
	"thingue-launcher/common/domain"
	"thingue-launcher/common/model"
)

type TicketSelector struct {
	SID           string `json:"sid"`
	Name          string `json:"name"`
	PlayerCount   uint   `json:"playerCount"`
	LabelSelector string `json:"labelSelector"`
}

type NodeRegisterInfo struct {
	NodeID     uint
	DeviceInfo *domain.DeviceInfo
	Instances  []*model.ClientInstance
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
	WsId    int    `json:"wsId"`
	TraceId string `json:"traceId"`
	NodeId  uint   `json:"nodeId"`
}
