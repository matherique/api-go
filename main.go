package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/matherique/api-golang/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = string(8000)
	}

	port = fmt.Sprintf("%v", port)

	router := mux.NewRouter()

	handler.Home(router)
	handler.User(router)

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%s", "localhost", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Running on http://localhost:%s\n", port)
	log.Fatal(srv.ListenAndServe())

}
