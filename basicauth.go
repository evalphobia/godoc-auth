package main

import (
	"net/http"
	"os"

	"github.com/abbot/go-http-auth"
)

func BasicAuthHandler(h http.Handler) http.Handler {
	authenticator := auth.NewBasicAuthenticator("godoc", Secret)
	return http.HandlerFunc(authenticator.Wrap(func(w http.ResponseWriter, req *auth.AuthenticatedRequest) {
		h.ServeHTTP(w, &req.Request)
	}))
}

func Secret(user, realm string) string {
	name := os.Getenv("BASIC_AUTH_USERNAME")
	pass := os.Getenv("BASIC_AUTH_PASSWORD")
	if name != "" && pass != "" && user == name {
		return pass
	}
	return ""
}
