// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/grpcserver/api/v1/envoy.proto

package v1

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EnvoyDetails struct {
	// The name of the gateway-proxy pod running this instance of envoy.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Raw representation of this instance's config_dump.
	Raw *Raw `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	// Status of the instance and its relationship with Gloo.
	// Currently there are three possible values for the "Code" field on the status message.
	// In this context, their values have the following definitions:
	// ERROR:
	//   - The Proxy CRD associated with the envoy instance has status "Rejected"
	//   - OR the gateway-proxy pod envoy runs on has a phase other than "Running"
	// WARNING:
	//   - The Proxy CRD associated with the envoy instance has status "Pending"
	// OK:
	//   - The gateway-proxy pod is running and the proxy CRD has status "Accepted"
	Status               *Status  `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvoyDetails) Reset()         { *m = EnvoyDetails{} }
func (m *EnvoyDetails) String() string { return proto.CompactTextString(m) }
func (*EnvoyDetails) ProtoMessage()    {}
func (*EnvoyDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_a89e3a95b2dd1a92, []int{0}
}
func (m *EnvoyDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvoyDetails.Unmarshal(m, b)
}
func (m *EnvoyDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvoyDetails.Marshal(b, m, deterministic)
}
func (m *EnvoyDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvoyDetails.Merge(m, src)
}
func (m *EnvoyDetails) XXX_Size() int {
	return xxx_messageInfo_EnvoyDetails.Size(m)
}
func (m *EnvoyDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvoyDetails.DiscardUnknown(m)
}

var xxx_messageInfo_EnvoyDetails proto.InternalMessageInfo

