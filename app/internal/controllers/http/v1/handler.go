package rest

import (
	"MuchUp/app/internal/controllers/http/ws"
	"MuchUp/app/internal/domain/usecase"
	"MuchUp/app/pkg/logger"
	"os"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Handler struct {
	userUsecase    usecase.UserUsecase
	messageUsecase usecase.MessageUsecase
	logger         logger.Logger
	hub            *ws.Hub
}
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}
type CreateUserRequest struct {
	Name               string                 `json:"name" validate:"required,min=2,max=50"`
	Email              string                 `json:"email" validate:"required,email"`
	PasswordHash       string                 `json:"password" validate:"required,min=8"`
	UsagePurpose       string                 `json:"usage_purpose"`
	PersonalityProfile map[string]interface{} `json:"personality_profile"`
}
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

var googleOAuthConfig = &oauth2.Config{
	ClientID:     os.Getenv("CLIENT_ID"),
	ClientSecret: os.Getenv("CLIENT_SECRET"),
	RedirectURL:  "https://localhost:8080/api/v1/auth/google/callback",
	Scopes:       []string{"email", "openid"},
	Endpoint:     google.Endpoint,
}

type CreateMessageRequest struct {
	Content string `json:"content" validate:"required,min=1,max=1000"`
	GroupID string `json:"group_id" validate:"required"`
}
type UpdateUserRequest struct {
	Name  string `json:"name,omitempty" validate:"omitempty,min=2,max=50"`
	Email string `json:"email,omitempty" validate:"omitempty,email"`
}

func NewHandler(

	userUsecase usecase.UserUsecase,
	messageUsecase usecase.MessageUsecase,
	logger logger.Logger,

) *Handler {
	hub := &ws.Hub{
		Clients:    make(map[*ws.Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *ws.Client),
		Unregister: make(chan *ws.Client),
	}

	go hub.Run()

	return &Handler{
		userUsecase:    userUsecase,
		messageUsecase: messageUsecase,
		logger:         logger,
		hub:            hub,
	}
}

func (h *Handler) SetupRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}
