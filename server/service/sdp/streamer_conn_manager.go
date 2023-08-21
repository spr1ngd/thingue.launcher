package sdp

import (
	"github.com/gorilla/websocket"
	"thingue-launcher/server/service/instance"
)

type streamerConnManager struct {
	idStreamerMap map[string]*StreamerConnector
}

var StreamerConnManager = streamerConnManager{
	idStreamerMap: make(map[string]*StreamerConnector),
}

func (m *streamerConnManager) NewStreamerConnector(sid string, conn *websocket.Conn) *StreamerConnector {
	connector := &StreamerConnector{
		SID:              sid,
		conn:             conn,
		PlayerConnectors: make([]*PlayerConnector, 0),
	}
	m.idStreamerMap[sid] = connector
	instance.InstanceService.UpdateStreamerConnected(sid, true)
	return connector
}

func (m *streamerConnManager) GetConnectorById(sid string) *StreamerConnector {
	return m.idStreamerMap[sid]
}

func (m *streamerConnManager) DeleteConnector(connector *StreamerConnector) {
	// 1
}

func (m *streamerConnManager) OnStreamerDisconnect(connector *StreamerConnector) {
	for _, playerConnector := range connector.PlayerConnectors {
		playerConnector.Close()
	}
	delete(m.idStreamerMap, connector.SID)
	instance.InstanceService.UpdateStreamerConnected(connector.SID, false)
}
