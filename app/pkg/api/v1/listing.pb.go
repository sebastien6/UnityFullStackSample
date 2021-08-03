// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: listing.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_listing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_listing_proto_rawDescGZIP(), []int{0}
}

func (x *ID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Game *Game  `protobuf:"bytes,2,opt,name=game,proto3" json:"game,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_listing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_listing_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateReq) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Category            string    `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Title               string    `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Subtitle            string    `protobuf:"bytes,4,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Description         string    `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Images              []*Images `protobuf:"bytes,6,rep,name=images,proto3" json:"images,omitempty"`
	Type                int32     `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"`
	Tags                []string  `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	Author              string    `protobuf:"bytes,9,opt,name=author,proto3" json:"author,omitempty"`
	ReplayBundleUrlJson string    `protobuf:"bytes,10,opt,name=replayBundleUrlJson,proto3" json:"replayBundleUrlJson,omitempty"`
	Duration            float64   `protobuf:"fixed64,11,opt,name=duration,proto3" json:"duration,omitempty"`
	IsDownloadable      bool      `protobuf:"varint,12,opt,name=isDownloadable,proto3" json:"isDownloadable,omitempty"`
	IsStreamable        bool      `protobuf:"varint,13,opt,name=isStreamable,proto3" json:"isStreamable,omitempty"`
	Version             string    `protobuf:"bytes,14,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_listing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_listing_proto_rawDescGZIP(), []int{2}
}

func (x *Game) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Game) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Game) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Game) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

func (x *Game) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Game) GetImages() []*Images {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *Game) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Game) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Game) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Game) GetReplayBundleUrlJson() string {
	if x != nil {
		return x.ReplayBundleUrlJson
	}
	return ""
}

func (x *Game) GetDuration() float64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Game) GetIsDownloadable() bool {
	if x != nil {
		return x.IsDownloadable
	}
	return false
}

func (x *Game) GetIsStreamable() bool {
	if x != nil {
		return x.IsStreamable
	}
	return false
}

