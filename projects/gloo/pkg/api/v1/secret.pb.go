// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: secret.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core_solo_io "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//
// @solo-kit:resource
// @solo-kit:resource.data_type
// @solo-kit:resource.short_name=sec
// @solo-kit:resource.plural_name=secrets
// @solo-kit:resource.group_name=gloo.solo.io
// @solo-kit:resource.version=v1
//
// Certain plugins such as the AWS Lambda Plugin require the use of secrets for authentication, configuration of SSL Certificates, and other data that should not be stored in plaintext configuration.
//
// Gloo runs an independent (goroutine) controller to monitor secrets. Secrets are stored in their own secret storage layer. Gloo can monitor secrets stored in the following secret storage services:
//
// Kubernetes Secrets
// Hashicorp Vault
// Plaintext files (recommended only for testing)
// Secrets must adhere to a structure, specified by the plugin that requires them.
//
// Gloo's secret backend can be configured in Gloo's bootstrap options
type Secret struct {
	// The data this Secret contains
	Data map[string]string `protobuf:"bytes,1,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Metadata contains the object metadata for this resource
	Metadata core_solo_io.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
}

func (m *Secret) Reset()                    { *m = Secret{} }
func (m *Secret) String() string            { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()               {}
func (*Secret) Descriptor() ([]byte, []int) { return fileDescriptorSecret, []int{0} }

func (m *Secret) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Secret) GetMetadata() core_solo_io.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core_solo_io.Metadata{}
}

func init() {
	proto.RegisterType((*Secret)(nil), "gloo.api.v1.Secret")
}
func (this *Secret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret)
	if !ok {
		that2, ok := that.(Secret)
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
	if len(this.Data) != len(that1.Data) {
		return false
	}
	for i := range this.Data {
		if this.Data[i] != that1.Data[i] {
			return false
		}
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("secret.proto", fileDescriptorSecret) }

var fileDescriptorSecret = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x4d, 0x2e,
	0x4a, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0xcf, 0xc9, 0xcf, 0xd7, 0x4b,
	0x2c, 0xc8, 0xd4, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x8b, 0xeb, 0x83, 0x58,
	0x10, 0x25, 0x52, 0x86, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc5,
	0xf9, 0x39, 0xf9, 0xba, 0x99, 0xf9, 0x10, 0x3a, 0x3b, 0xb3, 0x44, 0x3f, 0xb1, 0x20, 0x53, 0xbf,
	0xcc, 0x50, 0x3f, 0x37, 0xb5, 0x24, 0x31, 0x25, 0xb1, 0x24, 0x11, 0xa2, 0x45, 0x69, 0x1d, 0x23,
	0x17, 0x5b, 0x30, 0xd8, 0x1a, 0x21, 0x43, 0x2e, 0x16, 0x90, 0x84, 0x04, 0xa3, 0x02, 0xb3, 0x06,
	0xb7, 0x91, 0xac, 0x1e, 0x92, 0x7d, 0x7a, 0x10, 0x25, 0x7a, 0x2e, 0x89, 0x25, 0x89, 0xae, 0x79,
	0x25, 0x45, 0x95, 0x41, 0x60, 0xa5, 0x42, 0x16, 0x5c, 0x1c, 0x30, 0xf3, 0x24, 0xd8, 0x15, 0x18,
	0x35, 0xb8, 0x8d, 0xc4, 0xf4, 0x92, 0xf3, 0x8b, 0x52, 0xf5, 0x40, 0xb6, 0xea, 0x65, 0xe6, 0xeb,
	0xf9, 0x42, 0x65, 0x9d, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x82, 0xab, 0x96, 0x32, 0xe7, 0xe2,
	0x84, 0x1b, 0x26, 0x24, 0xc0, 0xc5, 0x9c, 0x9d, 0x5a, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19,
	0x04, 0x62, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x81, 0xc5, 0x20,
	0x1c, 0x2b, 0x26, 0x0b, 0x46, 0x27, 0xab, 0x15, 0x8f, 0xe4, 0x18, 0xa3, 0x4c, 0xf0, 0xf9, 0xb4,
	0xa0, 0x28, 0x3f, 0x2b, 0x35, 0xb9, 0xa4, 0x58, 0x1f, 0xe4, 0x03, 0xfd, 0x82, 0xec, 0x74, 0xa8,
	0xdf, 0x93, 0xd8, 0xc0, 0x7e, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x6c, 0x8c, 0x2e,
	0x59, 0x01, 0x00, 0x00,
}
