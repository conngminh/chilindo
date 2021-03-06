// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: src/pkg/proto/admin.proto

package admin

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

type CheckIsAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *CheckIsAuthRequest) Reset() {
	*x = CheckIsAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pkg_proto_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIsAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIsAuthRequest) ProtoMessage() {}

func (x *CheckIsAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_pkg_proto_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIsAuthRequest.ProtoReflect.Descriptor instead.
func (*CheckIsAuthRequest) Descriptor() ([]byte, []int) {
	return file_src_pkg_proto_admin_proto_rawDescGZIP(), []int{0}
}

func (x *CheckIsAuthRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CheckIsAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAuth bool `protobuf:"varint,1,opt,name=isAuth,proto3" json:"isAuth,omitempty"`
}

func (x *CheckIsAuthResponse) Reset() {
	*x = CheckIsAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pkg_proto_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIsAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIsAuthResponse) ProtoMessage() {}

func (x *CheckIsAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_pkg_proto_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIsAuthResponse.ProtoReflect.Descriptor instead.
func (*CheckIsAuthResponse) Descriptor() ([]byte, []int) {
	return file_src_pkg_proto_admin_proto_rawDescGZIP(), []int{1}
}

func (x *CheckIsAuthResponse) GetIsAuth() bool {
	if x != nil {
		return x.IsAuth
	}
	return false
}

var File_src_pkg_proto_admin_proto protoreflect.FileDescriptor

var file_src_pkg_proto_admin_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x12, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2d, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x32, 0x48, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x73, 0x41, 0x75, 0x74, 0x68, 0x12, 0x13, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x73, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x12, 0x5a, 0x10, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_pkg_proto_admin_proto_rawDescOnce sync.Once
	file_src_pkg_proto_admin_proto_rawDescData = file_src_pkg_proto_admin_proto_rawDesc
)

func file_src_pkg_proto_admin_proto_rawDescGZIP() []byte {
	file_src_pkg_proto_admin_proto_rawDescOnce.Do(func() {
		file_src_pkg_proto_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_pkg_proto_admin_proto_rawDescData)
	})
	return file_src_pkg_proto_admin_proto_rawDescData
}

var file_src_pkg_proto_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_pkg_proto_admin_proto_goTypes = []interface{}{
	(*CheckIsAuthRequest)(nil),  // 0: CheckIsAuthRequest
	(*CheckIsAuthResponse)(nil), // 1: CheckIsAuthResponse
}
var file_src_pkg_proto_admin_proto_depIdxs = []int32{
	0, // 0: AdminService.CheckIsAuth:input_type -> CheckIsAuthRequest
	1, // 1: AdminService.CheckIsAuth:output_type -> CheckIsAuthResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_pkg_proto_admin_proto_init() }
func file_src_pkg_proto_admin_proto_init() {
	if File_src_pkg_proto_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_pkg_proto_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIsAuthRequest); i {
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
		file_src_pkg_proto_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIsAuthResponse); i {
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
			RawDescriptor: file_src_pkg_proto_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_pkg_proto_admin_proto_goTypes,
		DependencyIndexes: file_src_pkg_proto_admin_proto_depIdxs,
		MessageInfos:      file_src_pkg_proto_admin_proto_msgTypes,
	}.Build()
	File_src_pkg_proto_admin_proto = out.File
	file_src_pkg_proto_admin_proto_rawDesc = nil
	file_src_pkg_proto_admin_proto_goTypes = nil
	file_src_pkg_proto_admin_proto_depIdxs = nil
}
