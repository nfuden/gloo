// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/static/static.proto

package static

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
func (m *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
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

	if h, ok := interface{}(m.GetUseTls()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUseTls()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUseTls(), target.GetUseTls()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetServiceSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetServiceSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetServiceSpec(), target.GetServiceSpec()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAutoSniRewrite()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAutoSniRewrite()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAutoSniRewrite(), target.GetAutoSniRewrite()) {
			return false
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

	if strings.Compare(m.GetAddr(), target.GetAddr()) != 0 {
		return false
	}

	if m.GetPort() != target.GetPort() {
		return false
	}

	if strings.Compare(m.GetSniAddr(), target.GetSniAddr()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetLoadBalancingWeight()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLoadBalancingWeight()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLoadBalancingWeight(), target.GetLoadBalancingWeight()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHealthCheckConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHealthCheckConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHealthCheckConfig(), target.GetHealthCheckConfig()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Host_HealthCheckConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Host_HealthCheckConfig)
	if !ok {
		that2, ok := that.(Host_HealthCheckConfig)
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

	if strings.Compare(m.GetPath(), target.GetPath()) != 0 {
		return false
	}

	if strings.Compare(m.GetMethod(), target.GetMethod()) != 0 {
		return false
	}

	return true
}