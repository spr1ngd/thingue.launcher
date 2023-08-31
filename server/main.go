package main

import (
	"embed"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"thingue-launcher/common/provider"
	"thingue-launcher/server/initialize"
)

var (
	GitCommit string
	BuildDate string
	//go:embed all:frontend/dist
	staticFiles embed.FS
)

func main() {
	provider.SetVersionBuildInfo(GitCommit, BuildDate)
	provider.SetWebStatic("frontend/dist", staticFiles)
	provider.InitConfig()
	fmt.Println(provider.AppConfig.LocalServer.BindAddr)
	fmt.Println(viper.GetString("agent.localserver.bindaddr"))
	showHelp := viper.GetBool("showHelp")
	if showHelp {
		pflag.Usage()
		return
	}
	showVersion := viper.GetBool("showVersion")
	if showVersion {
		fmt.Printf("ThingUE Server v%s\n", provider.VersionInfo.Version)
		fmt.Printf("Build Date %s\n", provider.VersionInfo.BuildDate)
		fmt.Printf("Git Commit %s\n", provider.VersionInfo.GitCommit)
		return
	}
	initialize.Server.Serve()
}
