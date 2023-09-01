package instance

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"thingue-launcher/common/message"
	"thingue-launcher/common/provider"
	"thingue-launcher/common/request"
	"thingue-launcher/common/response"
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
	if provider.AppConfig.RegisterUrl == "" {
		return "", errors.New("服务未连接")
	}
	parse, _ := url.Parse(provider.AppConfig.RegisterUrl)
	result, err := util.HttpGet(parse.JoinPath("/api/instance/getInstanceSid").String() +
		fmt.Sprintf("?nodeId=%d&instanceId=%d", nodeId, instanceId))
	fmt.Println("result", string(result))
	fmt.Println("err", err)
	if err == nil {
		res := response.Response{}
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
	registerInfo := request.NodeRegisterInfo{
		NodeID:     nodeId,
		DeviceInfo: GetDeviceInfo(),
		Instances:  RunnerManager.List(), //todo 去除不必要信息
	}
	reqData, _ := json.Marshal(registerInfo)
	result, err := util.HttpPost(s.BaseUrl.JoinPath("/api/instance/nodeRegister").String(), reqData)
	if err == nil {
		res := response.Response{}
		err = json.Unmarshal(result, &res)
		if err != nil {
			if res.Code != 200 {
				fmt.Println("注册信息发送失败")
			}
		}
	}
}

func (s *nodeService) SendProcessState(msg *message.NodeProcessStateUpdate) {
	if s.BaseUrl != nil {
		reqData, _ := json.Marshal(msg)
		util.HttpPost(s.BaseUrl.JoinPath("/api/instance/updateProcessState").String(), reqData)
	}
}

func (s *nodeService) UpdateStreamerConnected(msg *message.ServerStreamerConnectedUpdate) {
	runner := RunnerManager.GetRunnerById(msg.CID)
	if runner != nil {
		runner.StreamerConnected = msg.Connected
		RunnerManager.RunnerStatusUpdateChanel <- msg.CID
	}
}

func (s *nodeService) CollectLogs(traceId string) {
	fmt.Println("traceId", traceId)
	var filesToCompress []string
	instances := InstanceManager.List()
	for _, instance := range instances {
		filesToCompress = append(filesToCompress, getLogFiles(&instance)...)
	}

	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)
	for _, filePath := range filesToCompress {
		fmt.Println(filePath)
		_ = util.AddFileToZip(zipWriter, filePath)
	}
	_ = zipWriter.Close()

	apiUrl := s.BaseUrl.JoinPath("/api/instance/uploadLogs").String()
	req, _ := http.NewRequest("POST", apiUrl, &buf)
	req.Header.Set("traceId", traceId)
	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
