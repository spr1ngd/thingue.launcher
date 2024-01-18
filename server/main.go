package main

import (
	"embed"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"runtime"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/provider"
	"thingue-launcher/server/initialize"
	"thingue-launcher/server/sgcc"
)

var (
	GitCommit  string
	BuildDate  string
	AppVersion string
	//go:embed all:frontend/dist
	staticFiles embed.FS
)

func main() {
	// 设置信息
	constants.SetEmbed("frontend/dist", staticFiles)
	constants.SetVersionInfo(AppVersion, GitCommit, BuildDate)
	// 初始化配置
	provider.InitFlagConfig()
	// showHelp
	if viper.GetBool("showHelp") {
		pflag.Usage()
		return
	}
	// showVersion
	if viper.GetBool("showVersion") {
		fmt.Printf("ThingUE Server %s v%s\n", runtime.GOOS, constants.VersionInfo.Version)
		fmt.Printf("BuildDate %s\n", constants.VersionInfo.BuildDate)
		fmt.Printf("GitCommit %s\n", constants.VersionInfo.GitCommit)
		return
	}
	// 初始化日志
	logger.InitZapLogger(viper.GetString("logLevel"), "server.log")
	// 启动连接云端
	if viper.GetBool("enableSgcc") {
		sgcc.Init()
	}
	initialize.Server.Serve()
}
