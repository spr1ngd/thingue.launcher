package sdp

import (
	"github.com/gorilla/websocket"
	"thingue-launcher/common/util"
)

type PlayerConnector struct {
	PlayerId          uint
	StreamerConnector *StreamerConnector
	conn              *websocket.Conn
}

func (p *PlayerConnector) SendConfig() {
	p.SendMsg(util.MapDataToJsonStr(map[string]interface{}{
		"type":                  "config",
		"peerConnectionOptions": map[string]interface{}{},
	}))
}

func (p *PlayerConnector) HandleMsg(msgStr string) {
	msg := util.JsonStrToMapData(msgStr)
	msgType := msg["type"].(string)
	if msgType == "ping" {
		p.SendMsg(util.MapDataToJsonStr(map[string]interface{}{
			"type": "pong",
			"time": msg["time"],
		}))
	} else if msgType == "offer" {
		streamer := p.StreamerConnector
		streamer.AddPlayer(p)
		msg["playerId"] = p.PlayerId
		streamer.SendMsg(util.MapDataToJsonStr(msg))
	} else if msgType == "iceCandidate" {
		msg["playerId"] = p.PlayerId
		p.StreamerConnector.SendMsg(util.MapDataToJsonStr(msg))
	} else if msgType == "stats" {
	} else if msgType == "kick" {
		p.KickOthers()
	} else {
		p.SendCloseMsg(1008, "Unsupported message type")
	}
}

func (p *PlayerConnector) SendMsg(msg string) {
	p.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (p *PlayerConnector) Close() {
	p.conn.Close()
}

func (p *PlayerConnector) SendCloseMsg(code int, msg string) {
	p.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, msg))
}

func (p *PlayerConnector) SendPlayersCount() {
	players := p.StreamerConnector.PlayerConnectors
	msg := map[string]interface{}{
		"type":  "playerCount",
		"count": len(players),
	}
	for _, player := range players {
		player.SendMsg(util.MapDataToJsonStr(msg))
	}
}

func (p *PlayerConnector) KickOthers() {
	for _, kickPlayer := range p.StreamerConnector.PlayerConnectors {
		kickPlayer.SendCloseMsg(4000, "kicked")
	}
}
