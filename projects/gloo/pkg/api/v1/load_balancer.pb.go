// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: github.com/solo-io/gloo/projects/gloo/api/v1/load_balancer.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// LoadBalancerConfig is the settings for the load balancer used to send requests to the Upstream endpoints.
type LoadBalancerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configures envoy's panic threshold Percent between 0-100. Once the number of non health hosts
	// reaches this percentage, envoy disregards health information.
	// see more info [here](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/panic_threshold.html).
	HealthyPanicThreshold *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=healthy_panic_threshold,json=healthyPanicThreshold,proto3" json:"healthy_panic_threshold,omitempty"`
	// This allows batch updates of endpoints health/weight/metadata that happen during a time window.
	// this help lower cpu usage when endpoint change rate is high. defaults to 1 second.
	// Set to 0 to disable and have changes applied immediately.
	UpdateMergeWindow *duration.Duration `protobuf:"bytes,2,opt,name=update_merge_window,json=updateMergeWindow,proto3" json:"update_merge_window,omitempty"`
	// Types that are assignable to Type:
	//
	//	*LoadBalancerConfig_RoundRobin_
	//	*LoadBalancerConfig_LeastRequest_
	//	*LoadBalancerConfig_Random_
	//	*LoadBalancerConfig_RingHash_
	//	*LoadBalancerConfig_Maglev_
	Type isLoadBalancerConfig_Type `protobuf_oneof:"type"`
	// Types that are assignable to LocalityConfig:
	//
	//	*LoadBalancerConfig_LocalityWeightedLbConfig
	LocalityConfig isLoadBalancerConfig_LocalityConfig `protobuf_oneof:"locality_config"`
}

func (x *LoadBalancerConfig) Reset() {
	*x = LoadBalancerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig) ProtoMessage() {}

func (x *LoadBalancerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0}
}

func (x *LoadBalancerConfig) GetHealthyPanicThreshold() *wrappers.DoubleValue {
	if x != nil {
		return x.HealthyPanicThreshold
	}
	return nil
}

func (x *LoadBalancerConfig) GetUpdateMergeWindow() *duration.Duration {
	if x != nil {
		return x.UpdateMergeWindow
	}
	return nil
}

func (m *LoadBalancerConfig) GetType() isLoadBalancerConfig_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *LoadBalancerConfig) GetRoundRobin() *LoadBalancerConfig_RoundRobin {
	if x, ok := x.GetType().(*LoadBalancerConfig_RoundRobin_); ok {
		return x.RoundRobin
	}
	return nil
}

func (x *LoadBalancerConfig) GetLeastRequest() *LoadBalancerConfig_LeastRequest {
	if x, ok := x.GetType().(*LoadBalancerConfig_LeastRequest_); ok {
		return x.LeastRequest
	}
	return nil
}

func (x *LoadBalancerConfig) GetRandom() *LoadBalancerConfig_Random {
	if x, ok := x.GetType().(*LoadBalancerConfig_Random_); ok {
		return x.Random
	}
	return nil
}

func (x *LoadBalancerConfig) GetRingHash() *LoadBalancerConfig_RingHash {
	if x, ok := x.GetType().(*LoadBalancerConfig_RingHash_); ok {
		return x.RingHash
	}
	return nil
}

func (x *LoadBalancerConfig) GetMaglev() *LoadBalancerConfig_Maglev {
	if x, ok := x.GetType().(*LoadBalancerConfig_Maglev_); ok {
		return x.Maglev
	}
	return nil
}

func (m *LoadBalancerConfig) GetLocalityConfig() isLoadBalancerConfig_LocalityConfig {
	if m != nil {
		return m.LocalityConfig
	}
	return nil
}

func (x *LoadBalancerConfig) GetLocalityWeightedLbConfig() *empty.Empty {
	if x, ok := x.GetLocalityConfig().(*LoadBalancerConfig_LocalityWeightedLbConfig); ok {
		return x.LocalityWeightedLbConfig
	}
	return nil
}

type isLoadBalancerConfig_Type interface {
	isLoadBalancerConfig_Type()
}

