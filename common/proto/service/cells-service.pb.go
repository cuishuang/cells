// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: cells-service.proto

package service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OperationType int32

const (
	OperationType_OR  OperationType = 0
	OperationType_AND OperationType = 1
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "OR",
		1: "AND",
	}
	OperationType_value = map[string]int32{
		"OR":  0,
		"AND": 1,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_cells_service_proto_enumTypes[0].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_cells_service_proto_enumTypes[0]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_cells_service_proto_rawDescGZIP(), []int{0}
}

type ResourcePolicyAction int32

const (
	ResourcePolicyAction_ANY        ResourcePolicyAction = 0
	ResourcePolicyAction_OWNER      ResourcePolicyAction = 1
	ResourcePolicyAction_READ       ResourcePolicyAction = 2
	ResourcePolicyAction_WRITE      ResourcePolicyAction = 3
	ResourcePolicyAction_EDIT_RULES ResourcePolicyAction = 4
)

// Enum value maps for ResourcePolicyAction.
var (
	ResourcePolicyAction_name = map[int32]string{
		0: "ANY",
		1: "OWNER",
		2: "READ",
		3: "WRITE",
		4: "EDIT_RULES",
	}
	ResourcePolicyAction_value = map[string]int32{
		"ANY":        0,
		"OWNER":      1,
		"READ":       2,
		"WRITE":      3,
		"EDIT_RULES": 4,
	}
)

func (x ResourcePolicyAction) Enum() *ResourcePolicyAction {
	p := new(ResourcePolicyAction)
	*p = x
	return p
}

func (x ResourcePolicyAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourcePolicyAction) Descriptor() protoreflect.EnumDescriptor {
	return file_cells_service_proto_enumTypes[1].Descriptor()
}

func (ResourcePolicyAction) Type() protoreflect.EnumType {
	return &file_cells_service_proto_enumTypes[1]
}

func (x ResourcePolicyAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourcePolicyAction.Descriptor instead.
func (ResourcePolicyAction) EnumDescriptor() ([]byte, []int) {
	return file_cells_service_proto_rawDescGZIP(), []int{1}
}

type ResourcePolicy_PolicyEffect int32

const (
	ResourcePolicy_deny  ResourcePolicy_PolicyEffect = 0
	ResourcePolicy_allow ResourcePolicy_PolicyEffect = 1
)

// Enum value maps for ResourcePolicy_PolicyEffect.
var (
	ResourcePolicy_PolicyEffect_name = map[int32]string{
		0: "deny",
		1: "allow",
	}
	ResourcePolicy_PolicyEffect_value = map[string]int32{
		"deny":  0,
		"allow": 1,
	}
)

func (x ResourcePolicy_PolicyEffect) Enum() *ResourcePolicy_PolicyEffect {
	p := new(ResourcePolicy_PolicyEffect)
	*p = x
	return p
}

func (x ResourcePolicy_PolicyEffect) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourcePolicy_PolicyEffect) Descriptor() protoreflect.EnumDescriptor {
	return file_cells_service_proto_enumTypes[2].Descriptor()
}

func (ResourcePolicy_PolicyEffect) Type() protoreflect.EnumType {
	return &file_cells_service_proto_enumTypes[2]
}

func (x ResourcePolicy_PolicyEffect) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourcePolicy_PolicyEffect.Descriptor instead.
func (ResourcePolicy_PolicyEffect) EnumDescriptor() ([]byte, []int) {
	return file_cells_service_proto_rawDescGZIP(), []int{2, 0}
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubQueries          []*anypb.Any         `protobuf:"bytes,1,rep,name=SubQueries,proto3" json:"SubQueries,omitempty"`
	Operation           OperationType        `protobuf:"varint,2,opt,name=Operation,proto3,enum=service.OperationType" json:"Operation,omitempty"`
	ResourcePolicyQuery *ResourcePolicyQuery `protobuf:"bytes,3,opt,name=ResourcePolicyQuery,proto3" json:"ResourcePolicyQuery,omitempty"`
	Offset              int64                `protobuf:"varint,4,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit               int64                `protobuf:"varint,5,opt,name=Limit,proto3" json:"Limit,omitempty"`
	GroupBy             int32                `protobuf:"varint,6,opt,name=groupBy,proto3" json:"groupBy,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_cells_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_cells_service_proto_rawDescGZIP(), []int{0}
}

