package rpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
)

type InstanceService struct {
	pb.UnimplementedInstanceServiceServer
}

func (s InstanceService) RegisterAgent(context.Context, *pb.RegisterAgentRequest) (*pb.RegisterAgentResponse, error) {
	return &pb.RegisterAgentResponse{Id: 0}, nil
}

func (s InstanceService) AddInstance(context.Context, *pb.AddInstanceRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s InstanceService) DeleteInstance(context.Context, *pb.DeleteInstanceRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s InstanceService) UpdateConfig(context.Context, *pb.UpdateConfigRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s InstanceService) GetStreamerId(context.Context, *pb.GetStreamerIdRequest) (*pb.GetStreamerIdResponse, error) {
	return &pb.GetStreamerIdResponse{Id: ""}, nil
}

func (s InstanceService) UpdateProcessState(context.Context, *pb.UpdateProcessStateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s InstanceService) UpdateRestarting(context.Context, *pb.UpdateRestartingRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s InstanceService) ClearPakState(context.Context, *pb.ClearPakStateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
