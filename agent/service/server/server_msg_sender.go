package server

import (
	"encoding/json"
	"errors"
	"net/url"
	"thingue-launcher/agent/service/instance"
	"thingue-launcher/common/config"
	"thingue-launcher/common/model"
	"thingue-launcher/common/util"
)

func RegisterNode(nodeId uint) error {
	registerInfo := model.NodeRegisterInfo{
		NodeID:     nodeId,
		DeviceInfo: GetDeviceInfo(),
		Instances:  instance.RunnerManager.List(),
	}
	reqData, _ := json.Marshal(registerInfo)
	parse, _ := url.Parse(config.AppConfig.ServerUrl)
	result, err := util.HttpPost(parse.JoinPath("/api/instance/nodeRegister").String(), reqData)
	if err != nil {
		res := model.JsonStruct{}
		err = json.Unmarshal(result, &res)
		if err != nil {
			if res.Code != 200 {
				return errors.New("无法注册")
			}
		}
	}
	return err
}
