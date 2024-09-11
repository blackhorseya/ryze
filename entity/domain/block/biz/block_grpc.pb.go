// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: domain/block/biz/block.proto

package biz

import (
	context "context"
	model "github.com/blackhorseya/ryze/entity/domain/block/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BlockService_GetBlock_FullMethodName      = "/block.BlockService/GetBlock"
	BlockService_GetBlocks_FullMethodName     = "/block.BlockService/GetBlocks"
	BlockService_ScanBlock_FullMethodName     = "/block.BlockService/ScanBlock"
	BlockService_FoundNewBlock_FullMethodName = "/block.BlockService/FoundNewBlock"
)

// BlockServiceClient is the client API for BlockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service definition for handling blocks.
type BlockServiceClient interface {
	// Retrieves a single block by its ID.
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*model.Block, error)
	// Retrieves a stream of blocks within a specified height range.
	GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[model.Block], error)
	// Scans a range of blocks.
	ScanBlock(ctx context.Context, in *ScanBlockRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[model.Block], error)
	FoundNewBlock(ctx context.Context, in *FoundNewBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type blockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockServiceClient(cc grpc.ClientConnInterface) BlockServiceClient {
	return &blockServiceClient{cc}
}

func (c *blockServiceClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*model.Block, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(model.Block)
	err := c.cc.Invoke(ctx, BlockService_GetBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockServiceClient) GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[model.Block], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BlockService_ServiceDesc.Streams[0], BlockService_GetBlocks_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetBlocksRequest, model.Block]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BlockService_GetBlocksClient = grpc.ServerStreamingClient[model.Block]

func (c *blockServiceClient) ScanBlock(ctx context.Context, in *ScanBlockRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[model.Block], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BlockService_ServiceDesc.Streams[1], BlockService_ScanBlock_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ScanBlockRequest, model.Block]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BlockService_ScanBlockClient = grpc.ServerStreamingClient[model.Block]

func (c *blockServiceClient) FoundNewBlock(ctx context.Context, in *FoundNewBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BlockService_FoundNewBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockServiceServer is the server API for BlockService service.
// All implementations should embed UnimplementedBlockServiceServer
// for forward compatibility.
//
// Service definition for handling blocks.
type BlockServiceServer interface {
	// Retrieves a single block by its ID.
	GetBlock(context.Context, *GetBlockRequest) (*model.Block, error)
	// Retrieves a stream of blocks within a specified height range.
	GetBlocks(*GetBlocksRequest, grpc.ServerStreamingServer[model.Block]) error
	// Scans a range of blocks.
	ScanBlock(*ScanBlockRequest, grpc.ServerStreamingServer[model.Block]) error
	FoundNewBlock(context.Context, *FoundNewBlockRequest) (*emptypb.Empty, error)
}

// UnimplementedBlockServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBlockServiceServer struct{}

func (UnimplementedBlockServiceServer) GetBlock(context.Context, *GetBlockRequest) (*model.Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedBlockServiceServer) GetBlocks(*GetBlocksRequest, grpc.ServerStreamingServer[model.Block]) error {
	return status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
func (UnimplementedBlockServiceServer) ScanBlock(*ScanBlockRequest, grpc.ServerStreamingServer[model.Block]) error {
	return status.Errorf(codes.Unimplemented, "method ScanBlock not implemented")
}
func (UnimplementedBlockServiceServer) FoundNewBlock(context.Context, *FoundNewBlockRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FoundNewBlock not implemented")
}
func (UnimplementedBlockServiceServer) testEmbeddedByValue() {}

// UnsafeBlockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockServiceServer will
// result in compilation errors.
type UnsafeBlockServiceServer interface {
	mustEmbedUnimplementedBlockServiceServer()
}

func RegisterBlockServiceServer(s grpc.ServiceRegistrar, srv BlockServiceServer) {
	// If the following call pancis, it indicates UnimplementedBlockServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
	return srv.(BlockServiceServer).GetBlocks(m, &grpc.GenericServerStream[GetBlocksRequest, model.Block]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BlockService_GetBlocksServer = grpc.ServerStreamingServer[model.Block]

func _BlockService_ScanBlock_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ScanBlockRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockServiceServer).ScanBlock(m, &grpc.GenericServerStream[ScanBlockRequest, model.Block]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BlockService_ScanBlockServer = grpc.ServerStreamingServer[model.Block]

func _BlockService_FoundNewBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FoundNewBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockServiceServer).FoundNewBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockService_FoundNewBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockServiceServer).FoundNewBlock(ctx, req.(*FoundNewBlockRequest))
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
			MethodName: "FoundNewBlock",
			Handler:    _BlockService_FoundNewBlock_Handler,
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
	Metadata: "domain/block/biz/block.proto",
}
