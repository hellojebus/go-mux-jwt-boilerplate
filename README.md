# GoLang Mux, Gorm, JWT REST API Boilerplate 

The purpose of this web app is for me to learn Go REST API Development using Mux (router), Gorm (ORM), JSON Web Tokens (JWT), and Golang development in general. The application is optimized for Heroku deployment using the Go Build Kit and Godep for dependency management.

**This is for educational purposes only and probably unsuitable for production** 

## What's included

Basic CRUD* routes for user management _*Missing the U in CRUD ;)_ 

* Show Users `GET /users`
* Show User `GET /users/{userId}`
* Create User `POST /users`
* User Login `POST /users/login`
* Delete User `DELETE /users/{userId}`
* Update User `PUT /users/{userId}` * Note only the user can update their own name

Several routes are protected and require JWT tokens, which can be generated using the login route.
You will need to create a user by sending a post request to the createUser route.

## Configuration

Make sure to copy `.env.sample` to `.env` and update all fields (DB fields are important!)

**Please note that this is using the MySQL driver, if you prefer to use another driver, update `db.go` accordingly**

Gorm is setup to automigrate the Users table, so it should be plug and play.

## Installing and running

Make sure to have all required external deps. Look at Godeps config file to view them all.

To run using Gin: `gin run main.go`

To run using Go: `go run *.go`

Gin is my preferred method for running the application locally as it will watch your files for changes and automatically re-run the application.

To view application view `localhost:3000 or locahost:YOUR_PORT_ENV` in browser
  
## Todos
 
[] Add testing <br>
[] Add Home route <br>
[] Add User Role to illustrate how relationships work with Gorm<br>