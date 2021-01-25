// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1sets is a generated GoMock package.
package mock_v1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.gloo.solo.io/v1"
	v1sets "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.gloo.solo.io/v1/sets"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
)

// MockFederatedUpstreamSet is a mock of FederatedUpstreamSet interface
type MockFederatedUpstreamSet struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedUpstreamSetMockRecorder
}

// MockFederatedUpstreamSetMockRecorder is the mock recorder for MockFederatedUpstreamSet
type MockFederatedUpstreamSetMockRecorder struct {
	mock *MockFederatedUpstreamSet
}

// NewMockFederatedUpstreamSet creates a new mock instance
func NewMockFederatedUpstreamSet(ctrl *gomock.Controller) *MockFederatedUpstreamSet {
	mock := &MockFederatedUpstreamSet{ctrl: ctrl}
	mock.recorder = &MockFederatedUpstreamSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedUpstreamSet) EXPECT() *MockFederatedUpstreamSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockFederatedUpstreamSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockFederatedUpstreamSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockFederatedUpstreamSet)(nil).Keys))
}

// List mocks base method
func (m *MockFederatedUpstreamSet) List(filterResource ...func(*v1.FederatedUpstream) bool) []*v1.FederatedUpstream {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.FederatedUpstream)
	return ret0
}

// List indicates an expected call of List
func (mr *MockFederatedUpstreamSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFederatedUpstreamSet)(nil).List), filterResource...)
}

// Map mocks base method
func (m *MockFederatedUpstreamSet) Map() map[string]*v1.FederatedUpstream {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.FederatedUpstream)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockFederatedUpstreamSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockFederatedUpstreamSet)(nil).Map))
}

// Insert mocks base method
func (m *MockFederatedUpstreamSet) Insert(federatedUpstream ...*v1.FederatedUpstream) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range federatedUpstream {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockFederatedUpstreamSetMockRecorder) Insert(federatedUpstream ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockFederatedUpstreamSet)(nil).Insert), federatedUpstream...)
}

// Equal mocks base method
func (m *MockFederatedUpstreamSet) Equal(federatedUpstreamSet v1sets.FederatedUpstreamSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", federatedUpstreamSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockFederatedUpstreamSetMockRecorder) Equal(federatedUpstreamSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockFederatedUpstreamSet)(nil).Equal), federatedUpstreamSet)
}

// Has mocks base method
func (m *MockFederatedUpstreamSet) Has(federatedUpstream ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", federatedUpstream)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockFederatedUpstreamSetMockRecorder) Has(federatedUpstream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockFederatedUpstreamSet)(nil).Has), federatedUpstream)
}

// Delete mocks base method
func (m *MockFederatedUpstreamSet) Delete(federatedUpstream ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", federatedUpstream)
}

// Delete indicates an expected call of Delete
func (mr *MockFederatedUpstreamSetMockRecorder) Delete(federatedUpstream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFederatedUpstreamSet)(nil).Delete), federatedUpstream)
}

// Union mocks base method
func (m *MockFederatedUpstreamSet) Union(set v1sets.FederatedUpstreamSet) v1sets.FederatedUpstreamSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.FederatedUpstreamSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockFederatedUpstreamSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockFederatedUpstreamSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockFederatedUpstreamSet) Difference(set v1sets.FederatedUpstreamSet) v1sets.FederatedUpstreamSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.FederatedUpstreamSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockFederatedUpstreamSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockFederatedUpstreamSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockFederatedUpstreamSet) Intersection(set v1sets.FederatedUpstreamSet) v1sets.FederatedUpstreamSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.FederatedUpstreamSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockFederatedUpstreamSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockFederatedUpstreamSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockFederatedUpstreamSet) Find(id ezkube.ResourceId) (*v1.FederatedUpstream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.FederatedUpstream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockFederatedUpstreamSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockFederatedUpstreamSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockFederatedUpstreamSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockFederatedUpstreamSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockFederatedUpstreamSet)(nil).Length))
}

// Generic mocks base method
func (m *MockFederatedUpstreamSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockFederatedUpstreamSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockFederatedUpstreamSet)(nil).Generic))
}

// Delta mocks base method
func (m *MockFederatedUpstreamSet) Delta(newSet v1sets.FederatedUpstreamSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta
func (mr *MockFederatedUpstreamSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockFederatedUpstreamSet)(nil).Delta), newSet)
}

