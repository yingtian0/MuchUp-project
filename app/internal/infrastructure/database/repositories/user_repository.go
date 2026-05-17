package repositories

import (
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
	"MuchUp/app/internal/infrastructure/database/mapper"
	"MuchUp/app/internal/infrastructure/database/schema"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) CreateUser(user *entity.User) error {
	userShema := mapper.ToUserSchema(user)
	err := r.db.Create(userShema).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	var userSchema schema.UserSchema
	err := r.db.Where("email = ?", email).First(&userSchema).Error
	if err != nil {
		return nil, err
	}
	return mapper.ToUserEntity(&userSchema), nil
}
func (r *userRepository) GetUserByPhone(phone string) (*entity.User, error) {
	var userSchema schema.UserSchema
	err := r.db.Where("phone = ?", phone).First(&userSchema).Error
	if err != nil {
		return nil, err
	}
	return mapper.ToUserEntity(&userSchema), nil
}
func (r *userRepository) GetUserByID(id string) (*entity.User, error) {
	var userSchema schema.UserSchema
	err := r.db.Where("id = ?", id).First(&userSchema).Error
	if err != nil {
		return nil, err
	}
	return mapper.ToUserEntity(&userSchema), nil
}
func (r *userRepository) UpdateUser(user *entity.User) error {
	return r.db.Save(mapper.ToUserSchema(user)).Error
}
func (r *userRepository) DeleteUser(id string) error {
	return r.db.Delete(&schema.UserSchema{}, "id = ?", id).Error
}
func (r *userRepository) GetUsers(limit, offset int) ([]*entity.User, error) {
	return nil, nil
}
func (r *userRepository) GetUsersByRoom(roomID string) ([]*entity.User, error) {
	var userSchemas []schema.UserSchema
	err := r.db.
		Joins("JOIN user_rooms ON user_rooms.user_id = users.id").
		Where("user_rooms.room_id = ?", roomID).
		Find(&userSchemas).Error
	if err != nil {
		return nil, err
	}

	users := make([]*entity.User, 0, len(userSchemas))
	for i := range userSchemas {
		users = append(users, mapper.ToUserEntity(&userSchemas[i]))
	}
	return users, nil
}
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}
