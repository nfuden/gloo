// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/plugins/rbac/rbac.proto

package rbac

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// A JWT principal. To use this, JWT plugin MUST be enabled.
type JWTPrincipal struct {
	// Set of claims that make up this principal. Commonly, the 'iss' and 'sub' or 'email' claims are used.
	// all claims must be present on the JWT.
	Claims map[string]string `protobuf:"bytes,1,rep,name=claims,proto3" json:"claims,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Verify that the JWT came from a specific provider. This usually can be left empty
	// and a provider will be chosen automatically.
	Provider             string   `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWTPrincipal) Reset()         { *m = JWTPrincipal{} }
func (m *JWTPrincipal) String() string { return proto.CompactTextString(m) }
func (*JWTPrincipal) ProtoMessage()    {}
func (*JWTPrincipal) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b892a5350c8434, []int{0}
}
func (m *JWTPrincipal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWTPrincipal.Unmarshal(m, b)
}
func (m *JWTPrincipal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWTPrincipal.Marshal(b, m, deterministic)
}
func (m *JWTPrincipal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTPrincipal.Merge(m, src)
}
func (m *JWTPrincipal) XXX_Size() int {
	return xxx_messageInfo_JWTPrincipal.Size(m)
}
func (m *JWTPrincipal) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTPrincipal.DiscardUnknown(m)
}

var xxx_messageInfo_JWTPrincipal proto.InternalMessageInfo

func (m *JWTPrincipal) GetClaims() map[string]string {
	if m != nil {
		return m.Claims
	}
	return nil
}

func (m *JWTPrincipal) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

// An RBAC principal - the identity enitity (usually a user or a service account).
type Principal struct {
	JwtPrincipal         *JWTPrincipal `protobuf:"bytes,1,opt,name=jwt_principal,json=jwtPrincipal,proto3" json:"jwt_principal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b892a5350c8434, []int{1}
}
func (m *Principal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal.Unmarshal(m, b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
}
func (m *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(m, src)
}
func (m *Principal) XXX_Size() int {
	return xxx_messageInfo_Principal.Size(m)
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

func (m *Principal) GetJwtPrincipal() *JWTPrincipal {
	if m != nil {
		return m.JwtPrincipal
	}
	return nil
}

// What permissions should be granted. An empty field means allow-all.
// If more than one field is added, all of them need to match.
type Permissions struct {
	// Paths that have this prefix will be allowed.
	PathPrefix string `protobuf:"bytes,1,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	// What http methods (GET, POST, ...) are allowed.
	Methods              []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permissions) Reset()         { *m = Permissions{} }
func (m *Permissions) String() string { return proto.CompactTextString(m) }
func (*Permissions) ProtoMessage()    {}
func (*Permissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b892a5350c8434, []int{2}
}
func (m *Permissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permissions.Unmarshal(m, b)
}
func (m *Permissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permissions.Marshal(b, m, deterministic)
}
func (m *Permissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permissions.Merge(m, src)
}
func (m *Permissions) XXX_Size() int {
	return xxx_messageInfo_Permissions.Size(m)
}
func (m *Permissions) XXX_DiscardUnknown() {
	xxx_messageInfo_Permissions.DiscardUnknown(m)
}

var xxx_messageInfo_Permissions proto.InternalMessageInfo

func (m *Permissions) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

func (m *Permissions) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

type Policy struct {
	// Principals in this policy.
	Principals []*Principal `protobuf:"bytes,1,rep,name=principals,proto3" json:"principals,omitempty"`
	// Permissions granted to the principals.
	Permissions          *Permissions `protobuf:"bytes,2,opt,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b892a5350c8434, []int{3}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetPrincipals() []*Principal {
	if m != nil {
		return m.Principals
	}
	return nil
}

func (m *Policy) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type Settings struct {
	// Require RBAC for all vhosts. A vhost without an RBAC policy set will fallback to a deny-all policy.
	RequireRbac          bool     `protobuf:"varint,1,opt,name=require_rbac,json=requireRbac,proto3" json:"require_rbac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b892a5350c8434, []int{4}
}
func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetRequireRbac() bool {
	if m != nil {
		return m.RequireRbac
	}
	return false
}

