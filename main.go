package main

import (
	"os"
	"log"
	"net/http"
	"github.com/hellojebus/go-envoz-api/db"
)

func main(){

	//init router
	port := os.Getenv("PORT")
	router := NewRouter()

	//Setup database
	db.DB = db.SetupDB()
	defer db.DB.Close()

	//create http server
	log.Fatal(http.ListenAndServe(":"+port, router))
}