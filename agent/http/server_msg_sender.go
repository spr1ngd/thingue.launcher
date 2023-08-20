package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"thingue-launcher/common/config"
	"thingue-launcher/common/model"
	"thingue-launcher/common/util"
)

func GetInstanceSid(nodeId uint, instanceId uint) (string, error) {
	parse, _ := url.Parse(config.AppConfig.ServerUrl)
	result, err := util.HttpGet(parse.JoinPath("/api/instance/getInstanceSid").String() +
		fmt.Sprintf("nodeId=%d&instanceId=%d", nodeId, instanceId))
	if err != nil {
		res := model.JsonStruct{}
		err = json.Unmarshal(result, &res)
		if err != nil {
			if res.Code != 200 {
				err = errors.New("获取sid失败")
			} else {
				return res.Data.(string), err
			}
		}
	}
	return "", err
}
