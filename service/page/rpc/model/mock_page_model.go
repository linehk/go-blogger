// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/linehk/go-microservices-blogger/service/page/rpc/model (interfaces: PageModel)
//
// Generated by this command:
//
//	mockgen -destination=./mock_page_model.go -package=model -self_package=github.com/linehk/go-microservices-blogger/service/page/rpc/model github.com/linehk/go-microservices-blogger/service/page/rpc/model PageModel
//

// Package model is a generated GoMock package.
package model

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPageModel is a mock of PageModel interface.
type MockPageModel struct {
	ctrl     *gomock.Controller
	recorder *MockPageModelMockRecorder
}

// MockPageModelMockRecorder is the mock recorder for MockPageModel.
type MockPageModelMockRecorder struct {
	mock *MockPageModel
}

// NewMockPageModel creates a new mock instance.
func NewMockPageModel(ctrl *gomock.Controller) *MockPageModel {
	mock := &MockPageModel{ctrl: ctrl}
	mock.recorder = &MockPageModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPageModel) EXPECT() *MockPageModelMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockPageModel) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPageModelMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPageModel)(nil).Delete), arg0, arg1)
}

// FindOne mocks base method.
func (m *MockPageModel) FindOne(arg0 context.Context, arg1 int64) (*Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*Page)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockPageModelMockRecorder) FindOne(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockPageModel)(nil).FindOne), arg0, arg1)
}

// FindOneByBlogUuid mocks base method.
func (m *MockPageModel) FindOneByBlogUuid(arg0 context.Context, arg1 string) (*Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByBlogUuid", arg0, arg1)
	ret0, _ := ret[0].(*Page)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByBlogUuid indicates an expected call of FindOneByBlogUuid.
func (mr *MockPageModelMockRecorder) FindOneByBlogUuid(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByBlogUuid", reflect.TypeOf((*MockPageModel)(nil).FindOneByBlogUuid), arg0, arg1)
}

// FindOneByUuid mocks base method.
func (m *MockPageModel) FindOneByUuid(arg0 context.Context, arg1 string) (*Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByUuid", arg0, arg1)
	ret0, _ := ret[0].(*Page)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByUuid indicates an expected call of FindOneByUuid.
func (mr *MockPageModelMockRecorder) FindOneByUuid(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByUuid", reflect.TypeOf((*MockPageModel)(nil).FindOneByUuid), arg0, arg1)
}

// Insert mocks base method.
func (m *MockPageModel) Insert(arg0 context.Context, arg1 *Page) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockPageModelMockRecorder) Insert(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPageModel)(nil).Insert), arg0, arg1)
}

// Update mocks base method.
func (m *MockPageModel) Update(arg0 context.Context, arg1 *Page) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPageModelMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPageModel)(nil).Update), arg0, arg1)
}
