// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: entity/domain/block/model/block.proto

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

// Represents a block in the blockchain.
type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the block.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Height of the block in the blockchain.
	Height int64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	// Timestamp of when the block was created.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// List of transaction IDs included in the block.
	TransactionIds [][]byte `protobuf:"bytes,4,rep,name=transaction_ids,json=transactionIds,proto3" json:"transaction_ids,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_domain_block_model_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_entity_domain_block_model_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_entity_domain_block_model_block_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Block) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Block) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Block) GetTransactionIds() [][]byte {
	if x != nil {
		return x.TransactionIds
	}
	return nil
}

// Request message for retrieving a single block by its ID.
type GetBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the block to retrieve.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBlockRequest) Reset() {
	*x = GetBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_domain_block_model_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockRequest) ProtoMessage() {}

func (x *GetBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entity_domain_block_model_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockRequest.ProtoReflect.Descriptor instead.
func (*GetBlockRequest) Descriptor() ([]byte, []int) {
	return file_entity_domain_block_model_block_proto_rawDescGZIP(), []int{1}
}

func (x *GetBlockRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

// Request message for retrieving a range of blocks by their heights.
type GetBlocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Starting height of the range of blocks to retrieve.
	StartHeight int64 `protobuf:"varint,1,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	// Ending height of the range of blocks to retrieve.
	EndHeight int64 `protobuf:"varint,2,opt,name=end_height,json=endHeight,proto3" json:"end_height,omitempty"`
}

func (x *GetBlocksRequest) Reset() {
	*x = GetBlocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_domain_block_model_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocksRequest) ProtoMessage() {}

func (x *GetBlocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entity_domain_block_model_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlocksRequest.ProtoReflect.Descriptor instead.
func (*GetBlocksRequest) Descriptor() ([]byte, []int) {
	return file_entity_domain_block_model_block_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlocksRequest) GetStartHeight() int64 {
	if x != nil {
		return x.StartHeight
	}
	return 0
}

func (x *GetBlocksRequest) GetEndHeight() int64 {
	if x != nil {
		return x.EndHeight
	}
	return 0
}

var File_entity_domain_block_model_block_proto protoreflect.FileDescriptor

var file_entity_domain_block_model_block_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x92, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0x7a, 0x0a,
	0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x17,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x72,
	0x73, 0x65, 0x79, 0x61, 0x2f, 0x72, 0x79, 0x7a, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_domain_block_model_block_proto_rawDescOnce sync.Once
	file_entity_domain_block_model_block_proto_rawDescData = file_entity_domain_block_model_block_proto_rawDesc
)

func file_entity_domain_block_model_block_proto_rawDescGZIP() []byte {
	file_entity_domain_block_model_block_proto_rawDescOnce.Do(func() {
		file_entity_domain_block_model_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_domain_block_model_block_proto_rawDescData)
	})
	return file_entity_domain_block_model_block_proto_rawDescData
}

var file_entity_domain_block_model_block_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_entity_domain_block_model_block_proto_goTypes = []any{
	(*Block)(nil),                 // 0: block.Block
	(*GetBlockRequest)(nil),       // 1: block.GetBlockRequest
	(*GetBlocksRequest)(nil),      // 2: block.GetBlocksRequest
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_entity_domain_block_model_block_proto_depIdxs = []int32{
	3, // 0: block.Block.timestamp:type_name -> google.protobuf.Timestamp
	1, // 1: block.BlockService.GetBlock:input_type -> block.GetBlockRequest
	2, // 2: block.BlockService.GetBlocks:input_type -> block.GetBlocksRequest
	0, // 3: block.BlockService.GetBlock:output_type -> block.Block
	0, // 4: block.BlockService.GetBlocks:output_type -> block.Block
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_entity_domain_block_model_block_proto_init() }
func file_entity_domain_block_model_block_proto_init() {
	if File_entity_domain_block_model_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entity_domain_block_model_block_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Block); i {
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
		file_entity_domain_block_model_block_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetBlockRequest); i {
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
		file_entity_domain_block_model_block_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetBlocksRequest); i {
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
			RawDescriptor: file_entity_domain_block_model_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_entity_domain_block_model_block_proto_goTypes,
		DependencyIndexes: file_entity_domain_block_model_block_proto_depIdxs,
		MessageInfos:      file_entity_domain_block_model_block_proto_msgTypes,
	}.Build()
	File_entity_domain_block_model_block_proto = out.File
	file_entity_domain_block_model_block_proto_rawDesc = nil
	file_entity_domain_block_model_block_proto_goTypes = nil
	file_entity_domain_block_model_block_proto_depIdxs = nil
}