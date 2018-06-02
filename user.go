package main

type User struct {
	ID   int `gorm:"primary_key" json:"id"`
	Email string `gorm:"unique_index" json:"email"`
	Name string `json:"name"`
}

func (u User) hashPassword(password string) string {
	return password
}

func (u User) checkPassword(password string) string {
	return password
}