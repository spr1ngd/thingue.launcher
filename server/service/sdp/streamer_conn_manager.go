package sdp

import (
	"github.com/gorilla/websocket"
	"thingue-launcher/server/service/instance"
	"thingue-launcher/server/service/sdp/provider"
)

type streamerConnManager struct {
	idStreamerMap    map[string]*StreamerConnector
	heartbeatTimeout chan string
}

var StreamerConnManager = streamerConnManager{
	idStreamerMap:    make(map[string]*StreamerConnector),
	heartbeatTimeout: make(chan string),
}

func (m *streamerConnManager) NewStreamerConnector(sid string, conn *websocket.Conn) *StreamerConnector {
	connector := &StreamerConnector{
		SID:              sid,
		conn:             conn,
		PlayerConnectors: make([]*PlayerConnector, 0),
	}
	m.idStreamerMap[sid] = connector
	provider.StreamerConnProvider.AddConn(connector.SID, conn)
	instance.InstanceService.UpdateStreamerConnected(sid, true)
	return connector
}

func (m *streamerConnManager) GetConnectorById(sid string) *StreamerConnector {
	return m.idStreamerMap[sid]
}

func (m *streamerConnManager) OnStreamerDisconnect(connector *StreamerConnector) {
	for _, playerConnector := range connector.PlayerConnectors {
		playerConnector.Close()
	}
	delete(m.idStreamerMap, connector.SID)
	provider.StreamerConnProvider.RemoveConn(connector.SID)
	instance.InstanceService.UpdateStreamerConnected(connector.SID, false)
}
