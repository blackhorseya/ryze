// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: domain/transaction/biz/transaction.proto

package biz

import (
	model "github.com/blackhorseya/ryze/entity/domain/block/model"
	model1 "github.com/blackhorseya/ryze/entity/domain/transaction/model"
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

type ListTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListTransactionRequest) Reset() {
	*x = ListTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_transaction_biz_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransactionRequest) ProtoMessage() {}

func (x *ListTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_domain_transaction_biz_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransactionRequest.ProtoReflect.Descriptor instead.
func (*ListTransactionRequest) Descriptor() ([]byte, []int) {
	return file_domain_transaction_biz_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *ListTransactionRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListTransactionRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListTransactionsByAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Page      int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListTransactionsByAccountRequest) Reset() {
	*x = ListTransactionsByAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_transaction_biz_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransactionsByAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransactionsByAccountRequest) ProtoMessage() {}

func (x *ListTransactionsByAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_domain_transaction_biz_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransactionsByAccountRequest.ProtoReflect.Descriptor instead.
func (*ListTransactionsByAccountRequest) Descriptor() ([]byte, []int) {
	return file_domain_transaction_biz_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *ListTransactionsByAccountRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *ListTransactionsByAccountRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListTransactionsByAccountRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

var File_domain_transaction_biz_transaction_proto protoreflect.FileDescriptor

var file_domain_transaction_biz_transaction_proto_rawDesc = []byte{
	0x0a, 0x28, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x72,
	0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x32, 0x9f, 0x02, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x18, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x1a, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x12, 0x55, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x12, 0x68, 0x0a, 0x19, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x30, 0x01, 0x42, 0xab, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f,
	0x72, 0x73, 0x65, 0x79, 0x61, 0x2f, 0x72, 0x79, 0x7a, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x69, 0x7a, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xca, 0x02, 0x0b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x17, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_transaction_biz_transaction_proto_rawDescOnce sync.Once
	file_domain_transaction_biz_transaction_proto_rawDescData = file_domain_transaction_biz_transaction_proto_rawDesc
)

func file_domain_transaction_biz_transaction_proto_rawDescGZIP() []byte {
	file_domain_transaction_biz_transaction_proto_rawDescOnce.Do(func() {
		file_domain_transaction_biz_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_transaction_biz_transaction_proto_rawDescData)
	})
	return file_domain_transaction_biz_transaction_proto_rawDescData
}

var file_domain_transaction_biz_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_domain_transaction_biz_transaction_proto_goTypes = []any{
	(*ListTransactionRequest)(nil),           // 0: transaction.ListTransactionRequest
	(*ListTransactionsByAccountRequest)(nil), // 1: transaction.ListTransactionsByAccountRequest
	(*model.Block)(nil),                      // 2: block.Block
	(*model1.Transaction)(nil),               // 3: transaction.Transaction
}
var file_domain_transaction_biz_transaction_proto_depIdxs = []int32{
	2, // 0: transaction.TransactionService.ProcessBlockTransactions:input_type -> block.Block
	0, // 1: transaction.TransactionService.ListTransactions:input_type -> transaction.ListTransactionRequest
	1, // 2: transaction.TransactionService.ListTransactionsByAccount:input_type -> transaction.ListTransactionsByAccountRequest
	3, // 3: transaction.TransactionService.ProcessBlockTransactions:output_type -> transaction.Transaction
	3, // 4: transaction.TransactionService.ListTransactions:output_type -> transaction.Transaction
	3, // 5: transaction.TransactionService.ListTransactionsByAccount:output_type -> transaction.Transaction
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_domain_transaction_biz_transaction_proto_init() }
func file_domain_transaction_biz_transaction_proto_init() {
	if File_domain_transaction_biz_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_transaction_biz_transaction_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListTransactionRequest); i {
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
		file_domain_transaction_biz_transaction_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListTransactionsByAccountRequest); i {
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
			RawDescriptor: file_domain_transaction_biz_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_domain_transaction_biz_transaction_proto_goTypes,
		DependencyIndexes: file_domain_transaction_biz_transaction_proto_depIdxs,
		MessageInfos:      file_domain_transaction_biz_transaction_proto_msgTypes,
	}.Build()
	File_domain_transaction_biz_transaction_proto = out.File
	file_domain_transaction_biz_transaction_proto_rawDesc = nil
	file_domain_transaction_biz_transaction_proto_goTypes = nil
	file_domain_transaction_biz_transaction_proto_depIdxs = nil
}
