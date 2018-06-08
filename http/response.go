package http

import (
	"net/http"
	"encoding/json"
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