package util

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestUE(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	var ps *os.Process

	go func() {
		sig := <-interrupt
		fmt.Printf("接收到信号：%v\n", sig)

		// 在这里执行需要的清理操作
		fmt.Println("执行清理操作...")
		time.Sleep(2 * time.Second) // 模拟清理操作
		fmt.Println("清理操作完成")

		os.Exit(0) // 正常退出程序
	}()

	fmt.Println("程序运行中，请按 Ctrl+C 中断程序。")

	command := exec.Command("E:\\UE\\ue4-game\\game\\Binaries\\Win64\\game.exe", "-RenderOffScreen")
	command.Start()
	ps = command.Process
	fmt.Println("进程开始")
	command.Wait()
	fmt.Println("进程结束")
	ps.Kill()
}
