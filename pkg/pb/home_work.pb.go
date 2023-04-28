// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: home_work.proto

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

type HomeWork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	SportItem       *SportItem  `protobuf:"bytes,2,opt,name=sport_item,json=sportItem,proto3" json:"sport_item" bson:"sport_item"`
	SportItemId     string      `protobuf:"bytes,3,opt,name=sport_item_id,json=sportItemId,proto3" json:"sport_item_id" bson:"sport_item_id"`
	Teacher         string      `protobuf:"bytes,4,opt,name=teacher,proto3" json:"teacher" bson:"teacher"`
	StartTime       int64       `protobuf:"varint,5,opt,name=start_time,json=startTime,proto3" json:"start_time" bson:"start_time"`
	EndTime         int64       `protobuf:"varint,6,opt,name=end_time,json=endTime,proto3" json:"end_time" bson:"end_time"`
	QRCode          string      `protobuf:"bytes,7,opt,name=QR_code,json=QRCode,proto3" json:"QR_code" bson:"QR_code"`
	Number          int32       `protobuf:"varint,8,opt,name=number,proto3" json:"number" bson:"number"`
	Desc            string      `protobuf:"bytes,9,opt,name=desc,proto3" json:"desc" bson:"desc"`
	SolutionsInfoId []string    `protobuf:"bytes,10,rep,name=solutions_info_id,json=solutionsInfoId,proto3" json:"solutions_info_id" bson:"solutions_info_id"`
	SolutionsInfo   []*Solution `protobuf:"bytes,11,rep,name=solutions_info,json=solutionsInfo,proto3" json:"solutions_info" bson:"solutions_info"`
	Title           string      `protobuf:"bytes,12,opt,name=title,proto3" json:"title" bson:"title"`
}

func (x *HomeWork) Reset() {
	*x = HomeWork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_work_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeWork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeWork) ProtoMessage() {}

func (x *HomeWork) ProtoReflect() protoreflect.Message {
	mi := &file_home_work_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeWork.ProtoReflect.Descriptor instead.
func (*HomeWork) Descriptor() ([]byte, []int) {
	return file_home_work_proto_rawDescGZIP(), []int{0}
}

func (x *HomeWork) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HomeWork) GetSportItem() *SportItem {
	if x != nil {
		return x.SportItem
	}
	return nil
}

func (x *HomeWork) GetSportItemId() string {
	if x != nil {
		return x.SportItemId
	}
	return ""
}

func (x *HomeWork) GetTeacher() string {
	if x != nil {
		return x.Teacher
	}
	return ""
}

func (x *HomeWork) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *HomeWork) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *HomeWork) GetQRCode() string {
	if x != nil {
		return x.QRCode
	}
	return ""
}

func (x *HomeWork) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *HomeWork) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *HomeWork) GetSolutionsInfoId() []string {
	if x != nil {
		return x.SolutionsInfoId
	}
	return nil
}

func (x *HomeWork) GetSolutionsInfo() []*Solution {
	if x != nil {
		return x.SolutionsInfo
	}
	return nil
}

func (x *HomeWork) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type CreateHomeWorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *HomeWork `protobuf:"bytes,1,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *CreateHomeWorkRequest) Reset() {
	*x = CreateHomeWorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_work_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHomeWorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHomeWorkRequest) ProtoMessage() {}

func (x *CreateHomeWorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_home_work_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHomeWorkRequest.ProtoReflect.Descriptor instead.
func (*CreateHomeWorkRequest) Descriptor() ([]byte, []int) {
	return file_home_work_proto_rawDescGZIP(), []int{1}
}

func (x *CreateHomeWorkRequest) GetData() *HomeWork {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetHomeWorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *GetHomeWorkRequest) Reset() {
	*x = GetHomeWorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_work_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeWorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeWorkRequest) ProtoMessage() {}

func (x *GetHomeWorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_home_work_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeWorkRequest.ProtoReflect.Descriptor instead.
func (*GetHomeWorkRequest) Descriptor() ([]byte, []int) {
	return file_home_work_proto_rawDescGZIP(), []int{2}
}

func (x *GetHomeWorkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateHomeWorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Data *HomeWork `protobuf:"bytes,2,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *UpdateHomeWorkRequest) Reset() {
	*x = UpdateHomeWorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_work_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHomeWorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHomeWorkRequest) ProtoMessage() {}

