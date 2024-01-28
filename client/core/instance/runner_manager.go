package instance

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"thingue-launcher/client/global"
	"thingue-launcher/common/domain"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/common/logger"
	"time"
)

type runnerManager struct {
	IdRunnerMap                map[uint32]*Runner
	RunnerUnexpectedExitChanel chan uint32
	RunnerStatusUpdateChanel   chan uint32
	HaveInternalInstance       bool
	IsInternalInstanceStarted  bool
}

var RunnerManager = &runnerManager{
	IdRunnerMap:                make(map[uint32]*Runner),
	RunnerUnexpectedExitChanel: make(chan uint32),
	RunnerStatusUpdateChanel:   make(chan uint32),
}

func (m *runnerManager) Init() {
	// 从当前位置发现内置Runner
	pwd, _ := os.Getwd()
	files, _ := os.ReadDir(pwd)
	for _, entry := range files {
		if !entry.IsDir() {
			if "ThingUE.exe" == entry.Name() || "ThingUE.sh" == entry.Name() {
				m.HaveInternalInstance = true
				instance, err := ConfigManager.GetInternal()
				if err != nil {
					defaultInstance := ConfigManager.GetDefault()
					defaultInstance.Config.Name = "ThingUE"
					defaultInstance.Config.ExecPath = filepath.Join(pwd, entry.Name())
					defaultInstance.IsInternal = true
					id := ConfigManager.Create(defaultInstance)
					_ = m.NewRunner(id, defaultInstance.Config)
				} else {
					instance.Config.ExecPath = filepath.Join(pwd, entry.Name())
					_ = m.NewRunner(instance.ID, instance.Config)
					_ = ConfigManager.Update(instance)
				}
			}
		}
	}

	// 从持久化数据中实例化非内置Runners
	instances := ConfigManager.List()
	for _, item := range instances {
		if !item.IsInternal {
			var instance domain.Instance
			instance.FromInstanceConfig(item)
			err := m.NewRunner(instance.ID, instance.Config)
			if err != nil {
				logger.Zap.Error("从配置中初始化runner失败")
			}
		}
	}
}

func (m *runnerManager) List() []*domain.Instance {
	var instanceList = make([]*domain.Instance, 0)
	for _, item := range ConfigManager.List() {
		runner, err := m.GetRunnerById(item.ID)
		if err == nil {
			// 排除掉没有内置实例但数据中有内置实例配置的情况
			if !m.HaveInternalInstance && runner.IsInternal {
				continue
			}
			instanceList = append(instanceList, runner.Instance)
		}
	}
	return instanceList
}

func (m *runnerManager) NewRunner(id uint32, config domain.InstanceConfig) error {
	if _, ok := m.IdRunnerMap[id]; ok {
		return errors.New("无法重复创建")
	}
	r := &Runner{
		Instance: &domain.Instance{
			ID:     id,
			Config: config,
		},
		ExitSignalChannel: make(chan error, 1),
	}
	m.IdRunnerMap[id] = r
	return nil
}

func (m *runnerManager) GetRunnerById(id uint32) (*Runner, error) {
	value, ok := m.IdRunnerMap[id]
	if ok {
		return value, nil
	} else {
		return nil, errors.New("找不到实例")
	}
}

func (m *runnerManager) CloseAllRunner() {
	for _, runner := range m.IdRunnerMap {
		if runner.StateCode == 1 {
			_ = runner.Stop()
		}
	}
}

func (m *runnerManager) ExecCommand(id uint32, command string) {
	runner, err := m.GetRunnerById(id)
	if err == nil {
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
			logger.Zap.Infof("执行重启任务 %s", runner.Config.Name)
			err := runner.Stop()
			if err != nil {
				logger.Zap.Error(err)
				continue
			}
			_, err = global.GrpcClient.UpdateRestarting(context.Background(), &pb.UpdateRestartingRequest{
				StreamerId: runner.StreamerId,
				Restarting: true,
			})
			if err != nil {
				logger.Zap.Error(err)
				continue
			}
			time.Sleep(3 * time.Second) //kill发出停顿三秒，等待进程关闭
			err = runner.Start()
			if err != nil {
				logger.Zap.Errorf("重启失败 %s %s", runner.Config.Name, err)
				_, err = global.GrpcClient.UpdateRestarting(context.Background(), &pb.UpdateRestartingRequest{
					StreamerId: runner.StreamerId,
					Restarting: true,
				})
				if err != nil {
					logger.Zap.Error(err)
				}
			}
		}
	}
}

func (m *runnerManager) DeleteRunner(id uint32) error {
	runner, err := m.GetRunnerById(id)
	if err == nil {
		if runner.IsInternal {
			return errors.New("自动配置实例无法删除")
		}
		if runner.StateCode == 1 {
			return errors.New("实例正在运行，无法删除")
		}
		delete(m.IdRunnerMap, runner.ID)
	} else {
		return err
	}
	return nil
}

func (m *runnerManager) StartInternalRunner() {
	if m.HaveInternalInstance && !m.IsInternalInstanceStarted {
		internal, err := ConfigManager.GetInternal()
		if err == nil {
			runner, err := m.GetRunnerById(internal.ID)
			if err == nil {
				err := runner.Start()
				if err != nil {
					logger.Zap.Error("内置实例启动失败")
				}
				m.IsInternalInstanceStarted = true
			}
		}
	}
}
