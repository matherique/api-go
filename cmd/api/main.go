package main

import (
	"log"

	"github.com/matherique/api-go/cmd/server"
)

func main() {
	srv := server.NewServer()

	if err := srv.Listen(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
