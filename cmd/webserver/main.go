package main // ONLY main.go here!

import (
	"log"

	"github.com/yourusername/web-server/internal/server"
)

func main() {
	srv := server.New()
	log.Printf("Server starting on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
