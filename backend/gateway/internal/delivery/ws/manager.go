package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	time2 "quickflow/config/time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	http2 "quickflow/gateway/internal/delivery/http"
	"quickflow/gateway/internal/delivery/http/forms"
	forms2 "quickflow/gateway/internal/delivery/ws/forms"
	"quickflow/gateway/utils/validation"
	"quickflow/shared/logger"
	"quickflow/shared/models"
)

const (
	MessageEventRead    = "message_read"
	MessageEventDeleted = "message_delete"
	ChatEventDeleted    = "chat_delete"
	MessageEventSend    = "message"
)

type MessageEvent string

type WSConnectionManager struct {
	Connections map[uuid.UUID]*websocket.Conn
	mu          sync.RWMutex
}

func NewWSConnectionManager() *WSConnectionManager {
	return &WSConnectionManager{
		Connections: make(map[uuid.UUID]*websocket.Conn),
	}
}

// AddConnection adds a new user connection to the manager
func (wm *WSConnectionManager) AddConnection(userId uuid.UUID, conn *websocket.Conn) {
	wm.mu.Lock()
	wm.Connections[userId] = conn
	wm.mu.Unlock()
}

// RemoveAndCloseConnection removes a user connection from the manager and closes it
func (wm *WSConnectionManager) RemoveAndCloseConnection(userId uuid.UUID) {
	wm.mu.Lock()
	if _, exists := wm.Connections[userId]; exists {
		delete(wm.Connections, userId)
	}
	wm.mu.Unlock()
}

func (wm *WSConnectionManager) IsConnected(userId uuid.UUID) (*websocket.Conn, bool) {
	wm.mu.RLock()
	defer wm.mu.RUnlock()
	conn, exists := wm.Connections[userId]
	return conn, exists
}

// ---------------------------------------------------------

type InternalWSMessageHandler struct {
	WSConnectionManager *WSConnectionManager
	MessageUseCase      http2.MessageService
	profileUseCase      http2.ProfileUseCase
	ChatUseCase         http2.ChatUseCase
}

func NewInternalWSMessageHandler(wsConnManager *WSConnectionManager, messageUseCase http2.MessageService, profileUseCase http2.ProfileUseCase, chatUseCase http2.ChatUseCase) *InternalWSMessageHandler {
	return &InternalWSMessageHandler{
		WSConnectionManager: wsConnManager,
		MessageUseCase:      messageUseCase,
		profileUseCase:      profileUseCase,
		ChatUseCase:         chatUseCase,
	}
}

func (m *InternalWSMessageHandler) SendMessage(ctx context.Context, user models.User, payload json.RawMessage) error {
	var messageForm forms.MessageForm
	if err := json.Unmarshal(payload, &messageForm); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}
	if len(messageForm.Text)+len(messageForm.Media)+len(messageForm.Audio)+len(messageForm.File)+len(messageForm.Stickers) == 0 {
		return fmt.Errorf("message cannot be empty")
	}

	messageForm.SenderId = user.Id
	if messageForm.ChatId == uuid.Nil && messageForm.ReceiverId == uuid.Nil {
		logger.Error(ctx, "ChatId and ReceiverId cannot be both nil")
		return fmt.Errorf("chatId and receiverId cannot be both nil")
	}

	message := messageForm.ToMessageModel()
	if err := validation.ValidateMessage(message); err != nil {
		logger.Error(ctx, "Invalid message:", err)
		return fmt.Errorf("invalid message: %w", err)
	}

	var err error
	newMessage, err := m.MessageUseCase.SendMessage(ctx, &message, user.Id)
	if err != nil {
		log.Println("Failed to save message:", err)
		return fmt.Errorf("failed to save message: %w", err)
	}

	// retrieving info to send message to all chat users
	publicSenderInfo, err := m.profileUseCase.GetPublicUserInfo(ctx, user.Id)
	if err != nil {
		log.Println("Failed to get public sender info:", err)
		return fmt.Errorf("failed to get public sender info: %w", err)
	}
	chatParticipants, err := m.ChatUseCase.GetChatParticipants(ctx, newMessage.ChatID)
	if err != nil {
		log.Println("Failed to get chat participants:", err)
		return fmt.Errorf("failed to get chat participants: %w", err)
	}
	err = m.sendMessageToChat(ctx, *newMessage, publicSenderInfo, chatParticipants)
	if err != nil {
		log.Println("Failed to send message to chat:", err)
		return fmt.Errorf("failed to send message to chat: %w", err)
	}

	return nil
}

// SendMessageToChat sends a message to all participants in a chat
func (m *InternalWSMessageHandler) sendMessageToChat(ctx context.Context, message models.Message, publicSenderInfo models.PublicUserInfo, chatParticipants []uuid.UUID) error {
	for _, user := range chatParticipants {
		err := m.notifyMessageEvent(ctx, forms.ToMessageOut(message, publicSenderInfo), user, MessageEventSend)
		if err != nil {
			log.Println("Failed to send message to user:", user, err)
		}
	}

	return nil
}

