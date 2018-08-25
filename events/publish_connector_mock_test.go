// Code generated by MockGen. DO NOT EDIT.
// Source: events.go

// Package events is a generated GoMock package.
package events

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPublishConnector is a mock of PublishConnector interface
type MockPublishConnector struct {
	ctrl     *gomock.Controller
	recorder *MockPublishConnectorMockRecorder
}

// MockPublishConnectorMockRecorder is the mock recorder for MockPublishConnector
type MockPublishConnectorMockRecorder struct {
	mock *MockPublishConnector
}

// NewMockPublishConnector creates a new mock instance
func NewMockPublishConnector(ctrl *gomock.Controller) *MockPublishConnector {
	mock := &MockPublishConnector{ctrl: ctrl}
	mock.recorder = &MockPublishConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublishConnector) EXPECT() *MockPublishConnectorMockRecorder {
	return m.recorder
}

// PublishRaw mocks base method
func (m *MockPublishConnector) PublishRaw(target string, data []byte) error {
	ret := m.ctrl.Call(m, "PublishRaw", target, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishRaw indicates an expected call of PublishRaw
func (mr *MockPublishConnectorMockRecorder) PublishRaw(target, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishRaw", reflect.TypeOf((*MockPublishConnector)(nil).PublishRaw), target, data)
}