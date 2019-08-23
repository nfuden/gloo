// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/grpcserver/api/v1/types.proto

package v1

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Status_Code int32

const (
	Status_ERROR   Status_Code = 0
	Status_WARNING Status_Code = 1
	Status_OK      Status_Code = 2
)

var Status_Code_name = map[int32]string{
	0: "ERROR",
	1: "WARNING",
	2: "OK",
}

var Status_Code_value = map[string]int32{
	"ERROR":   0,
	"WARNING": 1,
	"OK":      2,
}

func (x Status_Code) String() string {
	return proto.EnumName(Status_Code_name, int32(x))
}

func (Status_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36bb7454e3c797e7, []int{1, 0}
}

// "Raw" representation of a resource, e.g. yaml representation of a kube custom resource
type Raw struct {
	// e.g. resource-name.yaml
	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	// Content of the file
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	// Error encountered while rendering content
	ContentRenderError   string   `protobuf:"bytes,3,opt,name=content_render_error,json=contentRenderError,proto3" json:"content_render_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Raw) Reset()         { *m = Raw{} }
func (m *Raw) String() string { return proto.CompactTextString(m) }
func (*Raw) ProtoMessage()    {}
func (*Raw) Descriptor() ([]byte, []int) {
	return fileDescriptor_36bb7454e3c797e7, []int{0}
}
func (m *Raw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Raw.Unmarshal(m, b)
}
func (m *Raw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Raw.Marshal(b, m, deterministic)
}
func (m *Raw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Raw.Merge(m, src)
}
func (m *Raw) XXX_Size() int {
	return xxx_messageInfo_Raw.Size(m)
}
func (m *Raw) XXX_DiscardUnknown() {
	xxx_messageInfo_Raw.DiscardUnknown(m)
}

var xxx_messageInfo_Raw proto.InternalMessageInfo

func (m *Raw) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *Raw) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Raw) GetContentRenderError() string {
	if m != nil {
		return m.ContentRenderError
	}
	return ""
}

type Status struct {
	// Red / Yellow / Green status of the resource
	Code Status_Code `protobuf:"varint,1,opt,name=code,proto3,enum=glooeeapi.solo.io.Status_Code" json:"code,omitempty"`
	// Optional error message or other information about resource status
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_36bb7454e3c797e7, []int{1}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() Status_Code {
	if m != nil {
		return m.Code
	}
	return Status_ERROR
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("glooeeapi.solo.io.Status_Code", Status_Code_name, Status_Code_value)
	proto.RegisterType((*Raw)(nil), "glooeeapi.solo.io.Raw")
	proto.RegisterType((*Status)(nil), "glooeeapi.solo.io.Status")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-projects/projects/grpcserver/api/v1/types.proto", fileDescriptor_36bb7454e3c797e7)
}

var fileDescriptor_36bb7454e3c797e7 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0xc1, 0x4a, 0x2b, 0x31,
	0x14, 0x86, 0xef, 0xb4, 0xbd, 0xad, 0x8d, 0x20, 0x35, 0x74, 0x31, 0x28, 0x14, 0xe9, 0x42, 0xdc,
	0x98, 0x68, 0x5d, 0xba, 0xaa, 0x52, 0x8a, 0x08, 0x2d, 0xc4, 0x85, 0xe0, 0xa6, 0xa4, 0x33, 0xc7,
	0x18, 0x99, 0x99, 0x13, 0x92, 0xb4, 0xe2, 0xc6, 0xe7, 0xf1, 0xb9, 0x7c, 0x12, 0x49, 0x66, 0xea,
	0xc6, 0xa5, 0xab, 0xfc, 0x27, 0xdf, 0x17, 0xf2, 0x73, 0xc8, 0x5c, 0x69, 0xff, 0xb2, 0x59, 0xb3,
	0x0c, 0x4b, 0xee, 0xb0, 0xc0, 0x73, 0x8d, 0xf5, 0x69, 0x2c, 0xbe, 0x42, 0xe6, 0x1d, 0xff, 0x09,
	0xca, 0x9a, 0xcc, 0x81, 0xdd, 0x82, 0xe5, 0xd2, 0x68, 0xbe, 0xbd, 0xe4, 0xfe, 0xdd, 0x80, 0x63,
	0xc6, 0xa2, 0x47, 0x7a, 0xa8, 0x0a, 0x44, 0x00, 0x69, 0x34, 0x0b, 0xef, 0x99, 0xc6, 0xa3, 0xa1,
	0x42, 0x85, 0x91, 0xf2, 0x90, 0x6a, 0x71, 0x5c, 0x91, 0xb6, 0x90, 0x6f, 0xf4, 0x98, 0xf4, 0x9f,
	0x75, 0x01, 0xab, 0x4a, 0x96, 0x90, 0x26, 0x27, 0xc9, 0x59, 0x5f, 0xec, 0x85, 0x8b, 0x85, 0x2c,
	0x81, 0xa6, 0xa4, 0x97, 0x61, 0xe5, 0xa1, 0xf2, 0x69, 0x2b, 0xa2, 0xdd, 0x48, 0x2f, 0xc8, 0xb0,
	0x89, 0x2b, 0x0b, 0x55, 0x0e, 0x76, 0x05, 0xd6, 0xa2, 0x4d, 0xdb, 0x51, 0xa3, 0x0d, 0x13, 0x11,
	0xcd, 0x02, 0x19, 0x7f, 0x90, 0xee, 0x83, 0x97, 0x7e, 0xe3, 0xe8, 0x84, 0x74, 0x32, 0xcc, 0xeb,
	0xdf, 0x0e, 0x26, 0x23, 0xf6, 0xab, 0x31, 0xab, 0x45, 0x76, 0x8b, 0x39, 0x88, 0xe8, 0x86, 0x26,
	0x25, 0x38, 0x27, 0x15, 0xec, 0x9a, 0x34, 0xe3, 0xf8, 0x94, 0x74, 0x82, 0x47, 0xfb, 0xe4, 0xff,
	0x4c, 0x88, 0xa5, 0x18, 0xfc, 0xa3, 0xfb, 0xa4, 0xf7, 0x38, 0x15, 0x8b, 0xbb, 0xc5, 0x7c, 0x90,
	0xd0, 0x2e, 0x69, 0x2d, 0xef, 0x07, 0xad, 0x9b, 0xe9, 0xe7, 0xd7, 0x28, 0x79, 0xba, 0xfe, 0xc3,
	0x9e, 0xd7, 0xdd, 0xb8, 0xb9, 0xab, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0x03, 0x10, 0xa4,
	0xad, 0x01, 0x00, 0x00,
}

func (this *Raw) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Raw)
	if !ok {
		that2, ok := that.(Raw)
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
	if this.FileName != that1.FileName {
		return false
	}
	if this.Content != that1.Content {
		return false
	}
	if this.ContentRenderError != that1.ContentRenderError {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Status) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Status)
	if !ok {
		that2, ok := that.(Status)
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
	if this.Code != that1.Code {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
