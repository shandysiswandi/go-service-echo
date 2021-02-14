# -------------------- including environment variable -------------------------------------------- #
include .env

# -------------------- define environment variable ----------------------------------------------- #
NAME					= $(APP_NAME)
PORT					= $(APP_PORT)
VERSION					= $(shell git describe --tags --always)

DOCKER_CONTAINER_NAME	= $(NAME)
DOCKER_IMAGE_NAME		= $(NAME):$(VERSION)

CONTAINER_NAME_EXIST	= $(shell docker ps -aq --filter name=${DOCKER_CONTAINER_NAME})
IMAGE_NAME_EXIST		= $(shell docker images -aq ${DOCKER_IMAGE_NAME})

# -------------------- define command target for docker ------------------------------------------ #
up: lint clean docker-build docker-run

clean: docker-destroy-container docker-destroy-image

docker-build:
	@docker build --build-arg TAGGED=builder-${DOCKER_IMAGE_NAME} --file Dockerfile --tag $(DOCKER_IMAGE_NAME) .
	@docker image prune --filter label=tagged=builder-${DOCKER_IMAGE_NAME} --force

docker-run:
	@docker run --detach --name $(DOCKER_CONTAINER_NAME) -p $(PORT):$(PORT) $(DOCKER_IMAGE_NAME)

docker-destroy-image:
	if [ -n "$(IMAGE_NAME_EXIST)" ]; then docker image rm $(IMAGE_NAME_EXIST) --force; fi;

docker-destroy-container:
	@if [ -n "$(CONTAINER_NAME_EXIST)" ]; then docker rm $(CONTAINER_NAME_EXIST) --force; fi;

# -------------------- define command target for docker-compose --------------------------------- #
compose: compose-down
	@clear
	@docker-compose up --build

compose-down:
	@clear
	@docker-compose down

# -------------------- define command target for generate rsa ------------------------------------ #
cert:
	@clear
	@openssl genrsa -out ./resource/key/private.pem 4096
	@openssl rsa -in ./resource/key/private.pem -pubout -out ./resource/key/public.pem

# -------------------- define command target for development & testing --------------------------- #
test: lint
	@clear
	@go test -v ./...

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
	@reflex -r '\.go' -s -- sh -c "go run main.go"

start: lint
	@clear
	@go run main.go
# -------------------------------------------------------------------------------------------------- #
