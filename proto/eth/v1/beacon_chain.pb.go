// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: proto/eth/v1/beacon_chain.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/prysmaticlabs/prysm/v4/proto/eth/ext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateId []byte `protobuf:"bytes,1,opt,name=state_id,json=stateId,proto3" json:"state_id,omitempty"`
}

func (x *StateRequest) Reset() {
	*x = StateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateRequest) ProtoMessage() {}

func (x *StateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateRequest.ProtoReflect.Descriptor instead.
func (*StateRequest) Descriptor() ([]byte, []int) {
	return file_proto_eth_v1_beacon_chain_proto_rawDescGZIP(), []int{0}
}

func (x *StateRequest) GetStateId() []byte {
	if x != nil {
		return x.StateId
	}
	return nil
}

type AttesterSlashingsPoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*AttesterSlashing `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *AttesterSlashingsPoolResponse) Reset() {
	*x = AttesterSlashingsPoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttesterSlashingsPoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttesterSlashingsPoolResponse) ProtoMessage() {}

func (x *AttesterSlashingsPoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttesterSlashingsPoolResponse.ProtoReflect.Descriptor instead.
func (*AttesterSlashingsPoolResponse) Descriptor() ([]byte, []int) {
	return file_proto_eth_v1_beacon_chain_proto_rawDescGZIP(), []int{1}
}

func (x *AttesterSlashingsPoolResponse) GetData() []*AttesterSlashing {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProposerSlashingPoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ProposerSlashing `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ProposerSlashingPoolResponse) Reset() {
	*x = ProposerSlashingPoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposerSlashingPoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposerSlashingPoolResponse) ProtoMessage() {}

func (x *ProposerSlashingPoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposerSlashingPoolResponse.ProtoReflect.Descriptor instead.
func (*ProposerSlashingPoolResponse) Descriptor() ([]byte, []int) {
	return file_proto_eth_v1_beacon_chain_proto_rawDescGZIP(), []int{2}
}

func (x *ProposerSlashingPoolResponse) GetData() []*ProposerSlashing {
	if x != nil {
		return x.Data
	}
	return nil
}

type ForkScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Fork `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ForkScheduleResponse) Reset() {
	*x = ForkScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForkScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForkScheduleResponse) ProtoMessage() {}

func (x *ForkScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForkScheduleResponse.ProtoReflect.Descriptor instead.
func (*ForkScheduleResponse) Descriptor() ([]byte, []int) {
	return file_proto_eth_v1_beacon_chain_proto_rawDescGZIP(), []int{3}
}

func (x *ForkScheduleResponse) GetData() []*Fork {
	if x != nil {
		return x.Data
	}
	return nil
}

type SpecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SpecResponse) Reset() {
	*x = SpecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecResponse) ProtoMessage() {}

func (x *SpecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecResponse.ProtoReflect.Descriptor instead.
func (*SpecResponse) Descriptor() ([]byte, []int) {
	return file_proto_eth_v1_beacon_chain_proto_rawDescGZIP(), []int{4}
}

func (x *SpecResponse) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type DepositContractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *DepositContract `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DepositContractResponse) Reset() {
	*x = DepositContractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositContractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositContractResponse) ProtoMessage() {}

func (x *DepositContractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositContractResponse.ProtoReflect.Descriptor instead.
func (*DepositContractResponse) Descriptor() ([]byte, []int) {
	return file_proto_eth_v1_beacon_chain_proto_rawDescGZIP(), []int{5}
}

func (x *DepositContractResponse) GetData() *DepositContract {
	if x != nil {
		return x.Data
	}
	return nil
}

type DepositContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId uint64 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *DepositContract) Reset() {
	*x = DepositContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositContract) ProtoMessage() {}

func (x *DepositContract) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositContract.ProtoReflect.Descriptor instead.
func (*DepositContract) Descriptor() ([]byte, []int) {
	return file_proto_eth_v1_beacon_chain_proto_rawDescGZIP(), []int{6}
}

