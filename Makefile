# -------------------- including environment variable -------------------- #
include .env

# -------------------- define environment variable ----------------------- #
NAME=$(DOCKER_IMAGE_NAME)
VERSION=$(shell git describe --tags --always)

CONTAINER_NAME=$(shell docker ps -aq --filter name=${NAME})
IMAGE_NAME=$(shell docker ps -aq --filter name=${NAME})

# -------------------- define command target ----------------------------- #
up: lint build run

build: clean
	@docker build --build-arg TAGGED=builder-${NAME} --file Dockerfile --tag $(NAME):$(VERSION) .
	@docker image prune --filter label=tagged=builder-${NAME} --force

run: destroy-container
	@docker run --detach --name $(NAME) -p $(APP_PORT):$(APP_PORT) $(NAME):$(VERSION)

clean:
	@echo "delete container if exist --force"
	@echo "delete image if exist --force"

destroy-image:
	@if [ -n "$(NAME)" ]; then docker image rm $(NAME) --force; fi;

destroy-container:
	@if [ -n "$(CONTAINER_NAME)" ]; then docker rm $(NAME) --force; fi;

cert:
	@openssl genrsa -out ./resource/key/private.pem 4096
	@openssl rsa -in ./resource/key/private.pem -pubout -out ./resource/key/public.pem

test: lint
	@clear
	@go test -timeout 90s --race -v ./...

test-cover: lint
	@clear
	@go test -timeout 90s ./... -coverprofile .coverage
	@go tool cover -html=.coverage

lint:
	@clear
	@golint ./...
	@go fmt ./...

dev: lint
	@clear
	@reflex -r '\.go' -s -- sh -c "go run ."

start: lint
	@clear
	@go run .
