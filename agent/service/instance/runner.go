package instance

import (
	"errors"
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"thingue-launcher/agent/global"
	"thingue-launcher/common/config"
	"thingue-launcher/common/model"
	"thingue-launcher/common/util"
	"time"
)

type Runner struct {
	*model.ClientInstance
	ExitSignalChannel chan error `json:"-"`
	process           *os.Process
}

func (r *Runner) Start() error {
	if r.process != nil {
		return errors.New("实例已在运行")
	}
	// 设置PixelStreamingURL
	var launchArguments []string
	appConfig := config.AppConfig
	sid, err := NodeService.GetInstanceSid(global.NODE_ID, r.ID)
	if err == nil {
		r.SID = sid
		wsUrl := util.HttpUrlToWsUrl(appConfig.ServerUrl, "/ws/streamer")
		launchArguments = append(r.LaunchArguments, "-PixelStreamingURL="+wsUrl+"/"+r.SID)
	} else {
		launchArguments = r.LaunchArguments
	}
	// 运行前
	fmt.Println(r.ExecPath, launchArguments)
	command := exec.Command(r.ExecPath, launchArguments...)
	err = command.Start()
	if err != nil {
		return err
	}
	r.Pid = command.Process.Pid
	r.process = command.Process
	r.updateStateCode(1)
	r.LastStartAt = time.Now()
	RunnerManager.RunnerStatusUpdateChanel <- r.ID
	go func() {
		exitCode := command.Wait()
		time.Sleep(3 * time.Second)
		r.Pid = 0
		r.process = nil
		r.LastStopAt = time.Now()
		select {
		case r.ExitSignalChannel <- exitCode:
			r.updateStateCode(0)
			RunnerManager.RunnerStatusUpdateChanel <- r.ID
			fmt.Println("退出码发送成功")
		default:
			r.updateStateCode(-1)
			RunnerManager.RunnerUnexpectedExitChanel <- r.ID
			fmt.Println("异常退出")
		}
	}()
	return nil
}

func (r *Runner) Stop() error {
	if r.StateCode != 1 {
		return errors.New("实例未在运行")
	}
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(r.Pid))
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("kill", "-TERM", strconv.Itoa(r.Pid))
	} else {
		return errors.New("不支持的系统")
	}
	cmd.SysProcAttr = &windows.SysProcAttr{HideWindow: true}
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	//err := r.process.Signal(syscall.SIGKILL)
	exitStatus := <-r.ExitSignalChannel
	fmt.Printf("%s进程退出%s\n", r.Name, exitStatus)
	return err
}

func (r *Runner) updateStateCode(stateCode int8) {
	r.StateCode = stateCode
	NodeService.SendProcessState(&model.ProcessStateUpdate{
		SID:       r.SID,
		StateCode: stateCode,
	})
}