func (x *Game) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Images struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Type int32  `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Images) Reset() {
	*x = Images{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Images) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Images) ProtoMessage() {}

func (x *Images) ProtoReflect() protoreflect.Message {
	mi := &file_listing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Images.ProtoReflect.Descriptor instead.
func (*Images) Descriptor() ([]byte, []int) {
	return file_listing_proto_rawDescGZIP(), []int{3}
}

func (x *Images) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Images) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Images) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type StreamGames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listings *Game `protobuf:"bytes,1,opt,name=listings,proto3" json:"listings,omitempty"`
}

func (x *StreamGames) Reset() {
	*x = StreamGames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamGames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamGames) ProtoMessage() {}

func (x *StreamGames) ProtoReflect() protoreflect.Message {
	mi := &file_listing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamGames.ProtoReflect.Descriptor instead.
func (*StreamGames) Descriptor() ([]byte, []int) {
	return file_listing_proto_rawDescGZIP(), []int{4}
}

func (x *StreamGames) GetListings() *Game {
	if x != nil {
		return x.Listings
	}
	return nil
}

type Games struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listings []*Game `protobuf:"bytes,1,rep,name=listings,proto3" json:"listings,omitempty"`
}

func (x *Games) Reset() {
	*x = Games{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Games) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Games) ProtoMessage() {}

func (x *Games) ProtoReflect() protoreflect.Message {
	mi := &file_listing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Games.ProtoReflect.Descriptor instead.
func (*Games) Descriptor() ([]byte, []int) {
	return file_listing_proto_rawDescGZIP(), []int{5}
}

func (x *Games) GetListings() []*Game {
	if x != nil {
		return x.Listings
	}
	return nil
}

var File_listing_proto protoreflect.FileDescriptor

var file_listing_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1d, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x84, 0x04, 0x0a, 0x04, 0x47,
	0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x23, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03, 0x52, 0x08, 0x73, 0x75, 0x62,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2a, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x22, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0e,
	0xfa, 0x42, 0x0b, 0x92, 0x01, 0x08, 0x22, 0x04, 0x72, 0x02, 0x10, 0x03, 0x28, 0x01, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03, 0x52, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x3d, 0x0a, 0x13, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x4a, 0x73, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x88, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52,
	0x13, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x55, 0x72, 0x6c,
	0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x26, 0x0a, 0x0e, 0x69, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x69, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x53, 0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x03,
	0x18, 0x14, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x88, 0x01, 0x01, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3b, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x47, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x22, 0x35, 0x0a, 0x05, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x08,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x32, 0xf0, 0x03, 0x0a, 0x0c, 0x47,
	0x61, 0x6d, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa4, 0x01, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x11, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x73, 0x22, 0x6d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x92, 0x41, 0x55, 0x12, 0x1d, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x61, 0x20, 0x66, 0x75, 0x6c, 0x6c, 0x20, 0x6c,
	0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x34, 0x55, 0x73,
	0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20,
	0x74, 0x6f, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x66, 0x75, 0x6c, 0x6c, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x67, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x8c, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x47,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x61, 0x6d, 0x65, 0x1a, 0x10, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x22, 0x58, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22,
	0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x3a, 0x01, 0x2a,
	0x92, 0x41, 0x3e, 0x12, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65,
	0x77, 0x20, 0x67, 0x61, 0x6d, 0x65, 0x1a, 0x29, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x67, 0x61, 0x6d,
	0x65, 0x12, 0xa9, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x11, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73,
	0x3a, 0x01, 0x2a, 0x92, 0x41, 0x57, 0x12, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x20, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x3e, 0x55,
	0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x65, 0x20, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x61, 0x20, 0x62,
	0x75, 0x6c, 0x6b, 0x20, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x9c, 0x01,
	0x5a, 0x0c, 0x2e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x92, 0x41,
	0x8a, 0x01, 0x12, 0x34, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x70,
	0x69, 0x22, 0x20, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x15, 0x73, 0x65,
	0x62, 0x2e, 0x64, 0x75, 0x72, 0x61, 0x6e, 0x64, 0x36, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x32, 0x14,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x2d, 0x6e, 0x64,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x14, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x2d, 0x6e, 0x64, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_listing_proto_rawDescOnce sync.Once
	file_listing_proto_rawDescData = file_listing_proto_rawDesc
)

func file_listing_proto_rawDescGZIP() []byte {
	file_listing_proto_rawDescOnce.Do(func() {
		file_listing_proto_rawDescData = protoimpl.X.CompressGZIP(file_listing_proto_rawDescData)
	})
	return file_listing_proto_rawDescData
}

var file_listing_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_listing_proto_goTypes = []interface{}{
	(*ID)(nil),            // 0: listing.v1.ID
	(*UpdateReq)(nil),     // 1: listing.v1.UpdateReq
	(*Game)(nil),          // 2: listing.v1.Game
	(*Images)(nil),        // 3: listing.v1.Images
	(*StreamGames)(nil),   // 4: listing.v1.StreamGames
	(*Games)(nil),         // 5: listing.v1.Games
	(*emptypb.Empty)(nil), // 6: google.protobuf.Empty
}
var file_listing_proto_depIdxs = []int32{
	2, // 0: listing.v1.UpdateReq.game:type_name -> listing.v1.Game
	3, // 1: listing.v1.Game.images:type_name -> listing.v1.Images
	2, // 2: listing.v1.StreamGames.listings:type_name -> listing.v1.Game
	2, // 3: listing.v1.Games.listings:type_name -> listing.v1.Game
	6, // 4: listing.v1.GamesService.GetGames:input_type -> google.protobuf.Empty
	2, // 5: listing.v1.GamesService.RegisterGame:input_type -> listing.v1.Game
	5, // 6: listing.v1.GamesService.RegisterGames:input_type -> listing.v1.Games
	5, // 7: listing.v1.GamesService.GetGames:output_type -> listing.v1.Games
	2, // 8: listing.v1.GamesService.RegisterGame:output_type -> listing.v1.Game
	5, // 9: listing.v1.GamesService.RegisterGames:output_type -> listing.v1.Games
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_listing_proto_init() }
func file_listing_proto_init() {
	if File_listing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_listing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_listing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_listing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
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
		file_listing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Images); i {
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
		file_listing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamGames); i {
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
		file_listing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Games); i {
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
			RawDescriptor: file_listing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_listing_proto_goTypes,
		DependencyIndexes: file_listing_proto_depIdxs,
		MessageInfos:      file_listing_proto_msgTypes,
	}.Build()
	File_listing_proto = out.File
	file_listing_proto_rawDesc = nil
	file_listing_proto_goTypes = nil
	file_listing_proto_depIdxs = nil
}
