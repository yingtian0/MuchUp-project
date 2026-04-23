package model

type Message struct {
	MessageID string `json:"message_id"`
	SenderID  string `json:"sender_id"`
	RoomID    string `json:"room_id"`
	Content   string `json:"content"`
}

type MessageSentEvent struct {
	EventType   string `json:"event_type"`
	MessageID   string `json:"message_id"`
	UserID      string `json:"user_id"`
	GroupID     string `json:"group_id"`
	RecipientID string `json:"recipient_id,omitempty"`
	Content     string `json:"content"`
	CreatedAt   int64  `json:"created_at"`
	Version     int    `json:"version"`
}
