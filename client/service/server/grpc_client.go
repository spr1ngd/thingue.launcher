package server

import pb "thingue-launcher/common/gen/proto/go/apis/v1"

type grpcClient struct {
	InstanceService pb.InstanceServiceClient
}

var GrpcClient = grpcClient{}