func (x *UpdateHomeWorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_home_work_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHomeWorkRequest.ProtoReflect.Descriptor instead.
func (*UpdateHomeWorkRequest) Descriptor() ([]byte, []int) {
	return file_home_work_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateHomeWorkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateHomeWorkRequest) GetData() *HomeWork {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteHomeWorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *DeleteHomeWorkRequest) Reset() {
	*x = DeleteHomeWorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_work_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHomeWorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHomeWorkRequest) ProtoMessage() {}

func (x *DeleteHomeWorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_home_work_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHomeWorkRequest.ProtoReflect.Descriptor instead.
func (*DeleteHomeWorkRequest) Descriptor() ([]byte, []int) {
	return file_home_work_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteHomeWorkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetHomeWorksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query" bson:"query"`
	Skip  int64  `protobuf:"varint,2,opt,name=skip,proto3" json:"skip" bson:"skip"`
	Limit int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit" bson:"limit"`
}

func (x *GetHomeWorksRequest) Reset() {
	*x = GetHomeWorksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_work_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeWorksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeWorksRequest) ProtoMessage() {}

func (x *GetHomeWorksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_home_work_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeWorksRequest.ProtoReflect.Descriptor instead.
func (*GetHomeWorksRequest) Descriptor() ([]byte, []int) {
	return file_home_work_proto_rawDescGZIP(), []int{5}
}

func (x *GetHomeWorksRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *GetHomeWorksRequest) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *GetHomeWorksRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetHomeWorksReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64       `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count" bson:"total_count"`
	Items      []*HomeWork `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *GetHomeWorksReply) Reset() {
	*x = GetHomeWorksReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_work_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeWorksReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeWorksReply) ProtoMessage() {}

func (x *GetHomeWorksReply) ProtoReflect() protoreflect.Message {
	mi := &file_home_work_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeWorksReply.ProtoReflect.Descriptor instead.
func (*GetHomeWorksReply) Descriptor() ([]byte, []int) {
	return file_home_work_proto_rawDescGZIP(), []int{6}
}

func (x *GetHomeWorksReply) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *GetHomeWorksReply) GetItems() []*HomeWork {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_home_work_proto protoreflect.FileDescriptor

var file_home_work_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x10, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x02, 0x0a, 0x08, 0x48, 0x6f, 0x6d, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70,
	0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x51, 0x52,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x51, 0x52, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12,
	0x2a, 0x0a, 0x11, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0e, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0d, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x39, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x55, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x58, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x6d,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xc1, 0x02, 0x0a,
	0x0f, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x3b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f,
	0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x22, 0x00, 0x12, 0x35, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f,
	0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48,
	0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x17,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x48, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x75, 0x73, 0x68, 0x6b, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_home_work_proto_rawDescOnce sync.Once
	file_home_work_proto_rawDescData = file_home_work_proto_rawDesc
)

func file_home_work_proto_rawDescGZIP() []byte {
	file_home_work_proto_rawDescOnce.Do(func() {
		file_home_work_proto_rawDescData = protoimpl.X.CompressGZIP(file_home_work_proto_rawDescData)
	})
	return file_home_work_proto_rawDescData
}

var file_home_work_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_home_work_proto_goTypes = []interface{}{
	(*HomeWork)(nil),              // 0: pb.HomeWork
	(*CreateHomeWorkRequest)(nil), // 1: pb.CreateHomeWorkRequest
	(*GetHomeWorkRequest)(nil),    // 2: pb.GetHomeWorkRequest
	(*UpdateHomeWorkRequest)(nil), // 3: pb.UpdateHomeWorkRequest
	(*DeleteHomeWorkRequest)(nil), // 4: pb.DeleteHomeWorkRequest
	(*GetHomeWorksRequest)(nil),   // 5: pb.GetHomeWorksRequest
	(*GetHomeWorksReply)(nil),     // 6: pb.GetHomeWorksReply
	(*SportItem)(nil),             // 7: pb.SportItem
	(*Solution)(nil),              // 8: pb.Solution
}
var file_home_work_proto_depIdxs = []int32{
	7,  // 0: pb.HomeWork.sport_item:type_name -> pb.SportItem
	8,  // 1: pb.HomeWork.solutions_info:type_name -> pb.Solution
	0,  // 2: pb.CreateHomeWorkRequest.data:type_name -> pb.HomeWork
	0,  // 3: pb.UpdateHomeWorkRequest.data:type_name -> pb.HomeWork
	0,  // 4: pb.GetHomeWorksReply.items:type_name -> pb.HomeWork
	1,  // 5: pb.HomeWorkManager.CreateHomeWork:input_type -> pb.CreateHomeWorkRequest
	2,  // 6: pb.HomeWorkManager.GetHomeWork:input_type -> pb.GetHomeWorkRequest
	3,  // 7: pb.HomeWorkManager.UpdateHomeWork:input_type -> pb.UpdateHomeWorkRequest
	4,  // 8: pb.HomeWorkManager.DeleteHomeWork:input_type -> pb.DeleteHomeWorkRequest
	5,  // 9: pb.HomeWorkManager.GetHomeWorks:input_type -> pb.GetHomeWorksRequest
	0,  // 10: pb.HomeWorkManager.CreateHomeWork:output_type -> pb.HomeWork
	0,  // 11: pb.HomeWorkManager.GetHomeWork:output_type -> pb.HomeWork
	0,  // 12: pb.HomeWorkManager.UpdateHomeWork:output_type -> pb.HomeWork
	0,  // 13: pb.HomeWorkManager.DeleteHomeWork:output_type -> pb.HomeWork
	6,  // 14: pb.HomeWorkManager.GetHomeWorks:output_type -> pb.GetHomeWorksReply
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_home_work_proto_init() }
func file_home_work_proto_init() {
	if File_home_work_proto != nil {
		return
	}
	file_sport_item_proto_init()
	file_solution_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_home_work_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeWork); i {
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
		file_home_work_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHomeWorkRequest); i {
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
		file_home_work_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHomeWorkRequest); i {
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
		file_home_work_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHomeWorkRequest); i {
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
		file_home_work_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHomeWorkRequest); i {
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
		file_home_work_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHomeWorksRequest); i {
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
		file_home_work_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHomeWorksReply); i {
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
			RawDescriptor: file_home_work_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_home_work_proto_goTypes,
		DependencyIndexes: file_home_work_proto_depIdxs,
		MessageInfos:      file_home_work_proto_msgTypes,
	}.Build()
	File_home_work_proto = out.File
	file_home_work_proto_rawDesc = nil
	file_home_work_proto_goTypes = nil
	file_home_work_proto_depIdxs = nil
}
