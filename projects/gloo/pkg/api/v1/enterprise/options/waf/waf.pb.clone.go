// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/waf/waf.proto

package waf

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_waf "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/waf"

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
func (m *Settings) Clone() proto.Message {
	var target *Settings
	if m == nil {
		return target
	}
	target = &Settings{}

	target.Disabled = m.GetDisabled()

	target.CustomInterventionMessage = m.GetCustomInterventionMessage()

	if h, ok := interface{}(m.GetCoreRuleSet()).(clone.Cloner); ok {
		target.CoreRuleSet = h.Clone().(*CoreRuleSet)
	} else {
		target.CoreRuleSet = proto.Clone(m.GetCoreRuleSet()).(*CoreRuleSet)
	}

	if m.GetRuleSets() != nil {
		target.RuleSets = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_waf.RuleSet, len(m.GetRuleSets()))
		for idx, v := range m.GetRuleSets() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.RuleSets[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_waf.RuleSet)
			} else {
				target.RuleSets[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_waf.RuleSet)
			}

		}
	}

	if m.GetConfigMapRuleSets() != nil {
		target.ConfigMapRuleSets = make([]*RuleSetFromConfigMap, len(m.GetConfigMapRuleSets()))
		for idx, v := range m.GetConfigMapRuleSets() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ConfigMapRuleSets[idx] = h.Clone().(*RuleSetFromConfigMap)
			} else {
				target.ConfigMapRuleSets[idx] = proto.Clone(v).(*RuleSetFromConfigMap)
			}

		}
	}

	if h, ok := interface{}(m.GetAuditLogging()).(clone.Cloner); ok {
		target.AuditLogging = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_waf.AuditLogging)
	} else {
		target.AuditLogging = proto.Clone(m.GetAuditLogging()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_waf.AuditLogging)
	}

	target.RequestHeadersOnly = m.GetRequestHeadersOnly()

	target.ResponseHeadersOnly = m.GetResponseHeadersOnly()

	return target
}

// Clone function
func (m *RuleSetFromConfigMap) Clone() proto.Message {
	var target *RuleSetFromConfigMap
	if m == nil {
		return target
	}
	target = &RuleSetFromConfigMap{}

	if h, ok := interface{}(m.GetConfigMapRef()).(clone.Cloner); ok {
		target.ConfigMapRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.ConfigMapRef = proto.Clone(m.GetConfigMapRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	if m.GetDataMapKeys() != nil {
		target.DataMapKeys = make([]string, len(m.GetDataMapKeys()))
		for idx, v := range m.GetDataMapKeys() {

			target.DataMapKeys[idx] = v

		}
	}

	return target
}

// Clone function
func (m *CoreRuleSet) Clone() proto.Message {
	var target *CoreRuleSet
	if m == nil {
		return target
	}
	target = &CoreRuleSet{}

	switch m.CustomSettingsType.(type) {

	case *CoreRuleSet_CustomSettingsString:

		target.CustomSettingsType = &CoreRuleSet_CustomSettingsString{
			CustomSettingsString: m.GetCustomSettingsString(),
		}

	case *CoreRuleSet_CustomSettingsFile:

		target.CustomSettingsType = &CoreRuleSet_CustomSettingsFile{
			CustomSettingsFile: m.GetCustomSettingsFile(),
		}

	}

	return target
}