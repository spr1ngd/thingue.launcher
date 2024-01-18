package provider

import (
	"github.com/gorilla/websocket"
	"thingue-launcher/common/logger"
)

var Conn *websocket.Conn

func SendCloudMessage(data []byte) {
	err := Conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		logger.Zap.Errorf("消息发送失败%s", string(data))
	}
}
