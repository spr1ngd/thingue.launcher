package sdp

import (
	"github.com/gorilla/websocket"
)

type streamerManager struct {
	idStreamerMap map[string]*StreamerConnector
}

var StreamerManager = streamerManager{
	idStreamerMap: make(map[string]*StreamerConnector),
}

func (m *streamerManager) NewStreamerConnector(id string, conn *websocket.Conn) *StreamerConnector {
	s := &StreamerConnector{
		ID:               id,
		conn:             conn,
		PlayerConnectors: make([]*PlayerConnector, 0),
	}
	m.idStreamerMap[id] = s
	return s
}

func (m *streamerManager) GetStreamerConnectorById(id string) *StreamerConnector {
	return m.idStreamerMap[id]
}

func (m *streamerManager) DeleteStreamerConnector(connector *StreamerConnector) {
	delete(m.idStreamerMap, connector.ID)
}
func (m *streamerManager) OnStreamerDisconnect(s *StreamerConnector) {

}
