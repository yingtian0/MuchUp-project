package usecase

import (
	"MuchUp/app/internal/domain/entity"
)

type UserUsecase interface {
	GetUserByID(id string) (*entity.UserProfile, error)
	GetUserByEmail(email string) (*entity.UserProfile, error)
	CreateUser(user *entity.UserProfile) (*entity.UserProfile, error)
	UpdateUser(user *entity.UserProfile) (*entity.UserProfile, error)
	DeleteUser(id string) error
	GetUsers(limit, offset uint) ([]*entity.UserProfile, error)
}
