// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.12.4
// source: ExecutionService.proto

package dynamic_executors

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

type ExecutionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassName   string `protobuf:"bytes,1,opt,name=className,proto3" json:"className,omitempty"`
	MethodName  string `protobuf:"bytes,2,opt,name=methodName,proto3" json:"methodName,omitempty"`
	JarLocation string `protobuf:"bytes,3,opt,name=jarLocation,proto3" json:"jarLocation,omitempty"`
	Data        string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ExecutionRequest) Reset() {
	*x = ExecutionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExecutionService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionRequest) ProtoMessage() {}

func (x *ExecutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ExecutionService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionRequest.ProtoReflect.Descriptor instead.
func (*ExecutionRequest) Descriptor() ([]byte, []int) {
	return file_ExecutionService_proto_rawDescGZIP(), []int{0}
}

func (x *ExecutionRequest) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *ExecutionRequest) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *ExecutionRequest) GetJarLocation() string {
	if x != nil {
		return x.JarLocation
	}
	return ""
}

func (x *ExecutionRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type ExecutionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode string `protobuf:"bytes,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	Data    string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ExecutionResponse) Reset() {
	*x = ExecutionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExecutionService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionResponse) ProtoMessage() {}

func (x *ExecutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ExecutionService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionResponse.ProtoReflect.Descriptor instead.
func (*ExecutionResponse) Descriptor() ([]byte, []int) {
	return file_ExecutionService_proto_rawDescGZIP(), []int{1}
}

func (x *ExecutionResponse) GetErrCode() string {
	if x != nil {
		return x.ErrCode
	}
	return ""
}

func (x *ExecutionResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_ExecutionService_proto protoreflect.FileDescriptor

var file_ExecutionService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6f, 0x72, 0x67, 0x2e, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x10, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6a, 0x61,
	0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6a, 0x61, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x41, 0x0a, 0x11, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x32, 0x60, 0x0a, 0x10, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x07, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x12, 0x1f, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x50, 0x01, 0x5a, 0x12, 0x2f, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ExecutionService_proto_rawDescOnce sync.Once
	file_ExecutionService_proto_rawDescData = file_ExecutionService_proto_rawDesc
)

func file_ExecutionService_proto_rawDescGZIP() []byte {
	file_ExecutionService_proto_rawDescOnce.Do(func() {
		file_ExecutionService_proto_rawDescData = protoimpl.X.CompressGZIP(file_ExecutionService_proto_rawDescData)
	})
	return file_ExecutionService_proto_rawDescData
}

var file_ExecutionService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ExecutionService_proto_goTypes = []interface{}{
	(*ExecutionRequest)(nil),  // 0: org.Executors.ExecutionRequest
	(*ExecutionResponse)(nil), // 1: org.Executors.ExecutionResponse
}
var file_ExecutionService_proto_depIdxs = []int32{
	0, // 0: org.Executors.ExecutionService.execute:input_type -> org.Executors.ExecutionRequest
	1, // 1: org.Executors.ExecutionService.execute:output_type -> org.Executors.ExecutionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ExecutionService_proto_init() }
func file_ExecutionService_proto_init() {
	if File_ExecutionService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ExecutionService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionRequest); i {
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
		file_ExecutionService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionResponse); i {
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
			RawDescriptor: file_ExecutionService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ExecutionService_proto_goTypes,
		DependencyIndexes: file_ExecutionService_proto_depIdxs,
		MessageInfos:      file_ExecutionService_proto_msgTypes,
	}.Build()
	File_ExecutionService_proto = out.File
	file_ExecutionService_proto_rawDesc = nil
	file_ExecutionService_proto_goTypes = nil
	file_ExecutionService_proto_depIdxs = nil
}