package main

type User struct {
	ID   int `gorm:"primary_key" json:"id"`
	Email string `gorm:"unique_index" json:"email"`
	Name string `json:"name"`
}


