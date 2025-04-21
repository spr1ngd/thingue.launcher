package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/provider"
	"thingue-launcher/server"
	"thingue-launcher/server/initialize"
)

var (
	serverLogLevel string
	turnServer     string
	turnUsername   string
	turnPassword   string
)

func init() {
	serverCmd.Flags().StringVarP(&provider.AppConfig.LocalServer.BindAddr, "bindAddr", "b", "0.0.0.0:8877", "设置服务绑定的地址与端口")
	serverCmd.Flags().BoolVarP(&provider.AppConfig.LocalServer.Tls, "tls", "b", false, "是否开启TLS")
	serverCmd.Flags().StringVar(&provider.AppConfig.LocalServer.ContentPath, "contentPath", "/", "设置服务路径前缀")
	serverCmd.Flags().StringVar(&serverLogLevel, "logLevel", "info", "设置日志级别")
	serverCmd.Flags().StringVar(&provider.AppConfig.LocalServer.StaticDir, "staticDir", "", "Path to directory containing the web static resources. Defaults use embed")
	serverCmd.Flags().StringVar(&turnServer, "turn-server", "", "turn服务地址")
	serverCmd.Flags().StringVar(&turnUsername, "turn-username", "", "turn服务用户名")
	serverCmd.Flags().StringVar(&turnPassword, "turn-password", "", "turn服务密码")
	rootCmd.AddCommand(serverCmd)
}

const tmpl = `iceServers:
  - urls:
    - turn:%s
    username: %s
    credential: %s`

var serverCmd = &cobra.Command{
	Use:   `server`,
	Short: "运行信令服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		if turnServer != "" && turnUsername != "" && turnPassword != "" {
			peerConnectionOptions := fmt.Sprintf(tmpl, turnServer, turnUsername, turnPassword)
			fmt.Println(peerConnectionOptions)
			provider.AppConfig.PeerConnectionOptions = peerConnectionOptions
		}
		server.Init()
		logger.InitZapLogger(serverLogLevel, "server.log")
		initialize.Server.Serve()
		return nil
	},
}
