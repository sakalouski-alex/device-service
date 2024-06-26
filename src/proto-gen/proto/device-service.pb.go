// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: proto/device-service.proto

package device_service

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

type Device_Type int32

const (
	Device_NO       Device_Type = 0
	Device_MOUSE    Device_Type = 1
	Device_KEYBOARD Device_Type = 2
)

// Enum value maps for Device_Type.
var (
	Device_Type_name = map[int32]string{
		0: "NO",
		1: "MOUSE",
		2: "KEYBOARD",
	}
	Device_Type_value = map[string]int32{
		"NO":       0,
		"MOUSE":    1,
		"KEYBOARD": 2,
	}
)

func (x Device_Type) Enum() *Device_Type {
	p := new(Device_Type)
	*p = x
	return p
}

func (x Device_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Device_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_device_service_proto_enumTypes[0].Descriptor()
}

func (Device_Type) Type() protoreflect.EnumType {
	return &file_proto_device_service_proto_enumTypes[0]
}

func (x Device_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Device_Type.Descriptor instead.
func (Device_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_device_service_proto_rawDescGZIP(), []int{5, 0}
}

type GetDeviceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDeviceListRequest) Reset() {
	*x = GetDeviceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_device_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceListRequest) ProtoMessage() {}

func (x *GetDeviceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_device_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceListRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceListRequest) Descriptor() ([]byte, []int) {
	return file_proto_device_service_proto_rawDescGZIP(), []int{0}
}

type DeviceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *DeviceListResponse) Reset() {
	*x = DeviceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_device_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceListResponse) ProtoMessage() {}

func (x *DeviceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_device_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceListResponse.ProtoReflect.Descriptor instead.
func (*DeviceListResponse) Descriptor() ([]byte, []int) {
	return file_proto_device_service_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceListResponse) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

type AddDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device *Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *AddDeviceRequest) Reset() {
	*x = AddDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_device_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDeviceRequest) ProtoMessage() {}

func (x *AddDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_device_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDeviceRequest.ProtoReflect.Descriptor instead.
func (*AddDeviceRequest) Descriptor() ([]byte, []int) {
	return file_proto_device_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddDeviceRequest) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

type DeleteDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDeviceRequest) Reset() {
	*x = DeleteDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_device_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeviceRequest) ProtoMessage() {}

func (x *DeleteDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_device_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeviceRequest.ProtoReflect.Descriptor instead.
func (*DeleteDeviceRequest) Descriptor() ([]byte, []int) {
	return file_proto_device_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteDeviceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type OperationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *OperationStatus) Reset() {
	*x = OperationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_device_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationStatus) ProtoMessage() {}

func (x *OperationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_device_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationStatus.ProtoReflect.Descriptor instead.
func (*OperationStatus) Descriptor() ([]byte, []int) {
	return file_proto_device_service_proto_rawDescGZIP(), []int{4}
}

func (x *OperationStatus) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type       Device_Type `protobuf:"varint,2,opt,name=type,proto3,enum=device_service.Device_Type" json:"type,omitempty"`
	DevicePath string      `protobuf:"bytes,3,opt,name=device_path,json=devicePath,proto3" json:"device_path,omitempty"`
	VendorId   string      `protobuf:"bytes,4,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
	ProductId  string      `protobuf:"bytes,5,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_device_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_proto_device_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_proto_device_service_proto_rawDescGZIP(), []int{5}
}

func (x *Device) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Device) GetType() Device_Type {
	if x != nil {
		return x.Type
	}
	return Device_NO
}

func (x *Device) GetDevicePath() string {
	if x != nil {
		return x.DevicePath
	}
	return ""
}

func (x *Device) GetVendorId() string {
	if x != nil {
		return x.VendorId
	}
	return ""
}

func (x *Device) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

var File_proto_device_service_proto protoreflect.FileDescriptor

var file_proto_device_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x16, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x12, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x10,
	0x41, 0x64, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0xcf, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x27, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x4f, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x4d, 0x4f, 0x55, 0x53, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x45, 0x59, 0x42,
	0x4f, 0x41, 0x52, 0x44, 0x10, 0x02, 0x32, 0x9c, 0x02, 0x0a, 0x15, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x59, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x24, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x09, 0x41,
	0x64, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x56, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6b, 0x61, 0x6c, 0x6f, 0x75, 0x73, 0x6b, 0x69, 0x2d, 0x61,
	0x6c, 0x65, 0x78, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x3b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_device_service_proto_rawDescOnce sync.Once
	file_proto_device_service_proto_rawDescData = file_proto_device_service_proto_rawDesc
)

func file_proto_device_service_proto_rawDescGZIP() []byte {
	file_proto_device_service_proto_rawDescOnce.Do(func() {
		file_proto_device_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_device_service_proto_rawDescData)
	})
	return file_proto_device_service_proto_rawDescData
}

var file_proto_device_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_device_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_device_service_proto_goTypes = []interface{}{
	(Device_Type)(0),             // 0: device_service.Device.Type
	(*GetDeviceListRequest)(nil), // 1: device_service.GetDeviceListRequest
	(*DeviceListResponse)(nil),   // 2: device_service.DeviceListResponse
	(*AddDeviceRequest)(nil),     // 3: device_service.AddDeviceRequest
	(*DeleteDeviceRequest)(nil),  // 4: device_service.DeleteDeviceRequest
	(*OperationStatus)(nil),      // 5: device_service.OperationStatus
	(*Device)(nil),               // 6: device_service.Device
}
var file_proto_device_service_proto_depIdxs = []int32{
	6, // 0: device_service.DeviceListResponse.devices:type_name -> device_service.Device
	6, // 1: device_service.AddDeviceRequest.device:type_name -> device_service.Device
	0, // 2: device_service.Device.type:type_name -> device_service.Device.Type
	1, // 3: device_service.DeviceProviderService.ListDevices:input_type -> device_service.GetDeviceListRequest
	3, // 4: device_service.DeviceProviderService.AddDevice:input_type -> device_service.AddDeviceRequest
	4, // 5: device_service.DeviceProviderService.DeleteDevice:input_type -> device_service.DeleteDeviceRequest
	2, // 6: device_service.DeviceProviderService.ListDevices:output_type -> device_service.DeviceListResponse
	5, // 7: device_service.DeviceProviderService.AddDevice:output_type -> device_service.OperationStatus
	5, // 8: device_service.DeviceProviderService.DeleteDevice:output_type -> device_service.OperationStatus
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_device_service_proto_init() }
func file_proto_device_service_proto_init() {
	if File_proto_device_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_device_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceListRequest); i {
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
		file_proto_device_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceListResponse); i {
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
		file_proto_device_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDeviceRequest); i {
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
		file_proto_device_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeviceRequest); i {
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
		file_proto_device_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationStatus); i {
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
		file_proto_device_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
			RawDescriptor: file_proto_device_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_device_service_proto_goTypes,
		DependencyIndexes: file_proto_device_service_proto_depIdxs,
		EnumInfos:         file_proto_device_service_proto_enumTypes,
		MessageInfos:      file_proto_device_service_proto_msgTypes,
	}.Build()
	File_proto_device_service_proto = out.File
	file_proto_device_service_proto_rawDesc = nil
	file_proto_device_service_proto_goTypes = nil
	file_proto_device_service_proto_depIdxs = nil
}
