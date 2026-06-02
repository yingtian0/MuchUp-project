package authzservice

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var signingKey = mustLoadSigningKey()

var tokenTTL = 1 * time.Hour

var secretKey []byte

func init() {
	if v := os.Getenv("JWT_SECRET"); v != "" {
		secretKey = []byte(v)
	}

	if v := os.Getenv("JWT_TTL"); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			tokenTTL = d
		}
	}
}

func mustLoadSigningKey() []byte {
	if value := os.Getenv("JWT_SECRET"); value != "" {
		return []byte(value)
	}

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(fmt.Sprintf("failed to generate JWT signing key: %v", err))
	}

	return key
}

// CustomClaim contains MuchUp-specific JWT claims.
type CustomClaim struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// IssueToken signs a JWT for the provided user identity.
func IssueToken(userID, username string) (string, error) {
	now := time.Now()
	expiresAt := now.Add(tokenTTL)
	claims := CustomClaim{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signed, nil
}
