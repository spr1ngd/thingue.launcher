package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有来源的 WebSocket 连接
		return true
	},
}

func handleStreamerWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}

	id := c.Param("id")
	fmt.Printf("端点已连接:%s\n", id)

	for {
		// 读取客户端发送的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("WebSocket read error:", err)
			break
		}

		// 处理接收到的消息
		fmt.Println(id, "端点收到消息:", string(msg))

		// 发送响应消息给客户端
		err = conn.WriteMessage(websocket.TextMessage, []byte("Server received: "+string(msg)))
		if err != nil {
			fmt.Println("WebSocket write error:", err)
			break
		}
	}
	conn.Close()
	fmt.Printf("端点已关闭:%s\n", id)
}

func handlePlayerWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}

	id := c.Param("id")
	fmt.Printf("端点已连接:%s\n", id)

	for {
		// 读取客户端发送的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("WebSocket read error:", err)
			break
		}

		// 处理接收到的消息
		fmt.Println(id, "端点收到消息:", string(msg))

		// 发送响应消息给客户端
		err = conn.WriteMessage(websocket.TextMessage, []byte("Server received: "+string(msg)))
		if err != nil {
			fmt.Println("WebSocket write error:", err)
			break
		}
	}
	conn.Close()
	fmt.Printf("端点已关闭:%s\n", id)
}
