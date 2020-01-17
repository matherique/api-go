package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matherique/api-golang/controllers"
)

// Router = All routers

type routerInfo struct {
	path    string
	handler func(w http.ResponseWriter, r *http.Request)
	met     []string
}

var routerList = []routerInfo{
	{path: "/", handler: controllers.HomeIndex, met: []string{"GET"}},
	{path: "/users", handler: controllers.UserIndex, met: []string{"GET"}},
	{path: "/users", handler: controllers.UserStore, met: []string{"POST"}},
}

// Loggin function
func withLogging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s", r.URL, r.Method)
		f(w, r)
	}
}

// GetRouter return Router pointer with all routers based in routerList
func GetRouter() *mux.Router {
	var router = mux.NewRouter()
	for _, ri := range routerList {
		router.HandleFunc(ri.path, withLogging(ri.handler)).Methods(ri.met...)
	}
	return router
}
