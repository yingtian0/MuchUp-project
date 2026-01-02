package authzservice

import (
	"context"
	"net/http"
)

type payLoad struct {
	Sub string 
	Email string
}

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if err := h(r.Context(), w, r); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


func  AuthorizationHandler (ctx context.Context,w http.ResponseWriter,r *http.Request,) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	claims,err := JWTVerification(*r)
	if err != nil {
		return err 
	}

	if claims.UserID != "" {
		w.WriteHeader(http.StatusUnauthorized)
		return nil 
	}
	if claims.RoomID != "" {
		w.WriteHeader(http.StatusForbidden)
	}

	w.WriteHeader(http.StatusOK)
	return nil 
}