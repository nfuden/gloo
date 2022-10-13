// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/multicluster.solo.io/v1alpha1"
	v1 "k8s.io/api/authentication/v1"
	admission "sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// MockMultiClusterAdmissionValidator is a mock of MultiClusterAdmissionValidator interface.
type MockMultiClusterAdmissionValidator struct {
	ctrl     *gomock.Controller
	recorder *MockMultiClusterAdmissionValidatorMockRecorder
}

// MockMultiClusterAdmissionValidatorMockRecorder is the mock recorder for MockMultiClusterAdmissionValidator.
type MockMultiClusterAdmissionValidatorMockRecorder struct {
	mock *MockMultiClusterAdmissionValidator
}

// NewMockMultiClusterAdmissionValidator creates a new mock instance.
func NewMockMultiClusterAdmissionValidator(ctrl *gomock.Controller) *MockMultiClusterAdmissionValidator {
	mock := &MockMultiClusterAdmissionValidator{ctrl: ctrl}
	mock.recorder = &MockMultiClusterAdmissionValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMultiClusterAdmissionValidator) EXPECT() *MockMultiClusterAdmissionValidatorMockRecorder {
	return m.recorder
}

// ActionIsAllowed mocks base method.
func (m *MockMultiClusterAdmissionValidator) ActionIsAllowed(ctx context.Context, roleBinding *v1alpha1.MultiClusterRoleBinding, req *admission.Request) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionIsAllowed", ctx, roleBinding, req)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionIsAllowed indicates an expected call of ActionIsAllowed.
func (mr *MockMultiClusterAdmissionValidatorMockRecorder) ActionIsAllowed(ctx, roleBinding, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionIsAllowed", reflect.TypeOf((*MockMultiClusterAdmissionValidator)(nil).ActionIsAllowed), ctx, roleBinding, req)
}

// GetMatchingMultiClusterRoleBindings mocks base method.
func (m *MockMultiClusterAdmissionValidator) GetMatchingMultiClusterRoleBindings(ctx context.Context, userInfo v1.UserInfo) ([]*v1alpha1.MultiClusterRoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchingMultiClusterRoleBindings", ctx, userInfo)
	ret0, _ := ret[0].([]*v1alpha1.MultiClusterRoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchingMultiClusterRoleBindings indicates an expected call of GetMatchingMultiClusterRoleBindings.
func (mr *MockMultiClusterAdmissionValidatorMockRecorder) GetMatchingMultiClusterRoleBindings(ctx, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchingMultiClusterRoleBindings", reflect.TypeOf((*MockMultiClusterAdmissionValidator)(nil).GetMatchingMultiClusterRoleBindings), ctx, userInfo)
}
