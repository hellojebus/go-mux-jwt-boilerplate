package main

import (
	"net/http"
	"fmt"
)

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("just passing through")
		next.ServeHTTP(w, r)
	})
}