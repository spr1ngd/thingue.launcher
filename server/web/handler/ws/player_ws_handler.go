package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"thingue-launcher/server/core"
	"time"
)

func (g *HandlerGroup) PlayerWebSocketHandler(c *gin.Context) {
	conn, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}

	playerConnector := core.PlayerConnManager.NewConnector(conn)
	// 关联Streamer
	ticket := c.Param("ticket")
	err = core.PlayerConnManager.SetStreamer(playerConnector, ticket)
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
		core.PlayerConnManager.OnPlayerDisConnect(playerConnector)
	} else {
		time.Sleep(5 * time.Second)
		playerConnector.Close()
	}
}
