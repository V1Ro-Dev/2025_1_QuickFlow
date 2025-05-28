package http_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"quickflow/shared/logger"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	http2 "quickflow/gateway/internal/delivery/http"
	"quickflow/gateway/internal/delivery/http/mocks"
	"quickflow/shared/models"
)

func TestFriendsHandler_GetFriends(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFriendsUseCase := mocks.NewMockFriendsUseCase(ctrl)
	mockWS := mocks.NewMockIWebSocketConnectionManager(ctrl)
	handler := http2.NewFriendsHandler(mockFriendsUseCase, mockWS)

	t.Run("OK (Current User)", func(t *testing.T) {
		userID := uuid.New()
		mockFriendsUseCase.EXPECT().
			GetFriendsInfo(gomock.Any(), userID.String(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.FriendInfo{}, 0, nil)
		mockWS.EXPECT().IsConnected(gomock.Any()).Return(nil, false).AnyTimes()

		req := httptest.NewRequest(http.MethodGet, "/api/friends", nil)
		ctx := context.WithValue(req.Context(), logger.Username, models.User{Id: userID, Username: "testuser"})
		req = req.WithContext(ctx)
		rr := httptest.NewRecorder()

		handler.GetFriends(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("OK (Specific User)", func(t *testing.T) {
		userID := uuid.New()
		targetUserID := uuid.New()
		mockFriendsUseCase.EXPECT().
			GetFriendsInfo(gomock.Any(), targetUserID.String(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.FriendInfo{}, 0, nil)
		mockWS.EXPECT().IsConnected(gomock.Any()).Return(nil, false).AnyTimes()

		req := httptest.NewRequest(http.MethodGet, "/api/friends?user_id="+targetUserID.String(), nil)
		ctx := context.WithValue(req.Context(), logger.Username, models.User{Id: userID, Username: "testuser"})
		req = req.WithContext(ctx)
		rr := httptest.NewRecorder()

		handler.GetFriends(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Error in UseCase", func(t *testing.T) {
		userID := uuid.New()
		mockFriendsUseCase.EXPECT().
			GetFriendsInfo(gomock.Any(), userID.String(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, 0, errors.New("some error"))

		req := httptest.NewRequest(http.MethodGet, "/api/friends", nil)
		ctx := context.WithValue(req.Context(), logger.Username, models.User{Id: userID, Username: "testuser"})
		req = req.WithContext(ctx)
		rr := httptest.NewRecorder()

		handler.GetFriends(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("No User in Context", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/friends", nil)
		rr := httptest.NewRecorder()

		handler.GetFriends(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestFriendsHandler_SendFriendRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFriendsUseCase := mocks.NewMockFriendsUseCase(ctrl)
	mockWS := mocks.NewMockIWebSocketConnectionManager(ctrl)
	handler := http2.NewFriendsHandler(mockFriendsUseCase, mockWS)

	userID := uuid.New()
	receiverID := uuid.New()

	testCases := []struct {
		name               string
		ctxUser            *models.User
		inputBody          string
		mockBehavior       func()
		expectedStatusCode int
	}{
		{
			name:      "OK (Success)",
			ctxUser:   &models.User{Id: userID, Username: "testuser"},
			inputBody: fmt.Sprintf(`{"receiver_id":"%s"}`, receiverID.String()),
			mockBehavior: func() {
				mockFriendsUseCase.EXPECT().
					SendFriendRequest(gomock.Any(), userID.String(), receiverID.String()).
					Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:      "Request Already Exists",
			ctxUser:   &models.User{Id: userID, Username: "testuser"},
			inputBody: fmt.Sprintf(`{"receiver_id":"%s"}`, receiverID.String()),
			mockBehavior: func() {
				mockFriendsUseCase.EXPECT().
					SendFriendRequest(gomock.Any(), userID.String(), receiverID.String()).
					Return(errors.New("request already exists"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:      "Bad JSON",
			ctxUser:   &models.User{Id: userID, Username: "testuser"},
			inputBody: `{"receiver_id": "invalid",`,
			mockBehavior: func() {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:      "Error in SendFriendRequest",
			ctxUser:   &models.User{Id: userID, Username: "testuser"},
			inputBody: fmt.Sprintf(`{"receiver_id":"%s"}`, receiverID.String()),
			mockBehavior: func() {
				mockFriendsUseCase.EXPECT().
					SendFriendRequest(gomock.Any(), userID.String(), receiverID.String()).
					Return(errors.New("internal error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()

			req := httptest.NewRequest(http.MethodPost, "/api/friends", bytes.NewBufferString(tc.inputBody))
			if tc.ctxUser != nil {
				ctx := context.WithValue(req.Context(), logger.Username, *tc.ctxUser)
				req = req.WithContext(ctx)
			}

			rr := httptest.NewRecorder()
			handler.SendFriendRequest(rr, req)

			assert.Equal(t, tc.expectedStatusCode, rr.Code)
		})
	}
}

func TestFriendsHandler_AcceptFriendRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFriendsUseCase := mocks.NewMockFriendsUseCase(ctrl)
	mockWS := mocks.NewMockIWebSocketConnectionManager(ctrl)
	handler := http2.NewFriendsHandler(mockFriendsUseCase, mockWS)

	userID := uuid.New()
	receiverID := uuid.New()

	testCases := []struct {
		name               string
		ctxUser            *models.User
		inputBody          string
		mockBehavior       func()
		expectedStatusCode int
	}{
		{
			name:      "OK (Success)",
			ctxUser:   &models.User{Id: userID, Username: "testuser"},
			inputBody: fmt.Sprintf(`{"receiver_id":"%s"}`, receiverID.String()),
			mockBehavior: func() {
				mockFriendsUseCase.EXPECT().
					AcceptFriendRequest(gomock.Any(), userID.String(), receiverID.String()).
					Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:      "Bad JSON",
			ctxUser:   &models.User{Id: userID, Username: "testuser"},
			inputBody: `{"receiver_id": "invalid",`,
			mockBehavior: func() {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:      "Error in AcceptFriendRequest",
			ctxUser:   &models.User{Id: userID, Username: "testuser"},
			inputBody: fmt.Sprintf(`{"receiver_id":"%s"}`, receiverID.String()),
			mockBehavior: func() {
				mockFriendsUseCase.EXPECT().
					AcceptFriendRequest(gomock.Any(), userID.String(), receiverID.String()).
					Return(errors.New("internal error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()

			req := httptest.NewRequest(http.MethodPost, "/api/friends/accept", bytes.NewBufferString(tc.inputBody))
			if tc.ctxUser != nil {
				ctx := context.WithValue(req.Context(), logger.Username, *tc.ctxUser)
				req = req.WithContext(ctx)
			}

			rr := httptest.NewRecorder()
			handler.AcceptFriendRequest(rr, req)

			assert.Equal(t, tc.expectedStatusCode, rr.Code)
		})
	}
}

func TestFriendsHandler_DeleteFriend(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFriendsUseCase := mocks.NewMockFriendsUseCase(ctrl)
	mockWS := mocks.NewMockIWebSocketConnectionManager(ctrl)
	handler := http2.NewFriendsHandler(mockFriendsUseCase, mockWS)

	userID := uuid.New()
	friendID := uuid.New()

	testCases := []struct {
		name               string
		ctxUser            *models.User
		inputBody          string
		mockBehavior       func()
		expectedStatusCode int
	}{
		{
			name:      "OK (Success)",
			ctxUser:   &models.User{Id: userID, Username: "testuser"},
			inputBody: fmt.Sprintf(`{"friend_id":"%s"}`, friendID.String()),
			mockBehavior: func() {
				mockFriendsUseCase.EXPECT().
					DeleteFriend(gomock.Any(), userID.String(), friendID.String()).
					Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:      "Bad JSON",
			ctxUser:   &models.User{Id: userID, Username: "testuser"},
			inputBody: `{"friend_id": "invalid",`,
			mockBehavior: func() {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:      "Error in DeleteFriend",
			ctxUser:   &models.User{Id: userID, Username: "testuser"},
			inputBody: fmt.Sprintf(`{"friend_id":"%s"}`, friendID.String()),
			mockBehavior: func() {
				mockFriendsUseCase.EXPECT().
					DeleteFriend(gomock.Any(), userID.String(), friendID.String()).
					Return(errors.New("internal error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()

			req := httptest.NewRequest(http.MethodDelete, "/api/friends", bytes.NewBufferString(tc.inputBody))
			if tc.ctxUser != nil {
				ctx := context.WithValue(req.Context(), logger.Username, *tc.ctxUser)
				req = req.WithContext(ctx)
			}

			rr := httptest.NewRecorder()
			handler.DeleteFriend(rr, req)

			assert.Equal(t, tc.expectedStatusCode, rr.Code)
		})
	}
}

func TestFriendsHandler_Unfollow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFriendsUseCase := mocks.NewMockFriendsUseCase(ctrl)
	mockWS := mocks.NewMockIWebSocketConnectionManager(ctrl)
	handler := http2.NewFriendsHandler(mockFriendsUseCase, mockWS)

	userID := uuid.New()
	friendID := uuid.New()

	testCases := []struct {
		name               string
		ctxUser            *models.User
		inputBody          string
		mockBehavior       func()
		expectedStatusCode int
	}{
		{
			name:      "OK (Success)",
			ctxUser:   &models.User{Id: userID, Username: "testuser"},
			inputBody: fmt.Sprintf(`{"friend_id":"%s"}`, friendID.String()),
			mockBehavior: func() {
				mockFriendsUseCase.EXPECT().
					Unfollow(gomock.Any(), userID.String(), friendID.String()).
					Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:      "Bad JSON",
			ctxUser:   &models.User{Id: userID, Username: "testuser"},
			inputBody: `{"friend_id": "invalid",`,
			mockBehavior: func() {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:      "Error in Unfollow",
			ctxUser:   &models.User{Id: userID, Username: "testuser"},
			inputBody: fmt.Sprintf(`{"friend_id":"%s"}`, friendID.String()),
			mockBehavior: func() {
				mockFriendsUseCase.EXPECT().
					Unfollow(gomock.Any(), userID.String(), friendID.String()).
					Return(errors.New("internal error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()

			req := httptest.NewRequest(http.MethodPost, "/api/friends/unfollow", bytes.NewBufferString(tc.inputBody))
			if tc.ctxUser != nil {
				ctx := context.WithValue(req.Context(), logger.Username, *tc.ctxUser)
				req = req.WithContext(ctx)
			}

			rr := httptest.NewRecorder()
			handler.Unfollow(rr, req)

			assert.Equal(t, tc.expectedStatusCode, rr.Code)
		})
	}
}
