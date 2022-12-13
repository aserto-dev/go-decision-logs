// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: aserto/decision-logs/v2/api/identity_context.proto

package api

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

// Identity types, describes the payload type of the identity field inside the IdentityContext message.
type IdentityType int32

const (
	// Unknown, value not set, requests will fail with identity type not set error.
	IdentityType_IDENTITY_TYPE_UNKNOWN IdentityType = 0
	// None, no explicit identity context set, equals anonymous.
	IdentityType_IDENTITY_TYPE_NONE IdentityType = 1
	// Sub(ject), identity field contains an oAUTH subject.
	IdentityType_IDENTITY_TYPE_SUB IdentityType = 2
	// JWT, identity field contains a JWT access token.
	IdentityType_IDENTITY_TYPE_JWT IdentityType = 3
)

// Enum value maps for IdentityType.
var (
	IdentityType_name = map[int32]string{
		0: "IDENTITY_TYPE_UNKNOWN",
		1: "IDENTITY_TYPE_NONE",
		2: "IDENTITY_TYPE_SUB",
		3: "IDENTITY_TYPE_JWT",
	}
	IdentityType_value = map[string]int32{
		"IDENTITY_TYPE_UNKNOWN": 0,
		"IDENTITY_TYPE_NONE":    1,
		"IDENTITY_TYPE_SUB":     2,
		"IDENTITY_TYPE_JWT":     3,
	}
)

func (x IdentityType) Enum() *IdentityType {
	p := new(IdentityType)
	*p = x
	return p
}

func (x IdentityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdentityType) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_decision_logs_v2_api_identity_context_proto_enumTypes[0].Descriptor()
}

func (IdentityType) Type() protoreflect.EnumType {
	return &file_aserto_decision_logs_v2_api_identity_context_proto_enumTypes[0]
}

func (x IdentityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdentityType.Descriptor instead.
func (IdentityType) EnumDescriptor() ([]byte, []int) {
	return file_aserto_decision_logs_v2_api_identity_context_proto_rawDescGZIP(), []int{0}
}

type IdentityContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string       `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Type     IdentityType `protobuf:"varint,2,opt,name=type,proto3,enum=aserto.decisionlogs.api.v2.IdentityType" json:"type,omitempty"`
}

func (x *IdentityContext) Reset() {
	*x = IdentityContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_decision_logs_v2_api_identity_context_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityContext) ProtoMessage() {}

func (x *IdentityContext) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_decision_logs_v2_api_identity_context_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityContext.ProtoReflect.Descriptor instead.
func (*IdentityContext) Descriptor() ([]byte, []int) {
	return file_aserto_decision_logs_v2_api_identity_context_proto_rawDescGZIP(), []int{0}
}

func (x *IdentityContext) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *IdentityContext) GetType() IdentityType {
	if x != nil {
		return x.Type
	}
	return IdentityType_IDENTITY_TYPE_UNKNOWN
}

var File_aserto_decision_logs_v2_api_identity_context_proto protoreflect.FileDescriptor

var file_aserto_decision_logs_v2_api_identity_context_proto_rawDesc = []byte{
	0x0a, 0x32, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2d, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x22, 0x6b, 0x0a, 0x0f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x6c,
	0x6f, 0x67, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x6f, 0x0a,
	0x0c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x15, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x44, 0x45, 0x4e,
	0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01,
	0x12, 0x15, 0x0a, 0x11, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x55, 0x42, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x44, 0x45, 0x4e, 0x54,
	0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x57, 0x54, 0x10, 0x03, 0x42, 0x65,
	0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f,
	0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0xaa, 0x02, 0x1a, 0x41, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x56,
	0x32, 0x2e, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_decision_logs_v2_api_identity_context_proto_rawDescOnce sync.Once
	file_aserto_decision_logs_v2_api_identity_context_proto_rawDescData = file_aserto_decision_logs_v2_api_identity_context_proto_rawDesc
)

func file_aserto_decision_logs_v2_api_identity_context_proto_rawDescGZIP() []byte {
	file_aserto_decision_logs_v2_api_identity_context_proto_rawDescOnce.Do(func() {
		file_aserto_decision_logs_v2_api_identity_context_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_decision_logs_v2_api_identity_context_proto_rawDescData)
	})
	return file_aserto_decision_logs_v2_api_identity_context_proto_rawDescData
}

var file_aserto_decision_logs_v2_api_identity_context_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_decision_logs_v2_api_identity_context_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aserto_decision_logs_v2_api_identity_context_proto_goTypes = []interface{}{
	(IdentityType)(0),       // 0: aserto.decisionlogs.api.v2.IdentityType
	(*IdentityContext)(nil), // 1: aserto.decisionlogs.api.v2.IdentityContext
}
var file_aserto_decision_logs_v2_api_identity_context_proto_depIdxs = []int32{
	0, // 0: aserto.decisionlogs.api.v2.IdentityContext.type:type_name -> aserto.decisionlogs.api.v2.IdentityType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aserto_decision_logs_v2_api_identity_context_proto_init() }
func file_aserto_decision_logs_v2_api_identity_context_proto_init() {
	if File_aserto_decision_logs_v2_api_identity_context_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_decision_logs_v2_api_identity_context_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityContext); i {
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
			RawDescriptor: file_aserto_decision_logs_v2_api_identity_context_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_decision_logs_v2_api_identity_context_proto_goTypes,
		DependencyIndexes: file_aserto_decision_logs_v2_api_identity_context_proto_depIdxs,
		EnumInfos:         file_aserto_decision_logs_v2_api_identity_context_proto_enumTypes,
		MessageInfos:      file_aserto_decision_logs_v2_api_identity_context_proto_msgTypes,
	}.Build()
	File_aserto_decision_logs_v2_api_identity_context_proto = out.File
	file_aserto_decision_logs_v2_api_identity_context_proto_rawDesc = nil
	file_aserto_decision_logs_v2_api_identity_context_proto_goTypes = nil
	file_aserto_decision_logs_v2_api_identity_context_proto_depIdxs = nil
}