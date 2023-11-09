package provider

import (
	"errors"
	"fmt"
	"thingue-launcher/common/message"

	"github.com/gorilla/websocket"
)

type clientConnProvider struct {
	ConnMap map[uint]*websocket.Conn
}

var ClientConnProvider = clientConnProvider{ConnMap: make(map[uint]*websocket.Conn)}

func (p *clientConnProvider) SendToClient(clientId uint, message *message.Message) error {
	conn := p.ConnMap[clientId]
	if conn != nil {
		return conn.WriteMessage(websocket.TextMessage, message.GetBytes())
	} else {
		return errors.New("客户端不存在")
	}
}

func (p *clientConnProvider) SendToAllClients(message *message.Message) {
	for _, conn := range p.ConnMap {
		err := conn.WriteMessage(websocket.TextMessage, message.GetBytes())
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (p *clientConnProvider) CloseAllConnection() {
	for _, conn := range p.ConnMap {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
}
