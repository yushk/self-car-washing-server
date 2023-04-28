// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: evaluate_info.proto

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

type EvaluateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	HomeWorkId string      `protobuf:"bytes,2,opt,name=home_work_id,json=homeWorkId,proto3" json:"home_work_id" bson:"home_work_id"`
	SolutionId string      `protobuf:"bytes,3,opt,name=solution_id,json=solutionId,proto3" json:"solution_id" bson:"solution_id"`
	UserId     string      `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id" bson:"user_id"`
	Name       string      `protobuf:"bytes,5,opt,name=name,proto3" json:"name" bson:"name"`
	Evaluates  []*Evaluate `protobuf:"bytes,6,rep,name=evaluates,proto3" json:"evaluates" bson:"evaluates"`
	Note       string      `protobuf:"bytes,7,opt,name=note,proto3" json:"note" bson:"note"`
}

func (x *EvaluateInfo) Reset() {
	*x = EvaluateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evaluate_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateInfo) ProtoMessage() {}

func (x *EvaluateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_evaluate_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateInfo.ProtoReflect.Descriptor instead.
func (*EvaluateInfo) Descriptor() ([]byte, []int) {
	return file_evaluate_info_proto_rawDescGZIP(), []int{0}
}

func (x *EvaluateInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EvaluateInfo) GetHomeWorkId() string {
	if x != nil {
		return x.HomeWorkId
	}
	return ""
}

func (x *EvaluateInfo) GetSolutionId() string {
	if x != nil {
		return x.SolutionId
	}
	return ""
}

func (x *EvaluateInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EvaluateInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EvaluateInfo) GetEvaluates() []*Evaluate {
	if x != nil {
		return x.Evaluates
	}
	return nil
}

func (x *EvaluateInfo) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

type CreateEvaluateInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *EvaluateInfo `protobuf:"bytes,1,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *CreateEvaluateInfoRequest) Reset() {
	*x = CreateEvaluateInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evaluate_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEvaluateInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEvaluateInfoRequest) ProtoMessage() {}

func (x *CreateEvaluateInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evaluate_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEvaluateInfoRequest.ProtoReflect.Descriptor instead.
func (*CreateEvaluateInfoRequest) Descriptor() ([]byte, []int) {
	return file_evaluate_info_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEvaluateInfoRequest) GetData() *EvaluateInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetEvaluateInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *GetEvaluateInfoRequest) Reset() {
	*x = GetEvaluateInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evaluate_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEvaluateInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEvaluateInfoRequest) ProtoMessage() {}

func (x *GetEvaluateInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evaluate_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEvaluateInfoRequest.ProtoReflect.Descriptor instead.
func (*GetEvaluateInfoRequest) Descriptor() ([]byte, []int) {
	return file_evaluate_info_proto_rawDescGZIP(), []int{2}
}

func (x *GetEvaluateInfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateEvaluateInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Data *EvaluateInfo `protobuf:"bytes,2,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *UpdateEvaluateInfoRequest) Reset() {
	*x = UpdateEvaluateInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evaluate_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEvaluateInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEvaluateInfoRequest) ProtoMessage() {}

func (x *UpdateEvaluateInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evaluate_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEvaluateInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateEvaluateInfoRequest) Descriptor() ([]byte, []int) {
	return file_evaluate_info_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEvaluateInfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateEvaluateInfoRequest) GetData() *EvaluateInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteEvaluateInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *DeleteEvaluateInfoRequest) Reset() {
	*x = DeleteEvaluateInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evaluate_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEvaluateInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEvaluateInfoRequest) ProtoMessage() {}

