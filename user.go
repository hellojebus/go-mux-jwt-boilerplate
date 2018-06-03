package main

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID   int `gorm:"primary_key" json:"id"`
	Email string `gorm:"unique_index" json:"email"`
	Name string `json:"name"`
	Hash string `json:"-"` //hides from any json marshalling output
}

func (u User) hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}

func (u User) checkPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password))
	return err == nil
}

func (u User) generateJWT() (string, error) {
	signingKey := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.ID,
		"name": u.Name,
		"email": u.Email,
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}