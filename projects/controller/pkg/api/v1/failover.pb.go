// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/controller/api/v1/failover.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	ssl "github.com/solo-io/gloo/projects/controller/pkg/api/v1/ssl"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Failover configuration for an upstream.
//
// Failover allows for optional fallback endpoints in the case that the primary set of endpoints is deemed
// unhealthy. As failover requires knowledge of the health of each set of endpoints, active or passive
// health checks must be configured on an upstream using failover in order for it to work properly.
//
// Failover closely resembles the Envoy config which this is translated to, with one notable exception.
// The priorities are not defined on the `LocalityLbEndpoints` but rather inferred from the list of
// `PrioritizedLocality`. More information on envoy prioritization can be found
// [here](https://www.envoyproxy.io/docs/envoy/v1.14.1/intro/arch_overview/upstream/load_balancing/priority#arch-overview-load-balancing-priority-levels).
// In practice this means that the priority of a given set of `LocalityLbEndpoints` is determined by its index in
// the list, first being `0` through `n-1`.
type Failover struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PrioritizedLocality is an implicitly prioritized list of lists of `LocalityLbEndpoints`. The priority of each
	// list of `LocalityLbEndpoints` is determined by its index in the list.
	PrioritizedLocalities []*Failover_PrioritizedLocality `protobuf:"bytes,1,rep,name=prioritized_localities,json=prioritizedLocalities,proto3" json:"prioritized_localities,omitempty"`
	// Load balancing policy settings.
	Policy *Failover_Policy `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *Failover) Reset() {
	*x = Failover{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Failover) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Failover) ProtoMessage() {}

func (x *Failover) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Failover.ProtoReflect.Descriptor instead.
func (*Failover) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDescGZIP(), []int{0}
}

func (x *Failover) GetPrioritizedLocalities() []*Failover_PrioritizedLocality {
	if x != nil {
		return x.PrioritizedLocalities
	}
	return nil
}

func (x *Failover) GetPolicy() *Failover_Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

// A group of endpoints belonging to a Locality.
// One can have multiple LocalityLbEndpoints for a locality, but this is
// generally only done if the different groups need to have different load
// balancing weights or different priorities.
type LocalityLbEndpoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifies where the parent upstream hosts run.
	Locality *Locality `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	// The group of endpoints belonging to the locality specified.
	// Note: If any address is DNS resolvable than `lb_endpoints[].load_balancing_weight` is not allowed on any of
	// this locality's endpoints.
	LbEndpoints []*LbEndpoint `protobuf:"bytes,2,rep,name=lb_endpoints,json=lbEndpoints,proto3" json:"lb_endpoints,omitempty"`
	// Optional: Per priority/region/zone/sub_zone weight; at least 1. The load
	// balancing weight for a locality is divided by the sum of the weights of all
	// localities at the same priority level to produce the effective percentage
	// of traffic for the locality.
	// To enable locality weighted load balancing, load_balancer_config.locality_weighted_lb_config must be set as well
	LoadBalancingWeight *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
}

func (x *LocalityLbEndpoints) Reset() {
	*x = LocalityLbEndpoints{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LocalityLbEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalityLbEndpoints) ProtoMessage() {}

func (x *LocalityLbEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalityLbEndpoints.ProtoReflect.Descriptor instead.
func (*LocalityLbEndpoints) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDescGZIP(), []int{1}
}

func (x *LocalityLbEndpoints) GetLocality() *Locality {
	if x != nil {
		return x.Locality
	}
	return nil
}

func (x *LocalityLbEndpoints) GetLbEndpoints() []*LbEndpoint {
	if x != nil {
		return x.LbEndpoints
	}
	return nil
}

func (x *LocalityLbEndpoints) GetLoadBalancingWeight() *wrapperspb.UInt32Value {
	if x != nil {
		return x.LoadBalancingWeight
	}
	return nil
}

// An Endpoint that Envoy can route traffic to.
type LbEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address (hostname or IP)
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Port the instance is listening on
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// The optional health check configuration is used as configuration for the
	// health checker to contact the health checked host.
	// This takes into effect only for upstreams with active health checking enabled
	HealthCheckConfig *LbEndpoint_HealthCheckConfig `protobuf:"bytes,3,opt,name=health_check_config,json=healthCheckConfig,proto3" json:"health_check_config,omitempty"`
	UpstreamSslConfig *ssl.UpstreamSslConfig        `protobuf:"bytes,4,opt,name=upstream_ssl_config,json=upstreamSslConfig,proto3" json:"upstream_ssl_config,omitempty"`
	// The optional load balancing weight of the upstream host; at least 1.
	// Envoy uses the load balancing weight in some of the built in load
	// balancers. The load balancing weight for an endpoint is divided by the sum
	// of the weights of all endpoints in the endpoint's locality to produce a
	// percentage of traffic for the endpoint. This percentage is then further
	// weighted by the endpoint's locality's load balancing weight from
	// LocalityLbEndpoints. If unspecified, each host is presumed to have equal
	// weight in a locality.
	LoadBalancingWeight *wrapperspb.UInt32Value `protobuf:"bytes,5,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	// Additional metadata to add to the endpoint. This metadata can be used in upstream HTTP filters
	// or other specific Envoy configurations.
	// The following keys are added by k8sgateway and are ignored if set:
	// - "envoy.transport_socket_match"
	// - "io.solo.health_checkers.advanced_http"
	Metadata map[string]*structpb.Struct `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LbEndpoint) Reset() {
	*x = LbEndpoint{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LbEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LbEndpoint) ProtoMessage() {}

func (x *LbEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LbEndpoint.ProtoReflect.Descriptor instead.
func (*LbEndpoint) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDescGZIP(), []int{2}
}

func (x *LbEndpoint) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *LbEndpoint) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *LbEndpoint) GetHealthCheckConfig() *LbEndpoint_HealthCheckConfig {
	if x != nil {
		return x.HealthCheckConfig
	}
	return nil
}

