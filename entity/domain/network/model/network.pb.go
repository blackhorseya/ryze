// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: domain/network/model/network.proto

package model

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

// Represents the statistics of the blockchain network.
type NetworkStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalBlocks       int64                  `protobuf:"varint,1,opt,name=total_blocks,json=totalBlocks,proto3" json:"total_blocks,omitempty"`
	TotalTransactions int64                  `protobuf:"varint,2,opt,name=total_transactions,json=totalTransactions,proto3" json:"total_transactions,omitempty"`
	TotalAccounts     int64                  `protobuf:"varint,3,opt,name=total_accounts,json=totalAccounts,proto3" json:"total_accounts,omitempty"`
	LatestBlockHeight uint32                 `protobuf:"varint,4,opt,name=latest_block_height,json=latestBlockHeight,proto3" json:"latest_block_height,omitempty"`
	LatestBlockTime   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=latest_block_time,json=latestBlockTime,proto3" json:"latest_block_time,omitempty"`
}

func (x *NetworkStats) Reset() {
	*x = NetworkStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_network_model_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkStats) ProtoMessage() {}

func (x *NetworkStats) ProtoReflect() protoreflect.Message {
	mi := &file_domain_network_model_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkStats.ProtoReflect.Descriptor instead.
func (*NetworkStats) Descriptor() ([]byte, []int) {
	return file_domain_network_model_network_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkStats) GetTotalBlocks() int64 {
	if x != nil {
		return x.TotalBlocks
	}
	return 0
}

func (x *NetworkStats) GetTotalTransactions() int64 {
	if x != nil {
		return x.TotalTransactions
	}
	return 0
}

func (x *NetworkStats) GetTotalAccounts() int64 {
	if x != nil {
		return x.TotalAccounts
	}
	return 0
}

func (x *NetworkStats) GetLatestBlockHeight() uint32 {
	if x != nil {
		return x.LatestBlockHeight
	}
	return 0
}

func (x *NetworkStats) GetLatestBlockTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LatestBlockTime
	}
	return nil
}

// Represents the status of a node in the blockchain network.
type NodeStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId      []byte                 `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	IsConnected bool                   `protobuf:"varint,2,opt,name=is_connected,json=isConnected,proto3" json:"is_connected,omitempty"`
	LastActive  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_active,json=lastActive,proto3" json:"last_active,omitempty"`
}

func (x *NodeStatus) Reset() {
	*x = NodeStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_network_model_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatus) ProtoMessage() {}

func (x *NodeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_domain_network_model_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatus.ProtoReflect.Descriptor instead.
func (*NodeStatus) Descriptor() ([]byte, []int) {
	return file_domain_network_model_network_proto_rawDescGZIP(), []int{1}
}

func (x *NodeStatus) GetNodeId() []byte {
	if x != nil {
		return x.NodeId
	}
	return nil
}

func (x *NodeStatus) GetIsConnected() bool {
	if x != nil {
		return x.IsConnected
	}
	return false
}

func (x *NodeStatus) GetLastActive() *timestamppb.Timestamp {
	if x != nil {
		return x.LastActive
	}
	return nil
}

var File_domain_network_model_network_proto protoreflect.FileDescriptor

var file_domain_network_model_network_proto_rawDesc = []byte{
	0x0a, 0x22, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff,
	0x01, 0x0a, 0x0c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x46, 0x0a, 0x11, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x85, 0x01, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61,
	0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x91, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x0c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x79,
	0x61, 0x2f, 0x72, 0x79, 0x7a, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0xa2, 0x02, 0x03, 0x4e, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0xca, 0x02, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0xe2, 0x02, 0x13, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_network_model_network_proto_rawDescOnce sync.Once
	file_domain_network_model_network_proto_rawDescData = file_domain_network_model_network_proto_rawDesc
)

func file_domain_network_model_network_proto_rawDescGZIP() []byte {
	file_domain_network_model_network_proto_rawDescOnce.Do(func() {
		file_domain_network_model_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_network_model_network_proto_rawDescData)
	})
	return file_domain_network_model_network_proto_rawDescData
}

var file_domain_network_model_network_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_domain_network_model_network_proto_goTypes = []any{
	(*NetworkStats)(nil),          // 0: network.NetworkStats
	(*NodeStatus)(nil),            // 1: network.NodeStatus
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_domain_network_model_network_proto_depIdxs = []int32{
	2, // 0: network.NetworkStats.latest_block_time:type_name -> google.protobuf.Timestamp
	2, // 1: network.NodeStatus.last_active:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_domain_network_model_network_proto_init() }
func file_domain_network_model_network_proto_init() {
	if File_domain_network_model_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_network_model_network_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NetworkStats); i {
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
		file_domain_network_model_network_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NodeStatus); i {
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
			RawDescriptor: file_domain_network_model_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_network_model_network_proto_goTypes,
		DependencyIndexes: file_domain_network_model_network_proto_depIdxs,
		MessageInfos:      file_domain_network_model_network_proto_msgTypes,
	}.Build()
	File_domain_network_model_network_proto = out.File
	file_domain_network_model_network_proto_rawDesc = nil
	file_domain_network_model_network_proto_goTypes = nil
	file_domain_network_model_network_proto_depIdxs = nil
}
