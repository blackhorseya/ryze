// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: entity/domain/block/model/block.proto

package model

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BlockService_GetBlock_FullMethodName           = "/block.BlockService/GetBlock"
	BlockService_GetBlocks_FullMethodName          = "/block.BlockService/GetBlocks"
	BlockService_ScanBlock_FullMethodName          = "/block.BlockService/ScanBlock"
	BlockService_FetchAndStoreBlock_FullMethodName = "/block.BlockService/FetchAndStoreBlock"
)

// BlockServiceClient is the client API for BlockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockServiceClient interface {
	// Retrieves a single block by its ID.
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*Block, error)
	// Retrieves a stream of blocks within a specified height range.
	GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (BlockService_GetBlocksClient, error)
	// Scans a range of blocks.
	ScanBlock(ctx context.Context, in *ScanBlockRequest, opts ...grpc.CallOption) (BlockService_ScanBlockClient, error)
	FetchAndStoreBlock(ctx context.Context, in *FetchAndStoreBlockRequest, opts ...grpc.CallOption) (*FetchAndStoreBlockResponse, error)
}

type blockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockServiceClient(cc grpc.ClientConnInterface) BlockServiceClient {
	return &blockServiceClient{cc}
}

func (c *blockServiceClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, BlockService_GetBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockServiceClient) GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (BlockService_GetBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlockService_ServiceDesc.Streams[0], BlockService_GetBlocks_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &blockServiceGetBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockService_GetBlocksClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type blockServiceGetBlocksClient struct {
	grpc.ClientStream
}

func (x *blockServiceGetBlocksClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockServiceClient) ScanBlock(ctx context.Context, in *ScanBlockRequest, opts ...grpc.CallOption) (BlockService_ScanBlockClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlockService_ServiceDesc.Streams[1], BlockService_ScanBlock_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &blockServiceScanBlockClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockService_ScanBlockClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type blockServiceScanBlockClient struct {
	grpc.ClientStream
}

func (x *blockServiceScanBlockClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockServiceClient) FetchAndStoreBlock(ctx context.Context, in *FetchAndStoreBlockRequest, opts ...grpc.CallOption) (*FetchAndStoreBlockResponse, error) {
	out := new(FetchAndStoreBlockResponse)
	err := c.cc.Invoke(ctx, BlockService_FetchAndStoreBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockServiceServer is the server API for BlockService service.
// All implementations should embed UnimplementedBlockServiceServer
// for forward compatibility
type BlockServiceServer interface {
	// Retrieves a single block by its ID.
	GetBlock(context.Context, *GetBlockRequest) (*Block, error)
	// Retrieves a stream of blocks within a specified height range.
	GetBlocks(*GetBlocksRequest, BlockService_GetBlocksServer) error
	// Scans a range of blocks.
	ScanBlock(*ScanBlockRequest, BlockService_ScanBlockServer) error
	FetchAndStoreBlock(context.Context, *FetchAndStoreBlockRequest) (*FetchAndStoreBlockResponse, error)
}

// UnimplementedBlockServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBlockServiceServer struct {
}

func (UnimplementedBlockServiceServer) GetBlock(context.Context, *GetBlockRequest) (*Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedBlockServiceServer) GetBlocks(*GetBlocksRequest, BlockService_GetBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
func (UnimplementedBlockServiceServer) ScanBlock(*ScanBlockRequest, BlockService_ScanBlockServer) error {
	return status.Errorf(codes.Unimplemented, "method ScanBlock not implemented")
}
func (UnimplementedBlockServiceServer) FetchAndStoreBlock(context.Context, *FetchAndStoreBlockRequest) (*FetchAndStoreBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchAndStoreBlock not implemented")
}

// UnsafeBlockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockServiceServer will
// result in compilation errors.
type UnsafeBlockServiceServer interface {
	mustEmbedUnimplementedBlockServiceServer()
}

func RegisterBlockServiceServer(s grpc.ServiceRegistrar, srv BlockServiceServer) {
	s.RegisterService(&BlockService_ServiceDesc, srv)
}

func _BlockService_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockServiceServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockService_GetBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockServiceServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockService_GetBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBlocksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockServiceServer).GetBlocks(m, &blockServiceGetBlocksServer{stream})
}

type BlockService_GetBlocksServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type blockServiceGetBlocksServer struct {
	grpc.ServerStream
}

func (x *blockServiceGetBlocksServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

func _BlockService_ScanBlock_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ScanBlockRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockServiceServer).ScanBlock(m, &blockServiceScanBlockServer{stream})
}

type BlockService_ScanBlockServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type blockServiceScanBlockServer struct {
	grpc.ServerStream
}

func (x *blockServiceScanBlockServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

func _BlockService_FetchAndStoreBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAndStoreBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockServiceServer).FetchAndStoreBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockService_FetchAndStoreBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockServiceServer).FetchAndStoreBlock(ctx, req.(*FetchAndStoreBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlockService_ServiceDesc is the grpc.ServiceDesc for BlockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "block.BlockService",
	HandlerType: (*BlockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlock",
			Handler:    _BlockService_GetBlock_Handler,
		},
		{
			MethodName: "FetchAndStoreBlock",
			Handler:    _BlockService_FetchAndStoreBlock_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlocks",
			Handler:       _BlockService_GetBlocks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ScanBlock",
			Handler:       _BlockService_ScanBlock_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "entity/domain/block/model/block.proto",
}
