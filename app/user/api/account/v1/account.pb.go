// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: account/v1/account.proto

package v1

import (
	pbentity "proxima/app/user/api/pbentity"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserRegisterReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty" v:"required|min-length:2"` // v:required|min-length:2
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" v:"required|min-length:6"` // v:required|min-length:6
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty" v:"required|email"`              // v:required|email
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRegisterReq) Reset() {
	*x = UserRegisterReq{}
	mi := &file_account_v1_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterReq) ProtoMessage() {}

func (x *UserRegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterReq.ProtoReflect.Descriptor instead.
func (*UserRegisterReq) Descriptor() ([]byte, []int) {
	return file_account_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *UserRegisterReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserRegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserRegisterReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UserRegisterRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRegisterRes) Reset() {
	*x = UserRegisterRes{}
	mi := &file_account_v1_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRegisterRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterRes) ProtoMessage() {}

func (x *UserRegisterRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterRes.ProtoReflect.Descriptor instead.
func (*UserRegisterRes) Descriptor() ([]byte, []int) {
	return file_account_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *UserRegisterRes) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserLoginReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty" v:"required|min-length:2"` // v:required|min-length:2
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" v:"required|min-length:6"` // v:required|min-length:6
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserLoginReq) Reset() {
	*x = UserLoginReq{}
	mi := &file_account_v1_account_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginReq) ProtoMessage() {}

func (x *UserLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_account_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginReq.ProtoReflect.Descriptor instead.
func (*UserLoginReq) Descriptor() ([]byte, []int) {
	return file_account_v1_account_proto_rawDescGZIP(), []int{2}
}

func (x *UserLoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLoginRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserLoginRes) Reset() {
	*x = UserLoginRes{}
	mi := &file_account_v1_account_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserLoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginRes) ProtoMessage() {}

func (x *UserLoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_account_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginRes.ProtoReflect.Descriptor instead.
func (*UserLoginRes) Descriptor() ([]byte, []int) {
	return file_account_v1_account_proto_rawDescGZIP(), []int{3}
}

func (x *UserLoginRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserInfoReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" v:"required"` // v:required
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoReq) Reset() {
	*x = UserInfoReq{}
	mi := &file_account_v1_account_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReq) ProtoMessage() {}

func (x *UserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_account_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReq.ProtoReflect.Descriptor instead.
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return file_account_v1_account_proto_rawDescGZIP(), []int{4}
}

func (x *UserInfoReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserInfoRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *pbentity.Users        `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoRes) Reset() {
	*x = UserInfoRes{}
	mi := &file_account_v1_account_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRes) ProtoMessage() {}

func (x *UserInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_account_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRes.ProtoReflect.Descriptor instead.
func (*UserInfoRes) Descriptor() ([]byte, []int) {
	return file_account_v1_account_proto_rawDescGZIP(), []int{5}
}

func (x *UserInfoRes) GetUser() *pbentity.Users {
	if x != nil {
		return x.User
	}
	return nil
}

var File_account_v1_account_proto protoreflect.FileDescriptor

const file_account_v1_account_proto_rawDesc = "" +
	"\n" +
	"\x18account/v1/account.proto\x12\n" +
	"account.v1\x1a\x14pbentity/users.proto\"_\n" +
	"\x0fUserRegisterReq\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\"!\n" +
	"\x0fUserRegisterRes\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"F\n" +
	"\fUserLoginReq\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"$\n" +
	"\fUserLoginRes\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"#\n" +
	"\vUserInfoReq\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"2\n" +
	"\vUserInfoRes\x12#\n" +
	"\x04user\x18\x01 \x01(\v2\x0f.pbentity.UsersR\x04user2\xd8\x01\n" +
	"\aAccount\x12J\n" +
	"\fUserRegister\x12\x1b.account.v1.UserRegisterReq\x1a\x1b.account.v1.UserRegisterRes\"\x00\x12A\n" +
	"\tUserLogin\x12\x18.account.v1.UserLoginReq\x1a\x18.account.v1.UserLoginRes\"\x00\x12>\n" +
	"\bUserInfo\x12\x17.account.v1.UserInfoReq\x1a\x17.account.v1.UserInfoRes\"\x00B!Z\x1fproxima/app/user/api/account/v1b\x06proto3"

var (
	file_account_v1_account_proto_rawDescOnce sync.Once
	file_account_v1_account_proto_rawDescData []byte
)

func file_account_v1_account_proto_rawDescGZIP() []byte {
	file_account_v1_account_proto_rawDescOnce.Do(func() {
		file_account_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_account_v1_account_proto_rawDesc), len(file_account_v1_account_proto_rawDesc)))
	})
	return file_account_v1_account_proto_rawDescData
}

var file_account_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_account_v1_account_proto_goTypes = []any{
	(*UserRegisterReq)(nil), // 0: account.v1.UserRegisterReq
	(*UserRegisterRes)(nil), // 1: account.v1.UserRegisterRes
	(*UserLoginReq)(nil),    // 2: account.v1.UserLoginReq
	(*UserLoginRes)(nil),    // 3: account.v1.UserLoginRes
	(*UserInfoReq)(nil),     // 4: account.v1.UserInfoReq
	(*UserInfoRes)(nil),     // 5: account.v1.UserInfoRes
	(*pbentity.Users)(nil),  // 6: pbentity.Users
}
var file_account_v1_account_proto_depIdxs = []int32{
	6, // 0: account.v1.UserInfoRes.user:type_name -> pbentity.Users
	0, // 1: account.v1.Account.UserRegister:input_type -> account.v1.UserRegisterReq
	2, // 2: account.v1.Account.UserLogin:input_type -> account.v1.UserLoginReq
	4, // 3: account.v1.Account.UserInfo:input_type -> account.v1.UserInfoReq
	1, // 4: account.v1.Account.UserRegister:output_type -> account.v1.UserRegisterRes
	3, // 5: account.v1.Account.UserLogin:output_type -> account.v1.UserLoginRes
	5, // 6: account.v1.Account.UserInfo:output_type -> account.v1.UserInfoRes
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_account_v1_account_proto_init() }
func file_account_v1_account_proto_init() {
	if File_account_v1_account_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_account_v1_account_proto_rawDesc), len(file_account_v1_account_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_v1_account_proto_goTypes,
		DependencyIndexes: file_account_v1_account_proto_depIdxs,
		MessageInfos:      file_account_v1_account_proto_msgTypes,
	}.Build()
	File_account_v1_account_proto = out.File
	file_account_v1_account_proto_goTypes = nil
	file_account_v1_account_proto_depIdxs = nil
}
