// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: proto/source/studyplaces/studyplaces.proto

package studyplaces

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

type EnrollmentRole int32

const (
	EnrollmentRole_TEACHER EnrollmentRole = 0
	EnrollmentRole_STUDENT EnrollmentRole = 1
	EnrollmentRole_STAFF   EnrollmentRole = 2
	EnrollmentRole_ADMIN   EnrollmentRole = 3
)

// Enum value maps for EnrollmentRole.
var (
	EnrollmentRole_name = map[int32]string{
		0: "TEACHER",
		1: "STUDENT",
		2: "STAFF",
		3: "ADMIN",
	}
	EnrollmentRole_value = map[string]int32{
		"TEACHER": 0,
		"STUDENT": 1,
		"STAFF":   2,
		"ADMIN":   3,
	}
)

func (x EnrollmentRole) Enum() *EnrollmentRole {
	p := new(EnrollmentRole)
	*p = x
	return p
}

func (x EnrollmentRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnrollmentRole) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_source_studyplaces_studyplaces_proto_enumTypes[0].Descriptor()
}

func (EnrollmentRole) Type() protoreflect.EnumType {
	return &file_proto_source_studyplaces_studyplaces_proto_enumTypes[0]
}

func (x EnrollmentRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnrollmentRole.Descriptor instead.
func (EnrollmentRole) EnumDescriptor() ([]byte, []int) {
	return file_proto_source_studyplaces_studyplaces_proto_rawDescGZIP(), []int{0}
}

type EnrollmentToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	StudyPlaceId string `protobuf:"bytes,2,opt,name=studyPlaceId,proto3" json:"studyPlaceId,omitempty"`
}

func (x *EnrollmentToken) Reset() {
	*x = EnrollmentToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_studyplaces_studyplaces_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollmentToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollmentToken) ProtoMessage() {}

func (x *EnrollmentToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_studyplaces_studyplaces_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollmentToken.ProtoReflect.Descriptor instead.
func (*EnrollmentToken) Descriptor() ([]byte, []int) {
	return file_proto_source_studyplaces_studyplaces_proto_rawDescGZIP(), []int{0}
}

func (x *EnrollmentToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *EnrollmentToken) GetStudyPlaceId() string {
	if x != nil {
		return x.StudyPlaceId
	}
	return ""
}

type Enrollment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId       string         `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	StudyPlaceId string         `protobuf:"bytes,3,opt,name=studyPlaceId,proto3" json:"studyPlaceId,omitempty"`
	UserName     string         `protobuf:"bytes,4,opt,name=userName,proto3" json:"userName,omitempty"`
	Role         EnrollmentRole `protobuf:"varint,5,opt,name=role,proto3,enum=studyplaces.EnrollmentRole" json:"role,omitempty"`
	TypeId       string         `protobuf:"bytes,6,opt,name=typeId,proto3" json:"typeId,omitempty"`
	Permissions  []string       `protobuf:"bytes,7,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Accepted     bool           `protobuf:"varint,8,opt,name=accepted,proto3" json:"accepted,omitempty"`
	Blocked      bool           `protobuf:"varint,9,opt,name=blocked,proto3" json:"blocked,omitempty"`
}

func (x *Enrollment) Reset() {
	*x = Enrollment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_studyplaces_studyplaces_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enrollment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enrollment) ProtoMessage() {}

func (x *Enrollment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_studyplaces_studyplaces_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enrollment.ProtoReflect.Descriptor instead.
func (*Enrollment) Descriptor() ([]byte, []int) {
	return file_proto_source_studyplaces_studyplaces_proto_rawDescGZIP(), []int{1}
}

func (x *Enrollment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Enrollment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Enrollment) GetStudyPlaceId() string {
	if x != nil {
		return x.StudyPlaceId
	}
	return ""
}

func (x *Enrollment) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Enrollment) GetRole() EnrollmentRole {
	if x != nil {
		return x.Role
	}
	return EnrollmentRole_TEACHER
}

func (x *Enrollment) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *Enrollment) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Enrollment) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

func (x *Enrollment) GetBlocked() bool {
	if x != nil {
		return x.Blocked
	}
	return false
}

type EnrollmentUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Login         string      `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	PictureUrl    string      `protobuf:"bytes,3,opt,name=pictureUrl,proto3" json:"pictureUrl,omitempty"`
	Email         string      `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	VerifiedEmail bool        `protobuf:"varint,5,opt,name=verifiedEmail,proto3" json:"verifiedEmail,omitempty"`
	Enrollment    *Enrollment `protobuf:"bytes,6,opt,name=enrollment,proto3" json:"enrollment,omitempty"`
}

