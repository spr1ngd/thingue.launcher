package instance

import (
	"thingue-launcher/common/model"
)

type syncRequest struct{}

var SyncRequest = syncRequest{}

func (s *syncRequest) GetSyncConfig() ([]string, error) {
	return nil, nil
}

func (s *syncRequest) GetCloudFiles(res string) ([]*model.CloudFile, error) {
	return nil, nil
}

func (s *syncRequest) UpdateCloudFiles(cloudRes string, files []*FileInfo) {
}

func (s *syncRequest) UploadFile(fileName string, cloudRes string, filePath string) {
}

func (s *syncRequest) DeleteCloudFiles(fileNames []string, cloudRes string) {
}
