// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/linehk/go-microservices-blogger/service/post/rpc/model (interfaces: PostModel)
//
// Generated by this command:
//
//	mockgen -destination=./mock_post_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/post/rpc/model github.com/linehk/go-microservices-blogger/service/post/rpc/model PostModel
//

// Package model is a generated GoMock package.
package model

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPostModel is a mock of PostModel interface.
type MockPostModel struct {
	ctrl     *gomock.Controller
	recorder *MockPostModelMockRecorder
}

// MockPostModelMockRecorder is the mock recorder for MockPostModel.
type MockPostModelMockRecorder struct {
	mock *MockPostModel
}

// NewMockPostModel creates a new mock instance.
func NewMockPostModel(ctrl *gomock.Controller) *MockPostModel {
	mock := &MockPostModel{ctrl: ctrl}
	mock.recorder = &MockPostModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostModel) EXPECT() *MockPostModelMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockPostModel) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPostModelMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostModel)(nil).Delete), arg0, arg1)
}

// FindOne mocks base method.
func (m *MockPostModel) FindOne(arg0 context.Context, arg1 int64) (*Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockPostModelMockRecorder) FindOne(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockPostModel)(nil).FindOne), arg0, arg1)
}

// FindOneByBlogUuid mocks base method.
func (m *MockPostModel) FindOneByBlogUuid(arg0 context.Context, arg1 string) (*Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByBlogUuid", arg0, arg1)
	ret0, _ := ret[0].(*Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByBlogUuid indicates an expected call of FindOneByBlogUuid.
func (mr *MockPostModelMockRecorder) FindOneByBlogUuid(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByBlogUuid", reflect.TypeOf((*MockPostModel)(nil).FindOneByBlogUuid), arg0, arg1)
}

// FindOneByUuid mocks base method.
func (m *MockPostModel) FindOneByUuid(arg0 context.Context, arg1 string) (*Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByUuid", arg0, arg1)
	ret0, _ := ret[0].(*Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByUuid indicates an expected call of FindOneByUuid.
func (mr *MockPostModelMockRecorder) FindOneByUuid(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByUuid", reflect.TypeOf((*MockPostModel)(nil).FindOneByUuid), arg0, arg1)
}

// Insert mocks base method.
func (m *MockPostModel) Insert(arg0 context.Context, arg1 *Post) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockPostModelMockRecorder) Insert(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPostModel)(nil).Insert), arg0, arg1)
}

// Update mocks base method.
func (m *MockPostModel) Update(arg0 context.Context, arg1 *Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPostModelMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostModel)(nil).Update), arg0, arg1)
}
