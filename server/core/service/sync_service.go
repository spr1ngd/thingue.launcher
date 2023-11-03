package service

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"thingue-launcher/common/model"
	"thingue-launcher/server/global"
	"time"
)

type syncService struct{}

var SyncService = syncService{}

func (s *syncService) GetSyncConfig() []string {
	var files []string
	file, err := os.Open("./thingue-launcher/storage/config.json")
	defer file.Close()
	if err == nil {
		var configs []string
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&configs); err == nil {
			for _, config := range configs {
				files = append(files, config)
			}
		}
	}
	return files
}

func (s *syncService) GetCloudFiles(resourceName string) []model.CloudFile {
	var cloudFiles []model.CloudFile
	global.STORAGE_DB.Where("resource_name = ?", resourceName).Find(&cloudFiles)
	return cloudFiles
}

func (s *syncService) UpdateCloudFiles(resourceName string, files []*model.CloudFile) {
	global.STORAGE_DB.Where("resource_name = ?", resourceName).Delete(&model.CloudFile{})
	for _, file := range files {
		file.ResourceName = resourceName
	}
	global.STORAGE_DB.Save(files)

	resource := model.CloudResource{
		Name: resourceName,
	}
	global.STORAGE_DB.Find(&resource)
	resource.LastUpdateAt = time.Now()
	global.STORAGE_DB.Save(&resource)
}

func (s *syncService) UploadFile(res string, _path string, reader io.ReadCloser) error {
	defer reader.Close()
	outPath := filepath.Join("thingue-launcher/storage/", res, _path)
	_, err := os.Stat(filepath.Dir(outPath))
	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
	}
	out, err := os.Create(outPath)
	defer out.Close()
	if err == nil {
		_, err = io.Copy(out, reader)
	}
	return err
}