type Config struct {
	// Named policies to apply.
	Policies             map[string]*Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b892a5350c8434, []int{5}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetPolicies() map[string]*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type VhostExtension struct {
	Config               *Config  `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VhostExtension) Reset()         { *m = VhostExtension{} }
func (m *VhostExtension) String() string { return proto.CompactTextString(m) }
func (*VhostExtension) ProtoMessage()    {}
func (*VhostExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b892a5350c8434, []int{6}
}
func (m *VhostExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VhostExtension.Unmarshal(m, b)
}
func (m *VhostExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VhostExtension.Marshal(b, m, deterministic)
}
func (m *VhostExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VhostExtension.Merge(m, src)
}
func (m *VhostExtension) XXX_Size() int {
	return xxx_messageInfo_VhostExtension.Size(m)
}
func (m *VhostExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_VhostExtension.DiscardUnknown(m)
}

var xxx_messageInfo_VhostExtension proto.InternalMessageInfo

func (m *VhostExtension) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

type RouteExtension struct {
	// Types that are valid to be assigned to Route:
	//	*RouteExtension_Disable
	//	*RouteExtension_Config
	Route                isRouteExtension_Route `protobuf_oneof:"route"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RouteExtension) Reset()         { *m = RouteExtension{} }
func (m *RouteExtension) String() string { return proto.CompactTextString(m) }
func (*RouteExtension) ProtoMessage()    {}
func (*RouteExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b892a5350c8434, []int{7}
}
func (m *RouteExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteExtension.Unmarshal(m, b)
}
func (m *RouteExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteExtension.Marshal(b, m, deterministic)
}
func (m *RouteExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteExtension.Merge(m, src)
}
func (m *RouteExtension) XXX_Size() int {
	return xxx_messageInfo_RouteExtension.Size(m)
}
func (m *RouteExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteExtension.DiscardUnknown(m)
}

var xxx_messageInfo_RouteExtension proto.InternalMessageInfo

type isRouteExtension_Route interface {
	isRouteExtension_Route()
	Equal(interface{}) bool
}

type RouteExtension_Disable struct {
	Disable bool `protobuf:"varint,1,opt,name=disable,proto3,oneof" json:"disable,omitempty"`
}
type RouteExtension_Config struct {
	Config *Config `protobuf:"bytes,2,opt,name=config,proto3,oneof" json:"config,omitempty"`
}

func (*RouteExtension_Disable) isRouteExtension_Route() {}
func (*RouteExtension_Config) isRouteExtension_Route()  {}

func (m *RouteExtension) GetRoute() isRouteExtension_Route {
	if m != nil {
		return m.Route
	}
	return nil
}

func (m *RouteExtension) GetDisable() bool {
	if x, ok := m.GetRoute().(*RouteExtension_Disable); ok {
		return x.Disable
	}
	return false
}

func (m *RouteExtension) GetConfig() *Config {
	if x, ok := m.GetRoute().(*RouteExtension_Config); ok {
		return x.Config
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteExtension) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteExtension_Disable)(nil),
		(*RouteExtension_Config)(nil),
	}
}

func init() {
	proto.RegisterType((*JWTPrincipal)(nil), "rbac.plugins.gloo.solo.io.JWTPrincipal")
	proto.RegisterMapType((map[string]string)(nil), "rbac.plugins.gloo.solo.io.JWTPrincipal.ClaimsEntry")
	proto.RegisterType((*Principal)(nil), "rbac.plugins.gloo.solo.io.Principal")
	proto.RegisterType((*Permissions)(nil), "rbac.plugins.gloo.solo.io.Permissions")
	proto.RegisterType((*Policy)(nil), "rbac.plugins.gloo.solo.io.Policy")
	proto.RegisterType((*Settings)(nil), "rbac.plugins.gloo.solo.io.Settings")
	proto.RegisterType((*Config)(nil), "rbac.plugins.gloo.solo.io.Config")
	proto.RegisterMapType((map[string]*Policy)(nil), "rbac.plugins.gloo.solo.io.Config.PoliciesEntry")
	proto.RegisterType((*VhostExtension)(nil), "rbac.plugins.gloo.solo.io.VhostExtension")
	proto.RegisterType((*RouteExtension)(nil), "rbac.plugins.gloo.solo.io.RouteExtension")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/plugins/rbac/rbac.proto", fileDescriptor_51b892a5350c8434)
}