func (x *DeleteEvaluateInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evaluate_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEvaluateInfoRequest.ProtoReflect.Descriptor instead.
func (*DeleteEvaluateInfoRequest) Descriptor() ([]byte, []int) {
	return file_evaluate_info_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteEvaluateInfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetEvaluateInfosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query" bson:"query"`
	Skip  int64  `protobuf:"varint,2,opt,name=skip,proto3" json:"skip" bson:"skip"`
	Limit int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit" bson:"limit"`
}

func (x *GetEvaluateInfosRequest) Reset() {
	*x = GetEvaluateInfosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evaluate_info_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEvaluateInfosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEvaluateInfosRequest) ProtoMessage() {}

func (x *GetEvaluateInfosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evaluate_info_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEvaluateInfosRequest.ProtoReflect.Descriptor instead.
func (*GetEvaluateInfosRequest) Descriptor() ([]byte, []int) {
	return file_evaluate_info_proto_rawDescGZIP(), []int{5}
}

func (x *GetEvaluateInfosRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *GetEvaluateInfosRequest) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *GetEvaluateInfosRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetEvaluateInfosReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64           `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count" bson:"total_count"`
	Items      []*EvaluateInfo `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *GetEvaluateInfosReply) Reset() {
	*x = GetEvaluateInfosReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evaluate_info_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEvaluateInfosReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEvaluateInfosReply) ProtoMessage() {}

func (x *GetEvaluateInfosReply) ProtoReflect() protoreflect.Message {
	mi := &file_evaluate_info_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEvaluateInfosReply.ProtoReflect.Descriptor instead.
func (*GetEvaluateInfosReply) Descriptor() ([]byte, []int) {
	return file_evaluate_info_proto_rawDescGZIP(), []int{6}
}

func (x *GetEvaluateInfosReply) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *GetEvaluateInfosReply) GetItems() []*EvaluateInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_evaluate_info_proto protoreflect.FileDescriptor

var file_evaluate_info_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x10, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x0c,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c,
	0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x09,
	0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x09, 0x65,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0x41, 0x0a, 0x19,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x19, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a, 0x19,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b,
	0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x60, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x62, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xbe, 0x02, 0x0a, 0x13, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x47,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70,
	0x62, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00,
	0x12, 0x47, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x1b, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x73, 0x68, 0x6b, 0x2f, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_evaluate_info_proto_rawDescOnce sync.Once
	file_evaluate_info_proto_rawDescData = file_evaluate_info_proto_rawDesc
)

func file_evaluate_info_proto_rawDescGZIP() []byte {
	file_evaluate_info_proto_rawDescOnce.Do(func() {
		file_evaluate_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_evaluate_info_proto_rawDescData)
	})
	return file_evaluate_info_proto_rawDescData
}

var file_evaluate_info_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_evaluate_info_proto_goTypes = []interface{}{
	(*EvaluateInfo)(nil),              // 0: pb.EvaluateInfo
	(*CreateEvaluateInfoRequest)(nil), // 1: pb.CreateEvaluateInfoRequest
	(*GetEvaluateInfoRequest)(nil),    // 2: pb.GetEvaluateInfoRequest
	(*UpdateEvaluateInfoRequest)(nil), // 3: pb.UpdateEvaluateInfoRequest
	(*DeleteEvaluateInfoRequest)(nil), // 4: pb.DeleteEvaluateInfoRequest
	(*GetEvaluateInfosRequest)(nil),   // 5: pb.GetEvaluateInfosRequest
	(*GetEvaluateInfosReply)(nil),     // 6: pb.GetEvaluateInfosReply
	(*Evaluate)(nil),                  // 7: pb.Evaluate
}
var file_evaluate_info_proto_depIdxs = []int32{
	7, // 0: pb.EvaluateInfo.evaluates:type_name -> pb.Evaluate
	0, // 1: pb.CreateEvaluateInfoRequest.data:type_name -> pb.EvaluateInfo
	0, // 2: pb.UpdateEvaluateInfoRequest.data:type_name -> pb.EvaluateInfo
	0, // 3: pb.GetEvaluateInfosReply.items:type_name -> pb.EvaluateInfo
	1, // 4: pb.EvaluateInfoManager.CreateEvaluateInfo:input_type -> pb.CreateEvaluateInfoRequest
	3, // 5: pb.EvaluateInfoManager.UpdateEvaluateInfo:input_type -> pb.UpdateEvaluateInfoRequest
	4, // 6: pb.EvaluateInfoManager.DeleteEvaluateInfo:input_type -> pb.DeleteEvaluateInfoRequest
	5, // 7: pb.EvaluateInfoManager.GetEvaluateInfos:input_type -> pb.GetEvaluateInfosRequest
	0, // 8: pb.EvaluateInfoManager.CreateEvaluateInfo:output_type -> pb.EvaluateInfo
	0, // 9: pb.EvaluateInfoManager.UpdateEvaluateInfo:output_type -> pb.EvaluateInfo
	0, // 10: pb.EvaluateInfoManager.DeleteEvaluateInfo:output_type -> pb.EvaluateInfo
	6, // 11: pb.EvaluateInfoManager.GetEvaluateInfos:output_type -> pb.GetEvaluateInfosReply
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_evaluate_info_proto_init() }
func file_evaluate_info_proto_init() {
	if File_evaluate_info_proto != nil {
		return
	}
	file_sport_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_evaluate_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluateInfo); i {
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
		file_evaluate_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEvaluateInfoRequest); i {
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
		file_evaluate_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEvaluateInfoRequest); i {
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
		file_evaluate_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEvaluateInfoRequest); i {
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
		file_evaluate_info_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEvaluateInfoRequest); i {
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
		file_evaluate_info_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEvaluateInfosRequest); i {
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
		file_evaluate_info_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEvaluateInfosReply); i {
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
			RawDescriptor: file_evaluate_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_evaluate_info_proto_goTypes,
		DependencyIndexes: file_evaluate_info_proto_depIdxs,
		MessageInfos:      file_evaluate_info_proto_msgTypes,
	}.Build()
	File_evaluate_info_proto = out.File
	file_evaluate_info_proto_rawDesc = nil
	file_evaluate_info_proto_goTypes = nil
	file_evaluate_info_proto_depIdxs = nil
}
