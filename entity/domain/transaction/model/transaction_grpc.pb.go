// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: entity/domain/transaction/model/transaction.proto

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
	TransactionService_GetTransaction_FullMethodName    = "/transaction.TransactionService/GetTransaction"
	TransactionService_GetTransactions_FullMethodName   = "/transaction.TransactionService/GetTransactions"
	TransactionService_CreateTransaction_FullMethodName = "/transaction.TransactionService/CreateTransaction"
)

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	// Retrieves a single transaction by its ID.
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
	// Retrieves all transactions within a specific block.
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (TransactionService_GetTransactionsClient, error)
	// Creates a new transaction and returns the created transaction.
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, TransactionService_GetTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (TransactionService_GetTransactionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &TransactionService_ServiceDesc.Streams[0], TransactionService_GetTransactions_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &transactionServiceGetTransactionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TransactionService_GetTransactionsClient interface {
	Recv() (*Transaction, error)
	grpc.ClientStream
}

type transactionServiceGetTransactionsClient struct {
	grpc.ClientStream
}

func (x *transactionServiceGetTransactionsClient) Recv() (*Transaction, error) {
	m := new(Transaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transactionServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, TransactionService_CreateTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations should embed UnimplementedTransactionServiceServer
// for forward compatibility
type TransactionServiceServer interface {
	// Retrieves a single transaction by its ID.
	GetTransaction(context.Context, *GetTransactionRequest) (*Transaction, error)
	// Retrieves all transactions within a specific block.
	GetTransactions(*GetTransactionsRequest, TransactionService_GetTransactionsServer) error
	// Creates a new transaction and returns the created transaction.
	CreateTransaction(context.Context, *CreateTransactionRequest) (*Transaction, error)
}

// UnimplementedTransactionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (UnimplementedTransactionServiceServer) GetTransaction(context.Context, *GetTransactionRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) GetTransactions(*GetTransactionsRequest, TransactionService_GetTransactionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedTransactionServiceServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_GetTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTransactionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransactionServiceServer).GetTransactions(m, &transactionServiceGetTransactionsServer{stream})
}

type TransactionService_GetTransactionsServer interface {
	Send(*Transaction) error
	grpc.ServerStream
}

type transactionServiceGetTransactionsServer struct {
	grpc.ServerStream
}

func (x *transactionServiceGetTransactionsServer) Send(m *Transaction) error {
	return x.ServerStream.SendMsg(m)
}

func _TransactionService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTransaction",
			Handler:    _TransactionService_GetTransaction_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _TransactionService_CreateTransaction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTransactions",
			Handler:       _TransactionService_GetTransactions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "entity/domain/transaction/model/transaction.proto",
}
