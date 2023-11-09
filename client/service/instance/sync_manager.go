package instance

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"thingue-launcher/common/util"
)

type syncManager struct{}

type FileInfo struct {
	FilePath string
	FileName string
	Hash     string
}

func (m *syncManager) StartUpload(id uint) (string, error) {
	var err error
	runner := RunnerManager.GetRunnerById(id)
	cloudRes := strings.TrimSpace(runner.CloudRes)
	if cloudRes == "" {
		return "", errors.New("云资源标识未设置")
	}

	cloudFileInfoSet, _, err := m.getCloudFiles(cloudRes)
	if err != nil {
		return "", err
	}
	localFileInfoSet, localFileInfos, err := m.getLocalFiles(filepath.Dir(runner.ExecPath))
	if err != nil {
		return "", err
	}
	var NewFileNames []string
	var DeleteFiles []string
	var ModFileNames []string
	// 检查删除了哪些文件
	for fileName := range cloudFileInfoSet {
		_, ok := localFileInfoSet[fileName]
		if !ok {
			DeleteFiles = append(DeleteFiles, fileName)
		}
	}
	// 检查新增修改了哪些文件
	for fileName, localFileInfo := range localFileInfoSet {
		cloudFileInfo, ok := cloudFileInfoSet[fileName]
		if ok {
			if localFileInfo.Hash != cloudFileInfo.Hash {
				SyncRequest.UploadFile(fileName, cloudRes, localFileInfo.FilePath)
				ModFileNames = append(ModFileNames, fileName)
			}
		} else {
			SyncRequest.UploadFile(fileName, cloudRes, localFileInfo.FilePath)
			NewFileNames = append(NewFileNames, fileName)
		}
	}
	var results []string
	if len(DeleteFiles) > 0 {
		results = append(results, fmt.Sprintf("删除%d个文件", len(DeleteFiles)))
		SyncRequest.DeleteCloudFiles(DeleteFiles, cloudRes)
	}
	if len(NewFileNames) > 0 {
		results = append(results, fmt.Sprintf("新增%d个文件", len(NewFileNames)))
	}
	if len(ModFileNames) > 0 {
		results = append(results, fmt.Sprintf("修改%d个文件", len(ModFileNames)))
	}
	if len(NewFileNames)+len(ModFileNames) > 0 || len(DeleteFiles) > 0 {
		// 更新记录
		SyncRequest.UpdateCloudFiles(cloudRes, localFileInfos)
	} else {
		results = append(results, "没有文件需要同步")
	}
	return strings.Join(results, "，"), nil
}

func (m *syncManager) StartUpdate(id uint) (string, error) {
	var err error
	runner := RunnerManager.GetRunnerById(id)
	cloudRes := strings.TrimSpace(runner.CloudRes)
	if cloudRes == "" {
		return "", errors.New("云资源标识未设置")
	}
	cloudFileInfoSet, _, err := m.getCloudFiles(cloudRes)
	if err != nil {
		return "", err
	}
	return m.updatePackage(cloudRes, filepath.Dir(runner.ExecPath), cloudFileInfoSet)
}

func (m *syncManager) UpdateCloudRes(cloudRes string) error {
	fmt.Println("检查更新")
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
	cloudFileSet, _, err := m.getCloudFiles(cloudRes)
	// 对每个包进行更新
	if err == nil {
		for _, packagePath := range uniquePackages {
			_, _ = m.updatePackage(cloudRes, filepath.Dir(packagePath), cloudFileSet)
		}
	}
	return err
}

func (m *syncManager) updatePackage(cloudRes string, packagePath string, cloudFileSet map[string]*FileInfo) (string, error) {
	localFileInfoSet, _, err := m.getLocalFiles(packagePath)
	if err != nil {
		return "", err
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
	var results []string
	if len(NewFileNames) > 0 {
		for _, fileName := range NewFileNames {
			m.downloadFile(cloudRes, fileName, packagePath)
		}
		results = append(results, fmt.Sprintf("新增%d个文件", len(NewFileNames)))
	}
	if len(ModFileNames) > 0 {
		for _, fileName := range ModFileNames {
			m.downloadFile(cloudRes, fileName, packagePath)
		}
		results = append(results, fmt.Sprintf("修改%d个文件", len(ModFileNames)))
	}
	if len(DeleteFiles) > 0 {
		for _, deleteFile := range DeleteFiles {
			os.Remove(deleteFile)
		}
		results = append(results, fmt.Sprintf("删除%d个文件", len(DeleteFiles)))
	}
	if len(results) == 0 {
		results = append(results, "没有文件需要同步")
	}
	return strings.Join(results, "，"), err
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
