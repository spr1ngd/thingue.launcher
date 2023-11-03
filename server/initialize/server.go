package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"thingue-launcher/common/model"
	"thingue-launcher/common/provider"
	coreprovider "thingue-launcher/server/core/provider"
	"thingue-launcher/server/global"
	"thingue-launcher/server/web/router"
)

type server struct {
	server            http.Server
	IsRunning         bool
	CloseReturnChanel chan string
	router            *gin.Engine
	isInitialized     bool
}

var Server = new(server)

func (s *server) Serve() {
	if !s.isInitialized { //如果是第一次没有初始化
		//s.router = router.BuildRouter(s.StaticFiles) //构建路由
		InitServerDB()  // 初始化gorm
		InitStorageDB() // 初始化gorm
		s.isInitialized = true
	}
	global.SERVER_DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Node{})
	global.SERVER_DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.ServerInstance{})
	s.server = http.Server{
		Addr:    provider.AppConfig.LocalServer.BindAddr,
		Handler: router.BuildRouter(),
	}
	s.IsRunning = true
	fmt.Println("thingue server listening at: ", s.server.Addr)
	err := s.server.ListenAndServe() //运行中阻塞
	s.IsRunning = false
	if s.CloseReturnChanel != nil {
		s.CloseReturnChanel <- err.Error()
	}
	fmt.Printf("server closed: %v\n", err)
}

func (s *server) Start() {
	if s.IsRunning {
		return
	}
	go func() {
		s.Serve()
	}()
}

func (s *server) Stop() {
	err := s.server.Close()
	coreprovider.NodeConnProvider.CloseAllConnection()
	coreprovider.AdminConnProvider.CloseAllConnection()
	if err != nil {
		fmt.Printf("server shutdown failed: %v\n", err)
	} else {
		fmt.Println("server gracefully stopped.")
	}
}
