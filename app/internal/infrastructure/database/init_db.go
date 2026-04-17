package database

import (
	"MuchUp/app/internal/domain/entity"

	"gorm.io/gorm"
)

func InitDB(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.ChatGroup{})
	db.AutoMigrate(&entity.Message{})
	return db
}
