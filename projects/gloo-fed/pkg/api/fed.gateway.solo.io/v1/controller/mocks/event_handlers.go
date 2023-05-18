// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.gateway.solo.io/v1"
	controller "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.gateway.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockFederatedGatewayEventHandler is a mock of FederatedGatewayEventHandler interface.
type MockFederatedGatewayEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedGatewayEventHandlerMockRecorder
}

// MockFederatedGatewayEventHandlerMockRecorder is the mock recorder for MockFederatedGatewayEventHandler.
type MockFederatedGatewayEventHandlerMockRecorder struct {
	mock *MockFederatedGatewayEventHandler
}

// NewMockFederatedGatewayEventHandler creates a new mock instance.
func NewMockFederatedGatewayEventHandler(ctrl *gomock.Controller) *MockFederatedGatewayEventHandler {
	mock := &MockFederatedGatewayEventHandler{ctrl: ctrl}
	mock.recorder = &MockFederatedGatewayEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedGatewayEventHandler) EXPECT() *MockFederatedGatewayEventHandlerMockRecorder {
	return m.recorder
}

// CreateFederatedGateway mocks base method.
func (m *MockFederatedGatewayEventHandler) CreateFederatedGateway(obj *v1.FederatedGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFederatedGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFederatedGateway indicates an expected call of CreateFederatedGateway.
func (mr *MockFederatedGatewayEventHandlerMockRecorder) CreateFederatedGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFederatedGateway", reflect.TypeOf((*MockFederatedGatewayEventHandler)(nil).CreateFederatedGateway), obj)
}

// DeleteFederatedGateway mocks base method.
func (m *MockFederatedGatewayEventHandler) DeleteFederatedGateway(obj *v1.FederatedGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFederatedGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFederatedGateway indicates an expected call of DeleteFederatedGateway.
func (mr *MockFederatedGatewayEventHandlerMockRecorder) DeleteFederatedGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFederatedGateway", reflect.TypeOf((*MockFederatedGatewayEventHandler)(nil).DeleteFederatedGateway), obj)
}

// GenericFederatedGateway mocks base method.
func (m *MockFederatedGatewayEventHandler) GenericFederatedGateway(obj *v1.FederatedGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericFederatedGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericFederatedGateway indicates an expected call of GenericFederatedGateway.
func (mr *MockFederatedGatewayEventHandlerMockRecorder) GenericFederatedGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericFederatedGateway", reflect.TypeOf((*MockFederatedGatewayEventHandler)(nil).GenericFederatedGateway), obj)
}

// UpdateFederatedGateway mocks base method.
func (m *MockFederatedGatewayEventHandler) UpdateFederatedGateway(old, new *v1.FederatedGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFederatedGateway", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederatedGateway indicates an expected call of UpdateFederatedGateway.
func (mr *MockFederatedGatewayEventHandlerMockRecorder) UpdateFederatedGateway(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedGateway", reflect.TypeOf((*MockFederatedGatewayEventHandler)(nil).UpdateFederatedGateway), old, new)
}

// MockFederatedGatewayEventWatcher is a mock of FederatedGatewayEventWatcher interface.
type MockFederatedGatewayEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedGatewayEventWatcherMockRecorder
}

// MockFederatedGatewayEventWatcherMockRecorder is the mock recorder for MockFederatedGatewayEventWatcher.
type MockFederatedGatewayEventWatcherMockRecorder struct {
	mock *MockFederatedGatewayEventWatcher
}

// NewMockFederatedGatewayEventWatcher creates a new mock instance.
func NewMockFederatedGatewayEventWatcher(ctrl *gomock.Controller) *MockFederatedGatewayEventWatcher {
	mock := &MockFederatedGatewayEventWatcher{ctrl: ctrl}
	mock.recorder = &MockFederatedGatewayEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedGatewayEventWatcher) EXPECT() *MockFederatedGatewayEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockFederatedGatewayEventWatcher) AddEventHandler(ctx context.Context, h controller.FederatedGatewayEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockFederatedGatewayEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockFederatedGatewayEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockFederatedMatchableHttpGatewayEventHandler is a mock of FederatedMatchableHttpGatewayEventHandler interface.
type MockFederatedMatchableHttpGatewayEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedMatchableHttpGatewayEventHandlerMockRecorder
}

