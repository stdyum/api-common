// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: proto/source/types_registry/types_registry.proto

package types_registry

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

type TypesIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	StudyPlaceId string   `protobuf:"bytes,2,opt,name=studyPlaceId,proto3" json:"studyPlaceId,omitempty"`
	GroupsIds    []string `protobuf:"bytes,3,rep,name=groupsIds,proto3" json:"groupsIds,omitempty"`
	RoomsIds     []string `protobuf:"bytes,4,rep,name=roomsIds,proto3" json:"roomsIds,omitempty"`
	StudentIds   []string `protobuf:"bytes,5,rep,name=studentIds,proto3" json:"studentIds,omitempty"`
	SubjectsIds  []string `protobuf:"bytes,6,rep,name=subjectsIds,proto3" json:"subjectsIds,omitempty"`
	TeachersIds  []string `protobuf:"bytes,7,rep,name=teachersIds,proto3" json:"teachersIds,omitempty"`
}

func (x *TypesIds) Reset() {
	*x = TypesIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypesIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypesIds) ProtoMessage() {}

func (x *TypesIds) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypesIds.ProtoReflect.Descriptor instead.
func (*TypesIds) Descriptor() ([]byte, []int) {
	return file_proto_source_types_registry_types_registry_proto_rawDescGZIP(), []int{0}
}

func (x *TypesIds) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TypesIds) GetStudyPlaceId() string {
	if x != nil {
		return x.StudyPlaceId
	}
	return ""
}

func (x *TypesIds) GetGroupsIds() []string {
	if x != nil {
		return x.GroupsIds
	}
	return nil
}

func (x *TypesIds) GetRoomsIds() []string {
	if x != nil {
		return x.RoomsIds
	}
	return nil
}

func (x *TypesIds) GetStudentIds() []string {
	if x != nil {
		return x.StudentIds
	}
	return nil
}

func (x *TypesIds) GetSubjectsIds() []string {
	if x != nil {
		return x.SubjectsIds
	}
	return nil
}

func (x *TypesIds) GetTeachersIds() []string {
	if x != nil {
		return x.TeachersIds
	}
	return nil
}

type TypesModels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups   map[string]*Group   `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Rooms    map[string]*Room    `protobuf:"bytes,2,rep,name=rooms,proto3" json:"rooms,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Students map[string]*Student `protobuf:"bytes,3,rep,name=students,proto3" json:"students,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Subjects map[string]*Subject `protobuf:"bytes,4,rep,name=subjects,proto3" json:"subjects,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Teachers map[string]*Teacher `protobuf:"bytes,5,rep,name=teachers,proto3" json:"teachers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TypesModels) Reset() {
	*x = TypesModels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypesModels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypesModels) ProtoMessage() {}

func (x *TypesModels) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypesModels.ProtoReflect.Descriptor instead.
func (*TypesModels) Descriptor() ([]byte, []int) {
	return file_proto_source_types_registry_types_registry_proto_rawDescGZIP(), []int{1}
}

func (x *TypesModels) GetGroups() map[string]*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *TypesModels) GetRooms() map[string]*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

func (x *TypesModels) GetStudents() map[string]*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

func (x *TypesModels) GetSubjects() map[string]*Subject {
	if x != nil {
		return x.Subjects
	}
	return nil
}

func (x *TypesModels) GetTeachers() map[string]*Teacher {
	if x != nil {
		return x.Teachers
	}
	return nil
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_proto_source_types_registry_types_registry_proto_rawDescGZIP(), []int{2}
}

func (x *Group) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_proto_source_types_registry_types_registry_proto_rawDescGZIP(), []int{3}
}

func (x *Room) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_proto_source_types_registry_types_registry_proto_rawDescGZIP(), []int{4}
}

func (x *Student) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Subject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Subject) Reset() {
	*x = Subject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subject) ProtoMessage() {}

func (x *Subject) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subject.ProtoReflect.Descriptor instead.
func (*Subject) Descriptor() ([]byte, []int) {
	return file_proto_source_types_registry_types_registry_proto_rawDescGZIP(), []int{5}
}

func (x *Subject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Teacher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Teacher) Reset() {
	*x = Teacher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Teacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Teacher) ProtoMessage() {}

func (x *Teacher) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Teacher.ProtoReflect.Descriptor instead.
func (*Teacher) Descriptor() ([]byte, []int) {
	return file_proto_source_types_registry_types_registry_proto_rawDescGZIP(), []int{6}
}

