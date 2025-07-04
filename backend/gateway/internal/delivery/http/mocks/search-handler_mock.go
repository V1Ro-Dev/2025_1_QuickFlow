// Code generated by MockGen. DO NOT EDIT.
// Source: .//gateway/internal/delivery/http/search-handler.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models "quickflow/shared/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSearchUseCase is a mock of SearchUseCase interface.
type MockSearchUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSearchUseCaseMockRecorder
}

// MockSearchUseCaseMockRecorder is the mock recorder for MockSearchUseCase.
type MockSearchUseCaseMockRecorder struct {
	mock *MockSearchUseCase
}

// NewMockSearchUseCase creates a new mock instance.
func NewMockSearchUseCase(ctrl *gomock.Controller) *MockSearchUseCase {
	mock := &MockSearchUseCase{ctrl: ctrl}
	mock.recorder = &MockSearchUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchUseCase) EXPECT() *MockSearchUseCaseMockRecorder {
	return m.recorder
}

// SearchSimilarUser mocks base method.
func (m *MockSearchUseCase) SearchSimilarUser(ctx context.Context, toSearch string, postsCount uint) ([]models.PublicUserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchSimilarUser", ctx, toSearch, postsCount)
	ret0, _ := ret[0].([]models.PublicUserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchSimilarUser indicates an expected call of SearchSimilarUser.
func (mr *MockSearchUseCaseMockRecorder) SearchSimilarUser(ctx, toSearch, postsCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchSimilarUser", reflect.TypeOf((*MockSearchUseCase)(nil).SearchSimilarUser), ctx, toSearch, postsCount)
}
