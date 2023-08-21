package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"thingue-launcher/server/service"
)

func (g *HandlerGroup) StreamerWebSocketHandler(c *gin.Context) {
	conn, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	id := c.Param("id")
	streamer := service.StreamerConnManager.NewStreamerConnector(id, conn)
	streamer.SendConfig()

	for {
		// 读取客户端发送的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("WebSocket read error:", err)
			break
		}
		// 将接收到的消息交给Streamer处理
		streamer.HandleMsg(string(msg))
	}
	conn.Close()
	service.StreamerConnManager.OnStreamerDisconnect(streamer)
}