func (x *Teacher) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Teacher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GroupId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	StudyPlaceId string `protobuf:"bytes,2,opt,name=studyPlaceId,proto3" json:"studyPlaceId,omitempty"`
	Uuid         string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GroupId) Reset() {
	*x = GroupId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupId) ProtoMessage() {}

func (x *GroupId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupId.ProtoReflect.Descriptor instead.
func (*GroupId) Descriptor() ([]byte, []int) {
	return file_proto_source_types_registry_types_registry_proto_rawDescGZIP(), []int{7}
}

func (x *GroupId) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GroupId) GetStudyPlaceId() string {
	if x != nil {
		return x.StudyPlaceId
	}
	return ""
}

func (x *GroupId) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type StudentId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	StudyPlaceId string `protobuf:"bytes,2,opt,name=studyPlaceId,proto3" json:"studyPlaceId,omitempty"`
	Uuid         string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *StudentId) Reset() {
	*x = StudentId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentId) ProtoMessage() {}

func (x *StudentId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentId.ProtoReflect.Descriptor instead.
func (*StudentId) Descriptor() ([]byte, []int) {
	return file_proto_source_types_registry_types_registry_proto_rawDescGZIP(), []int{8}
}

func (x *StudentId) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *StudentId) GetStudyPlaceId() string {
	if x != nil {
		return x.StudyPlaceId
	}
	return ""
}

func (x *StudentId) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type Students struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Student `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Students) Reset() {
	*x = Students{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Students) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Students) ProtoMessage() {}

func (x *Students) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Students.ProtoReflect.Descriptor instead.
func (*Students) Descriptor() ([]byte, []int) {
	return file_proto_source_types_registry_types_registry_proto_rawDescGZIP(), []int{9}
}

func (x *Students) GetList() []*Student {
	if x != nil {
		return x.List
	}
	return nil
}

type Groups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Group `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Groups) Reset() {
	*x = Groups{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Groups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Groups) ProtoMessage() {}

func (x *Groups) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_types_registry_types_registry_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Groups.ProtoReflect.Descriptor instead.
func (*Groups) Descriptor() ([]byte, []int) {
	return file_proto_source_types_registry_types_registry_proto_rawDescGZIP(), []int{10}
}

func (x *Groups) GetList() []*Group {
	if x != nil {
		return x.List
	}
	return nil
}

var File_proto_source_types_registry_types_registry_proto protoreflect.FileDescriptor

var file_proto_source_types_registry_types_registry_proto_rawDesc = []byte{
	0x0a, 0x30, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x22,
	0xe2, 0x01, 0x0a, 0x08, 0x54, 0x79, 0x70, 0x65, 0x73, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x79, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x49, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x49, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x49, 0x64, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x49, 0x64, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x49, 0x64, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x49,
	0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x49, 0x64,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x73, 0x49, 0x64, 0x73, 0x22, 0xe7, 0x05, 0x0a, 0x0b, 0x54, 0x79, 0x70, 0x65, 0x73, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x12, 0x3c, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x12, 0x39, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x42, 0x0a,
	0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x42, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x42, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x1a, 0x4d, 0x0a, 0x0b, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x75, 0x64,
	0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x51, 0x0a, 0x0d, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x51, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x75,
	0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x51, 0x0a, 0x0d, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2b,
	0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x04, 0x52,
	0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x07, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x75,
	0x64, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x59, 0x0a,
	0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x08, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x30,
	0x0a, 0x06, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x32, 0xdb, 0x01, 0x0a, 0x0d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x12, 0x42, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x42, 0x79,
	0x49, 0x64, 0x73, 0x12, 0x15, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x49, 0x64, 0x73, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x75,
	0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x2e, 0x73,
	0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x1a, 0x15, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12,
	0x16, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x00, 0x42, 0x11,
	0x5a, 0x0f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_source_types_registry_types_registry_proto_rawDescOnce sync.Once
	file_proto_source_types_registry_types_registry_proto_rawDescData = file_proto_source_types_registry_types_registry_proto_rawDesc
)

func file_proto_source_types_registry_types_registry_proto_rawDescGZIP() []byte {
	file_proto_source_types_registry_types_registry_proto_rawDescOnce.Do(func() {
		file_proto_source_types_registry_types_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_source_types_registry_types_registry_proto_rawDescData)
	})
	return file_proto_source_types_registry_types_registry_proto_rawDescData
}

