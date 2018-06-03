package main

import (
	"net/http"
)

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			NewErrorResponse(w, http.StatusUnauthorized, "Authentication failure")
			return
		}
		next.ServeHTTP(w, r)
	})
}