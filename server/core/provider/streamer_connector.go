package provider

import (
	"github.com/gorilla/websocket"
	"gopkg.in/yaml.v3"
	"thingue-launcher/common/domain"
	"thingue-launcher/common/message"
	"thingue-launcher/common/provider"
	"thingue-launcher/common/util"
	"time"
)

type StreamerConnector struct {
	SID                 string
	PlayerConnectors    []*PlayerConnector
	conn                *websocket.Conn
	heartbeatTimer      *time.Timer
	AutoStopTimer       *time.Timer
	EnableRelay         bool
	EnableRenderControl bool
	RenderingState      bool
}

func (s *StreamerConnector) SendPong(msg map[string]any) {
	s.SendMessage(util.MapToJson(map[string]interface{}{
		"type": "pong",
		"time": msg["time"],
	}))
}

func (s *StreamerConnector) State(msg map[string]any) {
	name := msg["name"].(string)
	value := msg["value"].(bool)
	if "streamingState" == name && !value {
		for _, connector := range s.PlayerConnectors {
			connector.SendMessage(util.MapToJson(msg))
		}
	}
}

func (s *StreamerConnector) SendConfig() {
	if s.EnableRelay && provider.AppConfig.PeerConnectionOptions != "" {
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

func (s *StreamerConnector) ForwardMessage(msg map[string]any) {
	playerId, err := util.GetPlayerIdFromMessage(msg)
	if err != nil {
		s.SendCloseMsg(1008, "不支持的消息类型")
	}
	delete(msg, "playerId")
	for _, player := range s.PlayerConnectors {
		if player.PlayerId == playerId {
			player.SendMessage(util.MapToJson(msg))
			return
		}
	}
}

func (s *StreamerConnector) PlayerDisconnect(disconnectPlayer *PlayerConnector) {
	for i, player := range s.PlayerConnectors {
		if player == disconnectPlayer {
			s.PlayerConnectors = append(s.PlayerConnectors[:i], s.PlayerConnectors[i+1:]...)
			s.SendMessage(util.MapToJson(map[string]any{
				"type":     "playerDisconnected",
				"playerId": player.PlayerId,
			}))
			return
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

func (s *StreamerConnector) SendMessage(msgStr []byte) {
	s.conn.WriteMessage(websocket.TextMessage, msgStr)
}

func (s *StreamerConnector) SendCloseMsg(code int, msg string) {
	s.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, msg))
}

func (s *StreamerConnector) SendCommand(command *message.Command) {
	s.conn.WriteMessage(websocket.TextMessage, command.GetBytes())
}

func (s *StreamerConnector) Close() {
	_ = s.conn.Close()
	delete(SdpConnProvider.idStreamerMap, s.SID)
}

func (s *StreamerConnector) ControlRendering(rendering bool) {
	if s.EnableRenderControl {
		command := message.Command{}
		command.BuildRenderingCommand(&message.RenderingParams{Value: rendering})
		s.SendCommand(&command)
	}
}

func (s *StreamerConnector) UpdateRenderingState(msg map[string]any) {
	s.RenderingState = msg["value"].(bool)
}