func (x *LbEndpoint) GetUpstreamSslConfig() *ssl.UpstreamSslConfig {
	if x != nil {
		return x.UpstreamSslConfig
	}
	return nil
}

func (x *LbEndpoint) GetLoadBalancingWeight() *wrapperspb.UInt32Value {
	if x != nil {
		return x.LoadBalancingWeight
	}
	return nil
}

func (x *LbEndpoint) GetMetadata() map[string]*structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Identifies location of where either Envoy runs or where upstream hosts run.
type Locality struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Region this zone belongs to.
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// Defines the local service zone where Envoy is running. The meaning of zone
	// is context dependent, e.g. [Availability Zone (AZ)](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html)
	// on AWS, [Zone](https://cloud.google.com/compute/docs/regions-zones/) on
	// GCP, etc.
	Zone string `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	// When used for locality of upstream hosts, this field further splits zone
	// into smaller chunks of sub-zones so they can be load balanced
	// independently.
	SubZone string `protobuf:"bytes,3,opt,name=sub_zone,json=subZone,proto3" json:"sub_zone,omitempty"`
}

func (x *Locality) Reset() {
	*x = Locality{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Locality) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Locality) ProtoMessage() {}

func (x *Locality) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Locality.ProtoReflect.Descriptor instead.
func (*Locality) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDescGZIP(), []int{3}
}

func (x *Locality) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Locality) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Locality) GetSubZone() string {
	if x != nil {
		return x.SubZone
	}
	return ""
}

type Failover_PrioritizedLocality struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalityEndpoints []*LocalityLbEndpoints `protobuf:"bytes,2,rep,name=locality_endpoints,json=localityEndpoints,proto3" json:"locality_endpoints,omitempty"`
}

func (x *Failover_PrioritizedLocality) Reset() {
	*x = Failover_PrioritizedLocality{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Failover_PrioritizedLocality) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Failover_PrioritizedLocality) ProtoMessage() {}

func (x *Failover_PrioritizedLocality) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Failover_PrioritizedLocality.ProtoReflect.Descriptor instead.
func (*Failover_PrioritizedLocality) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Failover_PrioritizedLocality) GetLocalityEndpoints() []*LocalityLbEndpoints {
	if x != nil {
		return x.LocalityEndpoints
	}
	return nil
}

type Failover_Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Priority levels and localities are considered overprovisioned with this
	// factor (in percentage). This means that we don't consider a priority
	// level or locality unhealthy until the fraction of healthy hosts
	// multiplied by the overprovisioning factor drops below 100.
	// With the default value 140(1.4), Envoy doesn't consider a priority level
	// or a locality unhealthy until their percentage of healthy hosts drops
	// below 72%. For example:
	//
	// .. code-block:: json
	//
	//	{ "overprovisioning_factor": 100 }
	//
	// Read more at priority levels and
	// localities.
	OverprovisioningFactor *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=overprovisioning_factor,json=overprovisioningFactor,proto3" json:"overprovisioning_factor,omitempty"`
}

func (x *Failover_Policy) Reset() {
	*x = Failover_Policy{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Failover_Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Failover_Policy) ProtoMessage() {}

func (x *Failover_Policy) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Failover_Policy.ProtoReflect.Descriptor instead.
func (*Failover_Policy) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Failover_Policy) GetOverprovisioningFactor() *wrapperspb.UInt32Value {
	if x != nil {
		return x.OverprovisioningFactor
	}
	return nil
}