// MockFederatedUpstreamGroupSet is a mock of FederatedUpstreamGroupSet interface
type MockFederatedUpstreamGroupSet struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedUpstreamGroupSetMockRecorder
}

// MockFederatedUpstreamGroupSetMockRecorder is the mock recorder for MockFederatedUpstreamGroupSet
type MockFederatedUpstreamGroupSetMockRecorder struct {
	mock *MockFederatedUpstreamGroupSet
}

// NewMockFederatedUpstreamGroupSet creates a new mock instance
func NewMockFederatedUpstreamGroupSet(ctrl *gomock.Controller) *MockFederatedUpstreamGroupSet {
	mock := &MockFederatedUpstreamGroupSet{ctrl: ctrl}
	mock.recorder = &MockFederatedUpstreamGroupSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedUpstreamGroupSet) EXPECT() *MockFederatedUpstreamGroupSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockFederatedUpstreamGroupSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockFederatedUpstreamGroupSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockFederatedUpstreamGroupSet)(nil).Keys))
}

// List mocks base method
func (m *MockFederatedUpstreamGroupSet) List(filterResource ...func(*v1.FederatedUpstreamGroup) bool) []*v1.FederatedUpstreamGroup {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.FederatedUpstreamGroup)
	return ret0
}

// List indicates an expected call of List
func (mr *MockFederatedUpstreamGroupSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFederatedUpstreamGroupSet)(nil).List), filterResource...)
}

// Map mocks base method
func (m *MockFederatedUpstreamGroupSet) Map() map[string]*v1.FederatedUpstreamGroup {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.FederatedUpstreamGroup)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockFederatedUpstreamGroupSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockFederatedUpstreamGroupSet)(nil).Map))
}

// Insert mocks base method
func (m *MockFederatedUpstreamGroupSet) Insert(federatedUpstreamGroup ...*v1.FederatedUpstreamGroup) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range federatedUpstreamGroup {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockFederatedUpstreamGroupSetMockRecorder) Insert(federatedUpstreamGroup ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockFederatedUpstreamGroupSet)(nil).Insert), federatedUpstreamGroup...)
}

// Equal mocks base method
func (m *MockFederatedUpstreamGroupSet) Equal(federatedUpstreamGroupSet v1sets.FederatedUpstreamGroupSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", federatedUpstreamGroupSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockFederatedUpstreamGroupSetMockRecorder) Equal(federatedUpstreamGroupSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockFederatedUpstreamGroupSet)(nil).Equal), federatedUpstreamGroupSet)
}

// Has mocks base method
func (m *MockFederatedUpstreamGroupSet) Has(federatedUpstreamGroup ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", federatedUpstreamGroup)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockFederatedUpstreamGroupSetMockRecorder) Has(federatedUpstreamGroup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockFederatedUpstreamGroupSet)(nil).Has), federatedUpstreamGroup)
}

// Delete mocks base method
func (m *MockFederatedUpstreamGroupSet) Delete(federatedUpstreamGroup ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", federatedUpstreamGroup)
}

// Delete indicates an expected call of Delete
func (mr *MockFederatedUpstreamGroupSetMockRecorder) Delete(federatedUpstreamGroup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFederatedUpstreamGroupSet)(nil).Delete), federatedUpstreamGroup)
}

// Union mocks base method
func (m *MockFederatedUpstreamGroupSet) Union(set v1sets.FederatedUpstreamGroupSet) v1sets.FederatedUpstreamGroupSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.FederatedUpstreamGroupSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockFederatedUpstreamGroupSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockFederatedUpstreamGroupSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockFederatedUpstreamGroupSet) Difference(set v1sets.FederatedUpstreamGroupSet) v1sets.FederatedUpstreamGroupSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.FederatedUpstreamGroupSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockFederatedUpstreamGroupSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockFederatedUpstreamGroupSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockFederatedUpstreamGroupSet) Intersection(set v1sets.FederatedUpstreamGroupSet) v1sets.FederatedUpstreamGroupSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.FederatedUpstreamGroupSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockFederatedUpstreamGroupSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockFederatedUpstreamGroupSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockFederatedUpstreamGroupSet) Find(id ezkube.ResourceId) (*v1.FederatedUpstreamGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.FederatedUpstreamGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockFederatedUpstreamGroupSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockFederatedUpstreamGroupSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockFederatedUpstreamGroupSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockFederatedUpstreamGroupSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockFederatedUpstreamGroupSet)(nil).Length))
}

