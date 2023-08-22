package ws

import (
	"github.com/gorilla/websocket"
	"thingue-launcher/common/util"
)

type nodeWsManager struct {
	ConnMap map[uint]*websocket.Conn
}

var NodeWsManager = nodeWsManager{ConnMap: make(map[uint]*websocket.Conn)}

func (m *nodeWsManager) SendToNode(nodeId uint, msg map[string]any) {
	conn := m.ConnMap[nodeId]
	if conn != nil {
		conn.WriteMessage(websocket.TextMessage, []byte(util.MapDataToJsonStr(msg)))
	}
}

func (m *nodeWsManager) send() {

}
