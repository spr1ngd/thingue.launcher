package provider

import (
	"errors"
	"github.com/gorilla/websocket"
	"thingue-launcher/common/message"
)

type nodeConnProvider struct {
	ConnMap map[uint]*websocket.Conn
}

var NodeConnProvider = nodeConnProvider{ConnMap: make(map[uint]*websocket.Conn)}

func (m *nodeConnProvider) SendToNode(nodeId uint, message *message.Message) error {
	conn := m.ConnMap[nodeId]
	if conn != nil {
		return conn.WriteMessage(websocket.TextMessage, message.GetBytes())
	} else {
		return errors.New("节点不存在")
	}
}

func (m *nodeConnProvider) send() {

}
