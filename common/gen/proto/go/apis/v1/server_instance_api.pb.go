// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: apis/v1/server_instance_api.proto

package apisv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	v1 "thingue-launcher/common/gen/proto/go/types/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     uint32           `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	InstanceInfo *v1.InstanceInfo `protobuf:"bytes,2,opt,name=instance_info,json=instanceInfo,proto3" json:"instance_info,omitempty"`
}

func (x *AddInstanceRequest) Reset() {
	*x = AddInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_server_instance_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInstanceRequest) ProtoMessage() {}

func (x *AddInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_server_instance_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInstanceRequest.ProtoReflect.Descriptor instead.
func (*AddInstanceRequest) Descriptor() ([]byte, []int) {
	return file_apis_v1_server_instance_api_proto_rawDescGZIP(), []int{0}
}

func (x *AddInstanceRequest) GetClientId() uint32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *AddInstanceRequest) GetInstanceInfo() *v1.InstanceInfo {
	if x != nil {
		return x.InstanceInfo
	}
	return nil
}

type DeleteInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId   uint32 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	InstanceId uint32 `protobuf:"varint,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
}

func (x *DeleteInstanceRequest) Reset() {
	*x = DeleteInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_server_instance_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInstanceRequest) ProtoMessage() {}

func (x *DeleteInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_server_instance_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInstanceRequest.ProtoReflect.Descriptor instead.
func (*DeleteInstanceRequest) Descriptor() ([]byte, []int) {
	return file_apis_v1_server_instance_api_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteInstanceRequest) GetClientId() uint32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *DeleteInstanceRequest) GetInstanceId() uint32 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

type GetStreamerIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId   uint32 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	InstanceId uint32 `protobuf:"varint,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
}

func (x *GetStreamerIdRequest) Reset() {
	*x = GetStreamerIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_server_instance_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamerIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamerIdRequest) ProtoMessage() {}

func (x *GetStreamerIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_server_instance_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamerIdRequest.ProtoReflect.Descriptor instead.
func (*GetStreamerIdRequest) Descriptor() ([]byte, []int) {
	return file_apis_v1_server_instance_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetStreamerIdRequest) GetClientId() uint32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *GetStreamerIdRequest) GetInstanceId() uint32 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

type GetStreamerIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetStreamerIdResponse) Reset() {
	*x = GetStreamerIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_server_instance_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamerIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamerIdResponse) ProtoMessage() {}

func (x *GetStreamerIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_server_instance_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamerIdResponse.ProtoReflect.Descriptor instead.
func (*GetStreamerIdResponse) Descriptor() ([]byte, []int) {
	return file_apis_v1_server_instance_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetStreamerIdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateRestartingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamerId string `protobuf:"bytes,1,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`
	Restarting bool   `protobuf:"varint,2,opt,name=restarting,proto3" json:"restarting,omitempty"`
}

func (x *UpdateRestartingRequest) Reset() {
	*x = UpdateRestartingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_server_instance_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRestartingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRestartingRequest) ProtoMessage() {}

func (x *UpdateRestartingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_server_instance_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRestartingRequest.ProtoReflect.Descriptor instead.
func (*UpdateRestartingRequest) Descriptor() ([]byte, []int) {
	return file_apis_v1_server_instance_api_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRestartingRequest) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *UpdateRestartingRequest) GetRestarting() bool {
	if x != nil {
		return x.Restarting
	}
	return false
}

type UpdateProcessStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId   uint32 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	InstanceId uint32 `protobuf:"varint,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	StateCode  int32  `protobuf:"varint,3,opt,name=state_code,json=stateCode,proto3" json:"state_code,omitempty"`
	Pid        int32  `protobuf:"varint,4,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *UpdateProcessStateRequest) Reset() {
	*x = UpdateProcessStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_server_instance_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProcessStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProcessStateRequest) ProtoMessage() {}

func (x *UpdateProcessStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_server_instance_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProcessStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateProcessStateRequest) Descriptor() ([]byte, []int) {
	return file_apis_v1_server_instance_api_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateProcessStateRequest) GetClientId() uint32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *UpdateProcessStateRequest) GetInstanceId() uint32 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

