package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home handler")
}