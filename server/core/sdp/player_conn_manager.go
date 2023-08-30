package sdp

import (
	"errors"
	"github.com/gorilla/websocket"
	"strconv"
	"thingue-launcher/common/util"
	"thingue-launcher/server/core/service"
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
	sid, err := service.TicketService.GetSidByTicket(ticket)
	if err == nil {
		streamerConnector := StreamerConnManager.GetConnectorById(sid)
		if streamerConnector != nil {
			playerConnector.StreamerConnector = streamerConnector
			service.InstanceService.AddPlayer(sid, strconv.Itoa(int(playerConnector.PlayerId)))
		} else {
			err = errors.New("streamer已离线")
		}
	}
	return err
}

func (m *playerConnManager) OnPlayerDisConnect(playerConnector *PlayerConnector) {
	playerConnector.StreamerConnector.SendMessage(util.MapToJson(map[string]interface{}{
		"type":     "playerDisconnected",
		"playerId": playerConnector.PlayerId,
	}))
	service.InstanceService.RemovePlayer(
		playerConnector.StreamerConnector.SID,
		strconv.Itoa(int(playerConnector.PlayerId)),
	)
}
