// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: proto/source/schedule/schedule.proto

package schedule

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

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	StudyPlaceId string `protobuf:"bytes,2,opt,name=studyPlaceId,proto3" json:"studyPlaceId,omitempty"`
	Uuid         string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_schedule_schedule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_schedule_schedule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_proto_source_schedule_schedule_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UUID) GetStudyPlaceId() string {
	if x != nil {
		return x.StudyPlaceId
	}
	return ""
}

func (x *UUID) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type Lesson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StudyPlaceId   string `protobuf:"bytes,2,opt,name=studyPlaceId,proto3" json:"studyPlaceId,omitempty"`
	GroupId        string `protobuf:"bytes,3,opt,name=groupId,proto3" json:"groupId,omitempty"`
	RoomId         string `protobuf:"bytes,4,opt,name=roomId,proto3" json:"roomId,omitempty"`
	SubjectId      string `protobuf:"bytes,5,opt,name=subjectId,proto3" json:"subjectId,omitempty"`
	TeacherId      string `protobuf:"bytes,6,opt,name=teacherId,proto3" json:"teacherId,omitempty"`
	StartTime      int64  `protobuf:"varint,7,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime        int64  `protobuf:"varint,8,opt,name=endTime,proto3" json:"endTime,omitempty"`
	LessonIndex    int32  `protobuf:"varint,9,opt,name=lessonIndex,proto3" json:"lessonIndex,omitempty"`
	PrimaryColor   string `protobuf:"bytes,10,opt,name=primaryColor,proto3" json:"primaryColor,omitempty"`
	SecondaryColor string `protobuf:"bytes,11,opt,name=secondaryColor,proto3" json:"secondaryColor,omitempty"`
}

func (x *Lesson) Reset() {
	*x = Lesson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_schedule_schedule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lesson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lesson) ProtoMessage() {}

func (x *Lesson) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_schedule_schedule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lesson.ProtoReflect.Descriptor instead.
func (*Lesson) Descriptor() ([]byte, []int) {
	return file_proto_source_schedule_schedule_proto_rawDescGZIP(), []int{1}
}

func (x *Lesson) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Lesson) GetStudyPlaceId() string {
	if x != nil {
		return x.StudyPlaceId
	}
	return ""
}

func (x *Lesson) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Lesson) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *Lesson) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *Lesson) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

func (x *Lesson) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Lesson) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *Lesson) GetLessonIndex() int32 {
	if x != nil {
		return x.LessonIndex
	}
	return 0
}

func (x *Lesson) GetPrimaryColor() string {
	if x != nil {
		return x.PrimaryColor
	}
	return ""
}

func (x *Lesson) GetSecondaryColor() string {
	if x != nil {
		return x.SecondaryColor
	}
	return ""
}

type Lessons struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Lesson `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Lessons) Reset() {
	*x = Lessons{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_schedule_schedule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lessons) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lessons) ProtoMessage() {}

func (x *Lessons) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_schedule_schedule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lessons.ProtoReflect.Descriptor instead.
func (*Lessons) Descriptor() ([]byte, []int) {
	return file_proto_source_schedule_schedule_proto_rawDescGZIP(), []int{2}
}

func (x *Lessons) GetList() []*Lesson {
	if x != nil {
		return x.List
	}
	return nil
}

type EntriesFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	StudyPlaceId string   `protobuf:"bytes,2,opt,name=studyPlaceId,proto3" json:"studyPlaceId,omitempty"`
	TeacherId    string   `protobuf:"bytes,3,opt,name=teacherId,proto3" json:"teacherId,omitempty"`
	GroupIds     []string `protobuf:"bytes,4,rep,name=groupIds,proto3" json:"groupIds,omitempty"`
	SubjectId    string   `protobuf:"bytes,5,opt,name=subjectId,proto3" json:"subjectId,omitempty"`
	Cursor       string   `protobuf:"bytes,6,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit        int32    `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *EntriesFilter) Reset() {
	*x = EntriesFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_schedule_schedule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntriesFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntriesFilter) ProtoMessage() {}

func (x *EntriesFilter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_schedule_schedule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntriesFilter.ProtoReflect.Descriptor instead.
func (*EntriesFilter) Descriptor() ([]byte, []int) {
	return file_proto_source_schedule_schedule_proto_rawDescGZIP(), []int{3}
}

func (x *EntriesFilter) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *EntriesFilter) GetStudyPlaceId() string {
	if x != nil {
		return x.StudyPlaceId
	}
	return ""
}

func (x *EntriesFilter) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

func (x *EntriesFilter) GetGroupIds() []string {
	if x != nil {
		return x.GroupIds
	}
	return nil
}

