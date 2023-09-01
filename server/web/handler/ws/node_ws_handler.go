package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"thingue-launcher/common/message"
	"thingue-launcher/common/model"
	"thingue-launcher/server/core"
	"thingue-launcher/server/core/provider"
)

func (g *HandlerGroup) NodeWebSocketHandler(c *gin.Context) {
	conn, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	conn.SetCloseHandler(func(code int, text string) error {
		conn.Close()
		fmt.Println("连接关闭之前")
		return nil
	})
	node := model.Node{}
	core.NodeService.NodeOnline(&node)
	callbackMsg := message.ServerConnectCallback(node.ID)
	provider.NodeConnProvider.ConnMap[node.ID] = conn
	err = conn.WriteMessage(websocket.TextMessage, callbackMsg.Pack().GetBytes())
	if err == nil {
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
	}
	conn.Close()
	delete(provider.NodeConnProvider.ConnMap, node.ID)
	core.NodeService.NodeOffline(&node)
}
