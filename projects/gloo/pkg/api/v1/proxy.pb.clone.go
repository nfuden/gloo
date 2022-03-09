// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/proxy.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_empty "github.com/golang/protobuf/ptypes/empty"

	github_com_golang_protobuf_ptypes_struct "github.com/golang/protobuf/ptypes/struct"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_core_matchers "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_dynamic_forward_proxy "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/dynamic_forward_proxy"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *Proxy) Clone() proto.Message {
	var target *Proxy
	if m == nil {
		return target
	}
	target = &Proxy{}

	target.CompressedSpec = m.GetCompressedSpec()

	if m.GetListeners() != nil {
		target.Listeners = make([]*Listener, len(m.GetListeners()))
		for idx, v := range m.GetListeners() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Listeners[idx] = h.Clone().(*Listener)
			} else {
				target.Listeners[idx] = proto.Clone(v).(*Listener)
			}

		}
	}

	if h, ok := interface{}(m.GetNamespacedStatuses()).(clone.Cloner); ok {
		target.NamespacedStatuses = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	} else {
		target.NamespacedStatuses = proto.Clone(m.GetNamespacedStatuses()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	return target
}

// Clone function
func (m *Listener) Clone() proto.Message {
	var target *Listener
	if m == nil {
		return target
	}
	target = &Listener{}

	target.Name = m.GetName()

	target.BindAddress = m.GetBindAddress()

	target.BindPort = m.GetBindPort()

	if m.GetSslConfigurations() != nil {
		target.SslConfigurations = make([]*SslConfig, len(m.GetSslConfigurations()))
		for idx, v := range m.GetSslConfigurations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SslConfigurations[idx] = h.Clone().(*SslConfig)
			} else {
				target.SslConfigurations[idx] = proto.Clone(v).(*SslConfig)
			}

		}
	}

	if h, ok := interface{}(m.GetUseProxyProto()).(clone.Cloner); ok {
		target.UseProxyProto = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.UseProxyProto = proto.Clone(m.GetUseProxyProto()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*ListenerOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*ListenerOptions)
	}

	if h, ok := interface{}(m.GetRouteOptions()).(clone.Cloner); ok {
		target.RouteOptions = h.Clone().(*RouteConfigurationOptions)
	} else {
		target.RouteOptions = proto.Clone(m.GetRouteOptions()).(*RouteConfigurationOptions)
	}

	switch m.ListenerType.(type) {

	case *Listener_HttpListener:

		if h, ok := interface{}(m.GetHttpListener()).(clone.Cloner); ok {
			target.ListenerType = &Listener_HttpListener{
				HttpListener: h.Clone().(*HttpListener),
			}
		} else {
			target.ListenerType = &Listener_HttpListener{
				HttpListener: proto.Clone(m.GetHttpListener()).(*HttpListener),
			}
		}

	case *Listener_TcpListener:

		if h, ok := interface{}(m.GetTcpListener()).(clone.Cloner); ok {
			target.ListenerType = &Listener_TcpListener{
				TcpListener: h.Clone().(*TcpListener),
			}
		} else {
			target.ListenerType = &Listener_TcpListener{
				TcpListener: proto.Clone(m.GetTcpListener()).(*TcpListener),
			}
		}

	case *Listener_HybridListener:

		if h, ok := interface{}(m.GetHybridListener()).(clone.Cloner); ok {
			target.ListenerType = &Listener_HybridListener{
				HybridListener: h.Clone().(*HybridListener),
			}
		} else {
			target.ListenerType = &Listener_HybridListener{
				HybridListener: proto.Clone(m.GetHybridListener()).(*HybridListener),
			}
		}

	}

	switch m.OpaqueMetadata.(type) {

	case *Listener_Metadata:

		if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
			target.OpaqueMetadata = &Listener_Metadata{
				Metadata: h.Clone().(*github_com_golang_protobuf_ptypes_struct.Struct),
			}
		} else {
			target.OpaqueMetadata = &Listener_Metadata{
				Metadata: proto.Clone(m.GetMetadata()).(*github_com_golang_protobuf_ptypes_struct.Struct),
			}
		}

	case *Listener_MetadataStatic:

		if h, ok := interface{}(m.GetMetadataStatic()).(clone.Cloner); ok {
			target.OpaqueMetadata = &Listener_MetadataStatic{
				MetadataStatic: h.Clone().(*SourceMetadata),
			}
		} else {
			target.OpaqueMetadata = &Listener_MetadataStatic{
				MetadataStatic: proto.Clone(m.GetMetadataStatic()).(*SourceMetadata),
			}
		}

	}

	return target
}

