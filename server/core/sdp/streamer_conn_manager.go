package sdp

import (
	"fmt"
	"github.com/gorilla/websocket"
	"thingue-launcher/common/request"
	"thingue-launcher/server/core/provider"
	"thingue-launcher/server/core/service"
	"time"
)

type streamerConnManager struct {
	idStreamerMap    map[string]*StreamerConnector
	heartbeatTimeout chan string
}

var StreamerConnManager = streamerConnManager{
	idStreamerMap:    make(map[string]*StreamerConnector),
	heartbeatTimeout: make(chan string),
}

func (m *streamerConnManager) NewStreamerConnector(sid string, conn *websocket.Conn) *StreamerConnector {
	connector := &StreamerConnector{
		SID:              sid,
		conn:             conn,
		PlayerConnectors: make([]*PlayerConnector, 0),
		autoStopTimer:    time.NewTimer(999 * time.Second),
	}
	connector.autoStopTimer.Stop()
	m.idStreamerMap[sid] = connector
	provider.StreamerConnProvider.AddConn(connector.SID, conn)
	service.InstanceService.UpdateStreamerConnected(sid, true)
	go func() {
		for {
			<-connector.autoStopTimer.C
			if len(connector.PlayerConnectors) == 0 {
				service.InstanceService.ProcessControl(request.ProcessControl{
					SID:     sid,
					Command: "STOP",
				})
				fmt.Println("检查完毕，自动停止控制指令发送")
			} else {
				fmt.Println("检查完毕，不需要自动停止")
			}
		}
	}()
	return connector
}

func (m *streamerConnManager) GetConnectorById(sid string) *StreamerConnector {
	return m.idStreamerMap[sid]
}

func (m *streamerConnManager) OnStreamerDisconnect(connector *StreamerConnector) {
	for _, playerConnector := range connector.PlayerConnectors {
		playerConnector.Close()
	}
	delete(m.idStreamerMap, connector.SID)
	provider.StreamerConnProvider.RemoveConn(connector.SID)
	service.InstanceService.UpdateStreamerConnected(connector.SID, false)
}
