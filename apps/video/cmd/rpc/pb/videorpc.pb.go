// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: rpc/pb/videorpc.proto

package pb

import (
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

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	OwnerId  int64  `protobuf:"varint,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	PlayUrl  string `protobuf:"bytes,4,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`
	CoverUrl string `protobuf:"bytes,5,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pb_videorpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pb_videorpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_rpc_pb_videorpc_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Video) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

type PublishVideoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	OwnerId  int64  `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	PlayUrl  string `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`
	CoverUrl string `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
}

func (x *PublishVideoReq) Reset() {
	*x = PublishVideoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pb_videorpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishVideoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishVideoReq) ProtoMessage() {}

func (x *PublishVideoReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pb_videorpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishVideoReq.ProtoReflect.Descriptor instead.
func (*PublishVideoReq) Descriptor() ([]byte, []int) {
	return file_rpc_pb_videorpc_proto_rawDescGZIP(), []int{1}
}

func (x *PublishVideoReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublishVideoReq) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *PublishVideoReq) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *PublishVideoReq) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

type PublishVideoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Video *Video `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *PublishVideoResp) Reset() {
	*x = PublishVideoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pb_videorpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishVideoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishVideoResp) ProtoMessage() {}

func (x *PublishVideoResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pb_videorpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishVideoResp.ProtoReflect.Descriptor instead.
func (*PublishVideoResp) Descriptor() ([]byte, []int) {
	return file_rpc_pb_videorpc_proto_rawDescGZIP(), []int{2}
}

func (x *PublishVideoResp) GetVideo() *Video {
	if x != nil {
		return x.Video
	}
	return nil
}

type VideoFeedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastTimeStamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=last_time_stamp,json=lastTimeStamp,proto3" json:"last_time_stamp,omitempty"`
}

func (x *VideoFeedReq) Reset() {
	*x = VideoFeedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pb_videorpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoFeedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoFeedReq) ProtoMessage() {}

func (x *VideoFeedReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pb_videorpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoFeedReq.ProtoReflect.Descriptor instead.
func (*VideoFeedReq) Descriptor() ([]byte, []int) {
	return file_rpc_pb_videorpc_proto_rawDescGZIP(), []int{3}
}

func (x *VideoFeedReq) GetLastTimeStamp() *timestamppb.Timestamp {
	if x != nil {
		return x.LastTimeStamp
	}
	return nil
}

type VideoFeedResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos []*Video `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
}

func (x *VideoFeedResp) Reset() {
	*x = VideoFeedResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pb_videorpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoFeedResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoFeedResp) ProtoMessage() {}

func (x *VideoFeedResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pb_videorpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoFeedResp.ProtoReflect.Descriptor instead.
func (*VideoFeedResp) Descriptor() ([]byte, []int) {
	return file_rpc_pb_videorpc_proto_rawDescGZIP(), []int{4}
}

func (x *VideoFeedResp) GetVideos() []*Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

type UserVideoListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserVideoListReq) Reset() {
	*x = UserVideoListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pb_videorpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserVideoListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVideoListReq) ProtoMessage() {}

func (x *UserVideoListReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pb_videorpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVideoListReq.ProtoReflect.Descriptor instead.
func (*UserVideoListReq) Descriptor() ([]byte, []int) {
	return file_rpc_pb_videorpc_proto_rawDescGZIP(), []int{5}
}

func (x *UserVideoListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserVideoListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos []*Video `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
}

func (x *UserVideoListResp) Reset() {
	*x = UserVideoListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pb_videorpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserVideoListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVideoListResp) ProtoMessage() {}

