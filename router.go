package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {

	//init router
	router := mux.NewRouter()

	AppRoutes = append(AppRoutes, userRoutes)

	for _, route := range AppRoutes {

		subRoute := router.PathPrefix(route.Prefix).Subrouter()

		for _, r := range route.SubRoutes {
			var handler http.Handler
			handler = r.HandlerFunc

			//check to see if route should be protected with jwt
			if r.Protected {
				handler = jwtMiddleware(r.HandlerFunc)
			}

			subRoute.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method).
				Name(r.Name)
		}

	}

	return router
}