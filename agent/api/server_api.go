package api

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"strings"
	"thingue-launcher/agent/constants"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/model"
	"thingue-launcher/agent/service"
	"thingue-launcher/common/app"
	"thingue-launcher/server"
	"thingue-launcher/server/core"
)

type serverApi struct {
	ctx context.Context
}

var ServerApi = new(serverApi)

func (s *serverApi) Init(ctx context.Context) {
	s.ctx = ctx
	appConfig := app.GetAppConfig()
	if appConfig.LocalServer.AutoStart {
		s.LocalServerStart()
	}
	if appConfig.ServerUrl != "" {
		err := s.ConnectServer(appConfig.ServerUrl)
		if err != nil {
			appConfig.ServerUrl = ""
			app.WriteConfig()
		}
	}
	global.RemoteServerConnCloseChanel = make(chan string)
	go func() {
		for {
			wsUrl := <-global.RemoteServerConnCloseChanel
			runtime.EventsEmit(s.ctx, constants.REMOTE_SERVER_CONN_CLOSE, wsUrl)
		}
	}()
}

func (s *serverApi) LocalServerStart() {
	runtime.EventsEmit(s.ctx, constants.LOCAL_SERVER_STATUS_UPDATE, true)
	server.Startup()
	runtime.EventsEmit(s.ctx, constants.LOCAL_SERVER_STATUS_UPDATE, false)
}

func (s *serverApi) LocalServerShutdown() {
	server.Shutdown()
}

func (s *serverApi) GetLocalServerStatus() bool {
	return core.ServerApp.GetLocalServerStatus()
}

func (s *serverApi) UpdateLocalServerConfig(localServerConfig app.LocalServer) {
	appConfig := app.GetAppConfig()
	appConfig.LocalServer = localServerConfig
	app.WriteConfig()
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

func (s *serverApi) ConnectServer(httpUrl string) error {
	return service.ServerConnManager.Connect(httpUrl)
}

func (s *serverApi) DisconnectServer() {
	service.ServerConnManager.Disconnect()
}
