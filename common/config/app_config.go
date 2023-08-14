package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var AppConfig = new(Config)

type Config struct {
	ServerUrl         string
	LocalServer       LocalServer
	SystemSettings    SystemSettings
	EnableRestartTask bool
}

type LocalServer struct {
	BasePath  string
	BindAddr  string
	Enable    bool
	AutoStart bool
}

type SystemSettings struct {
	RestartTaskCron string
}

func InitConfig() {
	_, statErr := os.Stat("config.yaml")
	if os.IsNotExist(statErr) {
		fmt.Println("文件不存在")
		file, _ := os.Create("config.yaml")
		defer file.Close()
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("agent.localserver.bindAddr", "0.0.0.0:8080")
	viper.SetDefault("agent.localserver.basepath", "/")
	viper.SetDefault("agent.localserver.enable", true)
	viper.ReadInConfig()
	err := viper.UnmarshalKey("agent", AppConfig)
	if err != nil {
		fmt.Println("配置文件解析失败")
	}
}

func WriteConfig() {
	viper.Set("agent", AppConfig)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
}
