.PHONY: build run destroy up

# Command Server
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

# Command Local
local:
	@clear
	@go run main.go