// MockFederatedMatchableHttpGatewayEventHandlerMockRecorder is the mock recorder for MockFederatedMatchableHttpGatewayEventHandler.
type MockFederatedMatchableHttpGatewayEventHandlerMockRecorder struct {
	mock *MockFederatedMatchableHttpGatewayEventHandler
}

// NewMockFederatedMatchableHttpGatewayEventHandler creates a new mock instance.
func NewMockFederatedMatchableHttpGatewayEventHandler(ctrl *gomock.Controller) *MockFederatedMatchableHttpGatewayEventHandler {
	mock := &MockFederatedMatchableHttpGatewayEventHandler{ctrl: ctrl}
	mock.recorder = &MockFederatedMatchableHttpGatewayEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedMatchableHttpGatewayEventHandler) EXPECT() *MockFederatedMatchableHttpGatewayEventHandlerMockRecorder {
	return m.recorder
}

// CreateFederatedMatchableHttpGateway mocks base method.
func (m *MockFederatedMatchableHttpGatewayEventHandler) CreateFederatedMatchableHttpGateway(obj *v1.FederatedMatchableHttpGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFederatedMatchableHttpGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFederatedMatchableHttpGateway indicates an expected call of CreateFederatedMatchableHttpGateway.
func (mr *MockFederatedMatchableHttpGatewayEventHandlerMockRecorder) CreateFederatedMatchableHttpGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFederatedMatchableHttpGateway", reflect.TypeOf((*MockFederatedMatchableHttpGatewayEventHandler)(nil).CreateFederatedMatchableHttpGateway), obj)
}

// DeleteFederatedMatchableHttpGateway mocks base method.
func (m *MockFederatedMatchableHttpGatewayEventHandler) DeleteFederatedMatchableHttpGateway(obj *v1.FederatedMatchableHttpGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFederatedMatchableHttpGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFederatedMatchableHttpGateway indicates an expected call of DeleteFederatedMatchableHttpGateway.
func (mr *MockFederatedMatchableHttpGatewayEventHandlerMockRecorder) DeleteFederatedMatchableHttpGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFederatedMatchableHttpGateway", reflect.TypeOf((*MockFederatedMatchableHttpGatewayEventHandler)(nil).DeleteFederatedMatchableHttpGateway), obj)
}

// GenericFederatedMatchableHttpGateway mocks base method.
func (m *MockFederatedMatchableHttpGatewayEventHandler) GenericFederatedMatchableHttpGateway(obj *v1.FederatedMatchableHttpGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericFederatedMatchableHttpGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericFederatedMatchableHttpGateway indicates an expected call of GenericFederatedMatchableHttpGateway.
func (mr *MockFederatedMatchableHttpGatewayEventHandlerMockRecorder) GenericFederatedMatchableHttpGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericFederatedMatchableHttpGateway", reflect.TypeOf((*MockFederatedMatchableHttpGatewayEventHandler)(nil).GenericFederatedMatchableHttpGateway), obj)
}

// UpdateFederatedMatchableHttpGateway mocks base method.
func (m *MockFederatedMatchableHttpGatewayEventHandler) UpdateFederatedMatchableHttpGateway(old, new *v1.FederatedMatchableHttpGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFederatedMatchableHttpGateway", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederatedMatchableHttpGateway indicates an expected call of UpdateFederatedMatchableHttpGateway.
func (mr *MockFederatedMatchableHttpGatewayEventHandlerMockRecorder) UpdateFederatedMatchableHttpGateway(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedMatchableHttpGateway", reflect.TypeOf((*MockFederatedMatchableHttpGatewayEventHandler)(nil).UpdateFederatedMatchableHttpGateway), old, new)
}

// MockFederatedMatchableHttpGatewayEventWatcher is a mock of FederatedMatchableHttpGatewayEventWatcher interface.
type MockFederatedMatchableHttpGatewayEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedMatchableHttpGatewayEventWatcherMockRecorder
}

