package os_cmd

import (
	"golang.org/x/sys/windows"
	"os"
	"os/exec"
)

func SetCmd(cmd *exec.Cmd) *exec.Cmd {
	cmd.SysProcAttr = &windows.SysProcAttr{HideWindow: true}
	cmd.Stdout = os.Stdout
	return cmd
}
