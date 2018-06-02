package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {

	//init router
	router := mux.NewRouter()

	//routes
	router.HandleFunc("/users", UsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", UserHandler).Methods("GET")
	router.HandleFunc("/users", CreateUserHandler).Methods("POST")

	return router
}