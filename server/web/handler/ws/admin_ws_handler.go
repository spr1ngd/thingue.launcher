package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"thingue-launcher/common/message"
	"thingue-launcher/common/util"
	"thingue-launcher/server/core/provider"
)

func (g *HandlerGroup) AdminWebSocketHandler(c *gin.Context) {
	conn, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	id := provider.AdminConnProvider.AddConn(conn)
	conn.WriteMessage(websocket.TextMessage, util.MapToJson(map[string]interface{}{
		"type": "config",
		"data": id,
	}))
	for {
		// 读取客户端发送的消息
		var msg = message.Message{}
		err = conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("WebSocket read error:", err)
			break
		}
	}
	provider.AdminConnProvider.RemoveConn(id)
	conn.Close()
}
