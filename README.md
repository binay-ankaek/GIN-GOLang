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
    ```

### How to run

1. **Install dependencies:**

    ```sh
    go mod tidy
    go get -u github.com/gin-gonic/gin
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/postgres
    go get -u github.com/spf13/viper
    go get -u github.com/dgrijalva/jwt-go
    ...
    # See go.mod file for additional dependencies
    ```

2. **Run main.go:**

    ```sh
    go run cmd/helloapp/main.go
    ```

    or

    ```sh
    cd cmd/helloapp
    go run main.go
    ```

3. **Open [http://localhost:3000](http://localhost:3000) in your browser (default port is 8080).**

## API Routes

- **POST** `http://localhost:3000/api/auth/register`: Register user credentials
- **POST** `http://localhost:3000/api/auth/login`: Login user credentials
- **GET** `http://localhost:3000/api/auth/profile`: Get profile
- **PUT** `http://localhost:3000/api/auth/profile`: Update profile
- **POST** `http://localhost:3000/api/contact/add`: Add contact
- **DELETE** `http://localhost:3000/api/auth/profile`: Delete contact
- **GET** `http://localhost:3000/api/contact/`: Get all contacts
- **PUT** `http://localhost:3000/api/contact/37`: Update contact
- **DELETE** `http://localhost:3000/api/contact/27`: Delete contact
- **GET** `http://localhost:3000/api/contact/search?phone=9899999990`: Search user and contact

## Frontend

### How to run frontend

1. Navigate to the frontend directory:

    ```sh
    cd vue_frontend/usermgmt
    ```

2. Install dependencies:

    ```sh
    yarn install
    ```

3. Serve the application:

    ```sh
    yarn serve
    ```

4. Open [http://localhost:8080](http://localhost:8080) in your browser.

### API routes for frontend

1. **Register and Login:** `http://localhost:8080/register`
2. **Profile View and Update:** `http://localhost:8080/profile`
3. **Contact View and Update:** `http://localhost:8080/contacts`

## Project setup

```sh
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


