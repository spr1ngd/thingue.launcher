package sdp

import (
	"fmt"
	"github.com/gorilla/websocket"
	"thingue-launcher/common/model"
	"thingue-launcher/common/util"
)

type StreamerConnector struct {
	ID               string
	PlayerConnectors []*PlayerConnector
	conn             *websocket.Conn
	Instance         *model.ServerInstance
}

func (s *StreamerConnector) HandleMsg(msgStr string) {
	msg := util.JsonStrToMapData(msgStr)
	msgType := msg["type"].(string)
	if msgType == "ping" {
		s.SendMsg(util.MapDataToJsonStr(map[string]interface{}{
			"type": "pong",
			"time": msg["time"],
		}))
	} else if msgType == "answer" {
		playerId := uint(msg["playerId"].(float64))
		for _, player := range s.PlayerConnectors {
			if player.PlayerId == playerId {
				player.SendMsg(msgStr)
			}
		}
	} else if msgType == "iceCandidate" {
		playerId := uint(msg["playerId"].(float64))
		for _, player := range s.PlayerConnectors {
			if player.PlayerId == playerId {
				player.SendMsg(msgStr)
			}
		}
	} else if msgType == "disconnectPlayer" {
		fmt.Println(msg)
	} else if msgType == "state" {
		fmt.Println(msg)
	} else {
		s.SendCloseMsg(1008, "Unsupported message type")
	}
}

func (s *StreamerConnector) SendConfig() {
	s.SendMsg(util.MapDataToJsonStr(map[string]interface{}{
		"type":                  "config",
		"peerConnectionOptions": map[string]interface{}{},
	}))
}

func (s *StreamerConnector) SendMsg(msg string) {
	fmt.Println("\x1b[1m" + "\x1b[32m" + msg + "\x1b[0m")
	s.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (s *StreamerConnector) AddPlayer(p *PlayerConnector) {
	s.PlayerConnectors = append(s.PlayerConnectors, p)
}

func (s *StreamerConnector) DeletePlayer(p *PlayerConnector) {
	for indexToDelete, pItem := range s.PlayerConnectors {
		if p == pItem {
			s.PlayerConnectors = append(s.PlayerConnectors[:indexToDelete], s.PlayerConnectors[indexToDelete+1:]...)
		}
	}
}

func (s *StreamerConnector) SendCloseMsg(code int, msg string) {
	s.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, msg))
}

func (s *StreamerConnector) close() {
	s.conn.Close()
}
