.PHONY: run test build tidy fmt lint

run:
	go run ./app/cmd

test:
	go test ./...

build:
	go build ./...

tidy:
	go mod tidy

fmt:
	go fmt ./...

lint:
	go vet ./...
