package unreal

import (
	"context"
	"errors"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"thingue-launcher/agent/model"
	"thingue-launcher/common/config"
)

var idRunnerMap = map[uint]*Runner{}

type Runner struct {
	*model.Instance
	exitChannel chan error
	process     *os.Process
	Pid         int
	IsRunning   bool
}

func NewRunner(instance *model.Instance) error {
	if _, ok := idRunnerMap[instance.ID]; ok {
		return errors.New("无法重复创建")
	}
	r := &Runner{
		Instance:    instance,
		exitChannel: make(chan error),
	}
	idRunnerMap[r.ID] = r
	return nil
}

func GetRunnerById(id uint) *Runner {
	value, ok := idRunnerMap[id]
	if ok {
		return value
	} else {
		return nil
	}
}

func CloseAllRunner() {
	fmt.Printf("关闭所有正在运行的实例")
	for _, runner := range idRunnerMap {
		if runner.IsRunning {
			runner.stop()
		}
	}
}

func RestartAllRunner(ctx context.Context) {
	for _, runner := range idRunnerMap {
		if runner.IsRunning {
			_ = runner.stop()
			//time.Sleep(3 * time.Second) //kill发出停顿三秒，等待进程关闭
			err := runner.start(ctx)
			if err != nil {
				fmt.Printf("%s重启失败:%s\n", runner.Name, err)
			} else {
				fmt.Printf("%s重启成功\n", runner.Name)
			}
		}
	}
}

func (r *Runner) start(ctx context.Context) error {
	if r.process != nil {
		return errors.New("实例已在运行")
	}
	// 设置PixelStreamingURL
	var launchArguments []string
	appConfig := config.GetAppConfig()
	if appConfig.ServerUrl != "" {
		httpUrl := appConfig.ServerUrl
		wsUrl := strings.Replace(httpUrl, "http://", "ws://", 1)
		wsUrl = strings.Replace(wsUrl, "https://", "wss://", 1)
		if strings.HasSuffix(wsUrl, "/") {
			launchArguments = append(r.LaunchArguments, "-PixelStreamingURL="+wsUrl+"ws/streamer/"+r.Name)
		} else {
			launchArguments = append(r.LaunchArguments, "-PixelStreamingURL="+wsUrl+"/ws/streamer/"+r.Name)
		}
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
	runtime.EventsEmit(ctx, "runner_status_update")
	go func() {
		exitCode := command.Wait()
		//运行后
		r.Pid = 0
		r.process = nil
		r.IsRunning = false
		runtime.EventsEmit(ctx, "runner_status_update")
		r.exitChannel <- exitCode
	}()
	return nil
}

func (r *Runner) stop() error {
	if !r.IsRunning {
		return errors.New("实例未在运行")
	}
	err := r.process.Signal(syscall.SIGKILL)
	exitStatus := <-r.exitChannel
	fmt.Printf("%s进程退出%s\n", r.Name, exitStatus)
	return err
}

func (r *Runner) delete() error {
	if r.IsRunning {
		return errors.New("实例正在运行，无法删除")
	}
	delete(idRunnerMap, r.ID)
	return nil
}
