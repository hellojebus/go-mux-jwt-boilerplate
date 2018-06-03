package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {

	//init router
	router := mux.NewRouter()

	for _, r := range routes {
		var handler http.Handler

		handler = r.HandlerFunc
		if r.Protected {
			handler = jwtMiddleware(r.HandlerFunc)
		}

		router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(handler)

	}

	return router
}