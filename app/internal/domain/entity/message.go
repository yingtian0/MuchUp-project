package entity

import (
	"errors"
	"strings"
	"time"
)

type UserID string
type RoomID string
type MessageID string

type Message struct {
	ID              MessageID `json:"id"`
	ClientMessageID *string   `json:"client_message_id,omitempty"`
	SenderID        UserID    `json:"user_id"`
	RoomID          RoomID    `json:"room_id"`
	Text            *string   `json:"text,omitempty"`
	MediaURL        *string   `json:"media_url,omitempty"`
	StickerID       *string   `json:"sticker_id,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
	ID     MessageID
	RoomID RoomID

	SenderID   UserID
	SenderType SenderType

	Kind      MessageKind
	Text      *string
	MediaURL  *string
	StickerID *string

	Status MessageStatus

	ClientMessageID *string
	IdempotencyKey  *string

	StreamID *string
	Sequence int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewMessage(userID, roomID, text string) (*Message, error) {
	if len(text) == 0 {
func NewTextMessage(roomID RoomID, senderID UserID, text string, now time.Time) (*Message, error) {
	text = strings.TrimSpace(text)
	if roomID == "" {
		return nil, errors.New("room_id is required")
	}
	if senderID == "" {
		return nil, errors.New("sender_id is required")
	}
	if text == "" {
		return nil, errors.New("text is required")
	}
	if utf8.RuneCountInString(text) > 1000 {
		return nil, errors.New("text is too long")
	}

	now := time.Now()
	return &Message{
		ID:              MessageID(utils.GenerateUUID()),
		ClientMessageID: utils.StringPtr(utils.GenerateUUID()),
		SenderID:        UserID(userID),
		RoomID:          RoomID(roomID),
		Text:            &text,
		CreatedAt:       now,
		UpdatedAt:       now,
	}, nil
}

func (m *Message) CanSendMessage(senderID string) bool {
	if string(m.SenderID) != senderID {
		return false
	}
	if m.Text == nil && m.MediaURL == nil && m.StickerID == nil {
		return false
	}
	if m.Text != nil && len(*m.Text) > 1000 {
		return false
	}
	return true
}