func (x *Query) GetSubQueries() []*anypb.Any {
	if x != nil {
		return x.SubQueries
	}
	return nil
}

func (x *Query) GetOperation() OperationType {
	if x != nil {
		return x.Operation
	}
	return OperationType_OR
}

func (x *Query) GetResourcePolicyQuery() *ResourcePolicyQuery {
	if x != nil {
		return x.ResourcePolicyQuery
	}
	return nil
}

func (x *Query) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Query) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Query) GetGroupBy() int32 {
	if x != nil {
		return x.GroupBy
	}
	return 0
}

type ResourcePolicyQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subjects []string `protobuf:"bytes,1,rep,name=Subjects,proto3" json:"Subjects,omitempty"`
	Empty    bool     `protobuf:"varint,2,opt,name=Empty,proto3" json:"Empty,omitempty"`
	Any      bool     `protobuf:"varint,3,opt,name=Any,proto3" json:"Any,omitempty"`
}

func (x *ResourcePolicyQuery) Reset() {
	*x = ResourcePolicyQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourcePolicyQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcePolicyQuery) ProtoMessage() {}

func (x *ResourcePolicyQuery) ProtoReflect() protoreflect.Message {
	mi := &file_cells_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcePolicyQuery.ProtoReflect.Descriptor instead.
func (*ResourcePolicyQuery) Descriptor() ([]byte, []int) {
	return file_cells_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResourcePolicyQuery) GetSubjects() []string {
	if x != nil {
		return x.Subjects
	}
	return nil
}

func (x *ResourcePolicyQuery) GetEmpty() bool {
	if x != nil {
		return x.Empty
	}
	return false
}

func (x *ResourcePolicyQuery) GetAny() bool {
	if x != nil {
		return x.Any
	}
	return false
}

type ResourcePolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64                       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Resource       string                      `protobuf:"bytes,2,opt,name=Resource,proto3" json:"Resource,omitempty"`
	Action         ResourcePolicyAction        `protobuf:"varint,3,opt,name=Action,proto3,enum=service.ResourcePolicyAction" json:"Action,omitempty"`
	Subject        string                      `protobuf:"bytes,4,opt,name=Subject,proto3" json:"Subject,omitempty"`
	Effect         ResourcePolicy_PolicyEffect `protobuf:"varint,5,opt,name=Effect,proto3,enum=service.ResourcePolicy_PolicyEffect" json:"Effect,omitempty"`
	JsonConditions string                      `protobuf:"bytes,6,opt,name=JsonConditions,proto3" json:"JsonConditions,omitempty"`
}

func (x *ResourcePolicy) Reset() {
	*x = ResourcePolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourcePolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcePolicy) ProtoMessage() {}

func (x *ResourcePolicy) ProtoReflect() protoreflect.Message {
	mi := &file_cells_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcePolicy.ProtoReflect.Descriptor instead.
func (*ResourcePolicy) Descriptor() ([]byte, []int) {
	return file_cells_service_proto_rawDescGZIP(), []int{2}
}

func (x *ResourcePolicy) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResourcePolicy) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *ResourcePolicy) GetAction() ResourcePolicyAction {
	if x != nil {
		return x.Action
	}
	return ResourcePolicyAction_ANY
}

func (x *ResourcePolicy) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *ResourcePolicy) GetEffect() ResourcePolicy_PolicyEffect {
	if x != nil {
		return x.Effect
	}
	return ResourcePolicy_deny
}

func (x *ResourcePolicy) GetJsonConditions() string {
	if x != nil {
		return x.JsonConditions
	}
	return ""
}

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cells_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_cells_service_proto_rawDescGZIP(), []int{3}
}

func (x *StartRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OK bool `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cells_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_cells_service_proto_rawDescGZIP(), []int{4}
}

func (x *StartResponse) GetOK() bool {
	if x != nil {
		return x.OK
	}
	return false
}

type StopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cells_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRequest.ProtoReflect.Descriptor instead.
func (*StopRequest) Descriptor() ([]byte, []int) {
	return file_cells_service_proto_rawDescGZIP(), []int{5}
}

func (x *StopRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OK bool `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
}