// MockFederatedMatchableHttpGatewayEventWatcherMockRecorder is the mock recorder for MockFederatedMatchableHttpGatewayEventWatcher.
type MockFederatedMatchableHttpGatewayEventWatcherMockRecorder struct {
	mock *MockFederatedMatchableHttpGatewayEventWatcher
}

// NewMockFederatedMatchableHttpGatewayEventWatcher creates a new mock instance.
func NewMockFederatedMatchableHttpGatewayEventWatcher(ctrl *gomock.Controller) *MockFederatedMatchableHttpGatewayEventWatcher {
	mock := &MockFederatedMatchableHttpGatewayEventWatcher{ctrl: ctrl}
	mock.recorder = &MockFederatedMatchableHttpGatewayEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedMatchableHttpGatewayEventWatcher) EXPECT() *MockFederatedMatchableHttpGatewayEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockFederatedMatchableHttpGatewayEventWatcher) AddEventHandler(ctx context.Context, h controller.FederatedMatchableHttpGatewayEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockFederatedMatchableHttpGatewayEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockFederatedMatchableHttpGatewayEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockFederatedMatchableTcpGatewayEventHandler is a mock of FederatedMatchableTcpGatewayEventHandler interface.
type MockFederatedMatchableTcpGatewayEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedMatchableTcpGatewayEventHandlerMockRecorder
}

// MockFederatedMatchableTcpGatewayEventHandlerMockRecorder is the mock recorder for MockFederatedMatchableTcpGatewayEventHandler.
type MockFederatedMatchableTcpGatewayEventHandlerMockRecorder struct {
	mock *MockFederatedMatchableTcpGatewayEventHandler
}

// NewMockFederatedMatchableTcpGatewayEventHandler creates a new mock instance.
func NewMockFederatedMatchableTcpGatewayEventHandler(ctrl *gomock.Controller) *MockFederatedMatchableTcpGatewayEventHandler {
	mock := &MockFederatedMatchableTcpGatewayEventHandler{ctrl: ctrl}
	mock.recorder = &MockFederatedMatchableTcpGatewayEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedMatchableTcpGatewayEventHandler) EXPECT() *MockFederatedMatchableTcpGatewayEventHandlerMockRecorder {
	return m.recorder
}

// CreateFederatedMatchableTcpGateway mocks base method.
func (m *MockFederatedMatchableTcpGatewayEventHandler) CreateFederatedMatchableTcpGateway(obj *v1.FederatedMatchableTcpGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFederatedMatchableTcpGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFederatedMatchableTcpGateway indicates an expected call of CreateFederatedMatchableTcpGateway.
func (mr *MockFederatedMatchableTcpGatewayEventHandlerMockRecorder) CreateFederatedMatchableTcpGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFederatedMatchableTcpGateway", reflect.TypeOf((*MockFederatedMatchableTcpGatewayEventHandler)(nil).CreateFederatedMatchableTcpGateway), obj)
}

// DeleteFederatedMatchableTcpGateway mocks base method.
func (m *MockFederatedMatchableTcpGatewayEventHandler) DeleteFederatedMatchableTcpGateway(obj *v1.FederatedMatchableTcpGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFederatedMatchableTcpGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFederatedMatchableTcpGateway indicates an expected call of DeleteFederatedMatchableTcpGateway.
func (mr *MockFederatedMatchableTcpGatewayEventHandlerMockRecorder) DeleteFederatedMatchableTcpGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFederatedMatchableTcpGateway", reflect.TypeOf((*MockFederatedMatchableTcpGatewayEventHandler)(nil).DeleteFederatedMatchableTcpGateway), obj)
}

