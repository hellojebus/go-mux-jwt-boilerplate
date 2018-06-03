# GoLang Mux, Gorm, JWT REST API Boilerplate 

The purpose of this web app to learn Go REST API Development using Mux (router), Gorm (ORM), JSON Web Tokens (JWT), and Golang development in general. The application is optimized for Heroku deployment using go Go Build Kit and Godep for dependency management.

**This is for educational purposes only and probably unsuitable for production** 

## What's included

Basic CRUD* routes for user management

* Show Users `GET /users`
* Show User `GET /users/{userId}`
* Add User `POST /users`
* User Login `POST /users/login`

Several routes are protected and require JWT tokens, which can be generated using the login route.
You will need to create a user by sending a post request to the addUser route.

## Configuration

Make sure to copy `.env.sample` to `.env` and add correct config details

## Installing and running

To run using Gin: `gin run main.go`

To run using Go: `go run *.go`

Gin is my preferred method for running the application locally as it will watch your files for changes and automatically re-run the application.
  
## Todos
 
[]Add testing <br>
[]Add Delete User route <br>
[]Add Home route <br>
[]Add exp to JWT <br>