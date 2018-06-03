package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []User
	//since we're passing a pointer to users, db.Find assigns array to the address
	DB.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	DB.Where("id = ?", params["id"]).Find(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	user.Email = r.FormValue("email")
	user.Name = r.FormValue("name")
	//get password hash
	user.Hash = user.hashPassword(r.FormValue("password"))
	DB.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	DB.Where("email = ?", r.FormValue("email")).Find(&user)
	w.Header().Set("Content-Type", "application/json")
	if user.checkPassword(r.FormValue("password")) {
		json.NewEncoder(w).Encode(&user)
	} else {
		err := NewJSONError("Password incorrect")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&err)
	}
}