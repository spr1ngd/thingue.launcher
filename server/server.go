package server

import (
	"context"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"thingue-launcher/common/config"
	"thingue-launcher/server/initialize"
	"thingue-launcher/server/web/router"
	"time"
)

type Server struct {
	server            http.Server
	IsRunning         bool
	CloseReturnChanel chan string
	staticFiles       embed.FS
	router            *gin.Engine
	isInitialized     bool
}

//go:embed all:frontend/dist
var staticFiles embed.FS

var App = Server{staticFiles: staticFiles}

func (s *Server) Serve() {
	if !s.isInitialized { //如果是第一次没有初始化
		s.router = router.BuildRouter(s.staticFiles) //构建路由
		initialize.InitGorm()                        // 初始化gorm
		s.isInitialized = true
	}
	s.server = http.Server{
		Addr:    config.AppConfig.LocalServer.BindAddr,
		Handler: s.router,
	}
	s.IsRunning = true
	fmt.Println("thingue server listening at: ", s.server.Addr)
	err := s.server.ListenAndServe() //运行中阻塞
	s.IsRunning = false
	if s.CloseReturnChanel != nil {
		s.CloseReturnChanel <- err.Error()
	}
	fmt.Printf("Server closed: %v\n", err)
}

func (s *Server) Start() {
	if s.IsRunning {
		return
	}
	go func() {
		s.Serve()
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
	}
}
