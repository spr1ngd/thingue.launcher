package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type HandlerGroup struct{}

var WsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有来源的 WebSocket 连接
		return true
	},
}
