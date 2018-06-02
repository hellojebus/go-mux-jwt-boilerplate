package main

import (
	"os"
	"log"
	"net/http"
)

func main(){

	//init router
	port := os.Getenv("PORT")
	router := NewRouter()

	//Setup database
	DB = SetupDB()
	defer DB.Close()

	//create http server
	log.Fatal(http.ListenAndServe(":"+port, router))
}