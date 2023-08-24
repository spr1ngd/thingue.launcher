package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"thingue-launcher/server/service"
)

func (g *HandlerGroup) PlayerWebSocketHandler(c *gin.Context) {
	conn, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}

	playerConnector := service.PlayerConnManager.NewConnector(conn)
	// 关联Streamer
	ticket := c.Param("ticket")
	err = service.PlayerConnManager.SetStreamer(playerConnector, ticket)
	if err == nil {
		playerConnector.SendConfig()
		for {
			// 读取客户端发送的消息
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("WebSocket read error:", err)
				break
			}
			// 处理接收到的消息
			playerConnector.HandleMessage(msg)
		}
		playerConnector.Close()
		service.PlayerConnManager.OnPlayerDisConnect(playerConnector)
	} else {
		playerConnector.Close()
	}
}
