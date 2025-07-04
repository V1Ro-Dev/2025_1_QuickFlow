// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.7
// source: friends.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FriendRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ReceiverId    string                 `protobuf:"bytes,2,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FriendRequest) Reset() {
	*x = FriendRequest{}
	mi := &file_friends_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRequest) ProtoMessage() {}

func (x *FriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRequest.ProtoReflect.Descriptor instead.
func (*FriendRequest) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{0}
}

func (x *FriendRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FriendRequest) GetReceiverId() string {
	if x != nil {
		return x.ReceiverId
	}
	return ""
}

type GetFriendInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Firstname     string                 `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname      string                 `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	AvatarUrl     string                 `protobuf:"bytes,5,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	University    string                 `protobuf:"bytes,6,opt,name=university,proto3" json:"university,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFriendInfo) Reset() {
	*x = GetFriendInfo{}
	mi := &file_friends_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFriendInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendInfo) ProtoMessage() {}

func (x *GetFriendInfo) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendInfo.ProtoReflect.Descriptor instead.
func (*GetFriendInfo) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{1}
}

func (x *GetFriendInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetFriendInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetFriendInfo) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *GetFriendInfo) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *GetFriendInfo) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *GetFriendInfo) GetUniversity() string {
	if x != nil {
		return x.University
	}
	return ""
}

type GetFriendsInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit         string                 `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        string                 `protobuf:"bytes,3,opt,name=offset,proto3" json:"offset,omitempty"`
	ReqType       string                 `protobuf:"bytes,4,opt,name=reqType,proto3" json:"reqType,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFriendsInfoRequest) Reset() {
	*x = GetFriendsInfoRequest{}
	mi := &file_friends_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFriendsInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendsInfoRequest) ProtoMessage() {}

func (x *GetFriendsInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendsInfoRequest.ProtoReflect.Descriptor instead.
func (*GetFriendsInfoRequest) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{2}
}

func (x *GetFriendsInfoRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetFriendsInfoRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *GetFriendsInfoRequest) GetOffset() string {
	if x != nil {
		return x.Offset
	}
	return ""
}

func (x *GetFriendsInfoRequest) GetReqType() string {
	if x != nil {
		return x.ReqType
	}
	return ""
}

type GetFriendsInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Friends       []*GetFriendInfo       `protobuf:"bytes,1,rep,name=friends,proto3" json:"friends,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFriendsInfoResponse) Reset() {
	*x = GetFriendsInfoResponse{}
	mi := &file_friends_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFriendsInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendsInfoResponse) ProtoMessage() {}

func (x *GetFriendsInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendsInfoResponse.ProtoReflect.Descriptor instead.
func (*GetFriendsInfoResponse) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{3}
}

func (x *GetFriendsInfoResponse) GetFriends() []*GetFriendInfo {
	if x != nil {
		return x.Friends
	}
	return nil
}

func (x *GetFriendsInfoResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type IsRelationExistsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsExists      bool                   `protobuf:"varint,1,opt,name=is_exists,json=isExists,proto3" json:"is_exists,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsRelationExistsResponse) Reset() {
	*x = IsRelationExistsResponse{}
	mi := &file_friends_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsRelationExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsRelationExistsResponse) ProtoMessage() {}

func (x *IsRelationExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsRelationExistsResponse.ProtoReflect.Descriptor instead.
func (*IsRelationExistsResponse) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{4}
}

func (x *IsRelationExistsResponse) GetIsExists() bool {
	if x != nil {
		return x.IsExists
	}
	return false
}

type RelationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Relation      string                 `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RelationResponse) Reset() {
	*x = RelationResponse{}
	mi := &file_friends_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationResponse) ProtoMessage() {}

func (x *RelationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationResponse.ProtoReflect.Descriptor instead.
func (*RelationResponse) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{5}
}

func (x *RelationResponse) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

var File_friends_proto protoreflect.FileDescriptor

const file_friends_proto_rawDesc = "" +
	"\n" +
	"\rfriends.proto\x12\x0ffriends_service\x1a\x1bgoogle/protobuf/empty.proto\"I\n" +
	"\rFriendRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1f\n" +
	"\vreceiver_id\x18\x02 \x01(\tR\n" +
	"receiverId\"\xb4\x01\n" +
	"\rGetFriendInfo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x1c\n" +
	"\tfirstname\x18\x03 \x01(\tR\tfirstname\x12\x1a\n" +
	"\blastname\x18\x04 \x01(\tR\blastname\x12\x1d\n" +
	"\n" +
	"avatar_url\x18\x05 \x01(\tR\tavatarUrl\x12\x1e\n" +
	"\n" +
	"university\x18\x06 \x01(\tR\n" +
	"university\"x\n" +
	"\x15GetFriendsInfoRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\tR\x05limit\x12\x16\n" +
	"\x06offset\x18\x03 \x01(\tR\x06offset\x12\x18\n" +
	"\areqType\x18\x04 \x01(\tR\areqType\"s\n" +
	"\x16GetFriendsInfoResponse\x128\n" +
	"\afriends\x18\x01 \x03(\v2\x1e.friends_service.GetFriendInfoR\afriends\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"7\n" +
	"\x18IsRelationExistsResponse\x12\x1b\n" +
	"\tis_exists\x18\x01 \x01(\bR\bisExists\".\n" +
	"\x10RelationResponse\x12\x1a\n" +
	"\brelation\x18\x01 \x01(\tR\brelation2\xc2\x04\n" +
	"\x0eFriendsService\x12a\n" +
	"\x0eGetFriendsInfo\x12&.friends_service.GetFriendsInfoRequest\x1a'.friends_service.GetFriendsInfoResponse\x12K\n" +
	"\x11SendFriendRequest\x12\x1e.friends_service.FriendRequest\x1a\x16.google.protobuf.Empty\x12M\n" +
	"\x13AcceptFriendRequest\x12\x1e.friends_service.FriendRequest\x1a\x16.google.protobuf.Empty\x12B\n" +
	"\bUnfollow\x12\x1e.friends_service.FriendRequest\x1a\x16.google.protobuf.Empty\x12F\n" +
	"\fDeleteFriend\x12\x1e.friends_service.FriendRequest\x1a\x16.google.protobuf.Empty\x12T\n" +
	"\x0fGetUserRelation\x12\x1e.friends_service.FriendRequest\x1a!.friends_service.RelationResponse\x12O\n" +
	"\x15MarkReadFriendRequest\x12\x1e.friends_service.FriendRequest\x1a\x16.google.protobuf.EmptyB8Z6quickflow/friends_service/internal/delivery/grpc/protob\x06proto3"

var (
	file_friends_proto_rawDescOnce sync.Once
	file_friends_proto_rawDescData []byte
)

func file_friends_proto_rawDescGZIP() []byte {
	file_friends_proto_rawDescOnce.Do(func() {
		file_friends_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_friends_proto_rawDesc), len(file_friends_proto_rawDesc)))
	})
	return file_friends_proto_rawDescData
}

var file_friends_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_friends_proto_goTypes = []any{
	(*FriendRequest)(nil),            // 0: friends_service.FriendRequest
	(*GetFriendInfo)(nil),            // 1: friends_service.GetFriendInfo
	(*GetFriendsInfoRequest)(nil),    // 2: friends_service.GetFriendsInfoRequest
	(*GetFriendsInfoResponse)(nil),   // 3: friends_service.GetFriendsInfoResponse
	(*IsRelationExistsResponse)(nil), // 4: friends_service.IsRelationExistsResponse
	(*RelationResponse)(nil),         // 5: friends_service.RelationResponse
	(*emptypb.Empty)(nil),            // 6: google.protobuf.Empty
}
var file_friends_proto_depIdxs = []int32{
	1, // 0: friends_service.GetFriendsInfoResponse.friends:type_name -> friends_service.GetFriendInfo
	2, // 1: friends_service.FriendsService.GetFriendsInfo:input_type -> friends_service.GetFriendsInfoRequest
	0, // 2: friends_service.FriendsService.SendFriendRequest:input_type -> friends_service.FriendRequest
	0, // 3: friends_service.FriendsService.AcceptFriendRequest:input_type -> friends_service.FriendRequest
	0, // 4: friends_service.FriendsService.Unfollow:input_type -> friends_service.FriendRequest
	0, // 5: friends_service.FriendsService.DeleteFriend:input_type -> friends_service.FriendRequest
	0, // 6: friends_service.FriendsService.GetUserRelation:input_type -> friends_service.FriendRequest
	0, // 7: friends_service.FriendsService.MarkReadFriendRequest:input_type -> friends_service.FriendRequest
	3, // 8: friends_service.FriendsService.GetFriendsInfo:output_type -> friends_service.GetFriendsInfoResponse
	6, // 9: friends_service.FriendsService.SendFriendRequest:output_type -> google.protobuf.Empty
	6, // 10: friends_service.FriendsService.AcceptFriendRequest:output_type -> google.protobuf.Empty
	6, // 11: friends_service.FriendsService.Unfollow:output_type -> google.protobuf.Empty
	6, // 12: friends_service.FriendsService.DeleteFriend:output_type -> google.protobuf.Empty
	5, // 13: friends_service.FriendsService.GetUserRelation:output_type -> friends_service.RelationResponse
	6, // 14: friends_service.FriendsService.MarkReadFriendRequest:output_type -> google.protobuf.Empty
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_friends_proto_init() }
func file_friends_proto_init() {
	if File_friends_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_friends_proto_rawDesc), len(file_friends_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_friends_proto_goTypes,
		DependencyIndexes: file_friends_proto_depIdxs,
		MessageInfos:      file_friends_proto_msgTypes,
	}.Build()
	File_friends_proto = out.File
	file_friends_proto_goTypes = nil
	file_friends_proto_depIdxs = nil
}
