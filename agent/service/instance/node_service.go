package instance

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"thingue-launcher/common/config"
	"thingue-launcher/common/model"
	"thingue-launcher/common/model/message"
	"thingue-launcher/common/util"
)

type nodeService struct {
	BaseUrl *url.URL
}

var NodeService = new(nodeService)

func (s *nodeService) SetBaseUrl(baseurl string) {
	s.BaseUrl, _ = url.Parse(baseurl)
}

func (s *nodeService) GetInstanceSid(nodeId uint, instanceId uint) (string, error) {
	if config.AppConfig.ServerUrl == "" {
		return "", errors.New("服务未连接")
	}
	parse, _ := url.Parse(config.AppConfig.ServerUrl)
	result, err := util.HttpGet(parse.JoinPath("/api/instance/getInstanceSid").String() +
		fmt.Sprintf("?nodeId=%d&instanceId=%d", nodeId, instanceId))
	fmt.Println("result", string(result))
	fmt.Println("err", err)
	if err == nil {
		res := model.JsonStruct{}
		err = json.Unmarshal(result, &res)

		if err == nil {
			if res.Code != 200 {
				err = errors.New("获取sid失败")
			} else {
				fmt.Println("res", res.Data)
				return res.Data.(string), err
			}
		}
	}
	return "", err
}

func (s *nodeService) RegisterNode(nodeId uint) {
	registerInfo := model.NodeRegisterInfo{
		NodeID:     nodeId,
		DeviceInfo: GetDeviceInfo(),
		Instances:  RunnerManager.List(),
	}
	reqData, _ := json.Marshal(registerInfo)
	result, err := util.HttpPost(s.BaseUrl.JoinPath("/api/instance/nodeRegister").String(), reqData)
	if err == nil {
		res := model.JsonStruct{}
		err = json.Unmarshal(result, &res)
		if err != nil {
			if res.Code != 200 {
				fmt.Println("注册信息发送失败")
			}
		}
	}
}

func (s *nodeService) SendProcessState(request *model.ProcessStateUpdate) {
	reqData, _ := json.Marshal(request)
	util.HttpPost(s.BaseUrl.JoinPath("/api/instance/updateProcessState").String(), reqData)
}

func (s *nodeService) Control(data message.ControlMsg) {

}

func (s *nodeService) Update(data message.UpdateMsg) {

}
