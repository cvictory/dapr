// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.0
// source: dapr/proto/common/v1/common.proto

package common

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Type of HTTP 1.1 Methods
// RFC 7231: https://tools.ietf.org/html/rfc7231#page-24
type HTTPExtension_Verb int32

const (
	HTTPExtension_NONE    HTTPExtension_Verb = 0
	HTTPExtension_GET     HTTPExtension_Verb = 1
	HTTPExtension_HEAD    HTTPExtension_Verb = 2
	HTTPExtension_POST    HTTPExtension_Verb = 3
	HTTPExtension_PUT     HTTPExtension_Verb = 4
	HTTPExtension_DELETE  HTTPExtension_Verb = 5
	HTTPExtension_CONNECT HTTPExtension_Verb = 6
	HTTPExtension_OPTIONS HTTPExtension_Verb = 7
	HTTPExtension_TRACE   HTTPExtension_Verb = 8
)

// Enum value maps for HTTPExtension_Verb.
var (
	HTTPExtension_Verb_name = map[int32]string{
		0: "NONE",
		1: "GET",
		2: "HEAD",
		3: "POST",
		4: "PUT",
		5: "DELETE",
		6: "CONNECT",
		7: "OPTIONS",
		8: "TRACE",
	}
	HTTPExtension_Verb_value = map[string]int32{
		"NONE":    0,
		"GET":     1,
		"HEAD":    2,
		"POST":    3,
		"PUT":     4,
		"DELETE":  5,
		"CONNECT": 6,
		"OPTIONS": 7,
		"TRACE":   8,
	}
)

func (x HTTPExtension_Verb) Enum() *HTTPExtension_Verb {
	p := new(HTTPExtension_Verb)
	*p = x
	return p
}

func (x HTTPExtension_Verb) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HTTPExtension_Verb) Descriptor() protoreflect.EnumDescriptor {
	return file_dapr_proto_common_v1_common_proto_enumTypes[0].Descriptor()
}

func (HTTPExtension_Verb) Type() protoreflect.EnumType {
	return &file_dapr_proto_common_v1_common_proto_enumTypes[0]
}

func (x HTTPExtension_Verb) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HTTPExtension_Verb.Descriptor instead.
func (HTTPExtension_Verb) EnumDescriptor() ([]byte, []int) {
	return file_dapr_proto_common_v1_common_proto_rawDescGZIP(), []int{0, 0}
}

// Enum describing the supported concurrency for state.
type StateOptions_StateConcurrency int32

const (
	StateOptions_CONCURRENCY_UNSPECIFIED StateOptions_StateConcurrency = 0
	StateOptions_CONCURRENCY_FIRST_WRITE StateOptions_StateConcurrency = 1
	StateOptions_CONCURRENCY_LAST_WRITE  StateOptions_StateConcurrency = 2
)

// Enum value maps for StateOptions_StateConcurrency.
var (
	StateOptions_StateConcurrency_name = map[int32]string{
		0: "CONCURRENCY_UNSPECIFIED",
		1: "CONCURRENCY_FIRST_WRITE",
		2: "CONCURRENCY_LAST_WRITE",
	}
	StateOptions_StateConcurrency_value = map[string]int32{
		"CONCURRENCY_UNSPECIFIED": 0,
		"CONCURRENCY_FIRST_WRITE": 1,
		"CONCURRENCY_LAST_WRITE":  2,
	}
)

func (x StateOptions_StateConcurrency) Enum() *StateOptions_StateConcurrency {
	p := new(StateOptions_StateConcurrency)
	*p = x
	return p
}

func (x StateOptions_StateConcurrency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateOptions_StateConcurrency) Descriptor() protoreflect.EnumDescriptor {
	return file_dapr_proto_common_v1_common_proto_enumTypes[1].Descriptor()
}

func (StateOptions_StateConcurrency) Type() protoreflect.EnumType {
	return &file_dapr_proto_common_v1_common_proto_enumTypes[1]
}

