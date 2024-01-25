package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/message"
	"thingue-launcher/common/model"
	"thingue-launcher/server/core"
	"thingue-launcher/server/core/provider"
)

func (g *HandlerGroup) ClientWebSocketHandler(c *gin.Context) {
	conn, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Zap.Error("WebSocket upgrade error:", err)
		return
	}
	conn.SetCloseHandler(func(code int, text string) error {
		conn.Close()
		return nil
	})
	client := model.Client{}
	core.ClientService.CreateClient(&client)
	callbackMsg := message.ServerConnectCallback(client.ID)
	provider.ClientConnProvider.ConnMap[client.ID] = conn
	err = conn.WriteMessage(websocket.TextMessage, callbackMsg.Pack().GetBytes())
	if err == nil {
		for {
			// 读取客户端发送的消息
			msgType, _, err := conn.ReadMessage()
			if err != nil {
				logger.Zap.Error("WebSocket read error:", err)
				break
			}
			if msgType == websocket.PingMessage {
				conn.WriteMessage(websocket.PongMessage, []byte("ping pong"))
			}
			// todo 处理接收到的消息
		}
	}
	conn.Close()
	delete(provider.ClientConnProvider.ConnMap, client.ID)
	core.ClientService.DeleteClient(&client)
}
