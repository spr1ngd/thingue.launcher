package ws

import (
	"github.com/gorilla/websocket"
	"thingue-launcher/common/util"
)

type adminWsManager struct {
	Connections []*websocket.Conn
}

var AdminWsManager = new(adminWsManager)

func (m *adminWsManager) Broadcast() {
	for _, connection := range m.Connections {
		str := util.MapDataToJson(map[string]interface{}{
			"type": "update",
		})
		connection.WriteMessage(websocket.TextMessage, []byte(str))
	}
}
