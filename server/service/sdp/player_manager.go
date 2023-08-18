package sdp

import (
	"github.com/gorilla/websocket"
	"thingue-launcher/common/util"
)

type playerManager struct {
	playerIdCount uint
	idPlayerMap   map[uint]*PlayerConnector
}

var PlayerManager = playerManager{
	idPlayerMap: make(map[uint]*PlayerConnector),
}

func (m *playerManager) NewPlayerConnector(conn *websocket.Conn) *PlayerConnector {
	m.playerIdCount++
	p := &PlayerConnector{
		PlayerId: m.playerIdCount,
		conn:     conn,
	}
	m.idPlayerMap[p.PlayerId] = p
	return p
}

func (m *playerManager) OnPlayerDisConnect(playerConnector *PlayerConnector) {
	playerConnector.StreamerConnector.SendMsg(util.MapDataToJsonStr(map[string]interface{}{
		"type":     "playerDisconnected",
		"playerId": playerConnector.PlayerId,
	}))
}
