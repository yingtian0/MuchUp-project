package entity

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

type RoomID string
type MessageID string
type SenderType string
type MessageKind string
type AIMessageKind string
type MessageStatus string

const (
	SenderTypeUser SenderType = "USER"
	SenderTypeAI   SenderType = "AI"
)

const (
	MessageKindText    MessageKind = "TEXT"
	MessageKindMedia   MessageKind = "MEDIA"
	MessageKindSticker MessageKind = "STICKER"
)

const (
	MessageStatusPending MessageStatus = "PENDING"
	MessageStatusSent    MessageStatus = "SENT"
	MessageStatusFailed  MessageStatus = "FAILED"
	MessageStatusDeleted MessageStatus = "DELETED"
)

type Message struct {
	ID              MessageID
	ClientMessageID *string
	SenderID        UserID
	RoomID          RoomID
	Text            *string
	MediaURL        *string
	StickerID       *string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time

	SenderType SenderType

	Kind MessageKind

	Status MessageStatus

	IdempotencyKey *string

	StreamID *string
	Sequence int64
}

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

	return &Message{
		ID:              MessageID(uuid.NewString()),
		ClientMessageID: new(uuid.NewString()),
		SenderID:        senderID,
		RoomID:          roomID,
		Text:            &text,
		SenderType:      SenderTypeUser,
		Kind:            MessageKindText,
		Status:          MessageStatusSent,
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

func ReconstructMessage(
	id MessageID,
	clientMessageID *string,
	senderID UserID,
	roomID RoomID,
	text *string,
	mediaURL *string,
	stickerID *string,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
	senderType SenderType,
	kind MessageKind,
	status MessageStatus,
	idempotencyKey *string,
	streamID *string,
	sequence int64,
) *Message {
	return &Message{
		ID:              id,
		ClientMessageID: clientMessageID,
		SenderID:        senderID,
		RoomID:          roomID,
		Text:            text,
		MediaURL:        mediaURL,
		StickerID:       stickerID,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
		DeletedAt:       deletedAt,
		SenderType:      senderType,
		Kind:            kind,
		Status:          status,
		IdempotencyKey:  idempotencyKey,
		StreamID:        streamID,
		Sequence:        sequence,
	}
}
