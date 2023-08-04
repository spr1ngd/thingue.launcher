package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var appConfig AppConfig

type AppConfig struct {
	ServerUrl   string
	LocalServer LocalServer
}

type LocalServer struct {
	BasePath  string
	BindAddr  string
	Enable    bool
	AutoStart bool
}

func InitConfig() {
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
}

func GetAppConfig() *AppConfig {
	return &appConfig
}

func WriteConfig() {
	viper.Set("agent", &appConfig)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
}
