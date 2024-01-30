// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: types/v1/instance.proto

package typesv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InstanceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Pid               int32                  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	StateCode         int32                  `protobuf:"varint,3,opt,name=state_code,json=stateCode,proto3" json:"state_code,omitempty"`
	StreamerConnected bool                   `protobuf:"varint,4,opt,name=streamer_connected,json=streamerConnected,proto3" json:"streamer_connected,omitempty"`
	StreamerId        string                 `protobuf:"bytes,5,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`
	LastStartAt       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=last_start_at,json=lastStartAt,proto3" json:"last_start_at,omitempty"`
	LastStopAt        *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_stop_at,json=lastStopAt,proto3" json:"last_stop_at,omitempty"`
	Config            *InstanceConfig        `protobuf:"bytes,8,opt,name=config,proto3" json:"config,omitempty"`
	PlayerConfig      *PlayerConfig          `protobuf:"bytes,9,opt,name=player_config,json=playerConfig,proto3" json:"player_config,omitempty"`
}

func (x *InstanceInfo) Reset() {
	*x = InstanceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_instance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceInfo) ProtoMessage() {}

func (x *InstanceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_instance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceInfo.ProtoReflect.Descriptor instead.
func (*InstanceInfo) Descriptor() ([]byte, []int) {
	return file_types_v1_instance_proto_rawDescGZIP(), []int{0}
}

func (x *InstanceInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InstanceInfo) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *InstanceInfo) GetStateCode() int32 {
	if x != nil {
		return x.StateCode
	}
	return 0
}

func (x *InstanceInfo) GetStreamerConnected() bool {
	if x != nil {
		return x.StreamerConnected
	}
	return false
}

func (x *InstanceInfo) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *InstanceInfo) GetLastStartAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastStartAt
	}
	return nil
}

func (x *InstanceInfo) GetLastStopAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastStopAt
	}
	return nil
}

func (x *InstanceInfo) GetConfig() *InstanceConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *InstanceInfo) GetPlayerConfig() *PlayerConfig {
	if x != nil {
		return x.PlayerConfig
	}
	return nil
}

type InstanceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CloudRes               string   `protobuf:"bytes,2,opt,name=cloud_res,json=cloudRes,proto3" json:"cloud_res,omitempty"`
	ExecPath               string   `protobuf:"bytes,3,opt,name=exec_path,json=execPath,proto3" json:"exec_path,omitempty"`
	LaunchArguments        []string `protobuf:"bytes,4,rep,name=launch_arguments,json=launchArguments,proto3" json:"launch_arguments,omitempty"`
	Metadata               string   `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	PaksConfig             string   `protobuf:"bytes,6,opt,name=paks_config,json=paksConfig,proto3" json:"paks_config,omitempty"`
	FaultRecover           bool     `protobuf:"varint,7,opt,name=fault_recover,json=faultRecover,proto3" json:"fault_recover,omitempty"`
	EnableRelay            bool     `protobuf:"varint,8,opt,name=enable_relay,json=enableRelay,proto3" json:"enable_relay,omitempty"`
	EnableRenderControl    bool     `protobuf:"varint,9,opt,name=enable_render_control,json=enableRenderControl,proto3" json:"enable_render_control,omitempty"`
	EnableMultiuserControl bool     `protobuf:"varint,10,opt,name=enable_multiuser_control,json=enableMultiuserControl,proto3" json:"enable_multiuser_control,omitempty"`
	AutoControl            bool     `protobuf:"varint,11,opt,name=auto_control,json=autoControl,proto3" json:"auto_control,omitempty"`
	StopDelay              int32    `protobuf:"varint,12,opt,name=stop_delay,json=stopDelay,proto3" json:"stop_delay,omitempty"`
}

func (x *InstanceConfig) Reset() {
	*x = InstanceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_instance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceConfig) ProtoMessage() {}

func (x *InstanceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_instance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceConfig.ProtoReflect.Descriptor instead.
func (*InstanceConfig) Descriptor() ([]byte, []int) {
	return file_types_v1_instance_proto_rawDescGZIP(), []int{1}
}

func (x *InstanceConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InstanceConfig) GetCloudRes() string {
	if x != nil {
		return x.CloudRes
	}
	return ""
}

func (x *InstanceConfig) GetExecPath() string {
	if x != nil {
		return x.ExecPath
	}
	return ""
}

func (x *InstanceConfig) GetLaunchArguments() []string {
	if x != nil {
		return x.LaunchArguments
	}
	return nil
}

func (x *InstanceConfig) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *InstanceConfig) GetPaksConfig() string {
	if x != nil {
		return x.PaksConfig
	}
	return ""
}

func (x *InstanceConfig) GetFaultRecover() bool {
	if x != nil {
		return x.FaultRecover
	}
	return false
}

func (x *InstanceConfig) GetEnableRelay() bool {
	if x != nil {
		return x.EnableRelay
	}
	return false
}

func (x *InstanceConfig) GetEnableRenderControl() bool {
	if x != nil {
		return x.EnableRenderControl
	}
	return false
}

func (x *InstanceConfig) GetEnableMultiuserControl() bool {
	if x != nil {
		return x.EnableMultiuserControl
	}
	return false
}

func (x *InstanceConfig) GetAutoControl() bool {
	if x != nil {
		return x.AutoControl
	}
	return false
}

func (x *InstanceConfig) GetStopDelay() int32 {
	if x != nil {
		return x.StopDelay
	}
	return 0
}

type PlayerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchViewportRes bool   `protobuf:"varint,1,opt,name=match_viewport_res,json=matchViewportRes,proto3" json:"match_viewport_res,omitempty"`
	HideUi           bool   `protobuf:"varint,2,opt,name=hide_ui,json=hideUi,proto3" json:"hide_ui,omitempty"`
	IdleDisconnect   bool   `protobuf:"varint,3,opt,name=idle_disconnect,json=idleDisconnect,proto3" json:"idle_disconnect,omitempty"`
	IdleTimeout      uint32 `protobuf:"varint,4,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
}

