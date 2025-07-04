// Code generated by MockGen. DO NOT EDIT.
// Source: .//file_service/internal/usecase/file-usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models "quickflow/shared/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileStorage is a mock of FileStorage interface.
type MockFileStorage struct {
	ctrl     *gomock.Controller
	recorder *MockFileStorageMockRecorder
}

// MockFileStorageMockRecorder is the mock recorder for MockFileStorage.
type MockFileStorageMockRecorder struct {
	mock *MockFileStorage
}

// NewMockFileStorage creates a new mock instance.
func NewMockFileStorage(ctrl *gomock.Controller) *MockFileStorage {
	mock := &MockFileStorage{ctrl: ctrl}
	mock.recorder = &MockFileStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileStorage) EXPECT() *MockFileStorageMockRecorder {
	return m.recorder
}

// DeleteFile mocks base method.
func (m *MockFileStorage) DeleteFile(ctx context.Context, filename string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", ctx, filename)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockFileStorageMockRecorder) DeleteFile(ctx, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockFileStorage)(nil).DeleteFile), ctx, filename)
}

// GetFileURL mocks base method.
func (m *MockFileStorage) GetFileURL(ctx context.Context, filename string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileURL", ctx, filename)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileURL indicates an expected call of GetFileURL.
func (mr *MockFileStorageMockRecorder) GetFileURL(ctx, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileURL", reflect.TypeOf((*MockFileStorage)(nil).GetFileURL), ctx, filename)
}

// UploadFile mocks base method.
func (m *MockFileStorage) UploadFile(ctx context.Context, file *models.File) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", ctx, file)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockFileStorageMockRecorder) UploadFile(ctx, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockFileStorage)(nil).UploadFile), ctx, file)
}

// UploadManyImages mocks base method.
func (m *MockFileStorage) UploadManyImages(ctx context.Context, files []*models.File) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadManyImages", ctx, files)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadManyImages indicates an expected call of UploadManyImages.
func (mr *MockFileStorageMockRecorder) UploadManyImages(ctx, files interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadManyImages", reflect.TypeOf((*MockFileStorage)(nil).UploadManyImages), ctx, files)
}

// MockFileRepository is a mock of FileRepository interface.
type MockFileRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFileRepositoryMockRecorder
}

// MockFileRepositoryMockRecorder is the mock recorder for MockFileRepository.
type MockFileRepositoryMockRecorder struct {
	mock *MockFileRepository
}

// NewMockFileRepository creates a new mock instance.
func NewMockFileRepository(ctrl *gomock.Controller) *MockFileRepository {
	mock := &MockFileRepository{ctrl: ctrl}
	mock.recorder = &MockFileRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileRepository) EXPECT() *MockFileRepositoryMockRecorder {
	return m.recorder
}

// AddFileRecord mocks base method.
func (m *MockFileRepository) AddFileRecord(ctx context.Context, file *models.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFileRecord", ctx, file)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFileRecord indicates an expected call of AddFileRecord.
func (mr *MockFileRepositoryMockRecorder) AddFileRecord(ctx, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFileRecord", reflect.TypeOf((*MockFileRepository)(nil).AddFileRecord), ctx, file)
}

// AddFilesRecords mocks base method.
func (m *MockFileRepository) AddFilesRecords(ctx context.Context, files []*models.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFilesRecords", ctx, files)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFilesRecords indicates an expected call of AddFilesRecords.
func (mr *MockFileRepositoryMockRecorder) AddFilesRecords(ctx, files interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFilesRecords", reflect.TypeOf((*MockFileRepository)(nil).AddFilesRecords), ctx, files)
}

// MockFileValidator is a mock of FileValidator interface.
type MockFileValidator struct {
	ctrl     *gomock.Controller
	recorder *MockFileValidatorMockRecorder
}

// MockFileValidatorMockRecorder is the mock recorder for MockFileValidator.
type MockFileValidatorMockRecorder struct {
	mock *MockFileValidator
}

// NewMockFileValidator creates a new mock instance.
func NewMockFileValidator(ctrl *gomock.Controller) *MockFileValidator {
	mock := &MockFileValidator{ctrl: ctrl}
	mock.recorder = &MockFileValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileValidator) EXPECT() *MockFileValidatorMockRecorder {
	return m.recorder
}

// ValidateFile mocks base method.
func (m *MockFileValidator) ValidateFile(file *models.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateFile", file)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateFile indicates an expected call of ValidateFile.
func (mr *MockFileValidatorMockRecorder) ValidateFile(file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateFile", reflect.TypeOf((*MockFileValidator)(nil).ValidateFile), file)
}

// ValidateFileName mocks base method.
func (m *MockFileValidator) ValidateFileName(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateFileName", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateFileName indicates an expected call of ValidateFileName.
func (mr *MockFileValidatorMockRecorder) ValidateFileName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateFileName", reflect.TypeOf((*MockFileValidator)(nil).ValidateFileName), name)
}

// ValidateFiles mocks base method.
func (m *MockFileValidator) ValidateFiles(files []*models.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateFiles", files)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateFiles indicates an expected call of ValidateFiles.
func (mr *MockFileValidatorMockRecorder) ValidateFiles(files interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateFiles", reflect.TypeOf((*MockFileValidator)(nil).ValidateFiles), files)
}
