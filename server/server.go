package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var server http.Server

func Start(addr string, basePath string) {
	r := gin.Default()
	group := r.Group(basePath)
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//group.Static("/", "./frontend")
	server = http.Server{
		Addr:    addr,
		Handler: r,
	}
	server.ListenAndServe()
}

func Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown failed: %v\n", err)
	}
	fmt.Println("Server gracefully stopped.")
}
