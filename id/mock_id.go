// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/taeho-io/taeho-go/id (interfaces: ID)

// Package id is a generated GoMock package.
package id

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockID is a mock of ID interface
type MockID struct {
	ctrl     *gomock.Controller
	recorder *MockIDMockRecorder
}

// MockIDMockRecorder is the mock recorder for MockID
type MockIDMockRecorder struct {
	mock *MockID
}

// NewMockID creates a new mock instance
func NewMockID(ctrl *gomock.Controller) *MockID {
	mock := &MockID{ctrl: ctrl}
	mock.recorder = &MockIDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockID) EXPECT() *MockIDMockRecorder {
	return m.recorder
}

// Generate mocks base method
func (m *MockID) Generate() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate
func (mr *MockIDMockRecorder) Generate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockID)(nil).Generate))
}

// Must mocks base method
func (m *MockID) Must() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Must")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Must indicates an expected call of Must
func (mr *MockIDMockRecorder) Must() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Must", reflect.TypeOf((*MockID)(nil).Must))
}