var fileDescriptor_51b892a5350c8434 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x53, 0x35, 0x3f, 0xe3, 0xb4, 0x42, 0xab, 0x1e, 0x42, 0x0e, 0x90, 0x58, 0x08, 0x72,
	0xa9, 0x2d, 0xd2, 0x03, 0x14, 0x6e, 0x2d, 0x95, 0x22, 0x82, 0xc0, 0x5a, 0x10, 0x08, 0x0e, 0x44,
	0xb6, 0xb3, 0x75, 0x36, 0x75, 0xbc, 0xcb, 0xee, 0x3a, 0x6d, 0xde, 0x84, 0x47, 0xe0, 0x8c, 0xc4,
	0x0b, 0xf1, 0x24, 0xc8, 0xeb, 0x5f, 0x24, 0x48, 0x73, 0xb1, 0x76, 0x66, 0xe7, 0xfb, 0xe6, 0x9b,
	0xd1, 0xe7, 0x85, 0x77, 0x21, 0x55, 0x8b, 0xc4, 0xb7, 0x03, 0xb6, 0x72, 0x24, 0x8b, 0xd8, 0x09,
	0x65, 0x4e, 0x18, 0x31, 0xe6, 0x70, 0xc1, 0x96, 0x24, 0x50, 0x32, 0x8b, 0x3c, 0x4e, 0x9d, 0xf5,
	0x53, 0x87, 0xc4, 0x8a, 0x08, 0x2e, 0xa8, 0x24, 0x0e, 0x8f, 0x92, 0x90, 0xc6, 0xd2, 0x11, 0xbe,
	0x17, 0xe8, 0x8f, 0xcd, 0x05, 0x53, 0x0c, 0xdd, 0xcf, 0xce, 0xd9, 0xad, 0x9d, 0x82, 0xed, 0x94,
	0xd7, 0xa6, 0xac, 0x7f, 0x1c, 0xb2, 0x90, 0xe9, 0x2a, 0x27, 0x3d, 0x65, 0x00, 0xeb, 0xa7, 0x01,
	0xdd, 0xd7, 0x9f, 0x3e, 0xb8, 0x82, 0xc6, 0x01, 0xe5, 0x5e, 0x84, 0xa6, 0xd0, 0x0c, 0x22, 0x8f,
	0xae, 0x64, 0xcf, 0x18, 0xec, 0x8f, 0xcc, 0xf1, 0xa9, 0xfd, 0x5f, 0x4a, 0xbb, 0x0e, 0xb4, 0x2f,
	0x34, 0xea, 0x32, 0x56, 0x62, 0x83, 0x73, 0x0a, 0xd4, 0x87, 0x36, 0x17, 0x6c, 0x4d, 0xe7, 0x44,
	0xf4, 0x1a, 0x03, 0x63, 0xd4, 0xc1, 0x65, 0xdc, 0x3f, 0x03, 0xb3, 0x06, 0x41, 0xf7, 0x60, 0xff,
	0x9a, 0x6c, 0x7a, 0x86, 0xae, 0x4a, 0x8f, 0xe8, 0x18, 0x0e, 0xd6, 0x5e, 0x94, 0x90, 0x1c, 0x99,
	0x05, 0x2f, 0x1a, 0xcf, 0x0d, 0xeb, 0x33, 0x74, 0x2a, 0xc1, 0x6f, 0xe0, 0x70, 0x79, 0xa3, 0x66,
	0xbc, 0x48, 0x68, 0x0a, 0x73, 0xfc, 0x64, 0x47, 0xdd, 0xb8, 0xbb, 0xbc, 0x51, 0x65, 0x64, 0x4d,
	0xc0, 0x74, 0x89, 0x58, 0x51, 0x29, 0x29, 0x8b, 0x25, 0x7a, 0x08, 0x26, 0xf7, 0xd4, 0x62, 0xc6,
	0x05, 0xb9, 0xa2, 0xb7, 0xb9, 0x3a, 0x48, 0x53, 0xae, 0xce, 0xa0, 0x1e, 0xb4, 0x56, 0x44, 0x2d,
	0xd8, 0x5c, 0xf6, 0x1a, 0x83, 0xfd, 0x51, 0x07, 0x17, 0xa1, 0xf5, 0xdd, 0x80, 0xa6, 0xcb, 0x22,
	0x1a, 0x6c, 0xd0, 0x2b, 0x80, 0x52, 0x5e, 0xb1, 0xd7, 0x47, 0x5b, 0xf4, 0x55, 0xe2, 0x6a, 0x38,
	0x34, 0x01, 0x93, 0x57, 0xd2, 0xf4, 0x56, 0xcc, 0xf1, 0xe3, 0x6d, 0x34, 0x55, 0x35, 0xae, 0x43,
	0xad, 0x13, 0x68, 0xbf, 0x27, 0x4a, 0xd1, 0x38, 0x94, 0x68, 0x08, 0x5d, 0x41, 0xbe, 0x25, 0x54,
	0x90, 0x59, 0xca, 0xa4, 0x47, 0x6c, 0x63, 0x33, 0xcf, 0x61, 0xdf, 0x0b, 0xac, 0x5f, 0x06, 0x34,
	0x2f, 0x58, 0x7c, 0x45, 0x43, 0x34, 0x85, 0x36, 0x4f, 0x67, 0xa2, 0xa4, 0x98, 0xc3, 0xd9, 0x22,
	0x20, 0x03, 0xd9, 0x6e, 0x8e, 0xc8, 0xbc, 0x51, 0x12, 0xf4, 0xbf, 0xc2, 0xe1, 0x5f, 0x57, 0xff,
	0xf0, 0xc0, 0xb3, 0xba, 0x07, 0xcc, 0xf1, 0x70, 0xdb, 0xb4, 0x7a, 0xd7, 0x75, 0x9b, 0x4c, 0xe1,
	0xe8, 0xe3, 0x82, 0x49, 0x75, 0x79, 0xab, 0x48, 0x9c, 0x4e, 0x8e, 0xce, 0xa0, 0x19, 0x68, 0x4d,
	0xb9, 0x49, 0x86, 0x77, 0x8a, 0xc7, 0x39, 0xc0, 0x12, 0x70, 0x84, 0x59, 0xa2, 0x48, 0x45, 0xd6,
	0x87, 0xd6, 0x9c, 0x4a, 0xcf, 0x8f, 0x48, 0xb6, 0xb4, 0xc9, 0x1e, 0x2e, 0x12, 0xe8, 0x65, 0xd9,
	0xa8, 0xb1, 0x63, 0xa3, 0xc9, 0x5e, 0xd1, 0xea, 0xbc, 0x05, 0x07, 0x22, 0x6d, 0x75, 0xfe, 0xf6,
	0xc7, 0xef, 0x07, 0xc6, 0x97, 0xc9, 0x6e, 0x8f, 0x04, 0xbf, 0x0e, 0xef, 0x78, 0x28, 0xfc, 0xa6,
	0xfe, 0xe7, 0x4f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x03, 0xf2, 0x1f, 0xf1, 0x77, 0x04, 0x00,
	0x00,
}

