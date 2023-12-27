// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: admin_casts_service_v1_messages.proto

package protos

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

type UserErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UserErrorMessage) Reset() {
	*x = UserErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_casts_service_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserErrorMessage) ProtoMessage() {}

func (x *UserErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_admin_casts_service_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserErrorMessage.ProtoReflect.Descriptor instead.
func (*UserErrorMessage) Descriptor() ([]byte, []int) {
	return file_admin_casts_service_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *UserErrorMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetCastsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// use ',' as separator for multiple ids, or leave it blank if you want to get all the casts on the page
	MoviesIDs string `protobuf:"bytes,1,opt,name=MoviesIDs,json=movies_ids,proto3" json:"MoviesIDs,omitempty"`
	// must be in range 10-100
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// must be > 0
	Page int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetCastsRequest) Reset() {
	*x = GetCastsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_casts_service_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCastsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCastsRequest) ProtoMessage() {}

func (x *GetCastsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_casts_service_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCastsRequest.ProtoReflect.Descriptor instead.
func (*GetCastsRequest) Descriptor() ([]byte, []int) {
	return file_admin_casts_service_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GetCastsRequest) GetMoviesIDs() string {
	if x != nil {
		return x.MoviesIDs
	}
	return ""
}

func (x *GetCastsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetCastsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type Casts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Casts map[int32]*Cast `protobuf:"bytes,1,rep,name=casts,proto3" json:"casts,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Casts) Reset() {
	*x = Casts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_casts_service_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Casts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Casts) ProtoMessage() {}

func (x *Casts) ProtoReflect() protoreflect.Message {
	mi := &file_admin_casts_service_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Casts.ProtoReflect.Descriptor instead.
func (*Casts) Descriptor() ([]byte, []int) {
	return file_admin_casts_service_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *Casts) GetCasts() map[int32]*Cast {
	if x != nil {
		return x.Casts
	}
	return nil
}

type Cast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CastLabel string  `protobuf:"bytes,1,opt,name=CastLabel,json=cast_label,proto3" json:"CastLabel,omitempty"`
	MovieID   int32   `protobuf:"varint,2,opt,name=MovieID,json=movie_id,proto3" json:"MovieID,omitempty"`
	ActorsIDs []int32 `protobuf:"varint,3,rep,packed,name=ActorsIDs,json=actors_ids,proto3" json:"ActorsIDs,omitempty"`
}

func (x *Cast) Reset() {
	*x = Cast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_casts_service_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cast) ProtoMessage() {}

func (x *Cast) ProtoReflect() protoreflect.Message {
	mi := &file_admin_casts_service_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cast.ProtoReflect.Descriptor instead.
func (*Cast) Descriptor() ([]byte, []int) {
	return file_admin_casts_service_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *Cast) GetCastLabel() string {
	if x != nil {
		return x.CastLabel
	}
	return ""
}

func (x *Cast) GetMovieID() int32 {
	if x != nil {
		return x.MovieID
	}
	return 0
}

func (x *Cast) GetActorsIDs() []int32 {
	if x != nil {
		return x.ActorsIDs
	}
	return nil
}

type GetCastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId int32 `protobuf:"varint,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
}

func (x *GetCastRequest) Reset() {
	*x = GetCastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_casts_service_v1_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCastRequest) ProtoMessage() {}

func (x *GetCastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_casts_service_v1_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCastRequest.ProtoReflect.Descriptor instead.
func (*GetCastRequest) Descriptor() ([]byte, []int) {
	return file_admin_casts_service_v1_messages_proto_rawDescGZIP(), []int{4}
}

func (x *GetCastRequest) GetMovieId() int32 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

type SearchCastByLabelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// must be in range 10-100
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// must be > 0
	Page int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *SearchCastByLabelRequest) Reset() {
	*x = SearchCastByLabelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_casts_service_v1_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCastByLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCastByLabelRequest) ProtoMessage() {}

func (x *SearchCastByLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_casts_service_v1_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCastByLabelRequest.ProtoReflect.Descriptor instead.
func (*SearchCastByLabelRequest) Descriptor() ([]byte, []int) {
	return file_admin_casts_service_v1_messages_proto_rawDescGZIP(), []int{5}
}

func (x *SearchCastByLabelRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *SearchCastByLabelRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SearchCastByLabelRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type CreateCastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CastLabel string  `protobuf:"bytes,1,opt,name=CastLabel,json=cast_label,proto3" json:"CastLabel,omitempty"`
	MovieID   int32   `protobuf:"varint,2,opt,name=MovieID,json=movie_id,proto3" json:"MovieID,omitempty"`
	ActorsIDs []int32 `protobuf:"varint,3,rep,packed,name=ActorsIDs,json=actors_ids,proto3" json:"ActorsIDs,omitempty"`
}

func (x *CreateCastRequest) Reset() {
	*x = CreateCastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_casts_service_v1_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCastRequest) ProtoMessage() {}

func (x *CreateCastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_casts_service_v1_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCastRequest.ProtoReflect.Descriptor instead.
func (*CreateCastRequest) Descriptor() ([]byte, []int) {
	return file_admin_casts_service_v1_messages_proto_rawDescGZIP(), []int{6}
}

func (x *CreateCastRequest) GetCastLabel() string {
	if x != nil {
		return x.CastLabel
	}
	return ""
}

func (x *CreateCastRequest) GetMovieID() int32 {
	if x != nil {
		return x.MovieID
	}
	return 0
}

func (x *CreateCastRequest) GetActorsIDs() []int32 {
	if x != nil {
		return x.ActorsIDs
	}
	return nil
}

type UpdateLabelForCastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label   string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	MovieID int32  `protobuf:"varint,2,opt,name=MovieID,json=movie_id,proto3" json:"MovieID,omitempty"`
}

func (x *UpdateLabelForCastRequest) Reset() {
	*x = UpdateLabelForCastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_casts_service_v1_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLabelForCastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLabelForCastRequest) ProtoMessage() {}

func (x *UpdateLabelForCastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_casts_service_v1_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLabelForCastRequest.ProtoReflect.Descriptor instead.
func (*UpdateLabelForCastRequest) Descriptor() ([]byte, []int) {
	return file_admin_casts_service_v1_messages_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateLabelForCastRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *UpdateLabelForCastRequest) GetMovieID() int32 {
	if x != nil {
		return x.MovieID
	}
	return 0
}

type AddActorsToTheCastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieID   int32   `protobuf:"varint,1,opt,name=MovieID,json=movie_id,proto3" json:"MovieID,omitempty"`
	ActorsIDs []int32 `protobuf:"varint,2,rep,packed,name=ActorsIDs,json=actors_ids,proto3" json:"ActorsIDs,omitempty"`
}

func (x *AddActorsToTheCastRequest) Reset() {
	*x = AddActorsToTheCastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_casts_service_v1_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddActorsToTheCastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddActorsToTheCastRequest) ProtoMessage() {}

func (x *AddActorsToTheCastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_casts_service_v1_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddActorsToTheCastRequest.ProtoReflect.Descriptor instead.
func (*AddActorsToTheCastRequest) Descriptor() ([]byte, []int) {
	return file_admin_casts_service_v1_messages_proto_rawDescGZIP(), []int{8}
}

func (x *AddActorsToTheCastRequest) GetMovieID() int32 {
	if x != nil {
		return x.MovieID
	}
	return 0
}

func (x *AddActorsToTheCastRequest) GetActorsIDs() []int32 {
	if x != nil {
		return x.ActorsIDs
	}
	return nil
}

type RemoveActorsFromTheCastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieID int32 `protobuf:"varint,1,opt,name=MovieID,json=movie_id,proto3" json:"MovieID,omitempty"`
	// use ',' as separator for multiple ids
	ActorsIDs string `protobuf:"bytes,2,opt,name=ActorsIDs,json=actors_ids,proto3" json:"ActorsIDs,omitempty"`
}

func (x *RemoveActorsFromTheCastRequest) Reset() {
	*x = RemoveActorsFromTheCastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_casts_service_v1_messages_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveActorsFromTheCastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveActorsFromTheCastRequest) ProtoMessage() {}

func (x *RemoveActorsFromTheCastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_casts_service_v1_messages_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveActorsFromTheCastRequest.ProtoReflect.Descriptor instead.
func (*RemoveActorsFromTheCastRequest) Descriptor() ([]byte, []int) {
	return file_admin_casts_service_v1_messages_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveActorsFromTheCastRequest) GetMovieID() int32 {
	if x != nil {
		return x.MovieID
	}
	return 0
}

func (x *RemoveActorsFromTheCastRequest) GetActorsIDs() string {
	if x != nil {
		return x.ActorsIDs
	}
	return ""
}

type DeleteCastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId int32 `protobuf:"varint,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
}

func (x *DeleteCastRequest) Reset() {
	*x = DeleteCastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_casts_service_v1_messages_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCastRequest) ProtoMessage() {}

func (x *DeleteCastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_casts_service_v1_messages_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCastRequest.ProtoReflect.Descriptor instead.
func (*DeleteCastRequest) Descriptor() ([]byte, []int) {
	return file_admin_casts_service_v1_messages_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteCastRequest) GetMovieId() int32 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

var File_admin_casts_service_v1_messages_proto protoreflect.FileDescriptor

var file_admin_casts_service_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63,
	0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2c, 0x0a, 0x10,
	0x55, 0x73, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5a, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x05, 0x43, 0x61, 0x73, 0x74, 0x73,
	0x12, 0x3b, 0x0a, 0x05, 0x63, 0x61, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x73, 0x74, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x63, 0x61, 0x73, 0x74, 0x73, 0x1a, 0x53, 0x0a,
	0x0a, 0x43, 0x61, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x61, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x5f, 0x0a, 0x04, 0x43, 0x61, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x61,
	0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x61, 0x73, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x07, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x5f, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x49, 0x44,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x5f,
	0x69, 0x64, 0x73, 0x22, 0x2b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64,
	0x22, 0x5a, 0x0a, 0x18, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x61, 0x73, 0x74, 0x42, 0x79,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x6c, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x61, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x19, 0x0a, 0x07, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x49, 0x44, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x22, 0x4c, 0x0a, 0x19, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x46, 0x6f, 0x72, 0x43, 0x61, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x19, 0x0a,
	0x07, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x55, 0x0a, 0x19, 0x41, 0x64, 0x64, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x54, 0x6f, 0x54, 0x68, 0x65, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x07, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x09, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x22,
	0x5a, 0x0a, 0x1e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x46,
	0x72, 0x6f, 0x6d, 0x54, 0x68, 0x65, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x07, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x09,
	0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x22, 0x2e, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x42, 0x1f, 0x5a, 0x1d, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_casts_service_v1_messages_proto_rawDescOnce sync.Once
	file_admin_casts_service_v1_messages_proto_rawDescData = file_admin_casts_service_v1_messages_proto_rawDesc
)

func file_admin_casts_service_v1_messages_proto_rawDescGZIP() []byte {
	file_admin_casts_service_v1_messages_proto_rawDescOnce.Do(func() {
		file_admin_casts_service_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_casts_service_v1_messages_proto_rawDescData)
	})
	return file_admin_casts_service_v1_messages_proto_rawDescData
}

var file_admin_casts_service_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_admin_casts_service_v1_messages_proto_goTypes = []interface{}{
	(*UserErrorMessage)(nil),               // 0: admin_casts_service.UserErrorMessage
	(*GetCastsRequest)(nil),                // 1: admin_casts_service.GetCastsRequest
	(*Casts)(nil),                          // 2: admin_casts_service.Casts
	(*Cast)(nil),                           // 3: admin_casts_service.Cast
	(*GetCastRequest)(nil),                 // 4: admin_casts_service.GetCastRequest
	(*SearchCastByLabelRequest)(nil),       // 5: admin_casts_service.SearchCastByLabelRequest
	(*CreateCastRequest)(nil),              // 6: admin_casts_service.CreateCastRequest
	(*UpdateLabelForCastRequest)(nil),      // 7: admin_casts_service.UpdateLabelForCastRequest
	(*AddActorsToTheCastRequest)(nil),      // 8: admin_casts_service.AddActorsToTheCastRequest
	(*RemoveActorsFromTheCastRequest)(nil), // 9: admin_casts_service.RemoveActorsFromTheCastRequest
	(*DeleteCastRequest)(nil),              // 10: admin_casts_service.DeleteCastRequest
	nil,                                    // 11: admin_casts_service.Casts.CastsEntry
}
var file_admin_casts_service_v1_messages_proto_depIdxs = []int32{
	11, // 0: admin_casts_service.Casts.casts:type_name -> admin_casts_service.Casts.CastsEntry
	3,  // 1: admin_casts_service.Casts.CastsEntry.value:type_name -> admin_casts_service.Cast
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_admin_casts_service_v1_messages_proto_init() }
func file_admin_casts_service_v1_messages_proto_init() {
	if File_admin_casts_service_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_casts_service_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserErrorMessage); i {
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
		file_admin_casts_service_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCastsRequest); i {
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
		file_admin_casts_service_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Casts); i {
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
		file_admin_casts_service_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cast); i {
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
		file_admin_casts_service_v1_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCastRequest); i {
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
		file_admin_casts_service_v1_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCastByLabelRequest); i {
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
		file_admin_casts_service_v1_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCastRequest); i {
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
		file_admin_casts_service_v1_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLabelForCastRequest); i {
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
		file_admin_casts_service_v1_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddActorsToTheCastRequest); i {
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
		file_admin_casts_service_v1_messages_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveActorsFromTheCastRequest); i {
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
		file_admin_casts_service_v1_messages_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCastRequest); i {
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
			RawDescriptor: file_admin_casts_service_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_casts_service_v1_messages_proto_goTypes,
		DependencyIndexes: file_admin_casts_service_v1_messages_proto_depIdxs,
		MessageInfos:      file_admin_casts_service_v1_messages_proto_msgTypes,
	}.Build()
	File_admin_casts_service_v1_messages_proto = out.File
	file_admin_casts_service_v1_messages_proto_rawDesc = nil
	file_admin_casts_service_v1_messages_proto_goTypes = nil
	file_admin_casts_service_v1_messages_proto_depIdxs = nil
}