package database

import (
	"MuchUp/app/internal/infrastructure/database/schema"

	"gorm.io/gorm"
)

func InitDB(db *gorm.DB) error {
	err := db.AutoMigrate(&schema.UserSchema{}, &schema.RoomSchema{}, &schema.MessageSchema{})
	if err != nil {
		return err
	}
	return nil
}
