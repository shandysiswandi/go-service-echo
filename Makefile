.PHONY: build run destroy up

# Command Server
up: build run

build :
	@echo "Hellow"

run : destroy
	@echo "Hellow"

destroy :
	@echo "Hellow"

# Command Local
local:
	@clear
	@go run main.go