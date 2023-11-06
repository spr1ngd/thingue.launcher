package instance

import (
	"errors"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"thingue-launcher/common/model"
	"thingue-launcher/common/util"
)

type syncManager struct {
}

type FileInfo struct {
	FilePath string
	FileName string
	Hash     string
}

func (m *syncManager) StartUpload(id uint) error {
	runner := RunnerManager.GetRunnerById(id)
	res := strings.TrimSpace(runner.CloudRes)
	if res == "" {
		return errors.New("云资源标识未设置")
	}
	configs, err := SyncRequest.GetSyncConfig()
	SyncRequest.GetCloudFiles(res)
	var cloudFiles []model.CloudFile
	var uploadFiles []string
	var relFiles []string
	for _, config := range configs {
		absSyncPath := filepath.Join(filepath.Dir(runner.Instance.ExecPath), config)
		stat, err := os.Stat(absSyncPath)
		if os.IsNotExist(err) {
			continue
		}
		if stat.IsDir() {
			filepath.WalkDir(absSyncPath, func(path string, d fs.DirEntry, err error) error {
				if !d.IsDir() {
					rel, _ := filepath.Rel(absSyncPath, path)
					uploadFiles = append(uploadFiles, path)
					relFiles = append(relFiles, filepath.Join(config, rel))
					cloudFiles = append(cloudFiles, model.CloudFile{
						FileName: filepath.Join(config, rel),
						Hash:     util.CalculateFileHash(path),
					})
				}
				return err
			})
		} else {
			uploadFiles = append(uploadFiles, absSyncPath)
			relFiles = append(relFiles, config)
			cloudFiles = append(cloudFiles, model.CloudFile{
				FileName: config,
				Hash:     util.CalculateFileHash(absSyncPath),
			})
		}
	}
	// 上传文件
	for i, uploadFile := range uploadFiles {
		SyncRequest.UploadFile(relFiles[i], res, uploadFile)
	}
	// 更新记录
	SyncRequest.UpdateCloudFiles(res, cloudFiles)
	return err
}

func (m *syncManager) StartDownload(id uint) error {
	var err error
	runner := RunnerManager.GetRunnerById(id)
	res := strings.TrimSpace(runner.CloudRes)
	if res == "" {
		err = errors.New("云资源标识未设置")
	}
	files, err := SyncRequest.GetCloudFiles(res)
	for _, cloudFile := range files {
		downfile := filepath.Join(filepath.Dir(runner.ExecPath), cloudFile.FileName)
		_, err := os.Stat(filepath.Dir(downfile))
		if os.IsNotExist(err) {
			err = os.MkdirAll(filepath.Dir(downfile), os.ModePerm)
		}
		out, _ := os.Create(downfile + ".tmp")

		apiUrl := SyncRequest.BaseUrl.JoinPath("/storage", cloudFile.Res, strings.ReplaceAll(cloudFile.FileName, "\\", "/")).String()
		resp, err := http.Get(apiUrl)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		io.Copy(out, resp.Body)
		out.Close()
		os.Rename(downfile+".tmp", downfile)
	}
	return err
}

func (m *syncManager) StartUpdate(id uint) error {
	var err error
	runner := RunnerManager.GetRunnerById(id)
	cloudRes := strings.TrimSpace(runner.CloudRes)
	if cloudRes == "" {
		err = errors.New("云资源标识未设置")
	}
	cloudFileSet, cloudFileInfos, err := m.getCloudFiles(cloudRes)
	err = m.UpdatePackage(cloudRes, filepath.Dir(runner.ExecPath), cloudFileSet, cloudFileInfos)
	return err
}

func (m *syncManager) UpdateCloudRes(cloudRes string) error {
	instances := InstanceManager.GetByCloudRes(cloudRes)
	// 筛除使用相同包的实例
	uniqueMap := make(map[string]int)
	var uniquePackages []string
	for _, instance := range instances {
		uniqueMap[instance.ExecPath]++
		if uniqueMap[instance.ExecPath] == 1 {
			uniquePackages = append(uniquePackages, instance.ExecPath)
		}
	}
	// 对每个包进行更新
	cloudFileSet, cloudFileInfos, err := m.getCloudFiles(cloudRes)
	if err == nil {
		for _, packagePath := range uniquePackages {
			_ = m.UpdatePackage(cloudRes, packagePath, cloudFileSet, cloudFileInfos)
		}
	}
	return err
}

