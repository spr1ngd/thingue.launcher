package provider

import (
	"errors"
	"github.com/gorilla/websocket"
	"time"
)

type sdpConnProvider struct {
	playerIdCount uint
	idStreamerMap map[string]*StreamerConnector
	idPlayerMap   map[uint]*PlayerConnector
}

var SdpConnProvider = sdpConnProvider{
	idStreamerMap: make(map[string]*StreamerConnector),
	idPlayerMap:   make(map[uint]*PlayerConnector),
}

func (sdp *sdpConnProvider) NewStreamer(sid string, conn *websocket.Conn, enableRelay bool) *StreamerConnector {
	streamer := &StreamerConnector{
		SID:              sid,
		conn:             conn,
		PlayerConnectors: make([]*PlayerConnector, 0),
		AutoStopTimer:    time.NewTimer(999 * time.Second),
		EnableRelay:      enableRelay,
	}
	streamer.AutoStopTimer.Stop()
	sdp.idStreamerMap[streamer.SID] = streamer
	return streamer
}

func (sdp *sdpConnProvider) NewPlayer(conn *websocket.Conn) *PlayerConnector {
	sdp.playerIdCount++
	player := &PlayerConnector{
		PlayerId: sdp.playerIdCount,
		conn:     conn,
	}
	sdp.idPlayerMap[player.PlayerId] = player
	return player
}

func (sdp *sdpConnProvider) GetStreamer(id string) (*StreamerConnector, error) {
	connector, ok := sdp.idStreamerMap[id]
	if ok {
		return connector, nil
	} else {
		return nil, errors.New("streamer不存在")
	}
}

func (sdp *sdpConnProvider) GetPlayer(id uint) (*PlayerConnector, error) {
	connector, ok := sdp.idPlayerMap[id]
	if ok {
		return connector, nil
	} else {
		return nil, errors.New("player不存在")
	}
}
