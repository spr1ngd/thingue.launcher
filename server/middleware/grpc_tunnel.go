package middleware

import (
	"fmt"
	"github.com/jhump/grpctunnel"
	"log"
)

func CreateGrpcTunnelServiceHandler() *grpctunnel.TunnelServiceHandler {
	return grpctunnel.NewTunnelServiceHandler(
		grpctunnel.TunnelServiceHandlerOptions{
			OnReverseTunnelOpen: func(channel grpctunnel.TunnelChannel) {
				log.Printf("New Tunnel Opened%p\n", &channel)
				channel.Done()
				// 客户端池

				<-channel.Context().Done()
				fmt.Printf("quxiaoguanbi")
			},
			OnReverseTunnelClose: func(channel grpctunnel.TunnelChannel) {
				log.Printf("Tunnel Closed%p\n", channel)
			},
		},
	)
}
