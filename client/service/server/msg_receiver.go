package server

import (
	"errors"
	"thingue-launcher/client/global"
	"thingue-launcher/client/service/instance"
	"thingue-launcher/common/message"
	"thingue-launcher/common/message/types"
)

func MsgReceive(msg message.Message) error {
	switch msg.Type {
	case types.ServerConnectCallback:
		clientId := msg.Data.(float64)
		global.CLIENT_ID = uint(clientId)
		instance.ClientService.RegisterClient(global.CLIENT_ID)
	case types.ServerProcessControl:
		processControl := msg.RecvServerProcessControl()
		instance.RunnerManager.ExecCommand(
			processControl.CID, processControl.Command)
	case types.ServerStreamerConnectedUpdate:
		update := msg.RecvServerStreamerConnectedUpdate()
		instance.ClientService.UpdateStreamerConnected(update)
	case types.ServerCollectClientLogs:
		traceId := msg.Data.(string)
		instance.ClientService.CollectLogs(traceId)
	case types.SyncUpdate:
		res := msg.Data.(string)
		instance.SyncManager.UpdateCloudRes(res)
	default:
		return errors.New("不支持的消息类型")
	}
	return nil
}
