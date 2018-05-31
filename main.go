package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"os"
	"github.com/jpfuentes2/go-env"
	"path"
)

func main(){
	port := os.Getenv("PORT")
	router := mux.NewRouter()

	//get env vars
	pwd, _ := os.Getwd()
	env.ReadEnv(path.Join(pwd, ".env"))

	router.HandleFunc("/", HomeHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	//load up the env package only in local env
	for _, v := range os.Environ() {
		w.Write([]byte("env: " + v))
	}

}