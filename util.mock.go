// +build debug

// Code generated by MockGen. DO NOT EDIT.
// Source: util.go

package taskforce

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockutilI is a mock of utilI interface
type MockutilI struct {
	ctrl     *gomock.Controller
	recorder *MockutilIMockRecorder
}

// MockutilIMockRecorder is the mock recorder for MockutilI
type MockutilIMockRecorder struct {
	mock *MockutilI
}

// NewMockutilI creates a new mock instance
func NewMockutilI(ctrl *gomock.Controller) *MockutilI {
	mock := &MockutilI{ctrl: ctrl}
	mock.recorder = &MockutilIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockutilI) EXPECT() *MockutilIMockRecorder {
	return m.recorder
}

// DoPanic mocks base method
func (m *MockutilI) DoPanic(err error) {
	m.ctrl.Call(m, "DoPanic", err)
}

// DoPanic indicates an expected call of DoPanic
func (mr *MockutilIMockRecorder) DoPanic(err interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoPanic", reflect.TypeOf((*MockutilI)(nil).DoPanic), err)
}

// DoRecover mocks base method
func (m *MockutilI) DoRecover(r interface{}) {
	m.ctrl.Call(m, "DoRecover", r)
}

// DoRecover indicates an expected call of DoRecover
func (mr *MockutilIMockRecorder) DoRecover(r interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoRecover", reflect.TypeOf((*MockutilI)(nil).DoRecover), r)
}

// DisplaySuccess mocks base method
func (m *MockutilI) DisplaySuccess() {
	m.ctrl.Call(m, "DisplaySuccess")
}

// DisplaySuccess indicates an expected call of DisplaySuccess
func (mr *MockutilIMockRecorder) DisplaySuccess() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisplaySuccess", reflect.TypeOf((*MockutilI)(nil).DisplaySuccess))
}