// GenericFederatedMatchableTcpGateway mocks base method.
func (m *MockFederatedMatchableTcpGatewayEventHandler) GenericFederatedMatchableTcpGateway(obj *v1.FederatedMatchableTcpGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericFederatedMatchableTcpGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericFederatedMatchableTcpGateway indicates an expected call of GenericFederatedMatchableTcpGateway.
func (mr *MockFederatedMatchableTcpGatewayEventHandlerMockRecorder) GenericFederatedMatchableTcpGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericFederatedMatchableTcpGateway", reflect.TypeOf((*MockFederatedMatchableTcpGatewayEventHandler)(nil).GenericFederatedMatchableTcpGateway), obj)
}

// UpdateFederatedMatchableTcpGateway mocks base method.
func (m *MockFederatedMatchableTcpGatewayEventHandler) UpdateFederatedMatchableTcpGateway(old, new *v1.FederatedMatchableTcpGateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFederatedMatchableTcpGateway", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederatedMatchableTcpGateway indicates an expected call of UpdateFederatedMatchableTcpGateway.
func (mr *MockFederatedMatchableTcpGatewayEventHandlerMockRecorder) UpdateFederatedMatchableTcpGateway(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedMatchableTcpGateway", reflect.TypeOf((*MockFederatedMatchableTcpGatewayEventHandler)(nil).UpdateFederatedMatchableTcpGateway), old, new)
}

// MockFederatedMatchableTcpGatewayEventWatcher is a mock of FederatedMatchableTcpGatewayEventWatcher interface.
type MockFederatedMatchableTcpGatewayEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedMatchableTcpGatewayEventWatcherMockRecorder
}

// MockFederatedMatchableTcpGatewayEventWatcherMockRecorder is the mock recorder for MockFederatedMatchableTcpGatewayEventWatcher.
type MockFederatedMatchableTcpGatewayEventWatcherMockRecorder struct {
	mock *MockFederatedMatchableTcpGatewayEventWatcher
}

// NewMockFederatedMatchableTcpGatewayEventWatcher creates a new mock instance.
func NewMockFederatedMatchableTcpGatewayEventWatcher(ctrl *gomock.Controller) *MockFederatedMatchableTcpGatewayEventWatcher {
	mock := &MockFederatedMatchableTcpGatewayEventWatcher{ctrl: ctrl}
	mock.recorder = &MockFederatedMatchableTcpGatewayEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedMatchableTcpGatewayEventWatcher) EXPECT() *MockFederatedMatchableTcpGatewayEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockFederatedMatchableTcpGatewayEventWatcher) AddEventHandler(ctx context.Context, h controller.FederatedMatchableTcpGatewayEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockFederatedMatchableTcpGatewayEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockFederatedMatchableTcpGatewayEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockFederatedVirtualServiceEventHandler is a mock of FederatedVirtualServiceEventHandler interface.
type MockFederatedVirtualServiceEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedVirtualServiceEventHandlerMockRecorder
}

// MockFederatedVirtualServiceEventHandlerMockRecorder is the mock recorder for MockFederatedVirtualServiceEventHandler.
type MockFederatedVirtualServiceEventHandlerMockRecorder struct {
	mock *MockFederatedVirtualServiceEventHandler
}

// NewMockFederatedVirtualServiceEventHandler creates a new mock instance.
func NewMockFederatedVirtualServiceEventHandler(ctrl *gomock.Controller) *MockFederatedVirtualServiceEventHandler {
	mock := &MockFederatedVirtualServiceEventHandler{ctrl: ctrl}
	mock.recorder = &MockFederatedVirtualServiceEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedVirtualServiceEventHandler) EXPECT() *MockFederatedVirtualServiceEventHandlerMockRecorder {
	return m.recorder
}

