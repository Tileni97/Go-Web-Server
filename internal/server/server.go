package server

import (
	"net/http"
	"os"
	"web-server/internal/handlers"
)

func New() *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &http.Server{
		Addr:    ":" + port,
		Handler: handlers.NewRouter(),
	}
}
