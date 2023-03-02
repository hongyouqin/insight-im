// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: chat.proto

package msg

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

type SendMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	OperationID string   `protobuf:"bytes,2,opt,name=operationID,proto3" json:"operationID,omitempty"`
	Data        *MsgData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SendMsgReq) Reset() {
	*x = SendMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgReq) ProtoMessage() {}

func (x *SendMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgReq.ProtoReflect.Descriptor instead.
func (*SendMsgReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *SendMsgReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SendMsgReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

func (x *SendMsgReq) GetData() *MsgData {
	if x != nil {
		return x.Data
	}
	return nil
}

type SendMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode     int32  `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMsg      string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	ServerMsgID string `protobuf:"bytes,4,opt,name=serverMsgID,proto3" json:"serverMsgID,omitempty"`
	ClientMsgID string `protobuf:"bytes,5,opt,name=clientMsgID,proto3" json:"clientMsgID,omitempty"`
	SendTime    int64  `protobuf:"varint,6,opt,name=sendTime,proto3" json:"sendTime,omitempty"`
}

func (x *SendMsgResp) Reset() {
	*x = SendMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgResp) ProtoMessage() {}

func (x *SendMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgResp.ProtoReflect.Descriptor instead.
func (*SendMsgResp) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *SendMsgResp) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *SendMsgResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *SendMsgResp) GetServerMsgID() string {
	if x != nil {
		return x.ServerMsgID
	}
	return ""
}

func (x *SendMsgResp) GetClientMsgID() string {
	if x != nil {
		return x.ClientMsgID
	}
	return ""
}

func (x *SendMsgResp) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

type UserSendMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerMsgID string `protobuf:"bytes,1,opt,name=serverMsgID,proto3" json:"serverMsgID,omitempty"`
	ClientMsgID string `protobuf:"bytes,2,opt,name=clientMsgID,proto3" json:"clientMsgID,omitempty"`
	SendTime    int64  `protobuf:"varint,3,opt,name=sendTime,proto3" json:"sendTime,omitempty"`
}

func (x *UserSendMsgResp) Reset() {
	*x = UserSendMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSendMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSendMsgResp) ProtoMessage() {}

func (x *UserSendMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSendMsgResp.ProtoReflect.Descriptor instead.
func (*UserSendMsgResp) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *UserSendMsgResp) GetServerMsgID() string {
	if x != nil {
		return x.ServerMsgID
	}
	return ""
}

func (x *UserSendMsgResp) GetClientMsgID() string {
	if x != nil {
		return x.ClientMsgID
	}
	return ""
}

func (x *UserSendMsgResp) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

type GetMaxAndMinSeqReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	OperationID string `protobuf:"bytes,2,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
}

func (x *GetMaxAndMinSeqReq) Reset() {
	*x = GetMaxAndMinSeqReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMaxAndMinSeqReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMaxAndMinSeqReq) ProtoMessage() {}

func (x *GetMaxAndMinSeqReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMaxAndMinSeqReq.ProtoReflect.Descriptor instead.
func (*GetMaxAndMinSeqReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *GetMaxAndMinSeqReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetMaxAndMinSeqReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

type GetMaxAndMinSeqResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32  `protobuf:"varint,1,opt,name=ErrCode,proto3" json:"ErrCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	MaxSeq  uint32 `protobuf:"varint,3,opt,name=MaxSeq,proto3" json:"MaxSeq,omitempty"`
	MinSeq  uint32 `protobuf:"varint,4,opt,name=MinSeq,proto3" json:"MinSeq,omitempty"`
}

func (x *GetMaxAndMinSeqResp) Reset() {
	*x = GetMaxAndMinSeqResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMaxAndMinSeqResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMaxAndMinSeqResp) ProtoMessage() {}

func (x *GetMaxAndMinSeqResp) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMaxAndMinSeqResp.ProtoReflect.Descriptor instead.
func (*GetMaxAndMinSeqResp) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *GetMaxAndMinSeqResp) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *GetMaxAndMinSeqResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *GetMaxAndMinSeqResp) GetMaxSeq() uint32 {
	if x != nil {
		return x.MaxSeq
	}
	return 0
}

func (x *GetMaxAndMinSeqResp) GetMinSeq() uint32 {
	if x != nil {
		return x.MinSeq
	}
	return 0
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68,
	0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9f, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x71, 0x0a, 0x0f, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4e, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x41, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x53, 0x65, 0x71,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x77, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x41, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x53, 0x65, 0x71,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x45, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x45, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x61, 0x78, 0x53, 0x65, 0x71,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4d, 0x61, 0x78, 0x53, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x4d, 0x69, 0x6e, 0x53, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x4d, 0x69, 0x6e, 0x53, 0x65, 0x71, 0x32, 0x82, 0x01, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12,
	0x30, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x48, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x41, 0x6e, 0x64, 0x4d, 0x69,
	0x6e, 0x53, 0x65, 0x71, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x78, 0x41, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x53, 0x65, 0x71, 0x52, 0x65, 0x71, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x41, 0x6e,
	0x64, 0x4d, 0x69, 0x6e, 0x53, 0x65, 0x71, 0x52, 0x65, 0x73, 0x70, 0x42, 0x08, 0x5a, 0x06, 0x2e,
	0x2f, 0x3b, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chat_proto_goTypes = []interface{}{
	(*SendMsgReq)(nil),          // 0: proto.SendMsgReq
	(*SendMsgResp)(nil),         // 1: proto.SendMsgResp
	(*UserSendMsgResp)(nil),     // 2: proto.UserSendMsgResp
	(*GetMaxAndMinSeqReq)(nil),  // 3: proto.GetMaxAndMinSeqReq
	(*GetMaxAndMinSeqResp)(nil), // 4: proto.GetMaxAndMinSeqResp
	(*MsgData)(nil),             // 5: proto.MsgData
}
var file_chat_proto_depIdxs = []int32{
	5, // 0: proto.SendMsgReq.data:type_name -> proto.MsgData
	0, // 1: proto.Chat.SendMsg:input_type -> proto.SendMsgReq
	3, // 2: proto.Chat.GetMaxAndMinSeq:input_type -> proto.GetMaxAndMinSeqReq
	1, // 3: proto.Chat.SendMsg:output_type -> proto.SendMsgResp
	4, // 4: proto.Chat.GetMaxAndMinSeq:output_type -> proto.GetMaxAndMinSeqResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	file_msg_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgReq); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgResp); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSendMsgResp); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMaxAndMinSeqReq); i {
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
		file_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMaxAndMinSeqResp); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}