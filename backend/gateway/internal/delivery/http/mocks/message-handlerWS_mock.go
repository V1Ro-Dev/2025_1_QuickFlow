// Code generated by MockGen. DO NOT EDIT.
// Source: .//gateway/internal/delivery/http/message-handlerWS.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	json "encoding/json"
	models "quickflow/shared/models"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	websocket "github.com/gorilla/websocket"
)

// MockMessageService is a mock of MessageService interface.
type MockMessageService struct {
	ctrl     *gomock.Controller
	recorder *MockMessageServiceMockRecorder
}

// MockMessageServiceMockRecorder is the mock recorder for MockMessageService.
type MockMessageServiceMockRecorder struct {
	mock *MockMessageService
}

// NewMockMessageService creates a new mock instance.
func NewMockMessageService(ctrl *gomock.Controller) *MockMessageService {
	mock := &MockMessageService{ctrl: ctrl}
	mock.recorder = &MockMessageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageService) EXPECT() *MockMessageServiceMockRecorder {
	return m.recorder
}

// DeleteMessage mocks base method.
func (m *MockMessageService) DeleteMessage(ctx context.Context, messageId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessage", ctx, messageId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessage indicates an expected call of DeleteMessage.
func (mr *MockMessageServiceMockRecorder) DeleteMessage(ctx, messageId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*MockMessageService)(nil).DeleteMessage), ctx, messageId)
}

// GetLastReadTs mocks base method.
func (m *MockMessageService) GetLastReadTs(ctx context.Context, chatId, userId uuid.UUID) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastReadTs", ctx, chatId, userId)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastReadTs indicates an expected call of GetLastReadTs.
func (mr *MockMessageServiceMockRecorder) GetLastReadTs(ctx, chatId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastReadTs", reflect.TypeOf((*MockMessageService)(nil).GetLastReadTs), ctx, chatId, userId)
}

