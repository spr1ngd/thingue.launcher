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
	"thingue-launcher/common/message"
	"thingue-launcher/common/model"
	"thingue-launcher/common/util"
	"time"
)

type Runner struct {
	*model.ClientInstance
	ExitSignalChannel chan error `json:"-"`
	process           *os.Process
	faultCount        uint
}

func (r *Runner) Start() error {
	if r.process != nil {
		return errors.New("实例已在运行")
	}
	// 设置PixelStreamingURL
	var launchArguments []string
	sid, err := NodeService.GetInstanceSid(global.NODE_ID, r.CID)
	if err == nil {
		r.SID = sid
		wsUrl := util.HttpUrlToWsUrl(config.AppConfig.ServerUrl, "/ws/streamer")
		launchArguments = append(r.LaunchArguments, "-PixelStreamingURL="+wsUrl+"/"+r.SID)
	} else {
		return errors.New("服务未连接")
	}
	// 设置日志文件名称为实例名称
	if r.Name != "" {
		launchArguments = append(r.LaunchArguments, "LOG="+r.Name+".log")
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
	RunnerManager.RunnerStatusUpdateChanel <- r.CID
	go func() {
		exitCode := command.Wait()
		r.Pid = 0
		r.process = nil
		r.LastStopAt = time.Now()
		r.StreamerConnected = false
		select {
		case r.ExitSignalChannel <- exitCode:
			r.updateStateCode(0)
			RunnerManager.RunnerStatusUpdateChanel <- r.CID
			fmt.Println("退出码发送成功")
			r.faultCount = 0
		default:
			r.updateStateCode(-1)
			RunnerManager.RunnerUnexpectedExitChanel <- r.CID
			fmt.Println("异常退出", r.faultCount)
			if r.FaultRecover && r.faultCount < 3 {
				time.Sleep(3 * time.Second)
				r.Start()
			}
			r.faultCount++
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
	NodeService.SendProcessState(&message.NodeProcessStateUpdate{
		SID:       r.SID,
		StateCode: stateCode,
	})
}
