package server

import (
	"encoding/json"
	"errors"
	"thingue-launcher/common/app"
	"thingue-launcher/common/model"
	"thingue-launcher/common/util"
)

var appConfig = app.GetAppConfig()

func RegisterAgent(httpUrl string) error {
	registerInfo := model.AgentRegisterInfo{
		DeviceInfo: util.GetDeviceInfo(),
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
