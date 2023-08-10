package server

import (
	"encoding/json"
	"thingue-launcher/common/util"
)

type ResisterAgentRes struct {
	code int
	msg  string
	data bool
}

func RegisterAgent() {
	reqData, _ := json.Marshal(GetDeviceInfo())
	result, err := util.HttpPost("", reqData)
	if err != nil {
		res := ResisterAgentRes{}
		_ = json.Unmarshal(result, &res)
	}
}
