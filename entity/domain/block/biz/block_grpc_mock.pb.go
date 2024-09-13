// Code generated by protoc-gen-go-grpc-mock. DO NOT EDIT.
// source: domain/block/biz/block.proto

package biz

import (
	context "context"
	reflect "reflect"

	model "github.com/blackhorseya/ryze/entity/domain/block/model"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockBlockService_ScanBlockClient is a mock of BlockService_ScanBlockClient interface.
type MockBlockService_ScanBlockClient struct {
	ctrl     *gomock.Controller
	recorder *MockBlockService_ScanBlockClientMockRecorder
}

// MockBlockService_ScanBlockClientMockRecorder is the mock recorder for MockBlockService_ScanBlockClient.
type MockBlockService_ScanBlockClientMockRecorder struct {
	mock *MockBlockService_ScanBlockClient
}

// NewMockBlockService_ScanBlockClient creates a new mock instance.
func NewMockBlockService_ScanBlockClient(ctrl *gomock.Controller) *MockBlockService_ScanBlockClient {
	mock := &MockBlockService_ScanBlockClient{ctrl: ctrl}
	mock.recorder = &MockBlockService_ScanBlockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockService_ScanBlockClient) EXPECT() *MockBlockService_ScanBlockClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockBlockService_ScanBlockClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockBlockService_ScanBlockClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockBlockService_ScanBlockClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockBlockService_ScanBlockClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockBlockService_ScanBlockClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBlockService_ScanBlockClient)(nil).Context))
}

// Header mocks base method.
func (m *MockBlockService_ScanBlockClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockBlockService_ScanBlockClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockBlockService_ScanBlockClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockBlockService_ScanBlockClient) Recv() (*model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockBlockService_ScanBlockClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockBlockService_ScanBlockClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockBlockService_ScanBlockClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockBlockService_ScanBlockClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBlockService_ScanBlockClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockBlockService_ScanBlockClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockBlockService_ScanBlockClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBlockService_ScanBlockClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockBlockService_ScanBlockClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockBlockService_ScanBlockClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockBlockService_ScanBlockClient)(nil).Trailer))
}

// MockBlockService_ScanBlockServer is a mock of BlockService_ScanBlockServer interface.
type MockBlockService_ScanBlockServer struct {
	ctrl     *gomock.Controller
	recorder *MockBlockService_ScanBlockServerMockRecorder
}

// MockBlockService_ScanBlockServerMockRecorder is the mock recorder for MockBlockService_ScanBlockServer.
type MockBlockService_ScanBlockServerMockRecorder struct {
	mock *MockBlockService_ScanBlockServer
}

// NewMockBlockService_ScanBlockServer creates a new mock instance.
func NewMockBlockService_ScanBlockServer(ctrl *gomock.Controller) *MockBlockService_ScanBlockServer {
	mock := &MockBlockService_ScanBlockServer{ctrl: ctrl}
	mock.recorder = &MockBlockService_ScanBlockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockService_ScanBlockServer) EXPECT() *MockBlockService_ScanBlockServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockBlockService_ScanBlockServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockBlockService_ScanBlockServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBlockService_ScanBlockServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockBlockService_ScanBlockServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockBlockService_ScanBlockServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBlockService_ScanBlockServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockBlockService_ScanBlockServer) Send(arg0 *model.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockBlockService_ScanBlockServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBlockService_ScanBlockServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockBlockService_ScanBlockServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockBlockService_ScanBlockServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockBlockService_ScanBlockServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockBlockService_ScanBlockServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockBlockService_ScanBlockServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBlockService_ScanBlockServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockBlockService_ScanBlockServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockBlockService_ScanBlockServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockBlockService_ScanBlockServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockBlockService_ScanBlockServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockBlockService_ScanBlockServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockBlockService_ScanBlockServer)(nil).SetTrailer), arg0)
}

// MockBlockServiceClient is a mock of BlockServiceClient interface.
type MockBlockServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockBlockServiceClientMockRecorder
}

// MockBlockServiceClientMockRecorder is the mock recorder for MockBlockServiceClient.
type MockBlockServiceClientMockRecorder struct {
	mock *MockBlockServiceClient
}

// NewMockBlockServiceClient creates a new mock instance.
func NewMockBlockServiceClient(ctrl *gomock.Controller) *MockBlockServiceClient {
	mock := &MockBlockServiceClient{ctrl: ctrl}
	mock.recorder = &MockBlockServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockServiceClient) EXPECT() *MockBlockServiceClientMockRecorder {
	return m.recorder
}

// FoundNewBlock mocks base method.
func (m *MockBlockServiceClient) FoundNewBlock(ctx context.Context, in *FoundNewBlockRequest, opts ...grpc.CallOption) (*model.Block, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FoundNewBlock", varargs...)
	ret0, _ := ret[0].(*model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FoundNewBlock indicates an expected call of FoundNewBlock.
func (mr *MockBlockServiceClientMockRecorder) FoundNewBlock(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FoundNewBlock", reflect.TypeOf((*MockBlockServiceClient)(nil).FoundNewBlock), varargs...)
}

// ScanBlock mocks base method.
func (m *MockBlockServiceClient) ScanBlock(ctx context.Context, in *ScanBlockRequest, opts ...grpc.CallOption) (BlockService_ScanBlockClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanBlock", varargs...)
	ret0, _ := ret[0].(BlockService_ScanBlockClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScanBlock indicates an expected call of ScanBlock.
func (mr *MockBlockServiceClientMockRecorder) ScanBlock(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanBlock", reflect.TypeOf((*MockBlockServiceClient)(nil).ScanBlock), varargs...)
}

// MockBlockServiceServer is a mock of BlockServiceServer interface.
type MockBlockServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockBlockServiceServerMockRecorder
}

// MockBlockServiceServerMockRecorder is the mock recorder for MockBlockServiceServer.
type MockBlockServiceServerMockRecorder struct {
	mock *MockBlockServiceServer
}

// NewMockBlockServiceServer creates a new mock instance.
func NewMockBlockServiceServer(ctrl *gomock.Controller) *MockBlockServiceServer {
	mock := &MockBlockServiceServer{ctrl: ctrl}
	mock.recorder = &MockBlockServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockServiceServer) EXPECT() *MockBlockServiceServerMockRecorder {
	return m.recorder
}

// FoundNewBlock mocks base method.
func (m *MockBlockServiceServer) FoundNewBlock(ctx context.Context, in *FoundNewBlockRequest) (*model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FoundNewBlock", ctx, in)
	ret0, _ := ret[0].(*model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FoundNewBlock indicates an expected call of FoundNewBlock.
func (mr *MockBlockServiceServerMockRecorder) FoundNewBlock(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FoundNewBlock", reflect.TypeOf((*MockBlockServiceServer)(nil).FoundNewBlock), ctx, in)
}

// ScanBlock mocks base method.
func (m *MockBlockServiceServer) ScanBlock(blob *ScanBlockRequest, server BlockService_ScanBlockServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanBlock", blob, server)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanBlock indicates an expected call of ScanBlock.
func (mr *MockBlockServiceServerMockRecorder) ScanBlock(blob, server interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanBlock", reflect.TypeOf((*MockBlockServiceServer)(nil).ScanBlock), blob, server)
}