func (x *UserVideoListResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pb_videorpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVideoListResp.ProtoReflect.Descriptor instead.
func (*UserVideoListResp) Descriptor() ([]byte, []int) {
	return file_rpc_pb_videorpc_proto_rawDescGZIP(), []int{6}
}

func (x *UserVideoListResp) GetVideos() []*Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

var File_rpc_pb_videorpc_proto protoreflect.FileDescriptor

var file_rpc_pb_videorpc_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a,
	0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55,
	0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x22,
	0x7a, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x33, 0x0a, 0x10, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x1f, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x22, 0x52, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x42, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x32, 0x0a, 0x0d, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x65, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x22, 0x2b, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x06, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x32, 0xb5, 0x01,
	0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x70, 0x63, 0x12, 0x39, 0x0a, 0x0c, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x1a,
	0x14, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x65,
	0x65, 0x64, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x65, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_pb_videorpc_proto_rawDescOnce sync.Once
	file_rpc_pb_videorpc_proto_rawDescData = file_rpc_pb_videorpc_proto_rawDesc
)

func file_rpc_pb_videorpc_proto_rawDescGZIP() []byte {
	file_rpc_pb_videorpc_proto_rawDescOnce.Do(func() {
		file_rpc_pb_videorpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_pb_videorpc_proto_rawDescData)
	})
	return file_rpc_pb_videorpc_proto_rawDescData
}

var file_rpc_pb_videorpc_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rpc_pb_videorpc_proto_goTypes = []interface{}{
	(*Video)(nil),                 // 0: pb.Video
	(*PublishVideoReq)(nil),       // 1: pb.PublishVideoReq
	(*PublishVideoResp)(nil),      // 2: pb.PublishVideoResp
	(*VideoFeedReq)(nil),          // 3: pb.VideoFeedReq
	(*VideoFeedResp)(nil),         // 4: pb.VideoFeedResp
	(*UserVideoListReq)(nil),      // 5: pb.UserVideoListReq
	(*UserVideoListResp)(nil),     // 6: pb.UserVideoListResp
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_rpc_pb_videorpc_proto_depIdxs = []int32{
	0, // 0: pb.PublishVideoResp.video:type_name -> pb.Video
	7, // 1: pb.VideoFeedReq.last_time_stamp:type_name -> google.protobuf.Timestamp
	0, // 2: pb.VideoFeedResp.videos:type_name -> pb.Video
	0, // 3: pb.UserVideoListResp.videos:type_name -> pb.Video
	1, // 4: pb.videorpc.PublishVideo:input_type -> pb.PublishVideoReq
	3, // 5: pb.videorpc.VideoFeed:input_type -> pb.VideoFeedReq
	5, // 6: pb.videorpc.UserVideoList:input_type -> pb.UserVideoListReq
	2, // 7: pb.videorpc.PublishVideo:output_type -> pb.PublishVideoResp
	4, // 8: pb.videorpc.VideoFeed:output_type -> pb.VideoFeedResp
	6, // 9: pb.videorpc.UserVideoList:output_type -> pb.UserVideoListResp
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_pb_videorpc_proto_init() }
func file_rpc_pb_videorpc_proto_init() {
	if File_rpc_pb_videorpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_pb_videorpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_rpc_pb_videorpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishVideoReq); i {
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
		file_rpc_pb_videorpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishVideoResp); i {
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
		file_rpc_pb_videorpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoFeedReq); i {
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
		file_rpc_pb_videorpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoFeedResp); i {
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
		file_rpc_pb_videorpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserVideoListReq); i {
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
		file_rpc_pb_videorpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserVideoListResp); i {
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
			RawDescriptor: file_rpc_pb_videorpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_pb_videorpc_proto_goTypes,
		DependencyIndexes: file_rpc_pb_videorpc_proto_depIdxs,
		MessageInfos:      file_rpc_pb_videorpc_proto_msgTypes,
	}.Build()
	File_rpc_pb_videorpc_proto = out.File
	file_rpc_pb_videorpc_proto_rawDesc = nil
	file_rpc_pb_videorpc_proto_goTypes = nil
	file_rpc_pb_videorpc_proto_depIdxs = nil
}
