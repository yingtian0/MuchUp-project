package authzservice

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserExists = errors.New("user already exists")
var ErrInvalidCredentials = errors.New("invalid credentials")

type User struct {
	ID           string
	Username     string
	Email        string
	PasswordHash []byte
}

type UserStore struct {
	mu     sync.RWMutex
	byMail map[string]User
}

func NewUserStore() *UserStore {
	return &UserStore{byMail: make(map[string]User)}
}

func (s *UserStore) Create(username, email, password string) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.byMail[email]; exists {
		return User{}, ErrUserExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	user := User{
		ID:           uuid.NewString(),
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
	}

	s.byMail[email] = user
	return user, nil
}

func (s *UserStore) Authenticate(email, password string) (User, error) {
	s.mu.RLock()
	user, ok := s.byMail[email]
	s.mu.RUnlock()
	if !ok {
		return User{}, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		return User{}, ErrInvalidCredentials
	}

	return user, nil
}
