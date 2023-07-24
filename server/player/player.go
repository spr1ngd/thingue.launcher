package player

import (
	"github.com/gorilla/websocket"
	"thingue-launcher/server/streamer"
)

var playerIdCount = 0

type Player struct {
	PlayerId int
	Streamer *streamer.Streamer
	conn     *websocket.Conn
}

func NewPlayer(conn *websocket.Conn) *Player {
	playerIdCount++
	return &Player{
		PlayerId: playerIdCount,
		conn:     conn,
	}
}

func (p *Player) setStreamer(s *streamer.Streamer) {
	p.Streamer = s
}

func (p *Player) sendMsg(msg string) {
	p.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (p *Player) close() {
	p.conn.Close()
}

func (p *Player) kickOthers() {
	for _, kickPlayer := range p.Streamer.Players {
		kickPlayer.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(4000, "kicked"))
	}
}
