// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: stats/v1/stats.proto

package stats_v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetPlantsStatsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *TimeFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetPlantsStatsV1Request) Reset() {
	*x = GetPlantsStatsV1Request{}
	mi := &file_stats_v1_stats_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPlantsStatsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlantsStatsV1Request) ProtoMessage() {}

func (x *GetPlantsStatsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_stats_v1_stats_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlantsStatsV1Request.ProtoReflect.Descriptor instead.
func (*GetPlantsStatsV1Request) Descriptor() ([]byte, []int) {
	return file_stats_v1_stats_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlantsStatsV1Request) GetFilter() *TimeFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetPlantsStatsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64                                 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Stats []*GetPlantsStatsV1Response_StatsInfo `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *GetPlantsStatsV1Response) Reset() {
	*x = GetPlantsStatsV1Response{}
	mi := &file_stats_v1_stats_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPlantsStatsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlantsStatsV1Response) ProtoMessage() {}

func (x *GetPlantsStatsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_stats_v1_stats_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlantsStatsV1Response.ProtoReflect.Descriptor instead.
func (*GetPlantsStatsV1Response) Descriptor() ([]byte, []int) {
	return file_stats_v1_stats_proto_rawDescGZIP(), []int{1}
}

func (x *GetPlantsStatsV1Response) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetPlantsStatsV1Response) GetStats() []*GetPlantsStatsV1Response_StatsInfo {
	if x != nil {
		return x.Stats
	}
	return nil
}

type GetBuyStatsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *TimeFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetBuyStatsV1Request) Reset() {
	*x = GetBuyStatsV1Request{}
	mi := &file_stats_v1_stats_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBuyStatsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBuyStatsV1Request) ProtoMessage() {}

func (x *GetBuyStatsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_stats_v1_stats_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBuyStatsV1Request.ProtoReflect.Descriptor instead.
func (*GetBuyStatsV1Request) Descriptor() ([]byte, []int) {
	return file_stats_v1_stats_proto_rawDescGZIP(), []int{2}
}

func (x *GetBuyStatsV1Request) GetFilter() *TimeFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetBuyStatsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64                              `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Stats []*GetBuyStatsV1Response_StatsInfo `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *GetBuyStatsV1Response) Reset() {
	*x = GetBuyStatsV1Response{}
	mi := &file_stats_v1_stats_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBuyStatsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBuyStatsV1Response) ProtoMessage() {}

func (x *GetBuyStatsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_stats_v1_stats_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBuyStatsV1Response.ProtoReflect.Descriptor instead.
func (*GetBuyStatsV1Response) Descriptor() ([]byte, []int) {
	return file_stats_v1_stats_proto_rawDescGZIP(), []int{3}
}

func (x *GetBuyStatsV1Response) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetBuyStatsV1Response) GetStats() []*GetBuyStatsV1Response_StatsInfo {
	if x != nil {
		return x.Stats
	}
	return nil
}

type GetTradeStatsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *TimeFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetTradeStatsV1Request) Reset() {
	*x = GetTradeStatsV1Request{}
	mi := &file_stats_v1_stats_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTradeStatsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTradeStatsV1Request) ProtoMessage() {}

func (x *GetTradeStatsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_stats_v1_stats_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTradeStatsV1Request.ProtoReflect.Descriptor instead.
func (*GetTradeStatsV1Request) Descriptor() ([]byte, []int) {
	return file_stats_v1_stats_proto_rawDescGZIP(), []int{4}
}

func (x *GetTradeStatsV1Request) GetFilter() *TimeFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetTradeStatsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64                                `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Stats []*GetTradeStatsV1Response_StatsInfo `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *GetTradeStatsV1Response) Reset() {
	*x = GetTradeStatsV1Response{}
	mi := &file_stats_v1_stats_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTradeStatsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTradeStatsV1Response) ProtoMessage() {}

func (x *GetTradeStatsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_stats_v1_stats_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTradeStatsV1Response.ProtoReflect.Descriptor instead.
func (*GetTradeStatsV1Response) Descriptor() ([]byte, []int) {
	return file_stats_v1_stats_proto_rawDescGZIP(), []int{5}
}

func (x *GetTradeStatsV1Response) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetTradeStatsV1Response) GetStats() []*GetTradeStatsV1Response_StatsInfo {
	if x != nil {
		return x.Stats
	}
	return nil
}

type TimeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeFrom *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timeFrom,proto3" json:"timeFrom,omitempty"`
	TimeTo   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timeTo,proto3" json:"timeTo,omitempty"`
}

