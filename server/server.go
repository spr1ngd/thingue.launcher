package server

import (
	"context"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

func Startup(addr string, basePath string) {
	router := gin.Default()
	router.Use(CorsMiddleware())
	group := router.Group(basePath)
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	group.Any("/static/*filepath", func(c *gin.Context) {
		staticServer := http.FileServer(http.FS(staticFiles))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})
	group.GET("/ws/streamer/:id", handleStreamerWebSocket)
	group.GET("/ws/player/:streamerId", handlePlayerWebSocket)
	server = http.Server{
		Addr:    addr,
		Handler: router,
	}
	serverIsBoot = true
	err := server.ListenAndServe()
	serverIsBoot = false
	if err != nil {
		fmt.Printf("Server start failed: %v\n", err)
	}
}

func Start() {

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

func GetServerStatus() bool {
	return serverIsBoot
}
