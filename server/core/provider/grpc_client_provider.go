package provider

import pb "thingue-launcher/common/gen/proto/go/apis/v1"

type grpcClientProvider struct {
	ConnMap map[uint]pb.AgentServiceClient
}

var GrpcClientProvider = grpcClientProvider{ConnMap: make(map[uint]pb.AgentServiceClient)}