// The optional health check configuration.
type LbEndpoint_HealthCheckConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional alternative health check port value.
	//
	// By default the health check address port of an upstream host is the same
	// as the host's serving address port. This provides an alternative health
	// check port. Setting this with a non-zero value allows an upstream host
	// to have different health check address port.
	PortValue uint32 `protobuf:"varint,1,opt,name=port_value,json=portValue,proto3" json:"port_value,omitempty"`
	// By default, the host header for L7 health checks is controlled by cluster level configuration. Setting this
	// to a non-empty value allows overriding the cluster level configuration for a specific endpoint.
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Path to use when health checking this failover endpoint.
	// Default is empty path.
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// Method to use when health checking this failover endpoint. Defaults to `GET`.
	Method string `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *LbEndpoint_HealthCheckConfig) Reset() {
	*x = LbEndpoint_HealthCheckConfig{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LbEndpoint_HealthCheckConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LbEndpoint_HealthCheckConfig) ProtoMessage() {}

func (x *LbEndpoint_HealthCheckConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LbEndpoint_HealthCheckConfig.ProtoReflect.Descriptor instead.
func (*LbEndpoint_HealthCheckConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDescGZIP(), []int{2, 0}
}

func (x *LbEndpoint_HealthCheckConfig) GetPortValue() uint32 {
	if x != nil {
		return x.PortValue
	}
	return 0
}

func (x *LbEndpoint_HealthCheckConfig) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *LbEndpoint_HealthCheckConfig) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *LbEndpoint_HealthCheckConfig) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x73, 0x6c, 0x2f,
	0x73, 0x73, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf7, 0x02, 0x0a, 0x08, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x12,
	0x61, 0x0a, 0x16, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46,
	0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69,
	0x7a, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x15, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x67, 0x0a, 0x13, 0x50, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x50, 0x0a, 0x12, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x4c, 0x62, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52,
	0x11, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x1a, 0x68, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x5e, 0x0a, 0x17,
	0x6f, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x20, 0x00, 0x52, 0x16, 0x6f, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0xd8, 0x01, 0x0a,
	0x13, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x62, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x0c, 0x6c, 0x62, 0x5f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x62,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x6c, 0x62, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x50, 0x0a, 0x15, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x13, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e,
	0x67, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xd8, 0x04, 0x0a, 0x0a, 0x4c, 0x62, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x5a, 0x0a, 0x13, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x4c, 0x62, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x4f, 0x0a, 0x13, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x73, 0x73, 0x6c,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x55, 0x70, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x73, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11,
	0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x73, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x59, 0x0a, 0x15, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x69, 0x6e, 0x67, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x01, 0x52, 0x13, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x42, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x62,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x7a, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x54, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x51, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75,
	0x62, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x5a, 0x6f, 0x6e, 0x65, 0x42, 0x3e, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0,
	0xf5, 0x04, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_goTypes = []any{
	(*Failover)(nil),                     // 0: gloo.solo.io.Failover
	(*LocalityLbEndpoints)(nil),          // 1: gloo.solo.io.LocalityLbEndpoints
	(*LbEndpoint)(nil),                   // 2: gloo.solo.io.LbEndpoint
	(*Locality)(nil),                     // 3: gloo.solo.io.Locality
	(*Failover_PrioritizedLocality)(nil), // 4: gloo.solo.io.Failover.PrioritizedLocality
	(*Failover_Policy)(nil),              // 5: gloo.solo.io.Failover.Policy
	(*LbEndpoint_HealthCheckConfig)(nil), // 6: gloo.solo.io.LbEndpoint.HealthCheckConfig
	nil,                                  // 7: gloo.solo.io.LbEndpoint.MetadataEntry
	(*wrapperspb.UInt32Value)(nil),       // 8: google.protobuf.UInt32Value
	(*ssl.UpstreamSslConfig)(nil),        // 9: gloo.solo.io.UpstreamSslConfig
	(*structpb.Struct)(nil),              // 10: google.protobuf.Struct
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_depIdxs = []int32{
	4,  // 0: gloo.solo.io.Failover.prioritized_localities:type_name -> gloo.solo.io.Failover.PrioritizedLocality
	5,  // 1: gloo.solo.io.Failover.policy:type_name -> gloo.solo.io.Failover.Policy
	3,  // 2: gloo.solo.io.LocalityLbEndpoints.locality:type_name -> gloo.solo.io.Locality
	2,  // 3: gloo.solo.io.LocalityLbEndpoints.lb_endpoints:type_name -> gloo.solo.io.LbEndpoint
	8,  // 4: gloo.solo.io.LocalityLbEndpoints.load_balancing_weight:type_name -> google.protobuf.UInt32Value
	6,  // 5: gloo.solo.io.LbEndpoint.health_check_config:type_name -> gloo.solo.io.LbEndpoint.HealthCheckConfig
	9,  // 6: gloo.solo.io.LbEndpoint.upstream_ssl_config:type_name -> gloo.solo.io.UpstreamSslConfig
	8,  // 7: gloo.solo.io.LbEndpoint.load_balancing_weight:type_name -> google.protobuf.UInt32Value
	7,  // 8: gloo.solo.io.LbEndpoint.metadata:type_name -> gloo.solo.io.LbEndpoint.MetadataEntry
	1,  // 9: gloo.solo.io.Failover.PrioritizedLocality.locality_endpoints:type_name -> gloo.solo.io.LocalityLbEndpoints
	8,  // 10: gloo.solo.io.Failover.Policy.overprovisioning_factor:type_name -> google.protobuf.UInt32Value
	10, // 11: gloo.solo.io.LbEndpoint.MetadataEntry.value:type_name -> google.protobuf.Struct
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_failover_proto_depIdxs = nil
}