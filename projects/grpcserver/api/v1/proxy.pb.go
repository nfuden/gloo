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
	Proxy *v1.Proxy `protobuf:"bytes,1,opt,name=proxy,proto3" json:"proxy,omitempty"`
	Raw   *Raw      `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	// Status of the instance and its relationship with Gloo.
	// Currently there are three possible values for the "Code" field on the status message.
	// In this context, their values have the following definitions:
	// ERROR:
	//   - The proxy CRD has status "Rejected"
	// WARNING:
	//   - The proxy CRD has status "Pending"
	// OK:
	//   - The proxy CRD has status "Accepted"
	Status               *Status  `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *ProxyDetails) GetStatus() *Status {
	if m != nil {
		return m.Status
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
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xd1, 0xca, 0xd3, 0x30,
	0x14, 0x80, 0xad, 0xc5, 0x9f, 0xdf, 0xb3, 0x89, 0x9a, 0x89, 0x6c, 0xbd, 0x98, 0x52, 0x51, 0x26,
	0x6a, 0xca, 0xa6, 0x17, 0x82, 0x20, 0x4c, 0x06, 0xbb, 0xf1, 0x42, 0x22, 0x82, 0x28, 0x28, 0x59,
	0x3d, 0xab, 0xd1, 0x6d, 0x89, 0x49, 0xba, 0xb9, 0x17, 0xf1, 0x19, 0x7c, 0x0d, 0x5f, 0xc5, 0x27,
	0x91, 0xa4, 0xad, 0xab, 0xae, 0x53, 0x7e, 0x76, 0xd5, 0xa4, 0xe7, 0x3b, 0xe7, 0x7c, 0xe7, 0x10,
	0x98, 0x66, 0xc2, 0x7e, 0xcc, 0x67, 0x34, 0x95, 0xcb, 0xc4, 0xc8, 0x85, 0x7c, 0x20, 0x64, 0xf1,
	0x55, 0x5a, 0x7e, 0xc2, 0xd4, 0x9a, 0xe4, 0xf7, 0x21, 0xd3, 0x2a, 0x35, 0xa8, 0xd7, 0xa8, 0x13,
	0xae, 0x44, 0xb2, 0x1e, 0xba, 0xd0, 0xd7, 0x2d, 0x55, 0x5a, 0x5a, 0x49, 0xae, 0x66, 0x0b, 0x29,
	0x11, 0xb9, 0x12, 0xd4, 0xe5, 0x53, 0x21, 0xa3, 0x6b, 0x99, 0xcc, 0xa4, 0x8f, 0x26, 0xee, 0x54,
	0x80, 0xd1, 0xe3, 0x86, 0x8e, 0x2e, 0xb7, 0xd6, 0xc8, 0xdd, 0xf6, 0x5b, 0x44, 0xf7, 0x0f, 0xb9,
	0x7e, 0x16, 0xb6, 0xe2, 0x35, 0xce, 0x4b, 0xfa, 0xa8, 0xc9, 0xec, 0x56, 0xa1, 0x29, 0x0a, 0xc5,
	0xdf, 0x02, 0x68, 0xbf, 0x70, 0x1a, 0x13, 0xb4, 0x5c, 0x2c, 0x0c, 0xb9, 0x0b, 0x17, 0xbc, 0x56,
	0x37, 0xb8, 0x19, 0x0c, 0x5a, 0xa3, 0x0e, 0x75, 0xc2, 0xd5, 0xd4, 0xd4, 0xa3, 0xac, 0x20, 0xc8,
	0x00, 0x42, 0xcd, 0x37, 0xdd, 0xf3, 0x1e, 0xbc, 0x4e, 0xf7, 0x76, 0x44, 0x19, 0xdf, 0x30, 0x87,
	0x90, 0x21, 0x9c, 0x18, 0xcb, 0x6d, 0x6e, 0xba, 0xa1, 0x87, 0x7b, 0x0d, 0xf0, 0x4b, 0x0f, 0xb0,
	0x12, 0x8c, 0x9f, 0xc2, 0xe5, 0x29, 0xda, 0xa2, 0x1f, 0x7e, 0xc9, 0xd1, 0x58, 0x72, 0x0f, 0x42,
	0x8d, 0xf3, 0x52, 0xac, 0x47, 0x53, 0xa9, 0x71, 0xd7, 0x0a, 0x8d, 0xcc, 0x75, 0x8a, 0x0c, 0xe7,
	0xcc, 0x51, 0xf1, 0x6b, 0xb8, 0xb2, 0xcb, 0x37, 0x4a, 0xae, 0x0c, 0x92, 0x09, 0x5c, 0xf2, 0xe6,
	0xef, 0x3f, 0x14, 0xc3, 0x96, 0xa5, 0x6e, 0x34, 0xd8, 0xd4, 0x77, 0xc2, 0xda, 0xaa, 0x76, 0x8b,
	0x1f, 0x01, 0x79, 0x2e, 0x8c, 0x2f, 0x2d, 0xd0, 0x54, 0x72, 0x7d, 0x80, 0x15, 0x5f, 0xa2, 0x51,
	0x3c, 0x45, 0x57, 0x38, 0x1c, 0x5c, 0x64, 0xb5, 0x3f, 0xf1, 0x5b, 0xe8, 0xfc, 0x91, 0x75, 0x58,
	0x29, 0x3c, 0xb3, 0xd2, 0xe8, 0x47, 0x00, 0xa7, 0x3e, 0x3c, 0x56, 0x82, 0xbc, 0x82, 0xd3, 0x6a,
	0x72, 0x12, 0x37, 0xd4, 0xf9, 0x6b, 0xad, 0xd1, 0xad, 0x7f, 0x32, 0x85, 0x67, 0x7c, 0x8e, 0xbc,
	0x83, 0x56, 0x6d, 0x00, 0x72, 0xbb, 0x21, 0x6b, 0x7f, 0x2d, 0xd1, 0x9d, 0xff, 0x61, 0x55, 0xfd,
	0x67, 0xe3, 0xef, 0x3f, 0xfb, 0xc1, 0x9b, 0x27, 0x47, 0x3c, 0xec, 0xd9, 0x89, 0x7f, 0xd3, 0x0f,
	0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x31, 0xa7, 0x25, 0x07, 0xf8, 0x03, 0x00, 0x00,
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
	if !this.Status.Equal(that1.Status) {
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
