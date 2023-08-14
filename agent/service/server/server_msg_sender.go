package server

import (
	"encoding/json"
	"errors"
	"thingue-launcher/agent/service/instance"
	"thingue-launcher/common/model"
	"thingue-launcher/common/util"
)

func RegisterAgent(httpUrl string) error {
	registerInfo := model.AgentRegisterInfo{
		DeviceInfo: GetDeviceInfo(),
		Instances:  instance.RunnerManager.List(),
	}
	reqData, _ := json.Marshal(registerInfo)
	result, err := util.HttpPost(httpUrl+"api/agent/register", reqData)
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