func (x *TimeFilter) Reset() {
	*x = TimeFilter{}
	mi := &file_stats_v1_stats_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeFilter) ProtoMessage() {}

func (x *TimeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_stats_v1_stats_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeFilter.ProtoReflect.Descriptor instead.
func (*TimeFilter) Descriptor() ([]byte, []int) {
	return file_stats_v1_stats_proto_rawDescGZIP(), []int{6}
}

func (x *TimeFilter) GetTimeFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeFrom
	}
	return nil
}

func (x *TimeFilter) GetTimeTo() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeTo
	}
	return nil
}

type GetPlantsStatsV1Response_StatsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Species string                 `protobuf:"bytes,1,opt,name=species,proto3" json:"species,omitempty"`
	Count   int64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Date    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetPlantsStatsV1Response_StatsInfo) Reset() {
	*x = GetPlantsStatsV1Response_StatsInfo{}
	mi := &file_stats_v1_stats_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPlantsStatsV1Response_StatsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlantsStatsV1Response_StatsInfo) ProtoMessage() {}

func (x *GetPlantsStatsV1Response_StatsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_stats_v1_stats_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlantsStatsV1Response_StatsInfo.ProtoReflect.Descriptor instead.
func (*GetPlantsStatsV1Response_StatsInfo) Descriptor() ([]byte, []int) {
	return file_stats_v1_stats_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetPlantsStatsV1Response_StatsInfo) GetSpecies() string {
	if x != nil {
		return x.Species
	}
	return ""
}

func (x *GetPlantsStatsV1Response_StatsInfo) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetPlantsStatsV1Response_StatsInfo) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type GetBuyStatsV1Response_StatsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Count int64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetBuyStatsV1Response_StatsInfo) Reset() {
	*x = GetBuyStatsV1Response_StatsInfo{}
	mi := &file_stats_v1_stats_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBuyStatsV1Response_StatsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBuyStatsV1Response_StatsInfo) ProtoMessage() {}

func (x *GetBuyStatsV1Response_StatsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_stats_v1_stats_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBuyStatsV1Response_StatsInfo.ProtoReflect.Descriptor instead.
func (*GetBuyStatsV1Response_StatsInfo) Descriptor() ([]byte, []int) {
	return file_stats_v1_stats_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetBuyStatsV1Response_StatsInfo) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *GetBuyStatsV1Response_StatsInfo) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetTradeStatsV1Response_StatsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Count  int64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Status int64                  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetTradeStatsV1Response_StatsInfo) Reset() {
	*x = GetTradeStatsV1Response_StatsInfo{}
	mi := &file_stats_v1_stats_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTradeStatsV1Response_StatsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTradeStatsV1Response_StatsInfo) ProtoMessage() {}

func (x *GetTradeStatsV1Response_StatsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_stats_v1_stats_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTradeStatsV1Response_StatsInfo.ProtoReflect.Descriptor instead.
func (*GetTradeStatsV1Response_StatsInfo) Descriptor() ([]byte, []int) {
	return file_stats_v1_stats_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetTradeStatsV1Response_StatsInfo) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *GetTradeStatsV1Response_StatsInfo) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetTradeStatsV1Response_StatsInfo) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_stats_v1_stats_proto protoreflect.FileDescriptor

var file_stats_v1_stats_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x47, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xe1, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x1a,
	0x6b, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x44, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x42, 0x75, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x22, 0xc1, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x75, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x75, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x1a, 0x51, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xdd,
	0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x41, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x1a, 0x69, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x78,
	0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x32, 0x0a, 0x06, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x06, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x32, 0xe5, 0x02, 0x0a, 0x08, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x41, 0x50, 0x49, 0x12, 0x77, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x31, 0x12, 0x21, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x6b,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x75, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x31, 0x12,
	0x1e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x62, 0x75, 0x79, 0x12, 0x73, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x31, 0x12, 0x20,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x42, 0x13, 0x5a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stats_v1_stats_proto_rawDescOnce sync.Once
	file_stats_v1_stats_proto_rawDescData = file_stats_v1_stats_proto_rawDesc
)

func file_stats_v1_stats_proto_rawDescGZIP() []byte {
	file_stats_v1_stats_proto_rawDescOnce.Do(func() {
		file_stats_v1_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_stats_v1_stats_proto_rawDescData)
	})
	return file_stats_v1_stats_proto_rawDescData
}

