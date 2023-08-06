package server

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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
