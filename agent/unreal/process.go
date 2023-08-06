package unreal

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"thingue-launcher/agent/model"
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
	return idMap[id]
}

func (p *Process) start() error {
	fmt.Println(p.LaunchArguments)
	command := exec.Command(p.ExecPath, p.LaunchArguments...)
	err := command.Start()
	p.Pid = command.Process.Pid
	p.process = command.Process
	return err
}

func (p *Process) stop() error {
	// 根据pid获取进程
	//process, err1 := os.FindProcess(p.Pid)
	//if err1 != nil {
	//	return err1
	//}
	// 杀死进程
	err := p.process.Signal(syscall.SIGKILL)
	if err != nil {
		delete(idMap, p.ID)
	}
	return err
}

func (p *Process) restart() error {
	stopErr := p.process.Signal(syscall.SIGKILL)
	if stopErr != nil {
		return stopErr
	}
	return p.start()
}
