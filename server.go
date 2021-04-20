package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/matherique/api-go/routes"
)

const DEFAULT_PORT string = "8080"

func getPort() string {
	envPort := os.Getenv("PORT")

	if envPort != "" {
		return envPort
	}

	return DEFAULT_PORT
}

type server struct {
	mux *http.ServeMux
}

func newServer() server {
	return server{
		mux: http.NewServeMux(),
	}
}

func (s server) Listen() error {
	port := fmt.Sprintf(":%s", getPort())

	routes.LoadRoutes(s.mux)

	log.Printf("server listen on localhost%s\n", port)
	return http.ListenAndServe(port, s.mux)
}