type LoadBalancerConfig_RoundRobin_ struct {
	// Use round robin for load balancing. Round robin is the default load balancing method.
	RoundRobin *LoadBalancerConfig_RoundRobin `protobuf:"bytes,3,opt,name=round_robin,json=roundRobin,proto3,oneof"`
}

type LoadBalancerConfig_LeastRequest_ struct {
	// Use least request for load balancing.
	LeastRequest *LoadBalancerConfig_LeastRequest `protobuf:"bytes,4,opt,name=least_request,json=leastRequest,proto3,oneof"`
}

type LoadBalancerConfig_Random_ struct {
	// Use random for load balancing.
	Random *LoadBalancerConfig_Random `protobuf:"bytes,5,opt,name=random,proto3,oneof"`
}

type LoadBalancerConfig_RingHash_ struct {
	// Use ring hash for load balancing.
	RingHash *LoadBalancerConfig_RingHash `protobuf:"bytes,6,opt,name=ring_hash,json=ringHash,proto3,oneof"`
}

type LoadBalancerConfig_Maglev_ struct {
	// Use maglev for load balancing.
	Maglev *LoadBalancerConfig_Maglev `protobuf:"bytes,7,opt,name=maglev,proto3,oneof"`
}

func (*LoadBalancerConfig_RoundRobin_) isLoadBalancerConfig_Type() {}

func (*LoadBalancerConfig_LeastRequest_) isLoadBalancerConfig_Type() {}

func (*LoadBalancerConfig_Random_) isLoadBalancerConfig_Type() {}

func (*LoadBalancerConfig_RingHash_) isLoadBalancerConfig_Type() {}

func (*LoadBalancerConfig_Maglev_) isLoadBalancerConfig_Type() {}

type isLoadBalancerConfig_LocalityConfig interface {
	isLoadBalancerConfig_LocalityConfig()
}

type LoadBalancerConfig_LocalityWeightedLbConfig struct {
	// (Enterprise Only)
	// https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/locality_weight#locality-weighted-load-balancing
	// Locality weighted load balancing enables weighting assignments across different zones and geographical locations by using explicit weights.
	// This field is required to enable locality weighted load balancing
	LocalityWeightedLbConfig *empty.Empty `protobuf:"bytes,8,opt,name=locality_weighted_lb_config,json=localityWeightedLbConfig,proto3,oneof"`
}

func (*LoadBalancerConfig_LocalityWeightedLbConfig) isLoadBalancerConfig_LocalityConfig() {}

type LoadBalancerConfig_RoundRobin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for slow start mode. If this configuration is not set, slow start will not be not enabled.
	SlowStartConfig *LoadBalancerConfig_SlowStartConfig `protobuf:"bytes,1,opt,name=slow_start_config,json=slowStartConfig,proto3" json:"slow_start_config,omitempty"`
}

func (x *LoadBalancerConfig_RoundRobin) Reset() {
	*x = LoadBalancerConfig_RoundRobin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig_RoundRobin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig_RoundRobin) ProtoMessage() {}

func (x *LoadBalancerConfig_RoundRobin) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig_RoundRobin.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig_RoundRobin) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0, 0}
}

func (x *LoadBalancerConfig_RoundRobin) GetSlowStartConfig() *LoadBalancerConfig_SlowStartConfig {
	if x != nil {
		return x.SlowStartConfig
	}
	return nil
}

type LoadBalancerConfig_LeastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How many choices to take into account. defaults to 2.
	ChoiceCount uint32 `protobuf:"varint,1,opt,name=choice_count,json=choiceCount,proto3" json:"choice_count,omitempty"`
	// Configuration for slow start mode. If this configuration is not set, slow start will not be not enabled.
	SlowStartConfig *LoadBalancerConfig_SlowStartConfig `protobuf:"bytes,2,opt,name=slow_start_config,json=slowStartConfig,proto3" json:"slow_start_config,omitempty"`
}

func (x *LoadBalancerConfig_LeastRequest) Reset() {
	*x = LoadBalancerConfig_LeastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig_LeastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig_LeastRequest) ProtoMessage() {}

func (x *LoadBalancerConfig_LeastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig_LeastRequest.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig_LeastRequest) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0, 1}
}