func (x *EnrollmentUser) Reset() {
	*x = EnrollmentUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_studyplaces_studyplaces_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollmentUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollmentUser) ProtoMessage() {}

func (x *EnrollmentUser) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_studyplaces_studyplaces_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollmentUser.ProtoReflect.Descriptor instead.
func (*EnrollmentUser) Descriptor() ([]byte, []int) {
	return file_proto_source_studyplaces_studyplaces_proto_rawDescGZIP(), []int{2}
}

func (x *EnrollmentUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EnrollmentUser) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *EnrollmentUser) GetPictureUrl() string {
	if x != nil {
		return x.PictureUrl
	}
	return ""
}

func (x *EnrollmentUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EnrollmentUser) GetVerifiedEmail() bool {
	if x != nil {
		return x.VerifiedEmail
	}
	return false
}

func (x *EnrollmentUser) GetEnrollment() *Enrollment {
	if x != nil {
		return x.Enrollment
	}
	return nil
}

var File_proto_source_studyplaces_studyplaces_proto protoreflect.FileDescriptor

var file_proto_source_studyplaces_studyplaces_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x73,
	0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x79,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x74,
	0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x22, 0x4b, 0x0a, 0x0f, 0x45, 0x6e, 0x72,
	0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x79, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x95, 0x02, 0x0a, 0x0a, 0x45, 0x6e, 0x72, 0x6f, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x73, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x74,
	0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22, 0xcb,
	0x01, 0x0a, 0x0e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a,
	0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x37, 0x0a, 0x0a, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x0a, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2a, 0x40, 0x0a, 0x0e,
	0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x54, 0x45, 0x41, 0x43, 0x48, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x54, 0x55, 0x44, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x46,
	0x46, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x03, 0x32, 0x52,
	0x0a, 0x0b, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x43, 0x0a,
	0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1c, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x73, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x1a, 0x1b, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x73, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_source_studyplaces_studyplaces_proto_rawDescOnce sync.Once
	file_proto_source_studyplaces_studyplaces_proto_rawDescData = file_proto_source_studyplaces_studyplaces_proto_rawDesc
)

func file_proto_source_studyplaces_studyplaces_proto_rawDescGZIP() []byte {
	file_proto_source_studyplaces_studyplaces_proto_rawDescOnce.Do(func() {
		file_proto_source_studyplaces_studyplaces_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_source_studyplaces_studyplaces_proto_rawDescData)
	})
	return file_proto_source_studyplaces_studyplaces_proto_rawDescData
}

var file_proto_source_studyplaces_studyplaces_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_source_studyplaces_studyplaces_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_source_studyplaces_studyplaces_proto_goTypes = []interface{}{
	(EnrollmentRole)(0),     // 0: studyplaces.EnrollmentRole
	(*EnrollmentToken)(nil), // 1: studyplaces.EnrollmentToken
	(*Enrollment)(nil),      // 2: studyplaces.Enrollment
	(*EnrollmentUser)(nil),  // 3: studyplaces.EnrollmentUser
}
var file_proto_source_studyplaces_studyplaces_proto_depIdxs = []int32{
	0, // 0: studyplaces.Enrollment.role:type_name -> studyplaces.EnrollmentRole
	2, // 1: studyplaces.EnrollmentUser.enrollment:type_name -> studyplaces.Enrollment
	1, // 2: studyplaces.studyplaces.Auth:input_type -> studyplaces.EnrollmentToken
	3, // 3: studyplaces.studyplaces.Auth:output_type -> studyplaces.EnrollmentUser
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_source_studyplaces_studyplaces_proto_init() }
func file_proto_source_studyplaces_studyplaces_proto_init() {
	if File_proto_source_studyplaces_studyplaces_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_source_studyplaces_studyplaces_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollmentToken); i {
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
		file_proto_source_studyplaces_studyplaces_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enrollment); i {
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
		file_proto_source_studyplaces_studyplaces_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollmentUser); i {
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
			RawDescriptor: file_proto_source_studyplaces_studyplaces_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_source_studyplaces_studyplaces_proto_goTypes,
		DependencyIndexes: file_proto_source_studyplaces_studyplaces_proto_depIdxs,
		EnumInfos:         file_proto_source_studyplaces_studyplaces_proto_enumTypes,
		MessageInfos:      file_proto_source_studyplaces_studyplaces_proto_msgTypes,
	}.Build()
	File_proto_source_studyplaces_studyplaces_proto = out.File
	file_proto_source_studyplaces_studyplaces_proto_rawDesc = nil
	file_proto_source_studyplaces_studyplaces_proto_goTypes = nil
	file_proto_source_studyplaces_studyplaces_proto_depIdxs = nil
}
