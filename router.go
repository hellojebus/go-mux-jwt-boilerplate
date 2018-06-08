package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/hellojebus/go-envoz-api/user"
	customRouter "github.com/hellojebus/go-envoz-api/router"
	"github.com/hellojebus/go-envoz-api/middleware"
)

func NewRouter() *mux.Router {

	//init router
	router := mux.NewRouter()

	//append user routes
	customRouter.AppRoutes = append(customRouter.AppRoutes, user.Routes)

	for _, route := range customRouter.AppRoutes {

		//create subroute
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		//loop through each sub route
		for _, r := range route.SubRoutes {

			var handler http.Handler
			handler = r.HandlerFunc

			//check to see if route should be protected with jwt
			if r.Protected {
				handler = middleware.JWTMiddleware(r.HandlerFunc)
			}

			//attach sub route
			routePrefix.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method).
				Name(r.Name)
		}

	}

	return router
}