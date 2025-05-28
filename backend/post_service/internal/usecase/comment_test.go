package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"quickflow/gateway/utils/validation"
	"quickflow/post_service/internal/errors"
	"quickflow/post_service/internal/usecase"
	"quickflow/post_service/internal/usecase/mocks"
	"quickflow/shared/models"
)

func TestDeleteComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	commentRepo := mocks.NewMockCommentRepository(ctrl)
	fileService := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockPostValidator(ctrl)

	commentId := uuid.New()
	userId := uuid.New()

	// Ожидаемый вызов репозитория
	commentRepo.EXPECT().GetCommentFiles(gomock.Any(), commentId).Return([]string{"file1", "file2"}, nil)
	commentRepo.EXPECT().DeleteComment(gomock.Any(), commentId).Return(nil)
	fileService.EXPECT().DeleteFile(gomock.Any(), "file1").Return(nil)
	fileService.EXPECT().DeleteFile(gomock.Any(), "file2").Return(nil)

	// Создание объекта usecase
	service := usecase.NewCommentUseCase(commentRepo, fileService, validator)

	// Вызов функции
	err := service.DeleteComment(context.Background(), userId, commentId)

	// Проверки
	assert.NoError(t, err)
}

func TestFetchCommentsForPost_InvalidNumComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	commentRepo := mocks.NewMockCommentRepository(ctrl)
	fileService := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockPostValidator(ctrl)

	// Создание объекта usecase
	service := usecase.NewCommentUseCase(commentRepo, fileService, validator)

	// Ожидаемая ошибка для некорректного числа комментариев
	validator.EXPECT().ValidateFeedParams(gomock.Any(), gomock.Any()).Return(validation.ErrInvalidNumPosts)

	// Вызов функции
	_, err := service.FetchCommentsForPost(context.Background(), uuid.New(), 0, time.Now())

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, errors.ErrInvalidNumComments, err)
}

func TestLikeComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	commentRepo := mocks.NewMockCommentRepository(ctrl)
	fileService := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockPostValidator(ctrl)

	postId := uuid.New()
	userId := uuid.New()

	// Ожидаемый вызов репозитория
	commentRepo.EXPECT().CheckIfCommentLiked(gomock.Any(), postId, userId).Return(false, nil)
	commentRepo.EXPECT().LikeComment(gomock.Any(), postId, userId).Return(nil)

	// Создание объекта usecase
	service := usecase.NewCommentUseCase(commentRepo, fileService, validator)

	// Вызов функции
	err := service.LikeComment(context.Background(), postId, userId)

	// Проверки
	assert.NoError(t, err)
}

func TestUnlikeComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	commentRepo := mocks.NewMockCommentRepository(ctrl)
	fileService := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockPostValidator(ctrl)

	postId := uuid.New()
	userId := uuid.New()

	// Ожидаемый вызов репозитория
	commentRepo.EXPECT().CheckIfCommentLiked(gomock.Any(), postId, userId).Return(true, nil)
	commentRepo.EXPECT().UnlikeComment(gomock.Any(), postId, userId).Return(nil)

	// Создание объекта usecase
	service := usecase.NewCommentUseCase(commentRepo, fileService, validator)

	// Вызов функции
	err := service.UnlikeComment(context.Background(), postId, userId)

	// Проверки
	assert.NoError(t, err)
}

func TestGetComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	commentRepo := mocks.NewMockCommentRepository(ctrl)
	fileService := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockPostValidator(ctrl)

	commentId := uuid.New()

	// Ожидаемый вызов репозитория
	comment := models.Comment{
		Id:        commentId,
		UserId:    uuid.New(),
		PostId:    uuid.New(),
		Text:      "This is a comment",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		LikeCount: 0,
		IsLiked:   false,
	}
	commentRepo.EXPECT().GetComment(gomock.Any(), commentId).Return(comment, nil)

	// Создание объекта usecase
	service := usecase.NewCommentUseCase(commentRepo, fileService, validator)

	// Вызов функции
	result, err := service.GetComment(context.Background(), commentId, uuid.New())

	// Проверки
	assert.NoError(t, err)
	assert.Equal(t, comment, *result)
}

func TestGetLastPostComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	commentRepo := mocks.NewMockCommentRepository(ctrl)
	fileService := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockPostValidator(ctrl)

	postId := uuid.New()

	// Ожидаемый вызов репозитория
	comment := models.Comment{
		Id:        uuid.New(),
		UserId:    uuid.New(),
		PostId:    postId,
		Text:      "This is the last comment",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		LikeCount: 10,
		IsLiked:   true,
	}
	commentRepo.EXPECT().GetLastPostComment(gomock.Any(), postId).Return(&comment, nil)

	// Создание объекта usecase
	service := usecase.NewCommentUseCase(commentRepo, fileService, validator)

	// Вызов функции
	result, err := service.GetLastPostComment(context.Background(), postId)

	// Проверки
	assert.NoError(t, err)
	assert.Equal(t, comment, *result)
}
