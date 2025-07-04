// Code generated by MockGen. DO NOT EDIT.
// Source: .//shared/proto/friends_service/friends_grpc.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	proto "quickflow/shared/proto/friends_service"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockFriendsServiceClient is a mock of FriendsServiceClient interface.
type MockFriendsServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockFriendsServiceClientMockRecorder
}

// MockFriendsServiceClientMockRecorder is the mock recorder for MockFriendsServiceClient.
type MockFriendsServiceClientMockRecorder struct {
	mock *MockFriendsServiceClient
}

// NewMockFriendsServiceClient creates a new mock instance.
func NewMockFriendsServiceClient(ctrl *gomock.Controller) *MockFriendsServiceClient {
	mock := &MockFriendsServiceClient{ctrl: ctrl}
	mock.recorder = &MockFriendsServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFriendsServiceClient) EXPECT() *MockFriendsServiceClientMockRecorder {
	return m.recorder
}

// AcceptFriendRequest mocks base method.
func (m *MockFriendsServiceClient) AcceptFriendRequest(ctx context.Context, in *proto.FriendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AcceptFriendRequest", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptFriendRequest indicates an expected call of AcceptFriendRequest.
func (mr *MockFriendsServiceClientMockRecorder) AcceptFriendRequest(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptFriendRequest", reflect.TypeOf((*MockFriendsServiceClient)(nil).AcceptFriendRequest), varargs...)
}

// DeleteFriend mocks base method.
func (m *MockFriendsServiceClient) DeleteFriend(ctx context.Context, in *proto.FriendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFriend", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFriend indicates an expected call of DeleteFriend.
func (mr *MockFriendsServiceClientMockRecorder) DeleteFriend(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFriend", reflect.TypeOf((*MockFriendsServiceClient)(nil).DeleteFriend), varargs...)
}

// GetFriendsInfo mocks base method.
func (m *MockFriendsServiceClient) GetFriendsInfo(ctx context.Context, in *proto.GetFriendsInfoRequest, opts ...grpc.CallOption) (*proto.GetFriendsInfoResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFriendsInfo", varargs...)
	ret0, _ := ret[0].(*proto.GetFriendsInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFriendsInfo indicates an expected call of GetFriendsInfo.
func (mr *MockFriendsServiceClientMockRecorder) GetFriendsInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriendsInfo", reflect.TypeOf((*MockFriendsServiceClient)(nil).GetFriendsInfo), varargs...)
}

// GetUserRelation mocks base method.
func (m *MockFriendsServiceClient) GetUserRelation(ctx context.Context, in *proto.FriendRequest, opts ...grpc.CallOption) (*proto.RelationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserRelation", varargs...)
	ret0, _ := ret[0].(*proto.RelationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRelation indicates an expected call of GetUserRelation.
func (mr *MockFriendsServiceClientMockRecorder) GetUserRelation(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRelation", reflect.TypeOf((*MockFriendsServiceClient)(nil).GetUserRelation), varargs...)
}

// MarkReadFriendRequest mocks base method.
func (m *MockFriendsServiceClient) MarkReadFriendRequest(ctx context.Context, in *proto.FriendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MarkReadFriendRequest", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkReadFriendRequest indicates an expected call of MarkReadFriendRequest.
func (mr *MockFriendsServiceClientMockRecorder) MarkReadFriendRequest(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkReadFriendRequest", reflect.TypeOf((*MockFriendsServiceClient)(nil).MarkReadFriendRequest), varargs...)
}

// SendFriendRequest mocks base method.
func (m *MockFriendsServiceClient) SendFriendRequest(ctx context.Context, in *proto.FriendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendFriendRequest", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendFriendRequest indicates an expected call of SendFriendRequest.
func (mr *MockFriendsServiceClientMockRecorder) SendFriendRequest(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendFriendRequest", reflect.TypeOf((*MockFriendsServiceClient)(nil).SendFriendRequest), varargs...)
}

// Unfollow mocks base method.
func (m *MockFriendsServiceClient) Unfollow(ctx context.Context, in *proto.FriendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Unfollow", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unfollow indicates an expected call of Unfollow.
func (mr *MockFriendsServiceClientMockRecorder) Unfollow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unfollow", reflect.TypeOf((*MockFriendsServiceClient)(nil).Unfollow), varargs...)
}

// MockFriendsServiceServer is a mock of FriendsServiceServer interface.
type MockFriendsServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockFriendsServiceServerMockRecorder
}

// MockFriendsServiceServerMockRecorder is the mock recorder for MockFriendsServiceServer.
type MockFriendsServiceServerMockRecorder struct {
	mock *MockFriendsServiceServer
}

// NewMockFriendsServiceServer creates a new mock instance.
func NewMockFriendsServiceServer(ctrl *gomock.Controller) *MockFriendsServiceServer {
	mock := &MockFriendsServiceServer{ctrl: ctrl}
	mock.recorder = &MockFriendsServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFriendsServiceServer) EXPECT() *MockFriendsServiceServerMockRecorder {
	return m.recorder
}

// AcceptFriendRequest mocks base method.
func (m *MockFriendsServiceServer) AcceptFriendRequest(arg0 context.Context, arg1 *proto.FriendRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptFriendRequest", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptFriendRequest indicates an expected call of AcceptFriendRequest.
func (mr *MockFriendsServiceServerMockRecorder) AcceptFriendRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptFriendRequest", reflect.TypeOf((*MockFriendsServiceServer)(nil).AcceptFriendRequest), arg0, arg1)
}

// DeleteFriend mocks base method.
func (m *MockFriendsServiceServer) DeleteFriend(arg0 context.Context, arg1 *proto.FriendRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFriend", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFriend indicates an expected call of DeleteFriend.
func (mr *MockFriendsServiceServerMockRecorder) DeleteFriend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFriend", reflect.TypeOf((*MockFriendsServiceServer)(nil).DeleteFriend), arg0, arg1)
}

// GetFriendsInfo mocks base method.
func (m *MockFriendsServiceServer) GetFriendsInfo(arg0 context.Context, arg1 *proto.GetFriendsInfoRequest) (*proto.GetFriendsInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFriendsInfo", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetFriendsInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFriendsInfo indicates an expected call of GetFriendsInfo.
func (mr *MockFriendsServiceServerMockRecorder) GetFriendsInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriendsInfo", reflect.TypeOf((*MockFriendsServiceServer)(nil).GetFriendsInfo), arg0, arg1)
}

// GetUserRelation mocks base method.
func (m *MockFriendsServiceServer) GetUserRelation(arg0 context.Context, arg1 *proto.FriendRequest) (*proto.RelationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRelation", arg0, arg1)
	ret0, _ := ret[0].(*proto.RelationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRelation indicates an expected call of GetUserRelation.
func (mr *MockFriendsServiceServerMockRecorder) GetUserRelation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRelation", reflect.TypeOf((*MockFriendsServiceServer)(nil).GetUserRelation), arg0, arg1)
}

// MarkReadFriendRequest mocks base method.
func (m *MockFriendsServiceServer) MarkReadFriendRequest(arg0 context.Context, arg1 *proto.FriendRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkReadFriendRequest", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkReadFriendRequest indicates an expected call of MarkReadFriendRequest.
func (mr *MockFriendsServiceServerMockRecorder) MarkReadFriendRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkReadFriendRequest", reflect.TypeOf((*MockFriendsServiceServer)(nil).MarkReadFriendRequest), arg0, arg1)
}

// SendFriendRequest mocks base method.
func (m *MockFriendsServiceServer) SendFriendRequest(arg0 context.Context, arg1 *proto.FriendRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendFriendRequest", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendFriendRequest indicates an expected call of SendFriendRequest.
func (mr *MockFriendsServiceServerMockRecorder) SendFriendRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendFriendRequest", reflect.TypeOf((*MockFriendsServiceServer)(nil).SendFriendRequest), arg0, arg1)
}

// Unfollow mocks base method.
func (m *MockFriendsServiceServer) Unfollow(arg0 context.Context, arg1 *proto.FriendRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unfollow", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unfollow indicates an expected call of Unfollow.
func (mr *MockFriendsServiceServerMockRecorder) Unfollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unfollow", reflect.TypeOf((*MockFriendsServiceServer)(nil).Unfollow), arg0, arg1)
}

// mustEmbedUnimplementedFriendsServiceServer mocks base method.
func (m *MockFriendsServiceServer) mustEmbedUnimplementedFriendsServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedFriendsServiceServer")
}

// mustEmbedUnimplementedFriendsServiceServer indicates an expected call of mustEmbedUnimplementedFriendsServiceServer.
func (mr *MockFriendsServiceServerMockRecorder) mustEmbedUnimplementedFriendsServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedFriendsServiceServer", reflect.TypeOf((*MockFriendsServiceServer)(nil).mustEmbedUnimplementedFriendsServiceServer))
}

// MockUnsafeFriendsServiceServer is a mock of UnsafeFriendsServiceServer interface.
type MockUnsafeFriendsServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeFriendsServiceServerMockRecorder
}

// MockUnsafeFriendsServiceServerMockRecorder is the mock recorder for MockUnsafeFriendsServiceServer.
type MockUnsafeFriendsServiceServerMockRecorder struct {
	mock *MockUnsafeFriendsServiceServer
}

// NewMockUnsafeFriendsServiceServer creates a new mock instance.
func NewMockUnsafeFriendsServiceServer(ctrl *gomock.Controller) *MockUnsafeFriendsServiceServer {
	mock := &MockUnsafeFriendsServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeFriendsServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeFriendsServiceServer) EXPECT() *MockUnsafeFriendsServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedFriendsServiceServer mocks base method.
func (m *MockUnsafeFriendsServiceServer) mustEmbedUnimplementedFriendsServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedFriendsServiceServer")
}

// mustEmbedUnimplementedFriendsServiceServer indicates an expected call of mustEmbedUnimplementedFriendsServiceServer.
func (mr *MockUnsafeFriendsServiceServerMockRecorder) mustEmbedUnimplementedFriendsServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedFriendsServiceServer", reflect.TypeOf((*MockUnsafeFriendsServiceServer)(nil).mustEmbedUnimplementedFriendsServiceServer))
}
