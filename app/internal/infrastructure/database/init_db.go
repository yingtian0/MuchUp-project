package database

import (
	"MuchUp/app/internal/domain/entity"

	"gorm.io/gorm"
)

func InitDB(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.User{}, &entity.ChatGroup{}, &entity.Message{})
	if err != nil {
		return err
	}
	return nil
}
