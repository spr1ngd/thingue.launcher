package middleware

import (
	"github.com/jhump/grpctunnel"
	"log"
	"thingue-launcher/common/logger"
)

func CreateGrpcTunnelServiceHandler() *grpctunnel.TunnelServiceHandler {
	return grpctunnel.NewTunnelServiceHandler(
		grpctunnel.TunnelServiceHandlerOptions{
			OnReverseTunnelOpen: func(channel grpctunnel.TunnelChannel) {
				log.Printf("New Tunnel Opened%p\n", &channel)
				channel.Done()
				// 客户端池

				<-channel.Context().Done()
				logger.Zap.Info("完成")
			},
			OnReverseTunnelClose: func(channel grpctunnel.TunnelChannel) {
				logger.Zap.Info("Tunnel Closed%p\n", channel)
			},
		},
	)
}