// CreateFederatedVirtualService mocks base method.
func (m *MockFederatedVirtualServiceEventHandler) CreateFederatedVirtualService(obj *v1.FederatedVirtualService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFederatedVirtualService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFederatedVirtualService indicates an expected call of CreateFederatedVirtualService.
func (mr *MockFederatedVirtualServiceEventHandlerMockRecorder) CreateFederatedVirtualService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFederatedVirtualService", reflect.TypeOf((*MockFederatedVirtualServiceEventHandler)(nil).CreateFederatedVirtualService), obj)
}

// DeleteFederatedVirtualService mocks base method.
func (m *MockFederatedVirtualServiceEventHandler) DeleteFederatedVirtualService(obj *v1.FederatedVirtualService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFederatedVirtualService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFederatedVirtualService indicates an expected call of DeleteFederatedVirtualService.
func (mr *MockFederatedVirtualServiceEventHandlerMockRecorder) DeleteFederatedVirtualService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFederatedVirtualService", reflect.TypeOf((*MockFederatedVirtualServiceEventHandler)(nil).DeleteFederatedVirtualService), obj)
}

// GenericFederatedVirtualService mocks base method.
func (m *MockFederatedVirtualServiceEventHandler) GenericFederatedVirtualService(obj *v1.FederatedVirtualService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericFederatedVirtualService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericFederatedVirtualService indicates an expected call of GenericFederatedVirtualService.
func (mr *MockFederatedVirtualServiceEventHandlerMockRecorder) GenericFederatedVirtualService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericFederatedVirtualService", reflect.TypeOf((*MockFederatedVirtualServiceEventHandler)(nil).GenericFederatedVirtualService), obj)
}

// UpdateFederatedVirtualService mocks base method.
func (m *MockFederatedVirtualServiceEventHandler) UpdateFederatedVirtualService(old, new *v1.FederatedVirtualService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFederatedVirtualService", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederatedVirtualService indicates an expected call of UpdateFederatedVirtualService.
func (mr *MockFederatedVirtualServiceEventHandlerMockRecorder) UpdateFederatedVirtualService(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedVirtualService", reflect.TypeOf((*MockFederatedVirtualServiceEventHandler)(nil).UpdateFederatedVirtualService), old, new)
}

// MockFederatedVirtualServiceEventWatcher is a mock of FederatedVirtualServiceEventWatcher interface.
type MockFederatedVirtualServiceEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedVirtualServiceEventWatcherMockRecorder
}

// MockFederatedVirtualServiceEventWatcherMockRecorder is the mock recorder for MockFederatedVirtualServiceEventWatcher.
type MockFederatedVirtualServiceEventWatcherMockRecorder struct {
	mock *MockFederatedVirtualServiceEventWatcher
}

// NewMockFederatedVirtualServiceEventWatcher creates a new mock instance.
func NewMockFederatedVirtualServiceEventWatcher(ctrl *gomock.Controller) *MockFederatedVirtualServiceEventWatcher {
	mock := &MockFederatedVirtualServiceEventWatcher{ctrl: ctrl}
	mock.recorder = &MockFederatedVirtualServiceEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedVirtualServiceEventWatcher) EXPECT() *MockFederatedVirtualServiceEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockFederatedVirtualServiceEventWatcher) AddEventHandler(ctx context.Context, h controller.FederatedVirtualServiceEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockFederatedVirtualServiceEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockFederatedVirtualServiceEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockFederatedRouteTableEventHandler is a mock of FederatedRouteTableEventHandler interface.
type MockFederatedRouteTableEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedRouteTableEventHandlerMockRecorder
}

// MockFederatedRouteTableEventHandlerMockRecorder is the mock recorder for MockFederatedRouteTableEventHandler.
type MockFederatedRouteTableEventHandlerMockRecorder struct {
	mock *MockFederatedRouteTableEventHandler
}

