package sdp

import (
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"thingue-launcher/common/util"
	"thingue-launcher/server/core/service"
	"time"
)

type StreamerConnector struct {
	SID              string
	PlayerConnectors []*PlayerConnector
	conn             *websocket.Conn
	heartbeatTimer   *time.Timer
	autoStopTimer    *time.Timer
}

func (s *StreamerConnector) HandleMessage(msgStr []byte) {
	msg := util.JsonStrToMapData(msgStr)
	msgType := msg["type"].(string)
	if msgType == "ping" {
		s.SendMessage(util.MapToJson(map[string]interface{}{
			"type": "pong",
			"time": msg["time"],
		}))
	} else if msgType == "answer" {
		var playerId uint
		f, ok := msg["playerId"].(float64)
		if ok {
			playerId = uint(f)
		} else {
			parseUint, err := strconv.ParseUint(msg["playerId"].(string), 10, 32)
			if err != nil {
				s.SendCloseMsg(1008, "不支持的消息类型")
				return
			}
			playerId = uint(parseUint)
		}
		for _, player := range s.PlayerConnectors {
			if player.PlayerId == playerId {
				player.SendMessage(msgStr)
			}
		}
	} else if msgType == "iceCandidate" {
		var playerId uint
		f, ok := msg["playerId"].(float64)
		if ok {
			playerId = uint(f)
		} else {
			parseUint, err := strconv.ParseUint(msg["playerId"].(string), 10, 32)
			if err != nil {
				s.SendCloseMsg(1008, "不支持的消息类型")
				return
			}
			playerId = uint(parseUint)
		}
		for _, player := range s.PlayerConnectors {
			if player.PlayerId == playerId {
				player.SendMessage(msgStr)
			}
		}
	} else if msgType == "disconnectPlayer" {
		// todo
		fmt.Println(msg)
	} else if msgType == "state" {
		name := msg["name"].(string)
		value := msg["value"].(bool)
		if "streamingState" == name && !value {
			for _, connector := range s.PlayerConnectors {
				connector.SendMessage(msgStr)
			}
		}
		// todo
		fmt.Println(msg)
	} else if msgType == "rendering" {
		service.InstanceService.UpdateRendering(s.SID, msg["value"].(bool))
	} else if msgType == "hotReloadComplete" {
		service.InstanceService.UpdatePak(s.SID, "")
	} else if msgType == "loadComplete" {
		service.InstanceService.UpdatePak(s.SID, msg["bundleName"].(string))
	} else {
		s.SendCloseMsg(1008, "不支持的消息类型")
	}
}

func (s *StreamerConnector) SendConfig() {
	s.SendMessage(util.MapToJson(map[string]interface{}{
		"type":                  "config",
		"peerConnectionOptions": map[string]interface{}{},
	}))
}

func (s *StreamerConnector) SendMsg(msg string) {
	fmt.Println("\x1b[1m" + "\x1b[32m" + msg + "\x1b[0m")
	s.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (s *StreamerConnector) SendMessage(msg []byte) error {
	err := s.conn.WriteMessage(websocket.TextMessage, msg)
	return err
}

func (s *StreamerConnector) AddPlayer(p *PlayerConnector) {
	s.PlayerConnectors = append(s.PlayerConnectors, p)
}

func (s *StreamerConnector) removePlayer(p *PlayerConnector) {
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

func (s *StreamerConnector) StartTimer() {
	s.heartbeatTimer = time.NewTimer(10 * time.Second)
}

func (s *StreamerConnector) StopTimer() {}

func (s *StreamerConnector) ResetTimer() {}
