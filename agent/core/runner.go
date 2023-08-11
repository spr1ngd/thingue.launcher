package core

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"thingue-launcher/agent/model"
	"thingue-launcher/common/app"
	"thingue-launcher/common/util"
)

type Runner struct {
	*model.Instance
	ExitSignalChannel chan error
	process           *os.Process
	Pid               int
	IsRunning         bool
}

func (r *Runner) Start() error {
	if r.process != nil {
		return errors.New("实例已在运行")
	}
	// 设置PixelStreamingURL
	var launchArguments []string
	appConfig := app.GetAppConfig()
	if appConfig.ServerUrl != "" {
		wsUrl := util.HttpUrlToStreamerWsUrl(appConfig.ServerUrl)
		launchArguments = append(r.LaunchArguments, "-PixelStreamingURL="+wsUrl+"/"+r.Name)
	} else {
		launchArguments = r.LaunchArguments
	}
	// 运行前
	fmt.Println(r.ExecPath, launchArguments)
	command := exec.Command(r.ExecPath, launchArguments...)
	err := command.Start()
	if err != nil {
		return err
	}
	r.Pid = command.Process.Pid
	r.process = command.Process
	r.IsRunning = true
	go func() {
		exitCode := command.Wait()
		//运行后
		r.Pid = 0
		r.process = nil
		r.IsRunning = false
		r.ExitSignalChannel <- exitCode
	}()
	return nil
}

func (r *Runner) Stop() error {
	if !r.IsRunning {
		return errors.New("实例未在运行")
	}
	err := r.process.Signal(syscall.SIGKILL)
	exitStatus := <-r.ExitSignalChannel
	fmt.Printf("%s进程退出%s\n", r.Name, exitStatus)
	return err
}
