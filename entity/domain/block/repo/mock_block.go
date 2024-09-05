// Code generated by MockGen. DO NOT EDIT.
// Source: block.go
//
// Generated by this command:
//
//	mockgen -destination=./mock_block.go -package=repo -source=block.go
//

// Package repo is a generated GoMock package.
package repo

import (
	context "context"
	reflect "reflect"

	model "github.com/blackhorseya/ryze/entity/domain/block/model"
	gomock "go.uber.org/mock/gomock"
)

// MockIBlockRepo is a mock of IBlockRepo interface.
type MockIBlockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIBlockRepoMockRecorder
}

// MockIBlockRepoMockRecorder is the mock recorder for MockIBlockRepo.
type MockIBlockRepoMockRecorder struct {
	mock *MockIBlockRepo
}

// NewMockIBlockRepo creates a new mock instance.
func NewMockIBlockRepo(ctrl *gomock.Controller) *MockIBlockRepo {
	mock := &MockIBlockRepo{ctrl: ctrl}
	mock.recorder = &MockIBlockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBlockRepo) EXPECT() *MockIBlockRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIBlockRepo) Create(c context.Context, item *model.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIBlockRepoMockRecorder) Create(c, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIBlockRepo)(nil).Create), c, item)
}

// GetByID mocks base method.
func (m *MockIBlockRepo) GetByID(c context.Context, id string) (*model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", c, id)
	ret0, _ := ret[0].(*model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIBlockRepoMockRecorder) GetByID(c, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIBlockRepo)(nil).GetByID), c, id)
}

// List mocks base method.
func (m *MockIBlockRepo) List(c context.Context, condition ListCondition) ([]*model.Block, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", c, condition)
	ret0, _ := ret[0].([]*model.Block)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockIBlockRepoMockRecorder) List(c, condition any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIBlockRepo)(nil).List), c, condition)
}
