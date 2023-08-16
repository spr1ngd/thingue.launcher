package server

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"thingue-launcher/common/config"
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
	//1.如果连接成功尝试注册
	if err == nil {
		fmt.Printf("服务连接成功：%s\n", wsUrl)
		err = RegisterAgent(httpUrl)
	} else {
		fmt.Printf("服务连接失败：%s\n", wsUrl)
		return err
	}
	if err == nil {
		fmt.Printf("服务注册成功：%s\n", wsUrl)
		//2.1如果注册成功,保存连接信息
		appConfig.ServerUrl = httpUrl
		config.WriteConfig()
		//2.2如果注册成功,启动`消息接收goroutine`
		go func() {
			for {
				var data = map[string]interface{}{}
				readErr := m.ws.ReadJSON(&data)
				//接收消息
				if readErr != nil {
					//连接断开
					fmt.Println("无法读取消息：", err)
					break
				}
				fmt.Printf("收到响应：%v\n", data)
				//MsgReceive(response[:n])
			}
			fmt.Printf("服务连接断开：%s\n", wsUrl)
			m.RemoteServerConnCloseChanel <- wsUrl
			appConfig.ServerUrl = ""
			config.WriteConfig()
		}()
	} else {
		// 如果注册失败
		fmt.Printf("服务注册失败：%s\n", wsUrl)
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
