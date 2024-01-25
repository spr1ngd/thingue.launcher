package conn

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"thingue-launcher/client/core/instance"
	"thingue-launcher/client/global"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	types "thingue-launcher/common/gen/proto/go/types/v1"
)

type AgentService struct {
	pb.UnimplementedAgentServiceServer
}

func (s AgentService) GetAgentInfo(context context.Context, request *pb.GetAgentInfoRequest) (*pb.GetAgentInfoResponse, error) {
	global.ClientId = uint(request.ClientId)
	var instanceInfos []*types.InstanceInfo
	instances := instance.RunnerManager.List()
	for _, item := range instances {
		instanceInfo := &types.InstanceInfo{
			Id:                uint32(item.CID),
			Pid:               int32(item.Pid),
			StateCode:         int32(item.StateCode),
			StreamerConnected: item.StreamerConnected,
			StreamerId:        item.SID,
			LastStartAt:       timestamppb.New(item.LastStartAt),
			LastStopAt:        timestamppb.New(item.LastStopAt),
			InstanceConfig: &types.InstanceConfig{
				Name:                   item.Name,
				CloudRes:               item.CloudRes,
				ExecPath:               item.ExecPath,
				LaunchArguments:        item.LaunchArguments,
				Metadata:               item.Metadata,
				PaksConfig:             item.PaksConfig,
				FaultRecover:           item.FaultRecover,
				EnableRelay:            item.EnableRelay,
				EnableRenderControl:    item.EnableRenderControl,
				EnableMultiuserControl: item.EnableMultiuserControl,
				AutoControl:            item.AutoControl,
				StopDelay:              int32(item.StopDelay),
			},
			PlayerConfig: &types.PlayerConfig{
				MatchViewportRes: item.PlayerConfig.MatchViewportRes,
				HideUi:           item.PlayerConfig.HideUI,
				IdleDisconnect:   item.PlayerConfig.IdleDisconnect,
				IdleTimeout:      uint32(item.PlayerConfig.IdleTimeout),
			},
		}
		instanceInfos = append(instanceInfos, instanceInfo)
	}
	response := &pb.GetAgentInfoResponse{
		DeviceInfo: GetDeviceInfo(),
		Instances:  instanceInfos,
	}
	return response, nil
}

func (s AgentService) ControlProcess(ctx context.Context, request *pb.ControlProcessRequest) (*emptypb.Empty, error) {
	runner, err := instance.RunnerManager.GetRunnerById(uint(request.InstanceId))
	if err == nil {
		err = runner.Start()
	}
	return &emptypb.Empty{}, err
}

func (s AgentService) UpdateStreamerState(ctx context.Context, request *pb.UpdateStreamerStateRequest) (*emptypb.Empty, error) {
	cid := uint(request.InstanceId)
	runner, err := instance.RunnerManager.GetRunnerById(cid)
	if err == nil {
		runner.StreamerConnected = request.StreamerState
		instance.RunnerManager.RunnerStatusUpdateChanel <- cid
	}
	return &emptypb.Empty{}, err
}
