// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/service/ext_proc/v3/external_processor.proto

package v3

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
func (m *ProcessingRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ProcessingRequest)
	if !ok {
		that2, ok := that.(ProcessingRequest)
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

	if h, ok := interface{}(m.GetMetadataContext()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadataContext()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadataContext(), target.GetMetadataContext()) {
			return false
		}
	}

	if len(m.GetAttributes()) != len(target.GetAttributes()) {
		return false
	}
	for k, v := range m.GetAttributes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAttributes()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAttributes()[k]) {
				return false
			}
		}

	}

	if m.GetObservabilityMode() != target.GetObservabilityMode() {
		return false
	}

	switch m.Request.(type) {

	case *ProcessingRequest_RequestHeaders:
		if _, ok := target.Request.(*ProcessingRequest_RequestHeaders); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRequestHeaders()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRequestHeaders()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRequestHeaders(), target.GetRequestHeaders()) {
				return false
			}
		}

	case *ProcessingRequest_ResponseHeaders:
		if _, ok := target.Request.(*ProcessingRequest_ResponseHeaders); !ok {
			return false
		}

		if h, ok := interface{}(m.GetResponseHeaders()).(equality.Equalizer); ok {
			if !h.Equal(target.GetResponseHeaders()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetResponseHeaders(), target.GetResponseHeaders()) {
				return false
			}
		}

	case *ProcessingRequest_RequestBody:
		if _, ok := target.Request.(*ProcessingRequest_RequestBody); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRequestBody()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRequestBody()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRequestBody(), target.GetRequestBody()) {
				return false
			}
		}

	case *ProcessingRequest_ResponseBody:
		if _, ok := target.Request.(*ProcessingRequest_ResponseBody); !ok {
			return false
		}

		if h, ok := interface{}(m.GetResponseBody()).(equality.Equalizer); ok {
			if !h.Equal(target.GetResponseBody()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetResponseBody(), target.GetResponseBody()) {
				return false
			}
		}

	case *ProcessingRequest_RequestTrailers:
		if _, ok := target.Request.(*ProcessingRequest_RequestTrailers); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRequestTrailers()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRequestTrailers()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRequestTrailers(), target.GetRequestTrailers()) {
				return false
			}
		}

	case *ProcessingRequest_ResponseTrailers:
		if _, ok := target.Request.(*ProcessingRequest_ResponseTrailers); !ok {
			return false
		}

		if h, ok := interface{}(m.GetResponseTrailers()).(equality.Equalizer); ok {
			if !h.Equal(target.GetResponseTrailers()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetResponseTrailers(), target.GetResponseTrailers()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Request != target.Request {
			return false
		}
	}

	return true
}

// Equal function
func (m *ProcessingResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ProcessingResponse)
	if !ok {
		that2, ok := that.(ProcessingResponse)
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

	if h, ok := interface{}(m.GetDynamicMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDynamicMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDynamicMetadata(), target.GetDynamicMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetModeOverride()).(equality.Equalizer); ok {
		if !h.Equal(target.GetModeOverride()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetModeOverride(), target.GetModeOverride()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetOverrideMessageTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOverrideMessageTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOverrideMessageTimeout(), target.GetOverrideMessageTimeout()) {
			return false
		}
	}

	switch m.Response.(type) {

	case *ProcessingResponse_RequestHeaders:
		if _, ok := target.Response.(*ProcessingResponse_RequestHeaders); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRequestHeaders()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRequestHeaders()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRequestHeaders(), target.GetRequestHeaders()) {
				return false
			}
		}

	case *ProcessingResponse_ResponseHeaders:
		if _, ok := target.Response.(*ProcessingResponse_ResponseHeaders); !ok {
			return false
		}

		if h, ok := interface{}(m.GetResponseHeaders()).(equality.Equalizer); ok {
			if !h.Equal(target.GetResponseHeaders()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetResponseHeaders(), target.GetResponseHeaders()) {
				return false
			}
		}

	case *ProcessingResponse_RequestBody:
		if _, ok := target.Response.(*ProcessingResponse_RequestBody); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRequestBody()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRequestBody()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRequestBody(), target.GetRequestBody()) {
				return false
			}
		}

	case *ProcessingResponse_ResponseBody:
		if _, ok := target.Response.(*ProcessingResponse_ResponseBody); !ok {
			return false
		}

		if h, ok := interface{}(m.GetResponseBody()).(equality.Equalizer); ok {
			if !h.Equal(target.GetResponseBody()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetResponseBody(), target.GetResponseBody()) {
				return false
			}
		}

	case *ProcessingResponse_RequestTrailers:
		if _, ok := target.Response.(*ProcessingResponse_RequestTrailers); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRequestTrailers()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRequestTrailers()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRequestTrailers(), target.GetRequestTrailers()) {
				return false
			}
		}

	case *ProcessingResponse_ResponseTrailers:
		if _, ok := target.Response.(*ProcessingResponse_ResponseTrailers); !ok {
			return false
		}

		if h, ok := interface{}(m.GetResponseTrailers()).(equality.Equalizer); ok {
			if !h.Equal(target.GetResponseTrailers()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetResponseTrailers(), target.GetResponseTrailers()) {
				return false
			}
		}

	case *ProcessingResponse_ImmediateResponse:
		if _, ok := target.Response.(*ProcessingResponse_ImmediateResponse); !ok {
			return false
		}

		if h, ok := interface{}(m.GetImmediateResponse()).(equality.Equalizer); ok {
			if !h.Equal(target.GetImmediateResponse()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetImmediateResponse(), target.GetImmediateResponse()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Response != target.Response {
			return false
		}
	}

	return true
}

