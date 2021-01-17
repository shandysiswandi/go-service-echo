.PHONY: up build run destroy test test-cover lint local

up: build run

build :
	@echo "build"

run : destroy
	@echo "run"

destroy :
	@echo "destroy"

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

local:
	@clear
	@go run .