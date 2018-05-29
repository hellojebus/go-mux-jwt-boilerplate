package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home handler"))
}