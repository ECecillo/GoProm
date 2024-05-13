.PHONY: build-all build run docker-image dev

build:
	@go build -o bin/server cmd/main.go

run: build
	@./bin/server

docker-image:
	@docker build --tag api .

build-all: build docker-image
	@./bin/server