func (m *InternalWSMessageHandler) MarkMessageRead(ctx context.Context, user models.User, jsonPayload json.RawMessage) error {
	var payload forms2.MarkReadPayload

	if err := json.Unmarshal(jsonPayload, &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	if payload.MessageId == uuid.Nil || payload.ChatId == uuid.Nil {
		return fmt.Errorf("messageId or chatId is empty")
	}

	msg, err := m.MessageUseCase.GetMessageById(ctx, payload.MessageId)
	if err != nil {
		return fmt.Errorf("failed to get message by id: %w", err)
	}

	err = m.MessageUseCase.UpdateLastReadTs(ctx, payload.ChatId, user.Id, msg.CreatedAt, user.Id)
	if err != nil {
		return fmt.Errorf("failed to update last message read: %w", err)
	}

	// send message to message author
	messageReadForm := forms2.NotifyMessageRead{
		MessageId: payload.MessageId,
		Timestamp: msg.CreatedAt.Format(time2.TimeStampLayout),
		ChatId:    payload.ChatId,
		SenderId:  user.Id,
	}

	err = m.notifyMessageEvent(ctx, messageReadForm, msg.SenderID, MessageEventRead)
	if err != nil {
		return fmt.Errorf("failed to notify message read: %w", err)
	}
	return nil
}

func (m *InternalWSMessageHandler) notifyMessageEvent(_ context.Context, read interface{}, receiver uuid.UUID, eventType MessageEvent) error {
	conn, exists := m.WSConnectionManager.IsConnected(receiver)
	if !exists {
		return nil
	}

	out := struct {
		Type string      `json:"type"`
		Data interface{} `json:"payload"`
	}{string(eventType), read}

	msgJSON, err := json.Marshal(out)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	err = conn.WriteMessage(websocket.TextMessage, msgJSON)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}

func (m *InternalWSMessageHandler) DeleteMessage(ctx context.Context, user models.User, jsonPayload json.RawMessage) error {
	var payload forms2.DeleteMessagePayload

	if err := json.Unmarshal(jsonPayload, &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	if payload.MessageId == uuid.Nil {
		return fmt.Errorf("messageId is empty")
	}

	msg, err := m.MessageUseCase.GetMessageById(ctx, payload.MessageId)
	if err != nil {
		return fmt.Errorf("failed to get message by id: %w", err)
	}

	if msg.SenderID != user.Id {
		return fmt.Errorf("user did not send the message")
	}

	err = m.MessageUseCase.DeleteMessage(ctx, payload.MessageId)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}

	// get chat participants
	participants, err := m.ChatUseCase.GetChatParticipants(ctx, msg.ChatID)
	if err != nil {
		return fmt.Errorf("failed to get chat participants: %w", err)
	}

	response := forms2.NotifyDeleteMessage{
		MessageId: payload.MessageId,
		ChatId:    msg.ChatID,
	}

	for _, participant := range participants {
		err := m.notifyMessageEvent(ctx, response, participant, MessageEventDeleted)
		if err != nil {
			return fmt.Errorf("failed to notify message read: %w", err)
		}
	}

	return nil
}

func (m *InternalWSMessageHandler) DeleteChat(ctx context.Context, user models.User, jsonPayload json.RawMessage) error {
	var payload forms2.DeleteChatPayload

	if err := json.Unmarshal(jsonPayload, &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	if payload.ChatId == uuid.Nil {
		return fmt.Errorf("chatId is empty")
	}

	participants, err := m.ChatUseCase.GetChatParticipants(ctx, payload.ChatId)
	if err != nil {
		return fmt.Errorf("failed to get chat participants: %w", err)
	}

	// check if user in participants
	var flag bool
	for _, participant := range participants {
		if participant == user.Id {
			flag = true
			break
		}
	}

	if !flag {
		return fmt.Errorf("user not found in chat")
	}

	err = m.ChatUseCase.DeleteChat(ctx, payload.ChatId)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}

	for _, participant := range participants {
		err := m.notifyMessageEvent(ctx, payload, participant, ChatEventDeleted)
		if err != nil {
			return fmt.Errorf("failed to notify message read: %w", err)
		}
	}
	return nil
}

type PingHandler interface {
	Handle(ctx context.Context, conn *websocket.Conn)
}

// PingHandlerWS - Обработчик Ping сообщений
type PingHandlerWS struct{}

func NewPingHandlerWS() *PingHandlerWS {
	return &PingHandlerWS{}
}

func (wm *PingHandlerWS) Handle(ctx context.Context, conn *websocket.Conn) {
	conn.SetPongHandler(func(appData string) error {
		logger.Info(ctx, "Received pong:", appData)
		return nil
	})

	go func() {
		for {
			time.Sleep(30 * time.Second) // отправка ping каждые 30 секунд
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logger.Info(ctx, "Failed to send ping:", err)
				return
			}
		}
	}()
	return
}
