package service

import (
	"errors"
	"fmt"
	"thingue-launcher/common/request"
	"thingue-launcher/server/core/provider"
	"time"
)

type sdpService struct{}

var SdpService = sdpService{}

func (m *sdpService) OnStreamerConnect(streamer *provider.StreamerConnector) {
	InstanceService.UpdateStreamerConnected(streamer.SID, true)
	// 开启自动停止任务
	go func() {
		for {
			// todo 释放资源
			<-streamer.AutoStopTimer.C
			if len(streamer.PlayerConnectors) == 0 {
				InstanceService.ProcessControl(request.ProcessControl{
					SID:     streamer.SID,
					Command: "STOP",
				})
				fmt.Println("检查完毕，自动停止控制指令发送")
			} else {
				fmt.Println("检查完毕，不需要自动停止")
			}
		}
	}()
}

func (m *sdpService) OnStreamerDisconnect(streamer *provider.StreamerConnector) {
	for _, playerConnector := range streamer.PlayerConnectors {
		playerConnector.Close()
	}
	InstanceService.UpdateStreamerConnected(streamer.SID, false)
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
	sid, err := TicketService.GetSidByTicket(ticket)
	if err == nil {
		streamer, err := provider.SdpConnProvider.GetStreamer(sid)
		if err == nil {
			playerConnector.StreamerConnector = streamer
		} else {
			instance := InstanceService.GetInstanceBySid(sid)
			if instance.AutoControl {
				InstanceService.ProcessControl(request.ProcessControl{
					SID:     sid,
					Command: "START",
				})
				ticker := time.NewTicker(2 * time.Second)
				for {
					<-ticker.C
					streamer, err := provider.SdpConnProvider.GetStreamer(sid)
					if err == nil {
						playerConnector.StreamerConnector = streamer
						ticker.Stop()
						break
					}
				}
				fmt.Println("自动启动成功")
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
		InstanceService.UpdateRenderingState(streamer.SID, false)
	}
}

func (m *sdpService) OnPlayerPaired(player *provider.PlayerConnector) {
	player.StreamerConnector.PlayerConnectors = append(player.StreamerConnector.PlayerConnectors, player)
	InstanceService.UpdatePlayers(player.StreamerConnector)
	// 如果未开启渲染，则发消息开启
	if !player.StreamerConnector.RenderingState {
		player.StreamerConnector.ControlRendering(true)
		InstanceService.UpdateRenderingState(player.StreamerConnector.SID, true)
	}
}

func (m *sdpService) OnPlayerDisConnect(player *provider.PlayerConnector) {
	player.StreamerConnector.PlayerDisconnect(player)
	player.StreamerConnector.SendPlayersCount()
	instance := InstanceService.UpdatePlayers(player.StreamerConnector)
	if len(player.StreamerConnector.PlayerConnectors) == 0 {
		// 关闭渲染
		player.StreamerConnector.ControlRendering(false)
		InstanceService.UpdateRenderingState(player.StreamerConnector.SID, false)
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
