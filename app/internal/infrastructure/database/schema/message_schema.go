package schema

import (
	"time"

	"gorm.io/gorm"
)

type MessageSchema struct {
	ID        string     `gorm:"type:uuid;primaryKey"`
	Text      string     `gorm:"type:text;not null"`
	SenderID  *string    `gorm:"type:uuid"`
	RoomID    string     `gorm:"type:uuid"`
	Room      RoomSchema `gorm:"foreignKey:RoomID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (MessageSchema) TableName() string {
	return "messages"
}
