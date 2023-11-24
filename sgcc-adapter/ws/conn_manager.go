package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
	"thingue-launcher/common/util"
	"thingue-launcher/sgcc-adapter/provider"
	"thingue-launcher/sgcc-adapter/service"
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
	reconnectInterval:    2,
}

func (m *connManager) connect() error {
	m.connectLock.Lock()
	defer m.connectLock.Unlock()
	if m.IsConnected {
		return nil
	}
	wsURL := provider.Config.CloudServerURL
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
		// 连接时注册
		service.SGCC.Register()
	}
	return err
}

func (m *connManager) StartConnectTask() {
	m.reconnectTimer = time.NewTimer(time.Duration(m.reconnectInterval) * time.Second)
	go func() {
		for {
			t := <-m.reconnectTimer.C
			fmt.Println("连接开始", t.Format("2006-01-02 15:04:05"))
			err := m.connect()
			if err == nil {
				break
			} else {
				m.reconnectTimer.Reset(time.Duration(m.reconnectInterval) * time.Second)
				fmt.Printf("连接失败,%d秒后重试\n", m.reconnectInterval)
			}
		}
	}()
}

func (m *connManager) StartHeartbeatTask() {
	// 创建一个定时器，每隔一段时间发送心跳消息
	m.heartbeatTicker = time.NewTicker(40 * time.Second)
	go func() {
		for {
			if !m.IsConnected {
				m.heartbeatTicker.Stop()
				break
			}
			t := <-m.heartbeatTicker.C
			err := provider.Conn.WriteMessage(websocket.TextMessage,
				util.MapToJson(map[string]any{"type": "ping", "time": util.DateFormat(t)}))
			if err != nil {
				m.heartbeatTicker.Stop()
				err = provider.Conn.Close()
				fmt.Println(err)
				break
			}
		}
		fmt.Println("停止发送心跳")
	}()
}
