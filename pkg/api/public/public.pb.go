// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.1
// source: public.proto

package public

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

type AddPlantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AddPlantRequest) Reset() {
	*x = AddPlantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPlantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPlantRequest) ProtoMessage() {}

func (x *AddPlantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPlantRequest.ProtoReflect.Descriptor instead.
func (*AddPlantRequest) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{0}
}

func (x *AddPlantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddPlantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plant *Plant `protobuf:"bytes,1,opt,name=plant,proto3" json:"plant,omitempty"`
}

func (x *AddPlantResponse) Reset() {
	*x = AddPlantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPlantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPlantResponse) ProtoMessage() {}

func (x *AddPlantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPlantResponse.ProtoReflect.Descriptor instead.
func (*AddPlantResponse) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{1}
}

func (x *AddPlantResponse) GetPlant() *Plant {
	if x != nil {
		return x.Plant
	}
	return nil
}

type UpdatePlantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldName string `protobuf:"bytes,1,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	NewName string `protobuf:"bytes,2,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
}

func (x *UpdatePlantRequest) Reset() {
	*x = UpdatePlantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlantRequest) ProtoMessage() {}

func (x *UpdatePlantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlantRequest.ProtoReflect.Descriptor instead.
func (*UpdatePlantRequest) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePlantRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *UpdatePlantRequest) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

type UpdatePlantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plant *Plant `protobuf:"bytes,1,opt,name=plant,proto3" json:"plant,omitempty"`
}

func (x *UpdatePlantResponse) Reset() {
	*x = UpdatePlantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlantResponse) ProtoMessage() {}

func (x *UpdatePlantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlantResponse.ProtoReflect.Descriptor instead.
func (*UpdatePlantResponse) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePlantResponse) GetPlant() *Plant {
	if x != nil {
		return x.Plant
	}
	return nil
}

type GetPlantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPlantRequest) Reset() {
	*x = GetPlantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlantRequest) ProtoMessage() {}

func (x *GetPlantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlantRequest.ProtoReflect.Descriptor instead.
func (*GetPlantRequest) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{4}
}

func (x *GetPlantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetPlantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plant *Plant `protobuf:"bytes,1,opt,name=plant,proto3" json:"plant,omitempty"`
}

func (x *GetPlantResponse) Reset() {
	*x = GetPlantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlantResponse) ProtoMessage() {}

func (x *GetPlantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlantResponse.ProtoReflect.Descriptor instead.
func (*GetPlantResponse) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{5}
}

func (x *GetPlantResponse) GetPlant() *Plant {
	if x != nil {
		return x.Plant
	}
	return nil
}

type GetPlantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPlantsRequest) Reset() {
	*x = GetPlantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlantsRequest) ProtoMessage() {}

func (x *GetPlantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlantsRequest.ProtoReflect.Descriptor instead.
func (*GetPlantsRequest) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{6}
}

type GetPlantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plants []*Plant `protobuf:"bytes,1,rep,name=plants,proto3" json:"plants,omitempty"`
}

func (x *GetPlantsResponse) Reset() {
	*x = GetPlantsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlantsResponse) ProtoMessage() {}

func (x *GetPlantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlantsResponse.ProtoReflect.Descriptor instead.
func (*GetPlantsResponse) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{7}
}

func (x *GetPlantsResponse) GetPlants() []*Plant {
	if x != nil {
		return x.Plants
	}
	return nil
}

type DeletePlantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeletePlantRequest) Reset() {
	*x = DeletePlantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePlantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePlantRequest) ProtoMessage() {}

func (x *DeletePlantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePlantRequest.ProtoReflect.Descriptor instead.
func (*DeletePlantRequest) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePlantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeletePlantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePlantResponse) Reset() {
	*x = DeletePlantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePlantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePlantResponse) ProtoMessage() {}

func (x *DeletePlantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePlantResponse.ProtoReflect.Descriptor instead.
func (*DeletePlantResponse) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{9}
}

type Plant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Plant) Reset() {
	*x = Plant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plant) ProtoMessage() {}

func (x *Plant) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plant.ProtoReflect.Descriptor instead.
func (*Plant) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{10}
}

func (x *Plant) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Plant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_public_proto protoreflect.FileDescriptor