func (x *LoadBalancerConfig_LeastRequest) GetChoiceCount() uint32 {
	if x != nil {
		return x.ChoiceCount
	}
	return 0
}

func (x *LoadBalancerConfig_LeastRequest) GetSlowStartConfig() *LoadBalancerConfig_SlowStartConfig {
	if x != nil {
		return x.SlowStartConfig
	}
	return nil
}

type LoadBalancerConfig_Random struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoadBalancerConfig_Random) Reset() {
	*x = LoadBalancerConfig_Random{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig_Random) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig_Random) ProtoMessage() {}

func (x *LoadBalancerConfig_Random) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig_Random.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig_Random) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0, 2}
}

// Customizes the parameters used in the hashing algorithm to refine performance or resource usage.
type LoadBalancerConfig_RingHashConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Minimum hash ring size. The larger the ring is (that is, the more hashes there are for each provided host)
	// the better the request distribution will reflect the desired weights. Defaults to 1024 entries, and limited
	// to 8M entries.
	MinimumRingSize uint64 `protobuf:"varint,1,opt,name=minimum_ring_size,json=minimumRingSize,proto3" json:"minimum_ring_size,omitempty"`
	// Maximum hash ring size. Defaults to 8M entries, and limited to 8M entries, but can be lowered to further
	// constrain resource use.
	MaximumRingSize uint64 `protobuf:"varint,2,opt,name=maximum_ring_size,json=maximumRingSize,proto3" json:"maximum_ring_size,omitempty"`
}

func (x *LoadBalancerConfig_RingHashConfig) Reset() {
	*x = LoadBalancerConfig_RingHashConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig_RingHashConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig_RingHashConfig) ProtoMessage() {}

func (x *LoadBalancerConfig_RingHashConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig_RingHashConfig.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig_RingHashConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0, 3}
}

func (x *LoadBalancerConfig_RingHashConfig) GetMinimumRingSize() uint64 {
	if x != nil {
		return x.MinimumRingSize
	}
	return 0
}

func (x *LoadBalancerConfig_RingHashConfig) GetMaximumRingSize() uint64 {
	if x != nil {
		return x.MaximumRingSize
	}
	return 0
}

type LoadBalancerConfig_RingHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional, customizes the parameters used in the hashing algorithm
	RingHashConfig *LoadBalancerConfig_RingHashConfig `protobuf:"bytes,1,opt,name=ring_hash_config,json=ringHashConfig,proto3" json:"ring_hash_config,omitempty"`
}

func (x *LoadBalancerConfig_RingHash) Reset() {
	*x = LoadBalancerConfig_RingHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig_RingHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig_RingHash) ProtoMessage() {}

func (x *LoadBalancerConfig_RingHash) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig_RingHash.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig_RingHash) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0, 4}
}

func (x *LoadBalancerConfig_RingHash) GetRingHashConfig() *LoadBalancerConfig_RingHashConfig {
	if x != nil {
		return x.RingHashConfig
	}
	return nil
}

type LoadBalancerConfig_Maglev struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoadBalancerConfig_Maglev) Reset() {
	*x = LoadBalancerConfig_Maglev{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig_Maglev) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig_Maglev) ProtoMessage() {}

func (x *LoadBalancerConfig_Maglev) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig_Maglev.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig_Maglev) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0, 5}
}

type LoadBalancerConfig_SlowStartConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Represents the size of slow start window.
	// If set, the newly created host remains in slow start mode starting from its creation time
	// for the duration of slow start window.
	SlowStartWindow *duration.Duration `protobuf:"bytes,1,opt,name=slow_start_window,json=slowStartWindow,proto3" json:"slow_start_window,omitempty"`
	// This parameter controls the speed of traffic increase over the slow start window. Defaults to 1.0,
	// so that endpoint would get linearly increasing amount of traffic.
	// When increasing the value for this parameter, the speed of traffic ramp-up increases non-linearly.
	// The value of aggression parameter should be greater than 0.0.
	// By tuning the parameter, is possible to achieve polynomial or exponential shape of ramp-up curve.
	//
	// During slow start window, effective weight of an endpoint would be scaled with time factor and aggression:
	// “new_weight = weight * max(min_weight_percent, time_factor ^ (1 / aggression))“,
	// where “time_factor=(time_since_start_seconds / slow_start_time_seconds)“.
	//
	// As time progresses, more and more traffic would be sent to endpoint, which is in slow start window.
	// Once host exits slow start, time_factor and aggression no longer affect its weight.
	Aggression *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=aggression,proto3" json:"aggression,omitempty"`
	// Configures the minimum percentage of origin weight that avoids too small new weight,
	// which may cause endpoints in slow start mode receive no traffic in slow start window.
	// If not specified, the default is 10%.
	MinWeightPercent *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=min_weight_percent,json=minWeightPercent,proto3" json:"min_weight_percent,omitempty"`
}

