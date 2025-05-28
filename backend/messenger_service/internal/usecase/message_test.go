package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	messenger_errors "quickflow/messenger_service/internal/errors"
	"quickflow/messenger_service/internal/usecase"
	"quickflow/messenger_service/internal/usecase/mocks"
	"quickflow/shared/models"
)

func TestSaveMessage_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	messageRepo := mocks.NewMockMessageRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	chatRepo := mocks.NewMockChatRepository(ctrl)
	validator := mocks.NewMockMessageValidator(ctrl)

	// Подготовка тестовых данных
	message := models.Message{
		ID:         uuid.New(),
		Text:       "Hello, World!",
		SenderID:   uuid.New(),
		ReceiverID: uuid.New(),
		ChatID:     uuid.New(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Ожидания для моков
	validator.EXPECT().ValidateMessage(message).Return(nil)
	messageRepo.EXPECT().SaveMessage(context.Background(), message).Return(nil)
	messageRepo.EXPECT().GetMessageById(context.Background(), message.ID).Return(message, nil)

	// Создаем сервис
	messageService := usecase.NewMessageService(messageRepo, fileRepo, chatRepo, validator)

	// Вызов метода
	savedMessage, err := messageService.SaveMessage(context.Background(), message)

	// Проверки
	assert.NoError(t, err)
	assert.Equal(t, message, *savedMessage)
}

func TestSaveMessage_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	messageRepo := mocks.NewMockMessageRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	chatRepo := mocks.NewMockChatRepository(ctrl)
	validator := mocks.NewMockMessageValidator(ctrl)

	// Подготовка тестовых данных
	message := models.Message{
		ID:         uuid.New(),
		Text:       "",
		SenderID:   uuid.New(),
		ReceiverID: uuid.New(),
		ChatID:     uuid.New(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Ожидания для моков
	validator.EXPECT().ValidateMessage(message).Return(errors.New("validation error"))

	// Создаем сервис
	messageService := usecase.NewMessageService(messageRepo, fileRepo, chatRepo, validator)

	// Вызов метода
	savedMessage, err := messageService.SaveMessage(context.Background(), message)

	// Проверки
	assert.Error(t, err)
	assert.Nil(t, savedMessage)
}

func TestGetMessagesForChatOlder_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	messageRepo := mocks.NewMockMessageRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	chatRepo := mocks.NewMockChatRepository(ctrl)
	validator := mocks.NewMockMessageValidator(ctrl)

	// Подготовка тестовых данных
	chatId := uuid.New()
	userId := uuid.New()
	messages := []models.Message{
		{
			ID:         uuid.New(),
			Text:       "Old message",
			SenderID:   userId,
			ReceiverID: uuid.New(),
			ChatID:     chatId,
			CreatedAt:  time.Now().Add(-1 * time.Hour),
			UpdatedAt:  time.Now().Add(-1 * time.Hour),
		},
	}

	// Ожидания для моков
	chatRepo.EXPECT().IsParticipant(context.Background(), chatId, userId).Return(true, nil)
	messageRepo.EXPECT().GetMessagesForChatOlder(context.Background(), chatId, 5, gomock.Any()).Return(messages, nil)

	// Создаем сервис
	messageService := usecase.NewMessageService(messageRepo, fileRepo, chatRepo, validator)

	// Вызов метода
	resultMessages, err := messageService.GetMessagesForChatOlder(context.Background(), chatId, userId, 5, time.Now())

	// Проверки
	assert.NoError(t, err)
	assert.Len(t, resultMessages, 1)
	assert.Equal(t, messages[0].Text, resultMessages[0].Text)
}

func TestGetMessagesForChatOlder_NotParticipant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	messageRepo := mocks.NewMockMessageRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	chatRepo := mocks.NewMockChatRepository(ctrl)
	validator := mocks.NewMockMessageValidator(ctrl)

	// Подготовка тестовых данных
	chatId := uuid.New()
	userId := uuid.New()

	// Ожидания для моков
	chatRepo.EXPECT().IsParticipant(context.Background(), chatId, userId).Return(false, nil)

	// Создаем сервис
	messageService := usecase.NewMessageService(messageRepo, fileRepo, chatRepo, validator)

	// Вызов метода
	messages, err := messageService.GetMessagesForChatOlder(context.Background(), chatId, userId, 5, time.Now())

	// Проверки
	assert.Error(t, err)
	assert.Nil(t, messages)
	assert.Equal(t, err, messenger_errors.ErrNotParticipant)
}

func TestDeleteMessage_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	messageRepo := mocks.NewMockMessageRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	chatRepo := mocks.NewMockChatRepository(ctrl)
	validator := mocks.NewMockMessageValidator(ctrl)

	// Подготовка тестовых данных
	messageId := uuid.New()

	// Ожидания для моков
	messageRepo.EXPECT().DeleteMessage(context.Background(), messageId).Return(nil)

	// Создаем сервис
	messageService := usecase.NewMessageService(messageRepo, fileRepo, chatRepo, validator)

	// Вызов метода
	err := messageService.DeleteMessage(context.Background(), messageId)

	// Проверки
	assert.NoError(t, err)
}

func TestDeleteMessage_InvalidId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	messageRepo := mocks.NewMockMessageRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	chatRepo := mocks.NewMockChatRepository(ctrl)
	validator := mocks.NewMockMessageValidator(ctrl)

	// Подготовка тестовых данных
	invalidMessageId := uuid.Nil

	// Создаем сервис
	messageService := usecase.NewMessageService(messageRepo, fileRepo, chatRepo, validator)

	// Вызов метода
	err := messageService.DeleteMessage(context.Background(), invalidMessageId)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "messageId is empty")
}
