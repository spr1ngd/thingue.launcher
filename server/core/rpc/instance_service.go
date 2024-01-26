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

func (s InstanceService) RegisterAgent(context.Context, *pb.RegisterAgentRequest) (*pb.RegisterAgentResponse, error) {
	return &pb.RegisterAgentResponse{Id: 919}, nil
}

func (s InstanceService) AddInstance(ctx context.Context, req *pb.AddInstanceRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, core.InstanceService.AddInstance(req)
}

func (s InstanceService) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, core.InstanceService.DeleteInstance(req)
}

func (s InstanceService) UpdateConfig(context.Context, *pb.UpdateConfigRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s InstanceService) GetStreamerId(ctx context.Context, req *pb.GetStreamerIdRequest) (*pb.GetStreamerIdResponse, error) {
	streamerId, err := core.ClientService.GetInstanceStreamerId(req.ClientId, req.InstanceId)
	return &pb.GetStreamerIdResponse{Id: streamerId}, err
}

func (s InstanceService) UpdateProcessState(ctx context.Context, req *pb.UpdateProcessStateRequest) (*emptypb.Empty, error) {
	core.InstanceService.UpdateProcessState(req.ClientId, req.InstanceId, req.StateCode, req.Pid)
	return &emptypb.Empty{}, nil
}

func (s InstanceService) UpdateRestarting(context.Context, *pb.UpdateRestartingRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s InstanceService) ClearPakState(context.Context, *pb.ClearPakStateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
