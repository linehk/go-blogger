// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/linehk/go-microservices-blogger/service/comment/rpc/model (interfaces: CommentModel)
//
// Generated by this command:
//
//	mockgen -destination=./mock_comment_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/comment/rpc/model github.com/linehk/go-microservices-blogger/service/comment/rpc/model CommentModel
//

// Package model is a generated GoMock package.
package model

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCommentModel is a mock of CommentModel interface.
type MockCommentModel struct {
	ctrl     *gomock.Controller
	recorder *MockCommentModelMockRecorder
}

// MockCommentModelMockRecorder is the mock recorder for MockCommentModel.
type MockCommentModelMockRecorder struct {
	mock *MockCommentModel
}

// NewMockCommentModel creates a new mock instance.
func NewMockCommentModel(ctrl *gomock.Controller) *MockCommentModel {
	mock := &MockCommentModel{ctrl: ctrl}
	mock.recorder = &MockCommentModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentModel) EXPECT() *MockCommentModelMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCommentModel) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCommentModelMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCommentModel)(nil).Delete), arg0, arg1)
}

// FindOne mocks base method.
func (m *MockCommentModel) FindOne(arg0 context.Context, arg1 int64) (*Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockCommentModelMockRecorder) FindOne(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockCommentModel)(nil).FindOne), arg0, arg1)
}

// FindOneByUuid mocks base method.
func (m *MockCommentModel) FindOneByUuid(arg0 context.Context, arg1 string) (*Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByUuid", arg0, arg1)
	ret0, _ := ret[0].(*Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByUuid indicates an expected call of FindOneByUuid.
func (mr *MockCommentModelMockRecorder) FindOneByUuid(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByUuid", reflect.TypeOf((*MockCommentModel)(nil).FindOneByUuid), arg0, arg1)
}

// Insert mocks base method.
func (m *MockCommentModel) Insert(arg0 context.Context, arg1 *Comment) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockCommentModelMockRecorder) Insert(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockCommentModel)(nil).Insert), arg0, arg1)
}

// ListByBlogUuidAndPostUuid mocks base method.
func (m *MockCommentModel) ListByBlogUuidAndPostUuid(arg0 context.Context, arg1, arg2 string) ([]*Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByBlogUuidAndPostUuid", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByBlogUuidAndPostUuid indicates an expected call of ListByBlogUuidAndPostUuid.
func (mr *MockCommentModelMockRecorder) ListByBlogUuidAndPostUuid(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByBlogUuidAndPostUuid", reflect.TypeOf((*MockCommentModel)(nil).ListByBlogUuidAndPostUuid), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockCommentModel) Update(arg0 context.Context, arg1 *Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCommentModelMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCommentModel)(nil).Update), arg0, arg1)
}