// GetMessageById mocks base method.
func (m *MockMessageService) GetMessageById(ctx context.Context, messageId uuid.UUID) (*models.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageById", ctx, messageId)
	ret0, _ := ret[0].(*models.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageById indicates an expected call of GetMessageById.
func (mr *MockMessageServiceMockRecorder) GetMessageById(ctx, messageId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageById", reflect.TypeOf((*MockMessageService)(nil).GetMessageById), ctx, messageId)
}

// GetMessagesForChat mocks base method.
func (m *MockMessageService) GetMessagesForChat(ctx context.Context, chatId uuid.UUID, numMessages int, timestamp time.Time, userId uuid.UUID) ([]*models.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessagesForChat", ctx, chatId, numMessages, timestamp, userId)
	ret0, _ := ret[0].([]*models.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessagesForChat indicates an expected call of GetMessagesForChat.
func (mr *MockMessageServiceMockRecorder) GetMessagesForChat(ctx, chatId, numMessages, timestamp, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessagesForChat", reflect.TypeOf((*MockMessageService)(nil).GetMessagesForChat), ctx, chatId, numMessages, timestamp, userId)
}

// GetNumUnreadMessages mocks base method.
func (m *MockMessageService) GetNumUnreadMessages(ctx context.Context, chatId, userId uuid.UUID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumUnreadMessages", ctx, chatId, userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNumUnreadMessages indicates an expected call of GetNumUnreadMessages.
func (mr *MockMessageServiceMockRecorder) GetNumUnreadMessages(ctx, chatId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumUnreadMessages", reflect.TypeOf((*MockMessageService)(nil).GetNumUnreadMessages), ctx, chatId, userId)
}

// SendMessage mocks base method.
func (m *MockMessageService) SendMessage(ctx context.Context, message *models.Message, userId uuid.UUID) (*models.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", ctx, message, userId)
	ret0, _ := ret[0].(*models.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockMessageServiceMockRecorder) SendMessage(ctx, message, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockMessageService)(nil).SendMessage), ctx, message, userId)
}

// UpdateLastReadTs mocks base method.
func (m *MockMessageService) UpdateLastReadTs(ctx context.Context, chatId, userId uuid.UUID, timestamp time.Time, userAuthId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastReadTs", ctx, chatId, userId, timestamp, userAuthId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLastReadTs indicates an expected call of UpdateLastReadTs.
func (mr *MockMessageServiceMockRecorder) UpdateLastReadTs(ctx, chatId, userId, timestamp, userAuthId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastReadTs", reflect.TypeOf((*MockMessageService)(nil).UpdateLastReadTs), ctx, chatId, userId, timestamp, userAuthId)
}

// MockIWebSocketConnectionManager is a mock of IWebSocketConnectionManager interface.
type MockIWebSocketConnectionManager struct {
	ctrl     *gomock.Controller
	recorder *MockIWebSocketConnectionManagerMockRecorder
}

// MockIWebSocketConnectionManagerMockRecorder is the mock recorder for MockIWebSocketConnectionManager.
type MockIWebSocketConnectionManagerMockRecorder struct {
	mock *MockIWebSocketConnectionManager
}

// NewMockIWebSocketConnectionManager creates a new mock instance.
func NewMockIWebSocketConnectionManager(ctrl *gomock.Controller) *MockIWebSocketConnectionManager {
	mock := &MockIWebSocketConnectionManager{ctrl: ctrl}
	mock.recorder = &MockIWebSocketConnectionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWebSocketConnectionManager) EXPECT() *MockIWebSocketConnectionManagerMockRecorder {
	return m.recorder
}

// AddConnection mocks base method.
func (m *MockIWebSocketConnectionManager) AddConnection(userId uuid.UUID, conn *websocket.Conn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddConnection", userId, conn)
}

// AddConnection indicates an expected call of AddConnection.
func (mr *MockIWebSocketConnectionManagerMockRecorder) AddConnection(userId, conn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddConnection", reflect.TypeOf((*MockIWebSocketConnectionManager)(nil).AddConnection), userId, conn)
}

// IsConnected mocks base method.
func (m *MockIWebSocketConnectionManager) IsConnected(userId uuid.UUID) (*websocket.Conn, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConnected", userId)
	ret0, _ := ret[0].(*websocket.Conn)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// IsConnected indicates an expected call of IsConnected.
func (mr *MockIWebSocketConnectionManagerMockRecorder) IsConnected(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConnected", reflect.TypeOf((*MockIWebSocketConnectionManager)(nil).IsConnected), userId)
}

// RemoveAndCloseConnection mocks base method.
func (m *MockIWebSocketConnectionManager) RemoveAndCloseConnection(userId uuid.UUID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveAndCloseConnection", userId)
}

// RemoveAndCloseConnection indicates an expected call of RemoveAndCloseConnection.
func (mr *MockIWebSocketConnectionManagerMockRecorder) RemoveAndCloseConnection(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAndCloseConnection", reflect.TypeOf((*MockIWebSocketConnectionManager)(nil).RemoveAndCloseConnection), userId)
}

// MockIWebSocketRouter is a mock of IWebSocketRouter interface.
type MockIWebSocketRouter struct {
	ctrl     *gomock.Controller
	recorder *MockIWebSocketRouterMockRecorder
}

// MockIWebSocketRouterMockRecorder is the mock recorder for MockIWebSocketRouter.
type MockIWebSocketRouterMockRecorder struct {
	mock *MockIWebSocketRouter
}

// NewMockIWebSocketRouter creates a new mock instance.
func NewMockIWebSocketRouter(ctrl *gomock.Controller) *MockIWebSocketRouter {
	mock := &MockIWebSocketRouter{ctrl: ctrl}
	mock.recorder = &MockIWebSocketRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWebSocketRouter) EXPECT() *MockIWebSocketRouterMockRecorder {
	return m.recorder
}

// RegisterHandler mocks base method.

type CommandHandler func(ctx context.Context, user models.User, payload json.RawMessage) error

func (m *MockIWebSocketRouter) RegisterHandler(command string, handler CommandHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterHandler", command, handler)
}

// RegisterHandler indicates an expected call of RegisterHandler.
func (mr *MockIWebSocketRouterMockRecorder) RegisterHandler(command, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterHandler", reflect.TypeOf((*MockIWebSocketRouter)(nil).RegisterHandler), command, handler)
}

// Route mocks base method.
func (m *MockIWebSocketRouter) Route(ctx context.Context, command string, user models.User, payload json.RawMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Route", ctx, command, user, payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// Route indicates an expected call of Route.
func (mr *MockIWebSocketRouterMockRecorder) Route(ctx, command, user, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Route", reflect.TypeOf((*MockIWebSocketRouter)(nil).Route), ctx, command, user, payload)
}
