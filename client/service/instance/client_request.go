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

type clientService struct {
	*baseRequest
}

var ClientService = clientService{BaseRequest}

func (s *clientService) GetInstanceSid(clientId uint, instanceId uint) (string, error) {
	params := url.Values{}
	params.Add("clientId", strconv.Itoa(int(clientId)))
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

func (s *clientService) RegisterClient(clientId uint) {
	registerInfo := request.ClientRegisterInfo{
		ClientID:   clientId,
		DeviceInfo: GetDeviceInfo(),
		Instances:  RunnerManager.List(), //todo 去除不必要信息
	}
	reqData, _ := json.Marshal(registerInfo)
	result, err := s.HttpPost("/api/instance/clientRegister", reqData)
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

func (s *clientService) SendProcessState(sid string, stateCode int8, pid int) {
	reqData, _ := json.Marshal(&message.ClientProcessStateUpdate{
		SID:       sid,
		Pid:       pid,
		StateCode: stateCode,
	})
	_, err := s.HttpPost("/api/instance/updateProcessState", reqData)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *clientService) SetRestarting(sid string, restarting bool) {
	params := url.Values{}
	params.Add("sid", sid)
	params.Add("restarting", strconv.FormatBool(restarting))
	_, err := s.HttpGetWithParams("/api/instance/setRestarting", params)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *clientService) ClearPak(sid string) {
	params := url.Values{}
	params.Add("sid", sid)
	_, err := s.HttpGetWithParams("/api/instance/clearPak", params)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *clientService) UpdateStreamerConnected(msg *message.ServerStreamerConnectedUpdate) {
	runner := RunnerManager.GetRunnerById(msg.CID)
	if runner != nil {
		runner.StreamerConnected = msg.Connected
		RunnerManager.RunnerStatusUpdateChanel <- msg.CID
	}
}

func (s *clientService) CollectLogs(traceId string) {
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
