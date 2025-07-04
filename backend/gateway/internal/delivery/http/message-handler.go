package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/microcosm-cc/bluemonday"

	time2 "quickflow/config/time"
	"quickflow/gateway/internal/delivery/http/forms"
	errors2 "quickflow/gateway/internal/errors"
	http2 "quickflow/gateway/utils/http"
	"quickflow/shared/logger"
	"quickflow/shared/models"
)

type MessageHandler struct {
	messageUseCase MessageService
	authUseCase    AuthUseCase
	profileUseCase ProfileUseCase
	policy         *bluemonday.Policy
}

func NewMessageHandler(messageUseCase MessageService, authUseCase AuthUseCase, profileUseCase ProfileUseCase, policy *bluemonday.Policy) *MessageHandler {
	return &MessageHandler{
		messageUseCase: messageUseCase,
		authUseCase:    authUseCase,
		profileUseCase: profileUseCase,
		policy:         policy,
	}
}

// GetMessagesForChat returns messages for a specific chat
// @Summary Get messages for chat
// @Description Get messages for a specific chat
// @Tags Messages
// @Accept json
// @Produce json
// @Param chat_id path string true "Chat ID"
// @Param posts_count query int true "Number of messages"
// @Param ts query string false "Timestamp"
// @Success 200 {array} forms.MessageOut "List of messages"
// @Failure 400 {object} forms.ErrorForm "Invalid data"
// @Failure 403 {object} forms.ErrorForm "User is not a participant in the chat"
// @Failure 500 {object} forms.ErrorForm "Server error"
// @Router /api/chats/{chat_id}/messages [get]
// GetMessagesForChat godoc
func (m *MessageHandler) GetMessagesForChat(w http.ResponseWriter, r *http.Request) {
	ctx := http2.SetRequestId(r.Context())
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		logger.Error(ctx, "Failed to get user from context while fetching messages")
		http2.WriteJSONError(w, errors2.New(errors2.InternalErrorCode, "Failed to get user from context", http.StatusInternalServerError))
		return
	}

	chat := mux.Vars(r)["chat_id"]
	if len(chat) == 0 {
		logger.Info(ctx, "Fetch messages request without chat_id")
		http2.WriteJSONError(w, errors2.New(errors2.BadRequestErrorCode, "chat_id is required", http.StatusBadRequest))
		return
	}

	chatId, err := uuid.Parse(chat)
	if err != nil {
		logger.Info(ctx, "Failed to parse chat_id (uuid: %s): %s", chat, err.Error())
		http2.WriteJSONError(w, errors2.New(errors2.BadRequestErrorCode, "chat_id is not valid", http.StatusBadRequest))
		return
	}

	var messageForm forms.GetMessagesForm
	err = messageForm.GetParams(r.URL.Query())
	if err != nil {
		logger.Error(ctx, "Failed to parse query params: %v", err)
		http2.WriteJSONError(w, errors2.New(errors2.BadRequestErrorCode, "Failed to parse query params", http.StatusBadRequest))
		return
	}

	logger.Info(ctx, "Fetching feed for user %s with %d posts with timestamp %v (autogenerated: %t)", user.Username, messageForm.MessagesCount, messageForm.Ts, !r.URL.Query().Has("ts"))

	messages, err := m.messageUseCase.GetMessagesForChat(ctx, chatId, messageForm.MessagesCount, messageForm.Ts, user.Id)
	if err != nil {
		err := errors2.FromGRPCError(err)
		logger.Error(ctx, "Failed to fetch messages: %v", err)
		http2.WriteJSONError(w, err)
		return
	}

	logger.Info(ctx, "Fetched %d messages for user %s", len(messages), user.Username)

	senderIds := make([]uuid.UUID, 0, len(messages))
	for _, message := range messages {
		senderIds = append(senderIds, message.SenderID)
	}

	publicInfoMap := make(map[uuid.UUID]models.PublicUserInfo)
	if len(senderIds) != 0 {
		publicInfo, err := m.profileUseCase.GetPublicUsersInfo(ctx, senderIds)
		if err != nil {
			err := errors2.FromGRPCError(err)
			logger.Error(ctx, "Error while fetching last messages users info: %v", err)
			http2.WriteJSONError(w, err)
			return
		}
		for _, info := range publicInfo {
			publicInfoMap[info.Id] = info
		}
	}

	getLastReadTs, err := m.messageUseCase.GetLastReadTs(ctx, chatId, user.Id)

	out := forms.MessagesOut{
		Messages: forms.ToMessagesOut(messages, publicInfoMap),
	}
	if err == nil {
		out.LastReadTs = getLastReadTs.Format(time2.TimeStampLayout)
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := easyjson.MarshalToWriter(out, w); err != nil {
		logger.Error(ctx, "Failed to encode feed: %s", err.Error())
		http2.WriteJSONError(w, errors2.New(errors2.InternalErrorCode, "Failed to encode feed", http.StatusInternalServerError))
		return
	}
}