func (this *JWTPrincipal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*JWTPrincipal)
	if !ok {
		that2, ok := that.(JWTPrincipal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Claims) != len(that1.Claims) {
		return false
	}
	for i := range this.Claims {
		if this.Claims[i] != that1.Claims[i] {
			return false
		}
	}
	if this.Provider != that1.Provider {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Principal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Principal)
	if !ok {
		that2, ok := that.(Principal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.JwtPrincipal.Equal(that1.JwtPrincipal) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Permissions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Permissions)
	if !ok {
		that2, ok := that.(Permissions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PathPrefix != that1.PathPrefix {
		return false
	}
	if len(this.Methods) != len(that1.Methods) {
		return false
	}
	for i := range this.Methods {
		if this.Methods[i] != that1.Methods[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Policy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Policy)
	if !ok {
		that2, ok := that.(Policy)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Principals) != len(that1.Principals) {
		return false
	}
	for i := range this.Principals {
		if !this.Principals[i].Equal(that1.Principals[i]) {
			return false
		}
	}
	if !this.Permissions.Equal(that1.Permissions) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Settings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings)
	if !ok {
		that2, ok := that.(Settings)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.RequireRbac != that1.RequireRbac {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Config) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Config)
	if !ok {
		that2, ok := that.(Config)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Policies) != len(that1.Policies) {
		return false
	}
	for i := range this.Policies {
		if !this.Policies[i].Equal(that1.Policies[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VhostExtension) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VhostExtension)
	if !ok {
		that2, ok := that.(VhostExtension)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Config.Equal(that1.Config) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteExtension) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteExtension)
	if !ok {
		that2, ok := that.(RouteExtension)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Route == nil {
		if this.Route != nil {
			return false
		}
	} else if this.Route == nil {
		return false
	} else if !this.Route.Equal(that1.Route) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteExtension_Disable) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteExtension_Disable)
	if !ok {
		that2, ok := that.(RouteExtension_Disable)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Disable != that1.Disable {
		return false
	}
	return true
}
func (this *RouteExtension_Config) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteExtension_Config)
	if !ok {
		that2, ok := that.(RouteExtension_Config)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Config.Equal(that1.Config) {
		return false
	}
	return true
}
