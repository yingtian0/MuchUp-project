package authzservice

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var secret_key = []byte("jwt_scret_key")

type CustomCraim struct {
	UserID string
	RoomID string 
	jwt.RegisteredClaims
}

func  JWTVerification(r http.Request) (*CustomCraim,error){

	auth := r.Header.Get("Authorization")
	if auth == "" {
		return nil,fmt.Errorf("Authorization is empty")
	}
	parts := strings.Split(auth," ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, fmt.Errorf("JWT format is invalid")
	}

	tokenString := parts[1]

	token,err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{},error) {
		return secret_key,nil
	})
	if err != nil || !(token.Valid){
		return  nil, fmt.Errorf("faild to parse token error: %w",err)
	} else if claims,ok := token.Claims.(CustomCraim); ok {
		return &claims,nil
	}

	return &CustomCraim{} ,nil

}