func (x *PlayerConfig) Reset() {
	*x = PlayerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_instance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerConfig) ProtoMessage() {}

func (x *PlayerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_instance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerConfig.ProtoReflect.Descriptor instead.
func (*PlayerConfig) Descriptor() ([]byte, []int) {
	return file_types_v1_instance_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerConfig) GetMatchViewportRes() bool {
	if x != nil {
		return x.MatchViewportRes
	}
	return false
}

func (x *PlayerConfig) GetHideUi() bool {
	if x != nil {
		return x.HideUi
	}
	return false
}

func (x *PlayerConfig) GetIdleDisconnect() bool {
	if x != nil {
		return x.IdleDisconnect
	}
	return false
}

func (x *PlayerConfig) GetIdleTimeout() uint32 {
	if x != nil {
		return x.IdleTimeout
	}
	return 0
}

type DeviceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version    string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Workdir    string   `protobuf:"bytes,2,opt,name=workdir,proto3" json:"workdir,omitempty"`
	Hostname   string   `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Memory     string   `protobuf:"bytes,4,opt,name=memory,proto3" json:"memory,omitempty"`
	Cpus       []string `protobuf:"bytes,5,rep,name=cpus,proto3" json:"cpus,omitempty"`
	Gpus       []string `protobuf:"bytes,6,rep,name=gpus,proto3" json:"gpus,omitempty"`
	Ips        []string `protobuf:"bytes,7,rep,name=ips,proto3" json:"ips,omitempty"`
	OsArch     string   `protobuf:"bytes,8,opt,name=os_arch,json=osArch,proto3" json:"os_arch,omitempty"`
	OsType     string   `protobuf:"bytes,9,opt,name=os_type,json=osType,proto3" json:"os_type,omitempty"`
	SystemUser string   `protobuf:"bytes,10,opt,name=system_user,json=systemUser,proto3" json:"system_user,omitempty"`
}

func (x *DeviceInfo) Reset() {
	*x = DeviceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_instance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceInfo) ProtoMessage() {}

func (x *DeviceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_instance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceInfo.ProtoReflect.Descriptor instead.
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return file_types_v1_instance_proto_rawDescGZIP(), []int{3}
}

func (x *DeviceInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *DeviceInfo) GetWorkdir() string {
	if x != nil {
		return x.Workdir
	}
	return ""
}

func (x *DeviceInfo) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *DeviceInfo) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *DeviceInfo) GetCpus() []string {
	if x != nil {
		return x.Cpus
	}
	return nil
}

func (x *DeviceInfo) GetGpus() []string {
	if x != nil {
		return x.Gpus
	}
	return nil
}

func (x *DeviceInfo) GetIps() []string {
	if x != nil {
		return x.Ips
	}
	return nil
}

func (x *DeviceInfo) GetOsArch() string {
	if x != nil {
		return x.OsArch
	}
	return ""
}

func (x *DeviceInfo) GetOsType() string {
	if x != nil {
		return x.OsType
	}
	return ""
}

func (x *DeviceInfo) GetSystemUser() string {
	if x != nil {
		return x.SystemUser
	}
	return ""
}

var File_types_v1_instance_proto protoreflect.FileDescriptor

var file_types_v1_instance_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x03, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x11, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73,
	0x74, 0x6f, 0x70, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x74,
	0x6f, 0x70, 0x41, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0xbe, 0x03, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x63, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x29, 0x0a, 0x10, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x61,
	0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f,
	0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x61, 0x6b, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x61, 0x6b, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a, 0x0d,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x6c, 0x61, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x13, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x38, 0x0a, 0x18, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x64, 0x65,
	0x6c, 0x61, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x70, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x22, 0xa1, 0x01, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x76,
	0x69, 0x65, 0x77, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x10, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x56, 0x69, 0x65, 0x77, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x75, 0x69, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x65, 0x55, 0x69, 0x12, 0x27, 0x0a, 0x0f,
	0x69, 0x64, 0x6c, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x64, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x69, 0x64, 0x6c,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x81, 0x02, 0x0a, 0x0a, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x64, 0x69, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x70, 0x75, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x70, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x70, 0x75, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x67, 0x70, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x70, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x70, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x73, 0x5f,
	0x61, 0x72, 0x63, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x73, 0x41, 0x72,
	0x63, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x42, 0x95, 0x01, 0x0a,
	0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x75, 0x65, 0x2d, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x72,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x14, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_v1_instance_proto_rawDescOnce sync.Once
	file_types_v1_instance_proto_rawDescData = file_types_v1_instance_proto_rawDesc
)

func file_types_v1_instance_proto_rawDescGZIP() []byte {
	file_types_v1_instance_proto_rawDescOnce.Do(func() {
		file_types_v1_instance_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_v1_instance_proto_rawDescData)
	})
	return file_types_v1_instance_proto_rawDescData
}

var file_types_v1_instance_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_types_v1_instance_proto_goTypes = []interface{}{
	(*InstanceInfo)(nil),          // 0: types.v1.InstanceInfo
	(*InstanceConfig)(nil),        // 1: types.v1.InstanceConfig
	(*PlayerConfig)(nil),          // 2: types.v1.PlayerConfig
	(*DeviceInfo)(nil),            // 3: types.v1.DeviceInfo
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_types_v1_instance_proto_depIdxs = []int32{
	4, // 0: types.v1.InstanceInfo.last_start_at:type_name -> google.protobuf.Timestamp
	4, // 1: types.v1.InstanceInfo.last_stop_at:type_name -> google.protobuf.Timestamp
	1, // 2: types.v1.InstanceInfo.config:type_name -> types.v1.InstanceConfig
	2, // 3: types.v1.InstanceInfo.player_config:type_name -> types.v1.PlayerConfig
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_types_v1_instance_proto_init() }
func file_types_v1_instance_proto_init() {
	if File_types_v1_instance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_v1_instance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_types_v1_instance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_types_v1_instance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_types_v1_instance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_v1_instance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_v1_instance_proto_goTypes,
		DependencyIndexes: file_types_v1_instance_proto_depIdxs,
		MessageInfos:      file_types_v1_instance_proto_msgTypes,
	}.Build()
	File_types_v1_instance_proto = out.File
	file_types_v1_instance_proto_rawDesc = nil
	file_types_v1_instance_proto_goTypes = nil
	file_types_v1_instance_proto_depIdxs = nil
}
