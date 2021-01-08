# Go REST Echo
is

## Project Structure
    .
    ├── app
    ├── database
    │   ├── mongo
    │   ├── mysql
    │   └── postgres
    ├── domain
    │   ├── task
    |   |   ├── delivery
    |   |   ├── repository
    |   |   └── usecase
    │   └── user
    |   |   ├── delivery
    |   |   ├── repository
    |   |   └── usecase
    ├── entity
    │   ├── base
    │   |    ├── increment.go
    │   |    ├── timestamp.go
    │   |    └── user.go
    │   ├── task.go
    │   └── user.go
    ├── helper
    │   ├── bcrypt.go
    │   └── env.go
    ├── test
    │   ├── .env.test
    │   └── ...
    ├── .dockerignore
    ├── .env.example
    ├── .gitignore
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── Makefile
    └── README.md

## Route List