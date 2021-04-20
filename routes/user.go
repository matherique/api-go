package routes

import "net/http"

type userRoutes struct{}

func (u userRoutes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET /user"))
}