// Clone function
func (m *TcpListener) Clone() proto.Message {
	var target *TcpListener
	if m == nil {
		return target
	}
	target = &TcpListener{}

	if m.GetTcpHosts() != nil {
		target.TcpHosts = make([]*TcpHost, len(m.GetTcpHosts()))
		for idx, v := range m.GetTcpHosts() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.TcpHosts[idx] = h.Clone().(*TcpHost)
			} else {
				target.TcpHosts[idx] = proto.Clone(v).(*TcpHost)
			}

		}
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*TcpListenerOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*TcpListenerOptions)
	}

	target.StatPrefix = m.GetStatPrefix()

	return target
}

// Clone function
func (m *TcpHost) Clone() proto.Message {
	var target *TcpHost
	if m == nil {
		return target
	}
	target = &TcpHost{}

	target.Name = m.GetName()

	if h, ok := interface{}(m.GetSslConfig()).(clone.Cloner); ok {
		target.SslConfig = h.Clone().(*SslConfig)
	} else {
		target.SslConfig = proto.Clone(m.GetSslConfig()).(*SslConfig)
	}

	if h, ok := interface{}(m.GetDestination()).(clone.Cloner); ok {
		target.Destination = h.Clone().(*TcpHost_TcpAction)
	} else {
		target.Destination = proto.Clone(m.GetDestination()).(*TcpHost_TcpAction)
	}

	return target
}

// Clone function
func (m *HttpListener) Clone() proto.Message {
	var target *HttpListener
	if m == nil {
		return target
	}
	target = &HttpListener{}

	if m.GetVirtualHosts() != nil {
		target.VirtualHosts = make([]*VirtualHost, len(m.GetVirtualHosts()))
		for idx, v := range m.GetVirtualHosts() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.VirtualHosts[idx] = h.Clone().(*VirtualHost)
			} else {
				target.VirtualHosts[idx] = proto.Clone(v).(*VirtualHost)
			}

		}
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*HttpListenerOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*HttpListenerOptions)
	}

	target.StatPrefix = m.GetStatPrefix()

	return target
}

// Clone function
func (m *HybridListener) Clone() proto.Message {
	var target *HybridListener
	if m == nil {
		return target
	}
	target = &HybridListener{}

	if m.GetMatchedListeners() != nil {
		target.MatchedListeners = make([]*MatchedListener, len(m.GetMatchedListeners()))
		for idx, v := range m.GetMatchedListeners() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.MatchedListeners[idx] = h.Clone().(*MatchedListener)
			} else {
				target.MatchedListeners[idx] = proto.Clone(v).(*MatchedListener)
			}

		}
	}

	return target
}

// Clone function
func (m *MatchedListener) Clone() proto.Message {
	var target *MatchedListener
	if m == nil {
		return target
	}
	target = &MatchedListener{}

	if h, ok := interface{}(m.GetMatcher()).(clone.Cloner); ok {
		target.Matcher = h.Clone().(*Matcher)
	} else {
		target.Matcher = proto.Clone(m.GetMatcher()).(*Matcher)
	}

	if m.GetSslConfigurations() != nil {
		target.SslConfigurations = make([]*SslConfig, len(m.GetSslConfigurations()))
		for idx, v := range m.GetSslConfigurations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SslConfigurations[idx] = h.Clone().(*SslConfig)
			} else {
				target.SslConfigurations[idx] = proto.Clone(v).(*SslConfig)
			}

		}
	}

	switch m.ListenerType.(type) {

	case *MatchedListener_HttpListener:

		if h, ok := interface{}(m.GetHttpListener()).(clone.Cloner); ok {
			target.ListenerType = &MatchedListener_HttpListener{
				HttpListener: h.Clone().(*HttpListener),
			}
		} else {
			target.ListenerType = &MatchedListener_HttpListener{
				HttpListener: proto.Clone(m.GetHttpListener()).(*HttpListener),
			}
		}

	case *MatchedListener_TcpListener:

		if h, ok := interface{}(m.GetTcpListener()).(clone.Cloner); ok {
			target.ListenerType = &MatchedListener_TcpListener{
				TcpListener: h.Clone().(*TcpListener),
			}
		} else {
			target.ListenerType = &MatchedListener_TcpListener{
				TcpListener: proto.Clone(m.GetTcpListener()).(*TcpListener),
			}
		}

	}

	return target
}

