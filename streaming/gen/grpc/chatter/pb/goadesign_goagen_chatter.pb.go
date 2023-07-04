// Code generated with goa v3.12.1, DO NOT EDIT.
//
// chatter protocol buffer definition
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: goadesign_goagen_chatter.proto

package chatterpb

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_chatter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_chatter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_chatter_proto_rawDescGZIP(), []int{0}
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_chatter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_chatter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_chatter_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type EchoerStreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *EchoerStreamingRequest) Reset() {
	*x = EchoerStreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_chatter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoerStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoerStreamingRequest) ProtoMessage() {}

func (x *EchoerStreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_chatter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoerStreamingRequest.ProtoReflect.Descriptor instead.
func (*EchoerStreamingRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_chatter_proto_rawDescGZIP(), []int{2}
}

func (x *EchoerStreamingRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type EchoerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *EchoerResponse) Reset() {
	*x = EchoerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_chatter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoerResponse) ProtoMessage() {}

func (x *EchoerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_chatter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoerResponse.ProtoReflect.Descriptor instead.
func (*EchoerResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_chatter_proto_rawDescGZIP(), []int{3}
}

func (x *EchoerResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type ListenerStreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *ListenerStreamingRequest) Reset() {
	*x = ListenerStreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_chatter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenerStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenerStreamingRequest) ProtoMessage() {}

func (x *ListenerStreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_chatter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenerStreamingRequest.ProtoReflect.Descriptor instead.
func (*ListenerStreamingRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_chatter_proto_rawDescGZIP(), []int{4}
}

func (x *ListenerStreamingRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type ListenerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListenerResponse) Reset() {
	*x = ListenerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_chatter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenerResponse) ProtoMessage() {}

func (x *ListenerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_chatter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenerResponse.ProtoReflect.Descriptor instead.
func (*ListenerResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_chatter_proto_rawDescGZIP(), []int{5}
}

type SummaryStreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *SummaryStreamingRequest) Reset() {
	*x = SummaryStreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_chatter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryStreamingRequest) ProtoMessage() {}

func (x *SummaryStreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_chatter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryStreamingRequest.ProtoReflect.Descriptor instead.
func (*SummaryStreamingRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_chatter_proto_rawDescGZIP(), []int{6}
}

func (x *SummaryStreamingRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type ChatSummaryCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field []*ChatSummary `protobuf:"bytes,1,rep,name=field,proto3" json:"field,omitempty"`
}

func (x *ChatSummaryCollection) Reset() {
	*x = ChatSummaryCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_chatter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatSummaryCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatSummaryCollection) ProtoMessage() {}

func (x *ChatSummaryCollection) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_chatter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatSummaryCollection.ProtoReflect.Descriptor instead.
func (*ChatSummaryCollection) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_chatter_proto_rawDescGZIP(), []int{7}
}

func (x *ChatSummaryCollection) GetField() []*ChatSummary {
	if x != nil {
		return x.Field
	}
	return nil
}

type ChatSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message sent to the server
	Message_ string `protobuf:"bytes,1,opt,name=message_,json=message,proto3" json:"message_,omitempty"`
	// Length of the message sent
	Length *int32 `protobuf:"zigzag32,2,opt,name=length,proto3,oneof" json:"length,omitempty"`
	// Time at which the message was sent
	SentAt string `protobuf:"bytes,3,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *ChatSummary) Reset() {
	*x = ChatSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_chatter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatSummary) ProtoMessage() {}

func (x *ChatSummary) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_chatter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatSummary.ProtoReflect.Descriptor instead.
func (*ChatSummary) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_chatter_proto_rawDescGZIP(), []int{8}
}

func (x *ChatSummary) GetMessage_() string {
	if x != nil {
		return x.Message_
	}
	return ""
}

func (x *ChatSummary) GetLength() int32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

func (x *ChatSummary) GetSentAt() string {
	if x != nil {
		return x.SentAt
	}
	return ""
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_chatter_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_chatter_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_chatter_proto_rawDescGZIP(), []int{9}
}

type SubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message sent to the server
	Message_ string `protobuf:"bytes,1,opt,name=message_,json=message,proto3" json:"message_,omitempty"`
	Action   string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	// Time at which the message was added
	AddedAt string `protobuf:"bytes,3,opt,name=added_at,json=addedAt,proto3" json:"added_at,omitempty"`
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_chatter_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_chatter_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_chatter_proto_rawDescGZIP(), []int{10}
}

func (x *SubscribeResponse) GetMessage_() string {
	if x != nil {
		return x.Message_
	}
	return ""
}

func (x *SubscribeResponse) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *SubscribeResponse) GetAddedAt() string {
	if x != nil {
		return x.AddedAt
	}
	return ""
}

type HistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HistoryRequest) Reset() {
	*x = HistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_chatter_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryRequest) ProtoMessage() {}

func (x *HistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_chatter_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryRequest.ProtoReflect.Descriptor instead.
func (*HistoryRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_chatter_proto_rawDescGZIP(), []int{11}
}

type HistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message sent to the server
	Message_ string `protobuf:"bytes,1,opt,name=message_,json=message,proto3" json:"message_,omitempty"`
	// Length of the message sent
	Length *int32 `protobuf:"zigzag32,2,opt,name=length,proto3,oneof" json:"length,omitempty"`
	// Time at which the message was sent
	SentAt string `protobuf:"bytes,3,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *HistoryResponse) Reset() {
	*x = HistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_chatter_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryResponse) ProtoMessage() {}

func (x *HistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_chatter_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryResponse.ProtoReflect.Descriptor instead.
func (*HistoryResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_chatter_proto_rawDescGZIP(), []int{12}
}

func (x *HistoryResponse) GetMessage_() string {
	if x != nil {
		return x.Message_
	}
	return ""
}

func (x *HistoryResponse) GetLength() int32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

func (x *HistoryResponse) GetSentAt() string {
	if x != nil {
		return x.SentAt
	}
	return ""
}

var File_goadesign_goagen_chatter_proto protoreflect.FileDescriptor

var file_goadesign_goagen_chatter_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x63, 0x68, 0x61, 0x74, 0x74, 0x65, 0x72, 0x22, 0x0e, 0x0a, 0x0c, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x0d, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x22, 0x2e, 0x0a, 0x16, 0x45, 0x63, 0x68, 0x6f, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x22, 0x26, 0x0a, 0x0e, 0x45, 0x63, 0x68, 0x6f, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x30, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f,
	0x0a, 0x17, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22,
	0x43, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x22, 0x69, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b,
	0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x48, 0x00,
	0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x6e, 0x74, 0x41, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22,
	0x12, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x61, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x41, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6d, 0x0a, 0x0f, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x48, 0x00, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x32, 0xaa, 0x03, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x45,
	0x63, 0x68, 0x6f, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e,
	0x45, 0x63, 0x68, 0x6f, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x2e, 0x45, 0x63, 0x68, 0x6f, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28,
	0x01, 0x30, 0x01, 0x12, 0x4a, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12,
	0x21, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12,
	0x4d, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x20, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x01, 0x12, 0x44,
	0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x19, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goadesign_goagen_chatter_proto_rawDescOnce sync.Once
	file_goadesign_goagen_chatter_proto_rawDescData = file_goadesign_goagen_chatter_proto_rawDesc
)

func file_goadesign_goagen_chatter_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_chatter_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_chatter_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_chatter_proto_rawDescData)
	})
	return file_goadesign_goagen_chatter_proto_rawDescData
}

