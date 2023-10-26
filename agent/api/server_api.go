package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"net"
	"net/url"
	"strings"
	"thingue-launcher/agent/global"
	"thingue-launcher/agent/service"
	"thingue-launcher/agent/service/instance"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/model"
	"thingue-launcher/common/provider"
	"thingue-launcher/server/initialize"
)

type serverApi struct {
	ctx context.Context
}

var ServerApi = new(serverApi)

func (s *serverApi) Init(ctx context.Context) {
	s.ctx = ctx
	if provider.AppConfig.LocalServer.AutoStart {
		s.LocalServerStart()
	}
	service.ServerConnManager.Init()
	// 监听server连接状态
	go func() {
		for {
			wsUrl := <-service.ServerConnManager.ServerConnUpdateChanel
			runtime.EventsEmit(s.ctx, constants.REMOTE_SERVER_CONN_UPDATE, wsUrl)
		}
	}()
	// 监听localserver关闭
	initialize.Server.CloseReturnChanel = make(chan string)
	go func() {
		for {
			closeErr := <-initialize.Server.CloseReturnChanel
			runtime.EventsEmit(s.ctx, constants.LOCAL_SERVER_CLOSE, closeErr)
		}
	}()
}

func (s *serverApi) LocalServerStart() {
	initialize.Server.Start()
}

func (s *serverApi) LocalServerShutdown() {
	initialize.Server.Stop()
}

func (s *serverApi) GetLocalServerStatus() bool {
	return initialize.Server.IsRunning
}

func (s *serverApi) UpdateLocalServerConfig(localServerConfig provider.LocalServer) {
	appConfig := provider.AppConfig
	appConfig.LocalServer = localServerConfig
	provider.WriteConfigToFile()
}

func (s *serverApi) ListRemoteServer() []model.RemoteServer {
	var list []model.RemoteServer
	global.APP_DB.Find(&list)
	return list
}

func (s *serverApi) CreateRemoteServer(remoteServer model.RemoteServer) uint {
	global.APP_DB.Create(&remoteServer)
	return remoteServer.ID
}

func (s *serverApi) SaveRemoteServer(remoteServer model.RemoteServer) uint {
	global.APP_DB.Save(&remoteServer)
	return remoteServer.ID
}

func (s *serverApi) DeleteRemoteServer(id uint) {
	global.APP_DB.Delete(&model.RemoteServer{}, id)
}

func (s *serverApi) GetConnectServerOptions() []string {
	var options []string
	if s.GetLocalServerStatus() {
		serverUrl, err := s.GetLocalServerUrl()
		if err == nil {
			options = append(options, serverUrl.String())
		}
	}
	for _, remoteServer := range s.ListRemoteServer() {
		options = append(options, remoteServer.Url)
	}
	return options
}

func (s *serverApi) ConnectServer(httpUrl string) error {
	if service.ServerConnManager.IsConnected {
		service.ServerConnManager.Disconnect()
	}
	err := service.ServerConnManager.SetServerAddr(httpUrl)
	if err == nil {
		service.ServerConnManager.StartConnectTask()
	}
	return err
}

func (s *serverApi) DisconnectServer() {
	// 关闭已启动实例
	instance.RunnerManager.CloseAllRunner()
	service.ServerConnManager.Disconnect()
}

func (s *serverApi) GetLocalServerUrl() (*url.URL, error) {
	split := strings.Split(provider.AppConfig.LocalServer.BindAddr, ":")
	if split[1] == "" {
		return nil, errors.New("地址不正确")
	}
	var baseUrl string
	ip := net.ParseIP(split[0])
	if ip != nil && ip.String() != "0.0.0.0" {
		baseUrl = fmt.Sprintf("http://%s:%s", ip, split[1])
	} else {
		baseUrl = fmt.Sprintf("http://localhost:%s", split[1])
	}
	parse, err := url.Parse(baseUrl)
	if err != nil {
		return nil, errors.New("地址不正确")
	}
	return parse.JoinPath(provider.AppConfig.LocalServer.ContentPath), nil
}

func (s *serverApi) OpenLocalServerUrl() {
	localServerUrl, err := s.GetLocalServerUrl()
	if err == nil {
		path := localServerUrl.JoinPath("/static")
		runtime.BrowserOpenURL(s.ctx, path.String())
	}
}

func (s *serverApi) OpenInstancePreviewUrl(sid string) {
	serverUrl, err := service.ServerConnManager.GetConnectedUrl()
	if err == nil {
		path := serverUrl.JoinPath("/static/player.html")
		runtime.BrowserOpenURL(s.ctx, fmt.Sprintf("%s?sid=%s", path.String(), sid))
	}
}

func (s *serverApi) GetServerConnInfo() map[string]interface{} {
	return map[string]interface{}{
		"isConnected": service.ServerConnManager.IsConnected,
		"serverAddr":  service.ServerConnManager.ServerAddr,
	}
}
