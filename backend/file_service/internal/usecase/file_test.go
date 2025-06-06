package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	file_errors "quickflow/file_service/internal/errors"
	"quickflow/file_service/internal/usecase"
	"quickflow/file_service/internal/usecase/mocks"
	"quickflow/shared/models"
)

func TestFileUseCase_UploadFile_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокирование зависимостей
	mockFileStorage := mocks.NewMockFileStorage(ctrl)
	mockFileRepo := mocks.NewMockFileRepository(ctrl)
	mockFileValidator := mocks.NewMockFileValidator(ctrl)

	// Инициализация FileUseCase
	fileUseCase := usecase.NewFileUseCase(mockFileStorage, mockFileRepo, mockFileValidator)
	file := &models.File{Name: "test.txt", Size: 1024}
	fileURL := "http://example.com/test.txt"

	// Настройка моков
	mockFileValidator.EXPECT().ValidateFile(file).Return(nil)
	mockFileStorage.EXPECT().UploadFile(gomock.Any(), file).Return(fileURL, nil)
	mockFileRepo.EXPECT().AddFileRecord(gomock.Any(), file).Return(nil)

	// Вызов метода
	ctx := context.Background()
	url, err := fileUseCase.UploadFile(ctx, file)

	// Проверки
	assert.NoError(t, err)
	assert.Equal(t, fileURL, url)
}

func TestFileUseCase_UploadFile_FileIsNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокирование зависимостей
	mockFileStorage := mocks.NewMockFileStorage(ctrl)
	mockFileRepo := mocks.NewMockFileRepository(ctrl)
	mockFileValidator := mocks.NewMockFileValidator(ctrl)

	// Инициализация FileUseCase
	fileUseCase := usecase.NewFileUseCase(mockFileStorage, mockFileRepo, mockFileValidator)

	// Вызов метода с nil файлом
	ctx := context.Background()
	url, err := fileUseCase.UploadFile(ctx, nil)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, file_errors.ErrFileIsNil, err)
	assert.Equal(t, "", url)
}

func TestFileUseCase_UploadFile_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокирование зависимостей
	mockFileStorage := mocks.NewMockFileStorage(ctrl)
	mockFileRepo := mocks.NewMockFileRepository(ctrl)
	mockFileValidator := mocks.NewMockFileValidator(ctrl)

	// Инициализация FileUseCase
	fileUseCase := usecase.NewFileUseCase(mockFileStorage, mockFileRepo, mockFileValidator)

	// Тестовые данные
	file := &models.File{Name: "test.txt", Size: 1024}

	// Настройка моков
	mockFileValidator.EXPECT().ValidateFile(file).Return(errors.New("invalid file"))

	// Вызов метода
	ctx := context.Background()
	url, err := fileUseCase.UploadFile(ctx, file)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, "validation.ValidateFile: invalid file", err.Error())
	assert.Equal(t, "", url)
}

func TestFileUseCase_UploadManyMedia_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокирование зависимостей
	mockFileStorage := mocks.NewMockFileStorage(ctrl)
	mockFileRepo := mocks.NewMockFileRepository(ctrl)
	mockFileValidator := mocks.NewMockFileValidator(ctrl)

	// Инициализация FileUseCase
	fileUseCase := usecase.NewFileUseCase(mockFileStorage, mockFileRepo, mockFileValidator)

	// Тестовые данные
	files := []*models.File{
		{Name: "test1.jpg", Size: 1024},
		{Name: "test2.jpg", Size: 2048},
	}
	fileURLs := []string{"http://example.com/test1.jpg", "http://example.com/test2.jpg"}

	// Настройка моков
	mockFileValidator.EXPECT().ValidateFiles(files).Return(nil)
	mockFileStorage.EXPECT().UploadManyImages(gomock.Any(), files).Return(fileURLs, nil)
	mockFileRepo.EXPECT().AddFilesRecords(gomock.Any(), files).Return(nil)

	// Вызов метода
	ctx := context.Background()
	urls, err := fileUseCase.UploadManyMedia(ctx, files)

	// Проверки
	assert.NoError(t, err)
	assert.Equal(t, fileURLs, urls)
}

func TestFileUseCase_GetFileURL_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокирование зависимостей
	mockFileStorage := mocks.NewMockFileStorage(ctrl)
	mockFileValidator := mocks.NewMockFileValidator(ctrl)

	// Инициализация FileUseCase
	fileUseCase := usecase.NewFileUseCase(mockFileStorage, nil, mockFileValidator)

	// Тестовые данные
	fileName := "test.txt"
	fileURL := "http://example.com/test.txt"

	// Настройка моков
	mockFileValidator.EXPECT().ValidateFileName(fileName).Return(nil)
	mockFileStorage.EXPECT().GetFileURL(gomock.Any(), fileName).Return(fileURL, nil)

	// Вызов метода
	ctx := context.Background()
	url, err := fileUseCase.GetFileURL(ctx, fileName)

	// Проверки
	assert.NoError(t, err)
	assert.Equal(t, fileURL, url)
}

func TestFileUseCase_GetFileURL_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокирование зависимостей
	mockFileStorage := mocks.NewMockFileStorage(ctrl)
	mockFileValidator := mocks.NewMockFileValidator(ctrl)

	// Инициализация FileUseCase
	fileUseCase := usecase.NewFileUseCase(mockFileStorage, nil, mockFileValidator)

	// Тестовые данные
	fileName := "test.txt"

	// Настройка моков
	mockFileValidator.EXPECT().ValidateFileName(fileName).Return(errors.New("invalid filename"))

	// Вызов метода
	ctx := context.Background()
	url, err := fileUseCase.GetFileURL(ctx, fileName)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, "validation.ValidateFileName: invalid filename", err.Error())
	assert.Equal(t, "", url)
}

func TestFileUseCase_DeleteFile_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокирование зависимостей
	mockFileStorage := mocks.NewMockFileStorage(ctrl)
	mockFileValidator := mocks.NewMockFileValidator(ctrl)

	// Инициализация FileUseCase
	fileUseCase := usecase.NewFileUseCase(mockFileStorage, nil, mockFileValidator)

	// Тестовые данные
	fileName := "test.txt"

	// Настройка моков
	mockFileValidator.EXPECT().ValidateFileName(fileName).Return(nil)
	mockFileStorage.EXPECT().DeleteFile(gomock.Any(), fileName).Return(nil)

	// Вызов метода
	ctx := context.Background()
	err := fileUseCase.DeleteFile(ctx, fileName)

	// Проверки
	assert.NoError(t, err)
}

func TestFileUseCase_DeleteFile_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокирование зависимостей
	mockFileStorage := mocks.NewMockFileStorage(ctrl)
	mockFileValidator := mocks.NewMockFileValidator(ctrl)

	// Инициализация FileUseCase
	fileUseCase := usecase.NewFileUseCase(mockFileStorage, nil, mockFileValidator)

	// Тестовые данные
	fileName := "test.txt"

	// Настройка моков
	mockFileValidator.EXPECT().ValidateFileName(fileName).Return(errors.New("invalid filename"))

	// Вызов метода
	ctx := context.Background()
	err := fileUseCase.DeleteFile(ctx, fileName)

	// Проверки
	assert.Error(t, err)
	assert.Equal(t, "validation.ValidateFileName: invalid filename", err.Error())
}
