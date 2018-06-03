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
		"UsersIndex",
		"GET",
		"/users",
		UsersIndexHandler,
		true,
	},
	Route{
		"UsersShow",
		"GET",
		"/users/{userId}",
		UsersShowHandler,
		true,
	},
	Route{
		"UsersCreate",
		"POST",
		"/users",
		UsersCreateHandler,
		false,
	},
	Route{
		"UsersLogin",
		"POST",
		"/users/login",
		UsersLoginHandler,
		false,
	},
	Route{
		"UsersDelete",
		"DELETE",
		"/users/{userId}",
		UsersDelete,
		false,
	},
}