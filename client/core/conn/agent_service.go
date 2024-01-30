package conn

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"thingue-launcher/client/core/instance"
	"thingue-launcher/client/global"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
	types "thingue-launcher/common/gen/proto/go/types/v1"
	"thingue-launcher/common/util"
)

type AgentService struct {
	pb.UnimplementedAgentServiceServer
}

func (s AgentService) GetAgentInfo(ctx context.Context, request *pb.GetAgentInfoRequest) (*pb.GetAgentInfoResponse, error) {
	global.ClientId = request.ClientId
	var instanceInfos []*types.InstanceInfo
	for _, item := range instance.RunnerManager.List() {
		instanceInfos = append(instanceInfos, item.ToInstanceInfoTypes())
	}
	fmt.Printf("%+v", instanceInfos)
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

func (s AgentService) GetInstanceLogs(context.Context, *emptypb.Empty) (*pb.GetInstanceLogsResponse, error) {
	var filesToCompress []string
	instances := instance.RunnerManager.List()
	for _, item := range instances {
		filesToCompress = append(filesToCompress, instance.GetLogFiles(item)...)
	}

	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)
	for _, filePath := range filesToCompress {
		_ = util.AddFileToZip(zipWriter, filePath)
	}
	_ = zipWriter.Close()

	return &pb.GetInstanceLogsResponse{Data: buf.Bytes()}, nil
}
