// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: proto/eth/v2/ssz.proto

package eth

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SSZContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version             Version `protobuf:"varint,1,opt,name=version,proto3,enum=ethereum.eth.v2.Version" json:"version,omitempty"`
	ExecutionOptimistic bool    `protobuf:"varint,2,opt,name=execution_optimistic,json=executionOptimistic,proto3" json:"execution_optimistic,omitempty"`
	Data                []byte  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Finalized           bool    `protobuf:"varint,4,opt,name=finalized,proto3" json:"finalized,omitempty"`
}

func (x *SSZContainer) Reset() {
	*x = SSZContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v2_ssz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSZContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSZContainer) ProtoMessage() {}

func (x *SSZContainer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v2_ssz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSZContainer.ProtoReflect.Descriptor instead.
func (*SSZContainer) Descriptor() ([]byte, []int) {
	return file_proto_eth_v2_ssz_proto_rawDescGZIP(), []int{0}
}

func (x *SSZContainer) GetVersion() Version {
	if x != nil {
		return x.Version
	}
	return Version_PHASE0
}

func (x *SSZContainer) GetExecutionOptimistic() bool {
	if x != nil {
		return x.ExecutionOptimistic
	}
	return false
}

func (x *SSZContainer) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SSZContainer) GetFinalized() bool {
	if x != nil {
		return x.Finalized
	}
	return false
}

var File_proto_eth_v2_ssz_proto protoreflect.FileDescriptor

var file_proto_eth_v2_ssz_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x2f, 0x73,
	0x73, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0c, 0x53, 0x53, 0x5a, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x14, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x42,
	0x79, 0x0a, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e,
	0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x42, 0x08, 0x53, 0x73, 0x7a, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x79,
	0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f,
	0x76, 0x32, 0x3b, 0x65, 0x74, 0x68, 0xaa, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x2e, 0x45, 0x74, 0x68, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x5c, 0x45, 0x74, 0x68, 0x5c, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_eth_v2_ssz_proto_rawDescOnce sync.Once
	file_proto_eth_v2_ssz_proto_rawDescData = file_proto_eth_v2_ssz_proto_rawDesc
)

func file_proto_eth_v2_ssz_proto_rawDescGZIP() []byte {
	file_proto_eth_v2_ssz_proto_rawDescOnce.Do(func() {
		file_proto_eth_v2_ssz_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_eth_v2_ssz_proto_rawDescData)
	})
	return file_proto_eth_v2_ssz_proto_rawDescData
}

var file_proto_eth_v2_ssz_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_eth_v2_ssz_proto_goTypes = []interface{}{
	(*SSZContainer)(nil), // 0: ethereum.eth.v2.SSZContainer
	(Version)(0),         // 1: ethereum.eth.v2.Version
}
var file_proto_eth_v2_ssz_proto_depIdxs = []int32{
	1, // 0: ethereum.eth.v2.SSZContainer.version:type_name -> ethereum.eth.v2.Version
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_eth_v2_ssz_proto_init() }
func file_proto_eth_v2_ssz_proto_init() {
	if File_proto_eth_v2_ssz_proto != nil {
		return
	}
	file_proto_eth_v2_version_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_eth_v2_ssz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSZContainer); i {
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
			RawDescriptor: file_proto_eth_v2_ssz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_eth_v2_ssz_proto_goTypes,
		DependencyIndexes: file_proto_eth_v2_ssz_proto_depIdxs,
		MessageInfos:      file_proto_eth_v2_ssz_proto_msgTypes,
	}.Build()
	File_proto_eth_v2_ssz_proto = out.File
	file_proto_eth_v2_ssz_proto_rawDesc = nil
	file_proto_eth_v2_ssz_proto_goTypes = nil
	file_proto_eth_v2_ssz_proto_depIdxs = nil
}
