package usecase

import (
	"context"
	"errors"
	"quickflow/messenger_service/internal/usecase/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	messenger_errors "quickflow/messenger_service/internal/errors"
	"quickflow/shared/models"
)

func TestNewChatUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mocks.NewMockChatRepository(ctrl)
	mockFileRepo := mocks.NewMockFileService(ctrl)
	mockProfileRepo := mocks.NewMockProfileService(ctrl)
	mockMessageRepo := mocks.NewMockMessageRepository(ctrl)
	mockValidator := mocks.NewMockChatValidator(ctrl)

	service := NewChatUseCase(
		mockChatRepo,
		mockFileRepo,
		mockProfileRepo,
		mockMessageRepo,
		mockValidator,
	)

	assert.NotNil(t, service)
	assert.Equal(t, mockChatRepo, service.chatRepo)
	assert.Equal(t, mockFileRepo, service.fileRepo)
	assert.Equal(t, mockProfileRepo, service.profileRepo)
	assert.Equal(t, mockMessageRepo, service.messageRepo)
	assert.Equal(t, mockValidator, service.validator)
}

func TestCreateChat_Private(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mocks.NewMockChatRepository(ctrl)
	mockValidator := mocks.NewMockChatValidator(ctrl)
	service := NewChatUseCase(
		mockChatRepo,
		nil, nil, nil,
		mockValidator,
	)

	ctx := context.Background()
	chatInfo := models.ChatCreationInfo{
		Type: models.ChatTypePrivate,
	}

	mockValidator.EXPECT().
		ValidateChatCreationInfo(chatInfo).
		Return(nil)

	mockChatRepo.EXPECT().
		CreateChat(ctx, gomock.Any()).
		Do(func(_ context.Context, chat models.Chat) {
			assert.Equal(t, models.ChatTypePrivate, chat.Type)
			assert.NotEqual(t, uuid.Nil, chat.ID)
		}).
		Return(nil)

	result, err := service.CreateChat(ctx, chatInfo)

	assert.NoError(t, err)
	assert.Equal(t, models.ChatTypePrivate, result.Type)
	assert.NotEqual(t, uuid.Nil, result.ID)
}

func TestCreateChat_Group(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mocks.NewMockChatRepository(ctrl)
	mockFileRepo := mocks.NewMockFileService(ctrl)
	mockValidator := mocks.NewMockChatValidator(ctrl)
	service := NewChatUseCase(
		mockChatRepo,
		mockFileRepo, nil, nil,
		mockValidator,
	)

	ctx := context.Background()
	chatInfo := models.ChatCreationInfo{
		Type:   models.ChatTypeGroup,
		Name:   "Test Group",
		Avatar: &models.File{},
	}

	mockValidator.EXPECT().
		ValidateChatCreationInfo(chatInfo).
		Return(nil)

	mockFileRepo.EXPECT().
		UploadFile(ctx, chatInfo.Avatar).
		Return("avatar_url", nil)

	mockChatRepo.EXPECT().
		CreateChat(ctx, gomock.Any()).
		Do(func(_ context.Context, chat models.Chat) {
			assert.Equal(t, models.ChatTypeGroup, chat.Type)
			assert.Equal(t, "Test Group", chat.Name)
			assert.Equal(t, "avatar_url", chat.AvatarURL)
		}).
		Return(nil)

	result, err := service.CreateChat(ctx, chatInfo)

	assert.NoError(t, err)
	assert.Equal(t, models.ChatTypeGroup, result.Type)
	assert.Equal(t, "Test Group", result.Name)
	assert.Equal(t, "avatar_url", result.AvatarURL)
}

func TestCreateChat_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockValidator := mocks.NewMockChatValidator(ctrl)
	service := NewChatUseCase(
		nil, nil, nil, nil,
		mockValidator,
	)

	ctx := context.Background()
	chatInfo := models.ChatCreationInfo{}

	mockValidator.EXPECT().
		ValidateChatCreationInfo(chatInfo).
		Return(errors.New("validation error"))

	_, err := service.CreateChat(ctx, chatInfo)

	assert.Error(t, err)
	assert.Equal(t, messenger_errors.ErrInvalidChatCreationInfo, err)
}

func TestCreateChat_UploadError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileRepo := mocks.NewMockFileService(ctrl)
	mockValidator := mocks.NewMockChatValidator(ctrl)
	service := NewChatUseCase(
		nil,
		mockFileRepo, nil, nil,
		mockValidator,
	)

	ctx := context.Background()
	chatInfo := models.ChatCreationInfo{
		Type:   models.ChatTypeGroup,
		Avatar: &models.File{},
	}

	mockValidator.EXPECT().
		ValidateChatCreationInfo(chatInfo).
		Return(nil)

	mockFileRepo.EXPECT().
		UploadFile(ctx, chatInfo.Avatar).
		Return("", errors.New("upload error"))

	_, err := service.CreateChat(ctx, chatInfo)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "upload error")
}

