// Code generated with goa v3.11.3, DO NOT EDIT.
//
// sommelier protocol buffer definition
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: goadesign_goagen_sommelier.proto

package sommelierpb

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

type PickRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of bottle to pick
	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Varietals in preference order
	Varietal []string `protobuf:"bytes,2,rep,name=varietal,proto3" json:"varietal,omitempty"`
	// Winery of bottle to pick
	Winery *string `protobuf:"bytes,3,opt,name=winery,proto3,oneof" json:"winery,omitempty"`
}

func (x *PickRequest) Reset() {
	*x = PickRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_sommelier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PickRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PickRequest) ProtoMessage() {}

func (x *PickRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_sommelier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PickRequest.ProtoReflect.Descriptor instead.
func (*PickRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_sommelier_proto_rawDescGZIP(), []int{0}
}

func (x *PickRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *PickRequest) GetVarietal() []string {
	if x != nil {
		return x.Varietal
	}
	return nil
}

func (x *PickRequest) GetWinery() string {
	if x != nil && x.Winery != nil {
		return *x.Winery
	}
	return ""
}

type StoredBottleCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field []*StoredBottle `protobuf:"bytes,1,rep,name=field,proto3" json:"field,omitempty"`
}

func (x *StoredBottleCollection) Reset() {
	*x = StoredBottleCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_sommelier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredBottleCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredBottleCollection) ProtoMessage() {}

func (x *StoredBottleCollection) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_sommelier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredBottleCollection.ProtoReflect.Descriptor instead.
func (*StoredBottleCollection) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_sommelier_proto_rawDescGZIP(), []int{1}
}

func (x *StoredBottleCollection) GetField() []*StoredBottle {
	if x != nil {
		return x.Field
	}
	return nil
}