var file_public_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x67, 0x6f, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73,
	0x22, 0x25, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e,
	0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e,
	0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x05, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52,
	0x05, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73,
	0x22, 0x28, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2b, 0x0a, 0x05, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xbf,
	0x03, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x41, 0x70, 0x69, 0x12, 0x51, 0x0a, 0x08,
	0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6f,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x41,
	0x64, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5a, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x12, 0x24,
	0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6f, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x67, 0x6f,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x6f, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x74, 0x72, 0x6f, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x3b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_proto_rawDescOnce sync.Once
	file_public_proto_rawDescData = file_public_proto_rawDesc
)

func file_public_proto_rawDescGZIP() []byte {
	file_public_proto_rawDescOnce.Do(func() {
		file_public_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_proto_rawDescData)
	})
	return file_public_proto_rawDescData
}

var file_public_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_public_proto_goTypes = []interface{}{
	(*AddPlantRequest)(nil),     // 0: go.public.plants.AddPlantRequest
	(*AddPlantResponse)(nil),    // 1: go.public.plants.AddPlantResponse
	(*UpdatePlantRequest)(nil),  // 2: go.public.plants.UpdatePlantRequest
	(*UpdatePlantResponse)(nil), // 3: go.public.plants.UpdatePlantResponse
	(*GetPlantRequest)(nil),     // 4: go.public.plants.GetPlantRequest
	(*GetPlantResponse)(nil),    // 5: go.public.plants.GetPlantResponse
	(*GetPlantsRequest)(nil),    // 6: go.public.plants.GetPlantsRequest
	(*GetPlantsResponse)(nil),   // 7: go.public.plants.GetPlantsResponse
	(*DeletePlantRequest)(nil),  // 8: go.public.plants.DeletePlantRequest
	(*DeletePlantResponse)(nil), // 9: go.public.plants.DeletePlantResponse
	(*Plant)(nil),               // 10: go.public.plants.Plant
}
var file_public_proto_depIdxs = []int32{
	10, // 0: go.public.plants.AddPlantResponse.plant:type_name -> go.public.plants.Plant
	10, // 1: go.public.plants.UpdatePlantResponse.plant:type_name -> go.public.plants.Plant
	10, // 2: go.public.plants.GetPlantResponse.plant:type_name -> go.public.plants.Plant
	10, // 3: go.public.plants.GetPlantsResponse.plants:type_name -> go.public.plants.Plant
	0,  // 4: go.public.plants.PlantsApi.AddPlant:input_type -> go.public.plants.AddPlantRequest
	2,  // 5: go.public.plants.PlantsApi.UpdatePlant:input_type -> go.public.plants.UpdatePlantRequest
	4,  // 6: go.public.plants.PlantsApi.GetPlant:input_type -> go.public.plants.GetPlantRequest
	6,  // 7: go.public.plants.PlantsApi.GetPlants:input_type -> go.public.plants.GetPlantsRequest
	8,  // 8: go.public.plants.PlantsApi.DeletePlant:input_type -> go.public.plants.DeletePlantRequest
	1,  // 9: go.public.plants.PlantsApi.AddPlant:output_type -> go.public.plants.AddPlantResponse
	3,  // 10: go.public.plants.PlantsApi.UpdatePlant:output_type -> go.public.plants.UpdatePlantResponse
	5,  // 11: go.public.plants.PlantsApi.GetPlant:output_type -> go.public.plants.GetPlantResponse
	7,  // 12: go.public.plants.PlantsApi.GetPlants:output_type -> go.public.plants.GetPlantsResponse
	9,  // 13: go.public.plants.PlantsApi.DeletePlant:output_type -> go.public.plants.DeletePlantResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_public_proto_init() }
func file_public_proto_init() {
	if File_public_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPlantRequest); i {
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
		file_public_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPlantResponse); i {
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
		file_public_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlantRequest); i {
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
		file_public_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlantResponse); i {
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
		file_public_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlantRequest); i {
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
		file_public_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlantResponse); i {
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
		file_public_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlantsRequest); i {
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
		file_public_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlantsResponse); i {
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
		file_public_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePlantRequest); i {
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
		file_public_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePlantResponse); i {
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
		file_public_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plant); i {
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
			RawDescriptor: file_public_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_public_proto_goTypes,
		DependencyIndexes: file_public_proto_depIdxs,
		MessageInfos:      file_public_proto_msgTypes,
	}.Build()
	File_public_proto = out.File
	file_public_proto_rawDesc = nil
	file_public_proto_goTypes = nil
	file_public_proto_depIdxs = nil
}