// NewMockFederatedRouteTableEventHandler creates a new mock instance.
func NewMockFederatedRouteTableEventHandler(ctrl *gomock.Controller) *MockFederatedRouteTableEventHandler {
	mock := &MockFederatedRouteTableEventHandler{ctrl: ctrl}
	mock.recorder = &MockFederatedRouteTableEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedRouteTableEventHandler) EXPECT() *MockFederatedRouteTableEventHandlerMockRecorder {
	return m.recorder
}

// CreateFederatedRouteTable mocks base method.
func (m *MockFederatedRouteTableEventHandler) CreateFederatedRouteTable(obj *v1.FederatedRouteTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFederatedRouteTable", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFederatedRouteTable indicates an expected call of CreateFederatedRouteTable.
func (mr *MockFederatedRouteTableEventHandlerMockRecorder) CreateFederatedRouteTable(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFederatedRouteTable", reflect.TypeOf((*MockFederatedRouteTableEventHandler)(nil).CreateFederatedRouteTable), obj)
}

// DeleteFederatedRouteTable mocks base method.
func (m *MockFederatedRouteTableEventHandler) DeleteFederatedRouteTable(obj *v1.FederatedRouteTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFederatedRouteTable", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFederatedRouteTable indicates an expected call of DeleteFederatedRouteTable.
func (mr *MockFederatedRouteTableEventHandlerMockRecorder) DeleteFederatedRouteTable(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFederatedRouteTable", reflect.TypeOf((*MockFederatedRouteTableEventHandler)(nil).DeleteFederatedRouteTable), obj)
}

// GenericFederatedRouteTable mocks base method.
func (m *MockFederatedRouteTableEventHandler) GenericFederatedRouteTable(obj *v1.FederatedRouteTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericFederatedRouteTable", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericFederatedRouteTable indicates an expected call of GenericFederatedRouteTable.
func (mr *MockFederatedRouteTableEventHandlerMockRecorder) GenericFederatedRouteTable(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericFederatedRouteTable", reflect.TypeOf((*MockFederatedRouteTableEventHandler)(nil).GenericFederatedRouteTable), obj)
}

// UpdateFederatedRouteTable mocks base method.
func (m *MockFederatedRouteTableEventHandler) UpdateFederatedRouteTable(old, new *v1.FederatedRouteTable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFederatedRouteTable", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederatedRouteTable indicates an expected call of UpdateFederatedRouteTable.
func (mr *MockFederatedRouteTableEventHandlerMockRecorder) UpdateFederatedRouteTable(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedRouteTable", reflect.TypeOf((*MockFederatedRouteTableEventHandler)(nil).UpdateFederatedRouteTable), old, new)
}

// MockFederatedRouteTableEventWatcher is a mock of FederatedRouteTableEventWatcher interface.
type MockFederatedRouteTableEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedRouteTableEventWatcherMockRecorder
}

// MockFederatedRouteTableEventWatcherMockRecorder is the mock recorder for MockFederatedRouteTableEventWatcher.
type MockFederatedRouteTableEventWatcherMockRecorder struct {
	mock *MockFederatedRouteTableEventWatcher
}

// NewMockFederatedRouteTableEventWatcher creates a new mock instance.
func NewMockFederatedRouteTableEventWatcher(ctrl *gomock.Controller) *MockFederatedRouteTableEventWatcher {
	mock := &MockFederatedRouteTableEventWatcher{ctrl: ctrl}
	mock.recorder = &MockFederatedRouteTableEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedRouteTableEventWatcher) EXPECT() *MockFederatedRouteTableEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockFederatedRouteTableEventWatcher) AddEventHandler(ctx context.Context, h controller.FederatedRouteTableEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockFederatedRouteTableEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockFederatedRouteTableEventWatcher)(nil).AddEventHandler), varargs...)
}