// Clone function
func (m *Matcher) Clone() proto.Message {
	var target *Matcher
	if m == nil {
		return target
	}
	target = &Matcher{}

	if h, ok := interface{}(m.GetSslConfig()).(clone.Cloner); ok {
		target.SslConfig = h.Clone().(*SslConfig)
	} else {
		target.SslConfig = proto.Clone(m.GetSslConfig()).(*SslConfig)
	}

	if m.GetSourcePrefixRanges() != nil {
		target.SourcePrefixRanges = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange, len(m.GetSourcePrefixRanges()))
		for idx, v := range m.GetSourcePrefixRanges() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SourcePrefixRanges[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange)
			} else {
				target.SourcePrefixRanges[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange)
			}

		}
	}

	return target
}

// Clone function
func (m *VirtualHost) Clone() proto.Message {
	var target *VirtualHost
	if m == nil {
		return target
	}
	target = &VirtualHost{}

	target.Name = m.GetName()

	if m.GetDomains() != nil {
		target.Domains = make([]string, len(m.GetDomains()))
		for idx, v := range m.GetDomains() {

			target.Domains[idx] = v

		}
	}

	if m.GetRoutes() != nil {
		target.Routes = make([]*Route, len(m.GetRoutes()))
		for idx, v := range m.GetRoutes() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Routes[idx] = h.Clone().(*Route)
			} else {
				target.Routes[idx] = proto.Clone(v).(*Route)
			}

		}
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*VirtualHostOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*VirtualHostOptions)
	}

	switch m.OpaqueMetadata.(type) {

	case *VirtualHost_Metadata:

		if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
			target.OpaqueMetadata = &VirtualHost_Metadata{
				Metadata: h.Clone().(*github_com_golang_protobuf_ptypes_struct.Struct),
			}
		} else {
			target.OpaqueMetadata = &VirtualHost_Metadata{
				Metadata: proto.Clone(m.GetMetadata()).(*github_com_golang_protobuf_ptypes_struct.Struct),
			}
		}

	case *VirtualHost_MetadataStatic:

		if h, ok := interface{}(m.GetMetadataStatic()).(clone.Cloner); ok {
			target.OpaqueMetadata = &VirtualHost_MetadataStatic{
				MetadataStatic: h.Clone().(*SourceMetadata),
			}
		} else {
			target.OpaqueMetadata = &VirtualHost_MetadataStatic{
				MetadataStatic: proto.Clone(m.GetMetadataStatic()).(*SourceMetadata),
			}
		}

	}

	return target
}

