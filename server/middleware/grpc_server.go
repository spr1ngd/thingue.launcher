package middleware

import (
	"github.com/jhump/grpctunnel/tunnelpb"
	"google.golang.org/grpc"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/server/core/rpc"
)

var grpcServer *grpc.Server

func CreateGRPCServer() {
	// 初始化 gRPC 服务器
	grpcServer = grpc.NewServer()
	// 注册反向隧道
	tunnelpb.RegisterTunnelServiceServer(grpcServer, CreateGrpcTunnelServiceHandler().Service())
	// 注册服务端gRPC服务
	pb.RegisterServerInstanceServiceServer(grpcServer, &rpc.InstanceService{})
}

func RunGRPCServer() error {
	// 将 gRPC 服务器绑定到自定义的 Listener 上
	return grpcServer.Serve(grpcL)
}

func StopGRPCServer() {
	if grpcServer != nil {
		grpcServer.Stop()
	}
}
