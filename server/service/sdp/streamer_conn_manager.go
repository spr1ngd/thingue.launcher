package sdp

import (
	"github.com/gorilla/websocket"
)

type streamerConnManager struct {
	idStreamerMap map[string]*StreamerConnector
}

var StreamerConnManager = streamerConnManager{
	idStreamerMap: make(map[string]*StreamerConnector),
}

func (m *streamerConnManager) NewStreamerConnector(id string, conn *websocket.Conn) *StreamerConnector {
	s := &StreamerConnector{
		ID:               id,
		conn:             conn,
		PlayerConnectors: make([]*PlayerConnector, 0),
	}
	m.idStreamerMap[id] = s
	return s
}

func (m *streamerConnManager) GetStreamerConnectorById(id string) *StreamerConnector {
	return m.idStreamerMap[id]
}

func (m *streamerConnManager) DeleteStreamerConnector(connector *StreamerConnector) {
	delete(m.idStreamerMap, connector.ID)
}
func (m *streamerConnManager) OnStreamerDisconnect(s *StreamerConnector) {

}
