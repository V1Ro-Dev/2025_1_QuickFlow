package http

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	"quickflow/gateway/internal/delivery/http/mocks"
	"quickflow/shared/models"
)

func TestGetUserChats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокирование зависимостей
	mockChatUseCase := mocks.NewMockChatUseCase(ctrl)
	mockProfileUseCase := mocks.NewMockProfileUseCase(ctrl)
	mockMessageService := mocks.NewMockMessageService(ctrl)
	mockConnService := mocks.NewMockIWebSocketConnectionManager(ctrl)

	// Инициализация обработчика
	handler := NewChatHandler(mockChatUseCase, mockProfileUseCase, mockMessageService, mockConnService)

	// Генерация тестовых данных
	userID := uuid.New()
	username := "testuser"
	chatsCount := 5
	user := models.User{Id: userID, Username: username}

	// Создание мока для запроса
	req := httptest.NewRequest("GET", "/api/chats?chats_count=5", nil)
	req = req.WithContext(context.WithValue(req.Context(), "user", user))

	// Мокирование ответа
	w := httptest.NewRecorder()

	// Мокирование вызова методов
	mockChatUseCase.EXPECT().GetUserChats(gomock.Any(), userID, chatsCount, gomock.Any()).Return([]models.Chat{}, nil).AnyTimes()
	mockProfileUseCase.EXPECT().GetPublicUsersInfo(gomock.Any(), gomock.Any()).Return([]models.PublicUserInfo{}, nil).AnyTimes()
	mockMessageService.EXPECT().GetNumUnreadMessages(gomock.Any(), gomock.Any(), userID).Return(0, nil).AnyTimes()

	// Вызов обработчика
	handler.GetUserChats(w, req)

}

func TestGetUserChats_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокирование зависимостей
	mockChatUseCase := mocks.NewMockChatUseCase(ctrl)
	mockProfileUseCase := mocks.NewMockProfileUseCase(ctrl)
	mockMessageService := mocks.NewMockMessageService(ctrl)
	mockConnService := mocks.NewMockIWebSocketConnectionManager(ctrl)

	// Инициализация обработчика
	handler := NewChatHandler(mockChatUseCase, mockProfileUseCase, mockMessageService, mockConnService)

	// Генерация тестовых данных
	userID := uuid.New()
	username := "testuser"
	user := models.User{Id: userID, Username: username}

	// Создание мока для запроса
	req := httptest.NewRequest("GET", "/api/chats?chats_count=5", nil)
	req = req.WithContext(context.WithValue(req.Context(), "user", user))

	// Мокирование ответа
	w := httptest.NewRecorder()

	// Мокирование ошибок
	mockChatUseCase.EXPECT().GetUserChats(gomock.Any(), userID, gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))

	// Вызов обработчика
	handler.GetUserChats(w, req)
}

func TestGetNumUnreadChats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокирование зависимостей
	mockChatUseCase := mocks.NewMockChatUseCase(ctrl)

	// Инициализация обработчика
	handler := NewChatHandler(mockChatUseCase, nil, nil, nil)

	// Генерация тестовых данных
	userID := uuid.New()
	user := models.User{Id: userID}

	// Создание мока для запроса
	req := httptest.NewRequest("GET", "/api/chats/unread", nil)
	req = req.WithContext(context.WithValue(req.Context(), "user", user))

	// Мокирование ответа
	w := httptest.NewRecorder()

	// Мокирование вызова методов
	mockChatUseCase.EXPECT().GetNumUnreadChats(gomock.Any(), userID).Return(5, nil)

	// Вызов обработчика
	handler.GetNumUnreadChats(w, req)
}

func TestGetNumUnreadChats_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокирование зависимостей
	mockChatUseCase := mocks.NewMockChatUseCase(ctrl)

	// Инициализация обработчика
	handler := NewChatHandler(mockChatUseCase, nil, nil, nil)

	// Генерация тестовых данных
	userID := uuid.New()
	user := models.User{Id: userID}

	// Создание мока для запроса
	req := httptest.NewRequest("GET", "/api/chats/unread", nil)
	req = req.WithContext(context.WithValue(req.Context(), "user", user))

	// Мокирование ответа
	w := httptest.NewRecorder()

	// Мокирование ошибок
	mockChatUseCase.EXPECT().GetNumUnreadChats(gomock.Any(), userID).Return(0, errors.New("error"))

	// Вызов обработчика
	handler.GetNumUnreadChats(w, req)
}
