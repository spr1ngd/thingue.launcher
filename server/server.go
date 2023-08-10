package server

import (
	"context"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"thingue-launcher/common/app"
	"time"
)

//go:embed all:frontend/dist
var staticFiles embed.FS

var server http.Server

var serverIsBoot = false

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 缓存预检请求结果的时间，单位为秒

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}

func Startup() {
	appConfig := app.GetAppConfig()
	router := gin.Default()
	router.Use(CorsMiddleware())
	group := router.Group(appConfig.LocalServer.BasePath)
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	group.GET("/static/*filepath", func(c *gin.Context) {
		c.Request.URL.Path = "/frontend/dist" + c.Param("filepath")
		http.FileServer(http.FS(staticFiles)).ServeHTTP(c.Writer, c.Request)
	})
	group.GET("/ws/streamer/:id", StreamerWebSocketHandler)
	group.GET("/ws/player/:streamerId", PlayerWebSocketHandler)
	group.GET("/ws/agent", AgentWebSocketHandler)
	group.GET("/ws/admin", AdminWebSocketHandler)
	server = http.Server{
		Addr:    appConfig.LocalServer.BindAddr,
		Handler: router,
	}
	serverIsBoot = true
	err := server.ListenAndServe()
	serverIsBoot = false
	if err != nil {
		fmt.Printf("Server start failed: %v\n", err)
	}
}

func Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		fmt.Printf("Server shutdown failed: %v\n", err)
	} else {
		fmt.Println("Server gracefully stopped.")
		serverIsBoot = false
	}
}

func GetLocalServerStatus() bool {
	return serverIsBoot
}
