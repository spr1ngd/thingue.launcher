package instance

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/model"
	"thingue-launcher/common/response"
)

type syncRequest struct {
	*baseRequest
}

var SyncRequest = syncRequest{BaseRequest}

func (s *syncRequest) GetSyncConfig() ([]string, error) {
	resp, err := s.HttpGet("/api/sync/getSyncConfig")
	if err == nil {
		res := response.Response[[]string]{}
		err = json.Unmarshal(resp, &res)
		if err == nil {
			if res.Code != 200 {
				err = errors.New("获取同步配置失败")
			} else {
				if res.Data != nil {
					return res.Data, err
				}
			}
		}
	}
	return nil, err
}

func (s *syncRequest) GetCloudFiles(res string) ([]*model.CloudFile, error) {
	params := url.Values{}
	params.Add("res", res)
	resp, err := s.HttpGetWithParams("/api/sync/getCloudFiles", params)
	if err == nil {
		res := response.Response[[]*model.CloudFile]{}
		err = json.Unmarshal(resp, &res)
		if err == nil {
			if res.Code != 200 {
				//err = errors.New("获取sid失败")
			} else {
				return res.Data, err
			}
		}
	}
	return nil, err
}

func (s *syncRequest) UpdateCloudFiles(cloudRes string, files []*FileInfo) {
	reqData, _ := json.Marshal(files)
	params := url.Values{}
	params.Add("res", cloudRes)
	_, _ = s.HttpPostWithParams("/api/sync/updateCloudFiles", params, reqData)
}

func (s *syncRequest) UploadFile(fileName string, cloudRes string, filePath string) {
	var buf bytes.Buffer
	file, err := os.Open(filePath)
	if err != nil {
		logger.Zap.Error(err)
	}
	_, err = io.Copy(&buf, file)
	defer file.Close()

	apiUrl := s.BaseUrl.JoinPath("/api/sync/uploadFile").String()
	req, _ := http.NewRequest("POST", apiUrl, &buf)
	req.Header.Set("name", fileName)
	req.Header.Set("res", cloudRes)
	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		logger.Zap.Error(err)
		return
	}
}

func (s *syncRequest) DeleteCloudFiles(fileNames []string, cloudRes string) {
	reqData, _ := json.Marshal(fileNames)
	params := url.Values{}
	params.Add("res", cloudRes)
	_, _ = s.HttpPostWithParams("/api/sync/deleteCloudFiles", params, reqData)
}
