package main

import (
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
)

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			NewErrorResponse(w, http.StatusUnauthorized, "Authentication failure")
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString)
		if err != nil {
			NewErrorResponse(w, http.StatusUnauthorized, "Error verifying JWT token: " + err.Error())
			return
		}
		userId := claims.(jwt.MapClaims)["user_id"].(string)
		r.Header.Set("userId", userId)
		next.ServeHTTP(w, r)
	})
}