var file_stats_v1_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_stats_v1_stats_proto_goTypes = []any{
	(*GetPlantsStatsV1Request)(nil),            // 0: stats.v1.GetPlantsStatsV1Request
	(*GetPlantsStatsV1Response)(nil),           // 1: stats.v1.GetPlantsStatsV1Response
	(*GetBuyStatsV1Request)(nil),               // 2: stats.v1.GetBuyStatsV1Request
	(*GetBuyStatsV1Response)(nil),              // 3: stats.v1.GetBuyStatsV1Response
	(*GetTradeStatsV1Request)(nil),             // 4: stats.v1.GetTradeStatsV1Request
	(*GetTradeStatsV1Response)(nil),            // 5: stats.v1.GetTradeStatsV1Response
	(*TimeFilter)(nil),                         // 6: stats.v1.TimeFilter
	(*GetPlantsStatsV1Response_StatsInfo)(nil), // 7: stats.v1.GetPlantsStatsV1Response.StatsInfo
	(*GetBuyStatsV1Response_StatsInfo)(nil),    // 8: stats.v1.GetBuyStatsV1Response.StatsInfo
	(*GetTradeStatsV1Response_StatsInfo)(nil),  // 9: stats.v1.GetTradeStatsV1Response.StatsInfo
	(*timestamppb.Timestamp)(nil),              // 10: google.protobuf.Timestamp
}
var file_stats_v1_stats_proto_depIdxs = []int32{
	6,  // 0: stats.v1.GetPlantsStatsV1Request.filter:type_name -> stats.v1.TimeFilter
	7,  // 1: stats.v1.GetPlantsStatsV1Response.stats:type_name -> stats.v1.GetPlantsStatsV1Response.StatsInfo
	6,  // 2: stats.v1.GetBuyStatsV1Request.filter:type_name -> stats.v1.TimeFilter
	8,  // 3: stats.v1.GetBuyStatsV1Response.stats:type_name -> stats.v1.GetBuyStatsV1Response.StatsInfo
	6,  // 4: stats.v1.GetTradeStatsV1Request.filter:type_name -> stats.v1.TimeFilter
	9,  // 5: stats.v1.GetTradeStatsV1Response.stats:type_name -> stats.v1.GetTradeStatsV1Response.StatsInfo
	10, // 6: stats.v1.TimeFilter.timeFrom:type_name -> google.protobuf.Timestamp
	10, // 7: stats.v1.TimeFilter.timeTo:type_name -> google.protobuf.Timestamp
	10, // 8: stats.v1.GetPlantsStatsV1Response.StatsInfo.date:type_name -> google.protobuf.Timestamp
	10, // 9: stats.v1.GetBuyStatsV1Response.StatsInfo.date:type_name -> google.protobuf.Timestamp
	10, // 10: stats.v1.GetTradeStatsV1Response.StatsInfo.date:type_name -> google.protobuf.Timestamp
	0,  // 11: stats.v1.StatsAPI.GetPlantsStatsV1:input_type -> stats.v1.GetPlantsStatsV1Request
	2,  // 12: stats.v1.StatsAPI.GetBuyStatsV1:input_type -> stats.v1.GetBuyStatsV1Request
	4,  // 13: stats.v1.StatsAPI.GetTradeStatsV1:input_type -> stats.v1.GetTradeStatsV1Request
	1,  // 14: stats.v1.StatsAPI.GetPlantsStatsV1:output_type -> stats.v1.GetPlantsStatsV1Response
	3,  // 15: stats.v1.StatsAPI.GetBuyStatsV1:output_type -> stats.v1.GetBuyStatsV1Response
	5,  // 16: stats.v1.StatsAPI.GetTradeStatsV1:output_type -> stats.v1.GetTradeStatsV1Response
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_stats_v1_stats_proto_init() }
func file_stats_v1_stats_proto_init() {
	if File_stats_v1_stats_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stats_v1_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stats_v1_stats_proto_goTypes,
		DependencyIndexes: file_stats_v1_stats_proto_depIdxs,
		MessageInfos:      file_stats_v1_stats_proto_msgTypes,
	}.Build()
	File_stats_v1_stats_proto = out.File
	file_stats_v1_stats_proto_rawDesc = nil
	file_stats_v1_stats_proto_goTypes = nil
	file_stats_v1_stats_proto_depIdxs = nil
}
