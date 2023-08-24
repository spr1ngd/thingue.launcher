package provider

import (
	"errors"
	"github.com/gorilla/websocket"
	"thingue-launcher/common/message"
)

type streamerConnProvider struct {
	StreamerConnMap map[string]*websocket.Conn
}

var StreamerConnProvider = streamerConnProvider{StreamerConnMap: make(map[string]*websocket.Conn)}

func (p *streamerConnProvider) AddConn(sid string, conn *websocket.Conn) {
	p.StreamerConnMap[sid] = conn
}

func (p *streamerConnProvider) RemoveConn(sid string) {
	delete(p.StreamerConnMap, sid)
}

func (p *streamerConnProvider) SendCommand(sid string, command *message.Command) error {
	var err error
	conn := p.StreamerConnMap[sid]
	if conn != nil {
		err = conn.WriteMessage(websocket.TextMessage, command.GetBytes())
	} else {
		err = errors.New("streamer未连接")
	}
	return err
}
