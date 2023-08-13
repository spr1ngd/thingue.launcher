package core

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"thingue-launcher/common/app"
	"time"
)

type Server struct {
	server          http.Server
	serverIsRunning bool
	ExitChan        chan error
}

var ServerApp = new(Server)

func (s *Server) Start(staticFiles embed.FS) {
	s.server = http.Server{
		Addr:    app.GetAppConfig().LocalServer.BindAddr,
		Handler: BuildRouter(staticFiles),
	}
	s.serverIsRunning = true
	go func() {
		err := s.server.ListenAndServe() //运行中阻塞
		s.serverIsRunning = false
		if err != nil {
			fmt.Printf("Server closed: %v\n", err)
		}
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
