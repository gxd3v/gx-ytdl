// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.9
// source: codes/actions.proto

package protos

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

type ActionsEnum int32

const (
	ActionsEnum_HEALTH_CHECK         ActionsEnum = 0
	ActionsEnum_SESSION_ID           ActionsEnum = 1
	ActionsEnum_DOWNLOAD_VIDEO_AUDIO ActionsEnum = 2
	ActionsEnum_DOWNLOAD_VIDEO       ActionsEnum = 3
	ActionsEnum_DOWNLOAD_AUDIO       ActionsEnum = 4
	ActionsEnum_LIST_FILES           ActionsEnum = 5
	ActionsEnum_SEND_FILE_TO_CLIENT  ActionsEnum = 6
	ActionsEnum_DELETE_FILE          ActionsEnum = 7
	ActionsEnum_DELETE_SESSION       ActionsEnum = 8
)

// Enum value maps for ActionsEnum.
var (
	ActionsEnum_name = map[int32]string{
		0: "HEALTH_CHECK",
		1: "SESSION_ID",
		2: "DOWNLOAD_VIDEO_AUDIO",
		3: "DOWNLOAD_VIDEO",
		4: "DOWNLOAD_AUDIO",
		5: "LIST_FILES",
		6: "SEND_FILE_TO_CLIENT",
		7: "DELETE_FILE",
		8: "DELETE_SESSION",
	}
	ActionsEnum_value = map[string]int32{
		"HEALTH_CHECK":         0,
		"SESSION_ID":           1,
		"DOWNLOAD_VIDEO_AUDIO": 2,
		"DOWNLOAD_VIDEO":       3,
		"DOWNLOAD_AUDIO":       4,
		"LIST_FILES":           5,
		"SEND_FILE_TO_CLIENT":  6,
		"DELETE_FILE":          7,
		"DELETE_SESSION":       8,
	}
)

func (x ActionsEnum) Enum() *ActionsEnum {
	p := new(ActionsEnum)
	*p = x
	return p
}

func (x ActionsEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionsEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_codes_actions_proto_enumTypes[0].Descriptor()
}

func (ActionsEnum) Type() protoreflect.EnumType {
	return &file_codes_actions_proto_enumTypes[0]
}

func (x ActionsEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionsEnum.Descriptor instead.
func (ActionsEnum) EnumDescriptor() ([]byte, []int) {
	return file_codes_actions_proto_rawDescGZIP(), []int{0}
}

var file_codes_actions_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51232,
		Name:          "protos.action_code",
		Tag:           "bytes,51232,opt,name=action_code",
		Filename:      "codes/actions.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string action_code = 51232;
	E_ActionCode = &file_codes_actions_proto_extTypes[0]
)

var File_codes_actions_proto protoreflect.FileDescriptor

var file_codes_actions_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0xa2, 0x02, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12,
	0x1b, 0x0a, 0x0c, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x10,
	0x00, 0x1a, 0x09, 0x82, 0x82, 0x19, 0x05, 0x41, 0x30, 0x30, 0x30, 0x30, 0x12, 0x19, 0x0a, 0x0a,
	0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x1a, 0x09, 0x82, 0x82,
	0x19, 0x05, 0x41, 0x30, 0x30, 0x30, 0x31, 0x12, 0x23, 0x0a, 0x14, 0x44, 0x4f, 0x57, 0x4e, 0x4c,
	0x4f, 0x41, 0x44, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x10,
	0x02, 0x1a, 0x09, 0x82, 0x82, 0x19, 0x05, 0x41, 0x31, 0x30, 0x30, 0x31, 0x12, 0x1d, 0x0a, 0x0e,
	0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x03,
	0x1a, 0x09, 0x82, 0x82, 0x19, 0x05, 0x41, 0x31, 0x30, 0x30, 0x32, 0x12, 0x1d, 0x0a, 0x0e, 0x44,
	0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x10, 0x04, 0x1a,
	0x09, 0x82, 0x82, 0x19, 0x05, 0x41, 0x31, 0x30, 0x30, 0x33, 0x12, 0x19, 0x0a, 0x0a, 0x4c, 0x49,
	0x53, 0x54, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x53, 0x10, 0x05, 0x1a, 0x09, 0x82, 0x82, 0x19, 0x05,
	0x41, 0x32, 0x30, 0x30, 0x30, 0x12, 0x22, 0x0a, 0x13, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x46, 0x49,
	0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x06, 0x1a, 0x09,
	0x82, 0x82, 0x19, 0x05, 0x41, 0x32, 0x30, 0x30, 0x31, 0x12, 0x1a, 0x0a, 0x0b, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x07, 0x1a, 0x09, 0x82, 0x82, 0x19, 0x05,
	0x41, 0x33, 0x30, 0x30, 0x30, 0x12, 0x1d, 0x0a, 0x0e, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f,
	0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x1a, 0x09, 0x82, 0x82, 0x19, 0x05, 0x41,
	0x33, 0x30, 0x30, 0x31, 0x3a, 0x47, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa0, 0x90, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_codes_actions_proto_rawDescOnce sync.Once
	file_codes_actions_proto_rawDescData = file_codes_actions_proto_rawDesc
)

func file_codes_actions_proto_rawDescGZIP() []byte {
	file_codes_actions_proto_rawDescOnce.Do(func() {
		file_codes_actions_proto_rawDescData = protoimpl.X.CompressGZIP(file_codes_actions_proto_rawDescData)
	})
	return file_codes_actions_proto_rawDescData
}

var file_codes_actions_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_codes_actions_proto_goTypes = []interface{}{
	(ActionsEnum)(0),                      // 0: protos.ActionsEnum
	(*descriptorpb.EnumValueOptions)(nil), // 1: google.protobuf.EnumValueOptions
}
var file_codes_actions_proto_depIdxs = []int32{
	1, // 0: protos.action_code:extendee -> google.protobuf.EnumValueOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_codes_actions_proto_init() }
func file_codes_actions_proto_init() {
	if File_codes_actions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_codes_actions_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_codes_actions_proto_goTypes,
		DependencyIndexes: file_codes_actions_proto_depIdxs,
		EnumInfos:         file_codes_actions_proto_enumTypes,
		ExtensionInfos:    file_codes_actions_proto_extTypes,
	}.Build()
	File_codes_actions_proto = out.File
	file_codes_actions_proto_rawDesc = nil
	file_codes_actions_proto_goTypes = nil
	file_codes_actions_proto_depIdxs = nil
}
