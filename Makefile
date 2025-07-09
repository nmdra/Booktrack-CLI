.PHONY: generate run test

generate:
	sqlc generate

run:
	go run main.go

test:
	go test ./...

build:
	go build -o bin/booktrack-cli main.go
