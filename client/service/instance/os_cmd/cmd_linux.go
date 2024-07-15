package os_cmd

import (
	"os"
	"os/exec"
)

func SetCmd(cmd *exec.Cmd) *exec.Cmd {
	cmd.Stdout = os.Stdout
	return cmd
}
