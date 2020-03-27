// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_crd_uninstall is a generated GoMock package.
package mock_crd_uninstall

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	rest "k8s.io/client-go/rest"
)

// MockCrdRemover is a mock of CrdRemover interface.
type MockCrdRemover struct {
	ctrl     *gomock.Controller
	recorder *MockCrdRemoverMockRecorder
}

// MockCrdRemoverMockRecorder is the mock recorder for MockCrdRemover.
type MockCrdRemoverMockRecorder struct {
	mock *MockCrdRemover
}

// NewMockCrdRemover creates a new mock instance.
func NewMockCrdRemover(ctrl *gomock.Controller) *MockCrdRemover {
	mock := &MockCrdRemover{ctrl: ctrl}
	mock.recorder = &MockCrdRemoverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCrdRemover) EXPECT() *MockCrdRemoverMockRecorder {
	return m.recorder
}

// RemoveZephyrCrds mocks base method.
func (m *MockCrdRemover) RemoveZephyrCrds(clusterName string, remoteKubeConfig *rest.Config) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveZephyrCrds", clusterName, remoteKubeConfig)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveZephyrCrds indicates an expected call of RemoveZephyrCrds.
func (mr *MockCrdRemoverMockRecorder) RemoveZephyrCrds(clusterName, remoteKubeConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveZephyrCrds", reflect.TypeOf((*MockCrdRemover)(nil).RemoveZephyrCrds), clusterName, remoteKubeConfig)
}