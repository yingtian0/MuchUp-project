package mapper

import (
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/infrastructure/database/schema"
	"encoding/json"
)

func ToUserSchema(user *entity.User) *schema.UserSchema {

	return &schema.UserSchema{
		ID:                user.ID,
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
func ToUserEntity(userSchema *schema.UserSchema) *entity.User {
	var profile map[string]interface{}
	if userSchema.PersonalityProfile != nil {
		_ = json.Unmarshal(userSchema.PersonalityProfile, &profile)
	}
	return &entity.User{
		ID:            userSchema.ID,
		NickName:      userSchema.NickName,
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
