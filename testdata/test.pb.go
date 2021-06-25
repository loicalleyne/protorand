// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: testdata/test.proto

package testproto

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

type SomeEnum int32

const (
	SomeEnum_SOME_ENUM_UNKNOWN SomeEnum = 0
	SomeEnum_SOME_ENUM_VALUE_1 SomeEnum = 1
	SomeEnum_SOME_ENUM_VALUE_2 SomeEnum = 2
)

// Enum value maps for SomeEnum.
var (
	SomeEnum_name = map[int32]string{
		0: "SOME_ENUM_UNKNOWN",
		1: "SOME_ENUM_VALUE_1",
		2: "SOME_ENUM_VALUE_2",
	}
	SomeEnum_value = map[string]int32{
		"SOME_ENUM_UNKNOWN": 0,
		"SOME_ENUM_VALUE_1": 1,
		"SOME_ENUM_VALUE_2": 2,
	}
)

func (x SomeEnum) Enum() *SomeEnum {
	p := new(SomeEnum)
	*p = x
	return p
}

func (x SomeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SomeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_testdata_test_proto_enumTypes[0].Descriptor()
}

func (SomeEnum) Type() protoreflect.EnumType {
	return &file_testdata_test_proto_enumTypes[0]
}

func (x SomeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SomeEnum.Descriptor instead.
func (SomeEnum) EnumDescriptor() ([]byte, []int) {
	return file_testdata_test_proto_rawDescGZIP(), []int{0}
}

type SomeEnum2 int32

const (
	SomeEnum2_SOME_ENUM_2_UNKNOWN SomeEnum2 = 0
)

// Enum value maps for SomeEnum2.
var (
	SomeEnum2_name = map[int32]string{
		0: "SOME_ENUM_2_UNKNOWN",
	}
	SomeEnum2_value = map[string]int32{
		"SOME_ENUM_2_UNKNOWN": 0,
	}
)

func (x SomeEnum2) Enum() *SomeEnum2 {
	p := new(SomeEnum2)
	*p = x
	return p
}

func (x SomeEnum2) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SomeEnum2) Descriptor() protoreflect.EnumDescriptor {
	return file_testdata_test_proto_enumTypes[1].Descriptor()
}

func (SomeEnum2) Type() protoreflect.EnumType {
	return &file_testdata_test_proto_enumTypes[1]
}

func (x SomeEnum2) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SomeEnum2.Descriptor instead.
func (SomeEnum2) EnumDescriptor() ([]byte, []int) {
	return file_testdata_test_proto_rawDescGZIP(), []int{1}
}

type TestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SomeStr     string                  `protobuf:"bytes,1,opt,name=some_str,json=someStr,proto3" json:"some_str,omitempty"`
	SomeInt32   int32                   `protobuf:"varint,2,opt,name=some_int32,json=someInt32,proto3" json:"some_int32,omitempty"`
	SomeInt64   int32                   `protobuf:"varint,3,opt,name=some_int64,json=someInt64,proto3" json:"some_int64,omitempty"`
	SomeFloat32 float32                 `protobuf:"fixed32,4,opt,name=some_float32,json=someFloat32,proto3" json:"some_float32,omitempty"`
	SomeFloat64 float32                 `protobuf:"fixed32,5,opt,name=some_float64,json=someFloat64,proto3" json:"some_float64,omitempty"`
	SomeBool    bool                    `protobuf:"varint,6,opt,name=some_bool,json=someBool,proto3" json:"some_bool,omitempty"`
	SomeSlice   []string                `protobuf:"bytes,7,rep,name=some_slice,json=someSlice,proto3" json:"some_slice,omitempty"`
	SomeMsg     *ChildMessage           `protobuf:"bytes,8,opt,name=some_msg,json=someMsg,proto3" json:"some_msg,omitempty"`
	SomeMsgs    []*ChildMessage         `protobuf:"bytes,9,rep,name=some_msgs,json=someMsgs,proto3" json:"some_msgs,omitempty"`
	SomeMap     map[int32]*ChildMessage `protobuf:"bytes,10,rep,name=some_map,json=someMap,proto3" json:"some_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SomeEnum    SomeEnum                `protobuf:"varint,11,opt,name=some_enum,json=someEnum,proto3,enum=testproto.SomeEnum" json:"some_enum,omitempty"`
	SomeEnum2   SomeEnum2               `protobuf:"varint,12,opt,name=some_enum2,json=someEnum2,proto3,enum=testproto.SomeEnum2" json:"some_enum2,omitempty"`
	// Types that are assignable to SomeOneOf:
	//	*TestMessage_OneOfInt32
	//	*TestMessage_OneOfStr
	SomeOneOf isTestMessage_SomeOneOf `protobuf_oneof:"some_one_of"`
}

func (x *TestMessage) Reset() {
	*x = TestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessage) ProtoMessage() {}

func (x *TestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessage.ProtoReflect.Descriptor instead.
func (*TestMessage) Descriptor() ([]byte, []int) {
	return file_testdata_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestMessage) GetSomeStr() string {
	if x != nil {
		return x.SomeStr
	}
	return ""
}

func (x *TestMessage) GetSomeInt32() int32 {
	if x != nil {
		return x.SomeInt32
	}
	return 0
}

func (x *TestMessage) GetSomeInt64() int32 {
	if x != nil {
		return x.SomeInt64
	}
	return 0
}

func (x *TestMessage) GetSomeFloat32() float32 {
	if x != nil {
		return x.SomeFloat32
	}
	return 0
}

func (x *TestMessage) GetSomeFloat64() float32 {
	if x != nil {
		return x.SomeFloat64
	}
	return 0
}

func (x *TestMessage) GetSomeBool() bool {
	if x != nil {
		return x.SomeBool
	}
	return false
}

func (x *TestMessage) GetSomeSlice() []string {
	if x != nil {
		return x.SomeSlice
	}
	return nil
}

func (x *TestMessage) GetSomeMsg() *ChildMessage {
	if x != nil {
		return x.SomeMsg
	}
	return nil
}

func (x *TestMessage) GetSomeMsgs() []*ChildMessage {
	if x != nil {
		return x.SomeMsgs
	}
	return nil
}

func (x *TestMessage) GetSomeMap() map[int32]*ChildMessage {
	if x != nil {
		return x.SomeMap
	}
	return nil
}

func (x *TestMessage) GetSomeEnum() SomeEnum {
	if x != nil {
		return x.SomeEnum
	}
	return SomeEnum_SOME_ENUM_UNKNOWN
}

func (x *TestMessage) GetSomeEnum2() SomeEnum2 {
	if x != nil {
		return x.SomeEnum2
	}
	return SomeEnum2_SOME_ENUM_2_UNKNOWN
}

func (m *TestMessage) GetSomeOneOf() isTestMessage_SomeOneOf {
	if m != nil {
		return m.SomeOneOf
	}
	return nil
}

func (x *TestMessage) GetOneOfInt32() int32 {
	if x, ok := x.GetSomeOneOf().(*TestMessage_OneOfInt32); ok {
		return x.OneOfInt32
	}
	return 0
}

func (x *TestMessage) GetOneOfStr() string {
	if x, ok := x.GetSomeOneOf().(*TestMessage_OneOfStr); ok {
		return x.OneOfStr
	}
	return ""
}

type isTestMessage_SomeOneOf interface {
	isTestMessage_SomeOneOf()
}

type TestMessage_OneOfInt32 struct {
	OneOfInt32 int32 `protobuf:"varint,13,opt,name=one_of_int32,json=oneOfInt32,proto3,oneof"`
}

type TestMessage_OneOfStr struct {
	OneOfStr string `protobuf:"bytes,14,opt,name=one_of_str,json=oneOfStr,proto3,oneof"`
}

func (*TestMessage_OneOfInt32) isTestMessage_SomeOneOf() {}

func (*TestMessage_OneOfStr) isTestMessage_SomeOneOf() {}

type ChildMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SomeInt int32 `protobuf:"varint,1,opt,name=some_int,json=someInt,proto3" json:"some_int,omitempty"`
}

func (x *ChildMessage) Reset() {
	*x = ChildMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChildMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChildMessage) ProtoMessage() {}

func (x *ChildMessage) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChildMessage.ProtoReflect.Descriptor instead.
func (*ChildMessage) Descriptor() ([]byte, []int) {
	return file_testdata_test_proto_rawDescGZIP(), []int{1}
}

func (x *ChildMessage) GetSomeInt() int32 {
	if x != nil {
		return x.SomeInt
	}
	return 0
}

var File_testdata_test_proto protoreflect.FileDescriptor

var file_testdata_test_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa1, 0x05, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x6f, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x73, 0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f,
	0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x73, 0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x6d,
	0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x33, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x73, 0x6f, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x33, 0x32, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0b, 0x73, 0x6f, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x73, 0x6f, 0x6d, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x73,
	0x6f, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x73, 0x6f, 0x6d, 0x65, 0x4d, 0x73, 0x67, 0x12,
	0x34, 0x0a, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x68, 0x69, 0x6c, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x73, 0x6f, 0x6d,
	0x65, 0x4d, 0x73, 0x67, 0x73, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x6d, 0x61,
	0x70, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x6f, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x6f,
	0x6d, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x30, 0x0a, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x65, 0x6e,
	0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x08, 0x73,
	0x6f, 0x6d, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x33, 0x0a, 0x0a, 0x73, 0x6f, 0x6d, 0x65, 0x5f,
	0x65, 0x6e, 0x75, 0x6d, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x32, 0x52, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x32, 0x12, 0x22, 0x0a, 0x0c,
	0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6e, 0x65, 0x4f, 0x66, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6f, 0x6e, 0x65, 0x4f, 0x66, 0x53, 0x74, 0x72,
	0x1a, 0x53, 0x0a, 0x0c, 0x53, 0x6f, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68,
	0x69, 0x6c, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x6f, 0x6e,
	0x65, 0x5f, 0x6f, 0x66, 0x22, 0x29, 0x0a, 0x0c, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x2a,
	0x4f, 0x0a, 0x08, 0x53, 0x6f, 0x6d, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x15, 0x0a, 0x11, 0x53,
	0x4f, 0x4d, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x4f, 0x4d, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x31, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x4f, 0x4d,
	0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x32, 0x10, 0x02,
	0x2a, 0x24, 0x0a, 0x09, 0x53, 0x6f, 0x6d, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x32, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x4f, 0x4d, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x32, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x72, 0x79, 0x6f, 0x79, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testdata_test_proto_rawDescOnce sync.Once
	file_testdata_test_proto_rawDescData = file_testdata_test_proto_rawDesc
)

func file_testdata_test_proto_rawDescGZIP() []byte {
	file_testdata_test_proto_rawDescOnce.Do(func() {
		file_testdata_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata_test_proto_rawDescData)
	})
	return file_testdata_test_proto_rawDescData
}

var file_testdata_test_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_testdata_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_testdata_test_proto_goTypes = []interface{}{
	(SomeEnum)(0),        // 0: testproto.SomeEnum
	(SomeEnum2)(0),       // 1: testproto.SomeEnum2
	(*TestMessage)(nil),  // 2: testproto.TestMessage
	(*ChildMessage)(nil), // 3: testproto.ChildMessage
	nil,                  // 4: testproto.TestMessage.SomeMapEntry
}
var file_testdata_test_proto_depIdxs = []int32{
	3, // 0: testproto.TestMessage.some_msg:type_name -> testproto.ChildMessage
	3, // 1: testproto.TestMessage.some_msgs:type_name -> testproto.ChildMessage
	4, // 2: testproto.TestMessage.some_map:type_name -> testproto.TestMessage.SomeMapEntry
	0, // 3: testproto.TestMessage.some_enum:type_name -> testproto.SomeEnum
	1, // 4: testproto.TestMessage.some_enum2:type_name -> testproto.SomeEnum2
	3, // 5: testproto.TestMessage.SomeMapEntry.value:type_name -> testproto.ChildMessage
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_testdata_test_proto_init() }
func file_testdata_test_proto_init() {
	if File_testdata_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testdata_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMessage); i {
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
		file_testdata_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChildMessage); i {
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
	file_testdata_test_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TestMessage_OneOfInt32)(nil),
		(*TestMessage_OneOfStr)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testdata_test_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testdata_test_proto_goTypes,
		DependencyIndexes: file_testdata_test_proto_depIdxs,
		EnumInfos:         file_testdata_test_proto_enumTypes,
		MessageInfos:      file_testdata_test_proto_msgTypes,
	}.Build()
	File_testdata_test_proto = out.File
	file_testdata_test_proto_rawDesc = nil
	file_testdata_test_proto_goTypes = nil
	file_testdata_test_proto_depIdxs = nil
}
