package main

import (
	"log"
)

func main() {
	server := newServer()

	if err := server.Listen(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
