BINARY_NAME = pw

.PHONY: build run

build:
	go build -o bin/$(BINARY_NAME) .

run:
	go run .
