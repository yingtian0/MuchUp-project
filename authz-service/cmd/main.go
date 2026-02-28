package main

import (
	"log"
	"net/http"

	authz "muchup.com/authz"
)

func main() {
	store := authz.NewUserStore()

	mux := http.NewServeMux()
	mux.Handle("/auth/login", authz.LoginHandler(store))
	mux.Handle("/auth/signup", authz.SignupHandler(store))
	mux.Handle("/healthz", authz.HandlerFunc(authz.HealthHandler))

	server := &http.Server{
		Addr:    ":8099",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
