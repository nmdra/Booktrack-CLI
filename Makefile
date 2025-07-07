.PHONY: generate run test

generate:
	docker run --rm -v .:/src -w /src sqlc/sqlc generate

run:
	go run main.go

test:
	go test ./...

build:
	go build -o bin/Bookstore-CLI main.go
