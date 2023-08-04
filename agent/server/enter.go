package server

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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

func (s *Server) ServerStart() {
	runtime.EventsEmit(s.ctx, "local_server_status_update", true)
	server.Startup()
	runtime.EventsEmit(s.ctx, "local_server_status_update", false)
}

func (s *Server) ServerShutdown() {
	server.Shutdown()
}

func (s *Server) GetServerStatus() bool {
	return server.GetLocalServerStatus()
}

func (s *Server) UpdateLocalServerConfig(localServerConfig config.LocalServer) {
	appConfig := config.GetAppConfig()
	appConfig.LocalServer = localServerConfig
	config.WriteConfig()
}
