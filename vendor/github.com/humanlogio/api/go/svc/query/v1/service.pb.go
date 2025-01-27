// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: svc/query/v1/service.proto

package queryv1

import (
	v1 "github.com/humanlogio/api/go/types/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type SummarizeEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId   int64                  `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	From        *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=from,proto3,oneof" json:"from,omitempty"`
	To          *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=to,proto3,oneof" json:"to,omitempty"`
	BucketCount uint32                 `protobuf:"varint,4,opt,name=bucket_count,json=bucketCount,proto3" json:"bucket_count,omitempty"`
}

func (x *SummarizeEventsRequest) Reset() {
	*x = SummarizeEventsRequest{}
	mi := &file_svc_query_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SummarizeEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummarizeEventsRequest) ProtoMessage() {}

func (x *SummarizeEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_query_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummarizeEventsRequest.ProtoReflect.Descriptor instead.
func (*SummarizeEventsRequest) Descriptor() ([]byte, []int) {
	return file_svc_query_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *SummarizeEventsRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *SummarizeEventsRequest) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *SummarizeEventsRequest) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *SummarizeEventsRequest) GetBucketCount() uint32 {
	if x != nil {
		return x.BucketCount
	}
	return 0
}

type SummarizeEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketWidth *durationpb.Duration              `protobuf:"bytes,1,opt,name=bucket_width,json=bucketWidth,proto3" json:"bucket_width,omitempty"`
	Buckets     []*SummarizeEventsResponse_Bucket `protobuf:"bytes,2,rep,name=buckets,proto3" json:"buckets,omitempty"`
}

func (x *SummarizeEventsResponse) Reset() {
	*x = SummarizeEventsResponse{}
	mi := &file_svc_query_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SummarizeEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummarizeEventsResponse) ProtoMessage() {}

func (x *SummarizeEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_query_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummarizeEventsResponse.ProtoReflect.Descriptor instead.
func (*SummarizeEventsResponse) Descriptor() ([]byte, []int) {
	return file_svc_query_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *SummarizeEventsResponse) GetBucketWidth() *durationpb.Duration {
	if x != nil {
		return x.BucketWidth
	}
	return nil
}

func (x *SummarizeEventsResponse) GetBuckets() []*SummarizeEventsResponse_Bucket {
	if x != nil {
		return x.Buckets
	}
	return nil
}

type WatchQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId int64        `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Query     *v1.LogQuery `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *WatchQueryRequest) Reset() {
	*x = WatchQueryRequest{}
	mi := &file_svc_query_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchQueryRequest) ProtoMessage() {}

