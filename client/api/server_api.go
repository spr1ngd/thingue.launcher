package api

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"
	"thingue-launcher/client/global"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/domain"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/model"
	"thingue-launcher/common/provider"
	"thingue-launcher/server/controller"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/yaml.v3"
)

type serverApi struct {
	ctx context.Context
}

var ServerApi = new(serverApi)

func (s *serverApi) Init(ctx context.Context) {
	s.ctx = ctx
	if provider.AppConfig.LocalServer.AutoStart {
		err := s.LocalServerStart()
		if err != nil {
			logger.Zap.Error(err)
		}
	}
	// 监听localserver关闭
	go func() {
		for {
			closeErr := <-controller.Application.HttpCloseChanel
			runtime.EventsEmit(s.ctx, constants.LOCAL_SERVER_CLOSE, closeErr)
		}
	}()
}

func (s *serverApi) LocalServerStart() error {
	return controller.Application.Start()
}

func (s *serverApi) LocalServerShutdown() error {
	err := controller.Application.Stop()
	return err
}

func (s *serverApi) GetHttpServerStatus() bool {
	return controller.Application.HttpServerRunning
}

func (s *serverApi) GetGrpcServerStatus() bool {
	return controller.Application.GrpcServerRunning
}

func (s *serverApi) GetMqttServerStatus() bool {
	return false
}

func (s *serverApi) UpdateLocalServerConfig(localServerConfig provider.LocalServer) {
	provider.AppConfig.LocalServer = localServerConfig
	provider.WriteConfigToFile()
}

func (s *serverApi) UpdatePeerConnectionOptions(options string) error {
	err := yaml.Unmarshal([]byte(options), &domain.PeerConnectionOptions{})
	if err != nil {
		return errors.New("格式不正确")
	} else {
		provider.AppConfig.PeerConnectionOptions = options
		provider.WriteConfigToFile()
		return nil
	}
}

func (s *serverApi) ListRemoteServer() []model.RemoteServer {
	var list []model.RemoteServer
	global.AppDB.Find(&list)
	return list
}

func (s *serverApi) CreateRemoteServer(remoteServer model.RemoteServer) uint {
	global.AppDB.Create(&remoteServer)
	return remoteServer.ID
}

func (s *serverApi) SaveRemoteServer(remoteServer model.RemoteServer) uint {
	global.AppDB.Save(&remoteServer)
	return remoteServer.ID
}

func (s *serverApi) DeleteRemoteServer(id uint) {
	global.AppDB.Delete(&model.RemoteServer{}, id)
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
	ServerURL, err := url.Parse(provider.AppConfig.ServerURL)
	if err == nil {
		path := ServerURL.JoinPath("/static/player.html")
		runtime.BrowserOpenURL(s.ctx, fmt.Sprintf("%s?sid=%s", path.String(), sid))
	}
}

func (s *serverApi) GetConnectServerOptions() []string {
	var options []string
	if s.GetHttpServerStatus() && s.GetGrpcServerStatus() {
		serverUrl, err := s.GetLocalServerUrl()
		if err == nil {
			options = append(options, serverUrl.String())
		}
	}
	for _, remoteServer := range s.ListRemoteServer() {
		if remoteServer.Url != "" {
			options = append(options, remoteServer.Url)
		}
	}
	return options
}
