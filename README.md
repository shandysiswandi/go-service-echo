# Go REST Echo
is

## Project Structure
    .
    ├── app
    ├── config
    │   ├── .env.test
    │   ├── config_test.go
    │   └── config.go
    ├── db
    │   ├── db_test.go
    │   └── db.go
    ├── external
    │   └── .gitkeep
    ├── internal
    │   ├── users
    │   └── tasks
    ├── resource
    │   ├── ...
    │   └── .gitkeep
    ├── service
    │   └── .gitkeep
    ├── util
    │   └── .gitkeep
    ├── .dockerignore
    ├── .env.example
    ├── .gitignore
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── Makefile
    ├── mysql.sql
    ├── README.md
    ├── requests.http
    └── serverless.yml

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
