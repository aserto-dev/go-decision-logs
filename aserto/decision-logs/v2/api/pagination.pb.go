// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/decision-logs/v2/api/pagination.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PaginationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Size          int32                  `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	mi := &file_aserto_decision_logs_v2_api_pagination_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_decision_logs_v2_api_pagination_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_aserto_decision_logs_v2_api_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *PaginationRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PaginationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type PaginationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NextToken     string                 `protobuf:"bytes,1,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
	ResultSize    int32                  `protobuf:"varint,2,opt,name=result_size,json=resultSize,proto3" json:"result_size,omitempty"`
	TotalSize     int32                  `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationResponse) Reset() {
	*x = PaginationResponse{}
	mi := &file_aserto_decision_logs_v2_api_pagination_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResponse) ProtoMessage() {}

func (x *PaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_decision_logs_v2_api_pagination_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationResponse.ProtoReflect.Descriptor instead.
func (*PaginationResponse) Descriptor() ([]byte, []int) {
	return file_aserto_decision_logs_v2_api_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *PaginationResponse) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

func (x *PaginationResponse) GetResultSize() int32 {
	if x != nil {
		return x.ResultSize
	}
	return 0
}

func (x *PaginationResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

var File_aserto_decision_logs_v2_api_pagination_proto protoreflect.FileDescriptor

const file_aserto_decision_logs_v2_api_pagination_proto_rawDesc = "" +
	"\n" +
	",aserto/decision-logs/v2/api/pagination.proto\x12\x1aaserto.decisionlogs.api.v2\"=\n" +
	"\x11PaginationRequest\x12\x12\n" +
	"\x04size\x18\x01 \x01(\x05R\x04size\x12\x14\n" +
	"\x05token\x18\x02 \x01(\tR\x05token\"s\n" +
	"\x12PaginationResponse\x12\x1d\n" +
	"\n" +
	"next_token\x18\x01 \x01(\tR\tnextToken\x12\x1f\n" +
	"\vresult_size\x18\x02 \x01(\x05R\n" +
	"resultSize\x12\x1d\n" +
	"\n" +
	"total_size\x18\x03 \x01(\x05R\ttotalSizeBeZFgithub.com/aserto-dev/go-decision-logs/aserto/decision-logs/v2/api;api\xaa\x02\x1aAserto.DecisionLogs.V2.APIb\x06proto3"

var (
	file_aserto_decision_logs_v2_api_pagination_proto_rawDescOnce sync.Once
	file_aserto_decision_logs_v2_api_pagination_proto_rawDescData []byte
)

func file_aserto_decision_logs_v2_api_pagination_proto_rawDescGZIP() []byte {
	file_aserto_decision_logs_v2_api_pagination_proto_rawDescOnce.Do(func() {
		file_aserto_decision_logs_v2_api_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_decision_logs_v2_api_pagination_proto_rawDesc), len(file_aserto_decision_logs_v2_api_pagination_proto_rawDesc)))
	})
	return file_aserto_decision_logs_v2_api_pagination_proto_rawDescData
}

var file_aserto_decision_logs_v2_api_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aserto_decision_logs_v2_api_pagination_proto_goTypes = []any{
	(*PaginationRequest)(nil),  // 0: aserto.decisionlogs.api.v2.PaginationRequest
	(*PaginationResponse)(nil), // 1: aserto.decisionlogs.api.v2.PaginationResponse
}
var file_aserto_decision_logs_v2_api_pagination_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aserto_decision_logs_v2_api_pagination_proto_init() }
func file_aserto_decision_logs_v2_api_pagination_proto_init() {
	if File_aserto_decision_logs_v2_api_pagination_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_decision_logs_v2_api_pagination_proto_rawDesc), len(file_aserto_decision_logs_v2_api_pagination_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_decision_logs_v2_api_pagination_proto_goTypes,
		DependencyIndexes: file_aserto_decision_logs_v2_api_pagination_proto_depIdxs,
		MessageInfos:      file_aserto_decision_logs_v2_api_pagination_proto_msgTypes,
	}.Build()
	File_aserto_decision_logs_v2_api_pagination_proto = out.File
	file_aserto_decision_logs_v2_api_pagination_proto_goTypes = nil
	file_aserto_decision_logs_v2_api_pagination_proto_depIdxs = nil
}
