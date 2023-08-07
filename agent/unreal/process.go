package unreal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"thingue-launcher/agent/model"
	"thingue-launcher/common/config"
)

var idMap = map[uint]*Process{}

type Process struct {
	*model.Instance
	process *os.Process
}

func NewProcess(instance *model.Instance) *Process {
	p := &Process{
		Instance: instance,
	}
	idMap[p.ID] = p
	return p
}

func GetProcessById(id uint) *Process {
	value, ok := idMap[id]
	if ok {
		return value
	} else {
		return nil
	}
}

func (p *Process) start() error {
	if p.process != nil {
		return errors.New("进程已存在")
	}
	// 设置PixelStreamingURL
	var launchArguments []string
	appConfig := config.GetAppConfig()
	if appConfig.ServerUrl != "" {
		if strings.HasSuffix(appConfig.ServerUrl, "/") {
			launchArguments = append(p.LaunchArguments, "-PixelStreamingURL="+appConfig.ServerUrl+"ws/streamer/"+p.Name)
		} else {
			launchArguments = append(p.LaunchArguments, "-PixelStreamingURL="+appConfig.ServerUrl+"/ws/streamer/"+p.Name)
		}
	}
	// 运行
	fmt.Println(p.ExecPath, launchArguments)
	command := exec.Command(p.ExecPath, launchArguments...)
	err := command.Start()
	command.Wait()
	if err == nil {
		p.Pid = command.Process.Pid
		p.process = command.Process
	}
	return err
}

func (p *Process) stop() error {
	if p.process == nil {
		return errors.New("进程不存在")
	}
	// 杀死进程
	err := p.process.Signal(syscall.SIGKILL)
	// 重置
	p.destroy()
	return err
}

func (p *Process) restart() error {
	stopErr := p.process.Signal(syscall.SIGKILL)
	if stopErr != nil {
		return stopErr
	}
	return p.start()
}

func (p *Process) destroy() {
	p.Pid = 0
	p.process = nil
	delete(idMap, p.ID)
}
