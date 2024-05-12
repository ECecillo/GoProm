.PHONY: build run docker-image

build:
	@go build -o bin/server cmd/main.go

run: build
	@./bin/server

docker-image:
	@docker build --tag api .