func TestGetUserChats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mocks.NewMockChatRepository(ctrl)
	mockProfileRepo := mocks.NewMockProfileService(ctrl)
	mockMessageRepo := mocks.NewMockMessageRepository(ctrl)
	service := NewChatUseCase(
		mockChatRepo,
		nil, mockProfileRepo, mockMessageRepo,
		nil,
	)

	ctx := context.Background()
	userID := uuid.New()
	chatID := uuid.New()
	otherUserID := uuid.New()

	// Setup test data
	chats := []models.Chat{
		{
			ID:   chatID,
			Type: models.ChatTypePrivate,
		},
	}

	// Mock expectations
	mockChatRepo.EXPECT().
		GetUserChats(ctx, userID).
		Return(chats, nil)

	mockChatRepo.EXPECT().
		GetChatParticipants(gomock.Any(), chatID).
		Return([]uuid.UUID{userID, otherUserID}, nil).AnyTimes()

	mockProfileRepo.EXPECT().
		GetPublicUsersInfo(gomock.Any(), []uuid.UUID{userID, otherUserID}).
		Return([]models.PublicUserInfo{
			{
				Id:        userID,
				Firstname: "Me",
				Lastname:  "User",
			},
			{
				Id:        otherUserID,
				Firstname: "Other",
				Lastname:  "User",
				AvatarURL: "avatar_url",
			},
		}, nil)

	mockMessageRepo.EXPECT().
		GetLastChatMessage(gomock.Any(), chatID).
		Return(&models.Message{Text: "last message"}, nil)

	// Execute
	result, err := service.GetUserChats(ctx, userID)

	// Verify
	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestGetChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mocks.NewMockChatRepository(ctrl)
	service := NewChatUseCase(
		mockChatRepo,
		nil, nil, nil,
		nil,
	)

	ctx := context.Background()
	chatID := uuid.New()
	expectedChat := models.Chat{ID: chatID}

	mockChatRepo.EXPECT().
		GetChat(ctx, chatID).
		Return(expectedChat, nil)

	result, err := service.GetChat(ctx, chatID)

	assert.NoError(t, err)
	assert.Equal(t, chatID, result.ID)
}

func TestDeleteChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mocks.NewMockChatRepository(ctrl)
	service := NewChatUseCase(
		mockChatRepo,
		nil, nil, nil,
		nil,
	)

	ctx := context.Background()
	chatID := uuid.New()

	// Success case
	mockChatRepo.EXPECT().
		Exists(ctx, chatID).
		Return(true, nil)

	mockChatRepo.EXPECT().
		DeleteChat(ctx, chatID).
		Return(nil)

	err := service.DeleteChat(ctx, chatID)
	assert.NoError(t, err)

	// Chat doesn't exist
	mockChatRepo.EXPECT().
		Exists(ctx, chatID).
		Return(false, nil)

	err = service.DeleteChat(ctx, chatID)
	assert.Equal(t, messenger_errors.ErrNotFound, err)
}

func TestJoinChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mocks.NewMockChatRepository(ctrl)
	service := NewChatUseCase(
		mockChatRepo,
		nil, nil, nil,
		nil,
	)

	ctx := context.Background()
	chatID := uuid.New()
	userID := uuid.New()

	// Success case
	mockChatRepo.EXPECT().
		Exists(ctx, chatID).
		Return(true, nil)

	mockChatRepo.EXPECT().
		IsParticipant(ctx, chatID, userID).
		Return(false, nil)

	mockChatRepo.EXPECT().
		JoinChat(ctx, chatID, userID).
		Return(nil)

	err := service.JoinChat(ctx, chatID, userID)
	assert.NoError(t, err)

	// Already in chat
	mockChatRepo.EXPECT().
		Exists(ctx, chatID).
		Return(true, nil)

	mockChatRepo.EXPECT().
		IsParticipant(ctx, chatID, userID).
		Return(true, nil)

	err = service.JoinChat(ctx, chatID, userID)
	assert.Equal(t, messenger_errors.ErrAlreadyInChat, err)

	// Chat doesn't exist
	mockChatRepo.EXPECT().
		Exists(ctx, chatID).
		Return(false, nil)

	err = service.JoinChat(ctx, chatID, userID)
	assert.Equal(t, messenger_errors.ErrNotFound, err)
}

func TestLeaveChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mocks.NewMockChatRepository(ctrl)
	service := NewChatUseCase(
		mockChatRepo,
		nil, nil, nil,
		nil,
	)

	ctx := context.Background()
	chatID := uuid.New()
	userID := uuid.New()

	// Success case
	mockChatRepo.EXPECT().
		Exists(ctx, chatID).
		Return(true, nil)

	mockChatRepo.EXPECT().
		IsParticipant(ctx, chatID, userID).
		Return(true, nil)

	mockChatRepo.EXPECT().
		LeaveChat(ctx, chatID, userID).
		Return(nil)

	err := service.LeaveChat(ctx, chatID, userID)
	assert.NoError(t, err)

	// Not a participant
	mockChatRepo.EXPECT().
		Exists(ctx, chatID).
		Return(true, nil)

	mockChatRepo.EXPECT().
		IsParticipant(ctx, chatID, userID).
		Return(false, nil)

	err = service.LeaveChat(ctx, chatID, userID)
	assert.Equal(t, messenger_errors.ErrNotFound, err)

	// Chat doesn't exist
	mockChatRepo.EXPECT().
		Exists(ctx, chatID).
		Return(false, nil)

	err = service.LeaveChat(ctx, chatID, userID)
	assert.Equal(t, messenger_errors.ErrNotFound, err)
}

