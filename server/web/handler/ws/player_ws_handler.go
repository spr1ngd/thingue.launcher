package ws

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/util"
	"thingue-launcher/server/core/provider"
	"thingue-launcher/server/core/service"
	"time"
)

func (g *HandlerGroup) PlayerWebSocketHandler(c *gin.Context) {
	conn, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Zap.Error("WebSocket upgrade error:", err)
		return
	}
	userData := map[string]string{}
	_ = c.ShouldBindQuery(userData)
	player := provider.SdpConnProvider.NewPlayer(conn)
	player.UserData = userData
	// 连接Streamer
	err = service.SdpService.ConnectStreamer(player, c.Param("ticket"))
	if err == nil {
		player.SendConfig()
		for {
			// 接收消息
			_, msgStr, err := conn.ReadMessage()
			if err != nil {
				break
			}
			msg := util.JsonStrToMapData(msgStr)
			// 处理不同消息类型
			msgType := msg["type"].(string)
			if msgType == "ping" {
				player.SendPong(msg)
			} else if msgType == "listStreamers" {
				player.ListStreamers()
			} else if msgType == "offer" { // for old streamer
				player.Offer(msg)
				service.SdpService.OnPlayerPaired(player)
			} else if msgType == "subscribe" { // for new streamer
				player.Subscribe()
				service.SdpService.OnPlayerPaired(player)
			} else if msgType == "answer" { // for new streamer
				player.ForwardMessage(msg)
			} else if msgType == "iceCandidate" {
				player.ForwardMessage(msg)
			} else if msgType == "stats" {
				//todo
			} else if msgType == "kick" {
				player.KickOthers()
			} else {
				player.SendCloseMsg(1008, "不支持的消息类型")
			}
		}
		service.SdpService.OnPlayerDisConnect(player)
	} else {
		// 无法连接Streamer
		time.Sleep(3 * time.Second)
		player.Close()
	}
}
