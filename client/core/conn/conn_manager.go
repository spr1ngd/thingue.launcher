package conn

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/url"
	"sync"
	"thingue-launcher/client/global"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/provider"
	"time"
)

type connManager struct {
	grpcTarget           string
	reconnectTimer       *time.Timer
	reconnectInterval    int
	maxReconnectInterval int
	connectLock          sync.Mutex
	cc                   *grpc.ClientConn
}

var ConnManager = &connManager{
	reconnectInterval: 5,
}

func (c *connManager) Init(ctx context.Context) {
	if provider.AppConfig.ServerURL != "" {
		err := c.SetConnAddr(provider.AppConfig.ServerURL)
		if err != nil {
			logger.Zap.Error(err)
		}
		c.StartConnectTask()
	}
}

func (c *connManager) SetConnAddr(httpAddr string) error {
	if TunnelServer.IsConnected {
		logger.Zap.Error("连接未断开")
		return errors.New("连接未断开")
	} else {
		httpURL, err := url.Parse(httpAddr)
		if err != nil {
			return err
		}
		c.grpcTarget = httpURL.Host
		provider.AppConfig.ServerURL = httpAddr
		provider.WriteConfigToFile()
		return nil
	}
}

func (c *connManager) UnsetConnAddr() {
	c.grpcTarget = ""
	provider.AppConfig.ServerURL = ""
	provider.WriteConfigToFile()
}

func (c *connManager) Close() {
	if c.cc != nil {
		c.UnsetConnAddr()
		err := c.cc.Close()
		TunnelServer.CloseTunnelServer()
		if err != nil {
			log.Error(err)
		}
	}
}

func (c *connManager) StartConnectTask() {
	logger.Zap.Infof("====grpc连接任务开始")
	c.reconnectTimer = time.NewTimer(time.Duration(0) * time.Second)
	go func() {
		for {
			<-c.reconnectTimer.C
			if c.grpcTarget == "" {
				logger.Zap.Info("grpc连接任务取消")
				break
			}
			err := c.connect()
			if err != nil {
				// 连接失败，重置计时器
				c.reconnectTimer.Reset(time.Duration(c.reconnectInterval) * time.Second)
			} else {
				// 连接成功，结束任务
				break
			}
		}
		logger.Zap.Info("====grpc连接任务结束")
	}()
}

func (c *connManager) connect() error {
	c.connectLock.Lock()
	defer c.connectLock.Unlock()
	if TunnelServer.IsConnected {
		return nil
	}
	cc, err := grpc.Dial(
		c.grpcTarget,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	} else {
		c.cc = cc
		// 创建客户端
		client := pb.NewServerInstanceServiceClient(cc)
		global.GrpcClient = client
		// 测试调用
		agent, err := client.RegisterAgent(context.Background(), &pb.RegisterAgentRequest{
			Instances:  nil,
			DeviceInfo: nil,
		})
		if err == nil {
			fmt.Println("grpc测试调用成功", agent.Id)
		}
		// 反向隧道
		resultChan := make(chan error)
		go TunnelServer.ServeTunnelServer(resultChan, cc)
		select {
		case err := <-resultChan:
			return err
		case <-time.After(5 * time.Second):
			// 稳定5秒返回
			logger.Zap.Info("grpc反向隧道连接稳定")
			return nil
		}
	}
}
