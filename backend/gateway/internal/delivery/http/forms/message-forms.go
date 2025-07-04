package forms

import (
	"errors"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"

	time2 "quickflow/config/time"
	"quickflow/shared/models"
)

//easyjson:json
type GetMessagesForm struct {
	MessagesCount int       `json:"messages_count"`
	Ts            time.Time `json:"ts,omitempty"`
}

func (m *GetMessagesForm) GetParams(values url.Values) error {
	var (
		err         error
		numMessages int64
	)

	if !values.Has("messages_count") {
		return errors.New("messages_count parameter missing")
	}

	numMessages, err = strconv.ParseInt(values.Get("messages_count"), 10, 64)
	if err != nil {
		return errors.New("failed to parse messages_count")
	}

	m.MessagesCount = int(numMessages)

	ts, err := time.Parse(time2.TimeStampLayout, values.Get("ts"))
	if err != nil {
		ts = time.Now()
	}
	m.Ts = ts
	return nil
}

type FileOut struct {
	URL  string `json:"url"`
	Name string `json:"name,omitempty"`
}

func ToFileOut(file models.File) FileOut {
	return FileOut{
		URL:  file.URL,
		Name: file.Name,
	}
}

//easyjson:json
type MessageOut struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Text        string    `json:"text"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	MediaURLs   []FileOut `json:"media,omitempty"`
	AudioURLs   []FileOut `json:"audio,omitempty"`
	FileURLs    []FileOut `json:"files,omitempty"`
	StickerUrls []FileOut `json:"stickers,omitempty"`

	Sender PublicUserInfoOut `json:"sender"`
	ChatId uuid.UUID         `json:"chat_id"`
}

func ToMessageOut(message models.Message, info models.PublicUserInfo) MessageOut {
	mediaURLs := make([]FileOut, 0)
	audioURLs := make([]FileOut, 0)
	fileURLs := make([]FileOut, 0)
	stickerUrls := make([]FileOut, 0)

	for _, file := range message.Attachments {
		if file.DisplayType == models.DisplayTypeMedia {
			mediaURLs = append(mediaURLs, ToFileOut(*file))
		} else if file.DisplayType == models.DisplayTypeAudio {
			audioURLs = append(audioURLs, ToFileOut(*file))
		} else if file.DisplayType == models.DisplayTypeSticker {
			stickerUrls = append(stickerUrls, ToFileOut(*file))
		} else {
			fileURLs = append(fileURLs, ToFileOut(*file))
		}
	}

	return MessageOut{
		ID:          message.ID,
		Text:        message.Text,
		CreatedAt:   message.CreatedAt.Format(time2.TimeStampLayout),
		UpdatedAt:   message.UpdatedAt.Format(time2.TimeStampLayout),
		MediaURLs:   mediaURLs,
		AudioURLs:   audioURLs,
		FileURLs:    fileURLs,
		StickerUrls: stickerUrls,

		Sender: PublicUserInfoToOut(info, ""),
		ChatId: message.ChatID,
	}
}

func ToMessagesOut(messages []*models.Message, usersInfo map[uuid.UUID]models.PublicUserInfo) []MessageOut {
	var messagesOut []MessageOut
	for _, message := range messages {
		messagesOut = append(messagesOut, ToMessageOut(*message, usersInfo[message.SenderID]))
	}

	return messagesOut
}

//easyjson:json
type MessagesOut struct {
	Messages   []MessageOut `json:"messages"`
	LastReadTs string       `json:"last_read_ts,omitempty"`
}

//easyjson:json
type MessageForm struct {
	Text       string    `form:"text" json:"text,omitempty"`
	ChatId     uuid.UUID `form:"chat_id" json:"chat_id,omitempty"`
	Media      []string  `form:"media" json:"media,omitempty"`
	Audio      []string  `form:"audio" json:"audio,omitempty"`
	File       []string  `form:"files" json:"files,omitempty"`
	Stickers   []string  `form:"stickers" json:"stickers,omitempty"`
	ReceiverId uuid.UUID `json:"receiver_id,omitempty"`
	SenderId   uuid.UUID `json:"-"`
}

func (f *MessageForm) ToMessageModel() models.Message {
	var attachments []*models.File
	for _, file := range f.Media {
		attachments = append(attachments, &models.File{
			URL:         file,
			DisplayType: models.DisplayTypeMedia,
		})
	}

	for _, file := range f.Audio {
		attachments = append(attachments, &models.File{
			URL:         file,
			DisplayType: models.DisplayTypeAudio,
		})
	}

	for _, file := range f.File {
		attachments = append(attachments, &models.File{
			URL:         file,
			DisplayType: models.DisplayTypeFile,
		})
	}

	for _, file := range f.Stickers {
		attachments = append(attachments, &models.File{
			URL:         file,
			DisplayType: models.DisplayTypeSticker,
		})
	}

	return models.Message{
		ID:          uuid.New(),
		Text:        f.Text,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Attachments: attachments,
		ReceiverID:  f.ReceiverId,
		SenderID:    f.SenderId,
		ChatID:      f.ChatId,
	}
}