func (x *StopResponse) Reset() {
	*x = StopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopResponse) ProtoMessage() {}

func (x *StopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cells_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopResponse.ProtoReflect.Descriptor instead.
func (*StopResponse) Descriptor() ([]byte, []int) {
	return file_cells_service_proto_rawDescGZIP(), []int{6}
}

func (x *StopResponse) GetOK() bool {
	if x != nil {
		return x.OK
	}
	return false
}

// ModifyLoginRequest is used to send a ModifyLogin call
type ModifyLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldLogin string            `protobuf:"bytes,1,opt,name=OldLogin,proto3" json:"OldLogin,omitempty"`
	NewLogin string            `protobuf:"bytes,2,opt,name=NewLogin,proto3" json:"NewLogin,omitempty"`
	DryRun   bool              `protobuf:"varint,3,opt,name=DryRun,proto3" json:"DryRun,omitempty"`
	Options  map[string]string `protobuf:"bytes,4,rep,name=Options,proto3" json:"Options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ModifyLoginRequest) Reset() {
	*x = ModifyLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyLoginRequest) ProtoMessage() {}

func (x *ModifyLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cells_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyLoginRequest.ProtoReflect.Descriptor instead.
func (*ModifyLoginRequest) Descriptor() ([]byte, []int) {
	return file_cells_service_proto_rawDescGZIP(), []int{7}
}

func (x *ModifyLoginRequest) GetOldLogin() string {
	if x != nil {
		return x.OldLogin
	}
	return ""
}

func (x *ModifyLoginRequest) GetNewLogin() string {
	if x != nil {
		return x.NewLogin
	}
	return ""
}

func (x *ModifyLoginRequest) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

func (x *ModifyLoginRequest) GetOptions() map[string]string {
	if x != nil {
		return x.Options
	}
	return nil
}

// ModifyLoginResponse indicates if operation succeeded and adds optional message
type ModifyLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Messages []string `protobuf:"bytes,2,rep,name=Messages,proto3" json:"Messages,omitempty"`
}

func (x *ModifyLoginResponse) Reset() {
	*x = ModifyLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyLoginResponse) ProtoMessage() {}

func (x *ModifyLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cells_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyLoginResponse.ProtoReflect.Descriptor instead.
func (*ModifyLoginResponse) Descriptor() ([]byte, []int) {
	return file_cells_service_proto_rawDescGZIP(), []int{8}
}

func (x *ModifyLoginResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ModifyLoginResponse) GetMessages() []string {
	if x != nil {
		return x.Messages
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Code    uint32 `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Status  string `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Details string `protobuf:"bytes,4,opt,name=Details,proto3" json:"Details,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_cells_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_cells_service_proto_rawDescGZIP(), []int{9}
}

func (x *Error) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Error) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Error) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

var File_cells_service_proto protoreflect.FileDescriptor

var file_cells_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x02, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x34, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0a, 0x53, 0x75, 0x62, 0x51,
	0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x13,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x13, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x42, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x42, 0x79, 0x22, 0x59, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x41, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x41, 0x6e, 0x79, 0x22,
	0x98, 0x02, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x35,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x3c, 0x0a, 0x06, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x06, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x26, 0x0a,
	0x0e, 0x4a, 0x73, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x4a, 0x73, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x23, 0x0a, 0x0c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x10, 0x01, 0x22, 0x22, 0x0a, 0x0c, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1f,
	0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x4f, 0x4b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f, 0x4b, 0x22,
	0x21, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x4b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02,
	0x4f, 0x4b, 0x22, 0xe4, 0x01, 0x0a, 0x12, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x6c, 0x64,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x6c, 0x64,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x65, 0x77, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x44, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x12, 0x42, 0x0a, 0x07, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3a, 0x0a,
	0x0c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4b, 0x0a, 0x13, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x5d, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x2a, 0x20, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x4e, 0x44, 0x10, 0x01, 0x2a, 0x4f, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x07, 0x0a, 0x03, 0x41, 0x4e, 0x59, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x57, 0x4e, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x09, 0x0a,
	0x05, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x44, 0x49, 0x54,
	0x5f, 0x52, 0x55, 0x4c, 0x45, 0x53, 0x10, 0x04, 0x32, 0x7d, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x59, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x79, 0x64, 0x69, 0x6f, 0x2f, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cells_service_proto_rawDescOnce sync.Once
	file_cells_service_proto_rawDescData = file_cells_service_proto_rawDesc
)

