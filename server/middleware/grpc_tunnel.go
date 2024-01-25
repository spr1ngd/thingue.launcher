package middleware

import (
	"context"
	"fmt"
	"github.com/jhump/grpctunnel"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/common/logger"
)

func CreateGrpcTunnelServiceHandler() *grpctunnel.TunnelServiceHandler {
	return grpctunnel.NewTunnelServiceHandler(
		grpctunnel.TunnelServiceHandlerOptions{
			OnReverseTunnelOpen: func(channel grpctunnel.TunnelChannel) {
				logger.Zap.Infof("New Tunnel Opened%p", &channel)
				channel.Done()
				// 客户端池
				client := pb.NewAgentServiceClient(channel)
				response, err := client.GetAgentInfo(context.Background(), &pb.GetAgentInfoRequest{
					ClientId: 0,
				})
				fmt.Println("agent result", response)
				if err != nil {
					return
				}
				<-channel.Context().Done()
				logger.Zap.Info("完成")
			},
			OnReverseTunnelClose: func(channel grpctunnel.TunnelChannel) {
				logger.Zap.Infof("Tunnel Closed %p", channel)
			},
		},
	)
}
