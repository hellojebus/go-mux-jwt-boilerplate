package main

import (
	"net/http"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type ErrorResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
}

/*
HTTP Response handling for errors,
Returns valid JSON with error type and response code
*/

func NewErrorResponse(w http.ResponseWriter, statusCode int, response string){
	error := ErrorResponse{
		true,
		response,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&error)
	return
}


/*
Verify JWT token based on environment var (JWT_SECRET)
*/

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}