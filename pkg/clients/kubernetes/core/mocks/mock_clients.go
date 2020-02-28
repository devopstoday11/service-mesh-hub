// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_kubernetes_core is a generated GoMock package.
package mock_kubernetes_core

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockServiceClient is a mock of ServiceClient interface
type MockServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockServiceClientMockRecorder
}

// MockServiceClientMockRecorder is the mock recorder for MockServiceClient
type MockServiceClientMockRecorder struct {
	mock *MockServiceClient
}

// NewMockServiceClient creates a new mock instance
func NewMockServiceClient(ctrl *gomock.Controller) *MockServiceClient {
	mock := &MockServiceClient{ctrl: ctrl}
	mock.recorder = &MockServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceClient) EXPECT() *MockServiceClientMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockServiceClient) List(ctx context.Context, options ...client.ListOption) (*v1.ServiceList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v1.ServiceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockServiceClientMockRecorder) List(ctx interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServiceClient)(nil).List), varargs...)
}

// MockSecretsClient is a mock of SecretsClient interface
type MockSecretsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsClientMockRecorder
}

// MockSecretsClientMockRecorder is the mock recorder for MockSecretsClient
type MockSecretsClientMockRecorder struct {
	mock *MockSecretsClient
}

// NewMockSecretsClient creates a new mock instance
func NewMockSecretsClient(ctrl *gomock.Controller) *MockSecretsClient {
	mock := &MockSecretsClient{ctrl: ctrl}
	mock.recorder = &MockSecretsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecretsClient) EXPECT() *MockSecretsClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockSecretsClient) Create(ctx context.Context, secret *v1.Secret, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, secret}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockSecretsClientMockRecorder) Create(ctx, secret interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, secret}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSecretsClient)(nil).Create), varargs...)
}

// Update mocks base method
func (m *MockSecretsClient) Update(ctx context.Context, secret *v1.Secret, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, secret}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockSecretsClientMockRecorder) Update(ctx, secret interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, secret}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSecretsClient)(nil).Update), varargs...)
}

// Get mocks base method
func (m *MockSecretsClient) Get(ctx context.Context, name, namespace string) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, name, namespace)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSecretsClientMockRecorder) Get(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSecretsClient)(nil).Get), ctx, name, namespace)
}

// List mocks base method
func (m *MockSecretsClient) List(ctx context.Context, opts v10.ListOptions) (*v1.SecretList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, opts)
	ret0, _ := ret[0].(*v1.SecretList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockSecretsClientMockRecorder) List(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecretsClient)(nil).List), ctx, opts)
}
