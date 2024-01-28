package service

import (
	"errors"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/message"
	"thingue-launcher/common/request"
	"thingue-launcher/server/core/provider"
	"time"
)

type sdpService struct{}

var SdpService = sdpService{}

func (m *sdpService) OnStreamerConnect(streamer *provider.StreamerConnector) {
	InstanceService.UpdateStreamerConnected(streamer.StreamerId, true)
	// 开启自动停止任务
	go func() {
		for {
			// todo 释放资源
			<-streamer.AutoStopTimer.C
			getStreamer, err := provider.SdpConnProvider.GetStreamer(streamer.StreamerId)
			if err == nil {
				if streamer != getStreamer {
					logger.Zap.Warn("streamer已重启，自动停止任务关闭")
					break
				}
				if len(getStreamer.PlayerConnectors) == 0 {
					InstanceService.ProcessControl(request.ProcessControl{
						StreamerId: getStreamer.StreamerId,
						Command:    "STOP",
					})
					logger.Zap.Info("检查完毕，自动停止控制指令发送")
				} else {
					logger.Zap.Debug("检查完毕，不需要自动停止")
				}
			} else {
				logger.Zap.Info("streamer已停止，自动停止任务关闭")
				break
			}
		}
		logger.Zap.Info("自动停止协程结束，资源释放")
	}()
}

func (m *sdpService) OnStreamerDisconnect(streamer *provider.StreamerConnector) {
	for _, playerConnector := range streamer.PlayerConnectors {
		playerConnector.Close()
	}
	InstanceService.UpdateStreamerConnected(streamer.StreamerId, false)
	streamer.Close()
}

func (m *sdpService) ConnectStreamer(playerConnector *provider.PlayerConnector, ticket string) error {
	if ticket == "test" {
		streamer, err := provider.SdpConnProvider.GetStreamer("test")
		if err == nil {
			playerConnector.StreamerConnector = streamer
		}
		return err
	}
	streamerId, err := TicketService.GetStreamerByTicket(ticket)
	if err == nil {
		streamer, err := provider.SdpConnProvider.GetStreamer(streamerId)
		if err == nil {
			playerConnector.StreamerConnector = streamer
		} else {
			instance := InstanceService.GetInstanceByStreamerId(streamerId)
			if instance.AutoControl {
				InstanceService.ProcessControl(request.ProcessControl{
					StreamerId: streamerId,
					Command:    "START",
				})
				ticker := time.NewTicker(2 * time.Second)
				for {
					<-ticker.C
					streamer, err := provider.SdpConnProvider.GetStreamer(streamerId)
					if err == nil {
						playerConnector.StreamerConnector = streamer
						ticker.Stop()
						break
					}
				}
				logger.Zap.Info("自动启动成功")
			} else {
				err = errors.New("streamer未连接且未开启自动启动")
			}
		}
	}
	return err
}

func (m *sdpService) OnStreamerLoadBundleComplete(streamer *provider.StreamerConnector) {
	if len(streamer.PlayerConnectors) == 0 {
		streamer.ControlRendering(false)
		InstanceService.UpdateRenderingState(streamer.StreamerId, false)
	}
}

func (m *sdpService) OnPlayerPaired(player *provider.PlayerConnector) {
	player.StreamerConnector.PlayerConnectors = append(player.StreamerConnector.PlayerConnectors, player)
	player.StreamerConnector.SendPlayersCount()
	InstanceService.UpdatePlayers(player.StreamerConnector)
	// 如果未开启渲染，则发消息开启
	if !player.StreamerConnector.RenderingState {
		player.StreamerConnector.ControlRendering(true)
		InstanceService.UpdateRenderingState(player.StreamerConnector.StreamerId, true)
	}
}

func (m *sdpService) OnPlayerDisConnect(player *provider.PlayerConnector) {
	player.StreamerConnector.PlayerDisconnect(player)
	player.StreamerConnector.SendPlayersCount()
	instance := InstanceService.UpdatePlayers(player.StreamerConnector)
	if len(player.StreamerConnector.PlayerConnectors) == 0 {
		// 关闭渲染
		player.StreamerConnector.ControlRendering(false)
		InstanceService.UpdateRenderingState(player.StreamerConnector.StreamerId, false)
		if instance.AutoControl && instance.StopDelay >= 0 {
			// 启动自动启停延时任务
			player.StreamerConnector.AutoStopTimer.Reset(time.Duration(instance.StopDelay) * time.Second)
		}
	}
	player.Close()
}

func (m *sdpService) KickPlayerUser(userQueryMap map[string]string) (int, error) {
	if len(userQueryMap) == 0 {
		return 0, errors.New("参数不能为空")
	}
	players := provider.SdpConnProvider.GetPlayersByUserData(userQueryMap)
	if len(players) > 0 {
		for _, player := range players {
			player.Kick()
		}
		return len(players), nil
	} else {
		return 0, errors.New("没有匹配的连接")
	}
}

func (m *sdpService) OnStreamerNodeRestarted(streamer *provider.StreamerConnector) {
	instance := InstanceService.GetInstanceByStreamerId(streamer.StreamerId)
	restarting := provider.SdpConnProvider.GetStreamerRestartingState(streamer.StreamerId)
	if restarting && instance.PakValue != "" {
		logger.Zap.Infof("重启后加载 %s %s", instance.Name, instance.PakValue)
		command := message.Command{}
		command.BuildBundleLoadCommand(message.BundleLoadParams{Bundle: instance.PakValue})
		streamer.SendCommand(&command)
		provider.SdpConnProvider.SetStreamerRestartingState(streamer.StreamerId, false)
	} else if restarting {
		provider.SdpConnProvider.SetStreamerRestartingState(streamer.StreamerId, false)
		logger.Zap.Infof("重启后不需要加载pak %s %s", instance.Name, instance.PakValue)
	} else {
		logger.Zap.Warnf("非重启时忽略nodeRestarted消息 %s", instance.Name)
	}
	if streamer.EnableRenderControl && len(streamer.PlayerConnectors) == 0 {
		command := message.Command{}
		command.BuildRenderingCommand(&message.RenderingParams{Value: false})
		streamer.SendCommand(&command)
	}
}
