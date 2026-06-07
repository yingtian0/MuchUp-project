package main

import (
	"log"
	"net/http"
	"time"

	"MuchUp/authz"
)

func main() {
	store := authz.NewUserStore()

	mux := http.NewServeMux()
	mux.Handle("/v1/auth/login", authz.LoginHandler(store))
	mux.Handle("/v1/auth/signup", authz.SignupHandler(store))
	mux.Handle("/healthz", authz.HandlerFunc(authz.HealthHandler))

	server := &http.Server{
		Addr:              ":8099",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
