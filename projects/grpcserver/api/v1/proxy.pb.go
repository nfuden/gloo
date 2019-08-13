// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/grpcserver/api/v1/proxy.proto

package v1

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ProxyDetails struct {
	Proxy                *v1.Proxy `protobuf:"bytes,1,opt,name=proxy,proto3" json:"proxy,omitempty"`
	Raw                  *Raw      `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ProxyDetails) Reset()         { *m = ProxyDetails{} }
func (m *ProxyDetails) String() string { return proto.CompactTextString(m) }
func (*ProxyDetails) ProtoMessage()    {}
func (*ProxyDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_b062e5ddba82b2b7, []int{0}
}
func (m *ProxyDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyDetails.Unmarshal(m, b)
}
func (m *ProxyDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyDetails.Marshal(b, m, deterministic)
}
func (m *ProxyDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyDetails.Merge(m, src)
}
func (m *ProxyDetails) XXX_Size() int {
	return xxx_messageInfo_ProxyDetails.Size(m)
}
func (m *ProxyDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyDetails proto.InternalMessageInfo

func (m *ProxyDetails) GetProxy() *v1.Proxy {
	if m != nil {
		return m.Proxy
	}
	return nil
}

func (m *ProxyDetails) GetRaw() *Raw {
	if m != nil {
		return m.Raw
	}
	return nil
}

type GetProxyRequest struct {
	Ref                  *core.ResourceRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetProxyRequest) Reset()         { *m = GetProxyRequest{} }
func (m *GetProxyRequest) String() string { return proto.CompactTextString(m) }
func (*GetProxyRequest) ProtoMessage()    {}
func (*GetProxyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b062e5ddba82b2b7, []int{1}
}
func (m *GetProxyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProxyRequest.Unmarshal(m, b)
}
func (m *GetProxyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProxyRequest.Marshal(b, m, deterministic)
}
func (m *GetProxyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProxyRequest.Merge(m, src)
}
func (m *GetProxyRequest) XXX_Size() int {
	return xxx_messageInfo_GetProxyRequest.Size(m)
}
func (m *GetProxyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProxyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProxyRequest proto.InternalMessageInfo

func (m *GetProxyRequest) GetRef() *core.ResourceRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

type GetProxyResponse struct {
	ProxyDetails         *ProxyDetails `protobuf:"bytes,1,opt,name=proxy_details,json=proxyDetails,proto3" json:"proxy_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetProxyResponse) Reset()         { *m = GetProxyResponse{} }
func (m *GetProxyResponse) String() string { return proto.CompactTextString(m) }
func (*GetProxyResponse) ProtoMessage()    {}
func (*GetProxyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b062e5ddba82b2b7, []int{2}
}
func (m *GetProxyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProxyResponse.Unmarshal(m, b)
}
func (m *GetProxyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProxyResponse.Marshal(b, m, deterministic)
}
func (m *GetProxyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProxyResponse.Merge(m, src)
}
func (m *GetProxyResponse) XXX_Size() int {
	return xxx_messageInfo_GetProxyResponse.Size(m)
}
func (m *GetProxyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProxyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProxyResponse proto.InternalMessageInfo

func (m *GetProxyResponse) GetProxyDetails() *ProxyDetails {
	if m != nil {
		return m.ProxyDetails
	}
	return nil
}

type ListProxiesRequest struct {
	Namespaces           []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListProxiesRequest) Reset()         { *m = ListProxiesRequest{} }
func (m *ListProxiesRequest) String() string { return proto.CompactTextString(m) }
func (*ListProxiesRequest) ProtoMessage()    {}
func (*ListProxiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b062e5ddba82b2b7, []int{3}
}
func (m *ListProxiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProxiesRequest.Unmarshal(m, b)
}
func (m *ListProxiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProxiesRequest.Marshal(b, m, deterministic)
}
func (m *ListProxiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProxiesRequest.Merge(m, src)
}
func (m *ListProxiesRequest) XXX_Size() int {
	return xxx_messageInfo_ListProxiesRequest.Size(m)
}
func (m *ListProxiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProxiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListProxiesRequest proto.InternalMessageInfo

func (m *ListProxiesRequest) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

type ListProxiesResponse struct {
	ProxyDetails         []*ProxyDetails `protobuf:"bytes,1,rep,name=proxy_details,json=proxyDetails,proto3" json:"proxy_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListProxiesResponse) Reset()         { *m = ListProxiesResponse{} }
func (m *ListProxiesResponse) String() string { return proto.CompactTextString(m) }
func (*ListProxiesResponse) ProtoMessage()    {}
func (*ListProxiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b062e5ddba82b2b7, []int{4}
}
func (m *ListProxiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProxiesResponse.Unmarshal(m, b)
}
func (m *ListProxiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProxiesResponse.Marshal(b, m, deterministic)
}
func (m *ListProxiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProxiesResponse.Merge(m, src)
}
func (m *ListProxiesResponse) XXX_Size() int {
	return xxx_messageInfo_ListProxiesResponse.Size(m)
}
func (m *ListProxiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProxiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListProxiesResponse proto.InternalMessageInfo

func (m *ListProxiesResponse) GetProxyDetails() []*ProxyDetails {
	if m != nil {
		return m.ProxyDetails
	}
	return nil
}

func init() {
	proto.RegisterType((*ProxyDetails)(nil), "glooeeapi.solo.io.ProxyDetails")
	proto.RegisterType((*GetProxyRequest)(nil), "glooeeapi.solo.io.GetProxyRequest")
	proto.RegisterType((*GetProxyResponse)(nil), "glooeeapi.solo.io.GetProxyResponse")
	proto.RegisterType((*ListProxiesRequest)(nil), "glooeeapi.solo.io.ListProxiesRequest")
	proto.RegisterType((*ListProxiesResponse)(nil), "glooeeapi.solo.io.ListProxiesResponse")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-projects/projects/grpcserver/api/v1/proxy.proto", fileDescriptor_b062e5ddba82b2b7)
}

var fileDescriptor_b062e5ddba82b2b7 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xd1, 0x4a, 0xf3, 0x30,
	0x14, 0xc7, 0xbf, 0x7e, 0xe5, 0xfb, 0x98, 0xd9, 0x44, 0xcd, 0x44, 0x66, 0x2f, 0xa6, 0x54, 0x94,
	0x89, 0x9a, 0xe2, 0xf4, 0x42, 0x10, 0x84, 0xc9, 0x60, 0x37, 0x5e, 0x48, 0x41, 0x10, 0x05, 0xa5,
	0xab, 0x67, 0x35, 0xba, 0x2d, 0x31, 0xc9, 0x36, 0xf7, 0x46, 0xbe, 0x86, 0xaf, 0xe2, 0x93, 0x48,
	0xd2, 0xd5, 0x55, 0xd7, 0x29, 0xb2, 0xab, 0x26, 0x3d, 0xff, 0xff, 0xff, 0xfc, 0xce, 0x21, 0xa8,
	0x11, 0x51, 0x75, 0xdf, 0x6b, 0x92, 0x90, 0x75, 0x3c, 0xc9, 0xda, 0x6c, 0x8f, 0xb2, 0xf8, 0xcb,
	0x05, 0x7b, 0x80, 0x50, 0x49, 0xef, 0xe3, 0x10, 0x09, 0x1e, 0x4a, 0x10, 0x7d, 0x10, 0x5e, 0xc0,
	0xa9, 0xd7, 0xdf, 0xd7, 0xa5, 0xe7, 0x21, 0xe1, 0x82, 0x29, 0x86, 0x97, 0xa2, 0x36, 0x63, 0x00,
	0x01, 0xa7, 0x44, 0xfb, 0x09, 0x65, 0xce, 0x72, 0xc4, 0x22, 0x66, 0xaa, 0x9e, 0x3e, 0xc5, 0x42,
	0xe7, 0x28, 0xa3, 0xa3, 0xf6, 0xa6, 0x1a, 0xe9, 0xdb, 0x64, 0x0b, 0x67, 0x77, 0x1a, 0xeb, 0x23,
	0x55, 0x89, 0x5e, 0x40, 0x6b, 0xa4, 0x9e, 0x69, 0x32, 0x35, 0xe4, 0x20, 0xe3, 0x20, 0x37, 0x44,
	0x85, 0x73, 0x4d, 0x51, 0x07, 0x15, 0xd0, 0xb6, 0xc4, 0xdb, 0xe8, 0x9f, 0xa1, 0x2a, 0x59, 0xeb,
	0x56, 0x25, 0x5f, 0x2d, 0x12, 0xcd, 0x9b, 0x0c, 0x4d, 0x8c, 0xd4, 0x8f, 0x15, 0xb8, 0x82, 0x6c,
	0x11, 0x0c, 0x4a, 0x7f, 0x8d, 0x70, 0x85, 0x4c, 0xac, 0x88, 0xf8, 0xc1, 0xc0, 0xd7, 0x12, 0xf7,
	0x04, 0x2d, 0x34, 0x40, 0xc5, 0x66, 0x78, 0xea, 0x81, 0x54, 0x78, 0x07, 0xd9, 0x02, 0x5a, 0xa3,
	0x2e, 0xab, 0x24, 0x64, 0x02, 0xc6, 0x3e, 0x90, 0xac, 0x27, 0x42, 0xf0, 0xa1, 0xe5, 0x6b, 0x95,
	0x7b, 0x89, 0x16, 0xc7, 0x7e, 0xc9, 0x59, 0x57, 0x02, 0xae, 0xa3, 0x79, 0x83, 0x71, 0x7b, 0x17,
	0x93, 0x8f, 0xa2, 0xd6, 0x32, 0x38, 0xd2, 0x03, 0xfa, 0x05, 0x9e, 0xba, 0xb9, 0x87, 0x08, 0x9f,
	0x51, 0x69, 0xa2, 0x29, 0xc8, 0x04, 0xae, 0x8c, 0x50, 0x37, 0xe8, 0x80, 0xe4, 0x41, 0x08, 0x3a,
	0xd8, 0xae, 0xcc, 0xf9, 0xa9, 0x3f, 0xee, 0x35, 0x2a, 0x7e, 0x72, 0x4d, 0x47, 0xb2, 0x7f, 0x8d,
	0x54, 0x7d, 0xb5, 0x50, 0xce, 0x94, 0x6b, 0x9c, 0xe2, 0x0b, 0x94, 0x4b, 0x26, 0xc7, 0x6e, 0x46,
	0xce, 0x97, 0xb5, 0x3a, 0x1b, 0xdf, 0x6a, 0x62, 0x4e, 0xf7, 0x0f, 0xbe, 0x41, 0xf9, 0xd4, 0x00,
	0x78, 0x33, 0xc3, 0x35, 0xb9, 0x16, 0x67, 0xeb, 0x27, 0x59, 0x92, 0x7f, 0x5a, 0x7b, 0x79, 0x2b,
	0x5b, 0x57, 0xc7, 0x33, 0x3c, 0xd2, 0xe6, 0x7f, 0xf3, 0x3e, 0x0f, 0xde, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xc3, 0x30, 0x2e, 0x17, 0xc4, 0x03, 0x00, 0x00,
}

func (this *ProxyDetails) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProxyDetails)
	if !ok {
		that2, ok := that.(ProxyDetails)
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
	if !this.Proxy.Equal(that1.Proxy) {
		return false
	}
	if !this.Raw.Equal(that1.Raw) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GetProxyRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetProxyRequest)
	if !ok {
		that2, ok := that.(GetProxyRequest)
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
	if !this.Ref.Equal(that1.Ref) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GetProxyResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetProxyResponse)
	if !ok {
		that2, ok := that.(GetProxyResponse)
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
	if !this.ProxyDetails.Equal(that1.ProxyDetails) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ListProxiesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListProxiesRequest)
	if !ok {
		that2, ok := that.(ListProxiesRequest)
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
	if len(this.Namespaces) != len(that1.Namespaces) {
		return false
	}
	for i := range this.Namespaces {
		if this.Namespaces[i] != that1.Namespaces[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ListProxiesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListProxiesResponse)
	if !ok {
		that2, ok := that.(ListProxiesResponse)
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
	if len(this.ProxyDetails) != len(that1.ProxyDetails) {
		return false
	}
	for i := range this.ProxyDetails {
		if !this.ProxyDetails[i].Equal(that1.ProxyDetails[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProxyApiClient is the client API for ProxyApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxyApiClient interface {
	GetProxy(ctx context.Context, in *GetProxyRequest, opts ...grpc.CallOption) (*GetProxyResponse, error)
	ListProxies(ctx context.Context, in *ListProxiesRequest, opts ...grpc.CallOption) (*ListProxiesResponse, error)
}

type proxyApiClient struct {
	cc *grpc.ClientConn
}

func NewProxyApiClient(cc *grpc.ClientConn) ProxyApiClient {
	return &proxyApiClient{cc}
}

func (c *proxyApiClient) GetProxy(ctx context.Context, in *GetProxyRequest, opts ...grpc.CallOption) (*GetProxyResponse, error) {
	out := new(GetProxyResponse)
	err := c.cc.Invoke(ctx, "/glooeeapi.solo.io.ProxyApi/GetProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyApiClient) ListProxies(ctx context.Context, in *ListProxiesRequest, opts ...grpc.CallOption) (*ListProxiesResponse, error) {
	out := new(ListProxiesResponse)
	err := c.cc.Invoke(ctx, "/glooeeapi.solo.io.ProxyApi/ListProxies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyApiServer is the server API for ProxyApi service.
type ProxyApiServer interface {
	GetProxy(context.Context, *GetProxyRequest) (*GetProxyResponse, error)
	ListProxies(context.Context, *ListProxiesRequest) (*ListProxiesResponse, error)
}

// UnimplementedProxyApiServer can be embedded to have forward compatible implementations.
type UnimplementedProxyApiServer struct {
}

func (*UnimplementedProxyApiServer) GetProxy(ctx context.Context, req *GetProxyRequest) (*GetProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProxy not implemented")
}
func (*UnimplementedProxyApiServer) ListProxies(ctx context.Context, req *ListProxiesRequest) (*ListProxiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProxies not implemented")
}

func RegisterProxyApiServer(s *grpc.Server, srv ProxyApiServer) {
	s.RegisterService(&_ProxyApi_serviceDesc, srv)
}

func _ProxyApi_GetProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyApiServer).GetProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glooeeapi.solo.io.ProxyApi/GetProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyApiServer).GetProxy(ctx, req.(*GetProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProxyApi_ListProxies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProxiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyApiServer).ListProxies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glooeeapi.solo.io.ProxyApi/ListProxies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyApiServer).ListProxies(ctx, req.(*ListProxiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProxyApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "glooeeapi.solo.io.ProxyApi",
	HandlerType: (*ProxyApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProxy",
			Handler:    _ProxyApi_GetProxy_Handler,
		},
		{
			MethodName: "ListProxies",
			Handler:    _ProxyApi_ListProxies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/solo-io/solo-projects/projects/grpcserver/api/v1/proxy.proto",
}
