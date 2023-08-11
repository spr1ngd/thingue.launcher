package manager

import (
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

func (m *ServerConnManager) ConnectServer(httpUrl string) error {
	wsUrl := util.HttpUrlToAgentWsUrl(httpUrl)
	fmt.Printf("正在连接服务%s===========================================================\n", wsUrl)
	appConfig := app.GetAppConfig()
	ws, err := websocket.Dial(wsUrl, "", "http://localhost/")
	//1.如果连接成功尝试注册
	if err == nil {
		fmt.Printf("服务连接成功：%s\n", wsUrl)
		err = server.RegisterAgent(httpUrl)
	}
	if err == nil {
		//2.1如果注册成功,保存连接信息
		global.WS = ws
		appConfig.ServerUrl = httpUrl
		app.WriteConfig()
		//2.2如果注册成功,启动`消息接收goroutine`
		go func() {
			fmt.Println("开启接收消息")
			for {
				//接收消息
				response := make([]byte, 512)
				n, readErr := ws.Read(response)
				if readErr != nil {
					//连接断开
					fmt.Println("接收响应失败：", err)
					break
				}
				fmt.Printf("收到响应：%s\n", response[:n])
				//MsgReceive(response[:n])
			}
			fmt.Println("服务连接断开")
			global.WS = nil
			appConfig.ServerUrl = ""
			app.WriteConfig()
			//runtime.EventsEmit(s.ctx, constants.REMOTE_SERVER_CONNECTION_CLOSE)
		}()
	} else {
		fmt.Printf("服务连接失败：%s\n", wsUrl)
		//runtime.EventsEmit(s.ctx, constants.REMOTE_SERVER_CONNECTION_CLOSE)
		m.DisconnectServer()
	}
	return err
}

func (m *ServerConnManager) DisconnectServer() {
	if global.WS != nil {
		global.WS.Close()
	}
}
