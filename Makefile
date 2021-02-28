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

DOCKER_HUB_REPO			= shandysiswandi/go-service

# -------------------- define command target for docker ------------------------------------------ #
up: lint clean docker-build docker-push docker-run

clean: docker-remove-container docker-remove-image

docker-build:
	@docker build --build-arg TAGGED=builder-${DOCKER_IMAGE_NAME} --file Dockerfile --tag $(DOCKER_IMAGE_NAME) .
	# @docker image prune --filter label=tagged=builder-${DOCKER_IMAGE_NAME} --force

docker-push:
	@docker tag $(DOCKER_IMAGE_NAME) $(DOCKER_HUB_REPO):latest
	@docker tag $(DOCKER_IMAGE_NAME) $(DOCKER_HUB_REPO):$(VERSION)
	@docker push $(DOCKER_HUB_REPO)

docker-run:
	@docker run --rm --detach --name $(DOCKER_CONTAINER_NAME) -p $(PORT):$(PORT) $(DOCKER_HUB_REPO)

docker-remove-container:
	@if [ -n "$(CONTAINER_NAME_EXIST)" ]; then docker rm $(CONTAINER_NAME_EXIST) --force; fi;

docker-remove-image:
	@if [ -n "$(IMAGE_NAME_EXIST)" ]; then docker image rm $(IMAGE_NAME_EXIST) --force; fi;

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
	@go fmt ./...

dev: lint
	@clear
	@reflex -r '\.go' -s -- sh -c "go run main.go"

start: lint
	@clear
	@go run main.go

cli: lint
	@clear
	@go run cmd/main.go
# -------------------------------------------------------------------------------------------------- #
