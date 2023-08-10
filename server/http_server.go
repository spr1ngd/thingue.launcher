package server

import (
	"context"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"thingue-launcher/common/app"
	"thingue-launcher/common/model/message"
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
	baseGroup := router.Group(appConfig.LocalServer.BasePath)

	//API
	apiGroup := baseGroup.Group("/api")
	apiGroup.POST("/agent/register", func(c *gin.Context) {
		var deviceInfo message.DeviceInfo
		if err := c.ShouldBindJSON(&deviceInfo); err != nil {
		}
		fmt.Printf("服务端收到消息%v\n", deviceInfo)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//静态资源
	baseGroup.GET("/static/*filepath", func(c *gin.Context) {
		c.Request.URL.Path = "/frontend/dist" + c.Param("filepath")
		http.FileServer(http.FS(staticFiles)).ServeHTTP(c.Writer, c.Request)
	})

	//WebSocket
	wsGroup := baseGroup.Group("/ws")
	{
		wsGroup.GET("/streamer/:id", StreamerWebSocketHandler)
		wsGroup.GET("/player/:streamerId", PlayerWebSocketHandler)
		wsGroup.GET("/agent", AgentWebSocketHandler)
		wsGroup.GET("/admin", AdminWebSocketHandler)
	}

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
