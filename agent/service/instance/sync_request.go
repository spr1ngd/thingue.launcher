package instance

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
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

func (s *syncRequest) UpdateCloudFiles(res string, files []model.CloudFile) {
	reqData, _ := json.Marshal(files)
	params := url.Values{}
	params.Add("res", res)
	_, _ = s.HttpPostWithParams("/api/sync/updateCloudFiles", params, reqData)
}

func (s *syncRequest) UploadFile(path string, res string, filePath string) {
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
