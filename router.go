package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {

	//init router
	router := mux.NewRouter()

	for _, r := range routes {
		if r.Protected {
			router.
				Methods(r.Method).
				Path(r.Pattern).
				Name(r.Name).
				Handler(jwtMiddleware(r.HandlerFunc))
		} else {
			router.
				Methods(r.Method).
				Path(r.Pattern).
				Name(r.Name).
				Handler(r.HandlerFunc)
		}

	}

	return router
}