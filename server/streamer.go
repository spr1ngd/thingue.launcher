package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"thingue-launcher/server/util"
)

type Streamer struct {
	StreamerID string
	Players    []*Player
	conn       *websocket.Conn
}

var idStreamerMap = map[string]*Streamer{}

func NewStreamer(id string, conn *websocket.Conn) *Streamer {
	s := &Streamer{
		StreamerID: id,
		conn:       conn,
		Players:    make([]*Player, 0),
	}
	idStreamerMap[id] = s
	return s
}

func GetStreamerById(id string) *Streamer {
	return idStreamerMap[id]
}

func UnmapStreamerId(id string) {
	fmt.Printf("取消映射%s\n", id)
	delete(idStreamerMap, id)
}

func (s *Streamer) Init() {
	s.SendMsg(util.MapDataToJsonStr(map[string]interface{}{
		"type":                  "config",
		"peerConnectionOptions": map[string]interface{}{},
	}))
}

func (s *Streamer) SendMsg(msg string) {
	fmt.Println("\x1b[1m" + "\x1b[32m" + msg + "\x1b[0m")
	s.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (s *Streamer) AddPlayer(p *Player) {
	s.Players = append(s.Players, p)
}

func (s *Streamer) DeletePlayer(p *Player) {
	for indexToDelete, pItem := range s.Players {
		if p == pItem {
			s.Players = append(s.Players[:indexToDelete], s.Players[indexToDelete+1:]...)
		}
	}
}

func (s *Streamer) SendCloseMsg(code int, msg string) {
	s.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, msg))
}

func (s *Streamer) close(msg string) {
	s.conn.Close()
}
