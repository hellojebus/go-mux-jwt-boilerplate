package main

import "net/http"

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"UserIndex",
		"GET",
		"/users",
		UsersHandler,
	},
	Route{
		"UserShow",
		"GET",
		"/users/{userId}",
		UserHandler,
	},
	Route{
		"UserCreate",
		"POST",
		"/users",
		UserCreateHandler,
	},
	Route{
		"UserLogin",
		"POST",
		"/users/login",
		UserLoginHandler,
	},
}