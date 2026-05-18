package entity

import (
	"MuchUp/app/utils"
	"errors"
	"strings"
	"time"
	"unicode/utf8"
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
	AIMessageKindIcebreakInitialTopic  AIMessageKind = "ICEBREAK_INITIAL_TOPIC"
	AIMessageKindIcebreakFollowUp      AIMessageKind = "ICEBREAK_FOLLOW_UP"
	AIMessageKindIcebreakSilencePrompt AIMessageKind = "ICEBREAK_SILENCE_PROMPT"
	AIMessageKindIcebreakMission       AIMessageKind = "ICEBREAK_MISSION"
	AIMessageKindIcebreakSummary       AIMessageKind = "ICEBREAK_SUMMARY"
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

	Kind   MessageKind
	AIKind *AIMessageKind

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
		ID:              MessageID(utils.GenerateUUID()),
		ClientMessageID: utils.StringPtr(utils.GenerateUUID()),
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
