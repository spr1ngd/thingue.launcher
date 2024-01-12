package instance

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"os"
	"path/filepath"
	"thingue-launcher/common/domain"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/model"
	"time"
)

type runnerManager struct {
	IdRunnerMap                map[uint]*Runner
	RunnerUnexpectedExitChanel chan uint
	RunnerStatusUpdateChanel   chan uint
	HaveInternalInstance       bool
	IsInternalInstanceStarted  bool
}

var RunnerManager = &runnerManager{
	IdRunnerMap:                make(map[uint]*Runner),
	RunnerUnexpectedExitChanel: make(chan uint),
	RunnerStatusUpdateChanel:   make(chan uint),
}

func (m *runnerManager) Init() {
	// 从当前位置发现内置Runner
	pwd, _ := os.Getwd()
	files, _ := os.ReadDir(pwd)
	for _, entry := range files {
		if !entry.IsDir() {
			if "ThingUE.exe" == entry.Name() || "ThingUE.sh" == entry.Name() {
				m.HaveInternalInstance = true
				instance, err := InstanceManager.GetInternal()
				if err != nil {
					instance = &model.ClientInstance{
						Name:            "ThingUE",
						ExecPath:        filepath.Join(pwd, entry.Name()),
						LaunchArguments: []string{"-AudioMixer", "-RenderOffScreen", "-ForceRes", "-ResX=1920", "-ResY=1080"},
						IsInternal:      true,
						EnableRelay:     true,
						PlayerConfig: domain.PlayerConfig{
							MatchViewportRes: true,
							HideUI:           false,
							IdleDisconnect:   false,
							IdleTimeout:      5,
						},
					}
					InstanceManager.Create(instance)
					_ = m.NewRunner(instance)
				} else {
					instance.ExecPath = filepath.Join(pwd, entry.Name())
					_ = m.NewRunner(instance)
					_ = InstanceManager.SaveConfig(instance)
				}
			}
		}
	}

	// 从持久化数据中实例化非内置Runners
	instances := InstanceManager.List()
	for index := range instances {
		if !instances[index].IsInternal {
			_ = m.NewRunner(&instances[index])
		}
	}
}

func (m *runnerManager) List() []*domain.Instance {
	var instances = make([]*domain.Instance, 0)
	for _, instance := range InstanceManager.List() {
		runner := m.GetRunnerById(instance.CID)
		if runner != nil {
			if !m.HaveInternalInstance && runner.IsInternal { //处理内置
				continue
			}
			//todo 排查这里为什么会NullPointerException
			instances = append(instances, runner.Instance)
		}
	}
	return instances
}

func (m *runnerManager) NewRunner(clientInstance *model.ClientInstance) error {
	if _, ok := m.IdRunnerMap[clientInstance.CID]; ok {
		return errors.New("无法重复创建")
	}
	var instance = &domain.Instance{}
	mapstructure.Decode(clientInstance, instance)
	r := &Runner{
		Instance:          instance,
		ExitSignalChannel: make(chan error, 1),
	}
	m.IdRunnerMap[r.CID] = r
	return nil
}

func (m *runnerManager) GetRunnerById(id uint) *Runner {
	value, ok := m.IdRunnerMap[id]
	if ok {
		return value
	} else {
		return nil
	}
}

func (m *runnerManager) CloseAllRunner() {
	logger.Zap.Info("关闭所有正在运行的实例")
	for _, runner := range m.IdRunnerMap {
		if runner.StateCode == 1 {
			runner.Stop()
		}
	}
}

func (m *runnerManager) ExecCommand(id uint, command string) {
	runner := m.GetRunnerById(id)
	if runner != nil {
		if command == "START" {
			_ = runner.Start()
		} else if command == "STOP" {
			_ = runner.Stop()
		}
	}
}

func (m *runnerManager) RestartAllRunner() {
	for _, runner := range m.IdRunnerMap {
		if runner.StateCode == 1 {
			logger.Zap.Infof("执行重启任务 %s", runner.Name)
			_ = runner.Stop()
			ClientService.SetRestarting(runner.SID, true)
			time.Sleep(3 * time.Second) //kill发出停顿三秒，等待进程关闭
			err := runner.Start()
			if err != nil {
				logger.Zap.Errorf("重启失败 %s %s", runner.Name, err)
				ClientService.SetRestarting(runner.SID, false)
			}
		}
	}
}

func (m *runnerManager) DeleteRunner(id uint) error {
	runner := m.GetRunnerById(id)
	if runner != nil {
		if runner.IsInternal {
			return errors.New("自动配置实例无法删除")
		}
		if runner.StateCode == 1 {
			return errors.New("实例正在运行，无法删除")
		}
		delete(m.IdRunnerMap, runner.CID)
	} else {
		return errors.New("找不到实例")
	}
	return nil
}

func (m *runnerManager) StartInternalRunner() {
	if m.HaveInternalInstance && !m.IsInternalInstanceStarted {
		internal, err := InstanceManager.GetInternal()
		if err == nil {
			runner := m.GetRunnerById(internal.CID)
			if runner != nil {
				runner.Start()
				m.IsInternalInstanceStarted = true
			}
		}
	}
}
