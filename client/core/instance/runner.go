package instance

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"thingue-launcher/client/global"
	"thingue-launcher/common/domain"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/provider"
	"thingue-launcher/common/util"
	"time"
)

type Runner struct {
	*domain.Instance
	ExitSignalChannel chan error `json:"-"`
	process           *os.Process
	faultCount        uint
}

func (r *Runner) Start() error {
	if r.process != nil {
		return errors.New("实例已在运行")
	}
	var launchArguments []string
	// 设置PixelStreamingURL启动参数
	streamerIdResponse, err := global.GrpcClient.GetStreamerId(context.Background(), &pb.GetStreamerIdRequest{
		ClientId:   global.ClientId,
		InstanceId: r.ID,
	})
	if err == nil {
		r.StreamerId = streamerIdResponse.GetId()
		wsUrl := util.HttpUrlToWsUrl(provider.AppConfig.ServerURL, "/ws/streamer")
		launchArguments = append(r.InstanceConfig.LaunchArguments,
			fmt.Sprintf("-PixelStreamingURL=%s/%s", wsUrl, r.StreamerId))
	} else {
		return err
	}
	// 设置日志文件名称启动参数
	if r.InstanceConfig.Name != "" {
		launchArguments = append(launchArguments, fmt.Sprintf("LOG=%s.log", r.InstanceConfig.Name))
	}
	// 运行前
	logger.Zap.Debug(r.InstanceConfig.ExecPath, launchArguments)
	command := exec.Command(r.InstanceConfig.ExecPath, launchArguments...)
	err = command.Start()
	if err != nil {
		return err
	}
	r.Pid = int32(command.Process.Pid)
	r.process = command.Process
	r.updateProcessState(1)
	r.LastStartAt = time.Now()
	RunnerManager.RunnerStatusUpdateChanel <- r.ID
	logger.Zap.Infof("实例启动 %s", r.InstanceConfig.Name)
	go func() {
		exitCode := command.Wait()
		r.Pid = 0
		r.process = nil
		r.LastStopAt = time.Now()
		r.StreamerConnected = false
		select {
		case r.ExitSignalChannel <- exitCode:
			r.updateProcessState(0)
			RunnerManager.RunnerStatusUpdateChanel <- r.ID
			logger.Zap.Debugf("退出码发送成功 %s", r.InstanceConfig.Name)
			r.faultCount = 0
		default:
			r.updateProcessState(-1)
			RunnerManager.RunnerUnexpectedExitChanel <- r.ID
			logger.Zap.Warnf("实例异常退出 %s %d", r.InstanceConfig.Name, r.faultCount)
			if r.InstanceConfig.FaultRecover && r.faultCount < 3 {
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
		cmd = exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(int(r.Pid)))
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("pkill", "-TERM", "-P", strconv.Itoa(int(r.Pid)))
	} else {
		return errors.New("不支持的系统")
	}
	cmd.SysProcAttr = &windows.SysProcAttr{HideWindow: true}
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	exitStatus := <-r.ExitSignalChannel
	logger.Zap.Infof("实例停止 %s %s", r.InstanceConfig.Name, exitStatus)
	return err
}

func (r *Runner) updateProcessState(stateCode int32) {
	r.StateCode = stateCode
	_, err := global.GrpcClient.UpdateProcessState(context.Background(), &pb.UpdateProcessStateRequest{
		ClientId:   global.ClientId,
		InstanceId: r.ID,
		StateCode:  stateCode,
		Pid:        r.Pid,
	})
	if err != nil {
		logger.Zap.Error(err)
	}
}

func (r *Runner) OpenLog() error {
	file, err := getLogFile(r.Instance)
	if err == nil {
		var cmdName string
		if provider.AppConfig.SystemSettings.ExternalEditorPath == "" {
			cmdName = "code"
		} else {
			cmdName = provider.AppConfig.SystemSettings.ExternalEditorPath
		}
		cmd := exec.Command(cmdName, file)
		return cmd.Run()
	}
	return err
}
