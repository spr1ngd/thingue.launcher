package instance

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"thingue-launcher/common/model"
)

type syncManager struct {
}

func (m *syncManager) StartUpload(id uint) error {
	runner := RunnerManager.GetRunnerById(id)
	configs := NodeService.SyncGetSyncConfig()
	var cloudFiles []model.CloudFile
	var uploadFiles []string
	var relFiles []string
	for _, config := range configs {
		absSyncPath := filepath.Join(filepath.Dir(runner.Instance.ExecPath), config.(string))
		stat, _ := os.Stat(absSyncPath)
		if stat.IsDir() {
			filepath.WalkDir(absSyncPath, func(path string, d fs.DirEntry, err error) error {
				if !d.IsDir() {
					rel, _ := filepath.Rel(absSyncPath, path)
					uploadFiles = append(uploadFiles, path)
					relFiles = append(relFiles, filepath.Join(config.(string), rel))
					cloudFiles = append(cloudFiles, model.CloudFile{
						FileName: filepath.Join(config.(string), rel),
						Hash:     "",
						Res:      "default",
					})
				}
				return nil
			})
		} else {
			uploadFiles = append(uploadFiles, absSyncPath)
			relFiles = append(relFiles, config.(string))
			cloudFiles = append(cloudFiles, model.CloudFile{
				FileName: config.(string),
				Hash:     "",
				Res:      "default",
			})
		}
	}
	for i, uploadFile := range uploadFiles {
		fmt.Println(relFiles[i])
		NodeService.SyncUploadFile(relFiles[i], "default", uploadFile)
	}
	NodeService.SyncUpdateCloudFiles("default", cloudFiles)
	return nil
}

func (m *syncManager) StartDownload(id uint) error {
	runner := RunnerManager.GetRunnerById(id)
	files := NodeService.SyncGetCloudFiles("default")
	for _, file := range files {
		cloudFile := model.CloudFile{}
		_ = mapstructure.Decode(file, &cloudFile)
		downfile := filepath.Join(filepath.Dir(runner.ExecPath), cloudFile.FileName)
		_, err := os.Stat(filepath.Dir(downfile))
		if os.IsNotExist(err) {
			err = os.MkdirAll(filepath.Dir(downfile), os.ModePerm)
		}
		out, _ := os.Create(downfile + ".tmp")

		apiUrl := NodeService.BaseUrl.JoinPath("/storage", cloudFile.Res, strings.ReplaceAll(cloudFile.FileName, "\\", "/")).String()
		resp, err := http.Get(apiUrl)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		io.Copy(out, resp.Body)
		out.Close()
		os.Rename(downfile+".tmp", downfile)
	}
	return nil
}

var SyncManager = &syncManager{}
