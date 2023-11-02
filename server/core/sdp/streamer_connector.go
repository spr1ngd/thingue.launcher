package sdp

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/yaml.v3"
	"thingue-launcher/common/domain"
	"thingue-launcher/common/provider"
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
	} else if msgType == "offer" { // for new streamer
		s.ForwardMessage(msg)
	} else if msgType == "answer" { // for old streamer
		s.ForwardMessage(msg)
	} else if msgType == "iceCandidate" {
		s.ForwardMessage(msg)
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
	if provider.AppConfig.PeerConnectionOptions != "" {
		var options domain.PeerConnectionOptions
		err := yaml.Unmarshal([]byte(provider.AppConfig.PeerConnectionOptions), &options)
		if err == nil {
			s.SendMessage(util.MapToJson(map[string]interface{}{
				"type":                  "config",
				"peerConnectionOptions": options,
			}))
			return
		}
	}
	s.SendMessage(util.MapToJson(map[string]interface{}{
		"type":                  "config",
		"peerConnectionOptions": map[string]interface{}{},
	}))
}

func (s *StreamerConnector) SendMsg(msg string) {
	fmt.Println("\x1b[1m" + "\x1b[32m" + msg + "\x1b[0m")
	s.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (s *StreamerConnector) SendMessage(msgStr []byte) error {
	err := s.conn.WriteMessage(websocket.TextMessage, msgStr)
	return err
}

func (s *StreamerConnector) ForwardMessage(msg map[string]any) {
	playerId, err := getPlayerIdFromMessage(msg)
	if err != nil {
		s.SendCloseMsg(1008, "不支持的消息类型")
	}
	delete(msg, "playerId")
	for _, player := range s.PlayerConnectors {
		if player.PlayerId == playerId {
			player.SendMessage(util.MapToJson(msg))
		}
	}
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

func (s *StreamerConnector) SendPlayersCount() {
	players := s.PlayerConnectors
	msg := map[string]interface{}{
		"type":  "playerCount",
		"count": len(players),
	}
	for _, player := range players {
		player.SendMessage(util.MapToJson(msg))
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
