package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"thingue-launcher/server/sdp"
)

func StreamerWebSocketHandler(c *gin.Context) {
	conn, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}

	id := c.Param("id")
	fmt.Printf("端点已连接:%s\n", id)
	streamer := sdp.NewStreamer(id, conn)
	defer sdp.UnmapStreamerId(id)
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
		sdp.StreamerMsgHandler(streamer, string(msg))

	}
	conn.Close()
	fmt.Printf("端点已关闭:%s\n", id)
	sdp.OnStreamerDisconnect(streamer)
}
