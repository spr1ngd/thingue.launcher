package server

import (
	"crypto/tls"
	"errors"
	"net/url"
	"sync"
	"thingue-launcher/client/service/instance"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/message"
	"thingue-launcher/common/provider"
	"thingue-launcher/common/util"
	"time"

	"github.com/gorilla/websocket"
)

type connManager struct {
	conn                   *websocket.Conn
	ServerAddr             string
	connectLock            sync.Mutex
	IsConnected            bool
	reconnectTimer         *time.Timer
	reconnectInterval      int
	maxReconnectInterval   int
	ServerConnUpdateChanel chan string
	heartbeatTicker        *time.Ticker
}

var ConnManager = &connManager{
	conn:                   nil,
	maxReconnectInterval:   60,
	reconnectInterval:      2,
	ServerConnUpdateChanel: make(chan string, 1),
}

func (m *connManager) Init() {
	if provider.AppConfig.ServerURL != "" {
		_ = m.SetServerAddr(provider.AppConfig.ServerURL)
		m.StartConnectTask()
	}
}

func (m *connManager) connect() error {
	m.connectLock.Lock()
	defer m.connectLock.Unlock()
	if m.IsConnected {
		return nil
	}
	wsUrl := util.HttpUrlToWsUrl(m.ServerAddr, "/ws/client")
	dialer := websocket.DefaultDialer
	dialer.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, _, err := dialer.Dial(wsUrl, nil)
	if err == nil {
		m.IsConnected = true
		m.conn = conn
		instance.BaseRequest.SetBaseUrl(provider.AppConfig.ServerURL)
		m.ServerConnUpdateChanel <- wsUrl
		m.StartHeartbeatTask()
		go func() {
			for {
				var msg = message.Message{}
				readErr := m.conn.ReadJSON(&msg)
				if readErr != nil {
					break
				}
				//todo 处理异常
				_ = MsgReceive(msg)
			}
			_ = m.conn.Close()
			m.IsConnected = false
			instance.BaseRequest.UnsetBaseUrl()
			m.ServerConnUpdateChanel <- wsUrl
			if m.ServerAddr != "" {
				m.StartConnectTask()
			}
		}()
	}
	return err
}

func (m *connManager) Reconnect() {
	if m.conn != nil {
		_ = m.conn.Close()
	}
}

func (m *connManager) Disconnect() {
	m.ServerAddr = ""
	if m.conn != nil {
		_ = m.conn.Close()
	}
	provider.AppConfig.ServerURL = m.ServerAddr
	provider.WriteConfigToFile()
}

func (m *connManager) SetServerAddr(serverAddr string) error {
	if m.IsConnected {
		return errors.New("连接未断开")
	} else {
		m.ServerAddr = serverAddr
		provider.AppConfig.ServerURL = m.ServerAddr
		provider.WriteConfigToFile()
		return nil
	}
}

func (m *connManager) GetConnectedUrl() (*url.URL, error) {
	if m.IsConnected {
		parse, err := url.Parse(m.ServerAddr)
		return parse, err
	} else {
		return nil, errors.New("服务未连接")
	}
}

func (m *connManager) StartConnectTask() {
	m.reconnectTimer = time.NewTimer(time.Duration(m.reconnectInterval) * time.Second)
	go func() {
		for {
			<-m.reconnectTimer.C
			logger.Zap.Debug("连接开始")
			if m.ServerAddr == "" {
				break //不需要连接了
			}
			err := m.connect()
			if err == nil {
				break
			} else {
				m.reconnectTimer.Reset(time.Duration(m.reconnectInterval) * time.Second)
				logger.Zap.Debugf("连接失败,%d秒后重试\n", m.reconnectInterval)
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
			err := m.conn.WriteMessage(websocket.TextMessage, util.MapToJson(map[string]any{"type": "ping", "time": util.DateFormat(t)}))
			if err != nil {
				m.heartbeatTicker.Stop()
				err = m.conn.Close()
				logger.Zap.Error(err)
				break
			}
		}
		logger.Zap.Info("停止发送心跳")
	}()
}
