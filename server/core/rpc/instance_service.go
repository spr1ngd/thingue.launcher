package rpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	"thingue-launcher/server/core"
)

type InstanceService struct {
	pb.UnimplementedServerInstanceServiceServer
}

func (s InstanceService) RegisterAgent(context.Context, *pb.RegisterAgentRequest) (*pb.RegisterAgentResponse, error) {
	return &pb.RegisterAgentResponse{Id: 919}, nil
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

func (s InstanceService) GetStreamerId(ctx context.Context, req *pb.GetStreamerIdRequest) (*pb.GetStreamerIdResponse, error) {
	sid, err := core.ClientService.GetInstanceSid(strconv.Itoa(int(req.ClientId)), strconv.Itoa(int(req.InstanceId)))
	return &pb.GetStreamerIdResponse{Id: sid}, err
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
