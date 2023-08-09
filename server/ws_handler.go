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

func AgentWebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}

	for {
		// 读取客户端发送的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("WebSocket read error:", err)
			break
		}

		// 处理接收到的消息
		fmt.Println(string(msg))
	}
	conn.Close()
}

func StreamerWebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}

	id := c.Param("id")
	fmt.Printf("端点已连接:%s\n", id)
	streamer := NewStreamer(id, conn)
	defer UnmapStreamerId(id)
	streamer.Init()

	for {
		// 读取客户端发送的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("WebSocket read error:", err)
			break
		}

		// 处理接收到的消息
		fmt.Println(id, "端点收到消息:", string(msg))
		StreamerMsgHandler(streamer, string(msg))

	}
	conn.Close()
	fmt.Printf("端点已关闭:%s\n", id)
	OnStreamerDisconnect(streamer)
}

func PlayerWebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}

	streamerId := c.Param("streamerId")
	fmt.Printf("player端点已连接:%s\n", streamerId)
	player := NewPlayer(conn)
	streamer := GetStreamerById(streamerId)
	player.SetStreamer(streamer)
	player.Init()

	for {
		// 读取客户端发送的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("WebSocket read error:", err)
			break
		}

		// 处理接收到的消息
		fmt.Println(streamerId, "端点收到消息:", string(msg))
		PlayerMsgHandler(player, string(msg))
	}
	conn.Close()
	fmt.Printf("Player端点已关闭:%s\n", streamerId)
	OnPlayerDisConnect(player)
}
