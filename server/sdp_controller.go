package server

import (
	"fmt"
	"thingue-launcher/server/util"
)

func StreamerMsgHandler(streamer *Streamer, msgStr string) {
	msg := util.JsonStrToMapData(msgStr)
	msgType := msg["type"].(string)
	if msgType == "ping" {
		streamer.SendMsg(util.MapDataToJsonStr(map[string]interface{}{
			"type": "pong",
			"time": msg["time"],
		}))
	} else if msgType == "answer" {
		playerId := int(msg["playerId"].(float64))
		for _, player := range streamer.Players {
			if player.PlayerId == playerId {
				player.SendMsg(msgStr)
			}
		}
	} else if msgType == "iceCandidate" {
		playerId := int(msg["playerId"].(float64))
		for _, player := range streamer.Players {
			if player.PlayerId == playerId {
				player.SendMsg(msgStr)
			}
		}
	} else if msgType == "disconnectPlayer" {
		fmt.Println(msg)
	} else if msgType == "state" {
		fmt.Println(msg)
	} else {
		streamer.SendCloseMsg(1008, "Unsupported message type")
	}
}

func OnStreamerDisconnect(streamer *Streamer) {
}

func PlayerMsgHandler(player *Player, msgStr string) {
	msg := util.JsonStrToMapData(msgStr)
	msgType := msg["type"].(string)
	if msgType == "ping" {
		player.SendMsg(util.MapDataToJsonStr(map[string]interface{}{
			"type": "pong",
			"time": msg["time"],
		}))
	} else if msgType == "offer" {
		streamer := player.Streamer
		streamer.AddPlayer(player)
		msg["playerId"] = player.PlayerId
		streamer.SendMsg(util.MapDataToJsonStr(msg))
	} else if msgType == "iceCandidate" {
		msg["playerId"] = player.PlayerId
		player.Streamer.SendMsg(util.MapDataToJsonStr(msg))
	} else if msgType == "stats" {
	} else if msgType == "kick" {
		player.KickOthers()
	} else {
		player.SendCloseMsg(1008, "Unsupported message type")
	}
}

func OnPlayerDisConnect(player *Player) {
	player.Streamer.SendMsg(util.MapDataToJsonStr(map[string]interface{}{
		"type":     "playerDisconnected",
		"playerId": player.PlayerId,
	}))
}
