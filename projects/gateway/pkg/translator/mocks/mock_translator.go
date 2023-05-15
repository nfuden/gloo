// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gateway/pkg/translator (interfaces: Translator)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	v10 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	gloosnapshot "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	reporter "github.com/solo-io/solo-kit/pkg/api/v2/reporter"
)

// MockTranslator is a mock of Translator interface.
type MockTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslatorMockRecorder
}

// MockTranslatorMockRecorder is the mock recorder for MockTranslator.
type MockTranslatorMockRecorder struct {
	mock *MockTranslator
}

// NewMockTranslator creates a new mock instance.
func NewMockTranslator(ctrl *gomock.Controller) *MockTranslator {
	mock := &MockTranslator{ctrl: ctrl}
	mock.recorder = &MockTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslator) EXPECT() *MockTranslatorMockRecorder {
	return m.recorder
}

// Translate mocks base method.
func (m *MockTranslator) Translate(arg0 context.Context, arg1 string, arg2 *gloosnapshot.ApiSnapshot, arg3 v1.GatewayList) (*v10.Proxy, reporter.ResourceReports) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v10.Proxy)
	ret1, _ := ret[1].(reporter.ResourceReports)
	return ret0, ret1
}

// Translate indicates an expected call of Translate.
func (mr *MockTranslatorMockRecorder) Translate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockTranslator)(nil).Translate), arg0, arg1, arg2, arg3)
}