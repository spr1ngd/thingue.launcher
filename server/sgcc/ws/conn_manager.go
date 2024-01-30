package ws

import (
	"github.com/gorilla/websocket"
	"sync"
	"thingue-launcher/common/logger"
	"thingue-launcher/server/sgcc/config"
	"thingue-launcher/server/sgcc/provider"
	"thingue-launcher/server/sgcc/service"
	"time"
)

type connManager struct {
	connectLock          sync.Mutex
	IsConnected          bool
	reconnectTimer       *time.Timer
	reconnectInterval    int
	maxReconnectInterval int
	heartbeatTicker      *time.Ticker
}

var ConnManager = &connManager{
	maxReconnectInterval: 60,
	reconnectInterval:    5,
}

func (m *connManager) connect() error {
	m.connectLock.Lock()
	defer m.connectLock.Unlock()
	if m.IsConnected {
		return nil
	}
	wsURL := config.Config.CloudServerURL
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err == nil {
		m.IsConnected = true
		provider.Conn = conn
		go func() {
			for {
				msg := map[string]any{}
				readErr := conn.ReadJSON(&msg)
				if readErr != nil {
					break
				}
				//接收消息
				MsgReceive(msg)
			}
			_ = conn.Close()
			m.IsConnected = false
			m.StartConnectTask()
		}()
		// 注册
		service.SgccService.Register()
	}
	return err
}

func (m *connManager) StartConnectTask() {
	m.reconnectTimer = time.NewTimer(time.Duration(m.reconnectInterval) * time.Second)
	go func() {
		for {
			<-m.reconnectTimer.C
			logger.Zap.Debug("开始连接云端负载均衡")
			err := m.connect()
			if err == nil {
				break
			} else {
				m.reconnectTimer.Reset(time.Duration(m.reconnectInterval) * time.Second)
				logger.Zap.Debugf("连接云端负载均衡失败,%d秒后重试", m.reconnectInterval)
			}
		}
	}()
}
