package api

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"thingue-launcher/client/core"
	"thingue-launcher/client/core/conn"
	"thingue-launcher/client/core/instance"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/provider"
)

type connApi struct {
	ctx context.Context
}

var ConnApi = new(connApi)

func (c *connApi) Init(ctx context.Context) {
	c.ctx = ctx
	core.ConnManager.Init(c.ctx)
	//监听连接状态变化
	go func() {
		for {
			state := <-conn.TunnelServer.StateUpdateChanel
			url := provider.AppConfig.ServerURL
			runtime.EventsEmit(c.ctx, constants.CONN_STATE_UPDATE, state, url)
		}
	}()
}

func (c *connApi) ConnectServer(httpAddr string) error {
	if conn.TunnelServer.IsConnected {
		core.ConnManager.Close()
		conn.TunnelServer.Wg.Wait()
		logger.Zap.Debug("等待关闭完成")
	}
	err := core.ConnManager.SetConnAddr(httpAddr)
	if err == nil {
		core.ConnManager.StartConnectTask()
	}
	return err
}

func (c *connApi) DisconnectServer() {
	// 关闭已启动实例
	instance.RunnerManager.CloseAllRunner()
	core.ConnManager.Close()
}

func (c *connApi) GetConnState() map[string]any {
	return map[string]any{
		"isConnected": conn.TunnelServer.IsConnected,
		"serverAddr":  provider.AppConfig.ServerURL,
	}
}
