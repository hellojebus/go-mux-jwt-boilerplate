package main

type JSONError struct {
	Type string `json:"type"`
	Message string `json:"message"`
}

func NewJSONError(message string) JSONError {
	error := JSONError{
		"error",
		message,
	}
	return error
}