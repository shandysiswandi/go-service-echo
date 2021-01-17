.PHONY: up build run destroy test test-cover lint local

include .env

up: build run

build :
	@docker build -f Dockerfile -t "$(IMAGE_NAME)" .

run : destroy
	@docker run -d --name "$(IMAGE_NAME)" -p $(PORT):$(PORT) $(IMAGE_NAME)

destroy :
	@docker rm "$(IMAGE_NAME)" --force
# ------------------------------------------------------------ #
test:
	@clear
	@go test --race -v ./...

test-cover:
	@clear
	@go test ./... -coverprofile .coverage
	@go tool cover -html=.coverage
# ------------------------------------------------------------ #
lint:
	@clear
	@golint ./...
	@go fmt ./...

local: lint
	@clear
	@go run .
# ------------------------------------------------------------ #