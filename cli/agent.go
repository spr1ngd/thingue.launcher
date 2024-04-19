package main

import (
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"thingue-launcher/client"
	"thingue-launcher/client/service"
	"thingue-launcher/common/logger"
)

var (
	serverAddr        string
	agentLogLevel     string
	enableRestartTask bool
)

func init() {
	agentCmd.Flags().StringVarP(&serverAddr, "serverAddr", "s", "http://127.0.0.1:8877", "信令服务地址")
	agentCmd.Flags().StringVarP(&agentLogLevel, "logLevel", "l", "info", "日志级别 debug info warn error")
	agentCmd.Flags().BoolVarP(&enableRestartTask, "enableRestartTask", "r", false, "开启定时重启任务")
	rootCmd.AddCommand(agentCmd)
}

var agentCmd = &cobra.Command{
	Use:   `agent`,
	Short: "运行Agent连接信令服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.InitZapLogger(agentLogLevel, "agent.log")
		client.Init()
		service.RunnerManager.Init()
		// 监听实例异常退出
		go func() {
			for {
				id := <-service.RunnerManager.RunnerUnexpectedExitChanel
				logger.Zap.Warn(id, "实例异常退出")
			}
		}()
		// 监听实例状态变化
		go func() {
			for {
				id := <-service.RunnerManager.RunnerStatusUpdateChanel
				logger.Zap.Info(id, "实例状态变化")
			}
		}()
		// 监听服务连接状态
		go func() {
			for {
				wsUrl := <-service.ServerConnManager.ServerConnUpdateChanel
				logger.Zap.Info(wsUrl, "服务连接状态变化")
			}
		}()
		if enableRestartTask {
			// 开启自动重启任务
			logger.Zap.Info("开启自动重启任务")
			service.RunnerRestartTaskManager.Start()
		}
		//连接server
		logger.Zap.Info("连接server")
		err := service.ServerConnManager.SetServerAddr(serverAddr)
		if err != nil {
			panic(err)
		}
		service.ServerConnManager.StartConnectTask()

		// 创建一个通道来接收操作系统的信号
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		// 等待信号
		<-sig
		logger.Zap.Info("程序退出前执行清理操作")
		service.RunnerManager.CloseAllRunner()
		return nil
	},
}
