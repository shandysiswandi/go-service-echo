.PHONY: build run destroy up

# Command Server
up: build run

build :
	@echo "Hellow"

run : destroy
	@echo "Hellow"

destroy :
	@echo "Hellow"

test:
	@clear
	@go test --race -v ./...

# Command Local
local:
	@clear
	@go run main.go