// Clone function
func (m *Route) Clone() proto.Message {
	var target *Route
	if m == nil {
		return target
	}
	target = &Route{}

	if m.GetMatchers() != nil {
		target.Matchers = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_core_matchers.Matcher, len(m.GetMatchers()))
		for idx, v := range m.GetMatchers() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Matchers[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_core_matchers.Matcher)
			} else {
				target.Matchers[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_core_matchers.Matcher)
			}

		}
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*RouteOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*RouteOptions)
	}

	target.Name = m.GetName()

	switch m.Action.(type) {

	case *Route_RouteAction:

		if h, ok := interface{}(m.GetRouteAction()).(clone.Cloner); ok {
			target.Action = &Route_RouteAction{
				RouteAction: h.Clone().(*RouteAction),
			}
		} else {
			target.Action = &Route_RouteAction{
				RouteAction: proto.Clone(m.GetRouteAction()).(*RouteAction),
			}
		}

	case *Route_RedirectAction:

		if h, ok := interface{}(m.GetRedirectAction()).(clone.Cloner); ok {
			target.Action = &Route_RedirectAction{
				RedirectAction: h.Clone().(*RedirectAction),
			}
		} else {
			target.Action = &Route_RedirectAction{
				RedirectAction: proto.Clone(m.GetRedirectAction()).(*RedirectAction),
			}
		}

	case *Route_DirectResponseAction:

		if h, ok := interface{}(m.GetDirectResponseAction()).(clone.Cloner); ok {
			target.Action = &Route_DirectResponseAction{
				DirectResponseAction: h.Clone().(*DirectResponseAction),
			}
		} else {
			target.Action = &Route_DirectResponseAction{
				DirectResponseAction: proto.Clone(m.GetDirectResponseAction()).(*DirectResponseAction),
			}
		}

	case *Route_GraphqlSchemaRef:

		if h, ok := interface{}(m.GetGraphqlSchemaRef()).(clone.Cloner); ok {
			target.Action = &Route_GraphqlSchemaRef{
				GraphqlSchemaRef: h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		} else {
			target.Action = &Route_GraphqlSchemaRef{
				GraphqlSchemaRef: proto.Clone(m.GetGraphqlSchemaRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		}

	}

	switch m.OpaqueMetadata.(type) {

	case *Route_Metadata:

		if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
			target.OpaqueMetadata = &Route_Metadata{
				Metadata: h.Clone().(*github_com_golang_protobuf_ptypes_struct.Struct),
			}
		} else {
			target.OpaqueMetadata = &Route_Metadata{
				Metadata: proto.Clone(m.GetMetadata()).(*github_com_golang_protobuf_ptypes_struct.Struct),
			}
		}

	case *Route_MetadataStatic:

		if h, ok := interface{}(m.GetMetadataStatic()).(clone.Cloner); ok {
			target.OpaqueMetadata = &Route_MetadataStatic{
				MetadataStatic: h.Clone().(*SourceMetadata),
			}
		} else {
			target.OpaqueMetadata = &Route_MetadataStatic{
				MetadataStatic: proto.Clone(m.GetMetadataStatic()).(*SourceMetadata),
			}
		}

	}

	return target
}

// Clone function
func (m *RouteAction) Clone() proto.Message {
	var target *RouteAction
	if m == nil {
		return target
	}
	target = &RouteAction{}

	switch m.Destination.(type) {

	case *RouteAction_Single:

		if h, ok := interface{}(m.GetSingle()).(clone.Cloner); ok {
			target.Destination = &RouteAction_Single{
				Single: h.Clone().(*Destination),
			}
		} else {
			target.Destination = &RouteAction_Single{
				Single: proto.Clone(m.GetSingle()).(*Destination),
			}
		}

	case *RouteAction_Multi:

		if h, ok := interface{}(m.GetMulti()).(clone.Cloner); ok {
			target.Destination = &RouteAction_Multi{
				Multi: h.Clone().(*MultiDestination),
			}
		} else {
			target.Destination = &RouteAction_Multi{
				Multi: proto.Clone(m.GetMulti()).(*MultiDestination),
			}
		}

	case *RouteAction_UpstreamGroup:

		if h, ok := interface{}(m.GetUpstreamGroup()).(clone.Cloner); ok {
			target.Destination = &RouteAction_UpstreamGroup{
				UpstreamGroup: h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		} else {
			target.Destination = &RouteAction_UpstreamGroup{
				UpstreamGroup: proto.Clone(m.GetUpstreamGroup()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		}

	case *RouteAction_ClusterHeader:

		target.Destination = &RouteAction_ClusterHeader{
			ClusterHeader: m.GetClusterHeader(),
		}

	case *RouteAction_DynamicForwardProxy:

		if h, ok := interface{}(m.GetDynamicForwardProxy()).(clone.Cloner); ok {
			target.Destination = &RouteAction_DynamicForwardProxy{
				DynamicForwardProxy: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_dynamic_forward_proxy.PerRouteConfig),
			}
		} else {
			target.Destination = &RouteAction_DynamicForwardProxy{
				DynamicForwardProxy: proto.Clone(m.GetDynamicForwardProxy()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_dynamic_forward_proxy.PerRouteConfig),
			}
		}

	}

	return target
}

// Clone function
func (m *Destination) Clone() proto.Message {
	var target *Destination
	if m == nil {
		return target
	}
	target = &Destination{}

	if h, ok := interface{}(m.GetDestinationSpec()).(clone.Cloner); ok {
		target.DestinationSpec = h.Clone().(*DestinationSpec)
	} else {
		target.DestinationSpec = proto.Clone(m.GetDestinationSpec()).(*DestinationSpec)
	}

	if h, ok := interface{}(m.GetSubset()).(clone.Cloner); ok {
		target.Subset = h.Clone().(*Subset)
	} else {
		target.Subset = proto.Clone(m.GetSubset()).(*Subset)
	}

	switch m.DestinationType.(type) {

	case *Destination_Upstream:

		if h, ok := interface{}(m.GetUpstream()).(clone.Cloner); ok {
			target.DestinationType = &Destination_Upstream{
				Upstream: h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		} else {
			target.DestinationType = &Destination_Upstream{
				Upstream: proto.Clone(m.GetUpstream()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		}

	case *Destination_Kube:

		if h, ok := interface{}(m.GetKube()).(clone.Cloner); ok {
			target.DestinationType = &Destination_Kube{
				Kube: h.Clone().(*KubernetesServiceDestination),
			}
		} else {
			target.DestinationType = &Destination_Kube{
				Kube: proto.Clone(m.GetKube()).(*KubernetesServiceDestination),
			}
		}

	case *Destination_Consul:

		if h, ok := interface{}(m.GetConsul()).(clone.Cloner); ok {
			target.DestinationType = &Destination_Consul{
				Consul: h.Clone().(*ConsulServiceDestination),
			}
		} else {
			target.DestinationType = &Destination_Consul{
				Consul: proto.Clone(m.GetConsul()).(*ConsulServiceDestination),
			}
		}

	}

	return target
}

// Clone function
func (m *KubernetesServiceDestination) Clone() proto.Message {
	var target *KubernetesServiceDestination
	if m == nil {
		return target
	}
	target = &KubernetesServiceDestination{}

	if h, ok := interface{}(m.GetRef()).(clone.Cloner); ok {
		target.Ref = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.Ref = proto.Clone(m.GetRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	target.Port = m.GetPort()

	return target
}

// Clone function
func (m *ConsulServiceDestination) Clone() proto.Message {
	var target *ConsulServiceDestination
	if m == nil {
		return target
	}
	target = &ConsulServiceDestination{}

	target.ServiceName = m.GetServiceName()

	if m.GetTags() != nil {
		target.Tags = make([]string, len(m.GetTags()))
		for idx, v := range m.GetTags() {

			target.Tags[idx] = v

		}
	}

	if m.GetDataCenters() != nil {
		target.DataCenters = make([]string, len(m.GetDataCenters()))
		for idx, v := range m.GetDataCenters() {

			target.DataCenters[idx] = v

		}
	}

	return target
}

// Clone function
func (m *UpstreamGroup) Clone() proto.Message {
	var target *UpstreamGroup
	if m == nil {
		return target
	}
	target = &UpstreamGroup{}

	if m.GetDestinations() != nil {
		target.Destinations = make([]*WeightedDestination, len(m.GetDestinations()))
		for idx, v := range m.GetDestinations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Destinations[idx] = h.Clone().(*WeightedDestination)
			} else {
				target.Destinations[idx] = proto.Clone(v).(*WeightedDestination)
			}

		}
	}

	if h, ok := interface{}(m.GetNamespacedStatuses()).(clone.Cloner); ok {
		target.NamespacedStatuses = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	} else {
		target.NamespacedStatuses = proto.Clone(m.GetNamespacedStatuses()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	return target
}

// Clone function
func (m *MultiDestination) Clone() proto.Message {
	var target *MultiDestination
	if m == nil {
		return target
	}
	target = &MultiDestination{}

	if m.GetDestinations() != nil {
		target.Destinations = make([]*WeightedDestination, len(m.GetDestinations()))
		for idx, v := range m.GetDestinations() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Destinations[idx] = h.Clone().(*WeightedDestination)
			} else {
				target.Destinations[idx] = proto.Clone(v).(*WeightedDestination)
			}

		}
	}

	return target
}

// Clone function
func (m *WeightedDestination) Clone() proto.Message {
	var target *WeightedDestination
	if m == nil {
		return target
	}
	target = &WeightedDestination{}

	if h, ok := interface{}(m.GetDestination()).(clone.Cloner); ok {
		target.Destination = h.Clone().(*Destination)
	} else {
		target.Destination = proto.Clone(m.GetDestination()).(*Destination)
	}

	target.Weight = m.GetWeight()

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*WeightedDestinationOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*WeightedDestinationOptions)
	}

	return target
}

// Clone function
func (m *RedirectAction) Clone() proto.Message {
	var target *RedirectAction
	if m == nil {
		return target
	}
	target = &RedirectAction{}

	target.HostRedirect = m.GetHostRedirect()

	target.ResponseCode = m.GetResponseCode()

	target.HttpsRedirect = m.GetHttpsRedirect()

	target.StripQuery = m.GetStripQuery()

	switch m.PathRewriteSpecifier.(type) {

	case *RedirectAction_PathRedirect:

		target.PathRewriteSpecifier = &RedirectAction_PathRedirect{
			PathRedirect: m.GetPathRedirect(),
		}

	case *RedirectAction_PrefixRewrite:

		target.PathRewriteSpecifier = &RedirectAction_PrefixRewrite{
			PrefixRewrite: m.GetPrefixRewrite(),
		}

	}

	return target
}

// Clone function
func (m *DirectResponseAction) Clone() proto.Message {
	var target *DirectResponseAction
	if m == nil {
		return target
	}
	target = &DirectResponseAction{}

	target.Status = m.GetStatus()

	target.Body = m.GetBody()

	return target
}

// Clone function
func (m *SourceMetadata) Clone() proto.Message {
	var target *SourceMetadata
	if m == nil {
		return target
	}
	target = &SourceMetadata{}

	if m.GetSources() != nil {
		target.Sources = make([]*SourceMetadata_SourceRef, len(m.GetSources()))
		for idx, v := range m.GetSources() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Sources[idx] = h.Clone().(*SourceMetadata_SourceRef)
			} else {
				target.Sources[idx] = proto.Clone(v).(*SourceMetadata_SourceRef)
			}

		}
	}

	return target
}

// Clone function
func (m *TcpHost_TcpAction) Clone() proto.Message {
	var target *TcpHost_TcpAction
	if m == nil {
		return target
	}
	target = &TcpHost_TcpAction{}

	switch m.Destination.(type) {

	case *TcpHost_TcpAction_Single:

		if h, ok := interface{}(m.GetSingle()).(clone.Cloner); ok {
			target.Destination = &TcpHost_TcpAction_Single{
				Single: h.Clone().(*Destination),
			}
		} else {
			target.Destination = &TcpHost_TcpAction_Single{
				Single: proto.Clone(m.GetSingle()).(*Destination),
			}
		}

	case *TcpHost_TcpAction_Multi:

		if h, ok := interface{}(m.GetMulti()).(clone.Cloner); ok {
			target.Destination = &TcpHost_TcpAction_Multi{
				Multi: h.Clone().(*MultiDestination),
			}
		} else {
			target.Destination = &TcpHost_TcpAction_Multi{
				Multi: proto.Clone(m.GetMulti()).(*MultiDestination),
			}
		}

	case *TcpHost_TcpAction_UpstreamGroup:

		if h, ok := interface{}(m.GetUpstreamGroup()).(clone.Cloner); ok {
			target.Destination = &TcpHost_TcpAction_UpstreamGroup{
				UpstreamGroup: h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		} else {
			target.Destination = &TcpHost_TcpAction_UpstreamGroup{
				UpstreamGroup: proto.Clone(m.GetUpstreamGroup()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		}

	case *TcpHost_TcpAction_ForwardSniClusterName:

		if h, ok := interface{}(m.GetForwardSniClusterName()).(clone.Cloner); ok {
			target.Destination = &TcpHost_TcpAction_ForwardSniClusterName{
				ForwardSniClusterName: h.Clone().(*github_com_golang_protobuf_ptypes_empty.Empty),
			}
		} else {
			target.Destination = &TcpHost_TcpAction_ForwardSniClusterName{
				ForwardSniClusterName: proto.Clone(m.GetForwardSniClusterName()).(*github_com_golang_protobuf_ptypes_empty.Empty),
			}
		}

	}

	return target
}

// Clone function
func (m *SourceMetadata_SourceRef) Clone() proto.Message {
	var target *SourceMetadata_SourceRef
	if m == nil {
		return target
	}
	target = &SourceMetadata_SourceRef{}

	if h, ok := interface{}(m.GetResourceRef()).(clone.Cloner); ok {
		target.ResourceRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.ResourceRef = proto.Clone(m.GetResourceRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	target.ResourceKind = m.GetResourceKind()

	target.ObservedGeneration = m.GetObservedGeneration()

	return target
}
