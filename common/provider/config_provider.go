package provider

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"thingue-launcher/common/constants"
)

var AppConfig = new(Config)

var configFile = path.Join(constants.SAVE_DIR, constants.CONFIG_NAME)

type Config struct {
	ServerURL      string         `json:"serverURL" yaml:"serverURL"`
	LocalServer    LocalServer    `json:"localServer" yaml:"localServer"`
	SystemSettings SystemSettings `json:"systemSettings" yaml:"systemSettings"`
}

type LocalServer struct {
	ContentPath       string `json:"contentPath" yaml:"contentPath"`
	BindAddr          string `json:"bindAddr" yaml:"bindAddr"`
	AutoStart         bool   `json:"autoStart" yaml:"autoStart"`
	UseExternalStatic bool   `json:"useExternalStatic" yaml:"useExternalStatic"`
	StaticDir         string `json:"staticDir" yaml:"staticDir"`
}

type SystemSettings struct {
	EnableRestartTask bool   `json:"enableRestartTask" yaml:"enableRestartTask"`
	RestartTaskCron   string `json:"restartTaskCron" yaml:"restartTaskCron"`
}

func InitFlagConfig() {
	pflag.StringP("bind", "b", "0.0.0.0:8877", "Sets the server bind address")
	pflag.String("content-path", "/", "Sets the server base content path")
	pflag.String("static-dir", "", "Path to directory containing the web static resources. Defaults use embed")
	pflag.BoolP("help", "h", false, "show help")
	pflag.BoolP("version", "v", false, "show version")
	pflag.Parse()
	_ = viper.BindPFlag("showHelp", pflag.Lookup("help"))
	_ = viper.BindPFlag("showVersion", pflag.Lookup("version"))
	_ = viper.BindPFlag("bindAddr", pflag.Lookup("bind"))
	_ = viper.BindPFlag("contentPath", pflag.Lookup("content-path"))
	_ = viper.BindPFlag("staticDir", pflag.Lookup("static-dir"))
	_ = viper.Unmarshal(&AppConfig.LocalServer)
	if AppConfig.LocalServer.StaticDir != "" {
		AppConfig.LocalServer.UseExternalStatic = true
	}
}

func InitConfigFromFile() {
	viper.SetDefault("serverURL", "http://localhost:8877/")
	viper.SetDefault("localServer.bindAddr", "0.0.0.0:8877")
	viper.SetDefault("localServer.contentPath", "/")
	viper.SetDefault("localServer.useExternalStatic", false)
	viper.SetDefault("localServer.staticDir", "./thingue-launcher/static")
	viper.SetDefault("localServer.autoStart", true)
	// 加载配置文件
	viper.SetConfigFile(configFile)
	_ = viper.ReadInConfig()
	//_ = viper.UnmarshalKey("agent", &AppConfig)
	viper.Unmarshal(&AppConfig)
}

func WriteConfigToFile() {
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path.Dir(configFile), 0755)
		file, err := os.Create(configFile)
		defer file.Close()
		if err != nil {
			panic(err)
		}
	}
	data, err := yaml.Marshal(&AppConfig)
	err = os.WriteFile(configFile, data, 0755)
	if err != nil {
		fmt.Println(err.Error())
	}
}
