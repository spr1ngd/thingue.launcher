package ws

import (
	"github.com/mitchellh/mapstructure"
	"thingue-launcher/common/logger"
	"thingue-launcher/server/sgcc/message"
	"thingue-launcher/server/sgcc/service"
)

func MsgReceive(msg map[string]any) {
	msgType, ok := msg["type"].(string)
	if ok {
		var err error
		switch msgType {
		case "register_callback":
			callback := &message.RegisterCallback{}
			if err = mapstructure.Decode(msg, callback); err == nil {
				logger.Zap.Info("注册成功", callback.Id)
			}
		case "init":
			init := &message.Init{}
			if err = mapstructure.Decode(msg, init); err == nil {
				service.SgccService.Init()
			}
		case "deploy":
			deploy := &message.Deploy{}
			if err = mapstructure.Decode(msg, deploy); err == nil {
				service.SgccService.Deploy(deploy)
			}
		case "release":
			release := &message.Release{}
			if err = mapstructure.Decode(msg, release); err == nil {
				service.SgccService.Release(release.Nodes)
			}
		case "restart":
			restart := &message.Restart{}
			if err = mapstructure.Decode(msg, restart); err == nil {
				service.SgccService.Restart(restart.Nodes)
			}
		case "kill":
			kill := &message.Kill{}
			if err = mapstructure.Decode(msg, kill); err == nil {
				service.SgccService.Kill(kill.Nodes)
			}
		default:
			logger.Zap.Error("不支持的消息类型")
		}
		if err != nil {
			logger.Zap.Error(msgType, "格式不正确")
		}
	}
}
