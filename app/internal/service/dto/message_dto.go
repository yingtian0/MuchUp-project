package dto

import "time"

type EventType string

const (
	MessageEvent EventType = "message_event"
)

type SendChatMessageInput struct {
	EventType EventType
	SenderID  string
	RoomID    string
	MessageID string
	Content   string
	CreatedAt time.Time
}
