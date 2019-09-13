package external

import (
	context "context"
	choria "github.com/choria-io/go-choria/choria"
	agents "github.com/choria-io/go-choria/server/agents"
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
	reflect "reflect"
)

// MockAgentManager is a mock of AgentManager interface
type MockAgentManager struct {
	ctrl     *gomock.Controller
	recorder *MockAgentManagerMockRecorder
}

// MockAgentManagerMockRecorder is the mock recorder for MockAgentManager
type MockAgentManagerMockRecorder struct {
	mock *MockAgentManager
}

// NewMockAgentManager creates a new mock instance
func NewMockAgentManager(ctrl *gomock.Controller) *MockAgentManager {
	mock := &MockAgentManager{ctrl: ctrl}
	mock.recorder = &MockAgentManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAgentManager) EXPECT() *MockAgentManagerMockRecorder {
	return m.recorder
}

// RegisterAgent mocks base method
func (m *MockAgentManager) RegisterAgent(ctx context.Context, name string, agent agents.Agent, conn choria.AgentConnector) error {
	ret := m.ctrl.Call(m, "RegisterAgent", ctx, name, agent, conn)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterAgent indicates an expected call of RegisterAgent
func (mr *MockAgentManagerMockRecorder) RegisterAgent(ctx, name, agent, conn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAgent", reflect.TypeOf((*MockAgentManager)(nil).RegisterAgent), ctx, name, agent, conn)
}

// Logger mocks base method
func (m *MockAgentManager) Logger() *logrus.Entry {
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

// Logger indicates an expected call of Logger
func (mr *MockAgentManagerMockRecorder) Logger() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockAgentManager)(nil).Logger))
}

// Choria mocks base method
func (m *MockAgentManager) Choria() *choria.Framework {
	ret := m.ctrl.Call(m, "Choria")
	ret0, _ := ret[0].(*choria.Framework)
	return ret0
}

// Choria indicates an expected call of Choria
func (mr *MockAgentManagerMockRecorder) Choria() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Choria", reflect.TypeOf((*MockAgentManager)(nil).Choria))
}
