package ws

import (
	"github.com/gorilla/websocket"
	"thingue-launcher/common/message"
)

type nodeWsManager struct {
	ConnMap map[uint]*websocket.Conn
}

var NodeWsManager = nodeWsManager{ConnMap: make(map[uint]*websocket.Conn)}

func (m *nodeWsManager) SendToNode(nodeId uint, message *message.Message) {
	conn := m.ConnMap[nodeId]
	if conn != nil {
		conn.WriteMessage(websocket.TextMessage, message.GetBytes())
	}
}

func (m *nodeWsManager) send() {

}
