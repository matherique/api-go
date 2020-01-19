package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

	port = fmt.Sprintf(":%v", port)

	fmt.Printf("Running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, GetRouter()))

}