// Generic mocks base method
func (m *MockFederatedUpstreamGroupSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockFederatedUpstreamGroupSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockFederatedUpstreamGroupSet)(nil).Generic))
}

// Delta mocks base method
func (m *MockFederatedUpstreamGroupSet) Delta(newSet v1sets.FederatedUpstreamGroupSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta
func (mr *MockFederatedUpstreamGroupSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockFederatedUpstreamGroupSet)(nil).Delta), newSet)
}

// MockFederatedSettingsSet is a mock of FederatedSettingsSet interface
type MockFederatedSettingsSet struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedSettingsSetMockRecorder
}

// MockFederatedSettingsSetMockRecorder is the mock recorder for MockFederatedSettingsSet
type MockFederatedSettingsSetMockRecorder struct {
	mock *MockFederatedSettingsSet
}

// NewMockFederatedSettingsSet creates a new mock instance
func NewMockFederatedSettingsSet(ctrl *gomock.Controller) *MockFederatedSettingsSet {
	mock := &MockFederatedSettingsSet{ctrl: ctrl}
	mock.recorder = &MockFederatedSettingsSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatedSettingsSet) EXPECT() *MockFederatedSettingsSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockFederatedSettingsSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockFederatedSettingsSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockFederatedSettingsSet)(nil).Keys))
}

// List mocks base method
func (m *MockFederatedSettingsSet) List(filterResource ...func(*v1.FederatedSettings) bool) []*v1.FederatedSettings {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.FederatedSettings)
	return ret0
}

// List indicates an expected call of List
func (mr *MockFederatedSettingsSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFederatedSettingsSet)(nil).List), filterResource...)
}

// Map mocks base method
func (m *MockFederatedSettingsSet) Map() map[string]*v1.FederatedSettings {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.FederatedSettings)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockFederatedSettingsSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockFederatedSettingsSet)(nil).Map))
}

// Insert mocks base method
func (m *MockFederatedSettingsSet) Insert(federatedSettings ...*v1.FederatedSettings) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range federatedSettings {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockFederatedSettingsSetMockRecorder) Insert(federatedSettings ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockFederatedSettingsSet)(nil).Insert), federatedSettings...)
}

// Equal mocks base method
func (m *MockFederatedSettingsSet) Equal(federatedSettingsSet v1sets.FederatedSettingsSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", federatedSettingsSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockFederatedSettingsSetMockRecorder) Equal(federatedSettingsSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockFederatedSettingsSet)(nil).Equal), federatedSettingsSet)
}

// Has mocks base method
func (m *MockFederatedSettingsSet) Has(federatedSettings ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", federatedSettings)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockFederatedSettingsSetMockRecorder) Has(federatedSettings interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockFederatedSettingsSet)(nil).Has), federatedSettings)
}

// Delete mocks base method
func (m *MockFederatedSettingsSet) Delete(federatedSettings ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", federatedSettings)
}

// Delete indicates an expected call of Delete
func (mr *MockFederatedSettingsSetMockRecorder) Delete(federatedSettings interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFederatedSettingsSet)(nil).Delete), federatedSettings)
}

// Union mocks base method
func (m *MockFederatedSettingsSet) Union(set v1sets.FederatedSettingsSet) v1sets.FederatedSettingsSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.FederatedSettingsSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockFederatedSettingsSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockFederatedSettingsSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockFederatedSettingsSet) Difference(set v1sets.FederatedSettingsSet) v1sets.FederatedSettingsSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.FederatedSettingsSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockFederatedSettingsSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockFederatedSettingsSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockFederatedSettingsSet) Intersection(set v1sets.FederatedSettingsSet) v1sets.FederatedSettingsSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.FederatedSettingsSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockFederatedSettingsSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockFederatedSettingsSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockFederatedSettingsSet) Find(id ezkube.ResourceId) (*v1.FederatedSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.FederatedSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockFederatedSettingsSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockFederatedSettingsSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockFederatedSettingsSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockFederatedSettingsSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockFederatedSettingsSet)(nil).Length))
}

// Generic mocks base method
func (m *MockFederatedSettingsSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockFederatedSettingsSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockFederatedSettingsSet)(nil).Generic))
}

// Delta mocks base method
func (m *MockFederatedSettingsSet) Delta(newSet v1sets.FederatedSettingsSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta
func (mr *MockFederatedSettingsSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockFederatedSettingsSet)(nil).Delta), newSet)
}
