.PHONY: build run migration migration_down sqlc test test100 cover gen clean
.DEFAULT_GOAL := run

include .env

build:
	go build -o todo_server cmd/main.go

run: build
	./todo_server

style:
	gofmt -l .
	golint ./...

migration:
	cd ./internal/database/migrations && goose postgres postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable up

migration_down:
	cd ./internal/database/migrations && goose postgres postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable down

sqlc:
	sqlc generate

test:
	go test -v -count=1 ./...

test100:
	go test -v -count=100 ./...

cover:
	go test -short -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	rm coverage.out

gen:
	mockgen -source=internal/database/repo.go \
	-destination=internal/database/mocks/mock_repo.go

clean:
	rm -rf coverage.html
	rm -rf todo_server