func (x *WatchQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_query_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchQueryRequest.ProtoReflect.Descriptor instead.
func (*WatchQueryRequest) Descriptor() ([]byte, []int) {
	return file_svc_query_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *WatchQueryRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *WatchQueryRequest) GetQuery() *v1.LogQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

type WatchQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*v1.LogEventGroup `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *WatchQueryResponse) Reset() {
	*x = WatchQueryResponse{}
	mi := &file_svc_query_v1_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchQueryResponse) ProtoMessage() {}

func (x *WatchQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_query_v1_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchQueryResponse.ProtoReflect.Descriptor instead.
func (*WatchQueryResponse) Descriptor() ([]byte, []int) {
	return file_svc_query_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *WatchQueryResponse) GetEvents() []*v1.LogEventGroup {
	if x != nil {
		return x.Events
	}
	return nil
}

type SummarizeEventsResponse_Bucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	EventCount uint64                 `protobuf:"varint,2,opt,name=event_count,json=eventCount,proto3" json:"event_count,omitempty"`
}

func (x *SummarizeEventsResponse_Bucket) Reset() {
	*x = SummarizeEventsResponse_Bucket{}
	mi := &file_svc_query_v1_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SummarizeEventsResponse_Bucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummarizeEventsResponse_Bucket) ProtoMessage() {}

func (x *SummarizeEventsResponse_Bucket) ProtoReflect() protoreflect.Message {
	mi := &file_svc_query_v1_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummarizeEventsResponse_Bucket.ProtoReflect.Descriptor instead.
func (*SummarizeEventsResponse_Bucket) Descriptor() ([]byte, []int) {
	return file_svc_query_v1_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SummarizeEventsResponse_Bucket) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *SummarizeEventsResponse_Bucket) GetEventCount() uint64 {
	if x != nil {
		return x.EventCount
	}
	return 0
}

var File_svc_query_v1_service_proto protoreflect.FileDescriptor

var file_svc_query_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x76, 0x63, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x76,
	0x63, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x16, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x69, 0x7a, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x33, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x02,
	0x74, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x66, 0x72, 0x6f,
	0x6d, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x74, 0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x17, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x57, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x46, 0x0a, 0x07, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x07, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x1a, 0x55, 0x0a, 0x06, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x5c, 0x0a, 0x11, 0x57, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22,
	0x45, 0x0a, 0x12, 0x57, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xc1, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0f, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x69, 0x7a, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x69, 0x7a, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0xa5, 0x01, 0x0a, 0x10, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x42,
	0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x6d, 0x61,
	0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x76,
	0x63, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x51, 0x58, 0xaa, 0x02, 0x0c, 0x53, 0x76, 0x63, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x53, 0x76, 0x63, 0x5c, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x53, 0x76, 0x63, 0x5c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0e, 0x53, 0x76, 0x63, 0x3a, 0x3a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_query_v1_service_proto_rawDescOnce sync.Once
	file_svc_query_v1_service_proto_rawDescData = file_svc_query_v1_service_proto_rawDesc
)

func file_svc_query_v1_service_proto_rawDescGZIP() []byte {
	file_svc_query_v1_service_proto_rawDescOnce.Do(func() {
		file_svc_query_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_query_v1_service_proto_rawDescData)
	})
	return file_svc_query_v1_service_proto_rawDescData
}

var file_svc_query_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_svc_query_v1_service_proto_goTypes = []any{
	(*SummarizeEventsRequest)(nil),         // 0: svc.query.v1.SummarizeEventsRequest
	(*SummarizeEventsResponse)(nil),        // 1: svc.query.v1.SummarizeEventsResponse
	(*WatchQueryRequest)(nil),              // 2: svc.query.v1.WatchQueryRequest
	(*WatchQueryResponse)(nil),             // 3: svc.query.v1.WatchQueryResponse
	(*SummarizeEventsResponse_Bucket)(nil), // 4: svc.query.v1.SummarizeEventsResponse.Bucket
	(*timestamppb.Timestamp)(nil),          // 5: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),            // 6: google.protobuf.Duration
	(*v1.LogQuery)(nil),                    // 7: types.v1.LogQuery
	(*v1.LogEventGroup)(nil),               // 8: types.v1.LogEventGroup
}
var file_svc_query_v1_service_proto_depIdxs = []int32{
	5, // 0: svc.query.v1.SummarizeEventsRequest.from:type_name -> google.protobuf.Timestamp
	5, // 1: svc.query.v1.SummarizeEventsRequest.to:type_name -> google.protobuf.Timestamp
	6, // 2: svc.query.v1.SummarizeEventsResponse.bucket_width:type_name -> google.protobuf.Duration
	4, // 3: svc.query.v1.SummarizeEventsResponse.buckets:type_name -> svc.query.v1.SummarizeEventsResponse.Bucket
	7, // 4: svc.query.v1.WatchQueryRequest.query:type_name -> types.v1.LogQuery
	8, // 5: svc.query.v1.WatchQueryResponse.events:type_name -> types.v1.LogEventGroup
	5, // 6: svc.query.v1.SummarizeEventsResponse.Bucket.ts:type_name -> google.protobuf.Timestamp
	0, // 7: svc.query.v1.QueryService.SummarizeEvents:input_type -> svc.query.v1.SummarizeEventsRequest
	2, // 8: svc.query.v1.QueryService.WatchQuery:input_type -> svc.query.v1.WatchQueryRequest
	1, // 9: svc.query.v1.QueryService.SummarizeEvents:output_type -> svc.query.v1.SummarizeEventsResponse
	3, // 10: svc.query.v1.QueryService.WatchQuery:output_type -> svc.query.v1.WatchQueryResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_svc_query_v1_service_proto_init() }
func file_svc_query_v1_service_proto_init() {
	if File_svc_query_v1_service_proto != nil {
		return
	}
	file_svc_query_v1_service_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_query_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_query_v1_service_proto_goTypes,
		DependencyIndexes: file_svc_query_v1_service_proto_depIdxs,
		MessageInfos:      file_svc_query_v1_service_proto_msgTypes,
	}.Build()
	File_svc_query_v1_service_proto = out.File
	file_svc_query_v1_service_proto_rawDesc = nil
	file_svc_query_v1_service_proto_goTypes = nil
	file_svc_query_v1_service_proto_depIdxs = nil
}
