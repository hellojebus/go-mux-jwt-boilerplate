package main

import "net/http"

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
	Protected bool
}

type Routes []Route

var routes = Routes{
	Route{
		"UserIndex",
		"GET",
		"/users",
		UsersHandler,
		true,
	},
	Route{
		"UserShow",
		"GET",
		"/users/{userId}",
		UserHandler,
		false,
	},
	Route{
		"UserCreate",
		"POST",
		"/users",
		UserCreateHandler,
		false,
	},
	Route{
		"UserLogin",
		"POST",
		"/users/login",
		UserLoginHandler,
		false,
	},
}