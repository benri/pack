// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpack/pack/commands (interfaces: BuilderInspector)

// Package mocks is a generated GoMock package.
package mocks

import (
	image "github.com/buildpack/lifecycle/image"
	pack "github.com/buildpack/pack"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBuilderInspector is a mock of BuilderInspector interface
type MockBuilderInspector struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderInspectorMockRecorder
}

// MockBuilderInspectorMockRecorder is the mock recorder for MockBuilderInspector
type MockBuilderInspectorMockRecorder struct {
	mock *MockBuilderInspector
}

// NewMockBuilderInspector creates a new mock instance
func NewMockBuilderInspector(ctrl *gomock.Controller) *MockBuilderInspector {
	mock := &MockBuilderInspector{ctrl: ctrl}
	mock.recorder = &MockBuilderInspectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuilderInspector) EXPECT() *MockBuilderInspectorMockRecorder {
	return m.recorder
}

// Inspect mocks base method
func (m *MockBuilderInspector) Inspect(arg0 image.Image) (pack.Builder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inspect", arg0)
	ret0, _ := ret[0].(pack.Builder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect
func (mr *MockBuilderInspectorMockRecorder) Inspect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*MockBuilderInspector)(nil).Inspect), arg0)
}
