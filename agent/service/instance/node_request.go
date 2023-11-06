package instance

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"thingue-launcher/common/message"
	"thingue-launcher/common/request"
	"thingue-launcher/common/response"
	"thingue-launcher/common/util"
)

type nodeService struct {
	*baseRequest
}

var NodeService = nodeService{BaseRequest}

func (s *nodeService) GetInstanceSid(nodeId uint, instanceId uint) (string, error) {
	params := url.Values{}
	params.Add("nodeId", strconv.Itoa(int(nodeId)))
	params.Add("instanceId", strconv.Itoa(int(instanceId)))
	result, err := s.HttpGetWithParams("/api/instance/getInstanceSid", params)
	if err == nil {
		res := response.Response[string]{}
		err = json.Unmarshal(result, &res)
		if err == nil {
			if res.Code != 200 {
				err = errors.New("获取sid失败")
			} else {
				return res.Data, err
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
	result, err := s.HttpPost("/api/instance/nodeRegister", reqData)
	if err == nil {
		res := response.Response[any]{}
		err = json.Unmarshal(result, &res)
		if err != nil {
			if res.Code != 200 {
				fmt.Println("注册信息发送失败")
			}
		} else {
			RunnerManager.StartInternalRunner()
		}
	}
}

func (s *nodeService) SendProcessState(sid string, stateCode int8) {
	reqData, _ := json.Marshal(&message.NodeProcessStateUpdate{
		SID:       sid,
		StateCode: stateCode,
	})
	_, err := s.HttpPost("/api/instance/updateProcessState", reqData)
	if err != nil {
		fmt.Println(err)
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
	var filesToCompress []string
	instances := RunnerManager.List()
	for _, instance := range instances {
		filesToCompress = append(filesToCompress, getLogFiles(instance)...)
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
