package server

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/net/websocket"
	"strings"
	"thingue-launcher/agent/constants"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/model"
	"thingue-launcher/common/app"
	"thingue-launcher/common/util"
	"thingue-launcher/server"
)

type Server struct {
	ctx context.Context
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SetContext(ctx context.Context) {
	s.ctx = ctx
	appConfig := app.GetAppConfig()
	if appConfig.LocalServer.AutoStart {
		s.LocalServerStart()
	}
	if appConfig.ServerUrl != "" {
		fmt.Println("ServerUrl不是空")
		err := s.ConnectServer(appConfig.ServerUrl)
		if err != nil {

		}
	} else {
		fmt.Println("ServerUrl是空")
	}

}

func (s *Server) LocalServerStart() {
	runtime.EventsEmit(s.ctx, constants.LOCAL_SERVER_STATUS_UPDATE, true)
	server.Startup()
	runtime.EventsEmit(s.ctx, constants.LOCAL_SERVER_STATUS_UPDATE, false)
}

func (s *Server) LocalServerShutdown() {
	server.Shutdown()
}

func (s *Server) GetLocalServerStatus() bool {
	return server.GetLocalServerStatus()
}

func (s *Server) UpdateLocalServerConfig(localServerConfig app.LocalServer) {
	appConfig := app.GetAppConfig()
	appConfig.LocalServer = localServerConfig
	app.WriteConfig()
}

func (s *Server) ListRemoteServer() []model.RemoteServer {
	var list []model.RemoteServer
	global.APP_DB.Find(&list)
	return list
}

func (s *Server) CreateRemoteServer(remoteServer model.RemoteServer) uint {
	global.APP_DB.Create(&remoteServer)
	return remoteServer.ID
}

func (s *Server) SaveRemoteServer(remoteServer model.RemoteServer) uint {
	global.APP_DB.Save(&remoteServer)
	return remoteServer.ID
}

func (s *Server) DeleteRemoteServer(id uint) {
	global.APP_DB.Delete(&model.RemoteServer{}, id)
}

func (s *Server) GetConnectServerOptions() []string {
	var options []string
	if s.GetLocalServerStatus() {
		appConfig := app.GetAppConfig()
		port := strings.Split(appConfig.LocalServer.BindAddr, ":")[1]
		if strings.HasSuffix(port+appConfig.LocalServer.BasePath, "/") {
			options = append(options, "http://localhost:"+port+appConfig.LocalServer.BasePath)
		} else {
			options = append(options, "http://localhost:"+port+appConfig.LocalServer.BasePath+"/")
		}
	}
	for _, remoteServer := range s.ListRemoteServer() {
		options = append(options, remoteServer.Url)
	}
	return options
}

func (s *Server) ConnectServer(httpUrl string) error {
	wsUrl := util.HttpUrlToAgentWsUrl(httpUrl)
	fmt.Printf("正在连接服务%s===========================================================\n", wsUrl)
	appConfig := app.GetAppConfig()
	ws, err := websocket.Dial(wsUrl, "", "http://localhost/")
	//1.如果连接成功尝试注册
	if err == nil {
		fmt.Printf("服务连接成功：%s\n", wsUrl)
		err = RegisterAgent(httpUrl)
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
				MsgReceive(response[:n])
			}
			fmt.Println("服务连接断开")
			global.WS = nil
			appConfig.ServerUrl = ""
			app.WriteConfig()
			runtime.EventsEmit(s.ctx, constants.REMOTE_SERVER_CONNECTION_CLOSE)
		}()
	} else {
		fmt.Printf("服务连接失败：%s\n", wsUrl)
		runtime.EventsEmit(s.ctx, constants.REMOTE_SERVER_CONNECTION_CLOSE)
		s.DisconnectServer()
	}
	return err
}

func (s *Server) DisconnectServer() {
	if global.WS != nil {
		global.WS.Close()
	}
}
