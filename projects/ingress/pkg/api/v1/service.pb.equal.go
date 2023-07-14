// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/ingress/api/v1/service.proto

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
func (m *KubeService) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*KubeService)
	if !ok {
		that2, ok := that.(KubeService)
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

	if h, ok := interface{}(m.GetKubeServiceSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetKubeServiceSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetKubeServiceSpec(), target.GetKubeServiceSpec()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetKubeServiceStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetKubeServiceStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetKubeServiceStatus(), target.GetKubeServiceStatus()) {
			return false
		}
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

	return true
}