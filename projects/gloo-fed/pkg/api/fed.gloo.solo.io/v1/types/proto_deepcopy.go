// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package types

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/solo-io/protoc-gen-ext/pkg/clone"
)

// DeepCopyInto for the FederatedUpstream.Spec
func (in *FederatedUpstreamSpec) DeepCopyInto(out *FederatedUpstreamSpec) {
	var p *FederatedUpstreamSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*FederatedUpstreamSpec)
	} else {
		p = proto.Clone(in).(*FederatedUpstreamSpec)
	}
	*out = *p
}

// DeepCopyInto for the FederatedUpstream.Status
func (in *FederatedUpstreamStatus) DeepCopyInto(out *FederatedUpstreamStatus) {
	var p *FederatedUpstreamStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*FederatedUpstreamStatus)
	} else {
		p = proto.Clone(in).(*FederatedUpstreamStatus)
	}
	*out = *p
}

// DeepCopyInto for the FederatedUpstreamGroup.Spec
func (in *FederatedUpstreamGroupSpec) DeepCopyInto(out *FederatedUpstreamGroupSpec) {
	var p *FederatedUpstreamGroupSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*FederatedUpstreamGroupSpec)
	} else {
		p = proto.Clone(in).(*FederatedUpstreamGroupSpec)
	}
	*out = *p
}

// DeepCopyInto for the FederatedUpstreamGroup.Status
func (in *FederatedUpstreamGroupStatus) DeepCopyInto(out *FederatedUpstreamGroupStatus) {
	var p *FederatedUpstreamGroupStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*FederatedUpstreamGroupStatus)
	} else {
		p = proto.Clone(in).(*FederatedUpstreamGroupStatus)
	}
	*out = *p
}

// DeepCopyInto for the FederatedSettings.Spec
func (in *FederatedSettingsSpec) DeepCopyInto(out *FederatedSettingsSpec) {
	var p *FederatedSettingsSpec
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*FederatedSettingsSpec)
	} else {
		p = proto.Clone(in).(*FederatedSettingsSpec)
	}
	*out = *p
}

// DeepCopyInto for the FederatedSettings.Status
func (in *FederatedSettingsStatus) DeepCopyInto(out *FederatedSettingsStatus) {
	var p *FederatedSettingsStatus
	if h, ok := interface{}(in).(clone.Cloner); ok {
		p = h.Clone().(*FederatedSettingsStatus)
	} else {
		p = proto.Clone(in).(*FederatedSettingsStatus)
	}
	*out = *p
}