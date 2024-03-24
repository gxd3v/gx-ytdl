// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.9
// source: global.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   SuccessEnum `protobuf:"varint,1,opt,name=code,proto3,enum=protos.SuccessEnum" json:"code,omitempty"`
	Status string      `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_global_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_global_proto_rawDescGZIP(), []int{0}
}

func (x *Success) GetCode() SuccessEnum {
	if x != nil {
		return x.Code
	}
	return SuccessEnum_SESSION_ID
}

func (x *Success) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ErrorsEnum `protobuf:"varint,1,opt,name=code,proto3,enum=protos.ErrorsEnum" json:"code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_global_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_global_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetCode() ErrorsEnum {
	if x != nil {
		return x.Code
	}
	return ErrorsEnum_MALFORMED_MESSAGE
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Created         *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	TimesDownloaded int64                  `protobuf:"varint,3,opt,name=times_downloaded,json=timesDownloaded,proto3" json:"times_downloaded,omitempty"`
	Ttl             *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Size            int64                  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_global_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_global_proto_rawDescGZIP(), []int{2}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *File) GetTimesDownloaded() int64 {
	if x != nil {
		return x.TimesDownloaded
	}
	return 0
}

func (x *File) GetTtl() *timestamppb.Timestamp {
	if x != nil {
		return x.Ttl
	}
	return nil
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ActionCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=Code,json=code,proto3" json:"Code,omitempty"`
}

func (x *ActionCode) Reset() {
	*x = ActionCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionCode) ProtoMessage() {}

func (x *ActionCode) ProtoReflect() protoreflect.Message {
	mi := &file_global_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionCode.ProtoReflect.Descriptor instead.
func (*ActionCode) Descriptor() ([]byte, []int) {
	return file_global_proto_rawDescGZIP(), []int{3}
}

func (x *ActionCode) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type PanicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ErrorsEnum `protobuf:"varint,1,opt,name=code,proto3,enum=protos.ErrorsEnum" json:"code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PanicResponse) Reset() {
	*x = PanicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PanicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PanicResponse) ProtoMessage() {}

func (x *PanicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_global_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PanicResponse.ProtoReflect.Descriptor instead.
func (*PanicResponse) Descriptor() ([]byte, []int) {
	return file_global_proto_rawDescGZIP(), []int{4}
}

func (x *PanicResponse) GetCode() ErrorsEnum {
	if x != nil {
		return x.Code
	}
	return ErrorsEnum_MALFORMED_MESSAGE
}

func (x *PanicResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_global_proto protoreflect.FileDescriptor

var file_global_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x12, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x2f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4a, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x49, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03,
	0x74, 0x74, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x20, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x51, 0x0a, 0x0d, 0x50, 0x61, 0x6e,
	0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0b, 0x5a, 0x09,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_global_proto_rawDescOnce sync.Once
	file_global_proto_rawDescData = file_global_proto_rawDesc
)

func file_global_proto_rawDescGZIP() []byte {
	file_global_proto_rawDescOnce.Do(func() {
		file_global_proto_rawDescData = protoimpl.X.CompressGZIP(file_global_proto_rawDescData)
	})
	return file_global_proto_rawDescData
}

var file_global_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_global_proto_goTypes = []interface{}{
	(*Success)(nil),               // 0: protos.Success
	(*Error)(nil),                 // 1: protos.Error
	(*File)(nil),                  // 2: protos.File
	(*ActionCode)(nil),            // 3: protos.ActionCode
	(*PanicResponse)(nil),         // 4: protos.PanicResponse
	(SuccessEnum)(0),              // 5: protos.SuccessEnum
	(ErrorsEnum)(0),               // 6: protos.ErrorsEnum
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_global_proto_depIdxs = []int32{
	5, // 0: protos.Success.code:type_name -> protos.SuccessEnum
	6, // 1: protos.Error.code:type_name -> protos.ErrorsEnum
	7, // 2: protos.File.created:type_name -> google.protobuf.Timestamp
	7, // 3: protos.File.ttl:type_name -> google.protobuf.Timestamp
	6, // 4: protos.PanicResponse.code:type_name -> protos.ErrorsEnum
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_global_proto_init() }
func file_global_proto_init() {
	if File_global_proto != nil {
		return
	}
	file_codes_errors_proto_init()
	file_codes_success_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_global_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
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
		file_global_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_global_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_global_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionCode); i {
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
		file_global_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PanicResponse); i {
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
			RawDescriptor: file_global_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_global_proto_goTypes,
		DependencyIndexes: file_global_proto_depIdxs,
		MessageInfos:      file_global_proto_msgTypes,
	}.Build()
	File_global_proto = out.File
	file_global_proto_rawDesc = nil
	file_global_proto_goTypes = nil
	file_global_proto_depIdxs = nil
}
