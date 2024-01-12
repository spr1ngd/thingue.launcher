package main

import (
	"embed"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"runtime"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/provider"
	"thingue-launcher/server/initialize"
)

var (
	GitCommit  string
	BuildDate  string
	AppVersion string
	//go:embed all:frontend/dist
	staticFiles embed.FS
)

func main() {
	provider.SetVersionBuildInfo(AppVersion, GitCommit, BuildDate)
	provider.SetWebStatic("frontend/dist", staticFiles)
	provider.InitFlagConfig()
	showHelp := viper.GetBool("showHelp")
	if showHelp {
		pflag.Usage()
		return
	}
	showVersion := viper.GetBool("showVersion")
	if showVersion {
		fmt.Printf("ThingUE Server %s v%s\n", runtime.GOOS, provider.VersionInfo.Version)
		fmt.Printf("BuildDate %s\n", provider.VersionInfo.BuildDate)
		fmt.Printf("GitCommit %s\n", provider.VersionInfo.GitCommit)
		return
	}
	logger.InitZapLogger("info")
	initialize.Server.Serve()
}
