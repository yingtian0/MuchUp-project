package auth

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type TokenValidator interface {
	RefreshToken(ctx context.Context, tokenString string) (string, error)
	ValidateToken(ctx context.Context, tokenString string) (*JWTClaims, error)
}

type RoomMemberValidator interface {
}

type JWTClaims struct {
	UserID  string `json:"user_id"`
	GroupID string `json:"group_id"`
	jwt.RegisteredClaims
}
