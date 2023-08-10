package server

import (
	"github.com/gorilla/websocket"
	"thingue-launcher/common/util"
)

var playerIdCount = 0
var idPlayerMap = map[int]*Player{}

type Player struct {
	PlayerId int
	Streamer *Streamer
	conn     *websocket.Conn
}

func NewPlayer(conn *websocket.Conn) *Player {
	playerIdCount++
	p := &Player{
		PlayerId: playerIdCount,
		conn:     conn,
	}
	idPlayerMap[p.PlayerId] = p
	return p
}

func (p *Player) Init() {
	p.SendMsg(util.MapDataToJsonStr(map[string]interface{}{
		"type":                  "config",
		"peerConnectionOptions": map[string]interface{}{},
	}))
}

func (p *Player) SetStreamer(s *Streamer) {
	p.Streamer = s
}

func (p *Player) SendMsg(msg string) {
	p.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (p *Player) Close() {
	p.conn.Close()
}

func (p *Player) SendCloseMsg(code int, msg string) {
	p.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, msg))
}

func (p *Player) SendPlayersCount() {
	players := p.Streamer.Players
	msg := map[string]interface{}{
		"type":  "playerCount",
		"count": len(players),
	}
	for _, player := range players {
		player.SendMsg(util.MapDataToJsonStr(msg))
	}
}

func (p *Player) KickOthers() {
	for _, kickPlayer := range p.Streamer.Players {
		kickPlayer.SendCloseMsg(4000, "kicked")
	}
}
