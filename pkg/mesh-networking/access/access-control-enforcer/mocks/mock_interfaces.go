// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_access_policy_enforcer is a generated GoMock package.
package mock_access_policy_enforcer

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccessPolicyEnforcerLoop is a mock of AccessPolicyEnforcerLoop interface.
type MockAccessPolicyEnforcerLoop struct {
	ctrl     *gomock.Controller
	recorder *MockAccessPolicyEnforcerLoopMockRecorder
}

// MockAccessPolicyEnforcerLoopMockRecorder is the mock recorder for MockAccessPolicyEnforcerLoop.
type MockAccessPolicyEnforcerLoopMockRecorder struct {
	mock *MockAccessPolicyEnforcerLoop
}

// NewMockAccessPolicyEnforcerLoop creates a new mock instance.
func NewMockAccessPolicyEnforcerLoop(ctrl *gomock.Controller) *MockAccessPolicyEnforcerLoop {
	mock := &MockAccessPolicyEnforcerLoop{ctrl: ctrl}
	mock.recorder = &MockAccessPolicyEnforcerLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessPolicyEnforcerLoop) EXPECT() *MockAccessPolicyEnforcerLoopMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockAccessPolicyEnforcerLoop) Start(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockAccessPolicyEnforcerLoopMockRecorder) Start(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockAccessPolicyEnforcerLoop)(nil).Start), ctx)
}