package server

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/net/websocket"
	"log"
	"strings"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/model"
	"thingue-launcher/common/config"
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
	appConfig := config.GetAppConfig()
	if appConfig.LocalServer.AutoStart {
		s.LocalServerStart()
	}
	if appConfig.ServerUrl != "" {
		fmt.Println("ServerUrl不是空")
		s.ConnectServer(appConfig.ServerUrl)
	} else {
		fmt.Println("ServerUrl是空")
	}

}

func (s *Server) LocalServerStart() {
	runtime.EventsEmit(s.ctx, "local_server_status_update", true)
	server.Startup()
	runtime.EventsEmit(s.ctx, "local_server_status_update", false)
}

func (s *Server) LocalServerShutdown() {
	server.Shutdown()
}

func (s *Server) GetLocalServerStatus() bool {
	return server.GetLocalServerStatus()
}

func (s *Server) UpdateLocalServerConfig(localServerConfig config.LocalServer) {
	appConfig := config.GetAppConfig()
	appConfig.LocalServer = localServerConfig
	config.WriteConfig()
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
		appConfig := config.GetAppConfig()
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

func (s *Server) ConnectServer(httpUrl string) {
	wsUrl := strings.Replace(httpUrl, "http://", "ws://", 1)
	wsUrl = strings.Replace(wsUrl, "https://", "wss://", 1)
	fmt.Printf("正在连接%s\n", wsUrl)
	appConfig := config.GetAppConfig()
	ws, err := websocket.Dial(wsUrl, "", "http://localhost/")
	if err != nil {
		fmt.Printf("连接失败：%s\n", err)
		runtime.EventsEmit(s.ctx, "ServerConnectionClose")
		global.WS = nil
		appConfig.ServerUrl = ""
		config.WriteConfig()
		return
	}
	fmt.Printf("连接成功：%s\n", wsUrl)
	global.WS = ws
	appConfig.ServerUrl = httpUrl
	config.WriteConfig()
	for {
		response := make([]byte, 512)
		n, err := ws.Read(response)
		if err != nil {
			log.Println("接收响应失败：", err)
			break
		}
		fmt.Printf("收到响应：%s\n", response[:n])
		MsgReceive(response[:n])
	}
	global.WS = nil
	appConfig.ServerUrl = ""
	config.WriteConfig()
	ws.Close()
	runtime.EventsEmit(s.ctx, "ServerConnectionClose")
}

func (s *Server) DisconnectServer() {
	if global.WS != nil {
		global.WS.Close()
	}
}
