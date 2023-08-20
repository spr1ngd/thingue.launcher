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

	ticket := c.Param("ticket")
	player := service.PlayerConnManager.NewPlayerConnector(conn)
	// 关联Streamer
	sid, err := service.TicketService.GetSidByTicket(ticket)
	if err != nil {
		// todo 明确关闭原因
		player.Close()
	} else {
		streamer := service.StreamerConnManager.GetStreamerConnectorById(sid)
		if streamer != nil {
			player.SetStreamer(streamer)
			player.SendConfig()
			for {
				// 读取客户端发送的消息
				_, msg, err := conn.ReadMessage()
				if err != nil {
					fmt.Println("WebSocket read error:", err)
					break
				}

				// 处理接收到的消息
				fmt.Println("streamerId", "端点收到消息:", string(msg))
				player.HandleMsg(string(msg))
			}

			player.Close()
			fmt.Printf("Player连接已关闭:%s\n", streamer.Instance.Name)
			service.PlayerConnManager.OnPlayerDisConnect(player)
		}
	}
}