func (x StateOptions_StateConcurrency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateOptions_StateConcurrency.Descriptor instead.
func (StateOptions_StateConcurrency) EnumDescriptor() ([]byte, []int) {
	return file_dapr_proto_common_v1_common_proto_rawDescGZIP(), []int{5, 0}
}

// Enum describing the supported consistency for state.
type StateOptions_StateConsistency int32

const (
	StateOptions_CONSISTENCY_UNSPECIFIED StateOptions_StateConsistency = 0
	StateOptions_CONSISTENCY_EVENTUAL    StateOptions_StateConsistency = 1
	StateOptions_CONSISTENCY_STRONG      StateOptions_StateConsistency = 2
)

// Enum value maps for StateOptions_StateConsistency.
var (
	StateOptions_StateConsistency_name = map[int32]string{
		0: "CONSISTENCY_UNSPECIFIED",
		1: "CONSISTENCY_EVENTUAL",
		2: "CONSISTENCY_STRONG",
	}
	StateOptions_StateConsistency_value = map[string]int32{
		"CONSISTENCY_UNSPECIFIED": 0,
		"CONSISTENCY_EVENTUAL":    1,
		"CONSISTENCY_STRONG":      2,
	}
)

func (x StateOptions_StateConsistency) Enum() *StateOptions_StateConsistency {
	p := new(StateOptions_StateConsistency)
	*p = x
	return p
}

func (x StateOptions_StateConsistency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateOptions_StateConsistency) Descriptor() protoreflect.EnumDescriptor {
	return file_dapr_proto_common_v1_common_proto_enumTypes[2].Descriptor()
}

func (StateOptions_StateConsistency) Type() protoreflect.EnumType {
	return &file_dapr_proto_common_v1_common_proto_enumTypes[2]
}

func (x StateOptions_StateConsistency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateOptions_StateConsistency.Descriptor instead.
func (StateOptions_StateConsistency) EnumDescriptor() ([]byte, []int) {
	return file_dapr_proto_common_v1_common_proto_rawDescGZIP(), []int{5, 1}
}

// HTTPExtension includes HTTP verb and querystring
// when Dapr runtime delivers HTTP content.
//
// For example, when callers calls http invoke api
// POST http://localhost:3500/v1.0/invoke/<app_id>/method/<method>?query1=value1&query2=value2
//
// Dapr runtime will parse POST as a verb and extract querystring to quersytring map.
type HTTPExtension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. HTTP verb.
	Verb HTTPExtension_Verb `protobuf:"varint,1,opt,name=verb,proto3,enum=dapr.proto.common.v1.HTTPExtension_Verb" json:"verb,omitempty"`
	// querystring includes HTTP querystring.
	Querystring map[string]string `protobuf:"bytes,2,rep,name=querystring,proto3" json:"querystring,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HTTPExtension) Reset() {
	*x = HTTPExtension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_common_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPExtension) ProtoMessage() {}

func (x *HTTPExtension) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_common_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPExtension.ProtoReflect.Descriptor instead.
func (*HTTPExtension) Descriptor() ([]byte, []int) {
	return file_dapr_proto_common_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *HTTPExtension) GetVerb() HTTPExtension_Verb {
	if x != nil {
		return x.Verb
	}
	return HTTPExtension_NONE
}

func (x *HTTPExtension) GetQuerystring() map[string]string {
	if x != nil {
		return x.Querystring
	}
	return nil
}

// InvokeRequest is the message to invoke a method with the data.
// This message is used in InvokeService of Dapr gRPC Service and OnInvoke
// of AppCallback gRPC service.
type InvokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. method is a method name which will be invoked by caller.
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// Required. Bytes value or Protobuf message which caller sent.
	// Dapr treats Any.value as bytes type if Any.type_url is unset.
	Data *any.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// The type of data content.
	//
	// This field is required if data delivers http request body
	// Otherwise, this is optional.
	ContentType string `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// HTTP specific fields if request conveys http-compatible request.
	//
	// This field is required for http-compatible request. Otherwise,
	// this field is optional.
	HttpExtension *HTTPExtension `protobuf:"bytes,4,opt,name=http_extension,json=httpExtension,proto3" json:"http_extension,omitempty"`
}

func (x *InvokeRequest) Reset() {
	*x = InvokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_common_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeRequest) ProtoMessage() {}

func (x *InvokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_common_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeRequest.ProtoReflect.Descriptor instead.
func (*InvokeRequest) Descriptor() ([]byte, []int) {
	return file_dapr_proto_common_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *InvokeRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *InvokeRequest) GetData() *any.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *InvokeRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *InvokeRequest) GetHttpExtension() *HTTPExtension {
	if x != nil {
		return x.HttpExtension
	}
	return nil
}

// InvokeResponse is the response message inclduing data and its content type
// from app callback.
// This message is used in InvokeService of Dapr gRPC Service and OnInvoke
// of AppCallback gRPC service.
type InvokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The content body of InvokeService response.
	Data *any.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Required. The type of data content.
	ContentType string `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (x *InvokeResponse) Reset() {
	*x = InvokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_common_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeResponse) ProtoMessage() {}

func (x *InvokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_common_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeResponse.ProtoReflect.Descriptor instead.
func (*InvokeResponse) Descriptor() ([]byte, []int) {
	return file_dapr_proto_common_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *InvokeResponse) GetData() *any.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *InvokeResponse) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

// StateItem represents state key, value, and additional options to save state.
type StateItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The state key
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Required. The state data for key
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// The entity tag which represents the specific version of data.
	// The exact ETag format is defined by the corresponding data store.
	Etag *Etag `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	// The metadata which will be passed to state store component.
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Options for concurrency and consistency to save the state.
	Options *StateOptions `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *StateItem) Reset() {
	*x = StateItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_common_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateItem) ProtoMessage() {}

func (x *StateItem) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_common_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateItem.ProtoReflect.Descriptor instead.
func (*StateItem) Descriptor() ([]byte, []int) {
	return file_dapr_proto_common_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *StateItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *StateItem) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *StateItem) GetEtag() *Etag {
	if x != nil {
		return x.Etag
	}
	return nil
}

func (x *StateItem) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *StateItem) GetOptions() *StateOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

// Etag represents a state item version
type Etag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// value sets the etag value
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Etag) Reset() {
	*x = Etag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_common_v1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Etag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Etag) ProtoMessage() {}

func (x *Etag) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_common_v1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Etag.ProtoReflect.Descriptor instead.
func (*Etag) Descriptor() ([]byte, []int) {
	return file_dapr_proto_common_v1_common_proto_rawDescGZIP(), []int{4}
}

func (x *Etag) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// StateOptions configures concurrency and consistency for state operations
type StateOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Concurrency StateOptions_StateConcurrency `protobuf:"varint,1,opt,name=concurrency,proto3,enum=dapr.proto.common.v1.StateOptions_StateConcurrency" json:"concurrency,omitempty"`
	Consistency StateOptions_StateConsistency `protobuf:"varint,2,opt,name=consistency,proto3,enum=dapr.proto.common.v1.StateOptions_StateConsistency" json:"consistency,omitempty"`
}

func (x *StateOptions) Reset() {
	*x = StateOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_common_v1_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateOptions) ProtoMessage() {}

func (x *StateOptions) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_common_v1_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateOptions.ProtoReflect.Descriptor instead.
func (*StateOptions) Descriptor() ([]byte, []int) {
	return file_dapr_proto_common_v1_common_proto_rawDescGZIP(), []int{5}
}

func (x *StateOptions) GetConcurrency() StateOptions_StateConcurrency {
	if x != nil {
		return x.Concurrency
	}
	return StateOptions_CONCURRENCY_UNSPECIFIED
}

func (x *StateOptions) GetConsistency() StateOptions_StateConsistency {
	if x != nil {
		return x.Consistency
	}
	return StateOptions_CONSISTENCY_UNSPECIFIED
}

var File_dapr_proto_common_v1_common_proto protoreflect.FileDescriptor

var file_dapr_proto_common_v1_common_proto_rawDesc = []byte{
	0x0a, 0x21, 0x64, 0x61, 0x70, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x02, 0x0a, 0x0d, 0x48, 0x54, 0x54, 0x50, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x04, 0x76, 0x65, 0x72, 0x62, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x54, 0x54, 0x50,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x62, 0x52, 0x04,
	0x76, 0x65, 0x72, 0x62, 0x12, 0x56, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x64, 0x61, 0x70, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x54, 0x54, 0x50, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x3e, 0x0a, 0x10,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x67, 0x0a, 0x04,
	0x56, 0x65, 0x72, 0x62, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x41, 0x44, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x50,
	0x55, 0x54, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x05,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x06, 0x12, 0x0b, 0x0a,
	0x07, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52,
	0x41, 0x43, 0x45, 0x10, 0x08, 0x22, 0xc0, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a, 0x0e,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x54, 0x54, 0x50,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xa9, 0x02, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2e, 0x0a,
	0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x61,
	0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x74, 0x61, 0x67, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x49, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x61, 0x70, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x1c, 0x0a, 0x04, 0x45, 0x74, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x89, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x55, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x55, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33,
	0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79,
	0x22, 0x68, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4e, 0x43, 0x55, 0x52, 0x52, 0x45,
	0x4e, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4e, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59,
	0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x01, 0x12, 0x1a,
	0x0a, 0x16, 0x43, 0x4f, 0x4e, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x4c, 0x41,
	0x53, 0x54, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x02, 0x22, 0x61, 0x0a, 0x10, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1b,
	0x0a, 0x17, 0x43, 0x4f, 0x4e, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x43,
	0x4f, 0x4e, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x55, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4e, 0x53, 0x49, 0x53, 0x54,
	0x45, 0x4e, 0x43, 0x59, 0x5f, 0x53, 0x54, 0x52, 0x4f, 0x4e, 0x47, 0x10, 0x02, 0x42, 0x69, 0x0a,
	0x0a, 0x69, 0x6f, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x70, 0x72, 0x2f, 0x64, 0x61, 0x70, 0x72, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xaa, 0x02, 0x1b, 0x44, 0x61, 0x70,
	0x72, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dapr_proto_common_v1_common_proto_rawDescOnce sync.Once
	file_dapr_proto_common_v1_common_proto_rawDescData = file_dapr_proto_common_v1_common_proto_rawDesc
)

func file_dapr_proto_common_v1_common_proto_rawDescGZIP() []byte {
	file_dapr_proto_common_v1_common_proto_rawDescOnce.Do(func() {
		file_dapr_proto_common_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_dapr_proto_common_v1_common_proto_rawDescData)
	})
	return file_dapr_proto_common_v1_common_proto_rawDescData
}

var file_dapr_proto_common_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_dapr_proto_common_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_dapr_proto_common_v1_common_proto_goTypes = []interface{}{
	(HTTPExtension_Verb)(0),            // 0: dapr.proto.common.v1.HTTPExtension.Verb
	(StateOptions_StateConcurrency)(0), // 1: dapr.proto.common.v1.StateOptions.StateConcurrency
	(StateOptions_StateConsistency)(0), // 2: dapr.proto.common.v1.StateOptions.StateConsistency
	(*HTTPExtension)(nil),              // 3: dapr.proto.common.v1.HTTPExtension
	(*InvokeRequest)(nil),              // 4: dapr.proto.common.v1.InvokeRequest
	(*InvokeResponse)(nil),             // 5: dapr.proto.common.v1.InvokeResponse
	(*StateItem)(nil),                  // 6: dapr.proto.common.v1.StateItem
	(*Etag)(nil),                       // 7: dapr.proto.common.v1.Etag
	(*StateOptions)(nil),               // 8: dapr.proto.common.v1.StateOptions
	nil,                                // 9: dapr.proto.common.v1.HTTPExtension.QuerystringEntry
	nil,                                // 10: dapr.proto.common.v1.StateItem.MetadataEntry
	(*any.Any)(nil),                    // 11: google.protobuf.Any
}
var file_dapr_proto_common_v1_common_proto_depIdxs = []int32{
	0,  // 0: dapr.proto.common.v1.HTTPExtension.verb:type_name -> dapr.proto.common.v1.HTTPExtension.Verb
	9,  // 1: dapr.proto.common.v1.HTTPExtension.querystring:type_name -> dapr.proto.common.v1.HTTPExtension.QuerystringEntry
	11, // 2: dapr.proto.common.v1.InvokeRequest.data:type_name -> google.protobuf.Any
	3,  // 3: dapr.proto.common.v1.InvokeRequest.http_extension:type_name -> dapr.proto.common.v1.HTTPExtension
	11, // 4: dapr.proto.common.v1.InvokeResponse.data:type_name -> google.protobuf.Any
	7,  // 5: dapr.proto.common.v1.StateItem.etag:type_name -> dapr.proto.common.v1.Etag
	10, // 6: dapr.proto.common.v1.StateItem.metadata:type_name -> dapr.proto.common.v1.StateItem.MetadataEntry
	8,  // 7: dapr.proto.common.v1.StateItem.options:type_name -> dapr.proto.common.v1.StateOptions
	1,  // 8: dapr.proto.common.v1.StateOptions.concurrency:type_name -> dapr.proto.common.v1.StateOptions.StateConcurrency
	2,  // 9: dapr.proto.common.v1.StateOptions.consistency:type_name -> dapr.proto.common.v1.StateOptions.StateConsistency
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_dapr_proto_common_v1_common_proto_init() }
func file_dapr_proto_common_v1_common_proto_init() {
	if File_dapr_proto_common_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dapr_proto_common_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPExtension); i {
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
		file_dapr_proto_common_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeRequest); i {
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
		file_dapr_proto_common_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeResponse); i {
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
		file_dapr_proto_common_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateItem); i {
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
		file_dapr_proto_common_v1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Etag); i {
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
		file_dapr_proto_common_v1_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateOptions); i {
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
			RawDescriptor: file_dapr_proto_common_v1_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dapr_proto_common_v1_common_proto_goTypes,
		DependencyIndexes: file_dapr_proto_common_v1_common_proto_depIdxs,
		EnumInfos:         file_dapr_proto_common_v1_common_proto_enumTypes,
		MessageInfos:      file_dapr_proto_common_v1_common_proto_msgTypes,
	}.Build()
	File_dapr_proto_common_v1_common_proto = out.File
	file_dapr_proto_common_v1_common_proto_rawDesc = nil
	file_dapr_proto_common_v1_common_proto_goTypes = nil
	file_dapr_proto_common_v1_common_proto_depIdxs = nil
}
