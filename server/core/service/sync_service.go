package service

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/message"
	"thingue-launcher/common/model"
	"thingue-launcher/server/core/provider"
	"thingue-launcher/server/global"
	"time"
)

type syncService struct{}

var SyncService = syncService{}

func (s *syncService) ListCloudRes() []*model.CloudRes {
	var cloudResList []*model.CloudRes
	global.STORAGE_DB.Find(&cloudResList)
	return cloudResList
}

func (s *syncService) GetSyncConfig() []string {
	var files []string
	file, err := os.Open(constants.SAVE_DIR + "storage/config.json")
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

func (s *syncService) GetCloudFiles(res string) []model.CloudFile {
	var cloudFiles []model.CloudFile
	global.STORAGE_DB.Where("res = ?", res).Find(&cloudFiles)
	return cloudFiles
}

func (s *syncService) UpdateCloudFiles(res string, files []*model.CloudFile) {
	global.STORAGE_DB.Where("res = ?", res).Delete(&model.CloudFile{})
	for _, file := range files {
		file.Res = res
	}
	global.STORAGE_DB.Save(files)

	resource := model.CloudRes{
		Name: res,
	}
	global.STORAGE_DB.Find(&resource)
	resource.LastUpdateAt = time.Now()
	global.STORAGE_DB.Save(&resource)
	updateMsg := message.SyncUpdate(res)
	provider.ClientConnProvider.SendToAllClients(updateMsg.Pack())
}

func (s *syncService) UploadFile(res string, name string, reader io.ReadCloser) error {
	defer reader.Close()
	outPath := filepath.Join("thingue-launcher/storage/", res, name)
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

func (s *syncService) DeleteFiles(res string, names []string) {
	for _, name := range names {
		os.Remove(filepath.Join("thingue-launcher/storage/", res, name))
	}
}

func (s *syncService) DeleteRes(names []string) int64 {
	var deleteCount int64
	var res []model.CloudRes
	tx := global.STORAGE_DB.Delete(&res, names)
	tx.Count(&deleteCount)
	return deleteCount
}

func (s *syncService) CreateCloudRes(res *model.CloudRes) error {
	tx := global.STORAGE_DB.Create(res)
	return tx.Error
}

func (s *syncService) UpdateCloudRes(res *model.CloudRes) error {
	tx := global.STORAGE_DB.Updates(res)
	return tx.Error
}
