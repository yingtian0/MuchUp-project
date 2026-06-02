package mapper

import (
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/infrastructure/database/schema"
)

func ToMessageSchema(message *entity.Message) *schema.MessageSchema {
	senderID := string(message.SenderID)
	text := ""

	if message.Text != nil {
		text = *message.Text
	}

	return &schema.MessageSchema{
		ID:       string(message.ID),
		Text:     text,
		SenderID: &senderID,
		RoomID:   string(message.RoomID),
	}
}
func ToMessageEntity(messageSchema *schema.MessageSchema) *entity.Message {
	var senderID entity.UserID
	if messageSchema.SenderID != nil {
		senderID = entity.UserID(*messageSchema.SenderID)
	}

	return &entity.Message{
		ID:        entity.MessageID(messageSchema.ID),
		Text:      &messageSchema.Text,
		SenderID:  senderID,
		RoomID:    entity.RoomID(messageSchema.RoomID),
		CreatedAt: messageSchema.CreatedAt,
		UpdatedAt: messageSchema.UpdatedAt,
	}
}
