// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: email/rpc/email.proto

// proto 包名

package email

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

type UserEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserEmailRequest) Reset() {
	*x = UserEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_rpc_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEmailRequest) ProtoMessage() {}

func (x *UserEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_rpc_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEmailRequest.ProtoReflect.Descriptor instead.
func (*UserEmailRequest) Descriptor() ([]byte, []int) {
	return file_email_rpc_email_proto_rawDescGZIP(), []int{0}
}

func (x *UserEmailRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UserEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UserEmailResponse) Reset() {
	*x = UserEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_rpc_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEmailResponse) ProtoMessage() {}

func (x *UserEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_rpc_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEmailResponse.ProtoReflect.Descriptor instead.
func (*UserEmailResponse) Descriptor() ([]byte, []int) {
	return file_email_rpc_email_proto_rawDescGZIP(), []int{1}
}

func (x *UserEmailResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_email_rpc_email_proto protoreflect.FileDescriptor

var file_email_rpc_email_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x44,
	0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x25, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x47, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3e, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x17, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_email_rpc_email_proto_rawDescOnce sync.Once
	file_email_rpc_email_proto_rawDescData = file_email_rpc_email_proto_rawDesc
)

func file_email_rpc_email_proto_rawDescGZIP() []byte {
	file_email_rpc_email_proto_rawDescOnce.Do(func() {
		file_email_rpc_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_email_rpc_email_proto_rawDescData)
	})
	return file_email_rpc_email_proto_rawDescData
}

var file_email_rpc_email_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_email_rpc_email_proto_goTypes = []interface{}{
	(*UserEmailRequest)(nil),  // 0: email.UserEmailRequest
	(*UserEmailResponse)(nil), // 1: email.UserEmailResponse
}
var file_email_rpc_email_proto_depIdxs = []int32{
	0, // 0: email.email.SendEmail:input_type -> email.UserEmailRequest
	1, // 1: email.email.SendEmail:output_type -> email.UserEmailResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_email_rpc_email_proto_init() }
func file_email_rpc_email_proto_init() {
	if File_email_rpc_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_email_rpc_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEmailRequest); i {
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
		file_email_rpc_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEmailResponse); i {
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
			RawDescriptor: file_email_rpc_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_email_rpc_email_proto_goTypes,
		DependencyIndexes: file_email_rpc_email_proto_depIdxs,
		MessageInfos:      file_email_rpc_email_proto_msgTypes,
	}.Build()
	File_email_rpc_email_proto = out.File
	file_email_rpc_email_proto_rawDesc = nil
	file_email_rpc_email_proto_goTypes = nil
	file_email_rpc_email_proto_depIdxs = nil
}