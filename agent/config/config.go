package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Instances     []Instance
	RemoteServers []RemoteServer
	LocalServer   LocalServer
}

type LocalServer struct {
	BasePath  string
	BindAddr  string
	Enable    bool
	AutoStart bool
}

type RemoteServer struct {
	Name   string
	Url    string
	Enable bool
}

type Instance struct {
	Name                   string
	ExecPath               string
	RunParams              []string
	EnableMultiuserControl bool
	EnableAutoControl      bool
	AutoStopDelay          int
	MetaData               map[string]interface{}
}

var appConfig AppConfig

func InitAppConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("agent.localserver.bindAddr", "0.0.0.0:8080")
	viper.SetDefault("agent.localserver.basepath", "/")
	viper.SetDefault("agent.localserver.enable", true)
	viper.ReadInConfig()
	err := viper.UnmarshalKey("agent", &appConfig)
	if err != nil {
		fmt.Println("配置文件解析失败")
	}
	fmt.Println(appConfig)
}

func WriteConfig() {
	viper.Set("agent", &appConfig)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetAppConfig() *AppConfig {
	return &appConfig
}
