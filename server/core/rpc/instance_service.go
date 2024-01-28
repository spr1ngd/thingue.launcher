package rpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/server/core"
)

type InstanceService struct {
	pb.UnimplementedServerInstanceServiceServer
}

func (s InstanceService) AddInstance(ctx context.Context, req *pb.AddInstanceRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, core.InstanceService.AddInstance(req)
}

func (s InstanceService) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, core.InstanceService.DeleteInstance(req)
}

func (s InstanceService) UpdateConfig(ctx context.Context, req *pb.UpdateConfigRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, core.InstanceService.UpdateInstanceConfig(req)
}

func (s InstanceService) GetStreamerId(ctx context.Context, req *pb.GetStreamerIdRequest) (*pb.GetStreamerIdResponse, error) {
	streamerId, err := core.ClientService.GetInstanceStreamerId(req.ClientId, req.InstanceId)
	return &pb.GetStreamerIdResponse{Id: streamerId}, err
}

func (s InstanceService) UpdateProcessState(ctx context.Context, req *pb.UpdateProcessStateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, core.InstanceService.UpdateProcessState(req)
}

func (s InstanceService) UpdateRestarting(ctx context.Context, req *pb.UpdateRestartingRequest) (*emptypb.Empty, error) {
	core.InstanceService.UpdateRestarting(req)
	return &emptypb.Empty{}, nil
}

func (s InstanceService) ClearPakState(ctx context.Context, req *pb.ClearPakStateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, core.InstanceService.ClearPakState(req)
}
