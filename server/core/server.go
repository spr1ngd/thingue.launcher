package core

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"thingue-launcher/common/config"
	"time"
)

type Server struct {
	server            http.Server
	serverIsRunning   bool
	CloseReturnChanel chan string
	staticFiles       embed.FS
}

var ServerApp = new(Server)

func (s *Server) Init(staticFiles embed.FS) {
	s.staticFiles = staticFiles
}

func (s *Server) Start() {
	s.server = http.Server{
		Addr:    config.AppConfig.LocalServer.BindAddr,
		Handler: BuildRouter(s.staticFiles),
	}
	s.serverIsRunning = true
	go func() {
		err := s.server.ListenAndServe() //运行中阻塞
		s.serverIsRunning = false
		if s.CloseReturnChanel != nil {
			s.CloseReturnChanel <- err.Error()
		}
		fmt.Printf("Server closed: %v\n", err)
	}()
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := s.server.Shutdown(ctx)
	if err != nil {
		fmt.Printf("Server shutdown failed: %v\n", err)
	} else {
		fmt.Println("Server gracefully stopped.")
		s.serverIsRunning = false
	}
}

func (s *Server) GetLocalServerStatus() bool {
	return s.serverIsRunning
}
