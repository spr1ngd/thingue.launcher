package agent

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func UnrealStart(exePath string, param ...string) int {
	command := exec.Command(exePath, param...)
	err := command.Start()
	if err != nil {
		fmt.Println("启动进程时发生错误:", err)
		return 0
	}
	pid := command.Process.Pid
	fmt.Println("进程号", pid)
	return pid
}

func UnrealStop(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	// 杀死进程
	err = process.Signal(syscall.SIGKILL)
	if err != nil {
		return false
	}
	return true
}

func UnrealRestart() {

}
