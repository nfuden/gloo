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

// "Raw" representation of a resource, e.g. yaml representation of a kube custom resource
type Raw struct {
	// e.g. resource-name.yaml
	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	// Content of the file
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
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

func init() {
	proto.RegisterType((*Raw)(nil), "glooeeapi.solo.io.Raw")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-projects/projects/grpcserver/api/v1/types.proto", fileDescriptor_36bb7454e3c797e7)
}

var fileDescriptor_36bb7454e3c797e7 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0x87, 0xd0,
	0x05, 0x45, 0xf9, 0x59, 0xa9, 0xc9, 0x25, 0xc5, 0xfa, 0x70, 0x46, 0x7a, 0x51, 0x41, 0x72, 0x71,
	0x6a, 0x51, 0x59, 0x6a, 0x91, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x99, 0xa1, 0x7e, 0x49, 0x65, 0x41,
	0x6a, 0xb1, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x60, 0x7a, 0x4e, 0x7e, 0x7e, 0x6a, 0x6a,
	0x62, 0x41, 0xa6, 0x1e, 0x48, 0xbf, 0x5e, 0x66, 0xbe, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58,
	0x56, 0x1f, 0xc4, 0x82, 0x28, 0x54, 0xb2, 0xe1, 0x62, 0x0e, 0x4a, 0x2c, 0x17, 0x92, 0xe6, 0xe2,
	0x4c, 0xcb, 0xcc, 0x49, 0x8d, 0xcf, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0xe2, 0x00, 0x09, 0xf8, 0x25, 0xe6, 0xa6, 0x0a, 0x49, 0x70, 0xb1, 0x27, 0xe7, 0xe7, 0x95, 0xa4,
	0xe6, 0x95, 0x48, 0x30, 0x81, 0xa5, 0x60, 0x5c, 0x27, 0xc7, 0x15, 0x8f, 0xe4, 0x18, 0xa3, 0xac,
	0x29, 0x70, 0x75, 0x12, 0x1b, 0xd8, 0x1d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0xa3,
	0x4d, 0x1b, 0xfb, 0x00, 0x00, 0x00,
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