func (x *UpdateProcessStateRequest) GetStateCode() int32 {
	if x != nil {
		return x.StateCode
	}
	return 0
}

func (x *UpdateProcessStateRequest) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type UpdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId       uint32             `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	InstanceId     uint32             `protobuf:"varint,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	InstanceConfig *v1.InstanceConfig `protobuf:"bytes,3,opt,name=instance_config,json=instanceConfig,proto3" json:"instance_config,omitempty"`
	PlayerConfig   *v1.PlayerConfig   `protobuf:"bytes,4,opt,name=player_config,json=playerConfig,proto3" json:"player_config,omitempty"`
}

func (x *UpdateConfigRequest) Reset() {
	*x = UpdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_server_instance_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigRequest) ProtoMessage() {}

func (x *UpdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_server_instance_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_apis_v1_server_instance_api_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateConfigRequest) GetClientId() uint32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *UpdateConfigRequest) GetInstanceId() uint32 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

func (x *UpdateConfigRequest) GetInstanceConfig() *v1.InstanceConfig {
	if x != nil {
		return x.InstanceConfig
	}
	return nil
}

func (x *UpdateConfigRequest) GetPlayerConfig() *v1.PlayerConfig {
	if x != nil {
		return x.PlayerConfig
	}
	return nil
}

type ClearPakStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId   uint32 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	InstanceId uint32 `protobuf:"varint,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
}

func (x *ClearPakStateRequest) Reset() {
	*x = ClearPakStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_server_instance_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearPakStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearPakStateRequest) ProtoMessage() {}

func (x *ClearPakStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_server_instance_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearPakStateRequest.ProtoReflect.Descriptor instead.
func (*ClearPakStateRequest) Descriptor() ([]byte, []int) {
	return file_apis_v1_server_instance_api_proto_rawDescGZIP(), []int{7}
}

func (x *ClearPakStateRequest) GetClientId() uint32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *ClearPakStateRequest) GetInstanceId() uint32 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

var File_apis_v1_server_instance_api_proto protoreflect.FileDescriptor

var file_apis_v1_server_instance_api_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x55, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x17, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x8a, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x22, 0xd3, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x0d, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x54, 0x0a, 0x14, 0x43, 0x6c, 0x65, 0x61,
	0x72, 0x50, 0x61, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x32, 0xa3,
	0x04, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4e, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4c,
	0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x0d,
	0x43, 0x6c, 0x65, 0x61, 0x72, 0x50, 0x61, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x50, 0x61, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x97, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x75, 0x65, 0x2d, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x72,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69,
	0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x41, 0x70, 0x69, 0x73,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13,
	0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_v1_server_instance_api_proto_rawDescOnce sync.Once
	file_apis_v1_server_instance_api_proto_rawDescData = file_apis_v1_server_instance_api_proto_rawDesc
)

func file_apis_v1_server_instance_api_proto_rawDescGZIP() []byte {
	file_apis_v1_server_instance_api_proto_rawDescOnce.Do(func() {
		file_apis_v1_server_instance_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_v1_server_instance_api_proto_rawDescData)
	})
	return file_apis_v1_server_instance_api_proto_rawDescData
}

var file_apis_v1_server_instance_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_apis_v1_server_instance_api_proto_goTypes = []interface{}{
	(*AddInstanceRequest)(nil),        // 0: apis.v1.AddInstanceRequest
	(*DeleteInstanceRequest)(nil),     // 1: apis.v1.DeleteInstanceRequest
	(*GetStreamerIdRequest)(nil),      // 2: apis.v1.GetStreamerIdRequest
	(*GetStreamerIdResponse)(nil),     // 3: apis.v1.GetStreamerIdResponse
	(*UpdateRestartingRequest)(nil),   // 4: apis.v1.UpdateRestartingRequest
	(*UpdateProcessStateRequest)(nil), // 5: apis.v1.UpdateProcessStateRequest
	(*UpdateConfigRequest)(nil),       // 6: apis.v1.UpdateConfigRequest
	(*ClearPakStateRequest)(nil),      // 7: apis.v1.ClearPakStateRequest
	(*v1.InstanceInfo)(nil),           // 8: types.v1.InstanceInfo
	(*v1.InstanceConfig)(nil),         // 9: types.v1.InstanceConfig
	(*v1.PlayerConfig)(nil),           // 10: types.v1.PlayerConfig
	(*emptypb.Empty)(nil),             // 11: google.protobuf.Empty
}
var file_apis_v1_server_instance_api_proto_depIdxs = []int32{
	8,  // 0: apis.v1.AddInstanceRequest.instance_info:type_name -> types.v1.InstanceInfo
	9,  // 1: apis.v1.UpdateConfigRequest.instance_config:type_name -> types.v1.InstanceConfig
	10, // 2: apis.v1.UpdateConfigRequest.player_config:type_name -> types.v1.PlayerConfig
	0,  // 3: apis.v1.ServerInstanceService.AddInstance:input_type -> apis.v1.AddInstanceRequest
	1,  // 4: apis.v1.ServerInstanceService.DeleteInstance:input_type -> apis.v1.DeleteInstanceRequest
	6,  // 5: apis.v1.ServerInstanceService.UpdateConfig:input_type -> apis.v1.UpdateConfigRequest
	2,  // 6: apis.v1.ServerInstanceService.GetStreamerId:input_type -> apis.v1.GetStreamerIdRequest
	5,  // 7: apis.v1.ServerInstanceService.UpdateProcessState:input_type -> apis.v1.UpdateProcessStateRequest
	4,  // 8: apis.v1.ServerInstanceService.UpdateRestarting:input_type -> apis.v1.UpdateRestartingRequest
	7,  // 9: apis.v1.ServerInstanceService.ClearPakState:input_type -> apis.v1.ClearPakStateRequest
	11, // 10: apis.v1.ServerInstanceService.AddInstance:output_type -> google.protobuf.Empty
	11, // 11: apis.v1.ServerInstanceService.DeleteInstance:output_type -> google.protobuf.Empty
	11, // 12: apis.v1.ServerInstanceService.UpdateConfig:output_type -> google.protobuf.Empty
	3,  // 13: apis.v1.ServerInstanceService.GetStreamerId:output_type -> apis.v1.GetStreamerIdResponse
	11, // 14: apis.v1.ServerInstanceService.UpdateProcessState:output_type -> google.protobuf.Empty
	11, // 15: apis.v1.ServerInstanceService.UpdateRestarting:output_type -> google.protobuf.Empty
	11, // 16: apis.v1.ServerInstanceService.ClearPakState:output_type -> google.protobuf.Empty
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_apis_v1_server_instance_api_proto_init() }
func file_apis_v1_server_instance_api_proto_init() {
	if File_apis_v1_server_instance_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_v1_server_instance_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInstanceRequest); i {
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
		file_apis_v1_server_instance_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInstanceRequest); i {
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
		file_apis_v1_server_instance_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreamerIdRequest); i {
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
		file_apis_v1_server_instance_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreamerIdResponse); i {
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
		file_apis_v1_server_instance_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRestartingRequest); i {
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
		file_apis_v1_server_instance_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProcessStateRequest); i {
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
		file_apis_v1_server_instance_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigRequest); i {
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
		file_apis_v1_server_instance_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearPakStateRequest); i {
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
			RawDescriptor: file_apis_v1_server_instance_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apis_v1_server_instance_api_proto_goTypes,
		DependencyIndexes: file_apis_v1_server_instance_api_proto_depIdxs,
		MessageInfos:      file_apis_v1_server_instance_api_proto_msgTypes,
	}.Build()
	File_apis_v1_server_instance_api_proto = out.File
	file_apis_v1_server_instance_api_proto_rawDesc = nil
	file_apis_v1_server_instance_api_proto_goTypes = nil
	file_apis_v1_server_instance_api_proto_depIdxs = nil
}
