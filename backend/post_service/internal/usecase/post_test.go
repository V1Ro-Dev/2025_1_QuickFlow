package usecase_test

import (
    "context"
    "testing"
    "time"

    "github.com/golang/mock/gomock"
    "github.com/google/uuid"
    "github.com/stretchr/testify/assert"

    "quickflow/post_service/internal/errors"
    "quickflow/post_service/internal/usecase"
    "quickflow/post_service/internal/usecase/mocks"
    "quickflow/shared/models"
)

func TestAddPost(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    // Моки
    postRepo := mocks.NewMockPostRepository(ctrl)
    fileRepo := mocks.NewMockFileService(ctrl)
    validator := mocks.NewMockPostValidator(ctrl)

    // Создание тестовых данных
    post := models.Post{
        Id:          uuid.New(),
        CreatorId:   uuid.New(),
        CreatorType: models.PostUser,
        Desc:        "This is a test post",
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
        LikeCount:   0,
    }

    // Ожидаемый вызов репозитория
    postRepo.EXPECT().AddPost(gomock.Any(), gomock.Any()).Return(nil)
    postRepo.EXPECT().GetPost(gomock.Any(), gomock.Any()).Return(post, nil)

    // Создание объекта usecase
    service := usecase.NewPostUseCase(postRepo, fileRepo, validator)

    // Вызов функции
    result, err := service.AddPost(context.Background(), post)

    // Проверки
    assert.NoError(t, err)
    assert.Equal(t, post, *result)
}

func TestDeletePost(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    // Моки
    postRepo := mocks.NewMockPostRepository(ctrl)
    fileRepo := mocks.NewMockFileService(ctrl)
    validator := mocks.NewMockPostValidator(ctrl)

    postId := uuid.New()
    userId := uuid.New()

    // Ожидаемый вызов репозитория
    postRepo.EXPECT().GetPostFiles(gomock.Any(), postId).Return([]string{"file1", "file2"}, nil)
    postRepo.EXPECT().DeletePost(gomock.Any(), postId).Return(nil)
    fileRepo.EXPECT().DeleteFile(gomock.Any(), "file1").Return(nil)
    fileRepo.EXPECT().DeleteFile(gomock.Any(), "file2").Return(nil)

    // Создание объекта usecase
    service := usecase.NewPostUseCase(postRepo, fileRepo, validator)

    // Вызов функции
    err := service.DeletePost(context.Background(), userId, postId)

    // Проверки
    assert.NoError(t, err)
}

func TestFetchFeed_InvalidNumPosts(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    // Моки
    postRepo := mocks.NewMockPostRepository(ctrl)
    fileRepo := mocks.NewMockFileService(ctrl)
    validator := mocks.NewMockPostValidator(ctrl)

    // Создание объекта usecase
    service := usecase.NewPostUseCase(postRepo, fileRepo, validator)

    // Ожидаемая ошибка для некорректного числа постов
    validator.EXPECT().ValidateFeedParams(gomock.Any(), gomock.Any()).Return(errors.ErrInvalidNumPosts)

    // Вызов функции
    _, err := service.FetchFeed(context.Background(), uuid.New(), 0, time.Now())

    // Проверки
    assert.Error(t, err)
}

func TestLikePost(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    // Моки
    postRepo := mocks.NewMockPostRepository(ctrl)
    fileRepo := mocks.NewMockFileService(ctrl)
    validator := mocks.NewMockPostValidator(ctrl)

    postId := uuid.New()
    userId := uuid.New()

    // Ожидаемый вызов репозитория
    postRepo.EXPECT().CheckIfPostLiked(gomock.Any(), postId, userId).Return(false, nil)
    postRepo.EXPECT().LikePost(gomock.Any(), postId, userId).Return(nil)

    // Создание объекта usecase
    service := usecase.NewPostUseCase(postRepo, fileRepo, validator)

    // Вызов функции
    err := service.LikePost(context.Background(), postId, userId)

    // Проверки
    assert.NoError(t, err)
}

func TestUnlikePost(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    // Моки
    postRepo := mocks.NewMockPostRepository(ctrl)
    fileRepo := mocks.NewMockFileService(ctrl)
    validator := mocks.NewMockPostValidator(ctrl)

    postId := uuid.New()
    userId := uuid.New()

    // Ожидаемый вызов репозитория
    postRepo.EXPECT().CheckIfPostLiked(gomock.Any(), postId, userId).Return(true, nil)
    postRepo.EXPECT().UnlikePost(gomock.Any(), postId, userId).Return(nil)

    // Создание объекта usecase
    service := usecase.NewPostUseCase(postRepo, fileRepo, validator)

    // Вызов функции
    err := service.UnlikePost(context.Background(), postId, userId)

    // Проверки
    assert.NoError(t, err)
}

func TestUpdatePost(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    // Моки
    postRepo := mocks.NewMockPostRepository(ctrl)
    fileRepo := mocks.NewMockFileService(ctrl)
    validator := mocks.NewMockPostValidator(ctrl)

    postId := uuid.New()
    userId := uuid.New()

    // Создание данных для обновления
    postUpdate := models.PostUpdate{
        Id:   postId,
        Desc: "Updated description",
    }

    // Ожидаемый вызов репозитория
    postRepo.EXPECT().GetPost(gomock.Any(), postId).Return(models.Post{
        Id:          postId,
        CreatorId:   userId,
        CreatorType: models.PostUser,
        Desc:        "Old description",
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
        LikeCount:   0,
    }, nil).AnyTimes()
    postRepo.EXPECT().UpdatePost(gomock.Any(), postUpdate).Return(nil).AnyTimes()
    postRepo.EXPECT().GetPost(gomock.Any(), postId).Return(models.Post{
        Id:          postId,
        CreatorId:   userId,
        CreatorType: models.PostUser,
        Desc:        postUpdate.Desc,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
        LikeCount:   0,
    }, nil).AnyTimes()

    postRepo.EXPECT().GetPostFiles(gomock.Any(), postId).Return([]string{"file1", "file2"}, nil).AnyTimes()
    fileRepo.EXPECT().DeleteFile(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

    // Создание объекта usecase
    service := usecase.NewPostUseCase(postRepo, fileRepo, validator)

    // Вызов функции
    _, err := service.UpdatePost(context.Background(), postUpdate, userId)

    // Проверки
    assert.NoError(t, err)
}
