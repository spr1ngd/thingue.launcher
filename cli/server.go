package main

import (
	"github.com/spf13/cobra"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/provider"
	"thingue-launcher/server"
	"thingue-launcher/server/initialize"
)

var (
	serverLogLevel string
)

func init() {
	serverCmd.Flags().StringVarP(&provider.AppConfig.LocalServer.BindAddr, "bindAddr", "b", "0.0.0.0:8877", "设置服务绑定的地址与端口")
	serverCmd.Flags().StringVar(&provider.AppConfig.LocalServer.ContentPath, "contentPath", "/", "设置服务路径前缀")
	serverCmd.Flags().StringVar(&serverLogLevel, "logLevel", "info", "设置日志级别")
	serverCmd.Flags().StringVar(&provider.AppConfig.LocalServer.StaticDir, "staticDir", "", "Path to directory containing the web static resources. Defaults use embed")
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   `server`,
	Short: "运行信令服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		server.Init()
		logger.InitZapLogger(serverLogLevel, "server.log")
		initialize.Server.Serve()
		return nil
	},
}
