// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.9
// source: common/version.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// VersionInfo is the type that both `telepresence daemon` (the super-user
// daemon) and `telepresence conector` (the normal-user daemon) use
// when reporting their version to the user-facing CLI.
type VersionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ApiVersion is probably unescessary, as it only gets bumped for
	// things that are detectable other ways, but it's here anyway.
	//
	//   - api_version=1 was edgectl's original JSON-based API that was
	//     served on `/var/run/edgectl.socket`.
	//
	//   - api_version=2 was edgectl's gRPC-based (`package edgectl`) API
	//     that was served on `/var/run/edgectl-daemon.socket`.
	//
	//   - api_version=3 is the current Telepresence 2 gRPC-based
	//     (`package telepresence.{sub}`) API:
	//
	//   - `telepresence.connector` is served on `/tmp/telepresence-connector.socket`.
	//
	//   - `telepresence.daemon` is served on `/var/run/telepresence-daemon.socket`.
	//
	//   - `telepresence.manager` is served on TCP `:8081` (by default) on the traffic-manager Pod.
	//
	//   - `telepresence.systema` is served on TCP+TLS `app.getambassador.io:443` (by default).
	//
	//     This is largely just a rename and split of api_version=2,
	//     since the product is called "telepresence" now instead of
	//     "edgectl" and the "connector" and the "daemon" are now two
	//     separate things.
	ApiVersion int32 `protobuf:"varint,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// Version is a "vSEMVER" string of the product version number.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Executable is the path to the executable for the process.
	Executable string `protobuf:"bytes,3,opt,name=executable,proto3" json:"executable,omitempty"`
	// Name of the process (Client, User Daemon, Root Daemon, Traffic Manager)
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *VersionInfo) Reset() {
	*x = VersionInfo{}
	mi := &file_common_version_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionInfo) ProtoMessage() {}

func (x *VersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_version_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionInfo.ProtoReflect.Descriptor instead.
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return file_common_version_proto_rawDescGZIP(), []int{0}
}

func (x *VersionInfo) GetApiVersion() int32 {
	if x != nil {
		return x.ApiVersion
	}
	return 0
}

func (x *VersionInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *VersionInfo) GetExecutable() string {
	if x != nil {
		return x.Executable
	}
	return ""
}

func (x *VersionInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_common_version_proto protoreflect.FileDescriptor

var file_common_version_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x7c, 0x0a, 0x0b, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70,
	0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x69, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_version_proto_rawDescOnce sync.Once
	file_common_version_proto_rawDescData = file_common_version_proto_rawDesc
)

func file_common_version_proto_rawDescGZIP() []byte {
	file_common_version_proto_rawDescOnce.Do(func() {
		file_common_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_version_proto_rawDescData)
	})
	return file_common_version_proto_rawDescData
}

var file_common_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_version_proto_goTypes = []any{
	(*VersionInfo)(nil), // 0: telepresence.common.VersionInfo
}
var file_common_version_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_version_proto_init() }
func file_common_version_proto_init() {
	if File_common_version_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_version_proto_goTypes,
		DependencyIndexes: file_common_version_proto_depIdxs,
		MessageInfos:      file_common_version_proto_msgTypes,
	}.Build()
	File_common_version_proto = out.File
	file_common_version_proto_rawDesc = nil
	file_common_version_proto_goTypes = nil
	file_common_version_proto_depIdxs = nil
}
