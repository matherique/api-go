package api

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/api/user", userRoute)
}
