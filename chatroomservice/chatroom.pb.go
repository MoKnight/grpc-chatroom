// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: chatroom.proto

// 包名: chatroomservice

package chatroomservice

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

// 请求数据 Request格式定义
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TS             *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=TS,proto3" json:"TS,omitempty"`
	Type           int32                  `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`                          //1 for login,2 for talk, 3 for end
	Input          string                 `protobuf:"bytes,3,opt,name=Input,proto3" json:"Input,omitempty"`                         //发言信息
	Loginformation *Request_Loginfo       `protobuf:"bytes,4,opt,name=Loginformation,proto3,oneof" json:"Loginformation,omitempty"` //登录信息
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatroom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_chatroom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_chatroom_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetTS() *timestamppb.Timestamp {
	if x != nil {
		return x.TS
	}
	return nil
}

func (x *Request) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Request) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *Request) GetLoginformation() *Request_Loginfo {
	if x != nil {
		return x.Loginformation
	}
	return nil
}

// 响应数据Response格式定义
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TS     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=TS,proto3" json:"TS,omitempty"`
	Name   string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Output string                 `protobuf:"bytes,3,opt,name=Output,proto3" json:"Output,omitempty"` //发言信息
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatroom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_chatroom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_chatroom_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetTS() *timestamppb.Timestamp {
	if x != nil {
		return x.TS
	}
	return nil
}

func (x *Response) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Response) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type Request_Loginfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Request_Loginfo) Reset() {
	*x = Request_Loginfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatroom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Loginfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Loginfo) ProtoMessage() {}

func (x *Request_Loginfo) ProtoReflect() protoreflect.Message {
	mi := &file_chatroom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Loginfo.ProtoReflect.Descriptor instead.
func (*Request_Loginfo) Descriptor() ([]byte, []int) {
	return file_chatroom_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Request_Loginfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Request_Loginfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_chatroom_proto protoreflect.FileDescriptor

var file_chatroom_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x02, 0x54, 0x53, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x54, 0x53, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x4d, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x48, 0x00,
	0x52, 0x0e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x1a, 0x39, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x11,
	0x0a, 0x0f, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x62, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a,
	0x02, 0x54, 0x53, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x54, 0x53, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0x50, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x54, 0x61, 0x6c, 0x6b, 0x12, 0x18, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f,
	0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x63, 0x68, 0x61,
	0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chatroom_proto_rawDescOnce sync.Once
	file_chatroom_proto_rawDescData = file_chatroom_proto_rawDesc
)

func file_chatroom_proto_rawDescGZIP() []byte {
	file_chatroom_proto_rawDescOnce.Do(func() {
		file_chatroom_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatroom_proto_rawDescData)
	})
	return file_chatroom_proto_rawDescData
}

var file_chatroom_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chatroom_proto_goTypes = []interface{}{
	(*Request)(nil),               // 0: chatroomservice.Request
	(*Response)(nil),              // 1: chatroomservice.Response
	(*Request_Loginfo)(nil),       // 2: chatroomservice.Request.Loginfo
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_chatroom_proto_depIdxs = []int32{
	3, // 0: chatroomservice.Request.TS:type_name -> google.protobuf.Timestamp
	2, // 1: chatroomservice.Request.Loginformation:type_name -> chatroomservice.Request.Loginfo
	3, // 2: chatroomservice.Response.TS:type_name -> google.protobuf.Timestamp
	0, // 3: chatroomservice.ChatService.Talk:input_type -> chatroomservice.Request
	1, // 4: chatroomservice.ChatService.Talk:output_type -> chatroomservice.Response
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chatroom_proto_init() }
func file_chatroom_proto_init() {
	if File_chatroom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatroom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_chatroom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_chatroom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Loginfo); i {
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
	file_chatroom_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chatroom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chatroom_proto_goTypes,
		DependencyIndexes: file_chatroom_proto_depIdxs,
		MessageInfos:      file_chatroom_proto_msgTypes,
	}.Build()
	File_chatroom_proto = out.File
	file_chatroom_proto_rawDesc = nil
	file_chatroom_proto_goTypes = nil
	file_chatroom_proto_depIdxs = nil
}