var file_proto_source_types_registry_types_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_proto_source_types_registry_types_registry_proto_goTypes = []interface{}{
	(*TypesIds)(nil),    // 0: studyplaces.TypesIds
	(*TypesModels)(nil), // 1: studyplaces.TypesModels
	(*Group)(nil),       // 2: studyplaces.Group
	(*Room)(nil),        // 3: studyplaces.Room
	(*Student)(nil),     // 4: studyplaces.Student
	(*Subject)(nil),     // 5: studyplaces.Subject
	(*Teacher)(nil),     // 6: studyplaces.Teacher
	(*GroupId)(nil),     // 7: studyplaces.GroupId
	(*StudentId)(nil),   // 8: studyplaces.StudentId
	(*Students)(nil),    // 9: studyplaces.Students
	(*Groups)(nil),      // 10: studyplaces.Groups
	nil,                 // 11: studyplaces.TypesModels.GroupsEntry
	nil,                 // 12: studyplaces.TypesModels.RoomsEntry
	nil,                 // 13: studyplaces.TypesModels.StudentsEntry
	nil,                 // 14: studyplaces.TypesModels.SubjectsEntry
	nil,                 // 15: studyplaces.TypesModels.TeachersEntry
}
var file_proto_source_types_registry_types_registry_proto_depIdxs = []int32{
	11, // 0: studyplaces.TypesModels.groups:type_name -> studyplaces.TypesModels.GroupsEntry
	12, // 1: studyplaces.TypesModels.rooms:type_name -> studyplaces.TypesModels.RoomsEntry
	13, // 2: studyplaces.TypesModels.students:type_name -> studyplaces.TypesModels.StudentsEntry
	14, // 3: studyplaces.TypesModels.subjects:type_name -> studyplaces.TypesModels.SubjectsEntry
	15, // 4: studyplaces.TypesModels.teachers:type_name -> studyplaces.TypesModels.TeachersEntry
	4,  // 5: studyplaces.Students.list:type_name -> studyplaces.Student
	2,  // 6: studyplaces.Groups.list:type_name -> studyplaces.Group
	2,  // 7: studyplaces.TypesModels.GroupsEntry.value:type_name -> studyplaces.Group
	3,  // 8: studyplaces.TypesModels.RoomsEntry.value:type_name -> studyplaces.Room
	4,  // 9: studyplaces.TypesModels.StudentsEntry.value:type_name -> studyplaces.Student
	5,  // 10: studyplaces.TypesModels.SubjectsEntry.value:type_name -> studyplaces.Subject
	6,  // 11: studyplaces.TypesModels.TeachersEntry.value:type_name -> studyplaces.Teacher
	0,  // 12: studyplaces.typesRegistry.GetTypesByIds:input_type -> studyplaces.TypesIds
	7,  // 13: studyplaces.typesRegistry.GetStudentsInGroup:input_type -> studyplaces.GroupId
	8,  // 14: studyplaces.typesRegistry.GetStudentGroups:input_type -> studyplaces.StudentId
	1,  // 15: studyplaces.typesRegistry.GetTypesByIds:output_type -> studyplaces.TypesModels
	9,  // 16: studyplaces.typesRegistry.GetStudentsInGroup:output_type -> studyplaces.Students
	10, // 17: studyplaces.typesRegistry.GetStudentGroups:output_type -> studyplaces.Groups
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_proto_source_types_registry_types_registry_proto_init() }
func file_proto_source_types_registry_types_registry_proto_init() {
	if File_proto_source_types_registry_types_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_source_types_registry_types_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypesIds); i {
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
		file_proto_source_types_registry_types_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypesModels); i {
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
		file_proto_source_types_registry_types_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_proto_source_types_registry_types_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_proto_source_types_registry_types_registry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
		file_proto_source_types_registry_types_registry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subject); i {
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
		file_proto_source_types_registry_types_registry_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Teacher); i {
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
		file_proto_source_types_registry_types_registry_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupId); i {
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
		file_proto_source_types_registry_types_registry_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentId); i {
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
		file_proto_source_types_registry_types_registry_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Students); i {
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
		file_proto_source_types_registry_types_registry_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Groups); i {
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
			RawDescriptor: file_proto_source_types_registry_types_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_source_types_registry_types_registry_proto_goTypes,
		DependencyIndexes: file_proto_source_types_registry_types_registry_proto_depIdxs,
		MessageInfos:      file_proto_source_types_registry_types_registry_proto_msgTypes,
	}.Build()
	File_proto_source_types_registry_types_registry_proto = out.File
	file_proto_source_types_registry_types_registry_proto_rawDesc = nil
	file_proto_source_types_registry_types_registry_proto_goTypes = nil
	file_proto_source_types_registry_types_registry_proto_depIdxs = nil
}
