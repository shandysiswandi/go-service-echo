# -------------------- including environment variable -------------------- #
include .env

# -------------------- define environment variable ----------------------- #
CONTAINER_NAME=$(docker ps -aq --filter name=IMAGE_NAME)

# -------------------- define command target ----------------------------- #
up: lint build run

build:
	@docker build --build-arg TAGGED=builder-${IMAGE_NAME} -f Dockerfile -t "$(IMAGE_NAME)" .
	@docker image prune --filter label=tagged=builder-${IMAGE_NAME} --force

run: destroy
	@docker run -d --name "$(IMAGE_NAME)" -p $(PORT):$(PORT) $(IMAGE_NAME)

destroy:
	@if [ -n "$(CONTAINER_NAME)" ]; then docker rm "$(IMAGE_NAME)" --force; fi;

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