package main

import (
	"gopkg.in/yaml.v3"
	"log"
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
		log.Fatalf("Failed to read YAML file: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, provider.Config)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}
	// 初始化server
	service.ThingUE.Init()
	// 连接云端
	ws.ConnManager.StartConnectTask()
	<-done
}
