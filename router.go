package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {

	//init router
	router := mux.NewRouter()

	for _, r := range routes {
		router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(r.HandlerFunc)
	}

	return router
}