package entity

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

type MessageID string

type SenderType string

const (
	SenderTypeUser   SenderType = "USER"
	SenderTypeAI     SenderType = "AI"
	SenderTypeSystem SenderType = "SYSTEM"
)

type MessageStatus string

const (
	MessagePendingModeration MessageStatus = "PENDING_MODERATION"
	MessageVisible           MessageStatus = "VISIBLE"
	MessageHidden            MessageStatus = "HIDDEN"
	MessageDeleted           MessageStatus = "DELETED"
)

type MessageKind string

const (
	MessageKindText    MessageKind = "TEXT"
	MessageKindImage   MessageKind = "IMAGE"
	MessageKindSticker MessageKind = "STICKER"
	MessageKindSystem  MessageKind = "SYSTEM"
)

type Message struct {
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
		ID:         MessageID(uuid.NewString()),
		RoomID:     roomID,
		SenderID:   senderID,
		SenderType: SenderTypeUser,
		Kind:       MessageKindText,
		Text:       &text,
		Status:     MessageVisible,
		CreatedAt:  now,
		UpdatedAt:  now,
	}, nil
}
