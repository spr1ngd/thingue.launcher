package middleware

import (
	"context"
	"github.com/jhump/grpctunnel"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/model"
	"thingue-launcher/server/core"
	"thingue-launcher/server/core/provider"
)

func CreateGrpcTunnelServiceHandler() *grpctunnel.TunnelServiceHandler {
	return grpctunnel.NewTunnelServiceHandler(
		grpctunnel.TunnelServiceHandlerOptions{
			OnReverseTunnelOpen: func(channel grpctunnel.TunnelChannel) {
				//logger.Zap.Infof("New Tunnel Opened%p", &channel)
				channel.Done()
				// 客户端池
				client := &model.Client{}
				core.ClientService.CreateClient(client)
				agentClient := pb.NewAgentServiceClient(channel)
				response, err := agentClient.GetAgentInfo(context.Background(), &pb.GetAgentInfoRequest{
					ClientId: client.ID,
				})
				if err == nil {
					core.ClientService.RegisterClient(client, response)
					provider.GrpcClientProvider.ConnMap[client.ID] = agentClient
				} else {
					logger.Zap.Errorf("获取Agent信息失败 %s", err)
				}
				<-channel.Context().Done()
				core.ClientService.DeleteClient(client)
			},
			OnReverseTunnelClose: func(channel grpctunnel.TunnelChannel) {
				logger.Zap.Infof("Tunnel Closed %p", channel)
			},
		},
	)
}
