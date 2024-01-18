package sgcc

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"thingue-launcher/server/global"
	"thingue-launcher/server/sgcc/provider"
	"thingue-launcher/server/sgcc/service"
	"thingue-launcher/server/sgcc/ws"
)

func Init() {
	// 读取配置文件
	yamlFile, err := os.ReadFile("sgcc.yaml")
	if err != nil {
		fmt.Printf("Failed to read YAML file: %v\n", err)
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, provider.Config)
	if err != nil {
		fmt.Printf("Failed to unmarshal YAML: %v", err)
		panic(err)
	}
	// 创建服务
	global.SgccService = &service.SgccService{}
	// 连接云端
	ws.ConnManager.StartConnectTask()
}