func (m *syncManager) UpdatePackage(cloudRes string, packagePath string, cloudFileSet map[string]*FileInfo, cloudFileInfos []*FileInfo) error {
	localFileInfoSet, _, err := m.getLocalFiles(packagePath)
	if err != nil {
		return err
	}
	var NewFileNames []string
	var DeleteFiles []string
	var ModFileNames []string

	// 检查删除了哪些文件
	for name, localFileInfo := range localFileInfoSet {
		_, ok := cloudFileSet[name]
		if !ok {
			DeleteFiles = append(DeleteFiles, localFileInfo.FilePath)
		}
	}

	// 检查新增修改了哪些文件
	for name, cloudFileInfo := range cloudFileSet {
		localFileInfo, ok := localFileInfoSet[name]
		if ok {
			if cloudFileInfo.Hash != localFileInfo.Hash {
				ModFileNames = append(ModFileNames, cloudFileInfo.FileName)
			}
		} else {
			NewFileNames = append(NewFileNames, cloudFileInfo.FileName)
		}
	}

	if len(NewFileNames) > 0 {
		for _, fileName := range NewFileNames {
			m.downloadFile(cloudRes, fileName, packagePath)
		}
	}

	if len(ModFileNames) > 0 {
		for _, fileName := range ModFileNames {
			m.downloadFile(cloudRes, fileName, packagePath)
		}
	}

	if len(DeleteFiles) > 0 {
		for _, deleteFile := range DeleteFiles {
			os.Remove(deleteFile)
		}
	}
	return err
}

func (m *syncManager) getLocalFiles(packagePath string) (map[string]*FileInfo, []*FileInfo, error) {
	var localFileInfoSet = make(map[string]*FileInfo)
	var localFileInfos []*FileInfo
	configs, err := SyncRequest.GetSyncConfig()
	for _, config := range configs {
		configFilePath := filepath.Join(packagePath, config)
		stat, err := os.Stat(configFilePath)
		if os.IsNotExist(err) {
			continue
		}
		if stat.IsDir() {
			filepath.WalkDir(configFilePath, func(path string, d fs.DirEntry, err error) error {
				if !d.IsDir() {
					rel, _ := filepath.Rel(configFilePath, path)
					info := &FileInfo{
						FilePath: path,
						FileName: filepath.Join(config, rel),
						Hash:     util.CalculateFileHash(path),
					}
					localFileInfoSet[info.FileName] = info
					localFileInfos = append(localFileInfos, info)
				}
				return err
			})
		} else {
			info := &FileInfo{
				FilePath: configFilePath,
				FileName: config,
				Hash:     util.CalculateFileHash(configFilePath),
			}
			localFileInfoSet[info.FileName] = info
			localFileInfos = append(localFileInfos, info)
		}
	}
	return localFileInfoSet, localFileInfos, err
}

func (m *syncManager) getCloudFiles(cloudRes string) (map[string]*FileInfo, []*FileInfo, error) {
	var cloudFileSet = make(map[string]*FileInfo)
	var cloudFileInfos []*FileInfo
	cloudFiles, err := SyncRequest.GetCloudFiles(cloudRes)
	if err == nil {
		for _, cloudFile := range cloudFiles {
			info := &FileInfo{
				FilePath: "",
				FileName: cloudFile.FileName,
				Hash:     cloudFile.Hash,
			}
			cloudFileSet[info.FileName] = info
			cloudFileInfos = append(cloudFileInfos, info)
		}
	}
	return cloudFileSet, cloudFileInfos, err
}

func (m *syncManager) downloadFile(cloudRes string, fileName string, packagePath string) error {
	downfile := filepath.Join(packagePath, fileName)
	_, err := os.Stat(filepath.Dir(downfile))
	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(downfile), os.ModePerm)
	}
	out, _ := os.Create(downfile + ".tmp")

	apiUrl := SyncRequest.BaseUrl.JoinPath("/storage", cloudRes, strings.ReplaceAll(fileName, "\\", "/")).String()
	resp, err := http.Get(apiUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(out, resp.Body)
	out.Close()
	os.Rename(downfile+".tmp", downfile)
	return nil
}

var SyncManager = &syncManager{}
