package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"os/signal"
	"syscall"
	"thingue-launcher/sgcc-adapter/provider"
	"thingue-launcher/sgcc-adapter/service"
	"thingue-launcher/sgcc-adapter/ws"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()
	// 初始化配置
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("Failed to read YAML file: %v\n", err)
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, provider.Config)
	if err != nil {
		fmt.Printf("Failed to unmarshal YAML: %v", err)
		panic(err)
	}
	// 初始化server
	service.ThingUE.Init()
	// 连接云端
	ws.ConnManager.StartConnectTask()
	<-done
}