func file_cells_service_proto_rawDescGZIP() []byte {
	file_cells_service_proto_rawDescOnce.Do(func() {
		file_cells_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cells_service_proto_rawDescData)
	})
	return file_cells_service_proto_rawDescData
}

var file_cells_service_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_cells_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_cells_service_proto_goTypes = []interface{}{
	(OperationType)(0),               // 0: service.OperationType
	(ResourcePolicyAction)(0),        // 1: service.ResourcePolicyAction
	(ResourcePolicy_PolicyEffect)(0), // 2: service.ResourcePolicy.PolicyEffect
	(*Query)(nil),                    // 3: service.Query
	(*ResourcePolicyQuery)(nil),      // 4: service.ResourcePolicyQuery
	(*ResourcePolicy)(nil),           // 5: service.ResourcePolicy
	(*StartRequest)(nil),             // 6: service.StartRequest
	(*StartResponse)(nil),            // 7: service.StartResponse
	(*StopRequest)(nil),              // 8: service.StopRequest
	(*StopResponse)(nil),             // 9: service.StopResponse
	(*ModifyLoginRequest)(nil),       // 10: service.ModifyLoginRequest
	(*ModifyLoginResponse)(nil),      // 11: service.ModifyLoginResponse
	(*Error)(nil),                    // 12: service.Error
	nil,                              // 13: service.ModifyLoginRequest.OptionsEntry
	(*anypb.Any)(nil),                // 14: google.protobuf.Any
}
var file_cells_service_proto_depIdxs = []int32{
	14, // 0: service.Query.SubQueries:type_name -> google.protobuf.Any
	0,  // 1: service.Query.Operation:type_name -> service.OperationType
	4,  // 2: service.Query.ResourcePolicyQuery:type_name -> service.ResourcePolicyQuery
	1,  // 3: service.ResourcePolicy.Action:type_name -> service.ResourcePolicyAction
	2,  // 4: service.ResourcePolicy.Effect:type_name -> service.ResourcePolicy.PolicyEffect
	13, // 5: service.ModifyLoginRequest.Options:type_name -> service.ModifyLoginRequest.OptionsEntry
	6,  // 6: service.ServiceManager.Start:input_type -> service.StartRequest
	8,  // 7: service.ServiceManager.Stop:input_type -> service.StopRequest
	10, // 8: service.LoginModifier.ModifyLogin:input_type -> service.ModifyLoginRequest
	7,  // 9: service.ServiceManager.Start:output_type -> service.StartResponse
	9,  // 10: service.ServiceManager.Stop:output_type -> service.StopResponse
	11, // 11: service.LoginModifier.ModifyLogin:output_type -> service.ModifyLoginResponse
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_cells_service_proto_init() }
func file_cells_service_proto_init() {
	if File_cells_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cells_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_cells_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourcePolicyQuery); i {
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
		file_cells_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourcePolicy); i {
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
		file_cells_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_cells_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResponse); i {
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
		file_cells_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopRequest); i {
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
		file_cells_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopResponse); i {
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
		file_cells_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyLoginRequest); i {
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
		file_cells_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyLoginResponse); i {
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
		file_cells_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_cells_service_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_cells_service_proto_goTypes,
		DependencyIndexes: file_cells_service_proto_depIdxs,
		EnumInfos:         file_cells_service_proto_enumTypes,
		MessageInfos:      file_cells_service_proto_msgTypes,
	}.Build()
	File_cells_service_proto = out.File
	file_cells_service_proto_rawDesc = nil
	file_cells_service_proto_goTypes = nil
	file_cells_service_proto_depIdxs = nil
}
