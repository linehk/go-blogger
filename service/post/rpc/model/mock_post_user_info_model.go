// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/linehk/go-microservices-blogger/service/post/rpc/model (interfaces: PostUserInfoModel)
//
// Generated by this command:
//
//	mockgen -destination=./mock_post_user_info_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/post/rpc/model github.com/linehk/go-microservices-blogger/service/post/rpc/model PostUserInfoModel
//

// Package model is a generated GoMock package.
package model

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPostUserInfoModel is a mock of PostUserInfoModel interface.
type MockPostUserInfoModel struct {
	ctrl     *gomock.Controller
	recorder *MockPostUserInfoModelMockRecorder
}

// MockPostUserInfoModelMockRecorder is the mock recorder for MockPostUserInfoModel.
type MockPostUserInfoModelMockRecorder struct {
	mock *MockPostUserInfoModel
}

// NewMockPostUserInfoModel creates a new mock instance.
func NewMockPostUserInfoModel(ctrl *gomock.Controller) *MockPostUserInfoModel {
	mock := &MockPostUserInfoModel{ctrl: ctrl}
	mock.recorder = &MockPostUserInfoModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostUserInfoModel) EXPECT() *MockPostUserInfoModelMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockPostUserInfoModel) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPostUserInfoModelMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostUserInfoModel)(nil).Delete), arg0, arg1)
}

// FindOne mocks base method.
func (m *MockPostUserInfoModel) FindOne(arg0 context.Context, arg1 int64) (*PostUserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*PostUserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockPostUserInfoModelMockRecorder) FindOne(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockPostUserInfoModel)(nil).FindOne), arg0, arg1)
}

// FindOneByUserUuidAndBlogUuidAndPostUuid mocks base method.
func (m *MockPostUserInfoModel) FindOneByUserUuidAndBlogUuidAndPostUuid(arg0 context.Context, arg1, arg2, arg3 string) (*PostUserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByUserUuidAndBlogUuidAndPostUuid", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*PostUserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByUserUuidAndBlogUuidAndPostUuid indicates an expected call of FindOneByUserUuidAndBlogUuidAndPostUuid.
func (mr *MockPostUserInfoModelMockRecorder) FindOneByUserUuidAndBlogUuidAndPostUuid(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByUserUuidAndBlogUuidAndPostUuid", reflect.TypeOf((*MockPostUserInfoModel)(nil).FindOneByUserUuidAndBlogUuidAndPostUuid), arg0, arg1, arg2, arg3)
}

// FindOneByUuid mocks base method.
func (m *MockPostUserInfoModel) FindOneByUuid(arg0 context.Context, arg1 string) (*PostUserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByUuid", arg0, arg1)
	ret0, _ := ret[0].(*PostUserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByUuid indicates an expected call of FindOneByUuid.
func (mr *MockPostUserInfoModelMockRecorder) FindOneByUuid(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByUuid", reflect.TypeOf((*MockPostUserInfoModel)(nil).FindOneByUuid), arg0, arg1)
}

// Insert mocks base method.
func (m *MockPostUserInfoModel) Insert(arg0 context.Context, arg1 *PostUserInfo) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockPostUserInfoModelMockRecorder) Insert(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPostUserInfoModel)(nil).Insert), arg0, arg1)
}

// Update mocks base method.
func (m *MockPostUserInfoModel) Update(arg0 context.Context, arg1 *PostUserInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPostUserInfoModelMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostUserInfoModel)(nil).Update), arg0, arg1)
}
