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

func (g *HandlerGroup) ClientWebSocketHandler(c *gin.Context) {
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
	client := model.Client{}
	core.ClientService.ClientOnline(&client)
	callbackMsg := message.ServerConnectCallback(client.ID)
	provider.ClientConnProvider.ConnMap[client.ID] = conn
	err = conn.WriteMessage(websocket.TextMessage, callbackMsg.Pack().GetBytes())
	if err == nil {
		for {
			// 读取客户端发送的消息
			msgType, _, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("WebSocket read error:", err)
				break
			}
			if msgType == websocket.PingMessage {
				fmt.Println("ping")
				conn.WriteMessage(websocket.PongMessage, []byte("ping pong"))
			}
			// 处理接收到的消息
			// todo
			//fmt.Println(string(msgByte))
		}
	}
	conn.Close()
	delete(provider.ClientConnProvider.ConnMap, client.ID)
	core.ClientService.ClientOffline(&client)
}
