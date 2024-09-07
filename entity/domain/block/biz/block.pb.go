// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: domain/block/biz/block.proto

package biz

import (
	model "github.com/blackhorseya/ryze/entity/domain/block/model"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for retrieving a single block by its ID.
type GetBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workchain int32  `protobuf:"varint,1,opt,name=workchain,proto3" json:"workchain,omitempty"`
	Shard     int64  `protobuf:"varint,2,opt,name=shard,proto3" json:"shard,omitempty"`
	SeqNo     uint32 `protobuf:"varint,3,opt,name=seq_no,json=seqNo,proto3" json:"seq_no,omitempty"`
}

func (x *GetBlockRequest) Reset() {
	*x = GetBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_block_biz_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockRequest) ProtoMessage() {}

func (x *GetBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_domain_block_biz_block_proto_msgTypes[0]
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
	return file_domain_block_biz_block_proto_rawDescGZIP(), []int{0}
}

func (x *GetBlockRequest) GetWorkchain() int32 {
	if x != nil {
		return x.Workchain
	}
	return 0
}

func (x *GetBlockRequest) GetShard() int64 {
	if x != nil {
		return x.Shard
	}
	return 0
}

func (x *GetBlockRequest) GetSeqNo() uint32 {
	if x != nil {
		return x.SeqNo
	}
	return 0
}

// Request message for retrieving a range of blocks by their heights.
type GetBlocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Starting height of the range of blocks to retrieve.
	StartHeight uint32 `protobuf:"varint,1,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	// Ending height of the range of blocks to retrieve.
	EndHeight uint32 `protobuf:"varint,2,opt,name=end_height,json=endHeight,proto3" json:"end_height,omitempty"`
}

func (x *GetBlocksRequest) Reset() {
	*x = GetBlocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_block_biz_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocksRequest) ProtoMessage() {}

func (x *GetBlocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_domain_block_biz_block_proto_msgTypes[1]
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
	return file_domain_block_biz_block_proto_rawDescGZIP(), []int{1}
}

func (x *GetBlocksRequest) GetStartHeight() uint32 {
	if x != nil {
		return x.StartHeight
	}
	return 0
}

func (x *GetBlocksRequest) GetEndHeight() uint32 {
	if x != nil {
		return x.EndHeight
	}
	return 0
}

// Request message for scanning blocks.
type ScanBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Starting height of the range of blocks to scan.
	StartHeight uint32 `protobuf:"varint,1,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	// Ending height of the range of blocks to scan.
	EndHeight uint32 `protobuf:"varint,2,opt,name=end_height,json=endHeight,proto3" json:"end_height,omitempty"`
}

func (x *ScanBlockRequest) Reset() {
	*x = ScanBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_block_biz_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanBlockRequest) ProtoMessage() {}

func (x *ScanBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_domain_block_biz_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanBlockRequest.ProtoReflect.Descriptor instead.
func (*ScanBlockRequest) Descriptor() ([]byte, []int) {
	return file_domain_block_biz_block_proto_rawDescGZIP(), []int{2}
}

func (x *ScanBlockRequest) GetStartHeight() uint32 {
	if x != nil {
		return x.StartHeight
	}
	return 0
}

func (x *ScanBlockRequest) GetEndHeight() uint32 {
	if x != nil {
		return x.EndHeight
	}
	return 0
}

// Request message for fetching and storing a block.
type FetchAndStoreBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workchain int32  `protobuf:"varint,1,opt,name=workchain,proto3" json:"workchain,omitempty"`
	Shard     int64  `protobuf:"varint,2,opt,name=shard,proto3" json:"shard,omitempty"`
	SeqNo     uint32 `protobuf:"varint,3,opt,name=seq_no,json=seqNo,proto3" json:"seq_no,omitempty"`
}

func (x *FetchAndStoreBlockRequest) Reset() {
	*x = FetchAndStoreBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_block_biz_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchAndStoreBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAndStoreBlockRequest) ProtoMessage() {}

