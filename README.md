# GoLang Mux, Gorm, JWT REST API Boilerplate 

The purpose of this web app is for me to learn Go REST API Development using Mux (router), Gorm (ORM), JSON Web Tokens (JWT), and Golang development in general. The application is optimized for Heroku deployment using the Go Build Kit and Godep for dependency management.

**This is for educational purposes only and probably unsuitable for production** 

## What's included

Basic CRUD routes for user management  

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

## Installation

Make sure to have all required external deps. Look at Godeps config file to view them all.

**Preferred Method, Live Reloading (optional):**

Install Gin `go get github.com/codegangsta/gin`

Then run: `gin run main.go`

**Otherwise:**

To run using Go: `go run *.go`

To view application in browser: `localhost:3000 (gin run) or locahost:YOUR_PORT_ENV (go run)`
  
## Todos
 
[] Detach userHandlers DB functions from http for testing purposes<br>
[] Add User Role to illustrate how relationships work with Gorm <br>