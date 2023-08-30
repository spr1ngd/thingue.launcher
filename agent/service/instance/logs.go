package instance

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"thingue-launcher/common/model"
)

func getLogFiles(instance *model.ClientInstance) []string {
	var logFiles []string
	execDir := filepath.Dir(instance.ExecPath)
	execName := strings.TrimSuffix(filepath.Base(instance.ExecPath), filepath.Ext(instance.ExecPath))
	files, err := os.ReadDir(execDir)
	var hasEngineDir bool
	var hasGameDir bool
	for _, file := range files {
		if file.IsDir() && file.Name() == "Engine" {
			hasEngineDir = true
		}
		if file.IsDir() && file.Name() == execName {
			hasGameDir = true
		}
	}
	var logsDir string
	if hasGameDir && hasEngineDir {
		logs := filepath.Join(execDir, execName, "Saved/Logs")
		_, err := os.Stat(logs)
		if err != nil {
			logsDir = filepath.Join(filepath.Dir(filepath.Dir(execDir)), "Saved/Logs")
		} else {
			logsDir = logs
		}
	} else {
		logsDir = filepath.Join(filepath.Dir(filepath.Dir(execDir)), "Saved/Logs")
	}
	logFile, err := os.ReadDir(logsDir)
	if err == nil {
		for _, entry := range logFile {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".log") && (entry.Name() == instance.Name+".log" || strings.HasPrefix(entry.Name(), instance.Name+"-backup-")) {
				logFiles = append(logFiles, filepath.Join(logsDir, entry.Name()))
			}
		}
	} else {
		fmt.Println("找不到logs目录")
	}
	return logFiles
}
