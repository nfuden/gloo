// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/apiserver/api/rpc.edge.gloo/v1/glooinstance.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *GlooInstance) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooInstance)
	if !ok {
		that2, ok := that.(GlooInstance)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSpec(), target.GetSpec()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatus(), target.GetStatus()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListGlooInstancesRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListGlooInstancesRequest)
	if !ok {
		that2, ok := that.(ListGlooInstancesRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *ListGlooInstancesResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListGlooInstancesResponse)
	if !ok {
		that2, ok := that.(ListGlooInstancesResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetGlooInstances()) != len(target.GetGlooInstances()) {
		return false
	}
	for idx, v := range m.GetGlooInstances() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetGlooInstances()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetGlooInstances()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ClusterDetails) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ClusterDetails)
	if !ok {
		that2, ok := that.(ClusterDetails)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetCluster(), target.GetCluster()) != 0 {
		return false
	}

	if len(m.GetGlooInstances()) != len(target.GetGlooInstances()) {
		return false
	}
	for idx, v := range m.GetGlooInstances() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetGlooInstances()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetGlooInstances()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ListClusterDetailsRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListClusterDetailsRequest)
	if !ok {
		that2, ok := that.(ListClusterDetailsRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *ListClusterDetailsResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListClusterDetailsResponse)
	if !ok {
		that2, ok := that.(ListClusterDetailsResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetClusterDetails()) != len(target.GetClusterDetails()) {
		return false
	}
	for idx, v := range m.GetClusterDetails() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetClusterDetails()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetClusterDetails()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ConfigDump) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ConfigDump)
	if !ok {
		that2, ok := that.(ConfigDump)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetRaw(), target.GetRaw()) != 0 {
		return false
	}

	if strings.Compare(m.GetError(), target.GetError()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GetConfigDumpsRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetConfigDumpsRequest)
	if !ok {
		that2, ok := that.(GetConfigDumpsRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetGlooInstanceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstanceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstanceRef(), target.GetGlooInstanceRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetConfigDumpsResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetConfigDumpsResponse)
	if !ok {
		that2, ok := that.(GetConfigDumpsResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetConfigDumps()) != len(target.GetConfigDumps()) {
		return false
	}
	for idx, v := range m.GetConfigDumps() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetConfigDumps()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetConfigDumps()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GetUpstreamHostsRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetUpstreamHostsRequest)
	if !ok {
		that2, ok := that.(GetUpstreamHostsRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetGlooInstanceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstanceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstanceRef(), target.GetGlooInstanceRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetUpstreamHostsResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetUpstreamHostsResponse)
	if !ok {
		that2, ok := that.(GetUpstreamHostsResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetUpstreamHosts()) != len(target.GetUpstreamHosts()) {
		return false
	}
	for k, v := range m.GetUpstreamHosts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetUpstreamHosts()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetUpstreamHosts()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *HostList) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HostList)
	if !ok {
		that2, ok := that.(HostList)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetHosts()) != len(target.GetHosts()) {
		return false
	}
	for idx, v := range m.GetHosts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHosts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHosts()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *Host) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Host)
	if !ok {
		that2, ok := that.(Host)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetAddress(), target.GetAddress()) != 0 {
		return false
	}

	if m.GetPort() != target.GetPort() {
		return false
	}

	if m.GetWeight() != target.GetWeight() {
		return false
	}

	if h, ok := interface{}(m.GetProxyRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProxyRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProxyRef(), target.GetProxyRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GlooInstance_GlooInstanceSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooInstance_GlooInstanceSpec)
	if !ok {
		that2, ok := that.(GlooInstance_GlooInstanceSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetCluster(), target.GetCluster()) != 0 {
		return false
	}

	if m.GetIsEnterprise() != target.GetIsEnterprise() {
		return false
	}

	if h, ok := interface{}(m.GetControlPlane()).(equality.Equalizer); ok {
		if !h.Equal(target.GetControlPlane()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetControlPlane(), target.GetControlPlane()) {
			return false
		}
	}

	if len(m.GetProxies()) != len(target.GetProxies()) {
		return false
	}
	for idx, v := range m.GetProxies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetProxies()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetProxies()[idx]) {
				return false
			}
		}

	}

	if strings.Compare(m.GetRegion(), target.GetRegion()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetCheck()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCheck()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCheck(), target.GetCheck()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GlooInstance_GlooInstanceStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooInstance_GlooInstanceStatus)
	if !ok {
		that2, ok := that.(GlooInstance_GlooInstanceStatus)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *GlooInstance_GlooInstanceSpec_ControlPlane) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooInstance_GlooInstanceSpec_ControlPlane)
	if !ok {
		that2, ok := that.(GlooInstance_GlooInstanceSpec_ControlPlane)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetVersion(), target.GetVersion()) != 0 {
		return false
	}

	if strings.Compare(m.GetNamespace(), target.GetNamespace()) != 0 {
		return false
	}

	if len(m.GetWatchedNamespaces()) != len(target.GetWatchedNamespaces()) {
		return false
	}
	for idx, v := range m.GetWatchedNamespaces() {

		if strings.Compare(v, target.GetWatchedNamespaces()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *GlooInstance_GlooInstanceSpec_Proxy) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooInstance_GlooInstanceSpec_Proxy)
	if !ok {
		that2, ok := that.(GlooInstance_GlooInstanceSpec_Proxy)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetReplicas() != target.GetReplicas() {
		return false
	}

	if m.GetAvailableReplicas() != target.GetAvailableReplicas() {
		return false
	}

	if m.GetReadyReplicas() != target.GetReadyReplicas() {
		return false
	}

	if m.GetWasmEnabled() != target.GetWasmEnabled() {
		return false
	}

	if m.GetReadConfigMulticlusterEnabled() != target.GetReadConfigMulticlusterEnabled() {
		return false
	}

	if strings.Compare(m.GetVersion(), target.GetVersion()) != 0 {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetNamespace(), target.GetNamespace()) != 0 {
		return false
	}

	if m.GetWorkloadControllerType() != target.GetWorkloadControllerType() {
		return false
	}

	if len(m.GetZones()) != len(target.GetZones()) {
		return false
	}
	for idx, v := range m.GetZones() {

		if strings.Compare(v, target.GetZones()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetIngressEndpoints()) != len(target.GetIngressEndpoints()) {
		return false
	}
	for idx, v := range m.GetIngressEndpoints() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetIngressEndpoints()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetIngressEndpoints()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GlooInstance_GlooInstanceSpec_Check) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooInstance_GlooInstanceSpec_Check)
	if !ok {
		that2, ok := that.(GlooInstance_GlooInstanceSpec_Check)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetGateways()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGateways()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGateways(), target.GetGateways()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetVirtualServices()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServices()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServices(), target.GetVirtualServices()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRouteTables()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRouteTables()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRouteTables(), target.GetRouteTables()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAuthConfigs()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAuthConfigs()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAuthConfigs(), target.GetAuthConfigs()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSettings()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSettings()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSettings(), target.GetSettings()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetUpstreams()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpstreams()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpstreams(), target.GetUpstreams()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetUpstreamGroups()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpstreamGroups()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpstreamGroups(), target.GetUpstreamGroups()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetProxies()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProxies()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProxies(), target.GetProxies()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRateLimitConfigs()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRateLimitConfigs()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRateLimitConfigs(), target.GetRateLimitConfigs()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMatchableHttpGateways()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMatchableHttpGateways()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMatchableHttpGateways(), target.GetMatchableHttpGateways()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMatchableTcpGateways()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMatchableTcpGateways()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMatchableTcpGateways(), target.GetMatchableTcpGateways()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDeployments()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDeployments()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDeployments(), target.GetDeployments()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPods()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPods()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPods(), target.GetPods()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GlooInstance_GlooInstanceSpec_Proxy_IngressEndpoint) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooInstance_GlooInstanceSpec_Proxy_IngressEndpoint)
	if !ok {
		that2, ok := that.(GlooInstance_GlooInstanceSpec_Proxy_IngressEndpoint)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetAddress(), target.GetAddress()) != 0 {
		return false
	}

	if len(m.GetPorts()) != len(target.GetPorts()) {
		return false
	}
	for idx, v := range m.GetPorts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPorts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPorts()[idx]) {
				return false
			}
		}

	}

	if strings.Compare(m.GetServiceName(), target.GetServiceName()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GlooInstance_GlooInstanceSpec_Proxy_IngressEndpoint_Port) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooInstance_GlooInstanceSpec_Proxy_IngressEndpoint_Port)
	if !ok {
		that2, ok := that.(GlooInstance_GlooInstanceSpec_Proxy_IngressEndpoint_Port)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetPort() != target.GetPort() {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GlooInstance_GlooInstanceSpec_Check_Summary) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooInstance_GlooInstanceSpec_Check_Summary)
	if !ok {
		that2, ok := that.(GlooInstance_GlooInstanceSpec_Check_Summary)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetTotal() != target.GetTotal() {
		return false
	}

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetErrors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetErrors()[idx]) {
				return false
			}
		}

	}

	if len(m.GetWarnings()) != len(target.GetWarnings()) {
		return false
	}
	for idx, v := range m.GetWarnings() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetWarnings()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetWarnings()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GlooInstance_GlooInstanceSpec_Check_Summary_ResourceReport) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlooInstance_GlooInstanceSpec_Check_Summary_ResourceReport)
	if !ok {
		that2, ok := that.(GlooInstance_GlooInstanceSpec_Check_Summary_ResourceReport)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRef(), target.GetRef()) {
			return false
		}
	}

	if strings.Compare(m.GetMessage(), target.GetMessage()) != 0 {
		return false
	}

	return true
}
