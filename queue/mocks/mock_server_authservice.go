// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sonofbytes/gocrawl/queue/server (interfaces: AuthService)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAuthService is a mock of AuthService interface
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// SetConnection mocks base method
func (m *MockAuthService) SetConnection(arg0 string) error {
	ret := m.ctrl.Call(m, "SetConnection", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConnection indicates an expected call of SetConnection
func (mr *MockAuthServiceMockRecorder) SetConnection(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConnection", reflect.TypeOf((*MockAuthService)(nil).SetConnection), arg0)
}

// Validate mocks base method
func (m *MockAuthService) Validate(arg0 string) error {
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockAuthServiceMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockAuthService)(nil).Validate), arg0)
}