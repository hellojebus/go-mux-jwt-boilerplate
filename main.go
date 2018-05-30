package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"os"
	"github.com/jpfuentes2/go-env"
	"path"
	"strings"
)

func main(){
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	pwd, _ := os.Getwd()
	if strings.Contains(r.Host, "localhost") {
		env.ReadEnv(path.Join(pwd, ".env"))
		for _, v := range os.Environ() {
			w.Write([]byte("env: " + v))
		}
	} else {
		w.Write([]byte("host: " + r.Host))
	}

}