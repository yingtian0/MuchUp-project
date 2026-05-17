package main

import (
	"log"
	"net/http"

	authzservice "muchup/authz"
)

func main() {
	store := authzservice.NewUserStore()

	mux := http.NewServeMux()
	mux.Handle("/v1/auth/login", authzservice.LoginHandler(store))
	mux.Handle("/v1/auth/signup", authzservice.SignupHandler(store))
	mux.Handle("/healthz", authzservice.HandlerFunc(authzservice.HealthHandler))

	server := &http.Server{
		Addr:    ":8099",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
