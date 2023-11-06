package server

import (
	"errors"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/service/instance"
	"thingue-launcher/common/message"
	"thingue-launcher/common/message/types"
)

func MsgReceive(msg message.Message) error {
	switch msg.Type {
	case types.ServerConnectCallback:
		nodeId := msg.Data.(float64)
		global.NODE_ID = uint(nodeId)
		instance.NodeService.RegisterNode(global.NODE_ID)
	case types.ServerProcessControl:
		processControl := msg.RecvServerProcessControl()
		instance.RunnerManager.ExecCommand(
			processControl.CID, processControl.Command)
	case types.ServerStreamerConnectedUpdate:
		update := msg.RecvServerStreamerConnectedUpdate()
		instance.NodeService.UpdateStreamerConnected(update)
	case types.ServerCollectNodeLogs:
		traceId := msg.Data.(string)
		instance.NodeService.CollectLogs(traceId)
	case types.SyncUpdate:
		res := msg.Data.(string)
		instance.SyncManager.UpdateCloudRes(res)
	default:
		return errors.New("不支持的消息类型")
	}
	return nil
}