func (x *LoadBalancerConfig_SlowStartConfig) Reset() {
	*x = LoadBalancerConfig_SlowStartConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig_SlowStartConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig_SlowStartConfig) ProtoMessage() {}

func (x *LoadBalancerConfig_SlowStartConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig_SlowStartConfig.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig_SlowStartConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0, 6}
}

func (x *LoadBalancerConfig_SlowStartConfig) GetSlowStartWindow() *duration.Duration {
	if x != nil {
		return x.SlowStartWindow
	}
	return nil
}

func (x *LoadBalancerConfig_SlowStartConfig) GetAggression() *wrappers.DoubleValue {
	if x != nil {
		return x.Aggression
	}
	return nil
}

func (x *LoadBalancerConfig_SlowStartConfig) GetMinWeightPercent() *wrappers.DoubleValue {
	if x != nil {
		return x.MinWeightPercent
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65,
	0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe7, 0x0a, 0x0a, 0x12, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x54, 0x0a, 0x17, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x79, 0x5f, 0x70, 0x61, 0x6e, 0x69, 0x63, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79,
	0x50, 0x61, 0x6e, 0x69, 0x63, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x49,
	0x0a, 0x13, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x5f, 0x77,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x72, 0x67, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x4e, 0x0a, 0x0b, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f,
	0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x6f, 0x62, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x52, 0x6f, 0x62, 0x69, 0x6e, 0x12, 0x54, 0x0a, 0x0d, 0x6c, 0x65, 0x61,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x0c, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x41, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c,
	0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x48, 0x00, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x12, 0x48, 0x0a, 0x09, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68,
	0x48, 0x00, 0x52, 0x08, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x12, 0x41, 0x0a, 0x06,
	0x6d, 0x61, 0x67, 0x6c, 0x65, 0x76, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x61, 0x64,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d,
	0x61, 0x67, 0x6c, 0x65, 0x76, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x61, 0x67, 0x6c, 0x65, 0x76, 0x12,
	0x57, 0x0a, 0x1b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x01, 0x52, 0x18,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64,
	0x4c, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x6a, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x6e,
	0x64, 0x52, 0x6f, 0x62, 0x69, 0x6e, 0x12, 0x5c, 0x0a, 0x11, 0x73, 0x6c, 0x6f, 0x77, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x53, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0f, 0x73, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x8f, 0x01, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x5c, 0x0a, 0x11, 0x73, 0x6c, 0x6f, 0x77,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0f, 0x73, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x08, 0x0a, 0x06, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x1a, 0x68, 0x0a, 0x0e, 0x52, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6d,
	0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x65, 0x0a, 0x08, 0x52, 0x69,
	0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x12, 0x59, 0x0a, 0x10, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x52, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0e, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0x08, 0x0a, 0x06, 0x4d, 0x61, 0x67, 0x6c, 0x65, 0x76, 0x1a, 0xe2, 0x01, 0x0a, 0x0f,
	0x53, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x45, 0x0a, 0x11, 0x73, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x77, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x73, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x3c, 0x0a, 0x0a, 0x61, 0x67, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x61, 0x67, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x12, 0x6d, 0x69, 0x6e, 0x5f, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10,
	0x6d, 0x69, 0x6e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x3e, 0xb8, 0xf5, 0x04,
	0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_goTypes = []interface{}{
	(*LoadBalancerConfig)(nil),                 // 0: gloo.solo.io.LoadBalancerConfig
	(*LoadBalancerConfig_RoundRobin)(nil),      // 1: gloo.solo.io.LoadBalancerConfig.RoundRobin
	(*LoadBalancerConfig_LeastRequest)(nil),    // 2: gloo.solo.io.LoadBalancerConfig.LeastRequest
	(*LoadBalancerConfig_Random)(nil),          // 3: gloo.solo.io.LoadBalancerConfig.Random
	(*LoadBalancerConfig_RingHashConfig)(nil),  // 4: gloo.solo.io.LoadBalancerConfig.RingHashConfig
	(*LoadBalancerConfig_RingHash)(nil),        // 5: gloo.solo.io.LoadBalancerConfig.RingHash
	(*LoadBalancerConfig_Maglev)(nil),          // 6: gloo.solo.io.LoadBalancerConfig.Maglev
	(*LoadBalancerConfig_SlowStartConfig)(nil), // 7: gloo.solo.io.LoadBalancerConfig.SlowStartConfig
	(*wrappers.DoubleValue)(nil),               // 8: google.protobuf.DoubleValue
	(*duration.Duration)(nil),                  // 9: google.protobuf.Duration
	(*empty.Empty)(nil),                        // 10: google.protobuf.Empty
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_depIdxs = []int32{
	8,  // 0: gloo.solo.io.LoadBalancerConfig.healthy_panic_threshold:type_name -> google.protobuf.DoubleValue
	9,  // 1: gloo.solo.io.LoadBalancerConfig.update_merge_window:type_name -> google.protobuf.Duration
	1,  // 2: gloo.solo.io.LoadBalancerConfig.round_robin:type_name -> gloo.solo.io.LoadBalancerConfig.RoundRobin
	2,  // 3: gloo.solo.io.LoadBalancerConfig.least_request:type_name -> gloo.solo.io.LoadBalancerConfig.LeastRequest
	3,  // 4: gloo.solo.io.LoadBalancerConfig.random:type_name -> gloo.solo.io.LoadBalancerConfig.Random
	5,  // 5: gloo.solo.io.LoadBalancerConfig.ring_hash:type_name -> gloo.solo.io.LoadBalancerConfig.RingHash
	6,  // 6: gloo.solo.io.LoadBalancerConfig.maglev:type_name -> gloo.solo.io.LoadBalancerConfig.Maglev
	10, // 7: gloo.solo.io.LoadBalancerConfig.locality_weighted_lb_config:type_name -> google.protobuf.Empty
	7,  // 8: gloo.solo.io.LoadBalancerConfig.RoundRobin.slow_start_config:type_name -> gloo.solo.io.LoadBalancerConfig.SlowStartConfig
	7,  // 9: gloo.solo.io.LoadBalancerConfig.LeastRequest.slow_start_config:type_name -> gloo.solo.io.LoadBalancerConfig.SlowStartConfig
	4,  // 10: gloo.solo.io.LoadBalancerConfig.RingHash.ring_hash_config:type_name -> gloo.solo.io.LoadBalancerConfig.RingHashConfig
	9,  // 11: gloo.solo.io.LoadBalancerConfig.SlowStartConfig.slow_start_window:type_name -> google.protobuf.Duration
	8,  // 12: gloo.solo.io.LoadBalancerConfig.SlowStartConfig.aggression:type_name -> google.protobuf.DoubleValue
	8,  // 13: gloo.solo.io.LoadBalancerConfig.SlowStartConfig.min_weight_percent:type_name -> google.protobuf.DoubleValue
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig_RoundRobin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig_LeastRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig_Random); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig_RingHashConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig_RingHash); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig_Maglev); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig_SlowStartConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LoadBalancerConfig_RoundRobin_)(nil),
		(*LoadBalancerConfig_LeastRequest_)(nil),
		(*LoadBalancerConfig_Random_)(nil),
		(*LoadBalancerConfig_RingHash_)(nil),
		(*LoadBalancerConfig_Maglev_)(nil),
		(*LoadBalancerConfig_LocalityWeightedLbConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_depIdxs = nil
}
