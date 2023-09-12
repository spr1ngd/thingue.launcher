package server

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
	"thingue-launcher/agent/service/instance"
	"thingue-launcher/common/message"
	"thingue-launcher/common/provider"
	"thingue-launcher/common/util"
	"time"
)

type connManager struct {
	conn                   *websocket.Conn
	heartbeatTicker        *time.Ticker
	reconnectTimer         *time.Timer
	ServerConnUpdateChanel chan string
	CloseSignalChannel     chan string
	ActiveAddrUrl          string
	ReconnectInterval      int
	MaxReconnectInterval   int
}

var ConnManager = connManager{
	conn:                   nil,
	ServerConnUpdateChanel: make(chan string, 1),
	CloseSignalChannel:     make(chan string),
	MaxReconnectInterval:   60,
}
var connectLock sync.Mutex

func (m *connManager) Connect(httpUrl string) error {
	connectLock.Lock()
	defer connectLock.Unlock()
	if m.ActiveAddrUrl != "" {
		fmt.Println("无法重复连接")
		return nil
	}
	var err error
	wsUrl := util.HttpUrlToWsUrl(httpUrl, "/ws/agent")
	m.conn, _, err = websocket.DefaultDialer.Dial(wsUrl, nil)
	if err == nil {
		fmt.Printf("服务连接成功：%s\n", wsUrl)
		m.ServerConnUpdateChanel <- wsUrl
		//2.1如果连接成功,保存连接信息
		m.ActiveAddrUrl = httpUrl
		instance.NodeService.SetBaseUrl(httpUrl)
		//2.2如果连接成功,启动`消息接收goroutine`
		go func() {
			m.StartHeartbeat()
			for {
				var msg = message.Message{}
				readErr := m.conn.ReadJSON(&msg)
				//接收消息
				if readErr != nil {
					//连接断开
					fmt.Println("无法读取消息：", readErr, msg)
					break
				}
				fmt.Printf("收到响应：%v\n", msg)
				MsgReceive(msg)
			}
			fmt.Printf("服务连接断开：%s\n", wsUrl)
			_ = m.conn.Close()
			m.ActiveAddrUrl = ""
			// 关闭定时器
			m.heartbeatTicker.Stop()
			// 如果是异常断连，开启重连尝试
			m.ServerConnUpdateChanel <- wsUrl
			select {
			case m.CloseSignalChannel <- "exitCode":
				//正常退出1
			default:
				m.StartReconnect()
			}
		}()
	} else {
		// 如果连接失败
		fmt.Printf("服务连接失败：%s\n", wsUrl)
	}
	return err
}

func (m *connManager) Disconnect() error {
	if m.conn != nil {
		err := m.conn.Close()
		fmt.Printf("开始连接关闭%s\n", err)
		if err == nil {
			exitCode := <-m.CloseSignalChannel
			fmt.Printf("连接关闭%s\n", exitCode)
		}
		return err
	} else {
		return errors.New("未连接无需断开")
	}
}

func (m *connManager) StartReconnect() {
	fmt.Println("开始尝试重连")
	m.ReconnectInterval = 1
	m.reconnectTimer = time.NewTimer(time.Duration(m.ReconnectInterval) * time.Second)
	go func() {
		for {
			t := <-m.reconnectTimer.C
			fmt.Println("重连一次", t.Format("2006-01-02 15:04:05"))
			err := m.Connect(provider.AppConfig.RegisterUrl)
			if err == nil {
				fmt.Println("重连成功")
				break
			}
			nextInterval := util.IntMin(m.ReconnectInterval*2, m.MaxReconnectInterval)
			m.reconnectTimer.Reset(time.Duration(nextInterval) * time.Second)
		}
		fmt.Println("停止尝试重连")
	}()
}

func (m *connManager) StartHeartbeat() {
	// 创建一个定时器，每隔一段时间发送心跳消息
	m.heartbeatTicker = time.NewTicker(40 * time.Second)
	go func() {
		for {
			t := <-m.heartbeatTicker.C
			err := m.conn.WriteMessage(websocket.TextMessage, util.MapToJson(map[string]interface{}{"type": "ping", "time": t.Format("2006-01-02 15:04:05")}))
			if err != nil {
				m.heartbeatTicker.Stop()
				err = m.conn.Close()
				fmt.Println(err)
				break
			}
		}
		fmt.Println("停止发送心跳")
	}()
}
