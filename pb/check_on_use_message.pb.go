// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: check_on_use_message.proto

package pb

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

type CheckFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName      string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	VersionNumber string `protobuf:"bytes,2,opt,name=version_number,json=versionNumber,proto3" json:"version_number,omitempty"`
}

func (x *CheckFileRequest) Reset() {
	*x = CheckFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_on_use_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckFileRequest) ProtoMessage() {}

func (x *CheckFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_check_on_use_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckFileRequest.ProtoReflect.Descriptor instead.
func (*CheckFileRequest) Descriptor() ([]byte, []int) {
	return file_check_on_use_message_proto_rawDescGZIP(), []int{0}
}

func (x *CheckFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *CheckFileRequest) GetVersionNumber() string {
	if x != nil {
		return x.VersionNumber
	}
	return ""
}

type CheckFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsNewest bool `protobuf:"varint,1,opt,name=is_newest,json=isNewest,proto3" json:"is_newest,omitempty"`
}

func (x *CheckFileResponse) Reset() {
	*x = CheckFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_on_use_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckFileResponse) ProtoMessage() {}

func (x *CheckFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_check_on_use_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckFileResponse.ProtoReflect.Descriptor instead.
func (*CheckFileResponse) Descriptor() ([]byte, []int) {
	return file_check_on_use_message_proto_rawDescGZIP(), []int{1}
}

func (x *CheckFileResponse) GetIsNewest() bool {
	if x != nil {
		return x.IsNewest
	}
	return false
}

type PullFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *PullFileRequest) Reset() {
	*x = PullFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_on_use_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullFileRequest) ProtoMessage() {}

func (x *PullFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_check_on_use_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullFileRequest.ProtoReflect.Descriptor instead.
func (*PullFileRequest) Descriptor() ([]byte, []int) {
	return file_check_on_use_message_proto_rawDescGZIP(), []int{2}
}

func (x *PullFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type PullFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName      string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	VersionNumber string `protobuf:"bytes,2,opt,name=version_number,json=versionNumber,proto3" json:"version_number,omitempty"`
	Data          []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PullFileResponse) Reset() {
	*x = PullFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_on_use_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullFileResponse) ProtoMessage() {}

func (x *PullFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_check_on_use_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullFileResponse.ProtoReflect.Descriptor instead.
func (*PullFileResponse) Descriptor() ([]byte, []int) {
	return file_check_on_use_message_proto_rawDescGZIP(), []int{3}
}

func (x *PullFileResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *PullFileResponse) GetVersionNumber() string {
	if x != nil {
		return x.VersionNumber
	}
	return ""
}

func (x *PullFileResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_check_on_use_message_proto protoreflect.FileDescriptor

var file_check_on_use_message_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6f, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x10,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x6e, 0x65, 0x77, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x4e, 0x65, 0x77, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x0f, 0x50, 0x75, 0x6c, 0x6c, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6a, 0x0a, 0x10, 0x50, 0x75, 0x6c, 0x6c, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x32, 0x7c, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x6e, 0x55, 0x73, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0a, 0x69, 0x73, 0x55, 0x70, 0x54, 0x6f,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x08, 0x70, 0x75, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x2e, 0x50, 0x75, 0x6c, 0x6c,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x50, 0x75,
	0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_check_on_use_message_proto_rawDescOnce sync.Once
	file_check_on_use_message_proto_rawDescData = file_check_on_use_message_proto_rawDesc
)

func file_check_on_use_message_proto_rawDescGZIP() []byte {
	file_check_on_use_message_proto_rawDescOnce.Do(func() {
		file_check_on_use_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_check_on_use_message_proto_rawDescData)
	})
	return file_check_on_use_message_proto_rawDescData
}

var file_check_on_use_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_check_on_use_message_proto_goTypes = []interface{}{
	(*CheckFileRequest)(nil),  // 0: CheckFileRequest
	(*CheckFileResponse)(nil), // 1: CheckFileResponse
	(*PullFileRequest)(nil),   // 2: PullFileRequest
	(*PullFileResponse)(nil),  // 3: PullFileResponse
}
var file_check_on_use_message_proto_depIdxs = []int32{
	0, // 0: CheckOnUseServer.isUpToDate:input_type -> CheckFileRequest
	2, // 1: CheckOnUseServer.pullFile:input_type -> PullFileRequest
	1, // 2: CheckOnUseServer.isUpToDate:output_type -> CheckFileResponse
	3, // 3: CheckOnUseServer.pullFile:output_type -> PullFileResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_check_on_use_message_proto_init() }
func file_check_on_use_message_proto_init() {
	if File_check_on_use_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_check_on_use_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckFileRequest); i {
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
		file_check_on_use_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckFileResponse); i {
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
		file_check_on_use_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullFileRequest); i {
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
		file_check_on_use_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullFileResponse); i {
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
			RawDescriptor: file_check_on_use_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_check_on_use_message_proto_goTypes,
		DependencyIndexes: file_check_on_use_message_proto_depIdxs,
		MessageInfos:      file_check_on_use_message_proto_msgTypes,
	}.Build()
	File_check_on_use_message_proto = out.File
	file_check_on_use_message_proto_rawDesc = nil
	file_check_on_use_message_proto_goTypes = nil
	file_check_on_use_message_proto_depIdxs = nil
}
