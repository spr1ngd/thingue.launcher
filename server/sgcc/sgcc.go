package sgcc

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"thingue-launcher/common/constants"
	provider2 "thingue-launcher/common/provider"
	"thingue-launcher/server/core/service"
	"thingue-launcher/server/sgcc/config"
	"thingue-launcher/server/sgcc/provider"
	"thingue-launcher/server/sgcc/utils"
	"thingue-launcher/server/sgcc/ws"
	"time"
)

func Init() {
	// 读取配置文件
	yamlFile, err := os.ReadFile(path.Join(constants.SaveDir, "sgcc.yaml"))
	if err != nil {
		fmt.Printf("Failed to read YAML file: %v\n", err)
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, config.Config)
	provider2.AppConfig.PeerConnectionOptions = config.Config.PeerConnectionOptions
	if err != nil {
		fmt.Printf("Failed to unmarshal YAML: %v", err)
		panic(err)
	}
	// 连接云端
	ws.ConnManager.StartConnectTask()
	// 定时发送状态
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			_ = <-ticker.C
			instances := service.InstanceService.InstanceList()
			for _, instance := range instances {
				status := utils.GetInstanceNodeStatus(instance)
				provider.SendStatus(instance.StreamerId, instance.PakName, status)
			}
		}
	}()
}
