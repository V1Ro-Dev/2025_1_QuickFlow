// Code generated by MockGen. DO NOT EDIT.
// Source: .//user_service/internal/usecase/profile-usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models "quickflow/shared/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockProfileRepository is a mock of ProfileRepository interface.
type MockProfileRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProfileRepositoryMockRecorder
}

// MockProfileRepositoryMockRecorder is the mock recorder for MockProfileRepository.
type MockProfileRepositoryMockRecorder struct {
	mock *MockProfileRepository
}

// NewMockProfileRepository creates a new mock instance.
func NewMockProfileRepository(ctrl *gomock.Controller) *MockProfileRepository {
	mock := &MockProfileRepository{ctrl: ctrl}
	mock.recorder = &MockProfileRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProfileRepository) EXPECT() *MockProfileRepositoryMockRecorder {
	return m.recorder
}

// GetProfile mocks base method.
func (m *MockProfileRepository) GetProfile(ctx context.Context, userId uuid.UUID) (models.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", ctx, userId)
	ret0, _ := ret[0].(models.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockProfileRepositoryMockRecorder) GetProfile(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockProfileRepository)(nil).GetProfile), ctx, userId)
}

// GetPublicUserInfo mocks base method.
func (m *MockProfileRepository) GetPublicUserInfo(ctx context.Context, userId uuid.UUID) (models.PublicUserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicUserInfo", ctx, userId)
	ret0, _ := ret[0].(models.PublicUserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicUserInfo indicates an expected call of GetPublicUserInfo.
func (mr *MockProfileRepositoryMockRecorder) GetPublicUserInfo(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicUserInfo", reflect.TypeOf((*MockProfileRepository)(nil).GetPublicUserInfo), ctx, userId)
}

// GetPublicUsersInfo mocks base method.
func (m *MockProfileRepository) GetPublicUsersInfo(ctx context.Context, userIds []uuid.UUID) ([]models.PublicUserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicUsersInfo", ctx, userIds)
	ret0, _ := ret[0].([]models.PublicUserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicUsersInfo indicates an expected call of GetPublicUsersInfo.
func (mr *MockProfileRepositoryMockRecorder) GetPublicUsersInfo(ctx, userIds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicUsersInfo", reflect.TypeOf((*MockProfileRepository)(nil).GetPublicUsersInfo), ctx, userIds)
}

// SaveProfile mocks base method.
func (m *MockProfileRepository) SaveProfile(ctx context.Context, profile models.Profile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveProfile", ctx, profile)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveProfile indicates an expected call of SaveProfile.
func (mr *MockProfileRepositoryMockRecorder) SaveProfile(ctx, profile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveProfile", reflect.TypeOf((*MockProfileRepository)(nil).SaveProfile), ctx, profile)
}

// UpdateLastSeen mocks base method.
func (m *MockProfileRepository) UpdateLastSeen(ctx context.Context, userId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastSeen", ctx, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLastSeen indicates an expected call of UpdateLastSeen.
func (mr *MockProfileRepositoryMockRecorder) UpdateLastSeen(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastSeen", reflect.TypeOf((*MockProfileRepository)(nil).UpdateLastSeen), ctx, userId)
}

// UpdateProfileAvatar mocks base method.
func (m *MockProfileRepository) UpdateProfileAvatar(ctx context.Context, id uuid.UUID, url string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfileAvatar", ctx, id, url)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfileAvatar indicates an expected call of UpdateProfileAvatar.
func (mr *MockProfileRepositoryMockRecorder) UpdateProfileAvatar(ctx, id, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfileAvatar", reflect.TypeOf((*MockProfileRepository)(nil).UpdateProfileAvatar), ctx, id, url)
}

// UpdateProfileCover mocks base method.
func (m *MockProfileRepository) UpdateProfileCover(ctx context.Context, id uuid.UUID, url string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfileCover", ctx, id, url)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfileCover indicates an expected call of UpdateProfileCover.
func (mr *MockProfileRepositoryMockRecorder) UpdateProfileCover(ctx, id, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfileCover", reflect.TypeOf((*MockProfileRepository)(nil).UpdateProfileCover), ctx, id, url)
}

// UpdateProfileTextInfo mocks base method.
func (m *MockProfileRepository) UpdateProfileTextInfo(ctx context.Context, newProfile models.Profile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfileTextInfo", ctx, newProfile)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfileTextInfo indicates an expected call of UpdateProfileTextInfo.
func (mr *MockProfileRepositoryMockRecorder) UpdateProfileTextInfo(ctx, newProfile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfileTextInfo", reflect.TypeOf((*MockProfileRepository)(nil).UpdateProfileTextInfo), ctx, newProfile)
}

// MockFileService is a mock of FileService interface.
type MockFileService struct {
	ctrl     *gomock.Controller
	recorder *MockFileServiceMockRecorder
}

// MockFileServiceMockRecorder is the mock recorder for MockFileService.
type MockFileServiceMockRecorder struct {
	mock *MockFileService
}

// NewMockFileService creates a new mock instance.
func NewMockFileService(ctrl *gomock.Controller) *MockFileService {
	mock := &MockFileService{ctrl: ctrl}
	mock.recorder = &MockFileServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileService) EXPECT() *MockFileServiceMockRecorder {
	return m.recorder
}

// DeleteFile mocks base method.
func (m *MockFileService) DeleteFile(ctx context.Context, filename string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", ctx, filename)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockFileServiceMockRecorder) DeleteFile(ctx, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockFileService)(nil).DeleteFile), ctx, filename)
}

// UploadFile mocks base method.
func (m *MockFileService) UploadFile(ctx context.Context, file *models.File) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", ctx, file)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockFileServiceMockRecorder) UploadFile(ctx, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockFileService)(nil).UploadFile), ctx, file)
}