var file_goadesign_goagen_chatter_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_goadesign_goagen_chatter_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),             // 0: chatter.LoginRequest
	(*LoginResponse)(nil),            // 1: chatter.LoginResponse
	(*EchoerStreamingRequest)(nil),   // 2: chatter.EchoerStreamingRequest
	(*EchoerResponse)(nil),           // 3: chatter.EchoerResponse
	(*ListenerStreamingRequest)(nil), // 4: chatter.ListenerStreamingRequest
	(*ListenerResponse)(nil),         // 5: chatter.ListenerResponse
	(*SummaryStreamingRequest)(nil),  // 6: chatter.SummaryStreamingRequest
	(*ChatSummaryCollection)(nil),    // 7: chatter.ChatSummaryCollection
	(*ChatSummary)(nil),              // 8: chatter.ChatSummary
	(*SubscribeRequest)(nil),         // 9: chatter.SubscribeRequest
	(*SubscribeResponse)(nil),        // 10: chatter.SubscribeResponse
	(*HistoryRequest)(nil),           // 11: chatter.HistoryRequest
	(*HistoryResponse)(nil),          // 12: chatter.HistoryResponse
}
var file_goadesign_goagen_chatter_proto_depIdxs = []int32{
	8,  // 0: chatter.ChatSummaryCollection.field:type_name -> chatter.ChatSummary
	0,  // 1: chatter.Chatter.Login:input_type -> chatter.LoginRequest
	2,  // 2: chatter.Chatter.Echoer:input_type -> chatter.EchoerStreamingRequest
	4,  // 3: chatter.Chatter.Listener:input_type -> chatter.ListenerStreamingRequest
	6,  // 4: chatter.Chatter.Summary:input_type -> chatter.SummaryStreamingRequest
	9,  // 5: chatter.Chatter.Subscribe:input_type -> chatter.SubscribeRequest
	11, // 6: chatter.Chatter.History:input_type -> chatter.HistoryRequest
	1,  // 7: chatter.Chatter.Login:output_type -> chatter.LoginResponse
	3,  // 8: chatter.Chatter.Echoer:output_type -> chatter.EchoerResponse
	5,  // 9: chatter.Chatter.Listener:output_type -> chatter.ListenerResponse
	7,  // 10: chatter.Chatter.Summary:output_type -> chatter.ChatSummaryCollection
	10, // 11: chatter.Chatter.Subscribe:output_type -> chatter.SubscribeResponse
	12, // 12: chatter.Chatter.History:output_type -> chatter.HistoryResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_chatter_proto_init() }
func file_goadesign_goagen_chatter_proto_init() {
	if File_goadesign_goagen_chatter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_chatter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_goadesign_goagen_chatter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_goadesign_goagen_chatter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoerStreamingRequest); i {
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
		file_goadesign_goagen_chatter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoerResponse); i {
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
		file_goadesign_goagen_chatter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenerStreamingRequest); i {
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
		file_goadesign_goagen_chatter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenerResponse); i {
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
		file_goadesign_goagen_chatter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryStreamingRequest); i {
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
		file_goadesign_goagen_chatter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatSummaryCollection); i {
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
		file_goadesign_goagen_chatter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatSummary); i {
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
		file_goadesign_goagen_chatter_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_goadesign_goagen_chatter_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeResponse); i {
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
		file_goadesign_goagen_chatter_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryRequest); i {
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
		file_goadesign_goagen_chatter_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryResponse); i {
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
	file_goadesign_goagen_chatter_proto_msgTypes[8].OneofWrappers = []interface{}{}
	file_goadesign_goagen_chatter_proto_msgTypes[12].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goadesign_goagen_chatter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_chatter_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_chatter_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_chatter_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_chatter_proto = out.File
	file_goadesign_goagen_chatter_proto_rawDesc = nil
	file_goadesign_goagen_chatter_proto_goTypes = nil
	file_goadesign_goagen_chatter_proto_depIdxs = nil
}
