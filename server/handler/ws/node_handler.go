package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"thingue-launcher/common/model"
	"thingue-launcher/common/util"
	"thingue-launcher/server/global"
	"thingue-launcher/server/service/ws"
)

func (g *HandlerGroup) NodeWebSocketHandler(c *gin.Context) {
	conn, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	node := model.Node{}
	global.SERVER_DB.Create(&node)
	str := util.MapDataToJsonStr(map[string]any{
		"type": "ConnectCallback",
		"data": node.ID,
	})
	ws.NodeWsManager.ConnMap[node.ID] = conn
	err = conn.WriteMessage(websocket.TextMessage, []byte(str))
	if err != nil {
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
	delete(ws.NodeWsManager.ConnMap, node.ID)
	global.SERVER_DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&node.Instances)
	global.SERVER_DB.Delete(&node)
}
