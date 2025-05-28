package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	messenger_errors "quickflow/messenger_service/internal/errors"
	"quickflow/messenger_service/internal/usecase"
	"quickflow/messenger_service/internal/usecase/mocks"
	"quickflow/shared/models"
)

func TestAddStickerPack_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	stickerRepo := mocks.NewMockStickerRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockStickerPackValidator(ctrl)

	// Подготовка тестовых данных
	stickerPack := &models.StickerPack{
		Id:       uuid.New(),
		Name:     "Funny Stickers",
		Stickers: []*models.File{{URL: "sticker_url_1", DisplayType: models.DisplayTypeSticker, Name: "sticker1"}},
	}

	// Ожидания для моков
	validator.EXPECT().ValidateStickerPack(stickerPack).Return(nil)
	stickerRepo.EXPECT().AddStickerPack(context.Background(), gomock.Any()).Return(nil)

	// Создаем сервис
	stickerService := usecase.NewStickerService(stickerRepo, fileRepo, validator)

	// Вызов метода
	result, err := stickerService.AddStickerPack(context.Background(), stickerPack)

	// Проверки
	assert.NoError(t, err)
	assert.Equal(t, stickerPack, result)
}

func TestAddStickerPack_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	stickerRepo := mocks.NewMockStickerRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockStickerPackValidator(ctrl)

	// Подготовка тестовых данных
	stickerPack := &models.StickerPack{
		Id:       uuid.New(),
		Name:     "Funny Stickers",
		Stickers: []*models.File{{URL: "sticker_url_1", DisplayType: models.DisplayTypeSticker, Name: "sticker1"}},
	}

	// Ожидания для моков
	validator.EXPECT().ValidateStickerPack(stickerPack).Return(errors.New("validation error"))

	// Создаем сервис
	stickerService := usecase.NewStickerService(stickerRepo, fileRepo, validator)

	// Вызов метода
	result, err := stickerService.AddStickerPack(context.Background(), stickerPack)

	// Проверки
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetStickerPack_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	stickerRepo := mocks.NewMockStickerRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockStickerPackValidator(ctrl)

	// Подготовка тестовых данных
	packId := uuid.New()
	stickerPack := models.StickerPack{
		Id:       packId,
		Name:     "Funny Stickers",
		Stickers: []*models.File{{URL: "sticker_url_1", DisplayType: models.DisplayTypeSticker, Name: "sticker1"}},
	}

	// Ожидания для моков
	stickerRepo.EXPECT().GetStickerPack(context.Background(), packId).Return(stickerPack, nil)

	// Создаем сервис
	stickerService := usecase.NewStickerService(stickerRepo, fileRepo, validator)

	// Вызов метода
	result, err := stickerService.GetStickerPack(context.Background(), packId)

	// Проверки
	assert.NoError(t, err)
	assert.Equal(t, stickerPack, result)
}

func TestGetStickerPack_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	stickerRepo := mocks.NewMockStickerRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockStickerPackValidator(ctrl)

	// Подготовка тестовых данных
	packId := uuid.New()

	// Ожидания для моков
	stickerRepo.EXPECT().GetStickerPack(context.Background(), packId).Return(models.StickerPack{}, messenger_errors.ErrNotFound)

	// Создаем сервис
	stickerService := usecase.NewStickerService(stickerRepo, fileRepo, validator)

	// Вызов метода
	result, err := stickerService.GetStickerPack(context.Background(), packId)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, err, messenger_errors.ErrNotFound)
	assert.Equal(t, models.StickerPack{}, result)
}

func TestDeleteStickerPack_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	stickerRepo := mocks.NewMockStickerRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockStickerPackValidator(ctrl)

	// Подготовка тестовых данных
	userId := uuid.New()
	packId := uuid.New()

	// Ожидания для моков
	stickerRepo.EXPECT().BelongsTo(context.Background(), userId, packId).Return(true, nil)
	stickerRepo.EXPECT().DeleteStickerPack(context.Background(), userId, packId).Return(nil)

	// Создаем сервис
	stickerService := usecase.NewStickerService(stickerRepo, fileRepo, validator)

	// Вызов метода
	err := stickerService.DeleteStickerPack(context.Background(), userId, packId)

	// Проверки
	assert.NoError(t, err)
}

func TestDeleteStickerPack_NotOwner(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	stickerRepo := mocks.NewMockStickerRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockStickerPackValidator(ctrl)

	// Подготовка тестовых данных
	userId := uuid.New()
	packId := uuid.New()

	// Ожидания для моков
	stickerRepo.EXPECT().BelongsTo(context.Background(), userId, packId).Return(false, nil)

	// Создаем сервис
	stickerService := usecase.NewStickerService(stickerRepo, fileRepo, validator)

	// Вызов метода
	err := stickerService.DeleteStickerPack(context.Background(), userId, packId)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, err, messenger_errors.ErrNotOwnerOfStickerPack)
}

func TestGetStickerPackByName_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	stickerRepo := mocks.NewMockStickerRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockStickerPackValidator(ctrl)

	// Подготовка тестовых данных
	name := "Funny Stickers"
	stickerPack := models.StickerPack{
		Id:       uuid.New(),
		Name:     name,
		Stickers: []*models.File{{URL: "sticker_url_1", DisplayType: models.DisplayTypeSticker, Name: "sticker1"}},
	}

	// Ожидания для моков
	stickerRepo.EXPECT().GetStickerPackByName(context.Background(), name).Return(stickerPack, nil)

	// Создаем сервис
	stickerService := usecase.NewStickerService(stickerRepo, fileRepo, validator)

	// Вызов метода
	result, err := stickerService.GetStickerPackByName(context.Background(), name)

	// Проверки
	assert.NoError(t, err)
	assert.Equal(t, stickerPack, result)
}

func TestBelongsTo_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	stickerRepo := mocks.NewMockStickerRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockStickerPackValidator(ctrl)

	// Подготовка тестовых данных
	userId := uuid.New()
	packId := uuid.New()

	// Ожидания для моков
	stickerRepo.EXPECT().BelongsTo(context.Background(), userId, packId).Return(true, nil)

	// Создаем сервис
	stickerService := usecase.NewStickerService(stickerRepo, fileRepo, validator)

	// Вызов метода
	belongs, err := stickerService.BelongsTo(context.Background(), userId, packId)

	// Проверки
	assert.NoError(t, err)
	assert.True(t, belongs)
}

func TestBelongsTo_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Моки
	stickerRepo := mocks.NewMockStickerRepository(ctrl)
	fileRepo := mocks.NewMockFileService(ctrl)
	validator := mocks.NewMockStickerPackValidator(ctrl)

	// Подготовка тестовых данных
	userId := uuid.New()
	packId := uuid.New()

	// Ожидания для моков
	stickerRepo.EXPECT().BelongsTo(context.Background(), userId, packId).Return(false, messenger_errors.ErrNotFound)

	// Создаем сервис
	stickerService := usecase.NewStickerService(stickerRepo, fileRepo, validator)

	// Вызов метода
	belongs, err := stickerService.BelongsTo(context.Background(), userId, packId)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, err, messenger_errors.ErrNotFound)
	assert.False(t, belongs)
}
