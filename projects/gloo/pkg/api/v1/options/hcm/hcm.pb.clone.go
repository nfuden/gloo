// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/hcm/hcm.proto

package hcm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_duration "github.com/golang/protobuf/ptypes/duration"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/headers"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/protocol"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/protocol_upgrade"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/tracing"
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
func (m *HttpConnectionManagerSettings) Clone() proto.Message {
	var target *HttpConnectionManagerSettings
	if m == nil {
		return target
	}
	target = &HttpConnectionManagerSettings{}

	if h, ok := interface{}(m.GetSkipXffAppend()).(clone.Cloner); ok {
		target.SkipXffAppend = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.SkipXffAppend = proto.Clone(m.GetSkipXffAppend()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetVia()).(clone.Cloner); ok {
		target.Via = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	} else {
		target.Via = proto.Clone(m.GetVia()).(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	}

	if h, ok := interface{}(m.GetXffNumTrustedHops()).(clone.Cloner); ok {
		target.XffNumTrustedHops = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.XffNumTrustedHops = proto.Clone(m.GetXffNumTrustedHops()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetUseRemoteAddress()).(clone.Cloner); ok {
		target.UseRemoteAddress = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.UseRemoteAddress = proto.Clone(m.GetUseRemoteAddress()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetGenerateRequestId()).(clone.Cloner); ok {
		target.GenerateRequestId = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.GenerateRequestId = proto.Clone(m.GetGenerateRequestId()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetProxy_100Continue()).(clone.Cloner); ok {
		target.Proxy_100Continue = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.Proxy_100Continue = proto.Clone(m.GetProxy_100Continue()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetStreamIdleTimeout()).(clone.Cloner); ok {
		target.StreamIdleTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.StreamIdleTimeout = proto.Clone(m.GetStreamIdleTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetIdleTimeout()).(clone.Cloner); ok {
		target.IdleTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.IdleTimeout = proto.Clone(m.GetIdleTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetMaxRequestHeadersKb()).(clone.Cloner); ok {
		target.MaxRequestHeadersKb = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxRequestHeadersKb = proto.Clone(m.GetMaxRequestHeadersKb()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetRequestTimeout()).(clone.Cloner); ok {
		target.RequestTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.RequestTimeout = proto.Clone(m.GetRequestTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetRequestHeadersTimeout()).(clone.Cloner); ok {
		target.RequestHeadersTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.RequestHeadersTimeout = proto.Clone(m.GetRequestHeadersTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetDrainTimeout()).(clone.Cloner); ok {
		target.DrainTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.DrainTimeout = proto.Clone(m.GetDrainTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetDelayedCloseTimeout()).(clone.Cloner); ok {
		target.DelayedCloseTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.DelayedCloseTimeout = proto.Clone(m.GetDelayedCloseTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetServerName()).(clone.Cloner); ok {
		target.ServerName = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	} else {
		target.ServerName = proto.Clone(m.GetServerName()).(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	}

	if h, ok := interface{}(m.GetStripAnyHostPort()).(clone.Cloner); ok {
		target.StripAnyHostPort = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.StripAnyHostPort = proto.Clone(m.GetStripAnyHostPort()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetAcceptHttp_10()).(clone.Cloner); ok {
		target.AcceptHttp_10 = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.AcceptHttp_10 = proto.Clone(m.GetAcceptHttp_10()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetDefaultHostForHttp_10()).(clone.Cloner); ok {
		target.DefaultHostForHttp_10 = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	} else {
		target.DefaultHostForHttp_10 = proto.Clone(m.GetDefaultHostForHttp_10()).(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	}

	if h, ok := interface{}(m.GetAllowChunkedLength()).(clone.Cloner); ok {
		target.AllowChunkedLength = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.AllowChunkedLength = proto.Clone(m.GetAllowChunkedLength()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetEnableTrailers()).(clone.Cloner); ok {
		target.EnableTrailers = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.EnableTrailers = proto.Clone(m.GetEnableTrailers()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetTracing()).(clone.Cloner); ok {
		target.Tracing = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing.ListenerTracingSettings)
	} else {
		target.Tracing = proto.Clone(m.GetTracing()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing.ListenerTracingSettings)
	}

	target.ForwardClientCertDetails = m.GetForwardClientCertDetails()

	if h, ok := interface{}(m.GetSetCurrentClientCertDetails()).(clone.Cloner); ok {
		target.SetCurrentClientCertDetails = h.Clone().(*HttpConnectionManagerSettings_SetCurrentClientCertDetails)
	} else {
		target.SetCurrentClientCertDetails = proto.Clone(m.GetSetCurrentClientCertDetails()).(*HttpConnectionManagerSettings_SetCurrentClientCertDetails)
	}

	if h, ok := interface{}(m.GetPreserveExternalRequestId()).(clone.Cloner); ok {
		target.PreserveExternalRequestId = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.PreserveExternalRequestId = proto.Clone(m.GetPreserveExternalRequestId()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if m.GetUpgrades() != nil {
		target.Upgrades = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig, len(m.GetUpgrades()))
		for idx, v := range m.GetUpgrades() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Upgrades[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig)
			} else {
				target.Upgrades[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig)
			}

		}
	}

	if h, ok := interface{}(m.GetMaxConnectionDuration()).(clone.Cloner); ok {
		target.MaxConnectionDuration = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.MaxConnectionDuration = proto.Clone(m.GetMaxConnectionDuration()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetMaxStreamDuration()).(clone.Cloner); ok {
		target.MaxStreamDuration = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.MaxStreamDuration = proto.Clone(m.GetMaxStreamDuration()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetMaxHeadersCount()).(clone.Cloner); ok {
		target.MaxHeadersCount = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxHeadersCount = proto.Clone(m.GetMaxHeadersCount()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	target.HeadersWithUnderscoresAction = m.GetHeadersWithUnderscoresAction()

	if h, ok := interface{}(m.GetMaxRequestsPerConnection()).(clone.Cloner); ok {
		target.MaxRequestsPerConnection = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxRequestsPerConnection = proto.Clone(m.GetMaxRequestsPerConnection()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	target.ServerHeaderTransformation = m.GetServerHeaderTransformation()

	target.PathWithEscapedSlashesAction = m.GetPathWithEscapedSlashesAction()

	target.CodecType = m.GetCodecType()

	if h, ok := interface{}(m.GetMergeSlashes()).(clone.Cloner); ok {
		target.MergeSlashes = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.MergeSlashes = proto.Clone(m.GetMergeSlashes()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetNormalizePath()).(clone.Cloner); ok {
		target.NormalizePath = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.NormalizePath = proto.Clone(m.GetNormalizePath()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetUuidRequestIdConfig()).(clone.Cloner); ok {
		target.UuidRequestIdConfig = h.Clone().(*HttpConnectionManagerSettings_UuidRequestIdConfigSettings)
	} else {
		target.UuidRequestIdConfig = proto.Clone(m.GetUuidRequestIdConfig()).(*HttpConnectionManagerSettings_UuidRequestIdConfigSettings)
	}

	if h, ok := interface{}(m.GetHttp2ProtocolOptions()).(clone.Cloner); ok {
		target.Http2ProtocolOptions = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol.Http2ProtocolOptions)
	} else {
		target.Http2ProtocolOptions = proto.Clone(m.GetHttp2ProtocolOptions()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol.Http2ProtocolOptions)
	}

	if h, ok := interface{}(m.GetInternalAddressConfig()).(clone.Cloner); ok {
		target.InternalAddressConfig = h.Clone().(*HttpConnectionManagerSettings_InternalAddressConfig)
	} else {
		target.InternalAddressConfig = proto.Clone(m.GetInternalAddressConfig()).(*HttpConnectionManagerSettings_InternalAddressConfig)
	}

	if h, ok := interface{}(m.GetAppendXForwardedPort()).(clone.Cloner); ok {
		target.AppendXForwardedPort = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.AppendXForwardedPort = proto.Clone(m.GetAppendXForwardedPort()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetEarlyHeaderManipulation()).(clone.Cloner); ok {
		target.EarlyHeaderManipulation = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers.EarlyHeaderManipulation)
	} else {
		target.EarlyHeaderManipulation = proto.Clone(m.GetEarlyHeaderManipulation()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers.EarlyHeaderManipulation)
	}

	switch m.HeaderFormat.(type) {

	case *HttpConnectionManagerSettings_ProperCaseHeaderKeyFormat:

		if h, ok := interface{}(m.GetProperCaseHeaderKeyFormat()).(clone.Cloner); ok {
			target.HeaderFormat = &HttpConnectionManagerSettings_ProperCaseHeaderKeyFormat{
				ProperCaseHeaderKeyFormat: h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue),
			}
		} else {
			target.HeaderFormat = &HttpConnectionManagerSettings_ProperCaseHeaderKeyFormat{
				ProperCaseHeaderKeyFormat: proto.Clone(m.GetProperCaseHeaderKeyFormat()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue),
			}
		}

	case *HttpConnectionManagerSettings_PreserveCaseHeaderKeyFormat:

		if h, ok := interface{}(m.GetPreserveCaseHeaderKeyFormat()).(clone.Cloner); ok {
			target.HeaderFormat = &HttpConnectionManagerSettings_PreserveCaseHeaderKeyFormat{
				PreserveCaseHeaderKeyFormat: h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue),
			}
		} else {
			target.HeaderFormat = &HttpConnectionManagerSettings_PreserveCaseHeaderKeyFormat{
				PreserveCaseHeaderKeyFormat: proto.Clone(m.GetPreserveCaseHeaderKeyFormat()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue),
			}
		}

	}

	return target
}

// Clone function
func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) Clone() proto.Message {
	var target *HttpConnectionManagerSettings_SetCurrentClientCertDetails
	if m == nil {
		return target
	}
	target = &HttpConnectionManagerSettings_SetCurrentClientCertDetails{}

	if h, ok := interface{}(m.GetSubject()).(clone.Cloner); ok {
		target.Subject = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.Subject = proto.Clone(m.GetSubject()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetCert()).(clone.Cloner); ok {
		target.Cert = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.Cert = proto.Clone(m.GetCert()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetChain()).(clone.Cloner); ok {
		target.Chain = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.Chain = proto.Clone(m.GetChain()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetDns()).(clone.Cloner); ok {
		target.Dns = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.Dns = proto.Clone(m.GetDns()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetUri()).(clone.Cloner); ok {
		target.Uri = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.Uri = proto.Clone(m.GetUri()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	return target
}

// Clone function
func (m *HttpConnectionManagerSettings_UuidRequestIdConfigSettings) Clone() proto.Message {
	var target *HttpConnectionManagerSettings_UuidRequestIdConfigSettings
	if m == nil {
		return target
	}
	target = &HttpConnectionManagerSettings_UuidRequestIdConfigSettings{}

	if h, ok := interface{}(m.GetPackTraceReason()).(clone.Cloner); ok {
		target.PackTraceReason = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.PackTraceReason = proto.Clone(m.GetPackTraceReason()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetUseRequestIdForTraceSampling()).(clone.Cloner); ok {
		target.UseRequestIdForTraceSampling = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.UseRequestIdForTraceSampling = proto.Clone(m.GetUseRequestIdForTraceSampling()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	return target
}

// Clone function
func (m *HttpConnectionManagerSettings_CidrRange) Clone() proto.Message {
	var target *HttpConnectionManagerSettings_CidrRange
	if m == nil {
		return target
	}
	target = &HttpConnectionManagerSettings_CidrRange{}

	target.AddressPrefix = m.GetAddressPrefix()

	if h, ok := interface{}(m.GetPrefixLen()).(clone.Cloner); ok {
		target.PrefixLen = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.PrefixLen = proto.Clone(m.GetPrefixLen()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	return target
}

// Clone function
func (m *HttpConnectionManagerSettings_InternalAddressConfig) Clone() proto.Message {
	var target *HttpConnectionManagerSettings_InternalAddressConfig
	if m == nil {
		return target
	}
	target = &HttpConnectionManagerSettings_InternalAddressConfig{}

	if h, ok := interface{}(m.GetUnixSockets()).(clone.Cloner); ok {
		target.UnixSockets = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.UnixSockets = proto.Clone(m.GetUnixSockets()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if m.GetCidrRanges() != nil {
		target.CidrRanges = make([]*HttpConnectionManagerSettings_CidrRange, len(m.GetCidrRanges()))
		for idx, v := range m.GetCidrRanges() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.CidrRanges[idx] = h.Clone().(*HttpConnectionManagerSettings_CidrRange)
			} else {
				target.CidrRanges[idx] = proto.Clone(v).(*HttpConnectionManagerSettings_CidrRange)
			}

		}
	}

	return target
}
