// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_config is a generated GoMock package.
package mock_config

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWebhookConfigManager is a mock of WebhookConfigManager interface.
type MockWebhookConfigManager struct {
	ctrl     *gomock.Controller
	recorder *MockWebhookConfigManagerMockRecorder
}

// MockWebhookConfigManagerMockRecorder is the mock recorder for MockWebhookConfigManager.
type MockWebhookConfigManagerMockRecorder struct {
	mock *MockWebhookConfigManager
}

// NewMockWebhookConfigManager creates a new mock instance.
func NewMockWebhookConfigManager(ctrl *gomock.Controller) *MockWebhookConfigManager {
	mock := &MockWebhookConfigManager{ctrl: ctrl}
	mock.recorder = &MockWebhookConfigManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebhookConfigManager) EXPECT() *MockWebhookConfigManagerMockRecorder {
	return m.recorder
}

// EnsureCaCerts mocks base method.
func (m *MockWebhookConfigManager) EnsureCaCerts(ctx context.Context, secretName, secretNamespace, svcName, svcNamespace string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureCaCerts", ctx, secretName, secretNamespace, svcName, svcNamespace)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureCaCerts indicates an expected call of EnsureCaCerts.
func (mr *MockWebhookConfigManagerMockRecorder) EnsureCaCerts(ctx, secretName, secretNamespace, svcName, svcNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureCaCerts", reflect.TypeOf((*MockWebhookConfigManager)(nil).EnsureCaCerts), ctx, secretName, secretNamespace, svcName, svcNamespace)
}

// EnsureValidatingWebhookConfiguration mocks base method.
func (m *MockWebhookConfigManager) EnsureValidatingWebhookConfiguration(ctx context.Context, caBundle []byte, svcName, svcNamespace, vwcName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureValidatingWebhookConfiguration", ctx, caBundle, svcName, svcNamespace, vwcName)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureValidatingWebhookConfiguration indicates an expected call of EnsureValidatingWebhookConfiguration.
func (mr *MockWebhookConfigManagerMockRecorder) EnsureValidatingWebhookConfiguration(ctx, caBundle, svcName, svcNamespace, vwcName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureValidatingWebhookConfiguration", reflect.TypeOf((*MockWebhookConfigManager)(nil).EnsureValidatingWebhookConfiguration), ctx, caBundle, svcName, svcNamespace, vwcName)
}