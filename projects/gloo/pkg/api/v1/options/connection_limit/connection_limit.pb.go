// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/connection_limit/connection_limit.proto

package connection_limit

import (
	reflect "reflect"
	sync "sync"

	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// These options provide the ability to limit the active connections in envoy.
// Ref. https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/connection_limit_filter
type ConnectionLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of active connections for this gateway. When this limit is reached, any incoming connection
	// will be closed after delay duration.
	// Must be greater than or equal to one.
	MaxActiveConnections *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_active_connections,json=maxActiveConnections,proto3" json:"max_active_connections,omitempty"`
	// The time to wait before a connection is dropped. Useful for DoS prevention.
	// Defaults to zero and the connection will be closed immediately.
	DelayBeforeClose *duration.Duration `protobuf:"bytes,2,opt,name=delay_before_close,json=delayBeforeClose,proto3" json:"delay_before_close,omitempty"`
}

func (x *ConnectionLimit) Reset() {
	*x = ConnectionLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionLimit) ProtoMessage() {}

func (x *ConnectionLimit) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionLimit.ProtoReflect.Descriptor instead.
func (*ConnectionLimit) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionLimit) GetMaxActiveConnections() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxActiveConnections
	}
	return nil
}

func (x *ConnectionLimit) GetDelayBeforeClose() *duration.Duration {
	if x != nil {
		return x.DelayBeforeClose
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_rawDesc = []byte{
	0x0a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xae, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x52, 0x0a, 0x16, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x47, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x61,
	0x79, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x10, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x42, 0x57, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a,
	0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_goTypes = []interface{}{
	(*ConnectionLimit)(nil),      // 0: connection_limit.options.gloo.solo.io.ConnectionLimit
	(*wrappers.UInt32Value)(nil), // 1: google.protobuf.UInt32Value
	(*duration.Duration)(nil),    // 2: google.protobuf.Duration
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_depIdxs = []int32{
	1, // 0: connection_limit.options.gloo.solo.io.ConnectionLimit.max_active_connections:type_name -> google.protobuf.UInt32Value
	2, // 1: connection_limit.options.gloo.solo.io.ConnectionLimit.delay_before_close:type_name -> google.protobuf.Duration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionLimit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_connection_limit_connection_limit_proto_depIdxs = nil
}
