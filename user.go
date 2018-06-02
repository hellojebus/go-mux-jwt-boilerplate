package main

import "golang.org/x/crypto/bcrypt"

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