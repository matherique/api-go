package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := fmt.Sprintf(":%d", 8000)

	fmt.Printf("running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, GetRouter()))

}
