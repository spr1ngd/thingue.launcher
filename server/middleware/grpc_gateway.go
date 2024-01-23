package middleware

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/common/provider"
)

func createGrpcGatewayHandler() http.Handler {
	endpoint := provider.AppConfig.LocalServer.BindAddr
	// 创建新的 ServeMux
	mux := runtime.NewServeMux()

	// 连接 gRPC 服务器
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// 注册服务
	err := pb.RegisterInstanceServiceHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC gateway: %v", err)
	}
	prefixMux := http.NewServeMux()
	contentPath := provider.AppConfig.LocalServer.ContentPath
	prefix := contentPath + "grpc-gateway" // 定义前缀
	prefixMux.Handle(prefix+"/", http.StripPrefix(prefix, mux))
	return prefixMux
}