func TestGetChatParticipants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mocks.NewMockChatRepository(ctrl)
	service := NewChatUseCase(
		mockChatRepo,
		nil, nil, nil,
		nil,
	)

	ctx := context.Background()
	chatID := uuid.New()
	expectedParticipants := []uuid.UUID{uuid.New(), uuid.New()}

	mockChatRepo.EXPECT().
		GetChatParticipants(ctx, chatID).
		Return(expectedParticipants, nil)

	result, err := service.GetChatParticipants(ctx, chatID)

	assert.NoError(t, err)
	assert.Equal(t, expectedParticipants, result)
}

func TestGetPrivateChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mocks.NewMockChatRepository(ctrl)
	service := NewChatUseCase(
		mockChatRepo,
		nil, nil, nil,
		nil,
	)

	ctx := context.Background()
	user1 := uuid.New()
	user2 := uuid.New()
	expectedChat := models.Chat{ID: uuid.New()}

	mockChatRepo.EXPECT().
		GetPrivateChat(ctx, user1, user2).
		Return(expectedChat, nil)

	result, err := service.GetPrivateChat(ctx, user1, user2)

	assert.NoError(t, err)
	assert.Equal(t, expectedChat.ID, result.ID)
}

func TestGetNumUnreadChats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mocks.NewMockChatRepository(ctrl)
	service := NewChatUseCase(
		mockChatRepo,
		nil, nil, nil,
		nil,
	)

	ctx := context.Background()
	userID := uuid.New()
	expectedCount := 5

	mockChatRepo.EXPECT().
		GetNumUnreadChats(ctx, userID).
		Return(expectedCount, nil)

	result, err := service.GetNumUnreadChats(ctx, userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCount, result)
}

func TestGetUserChats_ErrorHandling(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mocks.NewMockChatRepository(ctrl)
	service := NewChatUseCase(
		mockChatRepo,
		nil, nil, nil,
		nil,
	)

	ctx := context.Background()
	userID := uuid.New()

	// Test GetUserChats error
	mockChatRepo.EXPECT().
		GetUserChats(ctx, userID).
		Return(nil, errors.New("database error"))

	_, err := service.GetUserChats(ctx, userID)
	assert.Error(t, err)

	// Test GetChatParticipants error
	chats := []models.Chat{{ID: uuid.New(), Type: models.ChatTypePrivate}}
	mockChatRepo.EXPECT().
		GetUserChats(ctx, userID).
		Return(chats, nil)

	mockChatRepo.EXPECT().
		GetChatParticipants(gomock.Any(), chats[0].ID).
		Return(nil, errors.New("participants error")).AnyTimes()

	_, err = service.GetUserChats(ctx, userID)
	assert.Error(t, err)
}

func TestGetUserChats_ConcurrentProcessing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChatRepo := mocks.NewMockChatRepository(ctrl)
	mockProfileRepo := mocks.NewMockProfileService(ctrl)
	mockMessageRepo := mocks.NewMockMessageRepository(ctrl)
	service := NewChatUseCase(
		mockChatRepo,
		nil, mockProfileRepo, mockMessageRepo,
		nil,
	)

	ctx := context.Background()
	userID := uuid.New()

	// Create multiple private chats
	chats := make([]models.Chat, 10)
	for i := 0; i < 10; i++ {
		chats[i] = models.Chat{
			ID:   uuid.New(),
			Type: models.ChatTypePrivate,
		}
	}

	// Mock expectations
	mockChatRepo.EXPECT().
		GetUserChats(ctx, userID).
		Return(chats, nil)

	// Each chat will have its own participants
	for _, chat := range chats {
		otherUser := uuid.New()
		mockChatRepo.EXPECT().
			GetChatParticipants(gomock.Any(), chat.ID).
			Return([]uuid.UUID{userID, otherUser}, nil).AnyTimes()

		mockProfileRepo.EXPECT().
			GetPublicUsersInfo(gomock.Any(), []uuid.UUID{userID, otherUser}).
			Return([]models.PublicUserInfo{
				{Id: userID},
				{Id: otherUser, Firstname: "User", Lastname: "Other"},
			}, nil)

		mockMessageRepo.EXPECT().
			GetLastChatMessage(gomock.Any(), chat.ID).
			Return(&models.Message{}, nil)
	}

	// Execute and verify concurrent processing
	result, err := service.GetUserChats(ctx, userID)
	assert.NoError(t, err)
	assert.Len(t, result, 10)
}