func (m *EnvoyDetails) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EnvoyDetails) GetRaw() *Raw {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *EnvoyDetails) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type ListEnvoyDetailsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListEnvoyDetailsRequest) Reset()         { *m = ListEnvoyDetailsRequest{} }
func (m *ListEnvoyDetailsRequest) String() string { return proto.CompactTextString(m) }
func (*ListEnvoyDetailsRequest) ProtoMessage()    {}
func (*ListEnvoyDetailsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a89e3a95b2dd1a92, []int{1}
}
func (m *ListEnvoyDetailsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEnvoyDetailsRequest.Unmarshal(m, b)
}
func (m *ListEnvoyDetailsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEnvoyDetailsRequest.Marshal(b, m, deterministic)
}
func (m *ListEnvoyDetailsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEnvoyDetailsRequest.Merge(m, src)
}
func (m *ListEnvoyDetailsRequest) XXX_Size() int {
	return xxx_messageInfo_ListEnvoyDetailsRequest.Size(m)
}
func (m *ListEnvoyDetailsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEnvoyDetailsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListEnvoyDetailsRequest proto.InternalMessageInfo

type ListEnvoyDetailsResponse struct {
	EnvoyDetails         []*EnvoyDetails `protobuf:"bytes,1,rep,name=envoy_details,json=envoyDetails,proto3" json:"envoy_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListEnvoyDetailsResponse) Reset()         { *m = ListEnvoyDetailsResponse{} }
func (m *ListEnvoyDetailsResponse) String() string { return proto.CompactTextString(m) }
func (*ListEnvoyDetailsResponse) ProtoMessage()    {}
func (*ListEnvoyDetailsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a89e3a95b2dd1a92, []int{2}
}
func (m *ListEnvoyDetailsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEnvoyDetailsResponse.Unmarshal(m, b)
}
func (m *ListEnvoyDetailsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEnvoyDetailsResponse.Marshal(b, m, deterministic)
}
func (m *ListEnvoyDetailsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEnvoyDetailsResponse.Merge(m, src)
}
func (m *ListEnvoyDetailsResponse) XXX_Size() int {
	return xxx_messageInfo_ListEnvoyDetailsResponse.Size(m)
}
func (m *ListEnvoyDetailsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEnvoyDetailsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListEnvoyDetailsResponse proto.InternalMessageInfo

func (m *ListEnvoyDetailsResponse) GetEnvoyDetails() []*EnvoyDetails {
	if m != nil {
		return m.EnvoyDetails
	}
	return nil
}

func init() {
	proto.RegisterType((*EnvoyDetails)(nil), "glooeeapi.solo.io.EnvoyDetails")
	proto.RegisterType((*ListEnvoyDetailsRequest)(nil), "glooeeapi.solo.io.ListEnvoyDetailsRequest")
	proto.RegisterType((*ListEnvoyDetailsResponse)(nil), "glooeeapi.solo.io.ListEnvoyDetailsResponse")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-projects/projects/grpcserver/api/v1/envoy.proto", fileDescriptor_a89e3a95b2dd1a92)
}

var fileDescriptor_a89e3a95b2dd1a92 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xff, 0xfd, 0x4f, 0x86, 0x66, 0x13, 0x34, 0x88, 0x76, 0x3b, 0xe8, 0xe8, 0xa9, 0x28,
	0x26, 0x6c, 0x7a, 0xf3, 0x34, 0x99, 0x78, 0xf1, 0x54, 0x6f, 0x5e, 0x34, 0xab, 0x2f, 0x35, 0xd2,
	0xf6, 0x8d, 0x49, 0xda, 0xd1, 0x93, 0x5f, 0xc7, 0xcf, 0xe5, 0x27, 0x91, 0xa6, 0x22, 0xd3, 0x55,
	0x18, 0x78, 0xea, 0x5b, 0x9e, 0xdf, 0xf3, 0x3e, 0x0f, 0x49, 0xc8, 0x75, 0x22, 0xed, 0x53, 0x31,
	0x67, 0x31, 0x66, 0xdc, 0x60, 0x8a, 0xa7, 0x12, 0x9b, 0xaf, 0xd2, 0xf8, 0x0c, 0xb1, 0x35, 0xfc,
	0x6b, 0x48, 0xb4, 0x8a, 0x0d, 0xe8, 0x12, 0x34, 0x17, 0x4a, 0xf2, 0x72, 0xcc, 0x21, 0x2f, 0xb1,
	0x62, 0x4a, 0xa3, 0x45, 0xba, 0x9b, 0xa4, 0x88, 0x00, 0x42, 0x49, 0x56, 0xfb, 0x99, 0xc4, 0xe1,
	0x5e, 0x82, 0x09, 0x3a, 0x95, 0xd7, 0x53, 0x03, 0x0e, 0xcf, 0xd7, 0x5e, 0x6f, 0x2b, 0x05, 0xa6,
	0x71, 0x05, 0xaf, 0xa4, 0x7f, 0x55, 0xa7, 0xcd, 0xc0, 0x0a, 0x99, 0x1a, 0x4a, 0xc9, 0x46, 0x2e,
	0x32, 0xf0, 0xbd, 0x91, 0x17, 0x6e, 0x45, 0x6e, 0xa6, 0x21, 0xe9, 0x68, 0xb1, 0xf0, 0xff, 0x8f,
	0xbc, 0xb0, 0x37, 0xd9, 0x67, 0x2b, 0x85, 0x58, 0x24, 0x16, 0x51, 0x8d, 0xd0, 0x31, 0xe9, 0x1a,
	0x2b, 0x6c, 0x61, 0xfc, 0x8e, 0x83, 0x07, 0x2d, 0xf0, 0xad, 0x03, 0xa2, 0x4f, 0x30, 0x18, 0x90,
	0x83, 0x1b, 0x69, 0xec, 0x72, 0x89, 0x08, 0x5e, 0x0a, 0x30, 0x36, 0x78, 0x20, 0xfe, 0xaa, 0x64,
	0x14, 0xe6, 0x06, 0xe8, 0x8c, 0x6c, 0xbb, 0x53, 0xba, 0x7f, 0x6c, 0x04, 0xdf, 0x1b, 0x75, 0xc2,
	0xde, 0xe4, 0xa8, 0x25, 0xf0, 0x9b, 0xbf, 0x0f, 0x4b, 0x7f, 0x93, 0x8a, 0x6c, 0x3a, 0x75, 0xaa,
	0x24, 0xcd, 0xc8, 0xce, 0xcf, 0x34, 0x7a, 0xdc, 0xb2, 0xee, 0x97, 0xb6, 0xc3, 0x93, 0xb5, 0xd8,
	0xa6, 0x7e, 0xf0, 0xef, 0x72, 0xfa, 0xf6, 0x7e, 0xe8, 0xdd, 0x5d, 0xfc, 0xe1, 0x99, 0xcc, 0xbb,
	0xee, 0x0a, 0xcf, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x59, 0x4a, 0x34, 0xf8, 0x6c, 0x02, 0x00,
	0x00,
}

func (this *EnvoyDetails) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EnvoyDetails)
	if !ok {
		that2, ok := that.(EnvoyDetails)
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
	if this.Name != that1.Name {
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
func (this *ListEnvoyDetailsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListEnvoyDetailsRequest)
	if !ok {
		that2, ok := that.(ListEnvoyDetailsRequest)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ListEnvoyDetailsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListEnvoyDetailsResponse)
	if !ok {
		that2, ok := that.(ListEnvoyDetailsResponse)
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
	if len(this.EnvoyDetails) != len(that1.EnvoyDetails) {
		return false
	}
	for i := range this.EnvoyDetails {
		if !this.EnvoyDetails[i].Equal(that1.EnvoyDetails[i]) {
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

// EnvoyApiClient is the client API for EnvoyApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EnvoyApiClient interface {
	ListEnvoyDetails(ctx context.Context, in *ListEnvoyDetailsRequest, opts ...grpc.CallOption) (*ListEnvoyDetailsResponse, error)
}

type envoyApiClient struct {
	cc *grpc.ClientConn
}

func NewEnvoyApiClient(cc *grpc.ClientConn) EnvoyApiClient {
	return &envoyApiClient{cc}
}

func (c *envoyApiClient) ListEnvoyDetails(ctx context.Context, in *ListEnvoyDetailsRequest, opts ...grpc.CallOption) (*ListEnvoyDetailsResponse, error) {
	out := new(ListEnvoyDetailsResponse)
	err := c.cc.Invoke(ctx, "/glooeeapi.solo.io.EnvoyApi/ListEnvoyDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvoyApiServer is the server API for EnvoyApi service.
type EnvoyApiServer interface {
	ListEnvoyDetails(context.Context, *ListEnvoyDetailsRequest) (*ListEnvoyDetailsResponse, error)
}

// UnimplementedEnvoyApiServer can be embedded to have forward compatible implementations.
type UnimplementedEnvoyApiServer struct {
}

func (*UnimplementedEnvoyApiServer) ListEnvoyDetails(ctx context.Context, req *ListEnvoyDetailsRequest) (*ListEnvoyDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnvoyDetails not implemented")
}

func RegisterEnvoyApiServer(s *grpc.Server, srv EnvoyApiServer) {
	s.RegisterService(&_EnvoyApi_serviceDesc, srv)
}

func _EnvoyApi_ListEnvoyDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnvoyDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvoyApiServer).ListEnvoyDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glooeeapi.solo.io.EnvoyApi/ListEnvoyDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvoyApiServer).ListEnvoyDetails(ctx, req.(*ListEnvoyDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EnvoyApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "glooeeapi.solo.io.EnvoyApi",
	HandlerType: (*EnvoyApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEnvoyDetails",
			Handler:    _EnvoyApi_ListEnvoyDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/solo-io/solo-projects/projects/grpcserver/api/v1/envoy.proto",
}
