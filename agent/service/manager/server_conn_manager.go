package manager

import (
	"errors"
	"fmt"
	"golang.org/x/net/websocket"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/server"
	"thingue-launcher/common/app"
	"thingue-launcher/common/util"
)

type ServerConnManager struct {
	ws *websocket.Conn
}

func (m *ServerConnManager) Connect(httpUrl string) error {
	var err error
	wsUrl := util.HttpUrlToAgentWsUrl(httpUrl)
	fmt.Printf("正在连接服务%s===========================================================\n", wsUrl)
	appConfig := app.GetAppConfig()
	m.ws, err = websocket.Dial(wsUrl, "", "http://localhost/")
	//1.如果连接成功尝试注册
	if err == nil {
		fmt.Printf("服务连接成功：%s\n", wsUrl)
		err = server.RegisterAgent(httpUrl)
	} else {
		fmt.Printf("服务连接失败：%s\n", wsUrl)
	}
	if err == nil {
		fmt.Printf("服务注册成功：%s\n", wsUrl)
		//2.1如果注册成功,保存连接信息
		appConfig.ServerUrl = httpUrl
		app.WriteConfig()
		//2.2如果注册成功,启动`消息接收goroutine`
		go func() {
			fmt.Println("开启接收消息")
			for {
				//接收消息
				response := make([]byte, 512)
				n, readErr := m.ws.Read(response)
				if readErr != nil {
					//连接断开
					fmt.Println("接收响应失败：", err)
					break
				}
				fmt.Printf("收到响应：%s\n", response[:n])
				//MsgReceive(response[:n])
			}
			fmt.Printf("服务连接断开：%s\n", wsUrl)
			global.RemoteServerConnCloseChanel <- wsUrl
			appConfig.ServerUrl = ""
			app.WriteConfig()
		}()
	} else {
		// 如果注册失败
		fmt.Printf("服务注册失败：%s\n", wsUrl)
		//runtime.EventsEmit(s.ctx, constants.REMOTE_SERVER_CONNECTION_CLOSE)
		m.Disconnect()
	}
	return err
}

func (m *ServerConnManager) Disconnect() error {
	if m.ws != nil {
		err := m.ws.Close()
		return err
	} else {
		return errors.New("未连接无需断开")
	}
}
