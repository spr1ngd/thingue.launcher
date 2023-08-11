package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"thingue-launcher/server/sdp"
)

func PlayerWebSocketHandler(c *gin.Context) {
	conn, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}

	streamerId := c.Param("streamerId")
	fmt.Printf("player端点已连接:%s\n", streamerId)
	player := sdp.NewPlayer(conn)
	streamer := sdp.GetStreamerById(streamerId)
	player.SetStreamer(streamer)
	player.Init()

	for {
		// 读取客户端发送的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("WebSocket read error:", err)
			break
		}

		// 处理接收到的消息
		fmt.Println(streamerId, "端点收到消息:", string(msg))
		sdp.PlayerMsgHandler(player, string(msg))
	}
	conn.Close()
	fmt.Printf("Player端点已关闭:%s\n", streamerId)
	sdp.OnPlayerDisConnect(player)
}
