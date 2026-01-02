package main

import (
	"log"
	"net/http"

	authz "muchup.com/authz"
)

func main () {

    mux := http.NewServeMux()

    mux.Handle("/authz",authz.HandlerFunc(authz.AuthorizationHandler))

    server := &http.Server{
	Addr: ":8099",
	Handler: mux,

    }
    log.Fatal(server.ListenAndServe())

}