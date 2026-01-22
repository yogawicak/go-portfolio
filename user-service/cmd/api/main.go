package main

import (
	"fmt"
	"log"
	"net/http"
)

type AppHandler = func(http.ResponseWriter, *http.Request)

// Struct to hold routes
type Route struct {
	Method  string
	Pattern string
	Handler AppHandler
}

type Router struct {
	Routes []Route
}

func (r *Router) AddRouter(method string, pattern string, handler AppHandler) {

}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Register Routes
	for _, route := range r.Routes {

		// Check matching HTTP Method
		if req.Method != route.Method {
			continue
		}

		route.Handler(w, req)
	}
}

func main() {
	router := &Router{}

	// router.Routes.

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
