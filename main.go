package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"os"
	"github.com/jpfuentes2/go-env"
	"path"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type User struct {
	ID   int
	Email string
}

//connect to db
var dbHost string = os.Getenv("DB_HOST")
var dbName string = os.Getenv("DB_NAME")
var dbUser string = os.Getenv("DB_USER")
var dbPassword string = os.Getenv("DB_PASSWORD")

func main(){

	//read env
	pwd, _ := os.Getwd()
	env.ReadEnv(path.Join(pwd, ".env"))


	db, err := gorm.Open("mysql", dbUser+":"+ dbPassword +"@tcp(" + dbHost+ ":3306)/"+ dbName + "?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	//init router
	port := os.Getenv("PORT")
	router := mux.NewRouter()

	if err != nil {
		fmt.Println(err)
	}

	router.HandleFunc("/", HomeHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open("mysql", dbUser+":"+ dbPassword +"@tcp(" + dbHost+ ":3306)/"+ dbName + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(w, db.First(1))
	defer db.Close()

	r.Body.Close()

}