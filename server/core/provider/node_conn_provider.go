package provider

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"thingue-launcher/common/message"
)

type nodeConnProvider struct {
	ConnMap map[uint]*websocket.Conn
}

var NodeConnProvider = nodeConnProvider{ConnMap: make(map[uint]*websocket.Conn)}

func (p *nodeConnProvider) SendToNode(nodeId uint, message *message.Message) error {
	conn := p.ConnMap[nodeId]
	if conn != nil {
		return conn.WriteMessage(websocket.TextMessage, message.GetBytes())
	} else {
		return errors.New("节点不存在")
	}
}

func (p *nodeConnProvider) SendToAllNode(message *message.Message) {
	for _, conn := range p.ConnMap {
		err := conn.WriteMessage(websocket.TextMessage, message.GetBytes())
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (p *nodeConnProvider) CloseAllConnection() {
	for _, conn := range p.ConnMap {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
}
