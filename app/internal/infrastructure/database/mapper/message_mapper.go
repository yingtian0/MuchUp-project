package mapper

import (
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/infrastructure/database/schema"
)

func ToMessageSchema(message *entity.Message) *schema.MessageSchema {
	return &schema.MessageSchema{
		ID:       message.MessageID,
		Text:     *message.Text,
		SenderID: &message.SenderID,
		GroupID:  message.GroupID,
	}
}
func ToMessageEntity(messageSchema *schema.MessageSchema) *entity.Message {
	return &entity.Message{
		MessageID: messageSchema.ID,
		Text:      &messageSchema.Text,
		SenderID:  *messageSchema.SenderID,
		GroupID:   messageSchema.GroupID,
		CreatedAt: messageSchema.CreatedAt,
		UpdatedAt: messageSchema.UpdatedAt,
	}
}
