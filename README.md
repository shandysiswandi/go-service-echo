# Go REST Echo
go started template for building micro service using framework echo.

## Project Structure
```
.
├── app
│   ├── context
│   ├── library
│   │   ├── redis
│   │   ├── sentry
│   │   └── token
│   ├── middlewares
│   ├── routes
│   └── validation
├── config
│   └── constant
├── db
├── external
│   └── jsonplaceholder
├── internal
│   ├── authentication
│   └── users
├── resource
│   ├── html
│   ├── log
│   ├── media
│   └── ssl
└── util
    ├── arrays
    ├── bcrypt
    ├── logger
    ├── numbers
    └── stringy
```

## Routes List
asas

## Commands List
1. This command will call `lint`, `build`, and `run` command
```bash
make up
```
2. This command will build this project to docker image & delete the builder image
```bash
make build
```
3. This command will call `destroy` & run container base on image
```bash
make run
```
4. This command will call `lint` & run test
```bash
make test
```
5. This command will call `lint` & run test then coverage file
```bash
make test-cover
```
6. This command will linting `*.go` and formatted
```bash
make lint
```
7. This command will call `lint` & start this app with hot reload if any change the code
```bash
make dev
```
8. This command will call `lint` & start this app
```bash
make start
```
