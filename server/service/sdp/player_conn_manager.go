package sdp

import (
	"errors"
	"github.com/gorilla/websocket"
	"strconv"
	"thingue-launcher/common/util"
	"thingue-launcher/server/service/instance"
)

type playerConnManager struct {
	playerIdCount uint
	idPlayerMap   map[uint]*PlayerConnector
}

var PlayerConnManager = playerConnManager{
	idPlayerMap: make(map[uint]*PlayerConnector),
}

func (m *playerConnManager) NewConnector(conn *websocket.Conn) *PlayerConnector {
	m.playerIdCount++
	p := &PlayerConnector{
		PlayerId: m.playerIdCount,
		conn:     conn,
	}
	m.idPlayerMap[p.PlayerId] = p
	return p
}

func (m *playerConnManager) SetStreamer(playerConnector *PlayerConnector, ticket string) error {
	sid, err := instance.TicketService.GetSidByTicket(ticket)
	if err == nil {
		streamerConnector := StreamerConnManager.GetConnectorById(sid)
		if streamerConnector != nil {
			playerConnector.StreamerConnector = streamerConnector
			instance.InstanceService.AddPlayer(sid, strconv.Itoa(int(playerConnector.PlayerId)))
		} else {
			err = errors.New("streamer已离线")
		}
	}
	return err
}

func (m *playerConnManager) OnPlayerDisConnect(playerConnector *PlayerConnector) {
	playerConnector.StreamerConnector.SendMsg(util.MapDataToJsonStr(map[string]interface{}{
		"type":     "playerDisconnected",
		"playerId": playerConnector.PlayerId,
	}))
	instance.InstanceService.RemovePlayer(
		playerConnector.StreamerConnector.SID,
		strconv.Itoa(int(playerConnector.PlayerId)),
	)
}
