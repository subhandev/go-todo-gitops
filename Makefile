.PHONY: run build tidy

run:
	go run ./app/cmd

build:
	go build ./...

tidy:
	go mod tidy