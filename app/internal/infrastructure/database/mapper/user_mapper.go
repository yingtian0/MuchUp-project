package mapper

import (
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/infrastructure/database/schema"
)

func ToUserSchema(user *entity.UserProfile) *schema.UserSchema {
	return &schema.UserSchema{
		ID:                string(user.ID),
		NickName:          user.NickName,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		PasswordHash:      user.PasswordHash,
		EmailVerified:     user.EmailVerified,
		PhoneVerified:     user.PhoneVerified,
		PrimaryAuthMethod: string(user.AuthMethod),
		AvatarURL:         user.AvatarURL,
		UsagePurpose:      user.UsagePurpose,
		IsActive:          user.IsActive,
		IsBanned:          user.IsBanned,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
	}
}
func ToUserEntity(userSchema *schema.UserSchema) *entity.UserProfile {
	return &entity.UserProfile{
		ID:            entity.UserID(userSchema.ID),
		NickName:      userSchema.NickName,
		DisplayName:   userSchema.NickName,
		Email:         userSchema.Email,
		PhoneNumber:   userSchema.PhoneNumber,
		PasswordHash:  userSchema.PasswordHash,
		AvatarURL:     userSchema.AvatarURL,
		UsagePurpose:  userSchema.UsagePurpose,
		IsActive:      userSchema.IsActive,
		IsBanned:      userSchema.IsBanned,
		EmailVerified: userSchema.EmailVerified,
		PhoneVerified: userSchema.PhoneVerified,
		AuthMethod:    entity.PrimaryAuthMethod(userSchema.PrimaryAuthMethod),
		CreatedAt:     userSchema.CreatedAt,
		UpdatedAt:     userSchema.UpdatedAt,
	}
}
