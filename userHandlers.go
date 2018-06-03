package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func UsersIndexHandler(w http.ResponseWriter, r *http.Request) {
	var users []User
	//since we're passing a pointer to users, db.Find assigns array to the address
	DB.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func UsersShowHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["userId"])
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
		token, err := user.generateJWT()
		if err != nil {
			NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
			return
		}
		json.NewEncoder(w).Encode(&token)
	} else {
		NewErrorResponse(w, http.StatusUnauthorized, "Password incorrect")
		return
	}
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	var users []User

	DB.First(&user, params["userId"])
	DB.Delete(&user)

	DB.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}