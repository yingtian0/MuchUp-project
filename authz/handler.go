package authz

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// HandlerFunc adapts a context-aware handler to http.Handler.
type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(r.Context(), w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(payload)
}

// HealthHandler returns a simple 200 response when the request context is active.
func HealthHandler(ctx context.Context, w http.ResponseWriter, _ *http.Request) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

// LoginHandler authenticates an existing user and returns a JWT.
func LoginHandler(store *UserStore) HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		if err := ctx.Err(); err != nil {
			return err
		}

		defer func() {
			_ = r.Body.Close()
		}()

		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return writeJSON(w, http.StatusBadRequest, errorResponse{Code: http.StatusBadRequest, Message: "invalid request body"})
		}

		if req.Email == "" || req.Password == "" {
			return writeJSON(w, http.StatusBadRequest, errorResponse{Code: http.StatusBadRequest, Message: "email and password are required"})
		}

		user, authErr := store.Authenticate(req.Email, req.Password)
		if authErr != nil {
			return writeJSON(w, http.StatusUnauthorized, errorResponse{Code: http.StatusUnauthorized, Message: authErr.Error()})
		}

		token, err := IssueToken(user.ID, user.Username)
		if err != nil {
			return err
		}

		return writeJSON(w, http.StatusOK, AuthResponse{
			Token:    token,
			UserID:   user.ID,
			Username: user.Username,
		})
	}
}

// SignupHandler creates a user and returns a JWT for the new account.
func SignupHandler(store *UserStore) HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		if err := ctx.Err(); err != nil {
			return err
		}

		defer func() {
			_ = r.Body.Close()
		}()

		var req SignupRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return writeJSON(w, http.StatusBadRequest, errorResponse{Code: http.StatusBadRequest, Message: "invalid request body"})
		}

		if req.Email == "" || req.Username == "" || req.Password == "" {
			return writeJSON(w, http.StatusBadRequest, errorResponse{Code: http.StatusBadRequest, Message: "email, username, and password are required"})
		}

		user, err := store.Create(req.Username, req.Email, req.Password)
		if err != nil {
			if errors.Is(err, ErrUserExists) {
				return writeJSON(w, http.StatusConflict, errorResponse{Code: http.StatusConflict, Message: err.Error()})
			}

			return err
		}

		token, err := IssueToken(user.ID, user.Username)
		if err != nil {
			return err
		}

		return writeJSON(w, http.StatusOK, AuthResponse{
			Token:    token,
			UserID:   user.ID,
			Username: user.Username,
		})
	}
}
