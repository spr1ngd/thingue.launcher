package initialize

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/model"
	"thingue-launcher/common/provider"
	coreprovider "thingue-launcher/server/core/provider"
	"thingue-launcher/server/global"
	"thingue-launcher/server/web/router"
)

type server struct {
	listen            *http.Server
	IsRunning         bool
	CloseReturnChanel chan string
	router            *gin.Engine
	isInitialized     bool
}

var Server = new(server)

func (s *server) Serve() {
	var err error
	//ORM+MQTT
	if !s.isInitialized { //如果是第一次没有初始化
		//s.router = router.BuildRouter(s.StaticFiles) //构建路由
		initServerDB() // 初始化gorm
		initStorageDB()
		initMqttServer()
		s.isInitialized = true
	}
	global.SERVER_DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Client{})
	global.SERVER_DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.ServerInstance{})
	// 构建gin路由
	s.router = router.BuildRouter()
	// Listen
	s.listen = &http.Server{
		Addr:    provider.AppConfig.LocalServer.BindAddr,
		Handler: s.router,
	}
	s.IsRunning = true
	logger.Zap.Info("thingue server listening at: ", s.listen.Addr)
	err = s.listen.ListenAndServe() //运行中阻塞
	s.IsRunning = false
	if s.CloseReturnChanel != nil {
		s.CloseReturnChanel <- err.Error()
	}
	logger.Zap.Info("server closed: %v\n", err)
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
	err := s.listen.Close()
	coreprovider.ClientConnProvider.CloseAllConnection()
	coreprovider.AdminConnProvider.CloseAllConnection()
	if err != nil {
		logger.Zap.Error("server shutdown failed: %v\n", err)
	} else {
		logger.Zap.Info("server gracefully stopped.")
	}
}
