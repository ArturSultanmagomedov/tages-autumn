// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: internal/api/api.proto

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

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	B    []byte `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_api_proto_msgTypes[0]
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
	return file_internal_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetB() []byte {
	if x != nil {
		return x.B
	}
	return nil
}

type Nothing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Nothing) Reset() {
	*x = Nothing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nothing) ProtoMessage() {}

func (x *Nothing) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nothing.ProtoReflect.Descriptor instead.
func (*Nothing) Descriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{1}
}

type FilesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *FilesList) Reset() {
	*x = FilesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesList) ProtoMessage() {}

func (x *FilesList) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesList.ProtoReflect.Descriptor instead.
func (*FilesList) Descriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *FilesList) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type FileName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FileName) Reset() {
	*x = FileName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileName) ProtoMessage() {}

func (x *FileName) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileName.ProtoReflect.Descriptor instead.
func (*FileName) Descriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *FileName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_internal_api_api_proto protoreflect.FileDescriptor

var file_internal_api_api_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x01, 0x62, 0x22, 0x09, 0x0a, 0x07, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x21, 0x0a,
	0x09, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x22, 0x1e, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x32, 0x7a, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x09, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x05, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x05, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x08, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x08, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x1a, 0x0a,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a,
	0x74, 0x61, 0x67, 0x65, 0x73, 0x2d, 0x61, 0x75, 0x74, 0x75, 0x6d, 0x6e, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internal_api_api_proto_rawDescOnce sync.Once
	file_internal_api_api_proto_rawDescData = file_internal_api_api_proto_rawDesc
)

func file_internal_api_api_proto_rawDescGZIP() []byte {
	file_internal_api_api_proto_rawDescOnce.Do(func() {
		file_internal_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_api_api_proto_rawDescData)
	})
	return file_internal_api_api_proto_rawDescData
}

var file_internal_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_api_api_proto_goTypes = []interface{}{
	(*File)(nil),      // 0: File
	(*Nothing)(nil),   // 1: Nothing
	(*FilesList)(nil), // 2: FilesList
	(*FileName)(nil),  // 3: FileName
}
var file_internal_api_api_proto_depIdxs = []int32{
	3, // 0: FileService.DownloadFile:input_type -> FileName
	0, // 1: FileService.UploadFile:input_type -> File
	1, // 2: FileService.GetFilesList:input_type -> Nothing
	0, // 3: FileService.DownloadFile:output_type -> File
	1, // 4: FileService.UploadFile:output_type -> Nothing
	2, // 5: FileService.GetFilesList:output_type -> FilesList
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_api_api_proto_init() }
func file_internal_api_api_proto_init() {
	if File_internal_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nothing); i {
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
		file_internal_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesList); i {
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
		file_internal_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileName); i {
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
			RawDescriptor: file_internal_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_api_api_proto_goTypes,
		DependencyIndexes: file_internal_api_api_proto_depIdxs,
		MessageInfos:      file_internal_api_api_proto_msgTypes,
	}.Build()
	File_internal_api_api_proto = out.File
	file_internal_api_api_proto_rawDesc = nil
	file_internal_api_api_proto_goTypes = nil
	file_internal_api_api_proto_depIdxs = nil
}