// A StoredBottle describes a bottle retrieved by the storage service.
type StoredBottle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique id of the bottle.
	Id string `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	// Name of bottle
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Winery that produces wine
	Winery *Winery `protobuf:"bytes,3,opt,name=winery,proto3" json:"winery,omitempty"`
	// Vintage of bottle
	Vintage uint32 `protobuf:"varint,4,opt,name=vintage,proto3" json:"vintage,omitempty"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*Component `protobuf:"bytes,5,rep,name=composition,proto3" json:"composition,omitempty"`
	// Description of bottle
	Description *string `protobuf:"bytes,6,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `protobuf:"varint,7,opt,name=rating,proto3,oneof" json:"rating,omitempty"`
}

func (x *StoredBottle) Reset() {
	*x = StoredBottle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_sommelier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredBottle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredBottle) ProtoMessage() {}

func (x *StoredBottle) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_sommelier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredBottle.ProtoReflect.Descriptor instead.
func (*StoredBottle) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_sommelier_proto_rawDescGZIP(), []int{2}
}

func (x *StoredBottle) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StoredBottle) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoredBottle) GetWinery() *Winery {
	if x != nil {
		return x.Winery
	}
	return nil
}

func (x *StoredBottle) GetVintage() uint32 {
	if x != nil {
		return x.Vintage
	}
	return 0
}

func (x *StoredBottle) GetComposition() []*Component {
	if x != nil {
		return x.Composition
	}
	return nil
}

func (x *StoredBottle) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *StoredBottle) GetRating() uint32 {
	if x != nil && x.Rating != nil {
		return *x.Rating
	}
	return 0
}

type Winery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of winery
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Region of winery
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// Country of winery
	Country string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	// Winery website URL
	Url *string `protobuf:"bytes,4,opt,name=url,proto3,oneof" json:"url,omitempty"`
}

func (x *Winery) Reset() {
	*x = Winery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_sommelier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Winery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Winery) ProtoMessage() {}

func (x *Winery) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_sommelier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Winery.ProtoReflect.Descriptor instead.
func (*Winery) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_sommelier_proto_rawDescGZIP(), []int{3}
}

func (x *Winery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Winery) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Winery) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Winery) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

type Component struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Grape varietal
	Varietal string `protobuf:"bytes,1,opt,name=varietal,proto3" json:"varietal,omitempty"`
	// Percentage of varietal in wine
	Percentage *uint32 `protobuf:"varint,2,opt,name=percentage,proto3,oneof" json:"percentage,omitempty"`
}

func (x *Component) Reset() {
	*x = Component{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_sommelier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Component) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Component) ProtoMessage() {}

func (x *Component) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_sommelier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Component.ProtoReflect.Descriptor instead.
func (*Component) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_sommelier_proto_rawDescGZIP(), []int{4}
}

func (x *Component) GetVarietal() string {
	if x != nil {
		return x.Varietal
	}
	return ""
}

func (x *Component) GetPercentage() uint32 {
	if x != nil && x.Percentage != nil {
		return *x.Percentage
	}
	return 0
}

var File_goadesign_goagen_sommelier_proto protoreflect.FileDescriptor

var file_goadesign_goagen_sommelier_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x6c, 0x69, 0x65, 0x72, 0x22, 0x73, 0x0a,
	0x0b, 0x50, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x61,
	0x6c, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x77, 0x69, 0x6e, 0x65,
	0x72, 0x79, 0x22, 0x47, 0x0a, 0x16, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x42, 0x6f, 0x74, 0x74,
	0x6c, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x6f,
	0x6d, 0x6d, 0x65, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x42, 0x6f,
	0x74, 0x74, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x8e, 0x02, 0x0a, 0x0c,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x42, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x29, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x57, 0x69, 0x6e,
	0x65, 0x72, 0x79, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x69,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x6f, 0x6d,
	0x6d, 0x65, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x88, 0x01,
	0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x6d, 0x0a, 0x06,
	0x57, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x15, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x5b, 0x0a, 0x09, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69,
	0x65, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69,
	0x65, 0x74, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x32, 0x4e, 0x0a, 0x09, 0x53, 0x6f, 0x6d, 0x6d,
	0x65, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x04, 0x50, 0x69, 0x63, 0x6b, 0x12, 0x16, 0x2e,
	0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x6c, 0x69, 0x65,
	0x72, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x42, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x73, 0x6f, 0x6d,
	0x6d, 0x65, 0x6c, 0x69, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goadesign_goagen_sommelier_proto_rawDescOnce sync.Once
	file_goadesign_goagen_sommelier_proto_rawDescData = file_goadesign_goagen_sommelier_proto_rawDesc
)

func file_goadesign_goagen_sommelier_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_sommelier_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_sommelier_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_sommelier_proto_rawDescData)
	})
	return file_goadesign_goagen_sommelier_proto_rawDescData
}

var file_goadesign_goagen_sommelier_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_goadesign_goagen_sommelier_proto_goTypes = []interface{}{
	(*PickRequest)(nil),            // 0: sommelier.PickRequest
	(*StoredBottleCollection)(nil), // 1: sommelier.StoredBottleCollection
	(*StoredBottle)(nil),           // 2: sommelier.StoredBottle
	(*Winery)(nil),                 // 3: sommelier.Winery
	(*Component)(nil),              // 4: sommelier.Component
}
var file_goadesign_goagen_sommelier_proto_depIdxs = []int32{
	2, // 0: sommelier.StoredBottleCollection.field:type_name -> sommelier.StoredBottle
	3, // 1: sommelier.StoredBottle.winery:type_name -> sommelier.Winery
	4, // 2: sommelier.StoredBottle.composition:type_name -> sommelier.Component
	0, // 3: sommelier.Sommelier.Pick:input_type -> sommelier.PickRequest
	1, // 4: sommelier.Sommelier.Pick:output_type -> sommelier.StoredBottleCollection
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_sommelier_proto_init() }
func file_goadesign_goagen_sommelier_proto_init() {
	if File_goadesign_goagen_sommelier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_sommelier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PickRequest); i {
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
		file_goadesign_goagen_sommelier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredBottleCollection); i {
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
		file_goadesign_goagen_sommelier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredBottle); i {
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
		file_goadesign_goagen_sommelier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Winery); i {
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
		file_goadesign_goagen_sommelier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Component); i {
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
	file_goadesign_goagen_sommelier_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_goadesign_goagen_sommelier_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_goadesign_goagen_sommelier_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_goadesign_goagen_sommelier_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goadesign_goagen_sommelier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_sommelier_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_sommelier_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_sommelier_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_sommelier_proto = out.File
	file_goadesign_goagen_sommelier_proto_rawDesc = nil
	file_goadesign_goagen_sommelier_proto_goTypes = nil
	file_goadesign_goagen_sommelier_proto_depIdxs = nil
}
