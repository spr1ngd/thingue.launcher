package provider

import (
	"github.com/gorilla/websocket"
	"math"
	"math/rand"
	"thingue-launcher/common/logger"
	"thingue-launcher/server/sgcc/message"
)

var Conn *websocket.Conn

func SendCloudMessage(data []byte) {
	err := Conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		logger.Zap.Errorf("消息发送失败%s", string(data))
	}
}

func SendStatus(node, assetId string, status int) {
	statistic := message.Statistic{
		Gpu:    int(math.Round(rand.Float64()*4 + 20)),
		Cpu:    int(math.Round(rand.Float64()*10 + 5)),
		Memory: rand.Float32()*1.5 + 1.5,
	}
	msg := message.Status{
		Type:      "status",
		Node:      node,
		Status:    status,
		AssetId:   assetId,
		Statistic: statistic,
	}
	SendCloudMessage(msg.GetBytes())
}
