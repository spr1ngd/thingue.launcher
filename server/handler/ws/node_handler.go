package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"thingue-launcher/common/model"
	"thingue-launcher/common/util"
	"thingue-launcher/server/global"
)

func (g *HandlerGroup) NodeWebSocketHandler(c *gin.Context) {
	conn, err := WsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	node := model.Node{}
	global.SERVER_DB.Create(&node)
	str := util.MapDataToJsonStr(map[string]interface{}{
		"type": "ConnectCallback",
		"data": node.ID,
	})
	conn.WriteMessage(websocket.TextMessage, []byte(str))
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
	conn.Close()
	global.SERVER_DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&node.Instances)
	global.SERVER_DB.Delete(&node)
}
