# -------------------- including environment variable -------------------- #
include .env

# -------------------- define environment variable ----------------------- #
CONTAINER_NAME=$(docker ps -aq --filter name=IMAGE_NAME)

# -------------------- define command target ----------------------------- #
up: build run

build:
	@docker build --build-arg builder=builder-${IMAGE_NAME} -f Dockerfile -t "$(IMAGE_NAME)" .
	@docker image prune --filter label=buildername=builder-${IMAGE_NAME} --force

run: destroy
	@docker run -d --name "$(IMAGE_NAME)" -p $(PORT):$(PORT) $(IMAGE_NAME)

destroy:
	@if [ -n "$(CONTAINER_NAME)" ]; then docker rm "$(IMAGE_NAME)" --force; fi;

test:
	@clear
	@go test --race -v ./...

test-cover:
	@clear
	@go test ./... -coverprofile .coverage
	@go tool cover -html=.coverage
	
lint:
	@clear
	@golint ./...
	@go fmt ./...

local: lint
	@clear
	@go run .