func (x *EntriesFilter) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *EntriesFilter) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *EntriesFilter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type Entries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*Entry `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Next  string   `protobuf:"bytes,2,opt,name=next,proto3" json:"next,omitempty"`
	Limit int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Entries) Reset() {
	*x = Entries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_schedule_schedule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entries) ProtoMessage() {}

func (x *Entries) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_schedule_schedule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entries.ProtoReflect.Descriptor instead.
func (*Entries) Descriptor() ([]byte, []int) {
	return file_proto_source_schedule_schedule_proto_rawDescGZIP(), []int{4}
}

func (x *Entries) GetList() []*Entry {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *Entries) GetNext() string {
	if x != nil {
		return x.Next
	}
	return ""
}

func (x *Entries) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherId string `protobuf:"bytes,1,opt,name=teacherId,proto3" json:"teacherId,omitempty"`
	GroupId   string `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	SubjectId string `protobuf:"bytes,3,opt,name=subjectId,proto3" json:"subjectId,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_source_schedule_schedule_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_source_schedule_schedule_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_proto_source_schedule_schedule_proto_rawDescGZIP(), []int{5}
}

func (x *Entry) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

func (x *Entry) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Entry) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

var File_proto_source_schedule_schedule_proto protoreflect.FileDescriptor

var file_proto_source_schedule_schedule_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x54, 0x0a, 0x04,
	0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74,
	0x75, 0x64, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x22, 0xd0, 0x02, 0x0a, 0x06, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x73, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x26, 0x0a,
	0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x2b, 0x0a, 0x07, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73,
	0x12, 0x20, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0xcf, 0x01, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74,
	0x75, 0x64, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x54, 0x0a, 0x07, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x1f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5d, 0x0a, 0x05, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x32, 0xa5, 0x01, 0x0a, 0x08, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55,
	0x55, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x73, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x1a, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22,
	0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_source_schedule_schedule_proto_rawDescOnce sync.Once
	file_proto_source_schedule_schedule_proto_rawDescData = file_proto_source_schedule_schedule_proto_rawDesc
)

func file_proto_source_schedule_schedule_proto_rawDescGZIP() []byte {
	file_proto_source_schedule_schedule_proto_rawDescOnce.Do(func() {
		file_proto_source_schedule_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_source_schedule_schedule_proto_rawDescData)
	})
	return file_proto_source_schedule_schedule_proto_rawDescData
}

var file_proto_source_schedule_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_source_schedule_schedule_proto_goTypes = []interface{}{
	(*UUID)(nil),          // 0: auth.UUID
	(*Lesson)(nil),        // 1: auth.Lesson
	(*Lessons)(nil),       // 2: auth.Lessons
	(*EntriesFilter)(nil), // 3: auth.EntriesFilter
	(*Entries)(nil),       // 4: auth.Entries
	(*Entry)(nil),         // 5: auth.Entry
}
var file_proto_source_schedule_schedule_proto_depIdxs = []int32{
	1, // 0: auth.Lessons.list:type_name -> auth.Lesson
	5, // 1: auth.Entries.list:type_name -> auth.Entry
	0, // 2: auth.schedule.GetLessonById:input_type -> auth.UUID
	3, // 3: auth.schedule.GetLessons:input_type -> auth.EntriesFilter
	3, // 4: auth.schedule.GetUniqueEntries:input_type -> auth.EntriesFilter
	1, // 5: auth.schedule.GetLessonById:output_type -> auth.Lesson
	2, // 6: auth.schedule.GetLessons:output_type -> auth.Lessons
	4, // 7: auth.schedule.GetUniqueEntries:output_type -> auth.Entries
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_source_schedule_schedule_proto_init() }
func file_proto_source_schedule_schedule_proto_init() {
	if File_proto_source_schedule_schedule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_source_schedule_schedule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_proto_source_schedule_schedule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lesson); i {
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
		file_proto_source_schedule_schedule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lessons); i {
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
		file_proto_source_schedule_schedule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntriesFilter); i {
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
		file_proto_source_schedule_schedule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entries); i {
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
		file_proto_source_schedule_schedule_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
			RawDescriptor: file_proto_source_schedule_schedule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_source_schedule_schedule_proto_goTypes,
		DependencyIndexes: file_proto_source_schedule_schedule_proto_depIdxs,
		MessageInfos:      file_proto_source_schedule_schedule_proto_msgTypes,
	}.Build()
	File_proto_source_schedule_schedule_proto = out.File
	file_proto_source_schedule_schedule_proto_rawDesc = nil
	file_proto_source_schedule_schedule_proto_goTypes = nil
	file_proto_source_schedule_schedule_proto_depIdxs = nil
}
