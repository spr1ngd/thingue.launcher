package manager

import (
	"errors"
	"fmt"
	"thingue-launcher/agent/core"
	"thingue-launcher/agent/model"
)

type RunnerManager struct {
	IdRunnerMap map[uint]*core.Runner
}

func (m *RunnerManager) NewRunner(instance *model.Instance) error {
	if _, ok := m.IdRunnerMap[instance.ID]; ok {
		return errors.New("无法重复创建")
	}
	r := &core.Runner{
		Instance:          instance,
		ExitSignalChannel: make(chan error),
	}
	m.IdRunnerMap[r.ID] = r
	return nil
}

func (m *RunnerManager) GetRunnerById(id uint) *core.Runner {
	value, ok := m.IdRunnerMap[id]
	if ok {
		return value
	} else {
		return nil
	}
}

func (m *RunnerManager) CloseAllRunner() {
	fmt.Printf("关闭所有正在运行的实例")
	for _, runner := range m.IdRunnerMap {
		if runner.IsRunning {
			runner.Stop()
		}
	}
}

func (m *RunnerManager) RestartAllRunner() {
	for _, runner := range m.IdRunnerMap {
		if runner.IsRunning {
			_ = runner.Stop()
			//time.Sleep(3 * time.Second) //kill发出停顿三秒，等待进程关闭
			err := runner.Start()
			if err != nil {
				fmt.Printf("%s重启失败:%s\n", runner.Name, err)
			} else {
				fmt.Printf("%s重启成功\n", runner.Name)
			}
		}
	}
}

func (m *RunnerManager) DeleteRunner(id uint) error {
	runner := m.GetRunnerById(id)
	if runner != nil {
		if runner.IsRunning {
			return errors.New("实例正在运行，无法删除")
		}
		delete(m.IdRunnerMap, runner.ID)
	} else {
		return errors.New("找不到实例")
	}
	return nil
}
