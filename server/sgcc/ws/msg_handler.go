package ws

import (
	"github.com/labstack/gommon/log"
	"github.com/mitchellh/mapstructure"
	"thingue-launcher/server/global"
	"thingue-launcher/server/sgcc/message"
)

func MsgReceive(msg map[string]any) {
	msgType, ok := msg["type"].(string)
	if ok {
		var err error
		switch msgType {
		case "register_callback":
			callback := &message.RegisterCallback{}
			if err = mapstructure.Decode(msg, callback); err == nil {
				log.Info("注册成功", callback.Id)
			}
		case "init":
			init := &message.Init{}
			if err = mapstructure.Decode(msg, init); err == nil {
				global.SgccService.Init()
			}
		case "deploy":
			deploy := &message.Deploy{}
			if err = mapstructure.Decode(msg, deploy); err == nil {
				global.SgccService.Deploy(deploy)
			}
		case "release":
			release := &message.Release{}
			if err = mapstructure.Decode(msg, release); err == nil {
				global.SgccService.Release(release.Nodes)
			}
		case "restart":
			restart := &message.Restart{}
			if err = mapstructure.Decode(msg, restart); err == nil {
				global.SgccService.Restart(restart.Nodes)
			}
		case "kill":
			kill := &message.Kill{}
			if err = mapstructure.Decode(msg, kill); err == nil {
				global.SgccService.Kill(kill.Nodes)
			}
		default:
			log.Info("不支持的消息类型")
		}
		if err != nil {
			log.Error(msgType, "格式不正确")
		}
	}
}
