package server

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"thingue-launcher/agent/service/instance"
	"thingue-launcher/common/config"
	"thingue-launcher/common/message"
	"thingue-launcher/common/util"
)

type serverConnManager struct {
	ws                          *websocket.Conn
	RemoteServerConnCloseChanel chan string
}

var ServerConnManager = serverConnManager{
	ws:                          nil,
	RemoteServerConnCloseChanel: make(chan string),
}

func (m *serverConnManager) Connect(httpUrl string) error {
	var err error
	wsUrl := util.HttpUrlToWsUrl(httpUrl, "/ws/agent")
	appConfig := config.AppConfig
	m.ws, _, err = websocket.DefaultDialer.Dial(wsUrl, nil)
	if err == nil {
		fmt.Printf("服务连接成功：%s\n", wsUrl)
		//2.1如果连接成功,保存连接信息
		appConfig.ServerUrl = httpUrl
		config.WriteConfig()
		instance.NodeService.SetBaseUrl(httpUrl)
		//2.2如果连接成功,启动`消息接收goroutine`
		go func() {
			for {
				var msg = message.Message{}
				readErr := m.ws.ReadJSON(&msg)
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
			m.RemoteServerConnCloseChanel <- wsUrl
			appConfig.ServerUrl = ""
			config.WriteConfig()
		}()
	} else {
		// 如果连接失败
		fmt.Printf("服务连接失败：%s\n", wsUrl)
		m.Disconnect()
	}
	return err
}

func (m *serverConnManager) Disconnect() error {
	if m.ws != nil {
		err := m.ws.Close()
		return err
	} else {
		return errors.New("未连接无需断开")
	}
}
