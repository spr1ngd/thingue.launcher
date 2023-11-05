package instance

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"thingue-launcher/common/message"
	"thingue-launcher/common/model"
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
	if provider.AppConfig.ServerURL == "" {
		return "", errors.New("服务未连接")
	}
	parse, _ := url.Parse(provider.AppConfig.ServerURL)
	result, err := util.HttpGet(parse.JoinPath("/api/instance/getInstanceSid").String() +
		fmt.Sprintf("?nodeId=%d&instanceId=%d", nodeId, instanceId))
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
		} else {
			RunnerManager.StartInternalRunner()
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

func (s *nodeService) SyncGetSyncConfig() []any {
	apiUrl := s.BaseUrl.JoinPath("/api/sync/getSyncConfig").String()
	resp, err := util.HttpGet(apiUrl)
	if err == nil {
		res := response.Response{}
		err = json.Unmarshal(resp, &res)
		if err == nil {
			if res.Code != 200 {
				//err = errors.New("获取sid失败")
			} else {
				fmt.Println("res", res.Data)
				return res.Data.([]any)
			}
		}
	}
	return nil
}

func (s *nodeService) SyncGetCloudFiles(res string) []any {
	apiUrl := s.BaseUrl.JoinPath("/api/sync/getCloudFiles").String()
	params := url.Values{}
	params.Add("res", res)
	resp, err := util.HttpGet(apiUrl + "?" + params.Encode())
	if err == nil {
		res := response.Response{}
		err = json.Unmarshal(resp, &res)
		if err == nil {
			if res.Code != 200 {
				//err = errors.New("获取sid失败")
			} else {
				fmt.Println("res", res.Data)
				return res.Data.([]any)
			}
		}
	}
	return nil
}

func (s *nodeService) SyncUpdateCloudFiles(res string, files []model.CloudFile) {
	reqData, _ := json.Marshal(files)
	//apiUrl := s.BaseUrl.JoinPath("/api/sync/updateCloudFiles?res=" + res).String()
	apiUrl := s.BaseUrl.JoinPath("/api/sync/updateCloudFiles").String()
	_, _ = util.HttpPost(apiUrl, reqData)
}

func (s *nodeService) SyncUploadFile(path string, res string, filePath string) {
	var buf bytes.Buffer
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	_, err = io.Copy(&buf, file)
	defer file.Close()

	apiUrl := s.BaseUrl.JoinPath("/api/sync/uploadFile").String()
	req, _ := http.NewRequest("POST", apiUrl, &buf)
	req.Header.Set("path", path)
	req.Header.Set("res", res)
	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
