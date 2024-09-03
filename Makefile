.PHONY: build-all build run docker-image dev clean clean-all

loadtest:
	@k6 run --vus 10 --duration 30s loadtesting.js

build:
	@go build -o bin/server cmd/main.go

run: build
	@./bin/server

docker-image:
	@docker build --tag api .

build-all: build docker-image
	@./bin/server

clean:
	@rm -rf bin

clean-all: clean
	@docker rmi api
