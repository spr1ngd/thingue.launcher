package sdp

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"thingue-launcher/common/request"
	"thingue-launcher/common/util"
	"thingue-launcher/server/core/service"
	"time"
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
	if ticket == "test" {
		streamerConnector := StreamerConnManager.GetConnectorById("test")
		if streamerConnector != nil {
			playerConnector.StreamerConnector = streamerConnector
		}
		return nil
	}
	sid, err := service.TicketService.GetSidByTicket(ticket)
	if err == nil {
		streamerConnector := StreamerConnManager.GetConnectorById(sid)
		if streamerConnector != nil {
			playerConnector.StreamerConnector = streamerConnector
			service.InstanceService.AddPlayer(sid, strconv.Itoa(int(playerConnector.PlayerId)))
		} else {
			instance := service.InstanceService.GetInstanceBySid(sid)
			if instance.AutoControl {
				service.InstanceService.ProcessControl(request.ProcessControl{
					SID:     sid,
					Command: "START",
				})
				ticker := time.NewTicker(2 * time.Second)
				for {
					<-ticker.C
					streamerConnector := StreamerConnManager.GetConnectorById(sid)
					if streamerConnector != nil {
						playerConnector.StreamerConnector = streamerConnector
						service.InstanceService.AddPlayer(sid, strconv.Itoa(int(playerConnector.PlayerId)))
						ticker.Stop()
						break
					}
				}
				fmt.Println("自动启动成功")
			} else {
				err = errors.New("streamer已离线")
			}
		}
	}
	return err
}

func (m *playerConnManager) OnPlayerDisConnect(playerConnector *PlayerConnector) {
	playerConnector.StreamerConnector.SendMessage(util.MapToJson(map[string]interface{}{
		"type":     "playerDisconnected",
		"playerId": playerConnector.PlayerId,
	}))
	playerConnector.StreamerConnector.SendPlayersCount()
	playerConnector.StreamerConnector.removePlayer(playerConnector)
	autoControl, delay := service.InstanceService.RemovePlayer(
		playerConnector.StreamerConnector.SID,
		strconv.Itoa(int(playerConnector.PlayerId)),
	)
	if autoControl && delay >= 0 {
		playerConnector.StreamerConnector.autoStopTimer.Reset(time.Duration(delay) * time.Second)
	}
}
