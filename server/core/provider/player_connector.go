package provider

import (
	"github.com/gorilla/websocket"
	"gopkg.in/yaml.v3"
	"thingue-launcher/common/domain"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/provider"
	"thingue-launcher/common/util"
	"time"
)

type PlayerConnector struct {
	PlayerId          uint
	StreamerConnector *StreamerConnector
	conn              *websocket.Conn
	UserData          map[string]string
	heartbeatTicker   *time.Ticker
}

func (p *PlayerConnector) SendConfig() {
	if p.StreamerConnector.EnableRelay && provider.AppConfig.PeerConnectionOptions != "" {
		var options domain.PeerConnectionOptions
		err := yaml.Unmarshal([]byte(provider.AppConfig.PeerConnectionOptions), &options)
		if err == nil {
			p.SendMessage(util.MapToJson(map[string]interface{}{
				"type":                  "config",
				"peerConnectionOptions": options,
			}))
			return
		}
	}
	p.SendMessage(util.MapToJson(map[string]interface{}{
		"type":                  "config",
		"peerConnectionOptions": map[string]interface{}{},
	}))
}

func (p *PlayerConnector) ForwardMessage(msg map[string]any) {
	msg["playerId"] = util.SanitizePlayerId(p.PlayerId)
	p.StreamerConnector.SendMessage(util.MapToJson(msg))
}

func (p *PlayerConnector) Offer(msg map[string]any) {
	p.ForwardMessage(msg)
}

func (p *PlayerConnector) Subscribe() {
	forwardMsg := map[string]any{}
	forwardMsg["type"] = "playerConnected"
	forwardMsg["playerId"] = util.SanitizePlayerId(p.PlayerId)
	forwardMsg["dataChannel"] = true
	forwardMsg["sfu"] = false
	forwardMsg["sendOffer"] = true
	p.StreamerConnector.SendMessage(util.MapToJson(forwardMsg))
}

func (p *PlayerConnector) ListStreamers() {
	//todo
	var ids []string
	ids = append(ids, "DefaultStreamer")
	backMsg := map[string]any{}
	backMsg["type"] = "streamerList"
	backMsg["ids"] = ids
	p.SendMessage(util.MapToJson(backMsg))
}

func (p *PlayerConnector) SendPong(msg map[string]any) {
	p.SendMessage(util.MapToJson(map[string]any{
		"type": "pong",
		"time": msg["time"],
	}))
}

func (p *PlayerConnector) SendMessage(msg []byte) {
	err := p.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		p.Close()
	}
}

func (p *PlayerConnector) SendPingMsg() error {
	return p.conn.WriteMessage(websocket.TextMessage, util.MapToJson(map[string]any{
		"type": "ping",
	}))
}

func (p *PlayerConnector) SendCloseMsg(code int, msg string) {
	err := p.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, msg))
	if err != nil {
		p.Close()
	}
}

func (p *PlayerConnector) KickOthers() {
	for _, kickPlayer := range p.StreamerConnector.PlayerConnectors {
		if p != kickPlayer {
			kickPlayer.SendCloseMsg(4000, "kicked")
		}
	}
	p.StreamerConnector.SendPlayersCount()
}

func (p *PlayerConnector) Kick() {
	p.SendCloseMsg(4000, "kicked")
	p.StreamerConnector.SendPlayersCount()
}

func (p *PlayerConnector) Close() {
	_ = p.conn.Close()
	delete(SdpConnProvider.idPlayerMap, p.PlayerId)
}

func (p *PlayerConnector) StartPingSendTask() {
	p.heartbeatTicker = time.NewTicker(5 * time.Second)
	go func() {
		for {
			_ = <-p.heartbeatTicker.C
			if p.conn == nil {
				p.heartbeatTicker.Stop()
				break
			} else {
				err := p.SendPingMsg()
				if err != nil {
					logger.Zap.Error("无法向Player发送心跳")
					p.Close()
					p.heartbeatTicker.Stop()
					break
				}
			}
		}
		logger.Zap.Debug("停止向Player发送心跳")
	}()
}
