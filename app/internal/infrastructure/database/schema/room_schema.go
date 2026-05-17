package schema

import (
	"time"

	"gorm.io/gorm"
)

type RoomSchema struct {
	ID        string       `gorm:"type:uuid;primaryKey"`
	Type      string       `gorm:"type:varchar(50);not null"`
	Status    string       `gorm:"type:varchar(50);not null"`
	Capacity  int          `gorm:"not null"`
	CreatedBy string       `gorm:"type:uuid;not null"`
	Users     []UserSchema `gorm:"many2many:user_rooms;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (RoomSchema) TableName() string {
	return "rooms"
}