func (x *FetchAndStoreBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_domain_block_biz_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAndStoreBlockRequest.ProtoReflect.Descriptor instead.
func (*FetchAndStoreBlockRequest) Descriptor() ([]byte, []int) {
	return file_domain_block_biz_block_proto_rawDescGZIP(), []int{3}
}

func (x *FetchAndStoreBlockRequest) GetWorkchain() int32 {
	if x != nil {
		return x.Workchain
	}
	return 0
}

func (x *FetchAndStoreBlockRequest) GetShard() int64 {
	if x != nil {
		return x.Shard
	}
	return 0
}

func (x *FetchAndStoreBlockRequest) GetSeqNo() uint32 {
	if x != nil {
		return x.SeqNo
	}
	return 0
}

// Response message for fetching and storing a block.
type FetchAndStoreBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status of the operation.
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// The fetched and stored block.
	Block *model.Block `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *FetchAndStoreBlockResponse) Reset() {
	*x = FetchAndStoreBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_block_biz_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchAndStoreBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAndStoreBlockResponse) ProtoMessage() {}

func (x *FetchAndStoreBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_domain_block_biz_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAndStoreBlockResponse.ProtoReflect.Descriptor instead.
func (*FetchAndStoreBlockResponse) Descriptor() ([]byte, []int) {
	return file_domain_block_biz_block_proto_rawDescGZIP(), []int{4}
}

func (x *FetchAndStoreBlockResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *FetchAndStoreBlockResponse) GetBlock() *model.Block {
	if x != nil {
		return x.Block
	}
	return nil
}

var File_domain_block_biz_block_proto protoreflect.FileDescriptor

var file_domain_block_biz_block_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x62,
	0x69, 0x7a, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x1e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x77, 0x6f, 0x72,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x73, 0x65, 0x71, 0x5f, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x65,
	0x71, 0x4e, 0x6f, 0x22, 0x54, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e,
	0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x65, 0x6e, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x54, 0x0a, 0x10, 0x53, 0x63, 0x61,
	0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22,
	0x66, 0x0a, 0x19, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x77, 0x6f, 0x72, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x77, 0x6f, 0x72, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x73, 0x65, 0x71, 0x4e, 0x6f, 0x22, 0x58, 0x0a, 0x1a, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x41, 0x6e, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a,
	0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x32, 0xb2, 0x01, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x12, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x12, 0x36,
	0x0a, 0x09, 0x53, 0x63, 0x61, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x42, 0x81, 0x01, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x79, 0x61, 0x2f, 0x72, 0x79, 0x7a, 0x65,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x62, 0x69, 0x7a, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02,
	0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0xca, 0x02, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0xe2, 0x02,
	0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_domain_block_biz_block_proto_rawDescOnce sync.Once
	file_domain_block_biz_block_proto_rawDescData = file_domain_block_biz_block_proto_rawDesc
)

func file_domain_block_biz_block_proto_rawDescGZIP() []byte {
	file_domain_block_biz_block_proto_rawDescOnce.Do(func() {
		file_domain_block_biz_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_block_biz_block_proto_rawDescData)
	})
	return file_domain_block_biz_block_proto_rawDescData
}

var file_domain_block_biz_block_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_domain_block_biz_block_proto_goTypes = []any{
	(*GetBlockRequest)(nil),            // 0: block.GetBlockRequest
	(*GetBlocksRequest)(nil),           // 1: block.GetBlocksRequest
	(*ScanBlockRequest)(nil),           // 2: block.ScanBlockRequest
	(*FetchAndStoreBlockRequest)(nil),  // 3: block.FetchAndStoreBlockRequest
	(*FetchAndStoreBlockResponse)(nil), // 4: block.FetchAndStoreBlockResponse
	(*model.Block)(nil),                // 5: block.Block
}
var file_domain_block_biz_block_proto_depIdxs = []int32{
	5, // 0: block.FetchAndStoreBlockResponse.block:type_name -> block.Block
	0, // 1: block.BlockService.GetBlock:input_type -> block.GetBlockRequest
	1, // 2: block.BlockService.GetBlocks:input_type -> block.GetBlocksRequest
	2, // 3: block.BlockService.ScanBlock:input_type -> block.ScanBlockRequest
	5, // 4: block.BlockService.GetBlock:output_type -> block.Block
	5, // 5: block.BlockService.GetBlocks:output_type -> block.Block
	5, // 6: block.BlockService.ScanBlock:output_type -> block.Block
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_domain_block_biz_block_proto_init() }
func file_domain_block_biz_block_proto_init() {
	if File_domain_block_biz_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_block_biz_block_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_domain_block_biz_block_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_domain_block_biz_block_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ScanBlockRequest); i {
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
		file_domain_block_biz_block_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*FetchAndStoreBlockRequest); i {
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
		file_domain_block_biz_block_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*FetchAndStoreBlockResponse); i {
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
			RawDescriptor: file_domain_block_biz_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_domain_block_biz_block_proto_goTypes,
		DependencyIndexes: file_domain_block_biz_block_proto_depIdxs,
		MessageInfos:      file_domain_block_biz_block_proto_msgTypes,
	}.Build()
	File_domain_block_biz_block_proto = out.File
	file_domain_block_biz_block_proto_rawDesc = nil
	file_domain_block_biz_block_proto_goTypes = nil
	file_domain_block_biz_block_proto_depIdxs = nil
}
