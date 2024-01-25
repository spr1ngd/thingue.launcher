package conn

import (
	"context"
	"github.com/jhump/grpctunnel"
	"github.com/jhump/grpctunnel/tunnelpb"
	"google.golang.org/grpc"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/common/logger"
)

type tunnelServer struct {
	tunnelServer *grpctunnel.ReverseTunnelServer
	IsConnected  bool
}

var TunnelServer = &tunnelServer{}

func (s *tunnelServer) CreateTunnelServer(cc grpc.ClientConnInterface) {
	// Register services for reverse tunnels.
	tunnelStub := tunnelpb.NewTunnelServiceClient(cc)
	s.tunnelServer = grpctunnel.NewReverseTunnelServer(tunnelStub)
	pb.RegisterAgentServiceServer(s.tunnelServer, &AgentService{})
}

func (s *tunnelServer) ServeTunnelServer(resultChan chan error, cc grpc.ClientConnInterface) {
	s.CreateTunnelServer(cc)
	// Open the reverse tunnel and serve requests.
	s.IsConnected = true
	started, err := s.tunnelServer.Serve(context.Background()) //启动成功会阻塞
	if started {
		if ConnManager.grpcTarget != "" {
			logger.Zap.Info("grpc反向隧道启动后关闭,自动重连")
			ConnManager.StartConnectTask()
		} else {
			logger.Zap.Info("grpc反向隧道启动后关闭")
		}
	} else {
		logger.Zap.Errorf("grpc反向隧道启动失败 %s", err)
		resultChan <- err
	}
	s.IsConnected = false
}

func (s *tunnelServer) CloseTunnelServer() {
	s.tunnelServer.GracefulStop()
}
