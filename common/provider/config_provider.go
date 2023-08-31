package provider

import (
	"fmt"
	"github.com/spf13/pflag"
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
	BasePath           string `json:"basePath"`
	BindAddr           string `json:"bindAddr"`
	AutoStart          bool   `json:"autoStart"`
	UseExternalStatic  bool   `json:"useExternalStatic"`
	ExternalStaticPath string `json:"externalStaticPath"`
}

type SystemSettings struct {
	RestartTaskCron string
}

func InitConfig() {
	// 加载命令行参数
	pflag.StringP("bind", "b", "0.0.0.0:8080", "Sets the server bind address")
	pflag.String("content-path", "/", "Sets the server base content path")
	pflag.String("static-dir", "", "Path to directory containing the web static resources. Defaults use embed")
	pflag.BoolP("help", "h", false, "show help")
	pflag.BoolP("version", "v", false, "show version")
	pflag.Parse()
	pFlagV := viper.New()
	_ = pFlagV.BindPFlag("showHelp", pflag.Lookup("help"))
	_ = pFlagV.BindPFlag("showVersion", pflag.Lookup("version"))
	_ = pFlagV.BindPFlag("bindAddr", pflag.Lookup("bind"))
	_ = pFlagV.BindPFlag("basePath", pflag.Lookup("content-path"))
	_ = pFlagV.BindPFlag("externalStaticPath", pflag.Lookup("static-dir"))
	//pFlagV.Unmarshal(&AppConfig.LocalServer)
	//设置默认值
	viper.SetDefault("agent.localserver.bindAddr", "0.0.0.0:80801")
	viper.SetDefault("agent.localserver.basepath", "/")
	viper.SetDefault("agent.localserver.useexternalstatic", false)
	viper.SetDefault("agent.localserver.externalstaticpath", "./thingue-launcher/static")
	fmt.Println(viper.GetString("agent.localserver.bindaddr"))
	// 加载配置文件
	viper.SetConfigFile("./thingue-launcher/config.yaml")
	_ = viper.ReadInConfig()
	err := viper.UnmarshalKey("agent", &AppConfig)
	if err != nil {
		fmt.Println("配置文件解析失败")
	}
}

func WriteConfig() {
	_, statErr := os.Stat("./thingue-launcher/config.yaml")
	if os.IsNotExist(statErr) {
		err := os.MkdirAll("./thingue-launcher", 0755)
		file, err := os.Create("./thingue-launcher/config.yaml")
		defer file.Close()
		if err != nil {
			panic(err)
		}
	}
	viper.Set("agent", AppConfig)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
}
