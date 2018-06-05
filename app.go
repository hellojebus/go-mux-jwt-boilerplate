package main

import (
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
)

type App struct {
	DB *gorm.DB
	Router *mux.Router
}

func (*App) init(dbHost, dbName, dbUser, dbPassword string){

}
