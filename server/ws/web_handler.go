package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AdminWebSocketHandler(c *gin.Context) {
	conn, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
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
