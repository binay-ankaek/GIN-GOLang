# User Contacts Management System

This is a user contacts management system built using the Gin framework, GORM, and hexagonal architecture.

## Features

1. User registration and login using phone and password.
2. User can manage their own profile.
3. User can add phone numbers to their contact list.
4. User can list, update, and delete contacts.
5. User can search their own contacts.
6. If a contact number is also registered, the user can see the profile of that contact.

## Getting Started

### Prerequisites

- Go 1.18+
- PostgreSQL (or any other database supported by GORM)
- Docker (optional, for database setup)

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/ankaEK/contact-manager.git
   cd helloapp


2.How to run:
 1. Install dependencies:
    go mod tidy
    go get -u github.com/gin-gonic/gin
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/postgres
    go get -u github.com/spf13/viper
    go get -u github.com/dgrijalva/jwt-go

 2.run main.go which is inside cmd/helloapp/ and type command go run cmd/helloapp/main.go
  or 
   cd cmd
   cd helloapp
   go run main.go
 3. Open http://localhost:3000 in your browser but by default it run on port:8080.

## frontend
how to run frontend
1. cd vue_frontend
2. cd usermgmt
3. yarn install
4. yarn serve

5.Open http://localhost:8080 in your browser

##API route for frontend
1. http://localhost:8080/register (for both register and login)
2. http://localhost:8080/profile(for profile view and update)
3. http://localhost:8080/contacts(for contact view and update)


## Project setup
```
yarn install
```

### Compiles and hot-reloads for development
```
yarn serve
```

### Compiles and minifies for production
```
yarn build
```

### Lints and fixes files
```
yarn lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
