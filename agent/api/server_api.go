package api

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"strings"
	"thingue-launcher/agent/constants"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/service"
	"thingue-launcher/common/model"
	"thingue-launcher/common/provider"
	"thingue-launcher/server/initialize"
)

type serverApi struct {
	ctx context.Context
}

var ServerApi = new(serverApi)

func (s *serverApi) Init(ctx context.Context) {
	s.ctx = ctx
	if provider.AppConfig.LocalServer.AutoStart {
		s.LocalServerStart()
	}
	if provider.AppConfig.ServerUrl != "" {
		err := s.ConnectServer(provider.AppConfig.ServerUrl)
		if err != nil {
			provider.AppConfig.ServerUrl = ""
			provider.WriteConfig()
		}
	}
	go func() {
		for {
			wsUrl := <-service.ServerConnManager.RemoteServerConnCloseChanel
			runtime.EventsEmit(s.ctx, constants.REMOTE_SERVER_CONN_CLOSE, wsUrl)
		}
	}()
	// 监听localserver关闭
	initialize.Server.CloseReturnChanel = make(chan string)
	go func() {
		for {
			closeErr := <-initialize.Server.CloseReturnChanel
			runtime.EventsEmit(s.ctx, constants.LOCAL_SERVER_CLOSE, closeErr)
		}
	}()
}

func (s *serverApi) LocalServerStart() {
	initialize.Server.Start()
}

func (s *serverApi) LocalServerShutdown() {
	initialize.Server.Stop()
}

func (s *serverApi) GetLocalServerStatus() bool {
	return initialize.Server.IsRunning
}

func (s *serverApi) UpdateLocalServerConfig(localServerConfig provider.LocalServer) {
	appConfig := provider.AppConfig
	appConfig.LocalServer = localServerConfig
	provider.WriteConfig()
}

func (s *serverApi) ListRemoteServer() []model.RemoteServer {
	var list []model.RemoteServer
	global.APP_DB.Find(&list)
	return list
}

func (s *serverApi) CreateRemoteServer(remoteServer model.RemoteServer) uint {
	global.APP_DB.Create(&remoteServer)
	return remoteServer.ID
}

func (s *serverApi) SaveRemoteServer(remoteServer model.RemoteServer) uint {
	global.APP_DB.Save(&remoteServer)
	return remoteServer.ID
}

func (s *serverApi) DeleteRemoteServer(id uint) {
	global.APP_DB.Delete(&model.RemoteServer{}, id)
}

func (s *serverApi) GetConnectServerOptions() []string {
	var options []string
	if s.GetLocalServerStatus() {
		appConfig := provider.AppConfig
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

func (s *serverApi) ConnectServer(httpUrl string) error {
	return service.ServerConnManager.Connect(httpUrl)
}

func (s *serverApi) DisconnectServer() {
	service.ServerConnManager.Disconnect()
}
