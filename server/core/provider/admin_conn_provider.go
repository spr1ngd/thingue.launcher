package provider

import (
	"github.com/gorilla/websocket"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/util"
)

type adminConnProvider struct {
	ConnMap     map[int]*websocket.Conn
	idGenerator *util.IDGenerator
}

var AdminConnProvider = adminConnProvider{
	ConnMap:     make(map[int]*websocket.Conn),
	idGenerator: util.NewIDGenerator(),
}

func (p *adminConnProvider) AddConn(conn *websocket.Conn) int {
	id := p.idGenerator.GenerateID()
	p.ConnMap[id] = conn
	return id
}

func (p *adminConnProvider) RemoveConn(id int) {
	delete(p.ConnMap, id)
	p.idGenerator.InvalidateID(id)
}

func (p *adminConnProvider) NotifyDownloadComplete(id int) {
	msg := util.MapToJson(map[string]interface{}{
		"type": "downloadComplete",
	})
	conn := p.ConnMap[id]
	if conn != nil {
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func (p *adminConnProvider) BroadcastUpdate() {
	for _, conn := range p.ConnMap {
		str := util.MapToJson(map[string]interface{}{
			"type": "update",
		})
		conn.WriteMessage(websocket.TextMessage, str)
	}
}

func (p *adminConnProvider) CloseAllConnection() {
	for _, conn := range p.ConnMap {
		err := conn.Close()
		if err != nil {
			logger.Zap.Error(err)
		}
	}
}
