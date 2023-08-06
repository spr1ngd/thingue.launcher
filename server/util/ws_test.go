package util

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"testing"
)

func TestConnect(t *testing.T) {
	ws, err := websocket.Dial("ws://127.0.0.1:8080/", "", "http://localhost/")
	if err != nil {
		log.Println("WebSocket连接失败：", err)
	}
	fmt.Println(ws.IsServerConn())
	fmt.Println(ws.IsClientConn())
	// 发送数据到WebSocket服务器
	message := "Hello, WebSocket!"
	if _, err := ws.Write([]byte(message)); err != nil {
		log.Println("发送消息失败：", err)
	}
	for {
		// 接收来自WebSocket服务器的响应
		response := make([]byte, 512)
		n, err := ws.Read(response)
		if err != nil {
			log.Println("接收响应失败：", err)
			break
		}
		// 打印WebSocket服务器的响应
		fmt.Printf("收到响应：%s\n", response[:n])
	}
	ws.Close()

}