func (x *DepositContract) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *DepositContract) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// Deprecated: Marked as deprecated in proto/eth/v1/beacon_chain.proto.
type WeakSubjectivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *WeakSubjectivityData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *WeakSubjectivityResponse) Reset() {
	*x = WeakSubjectivityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeakSubjectivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeakSubjectivityResponse) ProtoMessage() {}

func (x *WeakSubjectivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeakSubjectivityResponse.ProtoReflect.Descriptor instead.
func (*WeakSubjectivityResponse) Descriptor() ([]byte, []int) {
	return file_proto_eth_v1_beacon_chain_proto_rawDescGZIP(), []int{7}
}

func (x *WeakSubjectivityResponse) GetData() *WeakSubjectivityData {
	if x != nil {
		return x.Data
	}
	return nil
}

// Deprecated: Marked as deprecated in proto/eth/v1/beacon_chain.proto.
type WeakSubjectivityData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WsCheckpoint *Checkpoint `protobuf:"bytes,1,opt,name=ws_checkpoint,json=wsCheckpoint,proto3" json:"ws_checkpoint,omitempty"`
	StateRoot    []byte      `protobuf:"bytes,2,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
}

func (x *WeakSubjectivityData) Reset() {
	*x = WeakSubjectivityData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeakSubjectivityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeakSubjectivityData) ProtoMessage() {}

func (x *WeakSubjectivityData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v1_beacon_chain_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeakSubjectivityData.ProtoReflect.Descriptor instead.
func (*WeakSubjectivityData) Descriptor() ([]byte, []int) {
	return file_proto_eth_v1_beacon_chain_proto_rawDescGZIP(), []int{8}
}

func (x *WeakSubjectivityData) GetWsCheckpoint() *Checkpoint {
	if x != nil {
		return x.WsCheckpoint
	}
	return nil
}

func (x *WeakSubjectivityData) GetStateRoot() []byte {
	if x != nil {
		return x.StateRoot
	}
	return nil
}

var File_proto_eth_v1_beacon_chain_proto protoreflect.FileDescriptor

var file_proto_eth_v1_beacon_chain_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f,
	0x65, 0x78, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x74, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x56,
	0x0a, 0x1d, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0x73, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x55, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e,
	0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x53,
	0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x41, 0x0a,
	0x14, 0x46, 0x6f, 0x72, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6b, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x84, 0x01, 0x0a, 0x0c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37,
	0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4f, 0x0a, 0x17, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x46, 0x0a, 0x0f, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x59, 0x0a, 0x18, 0x57, 0x65, 0x61, 0x6b, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x61,
	0x6b, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x02, 0x18, 0x01, 0x22, 0x7b, 0x0a, 0x14, 0x57,
	0x65, 0x61, 0x6b, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x0d, 0x77, 0x73, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0c, 0x77, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72,
	0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6f, 0x74, 0x3a, 0x02, 0x18, 0x01, 0x42, 0x7d, 0x0a, 0x13, 0x6f, 0x72, 0x67, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42,
	0x10, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72,
	0x79, 0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68,
	0x2f, 0x76, 0x31, 0xaa, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x45,
	0x74, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x5c, 0x45, 0x74, 0x68, 0x5c, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_eth_v1_beacon_chain_proto_rawDescOnce sync.Once
	file_proto_eth_v1_beacon_chain_proto_rawDescData = file_proto_eth_v1_beacon_chain_proto_rawDesc
)

func file_proto_eth_v1_beacon_chain_proto_rawDescGZIP() []byte {
	file_proto_eth_v1_beacon_chain_proto_rawDescOnce.Do(func() {
		file_proto_eth_v1_beacon_chain_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_eth_v1_beacon_chain_proto_rawDescData)
	})
	return file_proto_eth_v1_beacon_chain_proto_rawDescData
}

var file_proto_eth_v1_beacon_chain_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_eth_v1_beacon_chain_proto_goTypes = []interface{}{
	(*StateRequest)(nil),                  // 0: ethereum.eth.v1.StateRequest
	(*AttesterSlashingsPoolResponse)(nil), // 1: ethereum.eth.v1.AttesterSlashingsPoolResponse
	(*ProposerSlashingPoolResponse)(nil),  // 2: ethereum.eth.v1.ProposerSlashingPoolResponse
	(*ForkScheduleResponse)(nil),          // 3: ethereum.eth.v1.ForkScheduleResponse
	(*SpecResponse)(nil),                  // 4: ethereum.eth.v1.SpecResponse
	(*DepositContractResponse)(nil),       // 5: ethereum.eth.v1.DepositContractResponse
	(*DepositContract)(nil),               // 6: ethereum.eth.v1.DepositContract
	(*WeakSubjectivityResponse)(nil),      // 7: ethereum.eth.v1.WeakSubjectivityResponse
	(*WeakSubjectivityData)(nil),          // 8: ethereum.eth.v1.WeakSubjectivityData
	nil,                                   // 9: ethereum.eth.v1.SpecResponse.DataEntry
	(*AttesterSlashing)(nil),              // 10: ethereum.eth.v1.AttesterSlashing
	(*ProposerSlashing)(nil),              // 11: ethereum.eth.v1.ProposerSlashing
	(*Fork)(nil),                          // 12: ethereum.eth.v1.Fork
	(*Checkpoint)(nil),                    // 13: ethereum.eth.v1.Checkpoint
}
var file_proto_eth_v1_beacon_chain_proto_depIdxs = []int32{
	10, // 0: ethereum.eth.v1.AttesterSlashingsPoolResponse.data:type_name -> ethereum.eth.v1.AttesterSlashing
	11, // 1: ethereum.eth.v1.ProposerSlashingPoolResponse.data:type_name -> ethereum.eth.v1.ProposerSlashing
	12, // 2: ethereum.eth.v1.ForkScheduleResponse.data:type_name -> ethereum.eth.v1.Fork
	9,  // 3: ethereum.eth.v1.SpecResponse.data:type_name -> ethereum.eth.v1.SpecResponse.DataEntry
	6,  // 4: ethereum.eth.v1.DepositContractResponse.data:type_name -> ethereum.eth.v1.DepositContract
	8,  // 5: ethereum.eth.v1.WeakSubjectivityResponse.data:type_name -> ethereum.eth.v1.WeakSubjectivityData
	13, // 6: ethereum.eth.v1.WeakSubjectivityData.ws_checkpoint:type_name -> ethereum.eth.v1.Checkpoint
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_eth_v1_beacon_chain_proto_init() }
func file_proto_eth_v1_beacon_chain_proto_init() {
	if File_proto_eth_v1_beacon_chain_proto != nil {
		return
	}
	file_proto_eth_v1_attestation_proto_init()
	file_proto_eth_v1_beacon_block_proto_init()
	file_proto_eth_v1_beacon_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_eth_v1_beacon_chain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateRequest); i {
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
		file_proto_eth_v1_beacon_chain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttesterSlashingsPoolResponse); i {
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
		file_proto_eth_v1_beacon_chain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposerSlashingPoolResponse); i {
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
		file_proto_eth_v1_beacon_chain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForkScheduleResponse); i {
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
		file_proto_eth_v1_beacon_chain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecResponse); i {
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
		file_proto_eth_v1_beacon_chain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositContractResponse); i {
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
		file_proto_eth_v1_beacon_chain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositContract); i {
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
		file_proto_eth_v1_beacon_chain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeakSubjectivityResponse); i {
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
		file_proto_eth_v1_beacon_chain_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeakSubjectivityData); i {
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
			RawDescriptor: file_proto_eth_v1_beacon_chain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_eth_v1_beacon_chain_proto_goTypes,
		DependencyIndexes: file_proto_eth_v1_beacon_chain_proto_depIdxs,
		MessageInfos:      file_proto_eth_v1_beacon_chain_proto_msgTypes,
	}.Build()
	File_proto_eth_v1_beacon_chain_proto = out.File
	file_proto_eth_v1_beacon_chain_proto_rawDesc = nil
	file_proto_eth_v1_beacon_chain_proto_goTypes = nil
	file_proto_eth_v1_beacon_chain_proto_depIdxs = nil
}
