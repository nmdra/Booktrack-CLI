.PHONY: generate run test

generate:
	sqlc generate

schema:
	cat db/migrations/*.up.sql > internal/db/schema.sql

run:
	go run main.go

test:
	go test ./...

build:
	go build -o bin/booktrack-cli main.go
