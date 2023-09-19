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
	p.SendMessage(util.MapToJson(map[string]interface{}{
		"type":                  "config",
		"peerConnectionOptions": map[string]interface{}{},
	}))
}

func (p *PlayerConnector) HandleMessage(msgStr []byte) {
	msg := util.JsonStrToMapData(msgStr)
	msgType := msg["type"].(string)
	if msgType == "ping" {
		p.SendMessage(util.MapToJson(map[string]interface{}{
			"type": "pong",
			"time": msg["time"],
		}))
	} else if msgType == "listStreamers" {
		//todo
		var ids []string
		ids = append(ids, "DefaultStreamer")
		backMsg := map[string]any{}
		backMsg["type"] = "streamerList"
		backMsg["ids"] = ids
		p.SendMessage(util.MapToJson(backMsg))
	} else if msgType == "offer" { // for old streamer
		p.StreamerConnector.AddPlayer(p)
		msg["playerId"] = sanitizePlayerId(p.PlayerId)
		p.StreamerConnector.SendMessage(util.MapToJson(msg))
		p.StreamerConnector.SendPlayersCount()
	} else if msgType == "subscribe" { // for new streamer
		p.StreamerConnector.AddPlayer(p)
		forwardMsg := map[string]any{}
		forwardMsg["type"] = "playerConnected"
		forwardMsg["playerId"] = sanitizePlayerId(p.PlayerId)
		forwardMsg["dataChannel"] = true
		forwardMsg["sfu"] = false
		forwardMsg["sendOffer"] = true
		p.StreamerConnector.SendMessage(util.MapToJson(forwardMsg))
		p.StreamerConnector.SendPlayersCount()
	} else if msgType == "answer" { // for new streamer
		msg["playerId"] = sanitizePlayerId(p.PlayerId)
		p.StreamerConnector.SendMessage(util.MapToJson(msg))
	} else if msgType == "iceCandidate" {
		msg["playerId"] = sanitizePlayerId(p.PlayerId)
		p.StreamerConnector.SendMessage(util.MapToJson(msg))
	} else if msgType == "stats" {
	} else if msgType == "kick" {
		p.KickOthers()
		p.StreamerConnector.SendPlayersCount()
	} else {
		p.SendCloseMsg(1008, "Unsupported message type")
	}
}

func (p *PlayerConnector) SendMessage(msg []byte) {
	p.conn.WriteMessage(websocket.TextMessage, msg)
}

func (p *PlayerConnector) Close() {
	p.conn.Close()
}

func (p *PlayerConnector) SendCloseMsg(code int, msg string) {
	p.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, msg))
}

func (p *PlayerConnector) KickOthers() {
	for _, kickPlayer := range p.StreamerConnector.PlayerConnectors {
		kickPlayer.SendCloseMsg(4000, "kicked")
	}
}
