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
	"thingue-launcher/client/service/server"
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
	streamerIdResponse, err := server.GrpcClient.InstanceService.GetStreamerId(context.Background(), &pb.GetStreamerIdRequest{
		ClientId:   uint32(global.ClientId),
		InstanceId: uint32(r.CID),
	})
	if err == nil {
		r.SID = streamerIdResponse.GetId()
		wsUrl := util.HttpUrlToWsUrl(provider.AppConfig.ServerURL, "/ws/streamer")
		launchArguments = append(r.LaunchArguments,
			fmt.Sprintf("-PixelStreamingURL=%s/%s", wsUrl, r.SID))
	} else {
		return err
	}
	// 设置日志文件名称启动参数
	if r.Name != "" {
		launchArguments = append(launchArguments, fmt.Sprintf("LOG=%s.log", r.Name))
	}
	// 运行前
	logger.Zap.Debug(r.ExecPath, launchArguments)
	command := exec.Command(r.ExecPath, launchArguments...)
	err = command.Start()
	if err != nil {
		return err
	}
	r.Pid = command.Process.Pid
	r.process = command.Process
	r.updateProcessState(1)
	r.LastStartAt = time.Now()
	RunnerManager.RunnerStatusUpdateChanel <- r.CID
	logger.Zap.Infof("实例启动 %s", r.Name)
	go func() {
		exitCode := command.Wait()
		r.Pid = 0
		r.process = nil
		r.LastStopAt = time.Now()
		r.StreamerConnected = false
		select {
		case r.ExitSignalChannel <- exitCode:
			r.updateProcessState(0)
			RunnerManager.RunnerStatusUpdateChanel <- r.CID
			logger.Zap.Debugf("退出码发送成功 %s", r.Name)
			r.faultCount = 0
		default:
			r.updateProcessState(-1)
			RunnerManager.RunnerUnexpectedExitChanel <- r.CID
			logger.Zap.Warnf("实例异常退出 %s %d", r.Name, r.faultCount)
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
		cmd = exec.Command("pkill", "-TERM", "-P", strconv.Itoa(r.Pid))
	} else {
		return errors.New("不支持的系统")
	}
	cmd.SysProcAttr = &windows.SysProcAttr{HideWindow: true}
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	exitStatus := <-r.ExitSignalChannel
	logger.Zap.Infof("实例停止 %s %s", r.Name, exitStatus)
	return err
}

func (r *Runner) updateProcessState(stateCode int8) {
	r.StateCode = stateCode
	_, err := server.GrpcClient.InstanceService.UpdateProcessState(context.Background(), &pb.UpdateProcessStateRequest{
		ClientId:   uint32(global.ClientId),
		InstanceId: uint32(r.CID),
		StateCode:  int32(stateCode),
		Pid:        int32(r.Pid),
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
