package provider

import (
	"errors"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/common/logger"
)

type grpcClientProvider struct {
	connMap map[uint32]pb.AgentServiceClient
}

var GrpcClientProvider = grpcClientProvider{connMap: make(map[uint32]pb.AgentServiceClient)}

func (g *grpcClientProvider) GetClient(clientId uint32) (pb.AgentServiceClient, error) {
	client, ok := g.connMap[clientId]
	if ok {
		return client, nil
	} else {
		logger.Zap.Errorf("找不到客户端,id=%d", clientId)
		return nil, errors.New("找不到客户端")
	}
}

func (g *grpcClientProvider) AddClient(clientId uint32, client pb.AgentServiceClient) {
	g.connMap[clientId] = client
}

func (g *grpcClientProvider) RemoveClient(clientId uint32) {
	delete(g.connMap, clientId)
}
