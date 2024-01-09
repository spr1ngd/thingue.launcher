package ws

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"thingue-launcher/common/util"
	"thingue-launcher/server/core/provider"
	"thingue-launcher/server/core/service"
	"time"
)

func (g *HandlerGroup) StreamerWebSocketHandler(c *gin.Context) {
	conn, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	sid := c.Param("id")
	instance := service.InstanceService.GetInstanceBySid(sid)
	if instance != nil {
		streamer := provider.SdpConnProvider.NewStreamer(sid, conn, instance.EnableRelay, instance.EnableRenderControl)
		streamer.SendConfig()
		service.SdpService.OnStreamerConnect(streamer)
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
				streamer.SendPong(msg)
			} else if msgType == "offer" { // for new streamer
				streamer.ForwardMessage(msg)
			} else if msgType == "answer" { // for old streamer
				streamer.ForwardMessage(msg)
			} else if msgType == "iceCandidate" {
				streamer.ForwardMessage(msg)
			} else if msgType == "disconnectPlayer" {
				// todo
				fmt.Println(msg)
			} else if msgType == "state" {
				// todo
				fmt.Println(msg)
			} else if msgType == "rendering" {
				state, err := getRenderingStateFromMsg(msg)
				if err == nil {
					streamer.UpdateRenderingState(state)
					service.InstanceService.UpdateRenderingState(streamer.SID, state)
				}
			} else if msgType == "hotReloadComplete" {
				service.InstanceService.UpdatePak(streamer.SID, "")
			} else if msgType == "nodeRestarted" {
				service.SdpService.OnStreamerNodeRestarted(streamer)
			} else if msgType == "loadComplete" {
				service.SdpService.OnStreamerLoadBundleComplete(streamer)
				service.InstanceService.UpdatePak(streamer.SID, msg["bundleName"].(string))
			} else {
				streamer.SendCloseMsg(1008, "不支持的消息类型")
			}
		}
		service.SdpService.OnStreamerDisconnect(streamer)
	} else {
		time.Sleep(3 * time.Second)
		fmt.Println("未注册的Streamer在尝试连接")
		_ = conn.Close()
	}

}

func getRenderingStateFromMsg(msg map[string]any) (bool, error) {
	value, ok := msg["value"].(bool)
	if ok {
		return value, nil
	} else {
		strValue, ok := msg["value"].(string)
		if ok {
			value, err := strconv.ParseBool(strValue)
			if err == nil {
				return value, nil
			}
		}
	}
	return false, errors.New("类型错误")
}
