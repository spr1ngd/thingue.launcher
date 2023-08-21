package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"thingue-launcher/server/service/ws"
)

func (g *HandlerGroup) AdminWebSocketHandler(c *gin.Context) {
	conn, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	ws.AdminWsManager.Connections = append(ws.AdminWsManager.Connections, conn)
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
	for i, connection := range ws.AdminWsManager.Connections {
		if connection == conn {
			ws.AdminWsManager.Connections = append(ws.AdminWsManager.Connections[:i], ws.AdminWsManager.Connections[i+1:]...)
		}
	}
	conn.Close()
}
