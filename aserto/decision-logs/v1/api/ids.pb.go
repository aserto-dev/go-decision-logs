// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: aserto/decision-logs/v1/api/ids.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IDType int32

const (
	IDType_ID_TYPE_UNKNOWN        IDType = 0
	IDType_ID_TYPE_ACCOUNT        IDType = 1
	IDType_ID_TYPE_TENANT         IDType = 2
	IDType_ID_TYPE_ERROR          IDType = 3
	IDType_ID_TYPE_POLICY         IDType = 4
	IDType_ID_TYPE_REQUEST        IDType = 5
	IDType_ID_TYPE_PROVIDER       IDType = 6
	IDType_ID_TYPE_CONNECTION     IDType = 7
	IDType_ID_TYPE_INVITE         IDType = 8
	IDType_ID_TYPE_POLICY_BUILDER IDType = 9
	IDType_ID_TYPE_POLICY_REPO    IDType = 10
)

// Enum value maps for IDType.
var (
	IDType_name = map[int32]string{
		0:  "ID_TYPE_UNKNOWN",
		1:  "ID_TYPE_ACCOUNT",
		2:  "ID_TYPE_TENANT",
		3:  "ID_TYPE_ERROR",
		4:  "ID_TYPE_POLICY",
		5:  "ID_TYPE_REQUEST",
		6:  "ID_TYPE_PROVIDER",
		7:  "ID_TYPE_CONNECTION",
		8:  "ID_TYPE_INVITE",
		9:  "ID_TYPE_POLICY_BUILDER",
		10: "ID_TYPE_POLICY_REPO",
	}
	IDType_value = map[string]int32{
		"ID_TYPE_UNKNOWN":        0,
		"ID_TYPE_ACCOUNT":        1,
		"ID_TYPE_TENANT":         2,
		"ID_TYPE_ERROR":          3,
		"ID_TYPE_POLICY":         4,
		"ID_TYPE_REQUEST":        5,
		"ID_TYPE_PROVIDER":       6,
		"ID_TYPE_CONNECTION":     7,
		"ID_TYPE_INVITE":         8,
		"ID_TYPE_POLICY_BUILDER": 9,
		"ID_TYPE_POLICY_REPO":    10,
	}
)

func (x IDType) Enum() *IDType {
	p := new(IDType)
	*p = x
	return p
}

func (x IDType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IDType) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_decision_logs_v1_api_ids_proto_enumTypes[0].Descriptor()
}

func (IDType) Type() protoreflect.EnumType {
	return &file_aserto_decision_logs_v1_api_ids_proto_enumTypes[0]
}

func (x IDType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IDType.Descriptor instead.
func (IDType) EnumDescriptor() ([]byte, []int) {
	return file_aserto_decision_logs_v1_api_ids_proto_rawDescGZIP(), []int{0}
}

var file_aserto_decision_logs_v1_api_ids_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*IDType)(nil),
		Field:         50003,
		Name:          "aserto.decisionlogs.api.v1.id_type",
		Tag:           "varint,50003,opt,name=id_type,enum=aserto.decisionlogs.api.v1.IDType",
		Filename:      "aserto/decision-logs/v1/api/ids.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50003,
		Name:          "aserto.decisionlogs.api.v1.id",
		Tag:           "bytes,50003,opt,name=id",
		Filename:      "aserto/decision-logs/v1/api/ids.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional aserto.decisionlogs.api.v1.IDType id_type = 50003;
	E_IdType = &file_aserto_decision_logs_v1_api_ids_proto_extTypes[0]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string id = 50003;
	E_Id = &file_aserto_decision_logs_v1_api_ids_proto_extTypes[1]
)

var File_aserto_decision_logs_v1_api_ids_proto protoreflect.FileDescriptor

var file_aserto_decision_logs_v1_api_ids_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2d, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xf9, 0x01, 0x0a, 0x06, 0x49, 0x44, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x13, 0x0a, 0x0f, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x11,
	0x0a, 0x0d, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x03, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x4c,
	0x49, 0x43, 0x59, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x10, 0x06,
	0x12, 0x16, 0x0a, 0x12, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x10, 0x08, 0x12, 0x1a, 0x0a, 0x16,
	0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x42,
	0x55, 0x49, 0x4c, 0x44, 0x45, 0x52, 0x10, 0x09, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x10,
	0x0a, 0x3a, 0x5f, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd3, 0x86, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x44, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x69, 0x64, 0x54, 0x79, 0x70, 0x65, 0x88,
	0x01, 0x01, 0x3a, 0x36, 0x0a, 0x02, 0x69, 0x64, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd3, 0x86, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x65, 0x5a, 0x46, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d,
	0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2d,
	0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69,
	0x3b, 0x61, 0x70, 0x69, 0xaa, 0x02, 0x1a, 0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x41, 0x50,
	0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_decision_logs_v1_api_ids_proto_rawDescOnce sync.Once
	file_aserto_decision_logs_v1_api_ids_proto_rawDescData = file_aserto_decision_logs_v1_api_ids_proto_rawDesc
)

func file_aserto_decision_logs_v1_api_ids_proto_rawDescGZIP() []byte {
	file_aserto_decision_logs_v1_api_ids_proto_rawDescOnce.Do(func() {
		file_aserto_decision_logs_v1_api_ids_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_decision_logs_v1_api_ids_proto_rawDescData)
	})
	return file_aserto_decision_logs_v1_api_ids_proto_rawDescData
}

var file_aserto_decision_logs_v1_api_ids_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_decision_logs_v1_api_ids_proto_goTypes = []interface{}{
	(IDType)(0),                           // 0: aserto.decisionlogs.api.v1.IDType
	(*descriptorpb.FieldOptions)(nil),     // 1: google.protobuf.FieldOptions
	(*descriptorpb.EnumValueOptions)(nil), // 2: google.protobuf.EnumValueOptions
}
var file_aserto_decision_logs_v1_api_ids_proto_depIdxs = []int32{
	1, // 0: aserto.decisionlogs.api.v1.id_type:extendee -> google.protobuf.FieldOptions
	2, // 1: aserto.decisionlogs.api.v1.id:extendee -> google.protobuf.EnumValueOptions
	0, // 2: aserto.decisionlogs.api.v1.id_type:type_name -> aserto.decisionlogs.api.v1.IDType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aserto_decision_logs_v1_api_ids_proto_init() }
func file_aserto_decision_logs_v1_api_ids_proto_init() {
	if File_aserto_decision_logs_v1_api_ids_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_decision_logs_v1_api_ids_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_aserto_decision_logs_v1_api_ids_proto_goTypes,
		DependencyIndexes: file_aserto_decision_logs_v1_api_ids_proto_depIdxs,
		EnumInfos:         file_aserto_decision_logs_v1_api_ids_proto_enumTypes,
		ExtensionInfos:    file_aserto_decision_logs_v1_api_ids_proto_extTypes,
	}.Build()
	File_aserto_decision_logs_v1_api_ids_proto = out.File
	file_aserto_decision_logs_v1_api_ids_proto_rawDesc = nil
	file_aserto_decision_logs_v1_api_ids_proto_goTypes = nil
	file_aserto_decision_logs_v1_api_ids_proto_depIdxs = nil
}
