package provider

import "github.com/gorilla/websocket"

var Conn *websocket.Conn

func SendCloudMessage(data []byte) {
	Conn.WriteMessage(websocket.TextMessage, data)
}
