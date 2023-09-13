# LogApiv0


LogApiv0 is a log management application. This application allows you to create, view, and manage logs.

## Getting Started

In this section, provide information on how to get your project up and running in a local environment.

### Requirements

To run the project, you'll need the following:

- [Go](https://golang.org/) (at least version 1.21.0)
- [Fiber](https://github.com/gofiber/fiber) web framework
- [MongoDB](https://github.com/mongodb/mongo-go-driver) Nosql Databases
- [Jwt](https://github.com/golang-jwt/jwt) JSON Web Tokens.
  

### Installation

To start the project, follow these steps:

- Replace the connection key in the database/dbconnect.go file with your own mongodb connection key
- Change the SECRET_KEY and ADMIN_KEY in the .env file

-Navigate to the project directory:

   ```bash
   1. cd LogApiv0
   2. Install project dependencies by running: go get {fiber,mongo,jwt}
   3. Start the application: go run cmd/main.go
   ```

## Usage
This application provides a RESTful API for log management and user management. Below are explanations of API endpoints and how to use them:

Admin Registration (Register)
HTTP Method: POST

This endpoint is only available to administrative users. Before using it, you need to get the jwt token returned from the login endpoint as admin. It is used to register new users. The submitted data must include new user information. Example JSON payload:
```json
{
    "jwt":       "admin-jwt-key"
    "username":  "new_user",
    "password":  "password123"
}
```
## User Login (Login)
HTTP Method: POST

This endpoint allows users to log in. The data sent should contain user credentials and is used for the login process.Returns jwt token as body.Example JSON payload:
```json
{
    "username": "user_name",
    "password": "password123"
}
```

## Log Entry (Log)
HTTP Method: POST

This endpoint creates log entries. Only authenticated users can perform this action. So, from this action, you need to log in and have jwt token. As a result of this action, a folder is opened with your username obtained from the jwt token you entered, a file in the form of unixtime.txt is created and the data received as a request is stored in this file. The data sent must include log information. Example JSON payload:
```json
{
    "jwt": "user_jwt",
    "data": "Log data goes here."
}
```


## User Deletion (Delete User)
HTTP Method: DELETE

This endpoint is available for admin users only. To delete a user, the data sent should include an admin JWT and the username to be deleted. Example JSON payload:
```json
{
    "jwt": "admin_jwt",
    "username": "user_to_delete"
}
```
## User Update (Update User)
HTTP Method: PUT

This endpoint allows users to update their passwords. The data sent should include the user JWT, the old password, and the new password. Example JSON payload:
```json
{
    "jwt": "user_jwt",
    "oldpwd": "old_password",
    "newpwd": "new_password"
}
```
