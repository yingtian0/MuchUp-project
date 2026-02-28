package authzservice

import (
	"context"
	"encoding/json"
	"net/http"
)

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

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func HealthHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func LoginHandler(store *UserStore) HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		if err := ctx.Err(); err != nil {
			return err
		}

		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, errorResponse{Code: 400, Message: "invalid request body"})
			return nil
		}

		if req.Email == "" || req.Password == "" {
			writeJSON(w, http.StatusBadRequest, errorResponse{Code: 400, Message: "email and password are required"})
			return nil
		}

		user, err := store.Authenticate(req.Email, req.Password)
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, errorResponse{Code: 401, Message: err.Error()})
			return nil
		}

		token, err := IssueToken(user.ID, user.Username)
		if err != nil {
			return err
		}

		writeJSON(w, http.StatusOK, AuthResponse{
			Token:    token,
			UserID:   user.ID,
			Username: user.Username,
		})
		return nil
	}
}

func SignupHandler(store *UserStore) HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		if err := ctx.Err(); err != nil {
			return err
		}

		var req SignupRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, errorResponse{Code: 400, Message: "invalid request body"})
			return nil
		}

		if req.Email == "" || req.Username == "" || req.Password == "" {
			writeJSON(w, http.StatusBadRequest, errorResponse{Code: 400, Message: "email, username, and password are required"})
			return nil
		}

		user, err := store.Create(req.Username, req.Email, req.Password)
		if err != nil {
			if err == ErrUserExists {
				writeJSON(w, http.StatusConflict, errorResponse{Code: 409, Message: err.Error()})
				return nil
			}
			return err
		}

		token, err := IssueToken(user.ID, user.Username)
		if err != nil {
			return err
		}

		writeJSON(w, http.StatusOK, AuthResponse{
			Token:    token,
			UserID:   user.ID,
			Username: user.Username,
		})
		return nil
	}
}
