package streamer

import (
	"github.com/gorilla/websocket"
	"thingue-launcher/server/player"
)

type Streamer struct {
	StreamerID string
	Players    []*player.Player
	conn       *websocket.Conn
}

func (s *Streamer) SendMsg(msg string) {
	s.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (s *Streamer) AddPlayer(p *player.Player) {
	s.Players = append(s.Players, p)
}

func (s *Streamer) DeletePlayer(p *player.Player) {
	for indexToDelete, pItem := range s.Players {
		if p == pItem {
			s.Players = append(s.Players[:indexToDelete], s.Players[indexToDelete+1:]...)
		}
	}
}

func (s *Streamer) close(msg string) {
	s.conn.Close()
}
