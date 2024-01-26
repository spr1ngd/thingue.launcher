package conn

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/protobuf/types/known/emptypb"
	"thingue-launcher/client/core/instance"
	"thingue-launcher/client/global"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	types "thingue-launcher/common/gen/proto/go/types/v1"
)

type AgentService struct {
	pb.UnimplementedAgentServiceServer
}

func (s AgentService) GetAgentInfo(ctx context.Context, request *pb.GetAgentInfoRequest) (*pb.GetAgentInfoResponse, error) {
	global.ClientId = request.ClientId
	var instanceInfos []*types.InstanceInfo
	instances := instance.RunnerManager.List()
	_ = mapstructure.Decode(instances, instanceInfos)
	response := &pb.GetAgentInfoResponse{
		DeviceInfo: GetDeviceInfo(),
		Instances:  instanceInfos,
	}
	return response, nil
}

func (s AgentService) ControlProcess(ctx context.Context, request *pb.ControlProcessRequest) (*emptypb.Empty, error) {
	runner, err := instance.RunnerManager.GetRunnerById(request.InstanceId)
	if err == nil {
		if request.Command == types.Command_COMMAND_START {
			err = runner.Start()
		} else if request.Command == types.Command_COMMAND_STOP {
			err = runner.Stop()
		}
	}
	return &emptypb.Empty{}, err
}

func (s AgentService) UpdateStreamerState(ctx context.Context, request *pb.UpdateStreamerStateRequest) (*emptypb.Empty, error) {
	runner, err := instance.RunnerManager.GetRunnerById(request.InstanceId)
	if err == nil {
		runner.StreamerConnected = request.StreamerState
		instance.RunnerManager.RunnerStatusUpdateChanel <- request.InstanceId
	}
	return &emptypb.Empty{}, err
}