// Equal function
func (m *HttpHeaders) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpHeaders)
	if !ok {
		that2, ok := that.(HttpHeaders)
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

	if h, ok := interface{}(m.GetHeaders()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHeaders()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHeaders(), target.GetHeaders()) {
			return false
		}
	}

	if len(m.GetAttributes()) != len(target.GetAttributes()) {
		return false
	}
	for k, v := range m.GetAttributes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAttributes()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAttributes()[k]) {
				return false
			}
		}

	}

	if m.GetEndOfStream() != target.GetEndOfStream() {
		return false
	}

	return true
}

// Equal function
func (m *HttpBody) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpBody)
	if !ok {
		that2, ok := that.(HttpBody)
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

	if bytes.Compare(m.GetBody(), target.GetBody()) != 0 {
		return false
	}

	if m.GetEndOfStream() != target.GetEndOfStream() {
		return false
	}

	return true
}

// Equal function
func (m *HttpTrailers) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpTrailers)
	if !ok {
		that2, ok := that.(HttpTrailers)
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

	if h, ok := interface{}(m.GetTrailers()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTrailers()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTrailers(), target.GetTrailers()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *HeadersResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeadersResponse)
	if !ok {
		that2, ok := that.(HeadersResponse)
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

	if h, ok := interface{}(m.GetResponse()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResponse()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResponse(), target.GetResponse()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *TrailersResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TrailersResponse)
	if !ok {
		that2, ok := that.(TrailersResponse)
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

	if h, ok := interface{}(m.GetHeaderMutation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHeaderMutation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHeaderMutation(), target.GetHeaderMutation()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *BodyResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*BodyResponse)
	if !ok {
		that2, ok := that.(BodyResponse)
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

	if h, ok := interface{}(m.GetResponse()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResponse()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResponse(), target.GetResponse()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *CommonResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CommonResponse)
	if !ok {
		that2, ok := that.(CommonResponse)
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

	if m.GetStatus() != target.GetStatus() {
		return false
	}

	if h, ok := interface{}(m.GetHeaderMutation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHeaderMutation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHeaderMutation(), target.GetHeaderMutation()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetBodyMutation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBodyMutation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBodyMutation(), target.GetBodyMutation()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTrailers()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTrailers()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTrailers(), target.GetTrailers()) {
			return false
		}
	}

	if m.GetClearRouteCache() != target.GetClearRouteCache() {
		return false
	}

	return true
}

// Equal function
func (m *ImmediateResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ImmediateResponse)
	if !ok {
		that2, ok := that.(ImmediateResponse)
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

	if h, ok := interface{}(m.GetStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatus(), target.GetStatus()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHeaders()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHeaders()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHeaders(), target.GetHeaders()) {
			return false
		}
	}

	if bytes.Compare(m.GetBody(), target.GetBody()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetGrpcStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGrpcStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGrpcStatus(), target.GetGrpcStatus()) {
			return false
		}
	}

	if strings.Compare(m.GetDetails(), target.GetDetails()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GrpcStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcStatus)
	if !ok {
		that2, ok := that.(GrpcStatus)
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

	if m.GetStatus() != target.GetStatus() {
		return false
	}

	return true
}

// Equal function
func (m *HeaderMutation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeaderMutation)
	if !ok {
		that2, ok := that.(HeaderMutation)
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

	if len(m.GetSetHeaders()) != len(target.GetSetHeaders()) {
		return false
	}
	for idx, v := range m.GetSetHeaders() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSetHeaders()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSetHeaders()[idx]) {
				return false
			}
		}

	}

	if len(m.GetRemoveHeaders()) != len(target.GetRemoveHeaders()) {
		return false
	}
	for idx, v := range m.GetRemoveHeaders() {

		if strings.Compare(v, target.GetRemoveHeaders()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *BodyMutation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*BodyMutation)
	if !ok {
		that2, ok := that.(BodyMutation)
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

	switch m.Mutation.(type) {

	case *BodyMutation_Body:
		if _, ok := target.Mutation.(*BodyMutation_Body); !ok {
			return false
		}

		if bytes.Compare(m.GetBody(), target.GetBody()) != 0 {
			return false
		}

	case *BodyMutation_ClearBody:
		if _, ok := target.Mutation.(*BodyMutation_ClearBody); !ok {
			return false
		}

		if m.GetClearBody() != target.GetClearBody() {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.Mutation != target.Mutation {
			return false
		}
	}

	return true
}