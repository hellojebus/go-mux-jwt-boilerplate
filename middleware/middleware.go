package middleware

import (
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	customHTTP "github.com/hellojebus/go-envoz-api/http"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Authentication failure")
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := VerifyToken(tokenString)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error verifying JWT token: " + err.Error())
			return
		}

		//pass userId claim to req
		//todo: find a better way to convert the claim to string
		userId := strconv.FormatFloat(claims.(jwt.MapClaims)["user_id"].(float64), 'g', 1, 64)
		r.Header.Set("userId", userId)
		next.ServeHTTP(w, r)
	})
}