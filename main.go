package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"os"
)

func main(){
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	w.Write([]byte("Hostname: "+ name))
}