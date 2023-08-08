// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/matcher/v3/string.proto

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
func (m *StringMatcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*StringMatcher)
	if !ok {
		that2, ok := that.(StringMatcher)
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

	if m.GetIgnoreCase() != target.GetIgnoreCase() {
		return false
	}

	switch m.MatchPattern.(type) {

	case *StringMatcher_Exact:
		if _, ok := target.MatchPattern.(*StringMatcher_Exact); !ok {
			return false
		}

		if strings.Compare(m.GetExact(), target.GetExact()) != 0 {
			return false
		}

	case *StringMatcher_Prefix:
		if _, ok := target.MatchPattern.(*StringMatcher_Prefix); !ok {
			return false
		}

		if strings.Compare(m.GetPrefix(), target.GetPrefix()) != 0 {
			return false
		}

	case *StringMatcher_Suffix:
		if _, ok := target.MatchPattern.(*StringMatcher_Suffix); !ok {
			return false
		}

		if strings.Compare(m.GetSuffix(), target.GetSuffix()) != 0 {
			return false
		}

	case *StringMatcher_SafeRegex:
		if _, ok := target.MatchPattern.(*StringMatcher_SafeRegex); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSafeRegex()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSafeRegex()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSafeRegex(), target.GetSafeRegex()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.MatchPattern != target.MatchPattern {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListStringMatcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListStringMatcher)
	if !ok {
		that2, ok := that.(ListStringMatcher)
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

	if len(m.GetPatterns()) != len(target.GetPatterns()) {
		return false
	}
	for idx, v := range m.GetPatterns() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPatterns()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPatterns()[idx]) {
				return false
			}
		}

	}

	return true
}